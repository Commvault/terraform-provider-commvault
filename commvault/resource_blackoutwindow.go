package commvault

import (
    "fmt"
    "strconv"
    "strings"

    "terraform-provider-commvault/commvault/handler"

    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceBlackoutWindow() *schema.Resource {
    return &schema.Resource{
        Create: resourceCreateBlackoutWindow,
        Read:   resourceReadBlackoutWindow,
        Update: resourceUpdateBlackoutWindow,
        Delete: resourceDeleteBlackoutWindow,

        Schema: map[string]*schema.Schema{
            "donotsubmitjob": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Allows or Denies submitting a job when the blackout window is in effect. If set to false, the job is submitted and resumed once the blackout window ends.",
            },
            "alldays": {
                Type:        schema.TypeSet,
                Optional:    true,
                Description: "Days of the week along with the time on which the black out window will be in effect.",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "days": {
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Description: "Days of the week when the blackout window will be in effect.",
                            Elem: &schema.Schema{
                                Type:    schema.TypeString,
                            },
                        },
                        "time": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "start": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Description: "the blackout window comes into effect at this point.",
                                    },
                                    "end": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Description: "the blackout window is no longer in effect from this point on.",
                                    },
                                },
                            },
                        },
                    },
                },
            },
            "betweendates": {
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "start": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Computed:    true,
                            Description: "the blackout window comes into effect at this point.",
                        },
                        "end": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Computed:    true,
                            Description: "the blackout window is no longer in effect from this point on.",
                        },
                    },
                },
            },
            "weeks": {
                Type:        schema.TypeSet,
                Optional:    true,
                Computed:    true,
                Description: "Refers to the weeks of the month that the blackout window will be in effect.",
                Elem: &schema.Schema{
                    Type:    schema.TypeString,
                },
            },
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "Name of the blackout window to be created.",
            },
            "company": {
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "name": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                        },
                        "id": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                        },
                    },
                },
            },
            "backupoperations": {
                Type:        schema.TypeSet,
                Optional:    true,
                Computed:    true,
                Description: "Refers to backup types to include in the blackout window",
                Elem: &schema.Schema{
                    Type:    schema.TypeString,
                },
            },
        },
    }
}

func resourceCreateBlackoutWindow(d *schema.ResourceData, m interface{}) error {
    //API: (POST) /V5/BlackoutWindow
    var response_id = strconv.Itoa(0)
    var t_donotsubmitjob *bool
    if val, ok := d.GetOk("donotsubmitjob"); ok {
        t_donotsubmitjob = handler.ToBooleanValue(val, false)
    }
    var t_alldays []handler.MsgDaysAndTimesSet
    if val, ok := d.GetOk("alldays"); ok {
        t_alldays = build_blackoutwindow_msgdaysandtimesset_array(d, val.(*schema.Set).List())
    }
    var t_betweendates *handler.MsgStartEnd
    if val, ok := d.GetOk("betweendates"); ok {
        t_betweendates = build_blackoutwindow_msgstartend(d, val.([]interface{}))
    }
    var t_weeks []string
    if val, ok := d.GetOk("weeks"); ok {
        t_weeks = handler.ToStringArray(val.(*schema.Set).List())
    }
    var t_name *string
    if val, ok := d.GetOk("name"); ok {
        t_name = handler.ToStringValue(val, false)
    }
    var t_company *handler.MsgIdName
    if val, ok := d.GetOk("company"); ok {
        t_company = build_blackoutwindow_msgidname(d, val.([]interface{}))
    }
    var t_backupoperations []string
    if val, ok := d.GetOk("backupoperations"); ok {
        t_backupoperations = handler.ToStringArray(val.(*schema.Set).List())
    }
    var req = handler.MsgCreateBlackoutWindowRequest{DoNotSubmitJob:t_donotsubmitjob, AllDays:t_alldays, BetweenDates:t_betweendates, Weeks:t_weeks, Name:t_name, Company:t_company, BackupOperations:t_backupoperations}
    resp, err := handler.CvCreateBlackoutWindow(req)
    if err != nil {
        return fmt.Errorf("operation [CreateBlackoutWindow] failed, Error %s", err)
    }
    if resp.Id != nil {
        response_id = strconv.Itoa(*resp.Id)
    }
    if response_id == "0" {
        return fmt.Errorf("operation [CreateBlackoutWindow] failed")
    } else {
        d.SetId(response_id)
        return resourceReadBlackoutWindow(d, m)
    }
}

func resourceReadBlackoutWindow(d *schema.ResourceData, m interface{}) error {
    //API: (GET) /V5/BlackoutWindow/{blackoutWindowId}
    resp, err := handler.CvGetBlackoutWindowDetails(d.Id())
    if err != nil {
        if strings.Contains(err.Error(), "status: 404") {
            handler.LogEntry("debug", "entity not present, removing from state")
            d.SetId("")
            return nil
        }
        return fmt.Errorf("operation [GetBlackoutWindowDetails] failed, Error %s", err)
    }
    if resp.DoNotSubmitJob != nil {
        d.Set("donotsubmitjob", strconv.FormatBool(*resp.DoNotSubmitJob))
    }
    if rtn, ok := serialize_blackoutwindow_msgdaysandtimesset_array(d, resp.AllDays); ok {
        d.Set("alldays", rtn)
    } else {
        d.Set("alldays", make([]map[string]interface{}, 0))
    }
    if rtn, ok := serialize_blackoutwindow_msgstartend(d, resp.BetweenDates); ok {
        d.Set("betweendates", rtn)
    } else {
        d.Set("betweendates", make([]map[string]interface{}, 0))
    }
    if resp.Weeks != nil {
        d.Set("weeks", resp.Weeks)
    }
    if resp.Name != nil {
        d.Set("name", resp.Name)
    }
    if rtn, ok := serialize_blackoutwindow_msgidname(d, resp.Company); ok {
        d.Set("company", rtn)
    } else {
        d.Set("company", make([]map[string]interface{}, 0))
    }
    if resp.BackupOperations != nil {
        d.Set("backupoperations", resp.BackupOperations)
    }
    return nil
}

func resourceUpdateBlackoutWindow(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V5/BlackoutWindow/{blackoutWindowId}
    var t_donotsubmitjob *bool
    if d.HasChange("donotsubmitjob") {
        val := d.Get("donotsubmitjob")
        t_donotsubmitjob = handler.ToBooleanValue(val, false)
    }
    var t_alldays []handler.MsgDaysAndTimesSet
    if d.HasChange("alldays") {
        val := d.Get("alldays")
        t_alldays = build_blackoutwindow_msgdaysandtimesset_array(d, val.(*schema.Set).List())
    }
    var t_betweendates *handler.MsgStartEnd
    if d.HasChange("betweendates") {
        val := d.Get("betweendates")
        t_betweendates = build_blackoutwindow_msgstartend(d, val.([]interface{}))
    }
    var t_newname *string
    if d.HasChange("name") {
        val := d.Get("name")
        t_newname = handler.ToStringValue(val, false)
    }
    var t_weeks []string
    if d.HasChange("weeks") {
        val := d.Get("weeks")
        t_weeks = handler.ToStringArray(val.(*schema.Set).List())
    }
    var t_company *handler.MsgIdName
    if d.HasChange("company") {
        val := d.Get("company")
        t_company = build_blackoutwindow_msgidname(d, val.([]interface{}))
    }
    var t_backupoperations []string
    if d.HasChange("backupoperations") {
        val := d.Get("backupoperations")
        t_backupoperations = handler.ToStringArray(val.(*schema.Set).List())
    }
    var req = handler.MsgModifyBlackoutWindowRequest{DoNotSubmitJob:t_donotsubmitjob, AllDays:t_alldays, BetweenDates:t_betweendates, NewName:t_newname, Weeks:t_weeks, Company:t_company, BackupOperations:t_backupoperations}
    _, err := handler.CvModifyBlackoutWindow(req, d.Id())
    if err != nil {
        return fmt.Errorf("operation [ModifyBlackoutWindow] failed, Error %s", err)
    }
    return resourceReadBlackoutWindow(d, m)
}

func resourceDeleteBlackoutWindow(d *schema.ResourceData, m interface{}) error {
    //API: (DELETE) /V4/BlackoutWindow/{blackoutWindowId}
    _, err := handler.CvDeleteBlackoutWindow(d.Id())
    if err != nil {
        return fmt.Errorf("operation [DeleteBlackoutWindow] failed, Error %s", err)
    }
    return nil
}

func build_blackoutwindow_msgidname(d *schema.ResourceData, r []interface{}) *handler.MsgIdName {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_name *string
        if val, ok := tmp["name"]; ok {
            t_name = handler.ToStringValue(val, true)
        }
        var t_id *int
        if val, ok := tmp["id"]; ok {
            t_id = handler.ToIntValue(val, true)
        }
        return &handler.MsgIdName{Name:t_name, Id:t_id}
    } else {
        return nil
    }
}

func build_blackoutwindow_msgstartend(d *schema.ResourceData, r []interface{}) *handler.MsgStartEnd {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_start *int64
        if val, ok := tmp["start"]; ok {
            t_start = handler.ToLongValue(val, true)
        }
        var t_end *int64
        if val, ok := tmp["end"]; ok {
            t_end = handler.ToLongValue(val, true)
        }
        return &handler.MsgStartEnd{Start:t_start, End:t_end}
    } else {
        return nil
    }
}

func build_blackoutwindow_msgdaysandtimesset_array(d *schema.ResourceData, r []interface{}) []handler.MsgDaysAndTimesSet {
    if r != nil {
        tmp := make([]handler.MsgDaysAndTimesSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_days []string
            if val, ok := raw_a["days"]; ok {
                t_days = handler.ToStringArray(val.(*schema.Set).List())
            }
            var t_time *handler.MsgStartEnd
            if val, ok := raw_a["time"]; ok {
                t_time = build_blackoutwindow_msgstartend(d, val.([]interface{}))
            }
            tmp[a] = handler.MsgDaysAndTimesSet{Days:t_days, Time:t_time}
        }
        return tmp
    } else {
        return nil
    }
}

func serialize_blackoutwindow_msgidname(d *schema.ResourceData, data *handler.MsgIdName) ([]map[string]interface{}, bool) {
    //MsgIdName
    //MsgIdName
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 1)
    val[0] = make(map[string]interface{})
    added := false
    if data.Name != nil {
        val[0]["name"] = data.Name
        added = true
    }
    if data.Id != nil {
        val[0]["id"] = data.Id
        added = true
    }
    if added {
        return val, true
    } else {
        return nil, false
    }
}

func serialize_blackoutwindow_msgstartend(d *schema.ResourceData, data *handler.MsgStartEnd) ([]map[string]interface{}, bool) {
    //MsgStartEnd
    //MsgStartEnd
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 1)
    val[0] = make(map[string]interface{})
    added := false
    if data.Start != nil {
        val[0]["start"] = data.Start
        added = true
    }
    if data.End != nil {
        val[0]["end"] = data.End
        added = true
    }
    if added {
        return val, true
    } else {
        return nil, false
    }
}

func serialize_blackoutwindow_msgdaysandtimesset_array(d *schema.ResourceData, data []handler.MsgDaysAndTimesSet) ([]map[string]interface{}, bool) {
    //MsgDaysAndTimesSet
    //MsgDaysAndTimesSet
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 0)
    for i := range data {
        tmp := make(map[string]interface{})
        added := false
        if data[i].Days != nil {
            tmp["days"] = data[i].Days
            added = true
        }
        if rtn, ok := serialize_blackoutwindow_msgstartend(d, data[i].Time); ok {
            tmp["time"] = rtn
            added = true
        }
        if added {
            val = append(val, tmp)
        }
    }
    return val, true
}
