package schema

//go:generate stringer -type=Type type.go

// Type is an enum of types that A schema represent
type Type int

const (
	TypeInvalid Type = iota
	TypeBool
	TypeInt
	TypeFloat
	TypeString
	TypeList
	TypeMap
	TypeSet
)
