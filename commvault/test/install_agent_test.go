package test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// TestAccOracleInstallAgent_basic tests the Oracle agent installation resource
// Note: This test is skipped by default - requires a valid target server
func TestAccOracleInstallAgent_basic(t *testing.T) {
	t.Skip("Skipping install agent test - requires target server setup")

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccOracleInstallAgentConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("commvault_oracle_install_agent.test", "client_name", "test-oracle-server"),
					resource.TestCheckResourceAttrSet("commvault_oracle_install_agent.test", "task_id"),
				),
			},
		},
	})
}

func testAccOracleInstallAgentConfig_basic() string {
	return ProviderConfig() + `
resource "commvault_oracle_install_agent" "test" {
  client_name         = "test-oracle-server"
  host_name           = "192.168.1.100"
  commserve_host_name = "your-commserve.example.com"
  user_name           = "root"
  password            = "test-password"
  install_os_type     = 2
  unix_group          = "oinstall"
}
`
}
