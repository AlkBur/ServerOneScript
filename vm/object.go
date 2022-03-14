package vm

import (
	"errors"
	"github.com/shopspring/decimal"
)

type ObjectType int

const (
	TypeString = iota
	TypeNumber
	TypeBool
	TypeVariant
)

func (t ObjectType) String() string {
	switch t {
	case TypeString:
		return "string"
	case TypeBool:
		return "boolean"
	case TypeNumber:
		return "number"
	case TypeVariant:
		return "variant"
	}
	return "unknown"
}

type Object interface {
	String() string
	Type() ObjectType
	Equal(Object) bool
}

//////// БАЗОВЫЕ ТИПЫ //////////

//Number - число
type Number struct {
	val decimal.Decimal
}

func NewObjectNumber(val decimal.Decimal) *Number {
	return &Number{
		val: val,
	}
}

func (obj *Number) String() string {
	return obj.val.String()
}

func (obj *Number) Type() ObjectType {
	return TypeNumber
}

func (obj *Number) Equal(val Object) bool {
	if obj.Type() == val.Type() {
		num := val.(*Number)
		return obj.val.Equal(num.val)
	}
	return false
}

//String - строка
type String struct {
	val string
}

func NewObjectString(val string) *String {
	return &String{
		val: val,
	}
}

func (obj *String) String() string {
	return obj.val
}

func (obj *String) Type() ObjectType {
	return TypeString
}

func (obj *String) Equal(val Object) bool {
	if obj.Type() == val.Type() {
		num := val.(*String)
		return obj.val == num.val
	}
	return false
}

//Bool - булево
type Bool struct {
	val bool
}

func NewObjectBool(val bool) *Bool {
	return &Bool{
		val: val,
	}
}

func (obj *Bool) String() string {
	if obj.val {
		return "Истина"
	}
	return "Ложь"
}

func (obj *Bool) Type() ObjectType {
	return TypeBool
}

func (obj *Bool) Equal(val Object) bool {
	if obj.Type() == val.Type() {
		num := val.(*Bool)
		return obj.val == num.val
	}
	return false
}

//Variant - переменная
type Variant struct {
	name string
}

func NewObjectVariant(name string) *Variant {
	return &Variant{
		name: name,
	}
}

func (obj *Variant) String() string {
	return obj.name
}

func (obj *Variant) Type() ObjectType {
	return TypeVariant
}

func (obj *Variant) Equal(val Object) bool {
	if obj.Type() == val.Type() {
		num := val.(*Variant)
		return obj.name == num.name
	}
	return false
}

func AddObject(v1, v2 Object) (result Object, err error) {
	switch v1.Type() {
	case TypeNumber:
		val := v1.(*Number).val
		switch v2.Type() {
		case TypeNumber:
			result = NewObjectNumber(val.Add(v2.(*Number).val))
		case TypeBool:
			if v2.(*Bool).val {
				result = NewObjectNumber(val.Add(decimal.NewFromInt32(1)))
			} else {
				result = NewObjectNumber(val)
			}
		default:
			err = errors.New("число с " + v2.Type().String())
		}
	case TypeBool:
		val := v1.(*Bool).val
		switch v2.Type() {
		case TypeNumber:
			if v2.(*Number).val.Cmp(decimal.NewFromInt32(0)) > 0 {
				result = NewObjectBool(true)
			} else {
				result = NewObjectBool(val)
			}
		case TypeBool:
			if v2.(*Bool).val {
				result = NewObjectBool(true)
			} else {
				result = NewObjectBool(val)
			}
		default:
			err = errors.New("boolean с " + v2.Type().String())
		}
	case TypeString:
		result = NewObjectString(v1.String() + v2.String())
	default:
		err = errors.New(v2.Type().String() + " с " + v2.Type().String())
	}
	return
}

func SubObject(v1, v2 Object) (result Object, err error) {
	switch v1.Type() {
	case TypeNumber:
		val := v1.(*Number).val
		switch v2.Type() {
		case TypeNumber:
			result = NewObjectNumber(val.Sub(v2.(*Number).val))
		case TypeBool:
			if v2.(*Bool).val {
				result = NewObjectNumber(val.Sub(decimal.NewFromInt32(1)))
			} else {
				result = NewObjectNumber(val)
			}
		default:
			err = errors.New("число с " + v2.Type().String())
		}
	case TypeBool:
		val := v1.(*Bool).val
		switch v2.Type() {
		case TypeNumber:
			if v2.(*Number).val.Cmp(decimal.NewFromInt32(0)) > 0 {
				result = NewObjectBool(false)
			} else {
				result = NewObjectBool(val)
			}
		case TypeBool:
			if v2.(*Bool).val {
				result = NewObjectBool(false)
			} else {
				result = NewObjectBool(val)
			}
		default:
			err = errors.New("boolean с " + v2.Type().String())
		}
	default:
		err = errors.New(v2.Type().String() + " с " + v2.Type().String())
	}
	return
}

func MulObject(v1, v2 Object) (result Object, err error) {
	switch v1.Type() {
	case TypeNumber:
		val := v1.(*Number).val
		switch v2.Type() {
		case TypeNumber:
			result = NewObjectNumber(val.Mul(v2.(*Number).val))
		case TypeBool:
			if v2.(*Bool).val {
				result = NewObjectNumber(val)
			} else {
				result = NewObjectNumber(decimal.NewFromInt32(0))
			}
		default:
			err = errors.New("число с " + v2.Type().String())
		}
	default:
		err = errors.New(v2.Type().String() + " с " + v2.Type().String())
	}
	return
}

func DivObject(v1, v2 Object) (result Object, err error) {
	switch v1.Type() {
	case TypeNumber:
		val := v1.(*Number).val
		switch v2.Type() {
		case TypeNumber:
			result = NewObjectNumber(val.Div(v2.(*Number).val))
		default:
			err = errors.New("число с " + v2.Type().String())
		}
	default:
		err = errors.New(v2.Type().String() + " с " + v2.Type().String())
	}
	return
}
