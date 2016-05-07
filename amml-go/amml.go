package amml

import (
	"strings"
	"errors"
)

type Amml struct {
	arrayResult []interface{} // private array reslt
				  //map_result map[string]interface{}

}

// Token container
type Token struct {
	tokenIndex int
	tokenType TokenType
	tokenString string
}

// Auto get result,  array or map
func (this *Amml) GetResult() interface{}{
	return this.arrayResult
}
// get array result if can
func (this *Amml) GetArray() (result []interface{}, err error) {
	if this.arrayResult == nil {
		return nil, errors.New("Nil Array Result")
	}
	return this.arrayResult, nil
}
// Constructor of Amml
func (this *Amml) init(str string) {
	//parser := NewTokenParser(str)

	arr := strings.Split(str, ",")
	arrObjs := make([]interface{}, len(arr))
	for i, item := range arr {
		arrObjs[i] = item //item.(interface{})

	}
	this.arrayResult = arrObjs
}

func NewAmml(str string) *Amml {
	amml := new(Amml)
	amml.init(str)

	return amml;
}