package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDatasourceOracleInstance(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDatasourceOracleInstance(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.commvault_oracle_instance.test", "client_name", os.Getenv(envClientName)),
					resource.TestCheckResourceAttr("data.commvault_oracle_instance.test", "instance_name", os.Getenv(envInstanceName)),
					resource.TestCheckResourceAttrSet("data.commvault_oracle_instance.test", "instance_id"),
				),
			},
		},
	})
}

func TestAccDatasourceOracleSubclient(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDatasourceOracleSubclient(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.commvault_oracle_subclient.test", "client_name", os.Getenv(envClientName)),
					resource.TestCheckResourceAttr("data.commvault_oracle_subclient.test", "instance_name", os.Getenv(envInstanceName)),
					resource.TestCheckResourceAttrSet("data.commvault_oracle_subclient.test", "subclient_id"),
				),
			},
		},
	})
}

func testAccDatasourceOracleInstance() string {
	return fmt.Sprintf(`%s
data "commvault_oracle_instance" "test" {
  client_name   = "%s"
  instance_name = "%s"
}
`, ProviderConfig(), os.Getenv(envClientName), os.Getenv(envInstanceName))
}

func testAccDatasourceOracleSubclient() string {
	return fmt.Sprintf(`%s
data "commvault_oracle_subclient" "test" {
  client_name    = "%s"
  instance_name  = "%s"
  subclient_name = "default"
}
`, ProviderConfig(), os.Getenv(envClientName), os.Getenv(envInstanceName))
}
