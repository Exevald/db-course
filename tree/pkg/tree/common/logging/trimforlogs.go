package logging

import (
	"fmt"
	"reflect"
	"strings"
)

type TrimForLogsOptions struct {
	MaxStringLength      int
	MaxSliceLength       int
	SensitiveFields      []string
	SensitiveKeys        []string
	SensitivePlaceholder string
}

var DefaultTrimForLogsOpts = TrimForLogsOptions{
	MaxStringLength: 255,
	MaxSliceLength:  10,
}

func TrimForLogs(v interface{}, opts TrimForLogsOptions) interface{} {
	return trimForLogs(reflect.ValueOf(v), opts)
}

//nolint:gocognit // function is easy to read as plain function
func trimForLogs(v reflect.Value, opts TrimForLogsOptions) interface{} {
	switch v.Kind() {
	case reflect.String:
		result := v.String()
		if opts.MaxStringLength != 0 && len(result) > opts.MaxStringLength {
			result = result[:opts.MaxStringLength] + "..."
		}
		return result
	case reflect.Struct:
		result := make(map[string]interface{})
		t := v.Type()
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			fieldType := t.Field(i)
			if !fieldType.IsExported() {
				continue
			}
			name := fieldType.Name
			if tag := fieldType.Tag.Get("json"); tag != "" {
				if strings.HasSuffix(tag, ",omitempty") && field.IsZero() {
					continue
				}
				name = strings.TrimSuffix(tag, ",omitempty")
			}
			if strSliceContains(fieldType.Name, opts.SensitiveFields) {
				result[name] = opts.SensitivePlaceholder
			} else {
				result[name] = trimForLogs(field, opts)
			}
		}
		return result
	case reflect.Array, reflect.Slice:
		var result []interface{}
		itemType := v.Type().Elem()
		if itemType.Kind() == reflect.Uint8 {
			return make([]interface{}, 0)
		}
		length := v.Len()
		if opts.MaxSliceLength != 0 && length > opts.MaxSliceLength {
			length = opts.MaxSliceLength
		}
		result = make([]interface{}, length)
		for i := 0; i < length; i++ {
			result[i] = trimForLogs(v.Index(i), opts)
		}
		if opts.MaxSliceLength != 0 && v.Len() > opts.MaxSliceLength {
			result = append(result, "…")
		}
		return result
	case reflect.Map:
		result := make(map[string]interface{})
		for _, k := range v.MapKeys() {
			key := fmt.Sprint(trimForLogs(k, opts))
			if strSliceContains(k.String(), opts.SensitiveKeys) {
				result[key] = opts.SensitivePlaceholder
			} else {
				result[key] = trimForLogs(v.MapIndex(k), opts)
			}
		}
		return result
	case reflect.Ptr:
		return trimForLogs(v.Elem(), opts)
	default:
		if v.IsValid() {
			return v.Interface()
		}
		return nil
	}
}

func strSliceContains(str string, slice []string) bool {
	for _, it := range slice {
		if it == str {
			return true
		}
	}
	return false
}
