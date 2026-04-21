package commvault

import (
	"fmt"

	"terraform-provider-commvault/commvault/handler"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceOracleRMANLogs() *schema.Resource {
	return &schema.Resource{
		Read: datasourceReadOracleRMANLogs,

		Schema: map[string]*schema.Schema{
			"job_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Job ID for which to fetch RMAN logs",
			},
			"log_content": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Content of the RMAN logs",
			},
			"error_string": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Error string if any",
			},
		},
	}
}

func datasourceReadOracleRMANLogs(d *schema.ResourceData, m interface{}) error {
	jobId := d.Get("job_id").(string)

	resp, err := handler.CvFetchRMANLogs(jobId)
	if err != nil {
		return fmt.Errorf("operation [FetchRMANLogs] failed, Error %s", err)
	}

	d.SetId(jobId)

	if resp.LogContent != nil {
		d.Set("log_content", *resp.LogContent)
	}
	if resp.ErrorString != nil {
		d.Set("error_string", *resp.ErrorString)
	}

	return nil
}
