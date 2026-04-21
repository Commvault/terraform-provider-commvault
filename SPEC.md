# Commvault Oracle Terraform Provider — Implementation Specification

> **Purpose**: This document is a comprehensive specification for how the Commvault Oracle OpenAPI YAML spec (`OracleDB.yaml`) was translated into a Terraform provider. It is designed so that a future AI agent (or human developer) can replicate this pattern for **any other Commvault agent type** (e.g., SQL Server, SAP HANA, PostgreSQL, MySQL) by following the same systematic approach.

---

## Table of Contents

1. [Architecture Overview](#1-architecture-overview)
2. [Source Material: The OpenAPI YAML Spec](#2-source-material-the-openapi-yaml-spec)
3. [Translation Pipeline: YAML → Go → Terraform](#3-translation-pipeline-yaml--go--terraform)
4. [Layer 1: Message Types (OracleMsg.go)](#4-layer-1-message-types-oraclemsggo)
5. [Layer 2: API Handler Functions (OracleHandler.go)](#5-layer-2-api-handler-functions-oraclehandlergo)
6. [Layer 3: Terraform Resources](#6-layer-3-terraform-resources)
7. [Layer 4: Terraform Data Sources](#7-layer-4-terraform-data-sources)
8. [Layer 5: Provider Registration](#8-layer-5-provider-registration)
9. [Layer 6: Tests](#9-layer-6-tests)
10. [Layer 7: Documentation](#10-layer-7-documentation)
11. [Authentication & Login](#11-authentication--login)
12. [Design Decisions & Gotchas](#12-design-decisions--gotchas)
13. [API Endpoint Mapping Reference](#13-api-endpoint-mapping-reference)
14. [Step-by-Step Replication Guide](#14-step-by-step-replication-guide)
15. [File Inventory](#15-file-inventory)

---

## 1. Architecture Overview

```
┌─────────────────────────────────────────────────────────────────┐
│  OracleDB.yaml (OpenAPI 3.0 Spec)                              │
│  - Defines endpoints, request/response schemas, examples        │
└──────────────────────────┬──────────────────────────────────────┘
                           │ Manual translation
                           ▼
┌─────────────────────────────────────────────────────────────────┐
│  commvault/handler/OracleMsg.go                                 │
│  - Go struct types mirroring every JSON request/response body   │
│  - All fields are pointer types (*string, *int, *bool)          │
│  - JSON tags use `json:"fieldName,omitempty"`                   │
└──────────────────────────┬──────────────────────────────────────┘
                           │
                           ▼
┌─────────────────────────────────────────────────────────────────┐
│  commvault/handler/OracleHandler.go                             │
│  - One Go function per API operation (operationId)              │
│  - Each function: marshal → HTTP call → unmarshal → return      │
│  - Uses existing makeHttpRequestErr() utility                   │
└──────────────────────────┬──────────────────────────────────────┘
                           │
                           ▼
┌─────────────────────────────────────────────────────────────────┐
│  commvault/resource_oracle_*.go                                 │
│  commvault/datasource_oracle_*.go                               │
│  - Terraform schema definitions (TypeString, TypeInt, TypeBool) │
│  - CRUD functions that call handler functions                   │
│  - Build Msg structs from schema → call handler → set state     │
└──────────────────────────┬──────────────────────────────────────┘
                           │
                           ▼
┌─────────────────────────────────────────────────────────────────┐
│  commvault/provider.go                                          │
│  - Registers resources and data sources in ResourcesMap /       │
│    DataSourcesMap                                               │
└─────────────────────────────────────────────────────────────────┘
```

The provider uses **Terraform Plugin SDK v1** (`github.com/hashicorp/terraform-plugin-sdk`), NOT the newer Plugin Framework. This is important because it determines the schema definition style and CRUD function signatures.

---

## 2. Source Material: The OpenAPI YAML Spec

### Spec Structure

The YAML spec (`OracleDB.yaml`) is an OpenAPI 3.0 document with:

- **`paths`**: Each path maps to one or more HTTP methods (GET, POST, DELETE)
- **`operationId`**: Unique identifier for each operation — becomes the Go function name
- **`requestBody`**: JSON schema for the POST body — becomes the `Msg*Request` struct
- **`responses.200.content.application/json.schema`**: JSON schema for success response — becomes the `Msg*Response` struct
- **`examples`**: Sample payloads — invaluable for understanding field nesting and default values

### Operations Defined in the Oracle Spec

| operationId | HTTP Method | Path | Purpose | Terraform Mapping |
|---|---|---|---|---|
| `InstallOracleAgent` | POST | `/Createtask` | Install Oracle agent on server | Resource: `commvault_oracle_install_agent` |
| `AddInstance` | POST | `/instance` | Create Oracle instance | Resource: `commvault_oracle_instance` (Create) |
| `fetchInstances` | GET | `/instance?clientName=&appName=Oracle` | List instances for a client | Used internally for ID lookup |
| `GetInstanceProp` | GET | `/instance/{instanceId}` | Get instance properties | Resource: Read / Data Source: Read |
| `ModifyInstance` | POST | `/instance/{instanceId}` | Update instance properties | Resource: Update |
| `delete-instance-instanceId` | DELETE | `/instance/{instanceId}` | Delete instance | Resource: Delete |
| `InstanceDiscover` | GET | `/client/{clientId}/instance/oracle/discover` | Discover Oracle instances | Not exposed (internal utility) |
| `fetchEntityId` | GET | `/GetId?agent=Oracle&clientName=&instanceName=` | Resolve names → IDs | Used internally throughout |
| `GetOracleBackupPieces` | GET | `/oracle/instance/{instanceId}/backupPieces` | List backup pieces | Data Source: `commvault_oracle_backup_pieces` |
| `DBBrowse` | POST | `/instance/DBBrowse/{instanceId}` | Browse Oracle database | Handler exists, not exposed as resource |
| `createSubclient` | POST | `/subclient` | Create subclient | Resource: `commvault_oracle_subclient` (Create) |
| `fetchSubclients` | GET | `/subclient?clientId=&applicationId=22&instanceId=` | List subclients | Used internally |
| `fetchSubclientProp` | GET | `/subclient/{subclientId}` | Get subclient properties | Resource: Read / Data Source: Read |
| `modifySubclient` | POST | `/subclient/{subclientId}` | Update subclient | Resource: Update |
| `deleteSubclient` | DELETE | `/subclient/{subclientId}` | Delete subclient | Resource: Delete |
| `OracleBackup` | POST | `/CreateTask` | Trigger backup job | Handler exists, NOT a resource (by design) |
| `oracleDBRestore` | POST | `/createTask` | Trigger restore job | Handler exists, NOT a resource (by design) |
| `fetchRMANLogs` | GET | `/Job/{JobId}/RMANLogs` | Fetch RMAN logs for a job | Data Source: `commvault_oracle_rman_logs` |

### Why Backup/Restore Are NOT Resources

Terraform resources represent **infrastructure state** — things that are created, read, updated, and deleted. Backup and restore are **operations** (fire-and-forget jobs), not stateful objects. They were intentionally excluded from the resource model.

---

## 3. Translation Pipeline: YAML → Go → Terraform

### Step-by-Step For Each YAML Endpoint

```
YAML operationId  →  Handler function name:  Cv{PascalCase(operationId)}
YAML requestBody  →  Msg type:               Msg{PascalCase(operationId)}Request
YAML response 200 →  Msg type:               Msg{PascalCase(operationId)}Response
YAML path params   →  Function params:        func Cv...(paramName string)
YAML query params  →  URL construction:       url += "?param=" + value
```

### Naming Conventions

| YAML Concept | Go Convention | Example |
|---|---|---|
| operationId: `AddInstance` | `CvCreateOracleInstance` | Prefixed with `Cv`, Oracle-specific |
| Request body schema | `MsgCreateOracleInstanceRequest` | Prefixed with `Msg` |
| Response schema | `MsgCreateOracleInstanceResponse` | Prefixed with `Msg` |
| Nested object | `MsgOracleInstanceDetails` | Descriptive, prefixed with `Msg` |
| Terraform resource | `commvault_oracle_instance` | Snake_case, prefixed with `commvault_` |
| Resource function | `resourceOracleInstance()` | Returns `*schema.Resource` |
| CRUD functions | `resourceCreateOracleInstance` | `resource{Create|Read|Update|Delete}{Name}` |
| Data source function | `datasourceOracleInstance()` | Returns `*schema.Resource` |
| Data source read | `datasourceReadOracleInstance` | `datasourceRead{Name}` |

---

## 4. Layer 1: Message Types (OracleMsg.go)

**File**: `commvault/handler/OracleMsg.go`  
**Package**: `handler`

### Rules for Translating YAML Schemas to Go Structs

1. **Every field is a pointer type**: `*string`, `*int`, `*bool`, `*StructType`
   - This allows distinguishing "not set" (nil) from "set to zero value"
   - Terraform's `omitempty` JSON tag skips nil fields in API requests

2. **JSON tags must exactly match the API field names** (case-sensitive):
   ```go
   // YAML field: "oracleHome" → Go field: OracleHome with json:"oracleHome"
   OracleHome *string `json:"oracleHome,omitempty"`
   
   // YAML field: "TNSAdminPath" → Go field: TNSAdminPath with json:"TNSAdminPath"
   // Note: The JSON tag preserves the API's casing, even if unusual
   TNSAdminPath *string `json:"TNSAdminPath,omitempty"`
   ```

3. **Nested objects become separate struct types**:
   ```yaml
   # YAML:
   oracleInstance:
     oracleUser:
       userName: string
       domainName: string
   ```
   ```go
   // Go:
   type MsgOracleInstanceDetails struct {
       OracleUser *MsgOracleUser `json:"oracleUser,omitempty"`
   }
   type MsgOracleUser struct {
       UserName   *string `json:"userName,omitempty"`
       DomainName *string `json:"domainName,omitempty"`
   }
   ```

4. **Arrays become slices of struct types**:
   ```yaml
   # YAML:
   attributes:
     type: array
     items:
       type: object
       properties:
         name: { type: string }
         value: { type: string }
   ```
   ```go
   // Go:
   Attributes []MsgAttribute `json:"attributes,omitempty"`
   ```

5. **Reuse shared types**: `MsgIdName` (with `name`/`id` fields) is defined once in `CvGeneratedMsg.go` and reused everywhere. Don't redefine it.

6. **Credential patterns**: The API supports two auth approaches:
   - **Inline credentials** (deprecated): `sqlConnect.userName`, `sqlConnect.password`
   - **Credential IDs** (preferred): `dbConnectCredInfo.credentialId`
   
   Model BOTH in Go structs so the Terraform resource can support either.

### Struct Hierarchy Example

```
MsgCreateOracleInstanceRequest
└── InstanceProperties *MsgOracleInstanceProperties
    ├── Instance *MsgOracleInstance
    │   ├── ClientName *string
    │   ├── ClientId *int
    │   ├── InstanceName *string
    │   ├── AppName *string        (always "Oracle")
    │   └── ApplicationId *int     (always 22)
    ├── OracleInstance *MsgOracleInstanceDetails
    │   ├── OracleHome *string
    │   ├── OracleUser *MsgOracleUser
    │   │   ├── UserName *string
    │   │   └── DomainName *string
    │   ├── SqlConnect *MsgOracleConnect
    │   │   ├── UserName *string
    │   │   ├── Password *string
    │   │   ├── DomainName *string
    │   │   └── ConfirmPassword *string
    │   ├── OracleWalletAuthentication *bool
    │   ├── UseCatalogConnect *bool
    │   ├── CatalogConnect *MsgOracleConnect
    │   ├── BlockSize *int
    │   ├── CrossCheckTimeout *int
    │   ├── TNSAdminPath *string
    │   ├── ArchiveLogDest *string
    │   ├── OsUserCredInfo *MsgCredentialInfo
    │   │   ├── CredentialId *int
    │   │   └── CredentialName *string
    │   ├── DbConnectCredInfo *MsgCredentialInfo
    │   └── CatalogConnectCredInfo *MsgCredentialInfo
    └── PlanEntity *MsgIdName
        ├── Name *string
        └── Id *int
```

### Request vs Response Types

The API often returns **more fields** in GET responses than it accepts in POST/PUT requests. Create separate types when needed:

- **Request**: `MsgOracleInstance` (for POST `/instance`) — has `ClientName`, `AppName`, `ApplicationId`
- **Response list**: `MsgOracleInstanceResp` (from GET `/instance?clientName=`) — has `InstanceId`, `ClientId`, etc.
- **Response detail**: `MsgOracleInstanceFullProperties` (from GET `/instance/{id}`) — has full `OracleInstance` sub-object

---

## 5. Layer 2: API Handler Functions (OracleHandler.go)

**File**: `commvault/handler/OracleHandler.go`  
**Package**: `handler`

### Pattern for Every Handler Function

```go
func Cv{OperationName}({params}) (*Msg{Operation}Response, error) {
    // 1. Marshal request body (for POST/PUT only)
    reqBody, _ := json.Marshal(req)
    
    // 2. Build URL from environment variable + path
    url := os.Getenv("CV_CSIP") + "/{endpoint}"
    
    // 3. Get auth token from environment
    token := os.Getenv("AuthToken")
    
    // 4. Make HTTP request using shared utility
    respBody, err := makeHttpRequestErr(url, method, JSON, reqBody, JSON, token, 0)
    
    // 5. Unmarshal response
    var respObj Msg{Operation}Response
    json.Unmarshal(respBody, &respObj)
    
    // 6. Return response and error
    return &respObj, err
}
```

### Key Infrastructure

- **`os.Getenv("CV_CSIP")`**: Base URL for the Commvault REST API, set during provider configuration
- **`os.Getenv("AuthToken")`**: Authentication token, set during login
- **`makeHttpRequestErr()`**: Shared HTTP utility defined in `CommvaultClient.go` — handles TLS, headers, etc.
- **`urlEscape()`**: URL-encodes query parameter values, defined in `DataSourceHandler.go`
- **`JSON` constant**: Content type constant for JSON requests/responses

### makeHttpRequestErr Signature

```go
func makeHttpRequestErr(url string, method string, acceptType int, body []byte, 
    contentType int, token string, timeout int) ([]byte, error)
```

Parameters:
- `url`: Full URL including query parameters
- `method`: `http.MethodGet`, `http.MethodPost`, `http.MethodDelete`
- `acceptType`: `JSON` or `XML` constant
- `body`: Marshalled request body (nil for GET/DELETE)
- `contentType`: `JSON` or `XML` constant
- `token`: Auth token string
- `timeout`: Request timeout (0 = default)

### URL Construction Patterns

```go
// Simple path parameter
url := os.Getenv("CV_CSIP") + "/instance/" + instanceId

// Query parameters
url := os.Getenv("CV_CSIP") + "/instance?clientName=" + urlEscape(clientName) + "&appName=Oracle"

// Optional query parameters (conditional append)
url := os.Getenv("CV_CSIP") + "/GetId?agent=Oracle"
if clientName != "" {
    url += "&clientName=" + urlEscape(clientName)
}
```

### Oracle-Specific Constants

- **Application ID**: `22` (Oracle's numeric ID in Commvault)
- **App Name**: `"Oracle"` (string identifier)
- These are hardcoded in the handler/resource layer, not configurable

---

## 6. Layer 3: Terraform Resources

### Resource Files

| File | Resource Name | CRUD |
|---|---|---|
| `resource_oracle_instance.go` | `commvault_oracle_instance` | Create, Read, Update, Delete |
| `resource_oracle_subclient.go` | `commvault_oracle_subclient` | Create, Read, Update, Delete |
| `resource_oracle_install_agent.go` | `commvault_oracle_install_agent` | Create, Read (no-op), Delete (state-only) |

### Resource Pattern

Every resource file follows this exact structure:

```go
package commvault

import (
    "fmt"
    "strconv"
    "terraform-provider-commvault/commvault/handler"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// 1. Schema definition function
func resource{Name}() *schema.Resource {
    return &schema.Resource{
        Create: resourceCreate{Name},
        Read:   resourceRead{Name},
        Update: resourceUpdate{Name},
        Delete: resourceDelete{Name},
        Schema: map[string]*schema.Schema{
            // ... field definitions
        },
    }
}

// 2. Create function
func resourceCreate{Name}(d *schema.ResourceData, m interface{}) error {
    // a. Read values from schema
    // b. Build Msg request struct
    // c. Call handler function
    // d. Extract ID from response → d.SetId()
    // e. Optionally call Update to set additional properties
    // f. Return Read function
}

// 3. Read function
func resourceRead{Name}(d *schema.ResourceData, m interface{}) error {
    // a. Call handler GET function with d.Id()
    // b. Set schema values from response using d.Set()
}

// 4. Update function
func resourceUpdate{Name}(d *schema.ResourceData, m interface{}) error {
    // a. Check d.HasChanges() for relevant fields
    // b. Build Msg request struct
    // c. Call handler modify function
    // d. Return Read function
}

// 5. Delete function
func resourceDelete{Name}(d *schema.ResourceData, m interface{}) error {
    // a. Call handler delete function with d.Id()
}
```

### Schema Field Mapping Rules

| YAML Type | Terraform Type | Go Read Pattern |
|---|---|---|
| `string` | `schema.TypeString` | `d.Get("field").(string)` |
| `integer` | `schema.TypeInt` | `d.Get("field").(int)` |
| `boolean` | `schema.TypeBool` | `d.Get("field").(bool)` |
| `object` (with id/name) | `schema.TypeList` with nested `schema.Resource` | `d.GetOk("field")` → `[]interface{}` |
| `array` | `schema.TypeList` | List of nested resources |

### Schema Field Attributes

| Attribute | When to Use |
|---|---|
| `Required: true` | Field is mandatory for API call |
| `Optional: true` | Field has a sensible default or is truly optional |
| `Computed: true` | Field is returned by API but not set by user |
| `ForceNew: true` | Changing this field requires destroying and recreating the resource |
| `Sensitive: true` | Field contains secrets (passwords) — hidden in plan output |
| `Default: value` | Default value when user doesn't specify |
| `Description: "..."` | Documentation string |

### The Create-then-Update Pattern

The Oracle instance API has a quirk: the CREATE endpoint (`POST /instance`) accepts minimal fields, but most configuration (oracleHome, sqlConnect, blockSize, etc.) must be set via a subsequent MODIFY call (`POST /instance/{id}`).

```go
func resourceCreateOracleInstance(d *schema.ResourceData, m interface{}) error {
    // Step 1: Create with minimal fields
    resp, err := handler.CvCreateOracleInstance(instanceReq)
    d.SetId(strconv.Itoa(*resp.Response.Entity.Id))
    
    // Step 2: Update with full properties
    return resourceUpdateOracleInstanceProperties(d, m, oracleHome)
}
```

This is a common pattern in Commvault APIs — the CREATE endpoint returns an ID, then you configure the entity with a separate POST.

### ID Resolution Pattern

When the CREATE response doesn't return an entity ID directly (which happens), fall back to the `GetId` endpoint:

```go
if resp.Response != nil && resp.Response.Entity != nil && resp.Response.Entity.Id != nil {
    d.SetId(strconv.Itoa(*resp.Response.Entity.Id))
} else {
    // Fallback: look up ID by name
    entityResp, err := handler.CvFetchOracleEntityId(clientName, instanceName, "")
    d.SetId(strconv.Itoa(*entityResp.InstanceId))
}
```

### Building Msg Structs from Schema (Conditional Object Construction)

When the API expects a nested object, but the Terraform user may not have set all fields, build the object conditionally:

```go
// Only build SqlConnect if user provided sql_connect_user
if v, ok := d.GetOk("sql_connect_user"); ok {
    oracleDetails.SqlConnect = &handler.MsgOracleConnect{
        UserName:   handler.ToStringValue(v, true),
        DomainName: handler.ToStringValue(d.Get("sql_connect_domain"), true),
    }
}
```

### Credential ID Preference Pattern

When both inline credentials AND credential IDs are supported, prefer credential IDs:

```go
if v, ok := d.GetOk("db_connect_credential_id"); ok {
    // Preferred: use credential ID reference
    oracleDetails.DbConnectCredInfo = &handler.MsgCredentialInfo{
        CredentialId: handler.ToIntValue(v, true),
    }
} else if v, ok := d.GetOk("sql_connect_user"); ok {
    // Fallback: use inline credentials (deprecated by API)
    oracleDetails.SqlConnect = &handler.MsgOracleConnect{
        UserName:   handler.ToStringValue(v, true),
        DomainName: handler.ToStringValue(d.Get("sql_connect_domain"), true),
    }
}
```

### Read Function: Setting State from API Response

Always nil-check every field before calling `d.Set()`:

```go
if oracle.OracleUser != nil && oracle.OracleUser.UserName != nil {
    d.Set("oracle_user", oracle.OracleUser.UserName)
}
```

For nested objects that become flat Terraform fields:

```go
// API returns: oracleInstance.sqlConnect.userName
// Terraform field: sql_connect_user (flat)
if oracle.SqlConnect != nil && oracle.SqlConnect.UserName != nil {
    d.Set("sql_connect_user", oracle.SqlConnect.UserName)
    if oracle.SqlConnect.DomainName != nil {
        d.Set("sql_connect_domain", oracle.SqlConnect.DomainName)
    }
}
```

### Install Agent Resource (Special Case)

The install agent resource is a **one-shot operation** — it triggers an installation job but there's no corresponding GET/UPDATE/DELETE for the installed agent. This maps to:

```go
Create: resourceCreateOracleInstallAgent,    // POST /Createtask → returns taskId
Read:   resourceReadOracleInstallAgent,      // No-op (returns nil)
Delete: resourceDeleteOracleInstallAgent,    // Just clears state: d.SetId("")
// No Update function — all fields are ForceNew: true
```

---

## 7. Layer 4: Terraform Data Sources

### Data Source Files

| File | Data Source Name | Purpose |
|---|---|---|
| `datasource_oracle_instance.go` | `commvault_oracle_instance` | Look up instance properties by client/instance name |
| `datasource_oracle_subclient.go` | `commvault_oracle_subclient` | Look up subclient properties by name |
| `datasource_oracle_backup_pieces.go` | `commvault_oracle_backup_pieces` | List backup pieces for an instance |
| `datasource_oracle_rman_logs.go` | `commvault_oracle_rman_logs` | Fetch RMAN logs for a job |

### Data Source Pattern

```go
func datasource{Name}() *schema.Resource {
    return &schema.Resource{
        Read: datasourceRead{Name},
        Schema: map[string]*schema.Schema{
            // Input fields: Required: true
            // Output fields: Computed: true
        },
    }
}

func datasourceRead{Name}(d *schema.ResourceData, m interface{}) error {
    // 1. Read input parameters from schema
    // 2. Resolve names to IDs (using CvFetchOracleEntityId)
    // 3. Call handler GET function
    // 4. Set computed fields using d.Set()
    // 5. Set resource ID using d.SetId()
}
```

### Key Difference from Resources

- Data sources are **read-only** — they only have a `Read` function
- Input fields are `Required: true` (user provides them)
- Output fields are `Computed: true` (API returns them)
- Data sources must set `d.SetId()` to a non-empty string or Terraform considers it an error

---

## 8. Layer 5: Provider Registration

**File**: `commvault/provider.go`

Add entries to the `ResourcesMap` and `DataSourcesMap` in the `Provider()` function:

```go
ResourcesMap: map[string]*schema.Resource{
    // ... existing resources ...
    "commvault_oracle_instance":       resourceOracleInstance(),
    "commvault_oracle_subclient":      resourceOracleSubclient(),
    "commvault_oracle_install_agent":  resourceOracleInstallAgent(),
},
DataSourcesMap: map[string]*schema.Resource{
    // ... existing data sources ...
    "commvault_oracle_instance":       datasourceOracleInstance(),
    "commvault_oracle_subclient":      datasourceOracleSubclient(),
    "commvault_oracle_backup_pieces":  datasourceOracleBackupPieces(),
    "commvault_oracle_rman_logs":      datasourceOracleRMANLogs(),
},
```

**Important**: Resource names and data source names CAN overlap (e.g., both `commvault_oracle_instance`). Terraform distinguishes them by context (`resource` block vs `data` block in HCL).

---

## 9. Layer 6: Tests

**Directory**: `commvault/test/`

### Test Infrastructure

| File | Purpose |
|---|---|
| `helpers_test.go` | Provider setup, env var reading, shared check functions |
| `instance_test.go` | Oracle instance CRUD acceptance tests |
| `subclient_test.go` | Oracle subclient CRUD acceptance tests |
| `install_agent_test.go` | Install agent test (skipped by default) |
| `.env.example` | Template for required environment variables |
| `run-tests.ps1` | PowerShell script to run tests |
| `README.md` | Test documentation |

### Environment Variables (ALL required, NO defaults)

```
CV_TEST_WEB_SERVICE_URL   # e.g., https://server.example.com/webconsole/api
CV_TEST_USER_NAME         # Commvault admin username
CV_TEST_PASSWORD          # Commvault admin password
CV_TEST_CLIENT_NAME       # Client with Oracle agent installed
CV_TEST_INSTANCE_NAME     # Oracle SID on the client
CV_TEST_ORACLE_HOME       # Oracle home path on the client
CV_TEST_ORACLE_USER       # Oracle OS user (e.g., "oracle")
```

### Test Helper Pattern

```go
// getEnvOrFail panics if env var is not set — forces explicit configuration
func getEnvOrFail(key string) string {
    value := os.Getenv(key)
    if value == "" {
        panic(fmt.Sprintf("Required environment variable %s is not set.", key))
    }
    return value
}

// Package-level vars read at init time
var (
    TestWebServiceURL = getEnvOrFail("CV_TEST_WEB_SERVICE_URL")
    TestClientName    = getEnvOrFail("CV_TEST_CLIENT_NAME")
    // ...
)
```

### Provider Setup for Tests

```go
var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
    testAccProvider = commvault.Provider()
    testAccProviders = map[string]terraform.ResourceProvider{
        "commvault": testAccProvider,
    }
}
```

### Acceptance Test Pattern

```go
func TestAccOracleInstance_basic(t *testing.T) {
    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },  // Skips if TF_ACC not set
        Providers:    testAccProviders,
        CheckDestroy: CheckResourceDestroy("commvault_oracle_instance"),
        Steps: []resource.TestStep{
            {
                Config: testAccOracleInstanceConfigBasic(),  // HCL config string
                Check: resource.ComposeTestCheckFunc(
                    CheckResourceExists("Oracle Instance", "commvault_oracle_instance.test"),
                    resource.TestCheckResourceAttr("commvault_oracle_instance.test", "client_name", TestClientName),
                ),
            },
        },
    })
}
```

### Test Config Functions

Test configs are generated via `fmt.Sprintf` with environment variable values:

```go
func testAccOracleInstanceConfigBasic() string {
    return fmt.Sprintf(`%s
resource "commvault_oracle_instance" "test" {
  client_name   = "%s"
  instance_name = "%s"
  oracle_home   = "%s"
  oracle_user   = "%s"
  block_size    = 65536
}
`, ProviderConfig(), TestClientName, TestInstanceName, TestOracleHome, TestOracleUser)
}
```

### Running Tests

```powershell
# Set environment variables first
$env:TF_ACC = "1"
$env:CV_TEST_WEB_SERVICE_URL = "https://server.example.com/webconsole/api"
# ... set all CV_TEST_* vars ...

# Run from project root
cd c:\Users\maheshp\terraform-provider-commvault
go test ./commvault/test/ -v -timeout 30m
```

---

## 10. Layer 7: Documentation

### Documentation Files

| File | Purpose |
|---|---|
| `docs/resources/oracle_instance.md` | Resource: Oracle Instance |
| `docs/resources/oracle_subclient.md` | Resource: Oracle Subclient |
| `docs/resources/oracle_install_agent.md` | Resource: Install Agent |
| `docs/data-sources/commvault_oracle_instance.md` | Data Source: Oracle Instance |
| `docs/data-sources/commvault_oracle_subclient.md` | Data Source: Oracle Subclient |
| `docs/data-sources/commvault_oracle_backup_pieces.md` | Data Source: Backup Pieces |
| `docs/data-sources/commvault_oracle_rman_logs.md` | Data Source: RMAN Logs |
| `docs/index.md` | Provider overview (updated with Oracle entries) |

### Doc Structure

Each doc includes:
1. **Title and description**
2. **Example Usage** (complete HCL block)
3. **Argument Reference** (Required/Optional fields with descriptions)
4. **Attribute Reference** (Computed fields)

---

## 11. Authentication & Login

### Login Flow

The provider authenticates during configuration via `LoginWithProviderCredentials()`:

1. User provides `web_service_url`, `user_name`, `password` in provider config
2. Provider calls `POST /login` with XML body containing base64-encoded password
3. Response returns an auth token
4. Token is stored in `os.Setenv("AuthToken", token)`
5. All subsequent API calls read `os.Getenv("AuthToken")`

### Password Encoding

The Commvault API expects passwords to be **base64 encoded** before sending:

```go
encodedPassword := base64.StdEncoding.EncodeToString([]byte(password))
```

### Login Endpoint Details

```
POST /login
Content-Type: application/xml
Accept: application/xml

Body: <DM2ContentIndexing_CheckCredentialReq username="admin" password="BASE64ENCODED" timeOut="10000" />

Response: <DM2ContentIndexing_CheckCredentialResp token="QSDK_TOKEN_STRING" ... />
```

### Environment Variables Used at Runtime

| Variable | Set By | Used By |
|---|---|---|
| `CV_CSIP` | Provider configure function | All handler functions (API base URL) |
| `AuthToken` | Login handler | All handler functions (auth header) |
| `CV_LOGGING` | User (optional) | `LogEntry()` for debug logging |

---

## 12. Design Decisions & Gotchas

### 1. Pointer Types Everywhere

All Msg struct fields use pointer types (`*string`, `*int`, `*bool`) because:
- `omitempty` JSON tag only works with pointer types for zero-value distinction
- A nil `*int` is omitted from JSON; a zero `int` is serialized as `0`
- The API treats absent fields differently from zero-value fields

### 2. ToBooleanValue Accepts Strings

The existing `handler.ToBooleanValue()` function expects a **string** argument (`"true"/"false"`), not a Go bool. This is because the Terraform SDK v1 internally stores booleans as strings. When reading from `d.Get()`, you need to be aware of this.

```go
// handler.ToBooleanValue signature:
func ToBooleanValue(val interface{}, omitempty bool) *bool {
    tmp := val.(string)  // Expects string "true" or "false"
    // ...
}
```

### 3. Application ID 22

Oracle's application ID in Commvault is always `22`. This is hardcoded:

```go
applicationId := 22
appName := "Oracle"
```

Other agents have different IDs (e.g., SQL Server = 81, File System = 33). When replicating for another agent, find the correct `applicationId` from the API spec or Commvault documentation.

### 4. TNSAdminPath Casing

The API uses `TNSAdminPath` (all caps TNS), not `tnsAdminPath`. The Go struct and JSON tag must match:

```go
TNSAdminPath *string `json:"TNSAdminPath,omitempty"`  // NOT tnsAdminPath
```

### 5. Create Returns ID in Different Places

Depending on the endpoint, the entity ID may be in:
- `response.entity.instanceId` (from POST /instance)
- `response.entity.id` (generic)
- Only obtainable via a follow-up `GET /GetId?agent=Oracle&clientName=...` call

Always implement the fallback `GetId` lookup.

### 6. POST for Updates (Not PUT)

The Commvault API uses `POST` (not `PUT`) for update operations:
- Create: `POST /instance` (no ID in URL)
- Update: `POST /instance/{instanceId}` (ID in URL)
- Delete: `DELETE /instance/{instanceId}`

### 7. Subclient API Uses applicationId=22 Filter

When listing subclients, the API requires `applicationId=22` as a query parameter to filter for Oracle:

```go
url := "/subclient?clientId=" + clientId + "&applicationId=22&instanceId=" + instanceId
```

### 8. Install Agent Uses /Createtask (Capital C)

The install agent endpoint is `/Createtask` (capital C), while backup/restore use `/CreateTask` (capital T). Be careful with casing.

---

## 13. API Endpoint Mapping Reference

### Instance CRUD

```
CREATE:  POST   /instance                              → CvCreateOracleInstance
READ:    GET    /instance/{instanceId}                 → CvGetOracleInstanceProperties
UPDATE:  POST   /instance/{instanceId}                 → CvModifyOracleInstance
DELETE:  DELETE  /instance/{instanceId}                 → CvDeleteOracleInstance
LIST:    GET    /instance?clientName=X&appName=Oracle  → CvFetchOracleInstances
```

### Subclient CRUD

```
CREATE:  POST   /subclient                                                    → CvCreateOracleSubclient
READ:    GET    /subclient/{subclientId}                                      → CvGetOracleSubclientProperties
UPDATE:  POST   /subclient/{subclientId}                                      → CvModifyOracleSubclient
DELETE:  DELETE  /subclient/{subclientId}                                      → CvDeleteOracleSubclient
LIST:    GET    /subclient?clientId=X&applicationId=22&instanceId=Y           → CvFetchOracleSubclients
```

### Utility Endpoints

```
LOOKUP:    GET  /GetId?agent=Oracle&clientName=X&instanceName=Y&subclient=Z  → CvFetchOracleEntityId
DISCOVER:  GET  /client/{clientId}/instance/oracle/discover                   → CvOracleInstanceDiscovery
BROWSE:    POST /instance/DBBrowse/{instanceId}                               → CvBrowseOracleDB
PIECES:    GET  /oracle/instance/{instanceId}/backupPieces                    → CvGetOracleBackupPieces
RMAN:      GET  /Job/{jobId}/RMANLogs                                        → CvFetchRMANLogs
```

### Operations (Not Resources)

```
INSTALL:  POST  /Createtask   → CvInstallOracleAgent
BACKUP:   POST  /CreateTask   → CvOracleBackup    (handler only, not a resource)
RESTORE:  POST  /createTask   → CvOracleRestore   (handler only, not a resource)
```

---

## 14. Step-by-Step Replication Guide

To add support for a new Commvault agent (e.g., "PostgreSQL"), follow these steps:

### Step 1: Obtain the OpenAPI Spec

Get the YAML spec for the target agent from Commvault. It will follow the same structure as `OracleDB.yaml`.

### Step 2: Identify Operations

List all `operationId` values. Classify each as:
- **CRUD** (create/read/update/delete entity) → Resource
- **Read-only query** → Data Source
- **Fire-and-forget operation** → Handler only (not a resource)

### Step 3: Create Message Types

**File**: `commvault/handler/{Agent}Msg.go`

For each request/response schema in the YAML:
1. Create a Go struct with pointer fields
2. Match JSON tags exactly to API field names
3. Use `omitempty` on all fields
4. Reuse `MsgIdName` and other shared types from `CvGeneratedMsg.go`
5. Create separate Request and Response types

### Step 4: Create Handler Functions

**File**: `commvault/handler/{Agent}Handler.go`

For each operation:
1. Create a function `Cv{OperationName}()`
2. Follow the marshal → URL → HTTP → unmarshal pattern
3. Use `makeHttpRequestErr()` for all HTTP calls
4. Use `os.Getenv("CV_CSIP")` for base URL, `os.Getenv("AuthToken")` for auth

### Step 5: Create Terraform Resources

**File**: `commvault/resource_{agent}_{entity}.go`

For each CRUD entity:
1. Define schema with proper types and attributes
2. Implement Create, Read, Update, Delete functions
3. Use handler functions for API calls
4. Handle ID resolution with fallback to GetId endpoint
5. Use the create-then-update pattern if the API requires it

### Step 6: Create Terraform Data Sources

**File**: `commvault/datasource_{agent}_{entity}.go`

For each read-only query:
1. Define schema with Required inputs and Computed outputs
2. Implement Read function
3. Resolve names to IDs, then fetch properties

### Step 7: Register in Provider

**File**: `commvault/provider.go`

Add entries to `ResourcesMap` and `DataSourcesMap`.

### Step 8: Create Tests

**Directory**: `commvault/test/`

1. Add env vars to `helpers_test.go` using `getEnvOrFail()`
2. Create `{entity}_test.go` with acceptance tests
3. Use `ProviderConfig()` for provider block in test configs
4. Update `.env.example` and `run-tests.ps1`

### Step 9: Create Documentation

**Directory**: `docs/resources/` and `docs/data-sources/`

1. Create markdown docs for each resource and data source
2. Update `docs/index.md` with new entries

### Step 10: Build and Verify

```powershell
cd c:\Users\maheshp\terraform-provider-commvault
go build -o terraform-provider-commvault.exe
```

---

## 15. File Inventory

### Handler Layer (commvault/handler/)

| File | Lines | Purpose |
|---|---|---|
| `OracleMsg.go` | 524 | All Go struct types for Oracle API request/response bodies |
| `OracleHandler.go` | 233 | All Go functions that make Oracle API HTTP calls |
| `CommvaultClient.go` | 242 | Shared HTTP utilities: `makeHttpRequestErr`, `ToStringValue`, `ToIntValue`, `ToBooleanValue` |
| `CvGeneratedMsg.go` | ~300 | Shared types: `MsgIdName`, etc. |
| `DataSourceHandler.go` | ~230 | Shared utilities: `urlEscape`, data source handlers |
| `LoginHandler.go` | ~98 | Login functions: `GenerateAuthToken`, `LoginWithProviderCredentials` |

### Resource Layer (commvault/)

| File | Lines | Purpose |
|---|---|---|
| `resource_oracle_instance.go` | 335 | Oracle instance resource (full CRUD) |
| `resource_oracle_subclient.go` | 353 | Oracle subclient resource (full CRUD) |
| `resource_oracle_install_agent.go` | 312 | Oracle agent installation resource (create-only) |

### Data Source Layer (commvault/)

| File | Lines | Purpose |
|---|---|---|
| `datasource_oracle_instance.go` | 113 | Oracle instance data source |
| `datasource_oracle_subclient.go` | 143 | Oracle subclient data source |
| `datasource_oracle_backup_pieces.go` | 126 | Oracle backup pieces data source |
| `datasource_oracle_rman_logs.go` | 54 | RMAN logs data source |

### Test Layer (commvault/test/)

| File | Lines | Purpose |
|---|---|---|
| `helpers_test.go` | 96 | Test infrastructure, env vars, provider setup |
| `instance_test.go` | 103 | 3 acceptance tests for Oracle instance |
| `subclient_test.go` | 38 | 1 acceptance test for Oracle subclient |
| `install_agent_test.go` | 42 | 1 acceptance test (skipped by default) |
| `.env.example` | — | Template for test environment variables |
| `run-tests.ps1` | — | PowerShell script to run tests |
| `README.md` | — | Test documentation |

### Documentation Layer (docs/)

| File | Purpose |
|---|---|
| `docs/index.md` | Provider overview (updated) |
| `docs/resources/oracle_instance.md` | Oracle instance resource docs |
| `docs/resources/oracle_subclient.md` | Oracle subclient resource docs |
| `docs/resources/oracle_install_agent.md` | Install agent resource docs |
| `docs/data-sources/commvault_oracle_instance.md` | Oracle instance data source docs |
| `docs/data-sources/commvault_oracle_subclient.md` | Oracle subclient data source docs |
| `docs/data-sources/commvault_oracle_backup_pieces.md` | Backup pieces data source docs |
| `docs/data-sources/commvault_oracle_rman_logs.md` | RMAN logs data source docs |

---

## Appendix A: Shared Type Reference

### MsgIdName (from CvGeneratedMsg.go)

```go
type MsgIdName struct {
    Name *string `json:"name,omitempty"`
    Id   *int    `json:"id,omitempty"`
}
```

Used for: plan references, storage policy references, entity references.

### Helper Functions (from CommvaultClient.go)

```go
func ToStringValue(val interface{}, omitempty bool) *string  // Convert interface{} to *string
func ToIntValue(val interface{}, omitempty bool) *int         // Convert interface{} to *int
func ToBooleanValue(val interface{}, omitempty bool) *bool    // Convert string "true"/"false" to *bool
func ToLongValue(val interface{}, omitempty bool) *int64      // Convert interface{} to *int64
```

The `omitempty` parameter controls whether zero values return nil:
- `omitempty=true`: empty string → nil, 0 → nil
- `omitempty=false`: always returns a pointer to the value

---

## Appendix B: YAML-to-Go Type Mapping Quick Reference

| YAML Type | YAML Format | Go Type |
|---|---|---|
| `string` | — | `*string` |
| `integer` | — | `*int` |
| `integer` | int64 | `*int64` |
| `boolean` | — | `*bool` |
| `number` | float | `*float32` |
| `number` | double | `*float64` |
| `object` | (named) | `*MsgSomeName` |
| `array` of objects | — | `[]MsgSomeName` |
| `array` of strings | — | `[]string` |
| `array` of integers | — | `[]int` |

---

## Appendix C: Terraform Resource Lifecycle for Oracle Instance

```
terraform plan
  └── Read (check current state)

terraform apply
  ├── Create
  │   ├── POST /instance (minimal: clientName, instanceName, appName, applicationId)
  │   ├── Extract instanceId from response (or fallback to GET /GetId)
  │   ├── POST /instance/{instanceId} (full properties: oracleHome, sqlConnect, etc.)
  │   └── GET /instance/{instanceId} (read back final state)
  └── Update (if resource already exists)
      ├── Check d.HasChanges() for modified fields
      ├── POST /instance/{instanceId} (modified properties only)
      └── GET /instance/{instanceId} (read back final state)

terraform destroy
  └── DELETE /instance/{instanceId}
```
