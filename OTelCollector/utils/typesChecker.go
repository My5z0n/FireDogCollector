package utils

import "errors"

func GetTypeName(v any) (error, string) {
	switch v.(type) {
	case string:
		return nil, "String"
	case bool:
		return nil, "Bool"
	case int64:
		return nil, "Int64"
	case float64:
		return nil, "Float64"
	default:
		return errors.New("Unsuported Data Type"), ""
	}
}
