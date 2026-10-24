package commvault

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

// The API returns rules in the opposite order to the configuration.
const vmGroupV2TestContentJSON = `{"ruleGroups":[{"matchRule":"ALL","rules":[{"condition":"EQUALS","type":"TAG","name":"environment","value":"production","displayName":"Environment","GUID":"tag-guid","description":"Production machines"},{"condition":"EQUALS","type":"DATACENTER","name":"Australia Southeast","value":"australiasoutheast"}]}]}`

func vmGroupV2TestConfig(key, overwrite string) map[string]interface{} {
	content := map[string]interface{}{
		"rulegroups": []interface{}{map[string]interface{}{
			"matchrule": "ALL",
			"rules": []interface{}{
				map[string]interface{}{"condition": "EQUALS", "type": "DATACENTER", "name": "Australia Southeast", "value": "australiasoutheast"},
				map[string]interface{}{
					"condition": "EQUALS", "type": "TAG", "name": "environment", "value": "production",
					"displayname": "Environment", "guid": "tag-guid", "description": "Production machines",
				},
			},
		}},
	}
	if overwrite != "" {
		content["overwrite"] = overwrite
	}
	return map[string]interface{}{
		"name": "azure-vms", "hypervisor": []interface{}{map[string]interface{}{"id": 12}},
		key: []interface{}{content},
	}
}

func vmGroupV2TestServer(t *testing.T, serve http.HandlerFunc) {
	t.Helper()
	server := httptest.NewServer(serve)
	t.Cleanup(server.Close)
	t.Setenv("CV_CSIP", server.URL)
	t.Setenv("AuthToken", "test-token")
	t.Setenv("CV_LOGGING", "false")
}

// vmGroupV2GetServer serves body for GET /V4/VmGroup/42 and fails on any other request.
func vmGroupV2GetServer(t *testing.T, body string) {
	t.Helper()
	vmGroupV2TestServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet || r.URL.Path != "/V4/VmGroup/42" {
			t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, body)
	})
}

func refreshTestResource(t *testing.T, resource *schema.Resource, config map[string]interface{}) *terraform.InstanceState {
	t.Helper()
	data := schema.TestResourceDataRaw(t, resource.Schema, config)
	data.SetId("42")
	state, err := resource.RefreshWithoutUpgrade(data.State(), nil)
	if err != nil {
		t.Fatalf("refresh: %s", err)
	}
	if state == nil || state.ID != "42" {
		t.Fatalf("refresh lost resource: %#v", state)
	}
	return state
}

func TestVMGroupV2ReadRulesDiff(t *testing.T) {
	for _, key := range []string{"content", "filters"} {
		for _, tc := range []struct {
			name, regionName, regionValue string
			wantDiff                      bool
		}{
			{name: "unchanged", regionName: "Australia Southeast", regionValue: "australiasoutheast"},
			{name: "changed region", regionName: "Australia East", regionValue: "australiaeast", wantDiff: true},
			{name: "API loses friendly name and omits value", regionName: "australiasoutheast", wantDiff: true},
		} {
			t.Run(key+"/"+tc.name, func(t *testing.T) {
				regionJSON := fmt.Sprintf(`"name":%q`, tc.regionName)
				if tc.regionValue != "" {
					regionJSON += fmt.Sprintf(`,"value":%q`, tc.regionValue)
				}
				remote := strings.Replace(vmGroupV2TestContentJSON, `"name":"Australia Southeast","value":"australiasoutheast"`, regionJSON, 1)
				vmGroupV2GetServer(t, fmt.Sprintf(`{"name":"azure-vms",%q:[%s]}`, key, remote))
				resource := resourceVMGroup_V2()
				config := vmGroupV2TestConfig(key, "false")
				state := refreshTestResource(t, resource, config)
				foundRegion := false
				for attribute, value := range state.Attributes {
					if strings.HasPrefix(attribute, key+".") && strings.HasSuffix(attribute, ".type") && value == "DATACENTER" {
						foundRegion = true
						prefix := strings.TrimSuffix(attribute, "type")
						if state.Attributes[prefix+"name"] != tc.regionName || state.Attributes[prefix+"value"] != tc.regionValue {
							t.Fatalf("read masked API region: name = %q, value = %q; want %q, %q", state.Attributes[prefix+"name"], state.Attributes[prefix+"value"], tc.regionName, tc.regionValue)
						}
					}
				}
				if !foundRegion {
					t.Fatal("read lost DATACENTER rule")
				}
				diff, err := resource.Diff(state, terraform.NewResourceConfigRaw(config), nil)
				if err != nil {
					t.Fatalf("diff: %s", err)
				}
				nameRemoved, nameAdded := false, false
				valueRemoved, valueAdded := tc.regionValue == "", false
				if diff != nil {
					for attribute, change := range diff.Attributes {
						if !strings.HasPrefix(attribute, key+".") {
							continue
						}
						if !tc.wantDiff {
							t.Fatalf("unchanged rules produced a diff for %s: %#v", attribute, change)
						}
						if strings.HasSuffix(attribute, ".name") {
							nameRemoved = nameRemoved || change.Old == tc.regionName && change.Old != change.New
							nameAdded = nameAdded || change.New == "Australia Southeast" && change.Old != change.New && !change.NewComputed
						}
						if strings.HasSuffix(attribute, ".value") {
							valueRemoved = valueRemoved || change.Old == tc.regionValue && change.Old != change.New
							valueAdded = valueAdded || change.New == "australiasoutheast" && change.Old != change.New && !change.NewComputed
						}
						if change.RequiresNew {
							t.Errorf("rule change requires replacement: %s", attribute)
						}
					}
				}
				if tc.wantDiff && (!nameRemoved || !nameAdded || !valueRemoved || !valueAdded) {
					t.Fatalf("region drift not detected: name removed/added = %t/%t, value removed/added = %t/%t; diff = %#v", nameRemoved, nameAdded, valueRemoved, valueAdded, diff)
				}
			})
		}
	}
}

func TestVMGroupV2ReadRuleMetadataDiff(t *testing.T) {
	for _, key := range []string{"content", "filters"} {
		for _, field := range []struct{ config, wire string }{
			{"displayname", "displayName"}, {"guid", "GUID"}, {"description", "description"},
		} {
			for _, tc := range []struct {
				name, prior, configured, remote string
				wantDiff                        bool
			}{
				{name: "server enrichment", remote: "server metadata"},
				{name: "unconfigured absent"},
				{name: "explicit unchanged", prior: "metadata", configured: "metadata", remote: "metadata"},
				{name: "explicit remote changed", prior: "metadata", configured: "metadata", remote: "changed metadata", wantDiff: true},
				{name: "explicit remote missing", prior: "metadata", configured: "metadata", wantDiff: true},
				{name: "explicit config changed", prior: "metadata", configured: "changed metadata", remote: "metadata", wantDiff: true},
				{name: "explicit config removed", prior: "metadata", remote: "metadata", wantDiff: true},
			} {
				t.Run(key+"/"+field.config+"/"+tc.name, func(t *testing.T) {
					makeConfig := func(metadata string) map[string]interface{} {
						config := vmGroupV2TestConfig(key, "false")
						block := config[key].([]interface{})[0].(map[string]interface{})
						group := block["rulegroups"].([]interface{})[0].(map[string]interface{})
						rule := group["rules"].([]interface{})[1].(map[string]interface{})
						for _, name := range []string{"displayname", "guid", "description"} {
							delete(rule, name)
						}
						if metadata != "" {
							rule[field.config] = metadata
						}
						return config
					}
					remoteMetadata := ""
					if tc.remote != "" {
						remoteMetadata = fmt.Sprintf(",%q:%q", field.wire, tc.remote)
					}
					vmGroupV2GetServer(t, fmt.Sprintf(`{"name":"azure-vms",%q:[{"ruleGroups":[{"matchRule":"ALL","rules":[{"condition":"EQUALS","type":"TAG","name":"environment","value":"production"%s},{"condition":"EQUALS","type":"DATACENTER","name":"Australia Southeast","value":"australiasoutheast"}]}]}]}`, key, remoteMetadata))
					resource := resourceVMGroup_V2()
					state := refreshTestResource(t, resource, makeConfig(tc.prior))
					diff, err := resource.Diff(state, terraform.NewResourceConfigRaw(makeConfig(tc.configured)), nil)
					if err != nil {
						t.Fatalf("diff: %s", err)
					}
					changes := 0
					metadataRemoved, metadataAdded := tc.remote == "", tc.configured == ""
					if diff != nil {
						for attribute, change := range diff.Attributes {
							if strings.HasPrefix(attribute, key+".") {
								changes++
								if strings.Contains(attribute, ".rules.") && strings.HasSuffix(attribute, "."+field.config) && change.Old != change.New {
									metadataRemoved = metadataRemoved || change.Old == tc.remote
									metadataAdded = metadataAdded || change.New == tc.configured && !change.NewComputed
								}
								if change.RequiresNew {
									t.Errorf("metadata change requires replacement: %s", attribute)
								}
							}
						}
					}
					if (changes != 0) != tc.wantDiff || tc.wantDiff && (!metadataRemoved || !metadataAdded) {
						t.Fatalf("rule metadata diff has %d attributes (removed remote: %t, added configured: %t); wantDiff = %t", changes, metadataRemoved, metadataAdded, tc.wantDiff)
					}
				})
			}
		}
	}
}

func TestVMGroupV2RulesRoundTrip(t *testing.T) {
	for _, method := range []string{http.MethodPost, http.MethodPut} {
		for _, key := range []string{"content", "filters"} {
			if method == http.MethodPost && key == "filters" {
				continue // The v5 create API only accepts content.
			}
			for _, overwrite := range []string{"", "false", "true"} {
				t.Run(method+"/"+key+"/overwrite="+overwrite, func(t *testing.T) {
					captured := make(chan json.RawMessage, 1)
					vmGroupV2TestServer(t, func(w http.ResponseWriter, r *http.Request) {
						w.Header().Set("Content-Type", "application/json")
						if r.Method == http.MethodGet && r.URL.Path == "/V4/VmGroup/42" {
							fmt.Fprintf(w, `{"name":"azure-vms",%q:[%s]}`, key, vmGroupV2TestContentJSON)
							return
						}
						path := "/V4/VmGroup/42"
						if method == http.MethodPost {
							path = "/v5/VmGroup"
						}
						if r.Method != method || r.URL.Path != path {
							t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
							http.NotFound(w, r)
							return
						}
						var request json.RawMessage
						if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
							t.Errorf("decode request: %s", err)
							http.Error(w, err.Error(), http.StatusBadRequest)
							return
						}
						captured <- request
						fmt.Fprint(w, `{"subclientId":42,"errorCode":0}`)
					})
					resource := resourceVMGroup_V2()
					config := vmGroupV2TestConfig(key, overwrite)
					data := schema.TestResourceDataRaw(t, resource.Schema, config)
					var operation func(*schema.ResourceData, interface{}) error = resource.Create
					if method == http.MethodPut {
						data.SetId("42")
						operation = resource.Update
					}
					if err := operation(data, nil); err != nil {
						t.Fatalf("%s: %s", method, err)
					}
					if data.Id() != "42" {
						t.Fatalf("resource ID = %q; want 42", data.Id())
					}
					var raw json.RawMessage
					select {
					case raw = <-captured:
					default:
						t.Fatal("no write request captured")
					}
					var wire map[string]interface{}
					if err := json.Unmarshal(raw, &wire); err != nil {
						t.Fatal(err)
					}
					block, ok := wire[key].(map[string]interface{})
					if !ok {
						t.Fatalf("missing wire %s object: %s", key, raw)
					}
					var wantWire map[string]interface{}
					if err := json.Unmarshal([]byte(vmGroupV2TestContentJSON), &wantWire); err != nil {
						t.Fatal(err)
					}
					if overwrite != "" {
						wantWire["overwrite"] = overwrite == "true"
					}
					// Compare raw keys, including omitted fields; model decoding is case-insensitive.
					groups, ok := block["ruleGroups"].([]interface{})
					if !ok || len(groups) != 1 {
						t.Fatalf("unexpected wire ruleGroups: %s", raw)
					}
					group, ok := groups[0].(map[string]interface{})
					if !ok {
						t.Fatalf("unexpected wire rule group: %s", raw)
					}
					rules, ok := group["rules"].([]interface{})
					if !ok || len(rules) != 2 {
						t.Fatalf("unexpected wire rules: %s", raw)
					}
					foundRegion := false
					for _, rawRule := range rules {
						rule, ok := rawRule.(map[string]interface{})
						if ok && rule["type"] == "DATACENTER" {
							foundRegion = true
							if rule["name"] != "Australia Southeast" || rule["GUID"] != "australiasoutheast" || rule["value"] != nil {
								t.Errorf("%s wire %s DATACENTER must send friendly name and GUID, not value: %#v", method, key, rule)
							}
						}
					}
					if !foundRegion {
						t.Fatalf("%s wire %s missing DATACENTER rule: %s", method, key, raw)
					}
					wantRules := wantWire["ruleGroups"].([]interface{})[0].(map[string]interface{})["rules"].([]interface{})
					for _, rawRule := range wantRules {
						rule := rawRule.(map[string]interface{})
						if rule["type"] == "DATACENTER" {
							rule["GUID"] = rule["value"]
							delete(rule, "value")
						}
					}
					if reflect.DeepEqual(rules[0], wantRules[1]) {
						wantRules[0], wantRules[1] = wantRules[1], wantRules[0]
					}
					if !reflect.DeepEqual(rules, wantRules) {
						t.Errorf("wire %s rules = %#v; want %#v", key, rules, wantRules)
					}
					gotOverwrite, hasOverwrite := block["overwrite"]
					if hasOverwrite != (overwrite != "") || gotOverwrite != wantWire["overwrite"] {
						t.Errorf("wire overwrite = %#v (present: %t); want %q", gotOverwrite, hasOverwrite, overwrite)
					}
					wantData := schema.TestResourceDataRaw(t, resource.Schema, config)
					wantData.SetId("42")
					state := data.State()
					for attribute, want := range wantData.State().Attributes {
						if strings.HasPrefix(attribute, key+".") && state.Attributes[attribute] != want {
							t.Fatalf("GET %s = %q; want %q", attribute, state.Attributes[attribute], want)
						}
					}
					diff, err := resource.Diff(state, terraform.NewResourceConfigRaw(config), nil)
					if err != nil {
						t.Fatal(err)
					}
					if diff != nil {
						for attribute, change := range diff.Attributes {
							if strings.HasPrefix(attribute, key+".") {
								t.Fatalf("round trip produced drift for %s: %#v", attribute, change)
							}
						}
					}
				})
			}
		}
	}
}

func TestVMGroupV2ReadOverwrite(t *testing.T) {
	for _, prior := range []string{"", "false", "true"} {
		for _, remote := range []string{"absent", "false", "true", "no content", "empty content", "empty block"} {
			t.Run("prior="+prior+"/remote="+remote, func(t *testing.T) {
				config := vmGroupV2TestConfig("content", prior)
				filterPrior := "true"
				if prior == "true" {
					filterPrior = "false"
				}
				config["filters"] = vmGroupV2TestConfig("filters", filterPrior)["filters"]
				block := vmGroupV2TestContentJSON
				if remote == "false" || remote == "true" {
					block = `{"overwrite":` + remote + `,` + block[1:]
				}
				response := fmt.Sprintf(`{"name":"azure-vms","content":[%s],"filters":[%s]}`, block, block)
				switch remote {
				case "no content":
					response = `{"name":"azure-vms"}`
				case "empty content":
					response = `{"name":"azure-vms","content":[],"filters":[]}`
				case "empty block":
					response = `{"name":"azure-vms","content":[{}],"filters":[{}]}`
				}
				vmGroupV2GetServer(t, response)
				resource := resourceVMGroup_V2()
				state := refreshTestResource(t, resource, config)
				for key, want := range map[string]string{"content": prior, "filters": filterPrior} {
					if remote == "no content" || remote == "empty content" || remote == "empty block" {
						if count := state.Attributes[key+".#"]; count != "0" {
							t.Errorf("%s retained removed content: count = %q", key, count)
						}
						continue
					}
					if remote != "absent" {
						want = remote
					}
					if got := state.Attributes[key+".0.overwrite"]; got != want {
						t.Errorf("%s overwrite = %q; want %q", key, got, want)
					}
				}
			})
		}
	}
}

func TestVMGroupV2DatacenterIdentifier(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input map[string]interface{}
		want  string
	}{
		{name: "region value", input: map[string]interface{}{"type": "DATACENTER", "name": "Australia Southeast", "value": "australiasoutheast"}, want: `{"type":"DATACENTER","name":"Australia Southeast","GUID":"australiasoutheast"}`},
		{name: "explicit guid", input: map[string]interface{}{"type": "DATACENTER", "name": "Australia Southeast", "guid": "explicit-id", "value": "australiasoutheast"}, want: `{"type":"DATACENTER","name":"Australia Southeast","GUID":"explicit-id","value":"australiasoutheast"}`},
		{name: "guid only", input: map[string]interface{}{"type": "DATACENTER", "guid": "datacenter-1"}, want: `{"type":"DATACENTER","GUID":"datacenter-1"}`},
		{name: "name only", input: map[string]interface{}{"type": "DATACENTER", "name": "datacenter-1"}, want: `{"type":"DATACENTER","name":"datacenter-1"}`},
		{name: "tag value", input: map[string]interface{}{"type": "TAG", "value": "production"}, want: `{"type":"TAG","value":"production"}`},
		{name: "power state value", input: map[string]interface{}{"type": "VM_POWER_STATE", "value": "on"}, want: `{"type":"VM_POWER_STATE","value":"on"}`},
	} {
		t.Run(tc.name, func(t *testing.T) {
			rules := build_vmgroup_v2_msgrulecontentset_array(nil, []interface{}{tc.input})
			raw, err := json.Marshal(rules[0])
			if err != nil {
				t.Fatal(err)
			}
			var got, want map[string]interface{}
			if err := json.Unmarshal(raw, &got); err != nil {
				t.Fatal(err)
			}
			if err := json.Unmarshal([]byte(tc.want), &want); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("request = %s; want %s", raw, tc.want)
			}
		})
	}
}

func TestVMGroupV2UpdateResponse(t *testing.T) {
	for _, tc := range []struct {
		name      string
		body      string
		wantError string
	}{
		{name: "top-level success", body: `{"errorCode":0}`},
		{name: "top-level error", body: `{"errorCode":17,"errorMessage":"rule rejected"}`, wantError: "code=17, message=rule rejected"},
		{name: "top-level error without message", body: `{"errorCode":17}`, wantError: "code=17"},
		{name: "wrapped success", body: `{"response":{"errorCode":0}}`},
		{name: "observed array success", body: `{"response":[{"warningCode":0,"errorCode":0,"warningMessage":""}]}`},
		{name: "array error", body: `{"response":[{"errorCode":17,"errorMessage":"rule rejected"}]}`, wantError: "rule rejected"},
		{name: "later array error", body: `{"response":[{"errorCode":0},{"errorCode":18}]}`, wantError: "code=18"},
		{name: "array error with top-level success", body: `{"errorCode":0,"response":[{"errorCode":17}]}`, wantError: "code=17"},
		{name: "top-level error with array success", body: `{"errorCode":18,"response":[{"errorCode":0}]}`, wantError: "code=18"},
		{name: "invalid response shape", body: `{"response":"invalid"}`, wantError: "cannot unmarshal"},
		{name: "wrapped error", body: `{"response":{"errorCode":17,"errorMessage":"rule rejected"}}`, wantError: "rule rejected"},
		{name: "wrapped error without message", body: `{"response":{"errorCode":17}}`, wantError: "code=17"},
		{name: "wrapped error with top-level success", body: `{"errorCode":0,"response":{"errorCode":17}}`, wantError: "code=17"},
		{name: "top-level error with wrapped success", body: `{"errorCode":18,"response":{"errorCode":0}}`, wantError: "code=18"},
		{name: "invalid JSON", body: `{`, wantError: "unexpected end of JSON input"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			vmGroupV2TestServer(t, func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/V4/VmGroup/42" {
					t.Errorf("unexpected path: %s", r.URL.Path)
				}
				w.Header().Set("Content-Type", "application/json")
				if r.Method == http.MethodPut {
					fmt.Fprint(w, tc.body)
					return
				}
				if r.Method != http.MethodGet || tc.wantError != "" {
					t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
				}
				fmt.Fprintf(w, `{"name":"azure-vms","content":[%s]}`, vmGroupV2TestContentJSON)
			})
			resource := resourceVMGroup_V2()
			data := schema.TestResourceDataRaw(t, resource.Schema, vmGroupV2TestConfig("content", "true"))
			data.SetId("42")
			err := resource.Update(data, nil)
			if tc.wantError == "" {
				if err != nil {
					t.Fatal(err)
				}
			} else if err == nil || !strings.Contains(err.Error(), tc.wantError) {
				t.Fatalf("update error = %v; want %q", err, tc.wantError)
			}
			if data.Id() != "42" {
				t.Errorf("update changed resource ID to %q", data.Id())
			}
		})
	}
}

func TestVMGroupV2ReadMetadataMatching(t *testing.T) {
	group := func(match string, rules ...interface{}) map[string]interface{} {
		return map[string]interface{}{"matchrule": match, "rules": rules}
	}
	rule := func(name, guid string) map[string]interface{} {
		return map[string]interface{}{"condition": "EQUALS", "type": "VM", "name": name, "guid": guid}
	}
	remote := `{"ruleGroups":[{"matchRule":"ANY","rules":[{"condition":"EQUALS","type":"VM","name":"shared","GUID":"remote-b","displayName":"Remote B","description":"Description B"}]},{"matchRule":"ALL","rules":[{"condition":"EQUALS","type":"VM","name":"shared","GUID":"remote-a","displayName":"Remote A","description":"Description A"}]}]}`
	for _, key := range []string{"content", "filters"} {
		for _, tc := range []struct {
			name   string
			groups []interface{}
		}{
			{name: "import"},
			{name: "no previous rules", groups: []interface{}{group("ALL")}},
			{name: "ambiguous prior selectors", groups: []interface{}{
				group("ALL", rule("shared", "prior-a")), group("ANY", rule("shared", "prior-b")),
			}},
			{name: "ambiguous remote selectors", groups: []interface{}{
				group("ALL", rule("shared", "")),
			}},
		} {
			t.Run(key+"/"+tc.name, func(t *testing.T) {
				response := remote
				if tc.name == "ambiguous remote selectors" {
					response = strings.ReplaceAll(response, `"matchRule":"ANY"`, `"matchRule":"ALL"`)
					response = strings.Replace(response, `]},{"matchRule":"ALL","rules":[`, `,`, 1)
				}
				vmGroupV2GetServer(t, fmt.Sprintf(`{"name":"azure-vms",%q:[%s]}`, key, response))
				resource := resourceVMGroup_V2()
				config := map[string]interface{}{}
				if tc.groups != nil {
					config[key] = []interface{}{map[string]interface{}{"rulegroups": tc.groups}}
				}
				state := refreshTestResource(t, resource, config)
				wantA, wantB := rule("shared", "remote-a"), rule("shared", "remote-b")
				wantA["displayname"], wantA["description"] = "Remote A", "Description A"
				wantB["displayname"], wantB["description"] = "Remote B", "Description B"
				wantGroups := []interface{}{group("ALL", wantA), group("ANY", wantB)}
				if tc.name == "ambiguous remote selectors" {
					wantGroups = []interface{}{group("ALL", wantA, wantB)}
				}
				want := schema.TestResourceDataRaw(t, resource.Schema, map[string]interface{}{
					key: []interface{}{map[string]interface{}{"rulegroups": wantGroups}},
				})
				want.SetId("42")
				gotAttributes, wantAttributes := map[string]string{}, map[string]string{}
				for attribute, value := range state.Attributes {
					if strings.HasPrefix(attribute, key+".0.rulegroups.") {
						gotAttributes[attribute] = value
					}
				}
				for attribute, value := range want.State().Attributes {
					if strings.HasPrefix(attribute, key+".0.rulegroups.") {
						wantAttributes[attribute] = value
					}
				}
				if !reflect.DeepEqual(gotAttributes, wantAttributes) {
					t.Fatalf("remote rules or metadata were masked: got %#v; want %#v", gotAttributes, wantAttributes)
				}
			})
		}
	}
}

func TestVMGroupV2ReadReorderedGroupsDiff(t *testing.T) {
	for _, key := range []string{"content", "filters"} {
		for _, change := range []string{"reorder only", "condition", "name", "type", "value", "matchrule", "regroup"} {
			t.Run(key+"/"+change, func(t *testing.T) {
				config := vmGroupV2TestConfig(key, "false")
				block := config[key].([]interface{})[0].(map[string]interface{})
				group := block["rulegroups"].([]interface{})[0].(map[string]interface{})
				rules := group["rules"].([]interface{})
				for _, field := range []string{"displayname", "guid", "description"} {
					delete(rules[1].(map[string]interface{}), field)
				}
				group["rules"] = rules[:1]
				block["rulegroups"] = []interface{}{group, map[string]interface{}{"matchrule": "ANY", "rules": rules[1:]}}
				remote := `{"ruleGroups":[{"matchRule":"ANY","rules":[{"condition":"EQUALS","type":"TAG","name":"environment","value":"production","displayName":"Environment","GUID":"tag-guid","description":"Production machines"}]},{"matchRule":"ALL","rules":[{"condition":"EQUALS","type":"DATACENTER","name":"Australia Southeast","value":"australiasoutheast","GUID":"region-guid"}]}]}`
				switch change {
				case "condition":
					remote = strings.Replace(remote, `"condition":"EQUALS"`, `"condition":"CONTAINS"`, 1)
				case "name":
					remote = strings.Replace(remote, "environment", "department", 1)
				case "type":
					remote = strings.Replace(remote, `"type":"TAG"`, `"type":"VM"`, 1)
				case "value":
					remote = strings.Replace(remote, "production", "development", 1)
				case "matchrule":
					remote = strings.Replace(remote, `"matchRule":"ANY"`, `"matchRule":"ALL"`, 1)
				case "regroup":
					remote = strings.Replace(remote, `]},{"matchRule":"ALL","rules":[`, `,`, 1)
				}
				vmGroupV2GetServer(t, fmt.Sprintf(`{"name":"azure-vms",%q:[%s]}`, key, remote))
				resource := resourceVMGroup_V2()
				state := refreshTestResource(t, resource, config)
				selectorChanged := change == "condition" || change == "name" || change == "type" || change == "value"
				metadata := map[string]string{}
				for attribute, value := range state.Attributes {
					if strings.HasPrefix(attribute, key+".") && strings.Contains(attribute, ".rules.") {
						for _, field := range []string{"guid", "displayname", "description"} {
							if strings.HasSuffix(attribute, "."+field) && value != "" {
								metadata[field] = value
							}
						}
					}
				}
				wantMetadata := map[string]string{}
				if selectorChanged {
					wantMetadata = map[string]string{"guid": "tag-guid", "displayname": "Environment", "description": "Production machines"}
				}
				if !reflect.DeepEqual(metadata, wantMetadata) {
					t.Errorf("metadata = %#v; want %#v", metadata, wantMetadata)
				}
				diff, err := resource.Diff(state, terraform.NewResourceConfigRaw(config), nil)
				if err != nil {
					t.Fatalf("diff: %s", err)
				}
				changed := false
				if diff != nil {
					for attribute, delta := range diff.Attributes {
						if strings.HasPrefix(attribute, key+".") {
							changed = true
							if delta.RequiresNew {
								t.Errorf("rule change requires replacement: %s", attribute)
							}
						}
					}
				}
				if changed != (change != "reorder only") {
					t.Fatalf("changed = %t; diff = %#v", changed, diff)
				}
			})
		}
	}
}
