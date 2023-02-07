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

func GetConsoleTypes(d *schema.ResourceData, values []MsgRestrictedConsoleTypesSet) ([]map[string]interface{}, bool) {
	items := make([]string, len(values))
	for a, raw_a := range values {
		items[a] = *raw_a.ConsoleType
	}

	data := make([]map[string]interface{}, 1)
	data[0] = map[string]interface{}{
		"consoletype": items,
	}

	return data, true
}

func SortPlanSchedules(d *schema.ResourceData, data []MsgPlanScheduleSet) []MsgPlanScheduleSet {
	if len(data) == 0 {
		return data
	}

	rpo, _ := d.Get("rpo").([]interface{})
	if len(rpo) == 0 {
		return data
	}
	t_rpo := rpo[0].(map[string]interface{})

	backupfrequency, h_backupfrequency := t_rpo["backupfrequency"].([]interface{})
	if !h_backupfrequency || len(backupfrequency) == 0 {
		return data
	}
	t_backupfrequency := backupfrequency[0].(map[string]interface{})

	schedules, h_schedules := t_backupfrequency["schedules"].([]interface{})
	if !h_schedules || len(schedules) == 0 {
		return data
	}

	curr_data := make([]MsgPlanScheduleSet, 0)
	missing_data := make([]string, 0)

	for _, iter_a := range schedules {
		raw_a := iter_a.(map[string]interface{})
		sched_name := raw_a["schedulename"].(string)
		tmp := nextPlanSchedules(sched_name, data)
		if tmp != nil {
			curr_data = append(curr_data, *tmp)
		} else {
			missing_data = append(missing_data, sched_name)
		}
	}

	for _, sched_name := range missing_data {
		tmp := nextPlanSchedules(sched_name, data)
		if tmp != nil {
			curr_data = append(curr_data, *tmp)
		}
	}

	return curr_data
}

func nextPlanSchedules(scheduleName string, data []MsgPlanScheduleSet) *MsgPlanScheduleSet {
	for _, iter_a := range data {
		if iter_a.ScheduleName != nil && *iter_a.ScheduleName == scheduleName {
			return &iter_a
		}
	}
	return nil
}
