package utils

import (
	"math"
	"errors"
)

func ToFloat(unk interface{}) (float64, error) {
	switch i := unk.(type) {
	case float64:
		return i, nil
	case float32:
	case int64:
	case int32:
	case int16:
	case int8:
	case int:
		return float64(i), nil
	}

	return math.NaN(), errors.New("ToFloat: unknown value is of incompatible type")
}
