package commvault

import (
	"encoding/json"
	"fmt"
	"strconv"

	"terraform-provider-commvault/commvault/handler"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceOracleSubclient() *schema.Resource {
	return &schema.Resource{
		Read: datasourceReadOracleSubclient,

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
			"subclient_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the Oracle subclient",
			},
			"subclient_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "ID of the Oracle subclient",
			},
			"instance_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "ID of the Oracle instance",
			},
			"client_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "ID of the client",
			},
			"enable_backup": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether backup is enabled for the subclient",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Description of the subclient",
			},
			"content_operation_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Subclient content operation type",
			},
			"content": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Oracle content entries used for ONLINE_SUBSET_DB backups",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oracle_content": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Oracle content payload",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"table_space": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Tablespace selected for online subset backup",
									},
								},
							},
						},
					},
				},
			},
			"backup_archive_log": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether archive logs are backed up",
			},
			"backup_sp_file": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether SP file is backed up",
			},
			"backup_control_file": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether control file is backed up",
			},
			"delete_archive_log_after_backup": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether archive logs are deleted after backup",
			},
			"backup_mode": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Oracle backup mode",
			},
			"selective_online_full": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Selective online full setting",
			},
			"data": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Oracle data option",
			},
			"lights_out_script": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Lights out script value",
			},
			"enable_table_browse": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable table browse",
			},
			"archive_delete": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Archive delete option",
			},
			"data_threshold_streams": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Data threshold streams setting",
			},
			"archive_delete_all": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Archive delete all option",
			},
			"plan_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Plan ID associated with this Oracle subclient",
			},
			"plan_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Plan name associated with this Oracle subclient",
			},
		},
	}
}

func datasourceReadOracleSubclient(d *schema.ResourceData, m interface{}) error {
	clientName := d.Get("client_name").(string)
	instanceName := d.Get("instance_name").(string)
	subclientName := d.Get("subclient_name").(string)

	// Get the subclient ID
	entityResp, err := handler.CvFetchOracleEntityId(clientName, instanceName, subclientName)
	if err != nil {
		return fmt.Errorf("failed to fetch Oracle subclient ID, Error %s", err)
	}

	if entityResp.SubclientId <= 0 {
		return fmt.Errorf("Oracle subclient %s not found on instance %s, client %s", subclientName, instanceName, clientName)
	}

	subclientId := entityResp.SubclientId
	d.SetId(strconv.Itoa(subclientId))
	d.Set("subclient_id", subclientId)

	if entityResp.InstanceId > 0 {
		d.Set("instance_id", entityResp.InstanceId)
	}
	if entityResp.ClientId > 0 {
		d.Set("client_id", entityResp.ClientId)
	}

	// Get subclient properties
	resp, err := handler.CvGetOracleSubclientProperties(strconv.Itoa(subclientId))
	if err != nil {
		return fmt.Errorf("operation [GetOracleSubclientProperties] failed, Error %s", err)
	}

	if len(resp.SubClientProperties) > 0 {
		var propsArr []map[string]interface{}
		var propsObj map[string]interface{}
		props := map[string]interface{}{}
		if err := json.Unmarshal(resp.SubClientProperties, &propsArr); err == nil {
			if len(propsArr) == 0 {
				return nil
			}
			props = propsArr[0]
		} else if err := json.Unmarshal(resp.SubClientProperties, &propsObj); err != nil {
			return fmt.Errorf("operation [GetOracleSubclientProperties] failed to parse response, Error %s", err)
		} else {
			props = propsObj
		}

		if common, ok := props["commonProperties"].(map[string]interface{}); ok {
			if v, ok := common["enableBackup"].(bool); ok {
				d.Set("enable_backup", v)
			}
			if v, ok := common["description"].(string); ok {
				d.Set("description", v)
			}
		}

		if plan, ok := props["planEntity"].(map[string]interface{}); ok {
			if v, ok := plan["planId"].(float64); ok {
				d.Set("plan_id", int(v))
			} else if v, ok := plan["id"].(float64); ok {
				d.Set("plan_id", int(v))
			}
			if v, ok := plan["planName"].(string); ok {
				d.Set("plan_name", v)
			} else if v, ok := plan["name"].(string); ok {
				d.Set("plan_name", v)
			}
		}

		if v, ok := props["contentOperationType"].(string); ok {
			d.Set("content_operation_type", v)
		}
		if contentRaw, ok := props["content"]; ok {
			d.Set("content", flattenOracleContent(contentRaw))
		}

		oracle, ok := props["oracleSubclientProp"].(map[string]interface{})
		if !ok {
			oracle, _ = props["oracleSubclient"].(map[string]interface{})
		}
		if oracle != nil {
			if parsed, ok := parseIntFromValue(oracle["backupMode"]); ok {
				d.Set("backup_mode", parsed)
			}
			if v, ok := oracle["backupArchiveLog"].(bool); ok {
				d.Set("backup_archive_log", v)
			}
			if v, ok := oracle["backupSPFile"].(bool); ok {
				d.Set("backup_sp_file", v)
			}
			if v, ok := oracle["backupControlFile"].(bool); ok {
				d.Set("backup_control_file", v)
			}
			if v, ok := oracle["deleteArchiveLogAfterBackup"].(bool); ok {
				d.Set("delete_archive_log_after_backup", v)
			}
			if v, ok := oracle["selectiveOnlineFull"].(bool); ok {
				d.Set("selective_online_full", v)
			}
			if v, ok := oracle["data"].(bool); ok {
				d.Set("data", v)
			}
			if v, ok := oracle["lightsOutScript"].(string); ok {
				d.Set("lights_out_script", v)
			}
			if v, ok := oracle["enableTableBrowse"].(bool); ok {
				d.Set("enable_table_browse", v)
			}
			if v, ok := oracle["archiveDelete"].(bool); ok {
				d.Set("archive_delete", v)
			}
			if parsed, ok := parseIntFromValue(oracle["dataThresholdStreams"]); ok {
				d.Set("data_threshold_streams", parsed)
			}
			if v, ok := oracle["archiveDeleteAll"].(bool); ok {
				d.Set("archive_delete_all", v)
			}
		}
	}

	return nil
}
