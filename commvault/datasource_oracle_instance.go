package commvault

import (
	"encoding/json"
	"fmt"
	"strconv"

	"terraform-provider-commvault/commvault/handler"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceOracleInstance() *schema.Resource {
	return &schema.Resource{
		Read: datasourceReadOracleInstance,

		Schema: map[string]*schema.Schema{
			"client_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the client where the Oracle instance is configured",
			},
			"instance_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the Oracle instance",
			},
			"instance_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "ID of the Oracle instance",
			},
			"oracle_home": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Path to the Oracle home directory",
			},
			"oracle_user": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Oracle OS user name",
			},
			"tns_admin_path": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Path to the TNS admin directory",
			},
			"block_size": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Block size for the Oracle instance",
			},
			"use_catalog_connect": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether catalog connect is used",
			},
			"archive_log_dest": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Archive log destination path",
			},
		},
	}
}

func datasourceReadOracleInstance(d *schema.ResourceData, m interface{}) error {
	clientName := d.Get("client_name").(string)
	instanceName := d.Get("instance_name").(string)

	// First, get the instance ID
	entityResp, err := handler.CvFetchOracleEntityId(clientName, instanceName, "")
	if err != nil {
		return fmt.Errorf("failed to fetch Oracle instance ID, Error %s", err)
	}

	if entityResp.InstanceId <= 0 {
		return fmt.Errorf("Oracle instance %s not found on client %s", instanceName, clientName)
	}

	instanceId := entityResp.InstanceId
	d.SetId(strconv.Itoa(instanceId))
	d.Set("instance_id", instanceId)

	// Get instance properties
	resp, err := handler.CvGetOracleInstanceProperties(strconv.Itoa(instanceId))
	if err != nil {
		return fmt.Errorf("operation [GetOracleInstanceProperties] failed, Error %s", err)
	}

	var props handler.MsgOracleInstanceFullProperties
	if len(resp.InstanceProperties) > 0 {
		var propsArr []handler.MsgOracleInstanceFullProperties
		if err := json.Unmarshal(resp.InstanceProperties, &propsArr); err == nil {
			if len(propsArr) == 0 {
				return nil
			}
			props = propsArr[0]
		} else if err := json.Unmarshal(resp.InstanceProperties, &props); err != nil {
			return fmt.Errorf("operation [GetOracleInstanceProperties] failed to parse response, Error %s", err)
		}
	}

	if props.OracleInstance != nil {
		oracle := props.OracleInstance
		if oracle.OracleHome != nil {
			d.Set("oracle_home", oracle.OracleHome)
		}
		if oracle.OracleUser != nil && oracle.OracleUser.UserName != nil {
			d.Set("oracle_user", oracle.OracleUser.UserName)
		}
		if oracle.TNSAdminPath != nil {
			d.Set("tns_admin_path", oracle.TNSAdminPath)
		}
		if oracle.BlockSize != nil {
			d.Set("block_size", oracle.BlockSize)
		}
		if oracle.UseCatalogConnect != nil {
			d.Set("use_catalog_connect", oracle.UseCatalogConnect)
		}
		if oracle.ArchiveLogDest != nil {
			d.Set("archive_log_dest", oracle.ArchiveLogDest)
		}
	}

	return nil
}
