package dataparsers

import (
	"mservice/models"
	"reflect"
)

// Dataset for executing array cyclic rotation sligtly differs
// from dataset for other tasks:
// [
//     [[]float64, float64]
//	   ...
// ]

func ParseForCyclicRotation(data [][]interface{}) *[]models.CyclicRotationData {
	res := make([]models.CyclicRotationData, len(data))

	for i := range data { // TODO: Swap copying for actuall conversion
		v := reflect.ValueOf(data[i][0])
		switch v.Kind() {
		case reflect.Slice:
			res[i].Arr = make([]interface{}, v.Len())
			for j := 0; j < v.Len(); j++ {
				res[i].Arr[j] = v.Index(j).Interface()
				res[i].Arr[j] = res[i].Arr[j].(float64)
			}
		}
		res[i].K = data[i][1].(float64)
	}

	return &res
}
