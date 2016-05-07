package amml

import (
	"testing"
	"fmt"
	"reflect"
)

func TestBasic(t *testing.T) {
	amml := NewAmml("a,b,1,5")

	ammlType := reflect.TypeOf((*Amml)(nil))
	if  reflect.TypeOf(amml)!= ammlType {
		fmt.Errorf("Expect type: %T", ammlType)
	}
	if (amml.array_result[0] != "a") {
		fmt.Errorf("Array error")
	}

	arrayType := reflect.TypeOf((*[]string)(nil))
	result := amml.GetResult()
	if result != arrayType {
		fmt.Errorf("Error GetResult Type")
	}
	resultArray := result.([]string)
	if (resultArray[0] != "a") {
		fmt.Errorf("Error GetResult data")
	}
}