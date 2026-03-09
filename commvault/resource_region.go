package commvault

import (
    "fmt"
    "strconv"
    "strings"

    "terraform-provider-commvault/commvault/handler"

    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceRegion() *schema.Resource {
    return &schema.Resource{
        Create: resourceCreateRegion,
        Read:   resourceReadRegion,
        Update: resourceUpdateRegion,
        Delete: resourceDeleteRegion,

        Schema: map[string]*schema.Schema{
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "Region name",
            },
            "globalconfiginfo": {
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "Only applicable to Global CommCells. Not applicable for SaaS.",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "scopefilterquery": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "CommCellEntityCache filter query string using for filtering the scope",
                        },
                        "companies": {
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Computed:    true,
                            Description: "List of companies where the global configuration should be applied",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                    },
                                    "guid": {
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
                        "applyonallcommcells": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "Decides whether the global configuration should be applied to all the Service commcells, including the newly created ones",
                        },
                        "commcells": {
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Computed:    true,
                            Description: "List of Service CommCells where the global configuration should be applied",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "displayname": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                    },
                                    "name": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                    },
                                    "guid": {
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
                        "scope": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "The entity level at which the config has to be applied.",
                        },
                        "name": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                        },
                        "applyonallcompanies": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "Decides whether the global configuration should be applied to all the companies, including the newly created ones",
                        },
                        "actiononlocalentity": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "Action that will be taken on the local entity that has the same name as the global entity that needs to be created [CREATE_NEW, TAKE_OVER, FAIL_IF_EXIST]",
                        },
                    },
                },
            },
            "locations": {
                Type:        schema.TypeSet,
                Required:    true,
                Description: "List of locations which are part of the region",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "country": {
                            Type:        schema.TypeString,
                            Required:    true,
                            Description: "Name of country",
                        },
                        "city": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: "Name of city",
                        },
                        "latitude": {
                            Type:        schema.TypeFloat,
                            Required:    true,
                            Description: "Latitude for the location",
                        },
                        "state": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: "Name of state",
                        },
                        "longitude": {
                            Type:        schema.TypeFloat,
                            Required:    true,
                            Description: "Longitude for the location",
                        },
                    },
                },
            },
            "type": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Type of the region [USER_CREATED, AWS, AZURE, OCI, GCP]",
            },
        },
    }
}

func resourceCreateRegion(d *schema.ResourceData, m interface{}) error {
    //API: (POST) /V4/Regions
    var response_id = strconv.Itoa(0)
    var t_name *string
    if val, ok := d.GetOk("name"); ok {
        t_name = handler.ToStringValue(val, false)
    }
    var t_globalconfiginfo *handler.MsgCreateGlobalConfigInfo
    if val, ok := d.GetOk("globalconfiginfo"); ok {
        t_globalconfiginfo = build_region_msgcreateglobalconfiginfo(d, val.([]interface{}))
    }
    var t_locations []handler.MsgLocationDetailsSet
    if val, ok := d.GetOk("locations"); ok {
        t_locations = build_region_msglocationdetailsset_array(d, val.(*schema.Set).List())
    }
    var t_type *string
    if val, ok := d.GetOk("type"); ok {
        t_type = handler.ToStringValue(val, false)
    }
    var req = handler.MsgCreateRegionRequest{Name:t_name, GlobalConfigInfo:t_globalconfiginfo, Locations:t_locations, Type:t_type}
    resp, err := handler.CvCreateRegion(req)
    if err != nil {
        return fmt.Errorf("operation [CreateRegion] failed, Error %s", err)
    }
    if resp.Id != nil {
        response_id = strconv.Itoa(*resp.Id)
    }
    if response_id == "0" {
        return fmt.Errorf("operation [CreateRegion] failed")
    } else {
        d.SetId(response_id)
        return resourceCreateUpdateRegion(d, m)
    }
}

func resourceReadRegion(d *schema.ResourceData, m interface{}) error {
    //API: (GET) /V4/Regions/{regionId}
    resp, err := handler.CvGetRegionDetails(d.Id())
    if err != nil {
        if strings.Contains(err.Error(), "status: 404") {
            handler.LogEntry("debug", "entity not present, removing from state")
            d.SetId("")
            return nil
        }
        return fmt.Errorf("operation [GetRegionDetails] failed, Error %s", err)
    }
    if resp.Name != nil {
        d.Set("name", resp.Name)
    }
    if rtn, ok := serialize_region_msgglobalconfiginfo(d, resp.GlobalConfigInfo); ok {
        d.Set("globalconfiginfo", rtn)
    } else {
        d.Set("globalconfiginfo", make([]map[string]interface{}, 0))
    }
    if rtn, ok := serialize_region_msglocationdetailswithzoneset_array(d, resp.Locations); ok {
        d.Set("locations", rtn)
    } else {
        d.Set("locations", make([]map[string]interface{}, 0))
    }
    return nil
}

func resourceUpdateRegion(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V4/Regions/{regionId}
    var t_newname *string
    if d.HasChange("name") {
        val := d.Get("name")
        t_newname = handler.ToStringValue(val, false)
    }
    var t_locations []handler.MsgLocationDetailsWithZoneSet
    if d.HasChange("locations") {
        val := d.Get("locations")
        t_locations = build_region_msglocationdetailswithzoneset_array(d, val.(*schema.Set).List())
    }
    var t_locationsoperationtype *string
    var c_locationsoperationtype string = "OVERWRITE"
    t_locationsoperationtype = &c_locationsoperationtype
    var req = handler.MsgUpdateRegionRequest{NewName:t_newname, Locations:t_locations, LocationsOperationType:t_locationsoperationtype}
    _, err := handler.CvUpdateRegion(req, d.Id())
    if err != nil {
        return fmt.Errorf("operation [UpdateRegion] failed, Error %s", err)
    }
    return resourceReadRegion(d, m)
}

func resourceCreateUpdateRegion(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V4/Regions/{regionId}
    var execUpdate bool = false
    var t_locationsoperationtype *string
    var c_locationsoperationtype string = "OVERWRITE"
    t_locationsoperationtype = &c_locationsoperationtype
    if execUpdate {
        var req = handler.MsgUpdateRegionRequest{LocationsOperationType:t_locationsoperationtype}
        _, err := handler.CvUpdateRegion(req, d.Id())
        if err != nil {
            return fmt.Errorf("operation [UpdateRegion] failed, Error %s", err)
        }
    }
    return resourceReadRegion(d, m)
}

func resourceDeleteRegion(d *schema.ResourceData, m interface{}) error {
    //API: (DELETE) /V4/Regions/{regionId}
    _, err := handler.CvDeleteRegion(d.Id())
    if err != nil {
        return fmt.Errorf("operation [DeleteRegion] failed, Error %s", err)
    }
    return nil
}

func build_region_msglocationdetailswithzoneset_array(d *schema.ResourceData, r []interface{}) []handler.MsgLocationDetailsWithZoneSet {
    if r != nil {
        tmp := make([]handler.MsgLocationDetailsWithZoneSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_continent *string
            if val, ok := raw_a["continent"]; ok {
                t_continent = handler.ToStringValue(val, true)
            }
            var t_country *string
            if val, ok := raw_a["country"]; ok {
                t_country = handler.ToStringValue(val, true)
            }
            var t_city *string
            if val, ok := raw_a["city"]; ok {
                t_city = handler.ToStringValue(val, true)
            }
            var t_zone *handler.MsgIdName
            if val, ok := raw_a["zone"]; ok {
                t_zone = build_region_msgidname(d, val.([]interface{}))
            }
            var t_latitude *float64
            if val, ok := raw_a["latitude"]; ok {
                t_latitude = handler.ToDoubleValue(val, true)
            }
            var t_state *string
            if val, ok := raw_a["state"]; ok {
                t_state = handler.ToStringValue(val, true)
            }
            var t_longitude *float64
            if val, ok := raw_a["longitude"]; ok {
                t_longitude = handler.ToDoubleValue(val, true)
            }
            tmp[a] = handler.MsgLocationDetailsWithZoneSet{Continent:t_continent, Country:t_country, City:t_city, Zone:t_zone, Latitude:t_latitude, State:t_state, Longitude:t_longitude}
        }
        return tmp
    } else {
        return nil
    }
}

func build_region_msgidname(d *schema.ResourceData, r []interface{}) *handler.MsgIdName {
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

func build_region_msglocationdetailsset_array(d *schema.ResourceData, r []interface{}) []handler.MsgLocationDetailsSet {
    if r != nil {
        tmp := make([]handler.MsgLocationDetailsSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_country *string
            if val, ok := raw_a["country"]; ok {
                t_country = handler.ToStringValue(val, true)
            }
            var t_city *string
            if val, ok := raw_a["city"]; ok {
                t_city = handler.ToStringValue(val, true)
            }
            var t_latitude *float64
            if val, ok := raw_a["latitude"]; ok {
                t_latitude = handler.ToDoubleValue(val, true)
            }
            var t_state *string
            if val, ok := raw_a["state"]; ok {
                t_state = handler.ToStringValue(val, true)
            }
            var t_longitude *float64
            if val, ok := raw_a["longitude"]; ok {
                t_longitude = handler.ToDoubleValue(val, true)
            }
            tmp[a] = handler.MsgLocationDetailsSet{Country:t_country, City:t_city, Latitude:t_latitude, State:t_state, Longitude:t_longitude}
        }
        return tmp
    } else {
        return nil
    }
}

func build_region_msgcreateglobalconfiginfo(d *schema.ResourceData, r []interface{}) *handler.MsgCreateGlobalConfigInfo {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_scopefilterquery *string
        if val, ok := tmp["scopefilterquery"]; ok {
            t_scopefilterquery = handler.ToStringValue(val, true)
        }
        var t_companies []handler.MsgGlobalConfigCompanyInfoSet
        if val, ok := tmp["companies"]; ok {
            t_companies = build_region_msgglobalconfigcompanyinfoset_array(d, val.(*schema.Set).List())
        }
        var t_applyonallcommcells *bool
        if val, ok := tmp["applyonallcommcells"]; ok {
            t_applyonallcommcells = handler.ToBooleanValue(val, true)
        }
        var t_commcells []handler.MsgGlobalConfigCommcellInfoSet
        if val, ok := tmp["commcells"]; ok {
            t_commcells = build_region_msgglobalconfigcommcellinfoset_array(d, val.(*schema.Set).List())
        }
        var t_scope *string
        if val, ok := tmp["scope"]; ok {
            t_scope = handler.ToStringValue(val, true)
        }
        var t_name *string
        if val, ok := tmp["name"]; ok {
            t_name = handler.ToStringValue(val, true)
        }
        var t_applyonallcompanies *bool
        if val, ok := tmp["applyonallcompanies"]; ok {
            t_applyonallcompanies = handler.ToBooleanValue(val, true)
        }
        var t_actiononlocalentity *string
        if val, ok := tmp["actiononlocalentity"]; ok {
            t_actiononlocalentity = handler.ToStringValue(val, true)
        }
        return &handler.MsgCreateGlobalConfigInfo{ScopeFilterQuery:t_scopefilterquery, Companies:t_companies, ApplyOnAllCommCells:t_applyonallcommcells, Commcells:t_commcells, Scope:t_scope, Name:t_name, ApplyOnAllCompanies:t_applyonallcompanies, ActionOnLocalEntity:t_actiononlocalentity}
    } else {
        return nil
    }
}

func build_region_msgglobalconfigcommcellinfoset_array(d *schema.ResourceData, r []interface{}) []handler.MsgGlobalConfigCommcellInfoSet {
    if r != nil {
        tmp := make([]handler.MsgGlobalConfigCommcellInfoSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_displayname *string
            if val, ok := raw_a["displayname"]; ok {
                t_displayname = handler.ToStringValue(val, true)
            }
            var t_name *string
            if val, ok := raw_a["name"]; ok {
                t_name = handler.ToStringValue(val, true)
            }
            var t_guid *string
            if val, ok := raw_a["guid"]; ok {
                t_guid = handler.ToStringValue(val, true)
            }
            var t_id *int
            if val, ok := raw_a["id"]; ok {
                t_id = handler.ToIntValue(val, true)
            }
            tmp[a] = handler.MsgGlobalConfigCommcellInfoSet{DisplayName:t_displayname, Name:t_name, Guid:t_guid, Id:t_id}
        }
        return tmp
    } else {
        return nil
    }
}

func build_region_msgglobalconfigcompanyinfoset_array(d *schema.ResourceData, r []interface{}) []handler.MsgGlobalConfigCompanyInfoSet {
    if r != nil {
        tmp := make([]handler.MsgGlobalConfigCompanyInfoSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_name *string
            if val, ok := raw_a["name"]; ok {
                t_name = handler.ToStringValue(val, true)
            }
            var t_guid *string
            if val, ok := raw_a["guid"]; ok {
                t_guid = handler.ToStringValue(val, true)
            }
            var t_id *int
            if val, ok := raw_a["id"]; ok {
                t_id = handler.ToIntValue(val, true)
            }
            tmp[a] = handler.MsgGlobalConfigCompanyInfoSet{Name:t_name, Guid:t_guid, Id:t_id}
        }
        return tmp
    } else {
        return nil
    }
}

func serialize_region_msglocationdetailswithzoneset_array(d *schema.ResourceData, data []handler.MsgLocationDetailsWithZoneSet) ([]map[string]interface{}, bool) {
    //MsgLocationDetailsSet
    //MsgLocationDetailsWithZoneSet
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 0)
    for i := range data {
        tmp := make(map[string]interface{})
        added := false
        if data[i].Country != nil {
            tmp["country"] = data[i].Country
            added = true
        }
        if data[i].City != nil {
            tmp["city"] = data[i].City
            added = true
        }
        if data[i].Latitude != nil {
            tmp["latitude"] = data[i].Latitude
            added = true
        }
        if data[i].State != nil {
            tmp["state"] = data[i].State
            added = true
        }
        if data[i].Longitude != nil {
            tmp["longitude"] = data[i].Longitude
            added = true
        }
        if added {
            val = append(val, tmp)
        }
    }
    return val, true
}

func serialize_region_msgglobalconfiginfo(d *schema.ResourceData, data *handler.MsgGlobalConfigInfo) ([]map[string]interface{}, bool) {
    //MsgCreateGlobalConfigInfo
    //MsgGlobalConfigInfo
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 1)
    val[0] = make(map[string]interface{})
    added := false
    if data.ScopeFilterQuery != nil {
        val[0]["scopefilterquery"] = data.ScopeFilterQuery
        added = true
    }
    if rtn, ok := serialize_region_msgglobalconfigcompanyinfoset_array(d, data.Companies); ok {
        val[0]["companies"] = rtn
        added = true
    }
    if data.ApplyOnAllCommCells != nil {
        val[0]["applyonallcommcells"] = strconv.FormatBool(*data.ApplyOnAllCommCells)
        added = true
    }
    if rtn, ok := serialize_region_msgglobalconfigcommcellinfoset_array(d, data.Commcells); ok {
        val[0]["commcells"] = rtn
        added = true
    }
    if data.Scope != nil {
        val[0]["scope"] = data.Scope
        added = true
    }
    if data.Name != nil {
        val[0]["name"] = data.Name
        added = true
    }
    if data.ApplyOnAllCompanies != nil {
        val[0]["applyonallcompanies"] = strconv.FormatBool(*data.ApplyOnAllCompanies)
        added = true
    }
    if added {
        return val, true
    } else {
        return nil, false
    }
}

func serialize_region_msgglobalconfigcommcellinfoset_array(d *schema.ResourceData, data []handler.MsgGlobalConfigCommcellInfoSet) ([]map[string]interface{}, bool) {
    //MsgCreateGlobalConfigInfo -> MsgGlobalConfigCommcellInfoSet
    //MsgGlobalConfigInfo -> MsgGlobalConfigCommcellInfoSet
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 0)
    for i := range data {
        tmp := make(map[string]interface{})
        added := false
        if data[i].DisplayName != nil {
            tmp["displayname"] = data[i].DisplayName
            added = true
        }
        if data[i].Name != nil {
            tmp["name"] = data[i].Name
            added = true
        }
        if data[i].Guid != nil {
            tmp["guid"] = data[i].Guid
            added = true
        }
        if data[i].Id != nil {
            tmp["id"] = data[i].Id
            added = true
        }
        if added {
            val = append(val, tmp)
        }
    }
    return val, true
}

func serialize_region_msgglobalconfigcompanyinfoset_array(d *schema.ResourceData, data []handler.MsgGlobalConfigCompanyInfoSet) ([]map[string]interface{}, bool) {
    //MsgCreateGlobalConfigInfo -> MsgGlobalConfigCompanyInfoSet
    //MsgGlobalConfigInfo -> MsgGlobalConfigCompanyInfoSet
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 0)
    for i := range data {
        tmp := make(map[string]interface{})
        added := false
        if data[i].Name != nil {
            tmp["name"] = data[i].Name
            added = true
        }
        if data[i].Guid != nil {
            tmp["guid"] = data[i].Guid
            added = true
        }
        if data[i].Id != nil {
            tmp["id"] = data[i].Id
            added = true
        }
        if added {
            val = append(val, tmp)
        }
    }
    return val, true
}
