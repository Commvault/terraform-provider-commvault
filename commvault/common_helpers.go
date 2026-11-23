package commvault

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func makeComputed(fields map[string]*schema.Schema) {
	for _, field := range fields {
		field.Required, field.Optional, field.ForceNew = false, false, false
		field.Computed = true
		field.Default, field.DefaultFunc, field.ValidateFunc = nil, nil, nil
		field.MinItems, field.MaxItems = 0, 0
		if nested, ok := field.Elem.(*schema.Resource); ok {
			makeComputed(nested.Schema)
		}
	}
}

func cloudOptionalBool(d *schema.ResourceData, key string) *bool {
	if value, ok := d.GetOkExists(key); ok {
		result := value.(bool)
		return &result
	}
	return nil
}

func cloudOptionalString(d *schema.ResourceData, key string) *string {
	if value, ok := d.GetOkExists(key); ok {
		result := value.(string)
		return &result
	}
	return nil
}

func boolPtrFromResource(d *schema.ResourceData, key string) *bool {
	// Always get the value directly from state to handle explicit false.
	val := d.Get(key).(bool)
	return &val
}

func intPtrFromResource(d *schema.ResourceData, key string) *int {
	if val, ok := d.GetOkExists(key); ok {
		tmp := val.(int)
		return &tmp
	}
	return nil
}

func stringPtrFromResource(d *schema.ResourceData, key string) *string {
	if val, ok := d.GetOk(key); ok {
		tmp := val.(string)
		return &tmp
	}
	return nil
}

func parseIntFromValue(v interface{}) (int, bool) {
	switch val := v.(type) {
	case float64:
		return int(val), true
	case int:
		return val, true
	case string:
		parsed, err := strconv.Atoi(val)
		if err == nil {
			return parsed, true
		}
	}
	return 0, false
}

func validateNonEmptyString(value interface{}, key string) ([]string, []error) {
	if strings.TrimSpace(value.(string)) == "" {
		return nil, []error{fmt.Errorf("%s must not be empty", key)}
	}
	return nil, nil
}

func validateIntRange(min, max int) schema.SchemaValidateFunc {
	return func(value interface{}, key string) ([]string, []error) {
		number := value.(int)
		if number < min || number > max {
			return nil, []error{fmt.Errorf("%s must be between %d and %d", key, min, max)}
		}
		return nil, nil
	}
}

func validateRDSStringEnum(allowed ...string) schema.SchemaValidateFunc {
	return func(value interface{}, key string) ([]string, []error) {
		for _, candidate := range allowed {
			if value.(string) == candidate {
				return nil, nil
			}
		}
		return nil, []error{fmt.Errorf("%s must be one of %v", key, allowed)}
	}
}

func refreshCloudInstanceRename(d *schema.ResourceData, m interface{}, expected string, read func(*schema.ResourceData, interface{}) error) error {
	for attempt := 0; attempt < 30; attempt++ {
		if err := read(d, m); err != nil {
			return err
		}
		if d.Id() == "" {
			return fmt.Errorf("renamed cloud instance disappeared during refresh")
		}
		if d.Get("instance_name").(string) == expected {
			return nil
		}
		// Reads can lag a successful rename; never repeat the write.
		time.Sleep(time.Second)
	}
	return fmt.Errorf("cloud instance %s rename was accepted but %q was not observed within 30 seconds; inspect before retrying", d.Id(), expected)
}

func rdsPutString(values map[string]interface{}, key string, value *string) {
	if value != nil {
		values[key] = *value
	}
}

func s3PutString(values map[string]interface{}, key string, value *string) {
	if value != nil {
		values[key] = *value
	}
}

func s3PutInt(values map[string]interface{}, key string, value *int) {
	if value != nil {
		values[key] = *value
	}
}
