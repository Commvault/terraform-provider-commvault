package main

import (
    "fmt"
    "github.com/TerraformProvider/handler"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceVMGroup() *schema.Resource {
    return &schema.Resource{
        Create: resourceVMGroupCreate,
        Read:   resourceVMGroupRead,
        Update: resourceVMGroupUpdate,
        Delete: resourceVMGroupDelete,

        Schema: map[string]*schema.Schema{
            "vmgroupname": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "clientid": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "planid":&schema.Schema{
                Type:schema.TypeString,
                Required:true,
            },
            "vms":&schema.Schema{
                Type:schema.TypeSet,
                Required:true,
                Elem: &schema.Schema{Type: schema.TypeString},
            },
        },
    }
}

func resourceVMGroupCreate(d *schema.ResourceData, m interface{}) error {
    vmGroupName := d.Get("vmgroupname").(string)
    if vmGroupName == "" {
        return fmt.Errorf("vmgroup Name cannot be empty")
    }
    clientid := d.Get("clientid").(string)
    planid := d.Get("planid").(string)
    vms := d.Get("vms").(*schema.Set).List()
    vmnames := make([]string, len(vms))
    for i,n := range vms {
        vmnames[i] = n.(string)
    }
    apiResp := handler.VMGroupCreate(vmGroupName,planid,clientid,vmnames)
    if apiResp.Response.ErrorCode != "0" {
        return fmt.Errorf("Error creating vmgroup.")
    }
    d.SetId(apiResp.Response.Entity.SubclientId)
    return resourceVMGroupRead(d,m)
}

func resourceVMGroupRead(d *schema.ResourceData, m interface{}) error {
    return nil
}

func resourceVMGroupUpdate(d *schema.ResourceData, m interface{}) error {
    return resourceVMGroupRead(d, m)
}

func resourceVMGroupDelete(d *schema.ResourceData, m interface{}) error {
    vmgroupid := d.Id()
    resp := handler.VMGroupDelete(vmgroupid)
    if resp.Response.ErrorCode != "0" {
        return fmt.Errorf("Couldnt delete vmgroupid " + vmgroupid)
    }
    d.SetId("")
    return nil
}



