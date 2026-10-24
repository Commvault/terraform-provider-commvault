package commvault

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"terraform-provider-commvault/commvault/handler"

	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestSerializeHypervisorAzureCredentials(t *testing.T) {
	id, name := 12, "azure-credential"
	zeroID, emptyName := 0, ""
	for _, tt := range []struct {
		name string
		data *handler.MsgVMExistingCredential
		want []map[string]interface{}
	}{
		{name: "nil"},
		{name: "empty", data: &handler.MsgVMExistingCredential{}},
		{
			name: "id and name",
			data: &handler.MsgVMExistingCredential{CredentialId: &id, CredentialName: &name},
			want: []map[string]interface{}{{"id": &id, "name": &name}},
		},
		{
			name: "id only",
			data: &handler.MsgVMExistingCredential{CredentialId: &id},
			want: []map[string]interface{}{{"id": &id}},
		},
		{
			name: "name only",
			data: &handler.MsgVMExistingCredential{CredentialName: &name},
			want: []map[string]interface{}{{"name": &name}},
		},
		{
			name: "explicit zero values",
			data: &handler.MsgVMExistingCredential{CredentialId: &zeroID, CredentialName: &emptyName},
			want: []map[string]interface{}{{"id": &zeroID, "name": &emptyName}},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := serialize_hypervisor_azure_msgvmexistingcredential(nil, tt.data)
			if ok != (tt.want != nil) || !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("serialize credentials = %#v, %t; want %#v, %t", got, ok, tt.want, tt.want != nil)
			}
		})
	}
}

func TestHypervisorAzureReadCredentialsDiff(t *testing.T) {
	for _, tt := range []struct {
		name     string
		remoteID int
	}{
		{name: "unchanged", remoteID: 12},
		{name: "remote credential changed", remoteID: 34},
	} {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodGet || r.URL.Path != "/V4/Hypervisor/42" {
					t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
					http.NotFound(w, r)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprintf(w, `{"name":"azure","credentials":{"credentialId":%d,"credentialName":"azure-credential"}}`, tt.remoteID)
			}))
			defer server.Close()
			t.Setenv("CV_CSIP", server.URL)
			t.Setenv("AuthToken", "test-token")
			t.Setenv("CV_LOGGING", "false")

			resource := resourceHypervisor_Azure()
			config := map[string]interface{}{
				"name":           "azure",
				"subscriptionid": "test-subscription",
				"credentials":    []interface{}{map[string]interface{}{"id": 12}},
			}
			state := refreshTestResource(t, resource, config)
			for key, want := range map[string]string{
				"credentials.#":      "1",
				"credentials.0.id":   strconv.Itoa(tt.remoteID),
				"credentials.0.name": "azure-credential",
			} {
				if got := state.Attributes[key]; got != want {
					t.Errorf("refreshed %s = %q; want %q", key, got, want)
				}
			}

			diff, err := resource.Diff(state, terraform.NewResourceConfigRaw(config), nil)
			if err != nil {
				t.Fatalf("diff: %s", err)
			}
			if tt.remoteID == 12 {
				if diff != nil {
					for key, change := range diff.Attributes {
						if strings.HasPrefix(key, "credentials.") {
							t.Fatalf("unchanged credentials produced a diff for %s: %#v", key, change)
						}
					}
				}
				return
			}
			if diff == nil || diff.Attributes["credentials.0.id"] == nil {
				t.Fatal("remote credential ID change was not detected")
			}
			change := diff.Attributes["credentials.0.id"]
			if change.Old != strconv.Itoa(tt.remoteID) || change.New != "12" || change.NewComputed || change.RequiresNew {
				t.Fatalf("unexpected credential ID diff: %#v", change)
			}
		})
	}
}
