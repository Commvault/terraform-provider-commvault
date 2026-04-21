package test

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// randSuffix returns a 6-character random hex string so subclient names are
// unique per test run and never clash with existing ones.
func randSuffix() string {
	b := make([]byte, 3)
	rand.Read(b)
	return hex.EncodeToString(b)
}

// uniqueSubclientName builds a unique subclient name for a test run.
func uniqueSubclientName(suffix string) string {
	return fmt.Sprintf("tf_test_%s", suffix)
}

func dbConnectCredentialBlock() string {
	if v := testDbConnectCredentialID(); v != "" {
		return fmt.Sprintf("\n  db_connect_credential_id = %s", v)
	}
	return ""
}

// ---------------------------------------------------------------------------
// Oracle Instance acceptance tests
//
// These tests use the REAL instance name from CV_TEST_INSTANCE_NAME because
// Commvault validates actual DB connectivity on create — a fake name returns
// error 4444 "Unable to establish connectivity".
//
// If the instance already exists (HTTP 409) the resource automatically adopts
// it (looks up the instance ID and proceeds with Update+Read).
//
// CheckDestroy is intentionally nil so the pre-existing production instance
// (e.g. "destdb") is NOT deleted when the test finishes.
// ---------------------------------------------------------------------------

func TestAccResourceOracleInstance_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: nil, // don't delete the pre-existing instance
		Steps: []resource.TestStep{
			{
				Config: testAccResourceOracleInstanceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testCheckResourceExists("Oracle Instance", "commvault_oracle_instance.test"),
					resource.TestCheckResourceAttr("commvault_oracle_instance.test", "client_name", os.Getenv(envClientName)),
					resource.TestCheckResourceAttr("commvault_oracle_instance.test", "instance_name", os.Getenv(envInstanceName)),
					resource.TestCheckResourceAttr("commvault_oracle_instance.test", "oracle_home", os.Getenv(envOracleHome)),
				),
			},
		},
	})
}

func TestAccResourceOracleInstance_withSqlConnect(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceOracleInstanceConfig_withSqlConnect(),
				Check: resource.ComposeTestCheckFunc(
					testCheckResourceExists("Oracle Instance", "commvault_oracle_instance.test"),
					resource.TestCheckResourceAttr("commvault_oracle_instance.test", "sql_connect_user", "/"),
				),
			},
		},
	})
}

func TestAccResourceOracleInstance_withCatalog(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceOracleInstanceConfig_withCatalog(),
				Check: resource.ComposeTestCheckFunc(
					testCheckResourceExists("Oracle Instance", "commvault_oracle_instance.test"),
					resource.TestCheckResourceAttr("commvault_oracle_instance.test", "use_catalog_connect", "true"),
					resource.TestCheckResourceAttr("commvault_oracle_instance.test", "catalog_connect_user", "rman"),
					resource.TestCheckResourceAttr("commvault_oracle_instance.test", "catalog_connect_domain", "RMANCAT"),
				),
			},
		},
	})
}

func TestAccResourceOracleInstance_withPlan(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceOracleInstanceConfig_withPlan(),
				Check: resource.ComposeTestCheckFunc(
					testCheckResourceExists("Oracle Instance", "commvault_oracle_instance.test"),
					resource.TestCheckResourceAttr("commvault_oracle_instance.test", "plan_id", testPlanID()),
				),
			},
		},
	})
}

// TestAccResourceOracleInstance_update is the primary end-to-end test:
// Step 1 — adopt/create the instance with default settings.
// Step 2 — update block_size and cross_check_timeout and verify they stick.
// Step 3 — restore original settings (clean up changes on the live instance).
func TestAccResourceOracleInstance_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceOracleInstanceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testCheckResourceExists("Oracle Instance", "commvault_oracle_instance.test"),
				),
			},
			{
				Config: testAccResourceOracleInstanceConfig_updated(),
				Check: resource.ComposeTestCheckFunc(
					testCheckResourceExists("Oracle Instance", "commvault_oracle_instance.test"),
					resource.TestCheckResourceAttr("commvault_oracle_instance.test", "block_size", "131072"),
					resource.TestCheckResourceAttr("commvault_oracle_instance.test", "cross_check_timeout", "1200"),
				),
			},
			// Step 3: restore original values so the live instance is left clean.
			{
				Config: testAccResourceOracleInstanceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testCheckResourceExists("Oracle Instance", "commvault_oracle_instance.test"),
				),
			},
		},
	})
}

// ---------------------------------------------------------------------------
// Config helpers — read env vars at call time (never at struct-literal time).
// ---------------------------------------------------------------------------

func testAccResourceOracleInstanceConfig_basic() string {
	return fmt.Sprintf(`%s
resource "commvault_oracle_instance" "test" {
  client_name   = "%s"
  instance_name = "%s"
  oracle_home   = "%s"
  oracle_user   = "%s"
	%s
  block_size    = 65536
}
`, ProviderConfig(), os.Getenv(envClientName), os.Getenv(envInstanceName), os.Getenv(envOracleHome), os.Getenv(envOracleUser), dbConnectCredentialBlock())
}

func testAccResourceOracleInstanceConfig_withSqlConnect() string {
	return fmt.Sprintf(`%s
resource "commvault_oracle_instance" "test" {
  client_name        = "%s"
  instance_name      = "%s"
  oracle_home        = "%s"
  oracle_user        = "%s"
  sql_connect_user   = "/"
	%s
  block_size         = 65536
}
`, ProviderConfig(), os.Getenv(envClientName), os.Getenv(envInstanceName), os.Getenv(envOracleHome), os.Getenv(envOracleUser), dbConnectCredentialBlock())
}

func testAccResourceOracleInstanceConfig_withCatalog() string {
	return fmt.Sprintf(`%s
resource "commvault_oracle_instance" "test" {
  client_name            = "%s"
  instance_name          = "%s"
  oracle_home            = "%s"
  oracle_user            = "%s"
  use_catalog_connect    = true
  catalog_connect_user   = "rman"
  catalog_connect_domain = "RMANCAT"
  block_size             = 65536
}
`, ProviderConfig(), os.Getenv(envClientName), os.Getenv(envInstanceName), os.Getenv(envOracleHome), os.Getenv(envOracleUser))
}

func testAccResourceOracleInstanceConfig_updated() string {
	return fmt.Sprintf(`%s
resource "commvault_oracle_instance" "test" {
  client_name         = "%s"
  instance_name       = "%s"
  oracle_home         = "%s"
  oracle_user         = "%s"
	%s
  block_size          = 131072
  cross_check_timeout = 1200
}
`, ProviderConfig(), os.Getenv(envClientName), os.Getenv(envInstanceName), os.Getenv(envOracleHome), os.Getenv(envOracleUser), dbConnectCredentialBlock())
}

func testAccResourceOracleInstanceConfig_withPlan() string {
	return fmt.Sprintf(`%s
resource "commvault_oracle_instance" "test" {
  client_name   = "%s"
  instance_name = "%s"
  oracle_home   = "%s"
  oracle_user   = "%s"
	%s
  plan_id       = %s
  block_size    = 65536
}
`, ProviderConfig(), os.Getenv(envClientName), os.Getenv(envInstanceName), os.Getenv(envOracleHome), os.Getenv(envOracleUser), dbConnectCredentialBlock(), testPlanID())
}
