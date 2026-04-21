package commvault

import (
	"encoding/json"
	"fmt"
	"strconv"

	"terraform-provider-commvault/commvault/handler"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceOracleSubclient() *schema.Resource {
	return &schema.Resource{
		Create: resourceCreateOracleSubclient,
		Read:   resourceReadOracleSubclient,
		Update: resourceUpdateOracleSubclient,
		Delete: resourceDeleteOracleSubclient,

		Schema: map[string]*schema.Schema{
			"subclient_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of the Oracle subclient",
			},
			"client_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of the client where the Oracle instance is configured",
			},
			"instance_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of the Oracle instance",
			},
			"instance_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "ID of the Oracle instance",
			},
			"client_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "ID of the client",
			},
			"storage_policy": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Storage policy for data backup",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Name of the storage policy",
						},
						"id": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "ID of the storage policy",
						},
					},
				},
			},
			"log_storage_policy": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Storage policy for log backup",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Name of the storage policy",
						},
						"id": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "ID of the storage policy",
						},
					},
				},
			},
			"enable_backup": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Enable or disable backup for the subclient",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of the subclient",
			},
			"content_operation_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Subclient content operation type (for example: ADD, OVERWRITE)",
			},
			"content": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "Oracle content entries used for ONLINE_SUBSET_DB backups",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oracle_content": {
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Description: "Oracle content payload",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"table_space": {
										Type:        schema.TypeString,
										Optional:    true,
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
				Optional:    true,
				Default:     true,
				Description: "Backup archive logs",
			},
			"backup_sp_file": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Backup SP file",
			},
			"backup_control_file": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Backup control file",
			},
			"delete_archive_log_after_backup": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Delete archive logs after backup",
			},
			"plan_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Plan ID associated with this Oracle subclient",
			},
			"backup_mode": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Oracle backup mode enum (for example: 0, 1, 2)",
			},
			"selective_online_full": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Enable selective online full backup behavior",
			},
			"data": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Include data files for Oracle subclient operations",
			},
			"lights_out_script": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Lights out script value from Oracle subclient properties",
			},
			"enable_table_browse": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Enable table-level browse for Oracle subclient",
			},
			"archive_delete": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Enable archive delete behavior",
			},
			"data_threshold_streams": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Data threshold streams value for Oracle subclient",
			},
			"archive_delete_all": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Enable archive delete all behavior",
			},
			"plan_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Plan name associated with this Oracle subclient",
			},
		},
	}
}

func resourceCreateOracleSubclient(d *schema.ResourceData, m interface{}) error {
	subclientName := d.Get("subclient_name").(string)
	clientName := d.Get("client_name").(string)
	instanceName := d.Get("instance_name").(string)
	appName := "Oracle"

	// Get instance ID and client ID if not provided
	var instanceId *int
	var clientId *int

	if val, ok := d.GetOk("instance_id"); ok {
		id := val.(int)
		instanceId = &id
	}
	if val, ok := d.GetOk("client_id"); ok {
		id := val.(int)
		clientId = &id
	}

	// If instance_id or client_id is not provided, fetch them
	if instanceId == nil || clientId == nil {
		entityResp, err := handler.CvFetchOracleEntityId(clientName, instanceName, "")
		if err != nil {
			return fmt.Errorf("failed to fetch Oracle entity IDs, Error %s", err)
		}
		if instanceId == nil && entityResp.InstanceId > 0 {
			id := entityResp.InstanceId
			instanceId = &id
		}
		if clientId == nil && entityResp.ClientId > 0 {
			id := entityResp.ClientId
			clientId = &id
		}
	}

	// Build storage policy
	var storagePolicyEntity *handler.MsgIdName
	if val, ok := d.GetOk("storage_policy"); ok {
		storagePolicyEntity = buildOracleIdName(val.([]interface{}))
	}

	var logStoragePolicyEntity *handler.MsgIdName
	if val, ok := d.GetOk("log_storage_policy"); ok {
		logStoragePolicyEntity = buildOracleIdName(val.([]interface{}))
	}

	enableBackup := d.Get("enable_backup").(bool)
	description := handler.ToStringValue(d.Get("description"), true)
	contentOperationType := "ADD"
	if v, ok := d.GetOk("content_operation_type"); ok {
		contentOperationType = v.(string)
	}
	content := buildOracleContentRaw(d.Get("content").([]interface{}))
	var backupMode *int
	if v, ok := d.GetOkExists("backup_mode"); ok {
		tmp := v.(int)
		backupMode = &tmp
	}
	backupArchiveLog := boolPtrFromResource(d, "backup_archive_log")
	backupSPFile := boolPtrFromResource(d, "backup_sp_file")
	backupControlFile := boolPtrFromResource(d, "backup_control_file")
	deleteArchiveLog := boolPtrFromResource(d, "delete_archive_log_after_backup")
	selectiveOnlineFull := boolPtrFromResource(d, "selective_online_full")
	data := boolPtrFromResource(d, "data")
	lightsOutScript := stringPtrFromResource(d, "lights_out_script")
	enableTableBrowse := boolPtrFromResource(d, "enable_table_browse")
	archiveDelete := boolPtrFromResource(d, "archive_delete")
	dataThresholdStreams := intPtrFromResource(d, "data_threshold_streams")
	archiveDeleteAll := boolPtrFromResource(d, "archive_delete_all")
	var planEntity *handler.MsgOraclePlanEntity
	if v, ok := d.GetOkExists("plan_id"); ok {
		planID := v.(int)
		planEntity = &handler.MsgOraclePlanEntity{PlanID: &planID}
		if planName := stringPtrFromResource(d, "plan_name"); planName != nil {
			planEntity.Name = planName
		}
	} else if planName := stringPtrFromResource(d, "plan_name"); planName != nil {
		planEntity = &handler.MsgOraclePlanEntity{Name: planName}
	}

	req := handler.MsgCreateOracleSubclientRequest{
		SubClientProperties: &handler.MsgOracleSubclientProperties{
			ContentOperationType: &contentOperationType,
			Content:              content,
			SubClientEntity: &handler.MsgOracleSubclientEntity{
				SubclientName: &subclientName,
				ClientName:    &clientName,
				InstanceName:  &instanceName,
				AppName:       &appName,
				InstanceId:    instanceId,
				ClientId:      clientId,
			},
			CommonProperties: &handler.MsgSubclientCommonProperties{
				StorageDevice: &handler.MsgStorageDevice{
					DataBackupStoragePolicy: storagePolicyEntity,
					LogBackupStoragePolicy:  logStoragePolicyEntity,
				},
				EnableBackup: &enableBackup,
				Description:  description,
			},
			OracleSubclientProp: &handler.MsgOracleSubclientDetails{
				BackupMode:                  backupMode,
				BackupArchiveLog:            backupArchiveLog,
				BackupSPFile:                backupSPFile,
				BackupControlFile:           backupControlFile,
				DeleteArchiveLogAfterBackup: deleteArchiveLog,
				SelectiveOnlineFull:         selectiveOnlineFull,
				Data:                        data,
				LightsOutScript:             lightsOutScript,
				EnableTableBrowse:           enableTableBrowse,
				ArchiveDelete:               archiveDelete,
				DataThresholdStreams:        dataThresholdStreams,
				ArchiveDeleteAll:            archiveDeleteAll,
			},
			PlanEntity: planEntity,
		},
	}

	resp, err := handler.CvCreateOracleSubclient(req)
	if err != nil {
		return fmt.Errorf("operation [CreateOracleSubclient] failed, Error %s", err)
	}

	if resp.Response != nil && resp.Response.Entity != nil && resp.Response.Entity.Id != nil {
		d.SetId(strconv.Itoa(*resp.Response.Entity.Id))
	} else {
		// Try to get the subclient ID by looking it up
		entityResp, err := handler.CvFetchOracleEntityId(clientName, instanceName, subclientName)
		if err != nil {
			return fmt.Errorf("operation [CreateOracleSubclient] failed to get subclient ID, Error %s", err)
		}
		if entityResp.SubclientId > 0 {
			d.SetId(strconv.Itoa(entityResp.SubclientId))
		} else {
			return fmt.Errorf("operation [CreateOracleSubclient] failed, unable to get subclient ID")
		}
	}

	// Set the computed values
	if instanceId != nil {
		d.Set("instance_id", *instanceId)
	}
	if clientId != nil {
		d.Set("client_id", *clientId)
	}

	return resourceReadOracleSubclient(d, m)
}

func resourceReadOracleSubclient(d *schema.ResourceData, m interface{}) error {
	resp, err := handler.CvGetOracleSubclientProperties(d.Id())
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

		if entity, ok := props["subClientEntity"].(map[string]interface{}); ok {
			if v, ok := entity["subclientName"].(string); ok {
				d.Set("subclient_name", v)
			}
			if v, ok := entity["clientName"].(string); ok {
				d.Set("client_name", v)
			}
			if v, ok := entity["instanceName"].(string); ok {
				d.Set("instance_name", v)
			}
			if v, ok := entity["instanceId"].(float64); ok {
				d.Set("instance_id", int(v))
			}
			if v, ok := entity["clientId"].(float64); ok {
				d.Set("client_id", int(v))
			}
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

		if v, ok := props["contentOperationType"].(string); ok {
			d.Set("content_operation_type", v)
		}
		if contentRaw, ok := props["content"]; ok {
			d.Set("content", flattenOracleContent(contentRaw))
		}
	}

	return nil
}

func resourceUpdateOracleSubclient(d *schema.ResourceData, m interface{}) error {
	if d.HasChanges("storage_policy", "log_storage_policy", "enable_backup", "description", "plan_id", "plan_name", "content_operation_type", "content", "backup_mode", "backup_archive_log", "backup_sp_file", "backup_control_file", "delete_archive_log_after_backup", "selective_online_full", "data", "lights_out_script", "enable_table_browse", "archive_delete", "data_threshold_streams", "archive_delete_all") {
		// Commvault expects a minimal payload for backup-toggle updates.
		if d.HasChange("enable_backup") && !d.HasChanges("storage_policy", "log_storage_policy", "description", "content_operation_type", "content", "backup_mode", "backup_archive_log", "backup_sp_file", "backup_control_file", "delete_archive_log_after_backup", "selective_online_full", "data", "lights_out_script", "enable_table_browse", "archive_delete", "data_threshold_streams", "archive_delete_all") {
			enableBackup := boolPtrFromResource(d, "enable_backup")
			var planEntity *handler.MsgOraclePlanEntity
			if v, ok := d.GetOkExists("plan_id"); ok {
				planID := v.(int)
				planEntity = &handler.MsgOraclePlanEntity{PlanID: &planID}
			}
			if planName := stringPtrFromResource(d, "plan_name"); planName != nil {
				if planEntity == nil {
					planEntity = &handler.MsgOraclePlanEntity{}
				}
				planEntity.Name = planName
			}

			req := handler.MsgModifyOracleSubclientRequest{
				SubClientProperties: &handler.MsgOracleSubclientFullProperties{
					CommonProperties: &handler.MsgSubclientCommonProperties{
						EnableBackup: enableBackup,
					},
					PlanEntity: planEntity,
				},
			}

			_, err := handler.CvModifyOracleSubclient(req, d.Id())
			if err != nil {
				return fmt.Errorf("operation [ModifyOracleSubclient] failed, Error %s", err)
			}

			return resourceReadOracleSubclient(d, m)
		}

		var storagePolicyEntity *handler.MsgIdName
		if val, ok := d.GetOk("storage_policy"); ok {
			storagePolicyEntity = buildOracleIdName(val.([]interface{}))
		}

		var logStoragePolicyEntity *handler.MsgIdName
		if val, ok := d.GetOk("log_storage_policy"); ok {
			logStoragePolicyEntity = buildOracleIdName(val.([]interface{}))
		}

		enableBackup := boolPtrFromResource(d, "enable_backup")
		description := handler.ToStringValue(d.Get("description"), true)
		contentOperationType := "ADD"
		if v, ok := d.GetOk("content_operation_type"); ok {
			contentOperationType = v.(string)
		}
		content := buildOracleContentRaw(d.Get("content").([]interface{}))
		var backupMode *int
		if v, ok := d.GetOkExists("backup_mode"); ok {
			tmp := v.(int)
			backupMode = &tmp
		}
		backupArchiveLog := boolPtrFromResource(d, "backup_archive_log")
		backupSPFile := boolPtrFromResource(d, "backup_sp_file")
		backupControlFile := boolPtrFromResource(d, "backup_control_file")
		deleteArchiveLog := boolPtrFromResource(d, "delete_archive_log_after_backup")
		selectiveOnlineFull := boolPtrFromResource(d, "selective_online_full")
		data := boolPtrFromResource(d, "data")
		lightsOutScript := stringPtrFromResource(d, "lights_out_script")
		enableTableBrowse := boolPtrFromResource(d, "enable_table_browse")
		archiveDelete := boolPtrFromResource(d, "archive_delete")
		dataThresholdStreams := intPtrFromResource(d, "data_threshold_streams")
		archiveDeleteAll := boolPtrFromResource(d, "archive_delete_all")
		var planEntity *handler.MsgOraclePlanEntity
		if v, ok := d.GetOkExists("plan_id"); ok {
			planID := v.(int)
			planEntity = &handler.MsgOraclePlanEntity{PlanID: &planID}
			if planName := stringPtrFromResource(d, "plan_name"); planName != nil {
				planEntity.Name = planName
			}
		} else if planName := stringPtrFromResource(d, "plan_name"); planName != nil {
			planEntity = &handler.MsgOraclePlanEntity{Name: planName}
		}

		req := handler.MsgModifyOracleSubclientRequest{
			SubClientProperties: &handler.MsgOracleSubclientFullProperties{
				ContentOperationType: &contentOperationType,
				Content:              content,
				CommonProperties: &handler.MsgSubclientCommonProperties{
					StorageDevice: &handler.MsgStorageDevice{
						DataBackupStoragePolicy: storagePolicyEntity,
						LogBackupStoragePolicy:  logStoragePolicyEntity,
					},
					EnableBackup: enableBackup,
					Description:  description,
				},
				OracleSubclientProp: &handler.MsgOracleSubclientDetails{
					BackupMode:                  backupMode,
					BackupArchiveLog:            backupArchiveLog,
					BackupSPFile:                backupSPFile,
					BackupControlFile:           backupControlFile,
					DeleteArchiveLogAfterBackup: deleteArchiveLog,
					SelectiveOnlineFull:         selectiveOnlineFull,
					Data:                        data,
					LightsOutScript:             lightsOutScript,
					EnableTableBrowse:           enableTableBrowse,
					ArchiveDelete:               archiveDelete,
					DataThresholdStreams:        dataThresholdStreams,
					ArchiveDeleteAll:            archiveDeleteAll,
				},
				PlanEntity: planEntity,
			},
		}

		_, err := handler.CvModifyOracleSubclient(req, d.Id())
		if err != nil {
			return fmt.Errorf("operation [ModifyOracleSubclient] failed, Error %s", err)
		}
	}

	return resourceReadOracleSubclient(d, m)
}

func resourceDeleteOracleSubclient(d *schema.ResourceData, m interface{}) error {
	_, err := handler.CvDeleteOracleSubclient(d.Id())
	if err != nil {
		return fmt.Errorf("operation [DeleteOracleSubclient] failed, Error %s", err)
	}
	return nil
}

// Helper function to build MsgIdName from terraform schema
func buildOracleIdName(r []interface{}) *handler.MsgIdName {
	if len(r) > 0 && r[0] != nil {
		tmp := r[0].(map[string]interface{})
		var name *string
		if val, ok := tmp["name"]; ok {
			name = handler.ToStringValue(val, true)
		}
		var id *int
		if val, ok := tmp["id"]; ok {
			id = handler.ToIntValue(val, true)
		}
		return &handler.MsgIdName{Name: name, Id: id}
	}
	return nil
}

func boolPtrFromResource(d *schema.ResourceData, key string) *bool {
	// Always get the value directly from state to handle explicit false
	val := d.Get(key).(bool)
	return &val
}

func intPtrFromResource(d *schema.ResourceData, key string) *int {
	if val, ok := d.GetOkExists(key); ok {
		tmp := val.(int)
		return &tmp
	}
	return nil
}

func stringPtrFromResource(d *schema.ResourceData, key string) *string {
	if val, ok := d.GetOk(key); ok {
		tmp := val.(string)
		return &tmp
	}
	return nil
}

func parseIntFromValue(v interface{}) (int, bool) {
	switch val := v.(type) {
	case float64:
		return int(val), true
	case int:
		return val, true
	case string:
		parsed, err := strconv.Atoi(val)
		if err == nil {
			return parsed, true
		}
	}
	return 0, false
}

func buildOracleContentRaw(contentItems []interface{}) []json.RawMessage {
	if len(contentItems) == 0 {
		return nil
	}
	rawItems := make([]json.RawMessage, 0, len(contentItems))
	for _, item := range contentItems {
		entry, ok := item.(map[string]interface{})
		if !ok {
			continue
		}
		payload := map[string]interface{}{}
		if oracleContentItems, ok := entry["oracle_content"].([]interface{}); ok && len(oracleContentItems) > 0 && oracleContentItems[0] != nil {
			if oracleContent, ok := oracleContentItems[0].(map[string]interface{}); ok {
				oraclePayload := map[string]interface{}{}
				if tableSpace, ok := oracleContent["table_space"].(string); ok && tableSpace != "" {
					oraclePayload["tableSpace"] = tableSpace
				}
				if len(oraclePayload) > 0 {
					payload["oracleContent"] = oraclePayload
				}
			}
		}
		if len(payload) == 0 {
			continue
		}
		if raw, err := json.Marshal(payload); err == nil {
			rawItems = append(rawItems, json.RawMessage(raw))
		}
	}
	if len(rawItems) == 0 {
		return nil
	}
	return rawItems
}

func flattenOracleContent(contentRaw interface{}) []map[string]interface{} {
	contentItems, ok := contentRaw.([]interface{})
	if !ok || len(contentItems) == 0 {
		return nil
	}
	flattened := make([]map[string]interface{}, 0, len(contentItems))
	for _, item := range contentItems {
		entry, ok := item.(map[string]interface{})
		if !ok {
			continue
		}
		oracleContentObj, ok := entry["oracleContent"].(map[string]interface{})
		if !ok {
			continue
		}
		oracleContentTF := map[string]interface{}{}
		if tableSpace, ok := oracleContentObj["tableSpace"].(string); ok {
			oracleContentTF["table_space"] = tableSpace
		}
		flattened = append(flattened, map[string]interface{}{
			"oracle_content": []interface{}{oracleContentTF},
		})
	}
	if len(flattened) == 0 {
		return nil
	}
	return flattened
}
