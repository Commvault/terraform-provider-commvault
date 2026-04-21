package commvault

import (
	"fmt"
	"strconv"

	"terraform-provider-commvault/commvault/handler"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceOracleInstallAgent() *schema.Resource {
	return &schema.Resource{
		Create: resourceCreateOracleInstallAgent,
		Read:   resourceReadOracleInstallAgent,
		Delete: resourceDeleteOracleInstallAgent,

		Schema: map[string]*schema.Schema{
			"client_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of the client where Oracle agent will be installed",
			},
			"host_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Hostname or IP address of the server where Oracle agent will be installed",
			},
			"display_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Display name for the client (defaults to client_name if not specified)",
			},
			"commserve_host_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "CommServe hostname",
			},
			"user_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Username for client authentication (SSH user for Unix/Linux)",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				ForceNew:    true,
				Description: "Password for client authentication",
			},
			"install_os_type": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     1,
				ForceNew:    true,
				Description: "OS type: 1 for Windows, 2 for Unix/Linux",
			},
			"plan_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "Plan ID to associate with the installed agent",
			},
			"unix_group": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "oinstall",
				ForceNew:    true,
				Description: "Unix group for Oracle installation (Unix/Linux only)",
			},
			"unix_group_access": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     7,
				ForceNew:    true,
				Description: "Unix group access permissions (Unix/Linux only)",
			},
			"unix_other_access": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     5,
				ForceNew:    true,
				Description: "Unix other access permissions (Unix/Linux only)",
			},
			"override_unix_group_permissions": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				ForceNew:    true,
				Description: "Whether to override Unix group and permissions",
			},
			"allow_multiple_instances": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				ForceNew:    true,
				Description: "Allow multiple instances on the same client",
			},
			"disable_os_firewall": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				ForceNew:    true,
				Description: "Disable OS firewall during installation",
			},
			"add_to_firewall_exclusion": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				ForceNew:    true,
				Description: "Add Commvault to firewall exclusion",
			},
			"force_reboot": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				ForceNew:    true,
				Description: "Force reboot after installation",
			},
			"stop_oracle_services": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				ForceNew:    true,
				Description: "Stop Oracle services during installation",
			},
			"override_client_info": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				ForceNew:    true,
				Description: "Override client information if client already exists",
			},
			"enable_firewall_config": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				ForceNew:    true,
				Description: "Enable firewall configuration",
			},
			"firewall_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				ForceNew:    true,
				Description: "Firewall port number",
			},
			// Computed attributes
			"task_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Task ID of the installation job",
			},
			"job_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Job ID of the installation job",
			},
		},
	}
}

func resourceCreateOracleInstallAgent(d *schema.ResourceData, m interface{}) error {
	clientName := d.Get("client_name").(string)
	hostName := d.Get("host_name").(string)
	displayName := d.Get("display_name").(string)
	if displayName == "" {
		displayName = clientName
	}
	commServeHostName := d.Get("commserve_host_name").(string)
	userName := d.Get("user_name").(string)
	password := d.Get("password").(string)
	installOsType := d.Get("install_os_type").(int)
	unixGroup := d.Get("unix_group").(string)
	unixGroupAccess := d.Get("unix_group_access").(int)
	unixOtherAccess := d.Get("unix_other_access").(int)
	overrideUnixGroupPermissions := d.Get("override_unix_group_permissions").(bool)
	allowMultipleInstances := d.Get("allow_multiple_instances").(bool)
	disableOsFirewall := d.Get("disable_os_firewall").(bool)
	addToFirewallExclusion := d.Get("add_to_firewall_exclusion").(bool)
	forceReboot := d.Get("force_reboot").(bool)
	stopOracleServices := d.Get("stop_oracle_services").(bool)
	overrideClientInfo := d.Get("override_client_info").(bool)
	enableFirewallConfig := d.Get("enable_firewall_config").(bool)
	firewallPort := d.Get("firewall_port").(int)

	// Build the request
	taskType := 1 // IMMEDIATE
	subTaskType := "ADMIN"
	operationType := "INSTALL_CLIENT"
	oracleComponent := "Oracle iDataAgent"
	userId := 1
	adminUserName := "admin"

	req := handler.MsgInstallOracleAgentRequest{
		TaskInfo: &handler.MsgInstallOracleAgentTaskInfo{
			Task: &handler.MsgInstallOracleAgentTask{
				TaskType: &taskType,
			},
			SubTasks: []handler.MsgInstallOracleAgentSubTask{
				{
					SubTask: &handler.MsgInstallOracleAgentSubTaskInfo{
						SubTaskType:   &subTaskType,
						OperationType: &operationType,
					},
					Options: &handler.MsgInstallOracleAgentOptions{
						AdminOpts: &handler.MsgInstallOracleAgentAdminOpts{
							ClientInstallOption: &handler.MsgClientInstallOption{
								InstallerOption: &handler.MsgInstallerOption{
									User: &handler.MsgInstallerUser{
										UserId:   &userId,
										UserName: &adminUserName,
									},
									CommServeHostName: &commServeHostName,
									InstallFlags: &handler.MsgInstallFlags{
										AllowMultipleInstances:          &allowMultipleInstances,
										DisableOSFirewall:               &disableOsFirewall,
										AddToFirewallExclusion:          &addToFirewallExclusion,
										ForceReboot:                     &forceReboot,
										StopOracleServices:              &stopOracleServices,
										OverrideClientInfo:              &overrideClientInfo,
										OverrideUnixGroupAndPermissions: &overrideUnixGroupPermissions,
										UnixGroup:                       &unixGroup,
										UnixGroupAccess:                 &unixGroupAccess,
										UnixOtherAccess:                 &unixOtherAccess,
										FirewallInstall: &handler.MsgFirewallInstall{
											EnableFirewallConfig: &enableFirewallConfig,
											PortNumber:           &firewallPort,
										},
									},
									ClientComposition: []handler.MsgClientCompositionItem{
										{
											Components: &handler.MsgComponentsInfo{
												ComponentInfo: []handler.MsgComponentInfoItem{
													{
														ComponentName: &oracleComponent,
													},
												},
											},
										},
									},
								},
								ClientDetails: []handler.MsgClientDetails{
									{
										ClientEntity: &handler.MsgClientEntity{
											ClientName:  &clientName,
											HostName:    &hostName,
											DisplayName: &displayName,
										},
									},
								},
								InstallOSType: &installOsType,
								ClientAuthForJob: &handler.MsgClientAuthForJob{
									UserName: &userName,
									Password: &password,
								},
							},
							UpdateOption: &handler.MsgUpdateOption{},
						},
					},
				},
			},
		},
	}

	// Add plan if specified
	if planId, ok := d.GetOk("plan_id"); ok {
		planIdVal := planId.(int)
		req.TaskInfo.SubTasks[0].Options.AdminOpts.UpdateOption = &handler.MsgUpdateOption{
			Plan: &handler.MsgPlanRef{
				PlanId: &planIdVal,
			},
		}
	}

	resp, err := handler.CvInstallOracleAgent(req)
	if err != nil {
		return fmt.Errorf("operation [InstallOracleAgent] failed, Error %s", err)
	}

	if resp.TaskId != nil {
		d.SetId(strconv.Itoa(*resp.TaskId))
		d.Set("task_id", *resp.TaskId)
	} else {
		return fmt.Errorf("operation [InstallOracleAgent] failed, no task ID returned")
	}

	if resp.JobIds != nil && len(resp.JobIds) > 0 {
		d.Set("job_id", resp.JobIds[0])
	}

	return nil
}

func resourceReadOracleInstallAgent(d *schema.ResourceData, m interface{}) error {
	// Installation is a one-time operation, nothing to read back
	// The resource state is maintained based on the task ID
	return nil
}

func resourceDeleteOracleInstallAgent(d *schema.ResourceData, m interface{}) error {
	// Installation cannot be "deleted" - the agent must be uninstalled separately
	// Just remove from state
	d.SetId("")
	return nil
}
