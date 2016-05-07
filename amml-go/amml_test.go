package amml

import (
	"testing"
	"fmt"
	"reflect"
	//"strings"
	//"go/parser"
)

var testArrayStr string = "中文,b,1,5"

func TestTokenParser(t *testing.T) {
	// token parser test
	parser := NewTokenizer(testArrayStr)
	if (len(parser.tokens) != 7) {
		t.Errorf("Amml tokens must be 7, but %d", len(parser.tokens))
		for _, token := range(parser.tokens) {
			fmt.Printf("%s", token.tokenString)
		}
	}
	if (parser.tokens[0].tokenString != "中文") {
		t.Errorf("token 0 error")
	}

}

func TestBasic(t *testing.T) {
	amml := NewAmml(testArrayStr)

	// test token reader
	//reader := strings.NewReader("ABCDEFGHIJKLMN中文测试")
	//for _, ch := range "中文测试abcdefg"{
	//	fmt.Printf("CH: %c, size: %s, err: %s\n", ch)
	//}
	ammlType := reflect.TypeOf((*Amml)(nil))
	if  reflect.TypeOf(amml)!= ammlType {
		t.Errorf("Expect type: %T", ammlType)
	}
	if (amml.arrayResult[0] != "中文") {
		t.Errorf("Array error")
	}

	arrayType := reflect.TypeOf((*[]string)(nil))
	result := amml.GetResult()
	if result != arrayType {
		fmt.Errorf("Error GetResult Type")
	}
	resultArray := result.([]interface{})
	if (resultArray[0] != "a") {
		fmt.Errorf("Error GetResult data")
	}
	resultArray, _ = amml.GetArray()
	if (resultArray[0] != "a") {
		fmt.Errorf("Error GetArray data")
	}
}