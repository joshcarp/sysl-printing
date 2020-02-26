package eval

import (
	sysl "github.com/joshcarp/sysl-printing/pkg/sysl"
	"github.com/pkg/errors"
)

type valueType int

// const definitions for various sysl.Value types
const (
	ValueNoArg valueType = -1
	ValueBool  valueType = iota
	ValueInt
	ValueFloat
	ValueString
	ValueStringDecimal
	ValueList
	ValueMap
	ValueSet
	ValueNull
)

//nolint:gochecknoglobals
var valueTypeNames = map[valueType]string{
	ValueNoArg:         "ValueNoArg",
	ValueBool:          "ValueBool",
	ValueInt:           "ValueInt",
	ValueFloat:         "ValueFloat",
	ValueString:        "ValueString",
	ValueStringDecimal: "ValueStringDecimal",
	ValueList:          "ValueList",
	ValueMap:           "ValueMap",
	ValueSet:           "ValueSet",
	ValueNull:          "ValueNull",
}

func (v valueType) String() string {
	return valueTypeNames[v]
}

func getValueType(v *sysl.Value) valueType {
	if v == nil {
		return ValueNoArg
	}
	switch v.Value.(type) {
	case *sysl.Value_B:
		return ValueBool
	case *sysl.Value_I:
		return ValueInt
	case *sysl.Value_D:
		return ValueFloat
	case *sysl.Value_S:
		return ValueString
	case *sysl.Value_Decimal:
		return ValueStringDecimal
	case *sysl.Value_Set:
		return ValueSet
	case *sysl.Value_List_:
		return ValueList
	case *sysl.Value_Map_:
		return ValueMap
	case *sysl.Value_Null_:
		return ValueNull
	default:
		panic(errors.Errorf("exprOp: getValueType: unhandled type: %v", v))
	}
}

func getContainedType(container *sysl.Value) valueType {
	var list []*sysl.Value
	switch x := container.Value.(type) {
	case *sysl.Value_List_:
		list = x.List.Value
	case *sysl.Value_Set:
		list = x.Set.Value
	default:
		return ValueNoArg
	}

	if len(list) == 0 {
		return ValueNoArg
	}
	return getValueType(list[0])
}
