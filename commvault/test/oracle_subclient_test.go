package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// ---------------------------------------------------------------------------
// Oracle Subclient acceptance tests
// ---------------------------------------------------------------------------

func TestAccResourceOracleSubclient_basic(t *testing.T) {
	suffix := randSuffix()
	scName := uniqueSubclientName(suffix)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckResourceDestroy("commvault_oracle_subclient"),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceOracleSubclientConfig_basic(scName),
				Check: resource.ComposeTestCheckFunc(
					testCheckResourceExists("Oracle Subclient", "commvault_oracle_subclient.test"),
					resource.TestCheckResourceAttr("commvault_oracle_subclient.test", "subclient_name", scName),
					resource.TestCheckResourceAttr("commvault_oracle_subclient.test", "client_name", os.Getenv(envClientName)),
					resource.TestCheckResourceAttr("commvault_oracle_subclient.test", "instance_name", os.Getenv(envInstanceName)),
					resource.TestCheckResourceAttr("commvault_oracle_subclient.test", "enable_backup", "true"),
				),
			},
		},
	})
}

func TestAccResourceOracleSubclient_update(t *testing.T) {
	suffix := randSuffix()
	scName := uniqueSubclientName(suffix)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckResourceDestroy("commvault_oracle_subclient"),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceOracleSubclientConfig_basic(scName),
				Check: resource.ComposeTestCheckFunc(
					testCheckResourceExists("Oracle Subclient", "commvault_oracle_subclient.test"),
				),
			},
			{
				Config: testAccResourceOracleSubclientConfig_updated(scName),
				Check: resource.ComposeTestCheckFunc(
					testCheckResourceExists("Oracle Subclient", "commvault_oracle_subclient.test"),
					resource.TestCheckResourceAttr("commvault_oracle_subclient.test", "backup_archive_log", "true"),
					resource.TestCheckResourceAttr("commvault_oracle_subclient.test", "delete_archive_log_after_backup", "true"),
					resource.TestCheckResourceAttr("commvault_oracle_subclient.test", "description", "Updated test subclient"),
				),
			},
		},
	})
}

// ---------------------------------------------------------------------------
// Config helpers
// ---------------------------------------------------------------------------

func testAccResourceOracleSubclientConfig_basic(scName string) string {
	return fmt.Sprintf(`%s
resource "commvault_oracle_subclient" "test" {
  subclient_name = "%s"
  client_name    = "%s"
  instance_name  = "%s"
  enable_backup  = true
}
`, ProviderConfig(), scName, os.Getenv(envClientName), os.Getenv(envInstanceName))
}

func testAccResourceOracleSubclientConfig_updated(scName string) string {
	return fmt.Sprintf(`%s
resource "commvault_oracle_subclient" "test" {
  subclient_name                  = "%s"
  client_name                     = "%s"
  instance_name                   = "%s"
  enable_backup                   = true
  description                     = "Updated test subclient"
  backup_archive_log              = true
  backup_sp_file                  = true
  backup_control_file             = true
  delete_archive_log_after_backup = true
}
`, ProviderConfig(), scName, os.Getenv(envClientName), os.Getenv(envInstanceName))
}
