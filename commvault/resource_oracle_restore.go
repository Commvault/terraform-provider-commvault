package commvault

import (
	"fmt"
	"strconv"

	"terraform-provider-commvault/commvault/handler"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceOracleRestore() *schema.Resource {
	return &schema.Resource{
		Create: resourceCreateOracleRestore,
		Read:   resourceReadOracleRestore,
		Delete: resourceDeleteOracleRestore,

		Schema: map[string]*schema.Schema{
			"client_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of the source client",
			},
			"client_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "ID of the source client",
			},
			"instance_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of the source Oracle instance",
			},
			"instance_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "ID of the source Oracle instance",
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
			"restore_type": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Default:     0,
				Description: "Type of restore operation",
			},
			"destination_client": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Name of the destination client for out-of-place restore",
			},
			"destination_instance": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Name of the destination Oracle instance for out-of-place restore",
			},
			"point_in_time": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "Point in time (Unix timestamp) for the restore",
			},
			"scn": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "System Change Number (SCN) for the restore",
			},
			"task_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "ID of the restore task",
			},
			"job_ids": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of job IDs created by the restore",
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
		},
	}
}

func resourceCreateOracleRestore(d *schema.ResourceData, m interface{}) error {
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

	restoreType := d.Get("restore_type").(int)
	taskType := 1
	initiatedFrom := 1
	purityType := 0
	subTaskType := 3
	operationType := 1

	restoreOpts := &handler.MsgOracleRestoreOptions{
		RestoreType: &restoreType,
	}

	if val, ok := d.GetOk("destination_client"); ok {
		destClient := val.(string)
		restoreOpts.DestinationClient = &destClient
	}
	if val, ok := d.GetOk("destination_instance"); ok {
		destInstance := val.(string)
		restoreOpts.DestinationInstance = &destInstance
	}
	if val, ok := d.GetOk("point_in_time"); ok {
		pit := val.(int)
		restoreOpts.PointInTime = &pit
	}
	if val, ok := d.GetOk("scn"); ok {
		scn := val.(string)
		restoreOpts.SCN = &scn
	}

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
						RestoreOpts: restoreOpts,
					},
				},
			},
		},
	}

	resp, err := handler.CvOracleRestore(req)
	if err != nil {
		return fmt.Errorf("operation [OracleRestore] failed, Error %s", err)
	}

	if resp.TaskId != nil && *resp.TaskId > 0 {
		d.SetId(strconv.Itoa(*resp.TaskId))
		d.Set("task_id", *resp.TaskId)
	} else {
		return fmt.Errorf("operation [OracleRestore] failed, no task ID returned")
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

func resourceReadOracleRestore(d *schema.ResourceData, m interface{}) error {
	// Restore is a one-time operation, so we just return the stored values
	return nil
}

func resourceDeleteOracleRestore(d *schema.ResourceData, m interface{}) error {
	// Restore cannot be deleted through this resource
	// We just remove it from state
	d.SetId("")
	return nil
}
