// Code generated by "stringer -type=Type type.go"; DO NOT EDIT

package schema

import "fmt"

const _Type_name = "TypeInvalidTypeBoolTypeIntTypeFloatTypeStringTypeListTypeMapTypeSet"

var _Type_index = [...]uint8{0, 11, 19, 26, 35, 45, 53, 60, 67}

func (i Type) String() string {
	if i < 0 || i >= Type(len(_Type_index)-1) {
		return fmt.Sprintf("Type(%d)", i)
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}
