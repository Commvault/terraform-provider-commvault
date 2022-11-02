package handler

import (
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func UpdateUserRequest(req *MsgModifyUserRequest, d *schema.ResourceData, m interface{}) error {
	if d.HasChange("password") {
		var val string
		if os.Getenv("CV_TER_PASSWORD") != "" {
			val = os.Getenv("CV_TER_PASSWORD")
		} else {
			val = os.Getenv("CV_PASSWORD")
		}
		if val == "" {
			return fmt.Errorf("cannot change password without a provider password or the environment variable CV_TER_PASSWORD")
		}
		req.ValidationPassword = new(string)
		req.ValidationPassword = &val
	}
	return nil
}

func GetConsoleTypes(values []MsgRestrictedConsoleTypesSet) []map[string]interface{} {
	items := make([]string, len(values))
	for a, raw_a := range values {
		items[a] = *raw_a.ConsoleType
	}

	data := make([]map[string]interface{}, 1)
	data[0] = map[string]interface{}{
		"consoletype": items,
	}

	return data
}
