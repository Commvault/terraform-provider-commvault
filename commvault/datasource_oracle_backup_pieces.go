package commvault

import (
	"fmt"
	"strconv"

	"terraform-provider-commvault/commvault/handler"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceOracleBackupPieces() *schema.Resource {
	return &schema.Resource{
		Read: datasourceReadOracleBackupPieces,

		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "ID of the Oracle instance",
			},
			"from_time": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Start time (epoch) for the backup pieces query. Defaults to 2 weeks ago if not provided.",
			},
			"to_time": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "End time (epoch) for the backup pieces query. Defaults to current time if not provided.",
			},
			"backup_pieces": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of Oracle backup pieces",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"backup_piece_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the backup piece",
						},
						"tag": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Tag of the backup piece",
						},
						"start_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Start time of the backup piece (epoch)",
						},
						"end_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "End time of the backup piece (epoch)",
						},
						"backup_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of backup",
						},
						"size": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Size of the backup piece in bytes",
						},
					},
				},
			},
		},
	}
}

func datasourceReadOracleBackupPieces(d *schema.ResourceData, m interface{}) error {
	instanceId := d.Get("instance_id").(int)

	fromTime := 0
	toTime := 0

	if val, ok := d.GetOk("from_time"); ok {
		fromTime = val.(int)
	}
	if val, ok := d.GetOk("to_time"); ok {
		toTime = val.(int)
	}

	resp, err := handler.CvGetOracleBackupPieces(strconv.Itoa(instanceId), fromTime, toTime)
	if err != nil {
		return fmt.Errorf("operation [GetOracleBackupPieces] failed, Error %s", err)
	}

	d.SetId(strconv.Itoa(instanceId))

	if resp.BackupPieces != nil {
		pieces := make([]map[string]interface{}, len(resp.BackupPieces))
		for i, piece := range resp.BackupPieces {
			p := make(map[string]interface{})
			if piece.BackupPieceName != nil {
				p["backup_piece_name"] = *piece.BackupPieceName
			}
			if piece.Tag != nil {
				p["tag"] = *piece.Tag
			}
			if piece.StartTime != nil {
				p["start_time"] = *piece.StartTime
			}
			if piece.EndTime != nil {
				p["end_time"] = *piece.EndTime
			}
			if piece.BackupType != nil {
				p["backup_type"] = *piece.BackupType
			}
			if piece.Size != nil {
				p["size"] = *piece.Size
			}
			pieces[i] = p
		}
		d.Set("backup_pieces", pieces)
	} else {
		d.Set("backup_pieces", make([]map[string]interface{}, 0))
	}

	return nil
}
