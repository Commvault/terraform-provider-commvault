package commvault

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

const vmGroupV2DiskFiltersJSON = `{"rules":[{"condition":"EQUALS","filterType":"DISK_TAG","name":"backup","value":"exclude"},{"condition":"EQUALS","filterType":"DISK_PATH","name":"/data","vmName":"test-vm","vmGuid":"test-guid","children":[{"condition":"EQUALS","filterType":"DISK_PATH","name":"/data/b"},{"condition":"EQUALS","filterType":"DISK_PATH","name":"/data/a"}]}]}`

func vmGroupV2DiskFilterConfig(overwrite string) map[string]interface{} {
	rule := func(name string) map[string]interface{} {
		value := map[string]interface{}{"condition": "EQUALS", "filtertype": "DISK_PATH", "name": name}
		if overwrite != "" {
			value["overwrite"] = overwrite
		}
		return value
	}
	parent := rule("/data")
	parent["vmname"], parent["vmguid"] = "test-vm", "test-guid"
	parent["children"] = []interface{}{rule("/data/a"), rule("/data/b")}
	tag := rule("backup")
	tag["filtertype"], tag["value"] = "DISK_TAG", "exclude"
	block := map[string]interface{}{"rules": []interface{}{parent, tag}}
	if overwrite != "" {
		block["overwrite"] = overwrite
	}
	return map[string]interface{}{
		"name": "test-vms", "hypervisor": []interface{}{map[string]interface{}{"id": 12}},
		"diskfilters": []interface{}{block},
	}
}

func TestVMGroupV2DiskFilterRoundTrip(t *testing.T) {
	for _, create := range []bool{false, true} {
		for _, overwrite := range []string{"", "false", "true"} {
			t.Run(fmt.Sprintf("create=%t/overwrite=%s", create, overwrite), func(t *testing.T) {
				var methods []string
				vmGroupV2TestServer(t, func(w http.ResponseWriter, r *http.Request) {
					methods = append(methods, r.Method)
					w.Header().Set("Content-Type", "application/json")
					if r.Method == http.MethodPost && create && r.URL.Path == "/v5/VmGroup" {
						fmt.Fprint(w, `{"subclientId":42,"errorCode":0}`)
						return
					}
					if r.URL.Path != "/V4/VmGroup/42" {
						t.Errorf("unexpected path: %s", r.URL.Path)
					}
					switch r.Method {
					case http.MethodPut:
						var request map[string]interface{}
						if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
							t.Error(err)
							return
						}
						block, ok := request["diskFilters"].(map[string]interface{})
						if !ok {
							t.Errorf("missing diskFilters: %#v", request)
							return
						}
						if create && len(request) != 1 {
							t.Errorf("post-create update included unrelated fields: %#v", request)
						}
						var checkOverwrite func(map[string]interface{})
						checkOverwrite = func(value map[string]interface{}) {
							got, present := value["overwrite"]
							if present != (overwrite != "") || present && got != (overwrite == "true") {
								t.Errorf("wire overwrite = %#v (present=%t); want %q", got, present, overwrite)
							}
							for _, key := range []string{"rules", "children"} {
								if nested, ok := value[key].([]interface{}); ok {
									for _, child := range nested {
										checkOverwrite(child.(map[string]interface{}))
									}
								}
							}
						}
						checkOverwrite(block)
						if rules, ok := block["rules"].([]interface{}); !ok || len(rules) != 2 {
							t.Errorf("wrong wire rules: %#v", block)
						} else {
							for _, raw := range rules {
								rule := raw.(map[string]interface{})
								switch rule["name"] {
								case "/data":
									children, ok := rule["children"].([]interface{})
									if !ok || len(children) != 2 || rule["vmName"] != "test-vm" || rule["vmGuid"] != "test-guid" {
										t.Errorf("parent rule lost fields or children: %#v", rule)
									}
								case "backup":
									if rule["filterType"] != "DISK_TAG" || rule["value"] != "exclude" {
										t.Errorf("tag filter lost fields: %#v", rule)
									}
								default:
									t.Errorf("unexpected rule: %#v", rule)
								}
							}
						}
						fmt.Fprint(w, `{"response":[{"errorCode":0}]}`)
					case http.MethodGet:
						fmt.Fprintf(w, `{"name":"test-vms","diskFilters":[%s]}`, vmGroupV2DiskFiltersJSON)
					default:
						t.Errorf("unexpected request: %s", r.Method)
					}
				})
				resource := resourceVMGroup_V2()
				config := vmGroupV2DiskFilterConfig(overwrite)
				data := schema.TestResourceDataRaw(t, resource.Schema, config)
				var operation func(*schema.ResourceData, interface{}) error = resource.Update
				if create {
					operation = resource.Create
				} else {
					data.SetId("42")
				}
				if err := operation(data, nil); err != nil {
					t.Fatal(err)
				}
				wantMethods := []string{http.MethodPut, http.MethodGet}
				if create {
					wantMethods = append([]string{http.MethodPost}, wantMethods...)
				}
				if !reflect.DeepEqual(methods, wantMethods) {
					t.Fatalf("requests = %v; want %v", methods, wantMethods)
				}
				state, err := resource.RefreshWithoutUpgrade(data.State(), nil)
				if err != nil {
					t.Fatal(err)
				}
				diff, err := resource.Diff(state, terraform.NewResourceConfigRaw(config), nil)
				if err != nil {
					t.Fatal(err)
				}
				if diff != nil {
					for key, change := range diff.Attributes {
						if strings.HasPrefix(key, "diskfilters.") {
							t.Fatalf("disk filter drift for %s: %#v", key, change)
						}
					}
				}
			})
		}
	}
}

func TestVMGroupV2DiskFilterReadOverwrite(t *testing.T) {
	for _, scenario := range []string{"reordered", "remote false", "remote true", "changed parent", "changed child", "removed child", "removed rule", "absent", "empty list", "empty block", "no prior state"} {
		t.Run(scenario, func(t *testing.T) {
			config := vmGroupV2DiskFilterConfig("true")
			block := config["diskfilters"].([]interface{})[0].(map[string]interface{})
			rules := block["rules"].([]interface{})
			rules[1].(map[string]interface{})["overwrite"] = "false"
			rules[0].(map[string]interface{})["children"].([]interface{})[1].(map[string]interface{})["overwrite"] = "false"
			remote := vmGroupV2DiskFiltersJSON
			switch scenario {
			case "remote false", "remote true":
				remote = strings.ReplaceAll(remote, `{`, `{"overwrite":`+strings.TrimPrefix(scenario, "remote ")+`,`)
			case "changed parent":
				remote = strings.Replace(remote, `"vmGuid":"test-guid"`, `"vmGuid":"different-guid"`, 1)
			case "changed child":
				remote = strings.Replace(remote, `"name":"/data/a"`, `"name":"/other"`, 1)
			case "removed child":
				remote = strings.Replace(remote, `,{"condition":"EQUALS","filterType":"DISK_PATH","name":"/data/a"}`, ``, 1)
			case "removed rule":
				remote = strings.Replace(remote, `{"condition":"EQUALS","filterType":"DISK_TAG","name":"backup","value":"exclude"},`, ``, 1)
			}
			response := fmt.Sprintf(`{"name":"test-vms","diskFilters":[%s]}`, remote)
			switch scenario {
			case "absent":
				response = `{"name":"test-vms"}`
			case "empty list":
				response = `{"name":"test-vms","diskFilters":[]}`
			case "empty block":
				response = `{"name":"test-vms","diskFilters":[{}]}`
			case "no prior state":
				delete(config, "diskfilters")
			}
			vmGroupV2GetServer(t, response)
			resource := resourceVMGroup_V2()
			data := schema.TestResourceDataRaw(t, resource.Schema, config)
			data.SetId("42")
			if err := resource.Read(data, nil); err != nil {
				t.Fatal(err)
			}
			state := data.State()
			if scenario == "absent" || scenario == "empty list" || scenario == "empty block" {
				if state.Attributes["diskfilters.#"] != "0" {
					t.Fatal("removed diskfilters restored from state")
				}
			} else {
				for key, value := range state.Attributes {
					if !strings.HasPrefix(key, "diskfilters.") || !strings.HasSuffix(key, ".overwrite") {
						continue
					}
					if strings.HasPrefix(scenario, "remote ") && value != strings.TrimPrefix(scenario, "remote ") {
						t.Errorf("%s = %q; API value must win", key, value)
					}
					if scenario == "no prior state" && value != "" {
						t.Errorf("invented overwrite for %s: %q", key, value)
					}
				}
			}
			diff, err := resource.Diff(state, terraform.NewResourceConfigRaw(config), nil)
			if err != nil {
				t.Fatal(err)
			}
			changed := false
			if diff != nil {
				for key := range diff.Attributes {
					changed = changed || strings.HasPrefix(key, "diskfilters.")
				}
			}
			wantChange := scenario != "reordered" && scenario != "no prior state"
			if changed != wantChange {
				t.Fatalf("disk filter diff = %t; want %t: %#v", changed, wantChange, diff)
			}
		})
	}
}

func TestVMGroupV2DiskFilterCreateUpdateFailure(t *testing.T) {
	vmGroupV2TestServer(t, func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			fmt.Fprint(w, `{"subclientId":42,"errorCode":0}`)
		case http.MethodPut:
			fmt.Fprint(w, `{"response":[{"errorCode":17,"errorMessage":"disk filter rejected"}]}`)
		default:
			t.Errorf("unexpected request after failed update: %s", r.Method)
		}
	})
	resource := resourceVMGroup_V2()
	data := schema.TestResourceDataRaw(t, resource.Schema, vmGroupV2DiskFilterConfig("false"))
	err := resource.Create(data, nil)
	if err == nil || !strings.Contains(err.Error(), "disk filter rejected") {
		t.Fatalf("create error = %v", err)
	}
	if data.Id() != "42" {
		t.Fatal("failed disk-filter update lost the newly created VM group ID")
	}
}

func TestVMGroupV2DiskFilterOverwriteAmbiguity(t *testing.T) {
	for _, scenario := range []string{"prior rules", "remote rules", "prior children", "remote children"} {
		t.Run(scenario, func(t *testing.T) {
			config := vmGroupV2DiskFilterConfig("true")
			block := config["diskfilters"].([]interface{})[0].(map[string]interface{})
			rule := map[string]interface{}{"condition": "EQUALS", "filtertype": "DISK_PATH", "name": "/data", "overwrite": "true"}
			other := map[string]interface{}{"condition": "EQUALS", "filtertype": "DISK_PATH", "name": "/data", "overwrite": "false"}
			remoteRule := `{"condition":"EQUALS","filterType":"DISK_PATH","name":"/data"}`
			remoteOther := `{"condition":"EQUALS","filterType":"DISK_PATH","name":"/data","overwrite":false}`
			priorRules := []interface{}{rule}
			remoteRules := remoteRule
			if strings.HasPrefix(scenario, "prior") {
				priorRules = append(priorRules, other)
			} else {
				remoteRules += "," + remoteOther
			}
			if strings.HasSuffix(scenario, "children") {
				priorRules = []interface{}{map[string]interface{}{"name": "parent", "filtertype": "DISK_PATH", "children": priorRules}}
				remoteRules = `{"name":"parent","filterType":"DISK_PATH","children":[` + remoteRules + `]}`
			}
			block["rules"] = priorRules
			vmGroupV2GetServer(t, fmt.Sprintf(`{"name":"test-vms","diskFilters":[{"rules":[%s]}]}`, remoteRules))
			resource := resourceVMGroup_V2()
			data := schema.TestResourceDataRaw(t, resource.Schema, config)
			data.SetId("42")
			if err := resource.Read(data, nil); err != nil {
				t.Fatal(err)
			}
			for key, value := range data.State().Attributes {
				if strings.HasPrefix(key, "diskfilters.0.rules.") && strings.HasSuffix(key, ".overwrite") && value == "true" {
					t.Fatalf("assigned prior overwrite to ambiguous rule: %s", key)
				}
			}
		})
	}
}
