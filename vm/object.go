package vm

const (
	ObjectDigit int = iota
	ObjectBool
	ObjectString
	ObjectBase
)

type Object struct {
	num  float64
	str  string
	b    bool
	Type int
}

func NewObjectDigit(val float64) *Object {
	return &Object{
		num:  val,
		Type: ObjectDigit,
	}
}

func NewObjectString(val string) *Object {
	return &Object{
		str:  val,
		Type: ObjectString,
	}
}

func NewObjectObjectBool(val bool) *Object {
	return &Object{
		b:    val,
		Type: ObjectBool,
	}
}

func NewObjectObjectBase() *Object {
	return &Object{
		Type: ObjectBase,
	}
}
