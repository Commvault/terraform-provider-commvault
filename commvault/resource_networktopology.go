package commvault

import (
    "fmt"
    "strconv"
    "strings"

    "terraform-provider-commvault/commvault/handler"

    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceNetworkTopology() *schema.Resource {
    return &schema.Resource{
        Create: resourceCreateNetworkTopology,
        Read:   resourceReadNetworkTopology,
        Update: resourceUpdateNetworkTopology,
        Delete: resourceDeleteNetworkTopology,

        Schema: map[string]*schema.Schema{
            "firewallgroups": {
                Type:        schema.TypeSet,
                Optional:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "grouptype": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: "A FirewallFwGroupType defines a specific group type within a topology payload. [INTERNAL, EXTERNAL, PROXIES, PROXY_PERIMETER, PROXY2, PROXY3]",
                        },
                        "clientgroupid": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: "The id of the client group",
                        },
                        "advancedoptions": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "keepaliveinterval": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Description: "The interval in seconds for sending keep-alive packets, to maintain the session if backup traffic has an extended pause.",
                                    },
                                    "tunnelport": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Description: "The port defined for communications",
                                    },
                                },
                            },
                        },
                        "mnemonic": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: "Custom names(mnemonics) given to pre-defined smart client groups [MY_COMMSERVE, MY_MEDIAAGENTS, MY_COMMSERVE_AND_MEDIAAGENTS, NONE]",
                        },
                    },
                },
            },
            "topologyname": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "",
            },
            "clienttype": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "The type of clients associated with the topology. [SERVER, LAPTOP]",
            },
            "usewildcardproxy": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Flag determining wether network gateways are used to connect all infrastructure machines",
            },
            "topologytype": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "The type of topology of client groups (and sometimes network gateway) for this instance. [NETWORK_GATEWAY, ONE_WAY, TWO_WAY, CASCADING_GATEWAYS, ONEWAY_FORWARDING, TRI_CASCADING_GATEWAYS, QUAD_CASCADING_GATEWAYS]",
            },
            "tunnelsperroute": {
                Type:        schema.TypeInt,
                Optional:    true,
                Computed:    true,
                Description: "The number of tunnel connections per route",
            },
            "tunnelprotocol": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "The protocol for outgoing communication [REGULAR, ENCRYPTED, AUTHENTICATED, RAW]",
            },
            "encrypttraffic": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Flag determining if we want the data from tunnel to use HTTPS protocol",
            },
        },
    }
}

func resourceCreateNetworkTopology(d *schema.ResourceData, m interface{}) error {
    //API: (POST) /V4/NetworkTopology
    var response_id = strconv.Itoa(0)
    var t_firewallgroups []handler.MsgFirewallTopologyGroupsSet
    if val, ok := d.GetOk("firewallgroups"); ok {
        t_firewallgroups = build_networktopology_msgfirewalltopologygroupsset_array(d, val.(*schema.Set).List())
    }
    var t_topologyname *string
    if val, ok := d.GetOk("topologyname"); ok {
        t_topologyname = handler.ToStringValue(val, false)
    }
    var t_clienttype *string
    if val, ok := d.GetOk("clienttype"); ok {
        t_clienttype = handler.ToStringValue(val, false)
    }
    var t_usewildcardproxy *bool
    if val, ok := d.GetOk("usewildcardproxy"); ok {
        t_usewildcardproxy = handler.ToBooleanValue(val, false)
    }
    var t_topologytype *string
    if val, ok := d.GetOk("topologytype"); ok {
        t_topologytype = handler.ToStringValue(val, false)
    }
    var t_tunnelsperroute *int
    if val, ok := d.GetOk("tunnelsperroute"); ok {
        t_tunnelsperroute = handler.ToIntValue(val, false)
    }
    var t_tunnelprotocol *string
    if val, ok := d.GetOk("tunnelprotocol"); ok {
        t_tunnelprotocol = handler.ToStringValue(val, false)
    }
    var t_encrypttraffic *bool
    if val, ok := d.GetOk("encrypttraffic"); ok {
        t_encrypttraffic = handler.ToBooleanValue(val, false)
    }
    var req = handler.MsgPOSTFirewallTopologyRequest{FirewallGroups:t_firewallgroups, TopologyName:t_topologyname, ClientType:t_clienttype, UseWildCardProxy:t_usewildcardproxy, TopologyType:t_topologytype, TunnelsPerRoute:t_tunnelsperroute, TunnelProtocol:t_tunnelprotocol, EncryptTraffic:t_encrypttraffic}
    resp, err := handler.CvPOSTFirewallTopology(req)
    if err != nil {
        return fmt.Errorf("operation [POSTFirewallTopology] failed, Error %s", err)
    }
    if resp.TopologyId != nil {
        response_id = strconv.Itoa(*resp.TopologyId)
    }
    if response_id == "0" {
        return fmt.Errorf("operation [POSTFirewallTopology] failed")
    } else {
        d.SetId(response_id)
        return resourceReadNetworkTopology(d, m)
    }
}

func resourceReadNetworkTopology(d *schema.ResourceData, m interface{}) error {
    //API: (GET) /V4/NetworkTopology/{topologyId}
    _, err := handler.CvGETFirewallTopologyDetails(d.Id())
    if err != nil {
        if strings.Contains(err.Error(), "status: 404") {
            handler.LogEntry("debug", "entity not present, removing from state")
            d.SetId("")
            return nil
        }
        return fmt.Errorf("operation [GETFirewallTopologyDetails] failed, Error %s", err)
    }
    return nil
}

func resourceUpdateNetworkTopology(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V4/NetworkTopology/{topologyId}
    var t_firewallgroups []handler.MsgFirewallTopologyGroupsSet
    if d.HasChange("firewallgroups") {
        val := d.Get("firewallgroups")
        t_firewallgroups = build_networktopology_msgfirewalltopologygroupsset_array(d, val.(*schema.Set).List())
    }
    var t_topologyname *string
    if d.HasChange("topologyname") {
        val := d.Get("topologyname")
        t_topologyname = handler.ToStringValue(val, false)
    }
    var t_clienttype *string
    if d.HasChange("clienttype") {
        val := d.Get("clienttype")
        t_clienttype = handler.ToStringValue(val, false)
    }
    var t_usewildcardproxy *bool
    if d.HasChange("usewildcardproxy") {
        val := d.Get("usewildcardproxy")
        t_usewildcardproxy = handler.ToBooleanValue(val, false)
    }
    var t_topologytype *string
    if d.HasChange("topologytype") {
        val := d.Get("topologytype")
        t_topologytype = handler.ToStringValue(val, false)
    }
    var t_tunnelsperroute *int
    if d.HasChange("tunnelsperroute") {
        val := d.Get("tunnelsperroute")
        t_tunnelsperroute = handler.ToIntValue(val, false)
    }
    var t_tunnelprotocol *string
    if d.HasChange("tunnelprotocol") {
        val := d.Get("tunnelprotocol")
        t_tunnelprotocol = handler.ToStringValue(val, false)
    }
    var t_encrypttraffic *bool
    if d.HasChange("encrypttraffic") {
        val := d.Get("encrypttraffic")
        t_encrypttraffic = handler.ToBooleanValue(val, false)
    }
    var req = handler.MsgPUTFirewallTopologyRequest{FirewallGroups:t_firewallgroups, TopologyName:t_topologyname, ClientType:t_clienttype, UseWildCardProxy:t_usewildcardproxy, TopologyType:t_topologytype, TunnelsPerRoute:t_tunnelsperroute, TunnelProtocol:t_tunnelprotocol, EncryptTraffic:t_encrypttraffic}
    _, err := handler.CvPUTFirewallTopology(req, d.Id())
    if err != nil {
        return fmt.Errorf("operation [PUTFirewallTopology] failed, Error %s", err)
    }
    return resourceReadNetworkTopology(d, m)
}

func resourceDeleteNetworkTopology(d *schema.ResourceData, m interface{}) error {
    //API: (DELETE) /V4/NetworkTopology/{topologyId}
    _, err := handler.CvDELETEFirewallTopology(d.Id())
    if err != nil {
        return fmt.Errorf("operation [DELETEFirewallTopology] failed, Error %s", err)
    }
    return nil
}

func build_networktopology_msgfirewalltopologygroupsset_array(d *schema.ResourceData, r []interface{}) []handler.MsgFirewallTopologyGroupsSet {
    if r != nil {
        tmp := make([]handler.MsgFirewallTopologyGroupsSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_grouptype *string
            if val, ok := raw_a["grouptype"]; ok {
                t_grouptype = handler.ToStringValue(val, true)
            }
            var t_clientgroupid *int
            if val, ok := raw_a["clientgroupid"]; ok {
                t_clientgroupid = handler.ToIntValue(val, true)
            }
            var t_advancedoptions *handler.MsgFirewallGroupAdvancedOptions
            if val, ok := raw_a["advancedoptions"]; ok {
                t_advancedoptions = build_networktopology_msgfirewallgroupadvancedoptions(d, val.([]interface{}))
            }
            var t_mnemonic *string
            if val, ok := raw_a["mnemonic"]; ok {
                t_mnemonic = handler.ToStringValue(val, true)
            }
            tmp[a] = handler.MsgFirewallTopologyGroupsSet{GroupType:t_grouptype, ClientGroupId:t_clientgroupid, AdvancedOptions:t_advancedoptions, Mnemonic:t_mnemonic}
        }
        return tmp
    } else {
        return nil
    }
}

func build_networktopology_msgfirewallgroupadvancedoptions(d *schema.ResourceData, r []interface{}) *handler.MsgFirewallGroupAdvancedOptions {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_keepaliveinterval *int
        if val, ok := tmp["keepaliveinterval"]; ok {
            t_keepaliveinterval = handler.ToIntValue(val, true)
        }
        var t_tunnelport *int
        if val, ok := tmp["tunnelport"]; ok {
            t_tunnelport = handler.ToIntValue(val, true)
        }
        return &handler.MsgFirewallGroupAdvancedOptions{KeepAliveInterval:t_keepaliveinterval, TunnelPort:t_tunnelport}
    } else {
        return nil
    }
}
