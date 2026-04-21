package commvault

import (
	"fmt"
	"strconv"

	"terraform-provider-commvault/commvault/handler"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceOracleBackup() *schema.Resource {
	return &schema.Resource{
		Create: resourceCreateOracleBackup,
		Read:   resourceReadOracleBackup,
		Delete: resourceDeleteOracleBackup,

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
				ForceNew:    true,
				Description: "ID of the client",
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
				ForceNew:    true,
				Description: "ID of the Oracle instance",
			},
			"subclient_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Name of the Oracle subclient",
			},
			"subclient_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "ID of the Oracle subclient",
			},
			"backup_level": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Default:     1,
				Description: "Backup level: 1 for Full, 2 for Incremental",
			},
			"task_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "ID of the backup task",
			},
			"job_ids": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of job IDs created by the backup",
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
		},
	}
}

func resourceCreateOracleBackup(d *schema.ResourceData, m interface{}) error {
	clientName := d.Get("client_name").(string)
	instanceName := d.Get("instance_name").(string)
	appName := "Oracle"
	applicationId := 22

	// Get IDs
	var clientId *int
	var instanceId *int
	var subclientId *int
	var subclientName *string

	if val, ok := d.GetOk("client_id"); ok {
		id := val.(int)
		clientId = &id
	}
	if val, ok := d.GetOk("instance_id"); ok {
		id := val.(int)
		instanceId = &id
	}
	if val, ok := d.GetOk("subclient_id"); ok {
		id := val.(int)
		subclientId = &id
	}
	if val, ok := d.GetOk("subclient_name"); ok {
		name := val.(string)
		subclientName = &name
	}

	// Fetch entity IDs if not provided
	if clientId == nil || instanceId == nil {
		scName := ""
		if subclientName != nil {
			scName = *subclientName
		}
		entityResp, err := handler.CvFetchOracleEntityId(clientName, instanceName, scName)
		if err != nil {
			return fmt.Errorf("failed to fetch Oracle entity IDs, Error %s", err)
		}
		if clientId == nil && entityResp.ClientId > 0 {
			id := entityResp.ClientId
			clientId = &id
		}
		if instanceId == nil && entityResp.InstanceId > 0 {
			id := entityResp.InstanceId
			instanceId = &id
		}
		if subclientId == nil && entityResp.SubclientId > 0 {
			id := entityResp.SubclientId
			subclientId = &id
		}
	}

	backupLevel := d.Get("backup_level").(int)
	taskType := 1
	initiatedFrom := 1
	purityType := 0
	subTaskType := 2
	operationType := 2

	req := handler.MsgOracleBackupRequest{
		TaskInfo: &handler.MsgOracleTaskInfo{
			Associations: []handler.MsgOracleAssociation{
				{
					ClientName:    &clientName,
					ClientId:      clientId,
					InstanceName:  &instanceName,
					InstanceId:    instanceId,
					SubclientName: subclientName,
					SubclientId:   subclientId,
					AppName:       &appName,
					ApplicationId: &applicationId,
				},
			},
			Task: &handler.MsgOracleTask{
				TaskType:      &taskType,
				InitiatedFrom: &initiatedFrom,
				PurityType:    &purityType,
			},
			SubTasks: []handler.MsgOracleSubTask{
				{
					SubTaskType:   &subTaskType,
					OperationType: &operationType,
					Options: &handler.MsgOracleSubTaskOptions{
						BackupOpts: &handler.MsgOracleBackupOptions{
							BackupLevel: &backupLevel,
						},
					},
				},
			},
		},
	}

	resp, err := handler.CvOracleBackup(req)
	if err != nil {
		return fmt.Errorf("operation [OracleBackup] failed, Error %s", err)
	}

	if resp.TaskId != nil && *resp.TaskId > 0 {
		d.SetId(strconv.Itoa(*resp.TaskId))
		d.Set("task_id", *resp.TaskId)
	} else {
		return fmt.Errorf("operation [OracleBackup] failed, no task ID returned")
	}

	if len(resp.JobIds) > 0 {
		d.Set("job_ids", resp.JobIds)
	}

	// Set computed values
	if clientId != nil {
		d.Set("client_id", *clientId)
	}
	if instanceId != nil {
		d.Set("instance_id", *instanceId)
	}
	if subclientId != nil {
		d.Set("subclient_id", *subclientId)
	}

	return nil
}

func resourceReadOracleBackup(d *schema.ResourceData, m interface{}) error {
	// Backup is a one-time operation, so we just return the stored values
	return nil
}

func resourceDeleteOracleBackup(d *schema.ResourceData, m interface{}) error {
	// Backup cannot be deleted through this resource
	// We just remove it from state
	d.SetId("")
	return nil
}
