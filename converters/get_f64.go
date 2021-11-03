package converters

import (
	"errors"
	"reflect"
)

func GetF64ArrAt(data [][]interface{}, index int) ([]float64, error) {
	var arr []float64
	var ok bool

	ra := reflect.ValueOf(data[index][0])
	switch ra.Kind() {
	case reflect.Slice:
		arr = make([]float64, ra.Len())
		for i := 0; i < ra.Len(); i++ {
			arr[i], ok = ra.Index(i).Interface().(float64)

			if !ok {
				return nil, errors.New("failed to convert to float64")
			}
		}
	default:
		return nil, errors.New("failed to convert to slice, no slice given by the server")
	}

	return arr, nil
}

func GetF64At(data [][]interface{}, index int) (float64, error) {

	var k float64
	if len(data[index]) > 1 {
		rk := reflect.ValueOf(data[index][1])
		switch rk.Kind() {
		case reflect.Float64:
			k = rk.Float()
		default:
			return 0, errors.New("failed to convert to float64, no float64 given by the server")
		}
	}

	return k, nil
}
