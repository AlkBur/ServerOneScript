package parser

import "reflect"

var binaryOperators = map[string]int{
	"or":  10,
	"или": 10,
	"and": 15,
	"и":   15,
	"=":   20,
	"<>":  20,
	"<":   20,
	">":   20,
	">=":  20,
	"<=":  20,
	"+":   30,
	"-":   30,
	"*":   60,
	"/":   60,
	"%":   60,
}

var unaryOperators = map[string]int{
	"not": 50,
	"не":  50,
}

func isNil(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	}
	return false
}

/// GetPrecedence - Возвращает приоритет текущего бинарного оператора.
func GetPrecedence(val string) int {
	// Удостоверимся, что это объявленный бинарный оператор.
	TokPrec, ok := binaryOperators[val]
	if !ok {
		return -1
	}
	return TokPrec
}
