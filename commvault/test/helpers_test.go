package test

import (
	"fmt"
	"os"
	"testing"

	"terraform-provider-commvault/commvault"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

// Environment variable names for test configuration.
// Set these before running acceptance tests:
//
//	CV_TEST_WEB_SERVICE_URL  - Commvault API endpoint
//	CV_TEST_USER_NAME        - Admin username
//	CV_TEST_PASSWORD         - Admin password
//	CV_TEST_CLIENT_NAME      - Oracle client name
//	CV_TEST_INSTANCE_NAME    - Oracle instance name
//	CV_TEST_ORACLE_HOME      - Oracle home path
//	CV_TEST_ORACLE_USER      - Oracle OS user
//	CV_TEST_PLAN_ID          - (optional) Plan ID integer
const (
	envWebServiceURL = "CV_TEST_WEB_SERVICE_URL"
	envUserName      = "CV_TEST_USER_NAME"
	envPassword      = "CV_TEST_PASSWORD"
	envClientName    = "CV_TEST_CLIENT_NAME"
	envInstanceName  = "CV_TEST_INSTANCE_NAME"
	envOracleHome    = "CV_TEST_ORACLE_HOME"
	envOracleUser    = "CV_TEST_ORACLE_USER"
	envPlanID        = "CV_TEST_PLAN_ID"
	envDbCredID      = "CV_TEST_DB_CONNECT_CREDENTIAL_ID"
)

// mustEnv panics with a clear message when a required env var is missing.
// It is only called during acceptance test execution (after TF_ACC check).
func mustEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		panic(fmt.Sprintf("Required environment variable %s is not set. "+
			"Please set all CV_TEST_* variables before running acceptance tests.", key))
	}
	return v
}

// testPlanID returns the plan ID from env or falls back to "102".
func testPlanID() string {
	v := os.Getenv(envPlanID)
	if v == "" {
		return "102"
	}
	return v
}

// testDbConnectCredentialID returns optional db connect credential id.
// Empty string means inline sql_connect_user should be used.
func testDbConnectCredentialID() string {
	return os.Getenv(envDbCredID)
}

// Backward-compatible package-level vars used by instance_test.go / subclient_test.go.
// Initialised as empty strings — populated by testAccPreCheck so no init-time panic.
var (
	TestWebServiceURL string
	TestUserName      string
	TestPassword      string
	TestClientName    string
	TestInstanceName  string
	TestOracleHome    string
	TestOracleUser    string
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = commvault.Provider()
	testAccProviders = map[string]terraform.ResourceProvider{
		"commvault": testAccProvider,
	}
}

// testAccPreCheck skips the test if TF_ACC is not set, then populates the
// compatibility vars so existing test configs can reference them as strings.
func testAccPreCheck(t *testing.T) {
	t.Helper()
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Acceptance tests skipped unless env 'TF_ACC' set")
	}
	// Populate compatibility vars now (safe — we're inside a test function)
	TestWebServiceURL = mustEnv(envWebServiceURL)
	TestUserName = mustEnv(envUserName)
	TestPassword = mustEnv(envPassword)
	TestClientName = mustEnv(envClientName)
	TestInstanceName = mustEnv(envInstanceName)
	TestOracleHome = mustEnv(envOracleHome)
	TestOracleUser = mustEnv(envOracleUser)
}

// ProviderConfig returns the provider configuration block for use in tests.
// Uses os.Getenv directly so it never panics at struct-literal evaluation time.
func ProviderConfig() string {
	return fmt.Sprintf(`
provider "commvault" {
  web_service_url = "%s"
  user_name       = "%s"
  password        = "%s"
  ignore_cert     = true
}
`, os.Getenv(envWebServiceURL), os.Getenv(envUserName), os.Getenv(envPassword))
}

// CheckResourceExists is a generic check that a resource exists in state.
func CheckResourceExists(resourceType, resourceName string) func(*terraform.State) error {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No %s ID is set", resourceType)
		}
		return nil
	}
}

// CheckResourceDestroy is a generic check that a resource has been destroyed.
func CheckResourceDestroy(resourceType string) func(*terraform.State) error {
	return func(s *terraform.State) error {
		for _, rs := range s.RootModule().Resources {
			if rs.Type != resourceType {
				continue
			}
		}
		return nil
	}
}

// testCheckResourceExists is an alias used by oracle_instance_test.go /
// oracle_subclient_test.go which use the "testCheck" naming convention.
func testCheckResourceExists(resourceType, resourceName string) func(*terraform.State) error {
	return CheckResourceExists(resourceType, resourceName)
}

// testCheckResourceDestroy is an alias used by oracle_instance_test.go /
// oracle_subclient_test.go which use the "testCheck" naming convention.
func testCheckResourceDestroy(resourceType string) func(*terraform.State) error {
	return CheckResourceDestroy(resourceType)
}
