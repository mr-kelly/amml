package amml

import "strings"

type Amml struct {
	array_result []string // private array reslt
	//map_result map[string]interface{}
}

// Auto get result,  array or map
func (this *Amml) GetResult() interface{}{
	return this.array_result
}
// Constructor of Amml
func (this *Amml) Init(str string) {
	var arr = strings.Split(str, ",")
	this.array_result = arr
}

func NewAmml(str string) *Amml {
	amml := new(Amml)
	amml.Init(str)

	return amml;
}