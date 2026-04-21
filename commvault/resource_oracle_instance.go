package commvault

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"terraform-provider-commvault/commvault/handler"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceOracleInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceCreateOracleInstance,
		Read:   resourceReadOracleInstance,
		Update: resourceUpdateOracleInstance,
		Delete: resourceDeleteOracleInstance,
		Importer: &schema.ResourceImporter{
			State: resourceImportOracleInstance,
		},

		Schema: map[string]*schema.Schema{
			"client_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of the client where the Oracle instance is configured",
			},
			"client_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "ID of the client where the Oracle instance is configured",
			},
			"instance_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of the Oracle instance (SID)",
			},
			"oracle_home": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Path to the Oracle home directory",
			},
			"oracle_user": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Oracle OS user name",
			},
			"oracle_wallet_authentication": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Set to true if Oracle wallet is configured for authentication",
			},
			"sql_connect_user": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// Commvault may return empty SQL user when credential-based DB auth is used.
					// Treat empty and "/" as equivalent to avoid perpetual diffs.
					return (old == "" && new == "/") || (old == "/" && new == "")
				},
				Description: "Oracle database username for SQL connect (use '/' for OS authentication)",
			},
			"sql_connect_domain": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Oracle SID/DB TNS entry for SQL connect",
			},
			"db_connect_credential_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Credential ID of type 'Oracle' for database connection (preferred over inline credentials)",
			},
			"tns_admin_path": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Path to the TNS admin directory",
			},
			"block_size": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     1048576,
				Description: "Block size for RMAN backup operations",
			},
			"cross_check_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     600,
				Description: "RMAN crosscheck timeout in seconds",
			},
			"use_catalog_connect": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether to use RMAN recovery catalog",
			},
			"catalog_connect_credential_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Credential ID of type 'Oracle Recovery Catalog' (preferred over inline credentials)",
			},
			"catalog_connect_user": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Recovery catalog username (deprecated - use catalog_connect_credential_id instead)",
			},
			"catalog_connect_domain": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Recovery catalog DB TNS entry (deprecated - use catalog_connect_credential_id instead)",
			},
			"os_user_credential_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Credential ID of type 'Windows Account' for impersonation (Windows clients only)",
			},
			"archive_log_dest": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Archive log destination path",
			},
			"plan_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Plan ID to associate with the instance",
			}},
	}
}

// resourceImportOracleInstance resolves "clientName/instanceName" → numeric instanceId
// so that `terraform import commvault_oracle_instance.x clientName/instanceName` works.
func resourceImportOracleInstance(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	// If the ID is already numeric, treat it as a direct instance ID.
	if _, err := strconv.Atoi(d.Id()); err == nil {
		return []*schema.ResourceData{d}, nil
	}
	// Otherwise expect "clientName/instanceName"
	parts := strings.SplitN(d.Id(), "/", 2)
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return nil, fmt.Errorf("import ID must be <instanceId> or <clientName>/<instanceName>, got: %s", d.Id())
	}
	clientName := parts[0]
	instanceName := parts[1]
	entityResp, err := handler.CvFetchOracleEntityId(clientName, instanceName, "")
	if err != nil {
		return nil, fmt.Errorf("failed to resolve Oracle instance ID for %s/%s: %s", clientName, instanceName, err)
	}
	if entityResp.InstanceId <= 0 {
		return nil, fmt.Errorf("no Oracle instance found for client=%s instance=%s", clientName, instanceName)
	}
	d.SetId(strconv.Itoa(entityResp.InstanceId))
	d.Set("client_name", clientName)
	d.Set("instance_name", instanceName)
	return []*schema.ResourceData{d}, nil
}

func resourceCreateOracleInstance(d *schema.ResourceData, m interface{}) error {
	clientName := d.Get("client_name").(string)
	instanceName := d.Get("instance_name").(string)
	oracleHome := d.Get("oracle_home").(string)

	appName := "Oracle"
	applicationId := 22

	// Build oracleInstance details for the CREATE payload — Commvault validates
	// DB connectivity during creation, so credentials must be included here.
	oracleDetails := &handler.MsgOracleInstanceDetails{
		OracleHome: &oracleHome,
	}

	if v, ok := d.GetOk("oracle_user"); ok {
		oracleDetails.OracleUser = &handler.MsgOracleUser{
			UserName: handler.ToStringValue(v, true),
		}
	}

	if v, ok := d.GetOk("db_connect_credential_id"); ok {
		oracleDetails.DbConnectCredInfo = &handler.MsgCredentialInfo{
			CredentialId: handler.ToIntValue(v, true),
		}
	} else {
		sqlUser := d.Get("sql_connect_user").(string)
		if sqlUser == "" {
			sqlUser = "/"
		}
		oracleDetails.SqlConnect = &handler.MsgOracleConnect{
			UserName:   &sqlUser,
			DomainName: handler.ToStringValue(d.Get("sql_connect_domain"), true),
		}
	}

	var planEntity *handler.MsgOraclePlanEntity
	if v, ok := d.GetOk("plan_id"); ok {
		planId := v.(int)
		planEntity = &handler.MsgOraclePlanEntity{PlanID: &planId}
	}

	instanceReq := handler.MsgCreateOracleInstanceRequest{
		InstanceProperties: &handler.MsgOracleInstanceProperties{
			Instance: &handler.MsgOracleInstance{
				ClientName:    &clientName,
				InstanceName:  &instanceName,
				AppName:       &appName,
				ApplicationId: &applicationId,
			},
			OracleInstance: oracleDetails,
			PlanEntity:     planEntity,
		},
	}

	resp, err := handler.CvCreateOracleInstance(instanceReq)
	if err != nil {
		return fmt.Errorf("operation [CreateOracleInstance] failed, Error %s", err)
	}

	var createItem handler.MsgCreateResponseItem
	parsed := false
	if len(resp.Response) > 0 {
		var arr []handler.MsgCreateResponseItem
		if err := json.Unmarshal(resp.Response, &arr); err == nil && len(arr) > 0 {
			createItem = arr[0]
			parsed = true
		} else {
			var obj handler.MsgCreateResponseItem
			if err := json.Unmarshal(resp.Response, &obj); err == nil {
				createItem = obj
				parsed = true
			}
		}
	}

	if !parsed {
		if resp.ErrorCode != nil && *resp.ErrorCode != 0 {
			errMsg := ""
			if resp.ErrorMessage != nil {
				errMsg = *resp.ErrorMessage
			} else if resp.ErrorString != nil {
				errMsg = *resp.ErrorString
			}
			return fmt.Errorf("operation [CreateOracleInstance] failed, errorCode %d: %s", *resp.ErrorCode, errMsg)
		}
		if resp.Entity != nil && resp.Entity.InstanceId != nil && *resp.Entity.InstanceId > 0 {
			d.SetId(strconv.Itoa(*resp.Entity.InstanceId))
		} else {
			return fmt.Errorf("operation [CreateOracleInstance] failed, unable to get instance ID from response")
		}
	} else {
		if ec := createItem.ErrorCode; ec != nil && *ec != 0 {
			errMsg := ""
			if createItem.ErrorString != nil {
				errMsg = *createItem.ErrorString
			}
			return fmt.Errorf("operation [CreateOracleInstance] failed, errorCode %d: %s", *ec, errMsg)
		}
		if createItem.Entity == nil || createItem.Entity.InstanceId == nil || *createItem.Entity.InstanceId <= 0 {
			return fmt.Errorf("operation [CreateOracleInstance] failed, unable to get instance ID from response")
		}
		d.SetId(strconv.Itoa(*createItem.Entity.InstanceId))
	}

	// MODIFY to apply any remaining properties (block_size, tns_admin_path, etc.)
	return resourceUpdateOracleInstanceProperties(d, m, oracleHome)
}

func resourceUpdateOracleInstanceProperties(d *schema.ResourceData, m interface{}, oracleHome string) error {
	tnsAdminPath := handler.ToStringValue(d.Get("tns_admin_path"), true)
	blockSize := handler.ToIntValue(d.Get("block_size"), true)
	// Must not omit false; API interprets missing field as existing/default behavior.
	useCatalogConnect := handler.ToBooleanValue(d.Get("use_catalog_connect"), false)
	archiveLogDest := handler.ToStringValue(d.Get("archive_log_dest"), true)
	oracleWalletAuth := handler.ToBooleanValue(d.Get("oracle_wallet_authentication"), true)
	crossCheckTimeout := handler.ToIntValue(d.Get("cross_check_timeout"), true)

	oracleDetails := &handler.MsgOracleInstanceDetails{
		OracleHome:                 &oracleHome,
		TNSAdminPath:               tnsAdminPath,
		BlockSize:                  blockSize,
		UseCatalogConnect:          useCatalogConnect,
		ArchiveLogDest:             archiveLogDest,
		OracleWalletAuthentication: oracleWalletAuth,
		CrossCheckTimeout:          crossCheckTimeout,
	}

	// Build OracleUser object if oracle_user is specified
	if v, ok := d.GetOk("oracle_user"); ok {
		oracleDetails.OracleUser = &handler.MsgOracleUser{
			UserName: handler.ToStringValue(v, true),
		}
	}

	// Bug fix 1: sql_connect_user has default "/", so d.GetOk returns ok=false for that default.
	// Always send sqlConnect when no credential ID is set, so the connect string is correct.
	if v, ok := d.GetOk("db_connect_credential_id"); ok {
		oracleDetails.DbConnectCredInfo = &handler.MsgCredentialInfo{
			CredentialId: handler.ToIntValue(v, true),
		}
	} else {
		// Always send sqlConnect - even the default "/" must be explicitly set
		sqlUser := d.Get("sql_connect_user").(string)
		if sqlUser == "" {
			sqlUser = "/"
		}
		oracleDetails.SqlConnect = &handler.MsgOracleConnect{
			UserName:   &sqlUser,
			DomainName: handler.ToStringValue(d.Get("sql_connect_domain"), true),
		}
	}

	// Build CatalogConnect object if catalog is used
	if useCatalogConnect != nil && *useCatalogConnect {
		if v, ok := d.GetOk("catalog_connect_credential_id"); ok {
			oracleDetails.CatalogConnectCredInfo = &handler.MsgCredentialInfo{
				CredentialId: handler.ToIntValue(v, true),
			}
		} else if v, ok := d.GetOk("catalog_connect_user"); ok {
			oracleDetails.CatalogConnect = &handler.MsgOracleConnect{
				UserName:   handler.ToStringValue(v, true),
				DomainName: handler.ToStringValue(d.Get("catalog_connect_domain"), true),
			}
		}
	}

	// OS User credential ID (Windows clients)
	if v, ok := d.GetOk("os_user_credential_id"); ok {
		oracleDetails.OsUserCredInfo = &handler.MsgCredentialInfo{
			CredentialId: handler.ToIntValue(v, true),
		}
	}

	// Bug fix 2: the Modify API requires instance identity fields (instanceId, clientId,
	// applicationId) in the instance block, otherwise oracleHome and other fields are ignored.
	instanceId, _ := strconv.Atoi(d.Id())
	clientId := d.Get("client_id").(int)
	clientName := d.Get("client_name").(string)
	appName := "Oracle"
	applicationId := 22

	instanceIdentity := &handler.MsgOracleInstanceResp{
		InstanceId:    &instanceId,
		ClientId:      &clientId,
		ClientName:    &clientName,
		AppName:       &appName,
		ApplicationId: &applicationId,
	}

	// Bug fix 3: plan_id must be included in the modify request inside planEntity,
	// but MsgOracleInstanceFullProperties was missing PlanEntity — now fixed in OracleMsg.go.
	var planEntity *handler.MsgOraclePlanEntity
	if v, ok := d.GetOk("plan_id"); ok {
		planId := v.(int)
		planEntity = &handler.MsgOraclePlanEntity{PlanID: &planId}
	}

	modifyReq := handler.MsgModifyOracleInstanceRequest{
		InstanceProperties: &handler.MsgOracleInstanceFullProperties{
			Instance:       instanceIdentity,
			OracleInstance: oracleDetails,
			PlanEntity:     planEntity,
		},
	}

	_, err := handler.CvModifyOracleInstance(modifyReq, d.Id())
	if err != nil {
		return fmt.Errorf("operation [ModifyOracleInstance] failed, Error %s", err)
	}

	return resourceReadOracleInstance(d, m)
}

func resourceReadOracleInstance(d *schema.ResourceData, m interface{}) error {
	resp, err := handler.CvGetOracleInstanceProperties(d.Id())
	if err != nil {
		return fmt.Errorf("operation [GetOracleInstanceProperties] failed, Error %s", err)
	}

	if len(resp.InstanceProperties) > 0 {
		var propsArr []map[string]interface{}
		var propsObj map[string]interface{}
		props := map[string]interface{}{}

		if err := json.Unmarshal(resp.InstanceProperties, &propsArr); err == nil {
			if len(propsArr) == 0 {
				return nil
			}
			props = propsArr[0]
		} else if err := json.Unmarshal(resp.InstanceProperties, &propsObj); err != nil {
			return fmt.Errorf("operation [GetOracleInstanceProperties] failed to parse response, Error %s", err)
		} else {
			props = propsObj
		}

		if instance, ok := props["instance"].(map[string]interface{}); ok {
			if v, ok := instance["clientName"].(string); ok {
				d.Set("client_name", v)
			}
			if v, ok := instance["clientId"].(float64); ok {
				d.Set("client_id", int(v))
			}
			if v, ok := instance["instanceName"].(string); ok {
				d.Set("instance_name", v)
			}
		}

		if oracle, ok := props["oracleInstance"].(map[string]interface{}); ok {
			if v, ok := oracle["oracleHome"].(string); ok {
				d.Set("oracle_home", v)
			}
			if oracleUser, ok := oracle["oracleUser"].(map[string]interface{}); ok {
				if v, ok := oracleUser["userName"].(string); ok {
					d.Set("oracle_user", v)
				}
			}
			if v, ok := oracle["oracleWalletAuthentication"].(bool); ok {
				d.Set("oracle_wallet_authentication", v)
			}
			if sql, ok := oracle["sqlConnect"].(map[string]interface{}); ok {
				if v, ok := sql["userName"].(string); ok {
					d.Set("sql_connect_user", v)
				}
				if v, ok := sql["domainName"].(string); ok {
					d.Set("sql_connect_domain", v)
				}
			}
			if cred, ok := oracle["dbConnectCredInfo"].(map[string]interface{}); ok {
				if v, ok := cred["credentialId"].(float64); ok {
					d.Set("db_connect_credential_id", int(v))
				}
			}
			if v, ok := oracle["TNSAdminPath"].(string); ok {
				d.Set("tns_admin_path", v)
			}
			if v, ok := oracle["blockSize"].(float64); ok {
				d.Set("block_size", int(v))
			}
			if v, ok := oracle["crossCheckTimeout"].(float64); ok {
				d.Set("cross_check_timeout", int(v))
			}
			if v, ok := oracle["useCatalogConnect"].(bool); ok {
				d.Set("use_catalog_connect", v)
			}
			if catalog, ok := oracle["catalogConnect"].(map[string]interface{}); ok {
				if v, ok := catalog["userName"].(string); ok {
					d.Set("catalog_connect_user", v)
				}
				if v, ok := catalog["domainName"].(string); ok {
					d.Set("catalog_connect_domain", v)
				}
			}
			if cred, ok := oracle["catalogConnectCredInfo"].(map[string]interface{}); ok {
				if v, ok := cred["credentialId"].(float64); ok {
					d.Set("catalog_connect_credential_id", int(v))
				}
			}
			if cred, ok := oracle["osUserCredInfo"].(map[string]interface{}); ok {
				if v, ok := cred["credentialId"].(float64); ok {
					d.Set("os_user_credential_id", int(v))
				}
			}
			if v, ok := oracle["archiveLogDest"].(string); ok {
				d.Set("archive_log_dest", v)
			} else if arr, ok := oracle["archiveLogDest"].([]interface{}); ok && len(arr) > 0 {
				if first, ok := arr[0].(string); ok {
					d.Set("archive_log_dest", first)
				}
			}
		}

		if plan, ok := props["planEntity"].(map[string]interface{}); ok {
			if v, ok := plan["planId"].(float64); ok {
				d.Set("plan_id", int(v))
			} else if v, ok := plan["id"].(float64); ok {
				d.Set("plan_id", int(v))
			}
		}
	}

	return nil
}

func resourceUpdateOracleInstance(d *schema.ResourceData, m interface{}) error {
	if d.HasChanges("oracle_home", "oracle_user", "sql_connect_user", "sql_connect_domain", "db_connect_credential_id",
		"tns_admin_path", "block_size", "cross_check_timeout", "use_catalog_connect",
		"catalog_connect_user", "catalog_connect_domain", "catalog_connect_credential_id",
		"os_user_credential_id", "archive_log_dest", "oracle_wallet_authentication", "plan_id") {

		oracleHome := d.Get("oracle_home").(string)
		return resourceUpdateOracleInstanceProperties(d, m, oracleHome)
	}

	return resourceReadOracleInstance(d, m)
}

func resourceDeleteOracleInstance(d *schema.ResourceData, m interface{}) error {
	_, err := handler.CvDeleteOracleInstance(d.Id())
	if err != nil {
		return fmt.Errorf("operation [DeleteOracleInstance] failed, Error %s", err)
	}
	return nil
}
