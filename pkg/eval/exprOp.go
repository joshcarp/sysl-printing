package eval

import (
	"encoding/json"
	"sort"

	sysl "github.com/joshcarp/sysl-printing/pkg/sysl"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func addInt64(lhs, rhs *sysl.Value) *sysl.Value {
	return MakeValueI64(lhs.GetI() + rhs.GetI())
}

func gtInt64(lhs, rhs *sysl.Value) *sysl.Value {
	return MakeValueBool(lhs.GetI() > rhs.GetI())
}

func ltInt64(lhs, rhs *sysl.Value) *sysl.Value {
	return MakeValueBool(lhs.GetI() < rhs.GetI())
}

func geInt64(lhs, rhs *sysl.Value) *sysl.Value {
	return MakeValueBool(lhs.GetI() >= rhs.GetI())
}

func leInt64(lhs, rhs *sysl.Value) *sysl.Value {
	return MakeValueBool(lhs.GetI() <= rhs.GetI())
}

func subInt64(lhs, rhs *sysl.Value) *sysl.Value {
	return MakeValueI64(lhs.GetI() - rhs.GetI())
}

func mulInt64(lhs, rhs *sysl.Value) *sysl.Value {
	return MakeValueI64(lhs.GetI() * rhs.GetI())
}

func divInt64(lhs, rhs *sysl.Value) *sysl.Value {
	return MakeValueI64(lhs.GetI() / rhs.GetI())
}

func modInt64(lhs, rhs *sysl.Value) *sysl.Value {
	return MakeValueI64(lhs.GetI() % rhs.GetI())
}

func addString(lhs, rhs *sysl.Value) *sysl.Value {
	return MakeValueString(lhs.GetS() + rhs.GetS())
}

func cmpString(lhs, rhs *sysl.Value) *sysl.Value {
	return MakeValueBool(lhs.GetS() == rhs.GetS())
}

func cmpInt(lhs, rhs *sysl.Value) *sysl.Value {
	return MakeValueBool(lhs.GetI() == rhs.GetI())
}

func cmpBool(lhs, rhs *sysl.Value) *sysl.Value {
	return MakeValueBool(lhs.GetB() == rhs.GetB())
}

func andBool(lhs, rhs *sysl.Value) *sysl.Value {
	return MakeValueBool(lhs.GetB() && rhs.GetB())
}

func cmpNullTrue(lhs, rhs *sysl.Value) *sysl.Value {
	return MakeValueBool(true)
}

func cmpNullFalse(lhs, rhs *sysl.Value) *sysl.Value {
	return MakeValueBool(false)
}

func cmpListNull(lhs, rhs *sysl.Value) *sysl.Value {
	return MakeValueBool(lhs.GetList() == nil && rhs.GetList() == nil)
}

func flattenListMap(
	ee *exprEval,
	assign Scope,
	list *sysl.Value,
	scopeVar string,
	rhs *sysl.Expr,
) *sysl.Value {
	listResult := MakeValueList()
	for _, l := range list.GetList().Value {
		assign[scopeVar] = l
		AppendItemToValueList(listResult.GetList(), Eval(ee, assign, rhs))
	}
	return listResult
}

func flattenListList(
	ee *exprEval,
	assign Scope,
	list *sysl.Value,
	scopeVar string,
	rhs *sysl.Expr,
) *sysl.Value {
	listResult := MakeValueList()
	for _, l := range list.GetList().Value {
		for _, ll := range l.GetList().Value {
			assign[scopeVar] = ll
			AppendItemToValueList(listResult.GetList(), Eval(ee, assign, rhs))
		}
	}
	return listResult
}

func flattenListSet(
	ee *exprEval,
	assign Scope,
	list *sysl.Value,
	scopeVar string,
	rhs *sysl.Expr,
) *sysl.Value {
	listResult := MakeValueList()
	for _, l := range list.GetList().Value {
		for _, ll := range l.GetSet().Value {
			assign[scopeVar] = ll
			AppendItemToValueList(listResult.GetList(), Eval(ee, assign, rhs))
		}
	}
	return listResult
}

func flattenSetList(
	ee *exprEval,
	assign Scope,
	list *sysl.Value,
	scopeVar string,
	rhs *sysl.Expr,
) *sysl.Value {
	setResult := MakeValueSet()
	for _, l := range list.GetSet().Value {
		for _, ll := range l.GetList().Value {
			assign[scopeVar] = ll
			AppendItemToValueList(setResult.GetSet(), Eval(ee, assign, rhs))
		}
	}
	return setResult
}

func flattenSetMap(
	ee *exprEval,
	assign Scope,
	list *sysl.Value,
	scopeVar string,
	rhs *sysl.Expr,
) *sysl.Value {
	setResult := MakeValueSet()
	for _, l := range list.GetSet().Value {
		if l.GetMap() == nil {
			panic(errors.Errorf("flattenSetMap: expecting map instead of %v ", l))
		}
		assign[scopeVar] = l
		res := Eval(ee, assign, rhs)
		switch x := res.Value.(type) {
		case *sysl.Value_Set:
			for _, ll := range x.Set.Value {
				AppendItemToValueList(setResult.GetSet(), ll)
			}
		default:
			AppendItemToValueList(setResult.GetSet(), res)
		}
	}
	return setResult
}

func flattenSetSet(ee *exprEval,
	assign Scope,
	list *sysl.Value,
	scopeVar string,
	rhs *sysl.Expr,
) *sysl.Value {
	setResult := MakeValueSet()
	for _, l := range list.GetSet().Value {
		for _, ll := range l.GetSet().Value {
			assign[scopeVar] = ll
			AppendItemToValueList(setResult.GetSet(), Eval(ee, assign, rhs))
		}
	}
	return setResult
}

func concat(lhs, rhs *sysl.Value_List) *sysl.Value {
	result := MakeValueList()
	{
		result := result.GetList()
		result.Value = lhs.Value
		result.Value = append(result.Value, rhs.Value...)
		logrus.Tracef("concatList: lhs %d | rhs %d = %d\n", len(lhs.Value), len(rhs.Value), len(result.Value))
	}
	return result
}

func concatListList(lhs, rhs *sysl.Value) *sysl.Value {
	return concat(lhs.GetList(), rhs.GetList())
}

func concatListSet(lhs, rhs *sysl.Value) *sysl.Value {
	return concat(lhs.GetList(), rhs.GetSet())
}

func setUnion(lhs, rhs *sysl.Value) *sysl.Value {
	itemType := getContainedType(lhs)
	if itemType == ValueNoArg {
		itemType = getContainedType(rhs)
	}

	if itemType == ValueNoArg {
		return MakeValueSet()
	}

	switch itemType {
	case ValueInt:
		unionSet := unionIntSets(intSet(lhs.GetSet().Value), intSet(rhs.GetSet().Value))
		logrus.Tracef("Union set: lhs %d, rhs %d res %d\n", len(lhs.GetSet().Value), len(rhs.GetSet().Value), len(unionSet))
		return intSetToValueSet(unionSet)
	case ValueString:
		unionSet := unionStringSets(stringSet(lhs.GetSet().Value), stringSet(rhs.GetSet().Value))
		logrus.Tracef("Union set: lhs %d, rhs %d res %d\n", len(lhs.GetSet().Value), len(rhs.GetSet().Value), len(unionSet))
		return stringSetToValueSet(unionSet)
	case ValueMap:
		unionSet := unionMapSets(mapSet(lhs.GetSet().Value), mapSet(rhs.GetSet().Value))
		logrus.Tracef("Union set: lhs %d, rhs %d res %d\n", len(lhs.GetSet().Value), len(rhs.GetSet().Value), len(unionSet))
		return mapSetToValueSet(unionSet)
	}
	panic(errors.Errorf("union of itemType: %s not supported", itemType.String()))
}

func stringInSet(lhs, rhs *sysl.Value) *sysl.Value {
	str := lhs.GetS()
	for _, v := range rhs.GetSet().Value {
		if str == v.GetS() {
			return MakeValueBool(true)
		}
	}
	return MakeValueBool(false)
}

func stringInMapKey(lhs, rhs *sysl.Value) *sysl.Value {
	str := lhs.GetS()
	_, has := rhs.GetMap().Items[str]
	return MakeValueBool(has)
}

func intSet(list []*sysl.Value) map[int64]struct{} {
	m := map[int64]struct{}{}
	for _, item := range list {
		m[item.GetI()] = struct{}{}
	}
	return m
}

func unionIntSets(lhs, rhs map[int64]struct{}) map[int64]struct{} {
	for key := range rhs {
		lhs[key] = struct{}{}
	}
	return lhs
}

func intSetToValueSet(lhs map[int64]struct{}) *sysl.Value {
	m := MakeValueSet()
	keys := make([]int, 0, len(lhs))
	for key := range lhs {
		keys = append(keys, int(key))
	}
	sort.Ints(keys)

	for _, key := range keys {
		AppendItemToValueList(m.GetSet(), MakeValueI64(int64(key)))
	}
	return m
}

func stringSet(list []*sysl.Value) map[string]struct{} {
	m := map[string]struct{}{}

	for _, item := range list {
		m[item.GetS()] = struct{}{}
	}
	return m
}

func unionStringSets(lhs, rhs map[string]struct{}) map[string]struct{} {
	for key := range rhs {
		lhs[key] = struct{}{}
	}
	return lhs
}

func stringSetToValueSet(lhs map[string]struct{}) *sysl.Value {
	m := MakeValueSet()

	// for stable output
	keys := make([]string, 0, len(lhs))
	for key := range lhs {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		AppendItemToValueList(m.GetSet(), MakeValueString(key))
	}
	return m
}

func mapSet(m []*sysl.Value) map[string]*sysl.Value {
	resultMap := map[string]*sysl.Value{}
	for _, item := range m {
		// Marshal() sorts the keys, so should get stable output.
		bytes, err := json.Marshal(item.GetMap().Items)
		if err == nil {
			resultMap[string(bytes)] = item
		}
	}
	return resultMap
}

func unionMapSets(lhs, rhs map[string]*sysl.Value) map[string]*sysl.Value {
	for key, val := range rhs {
		lhs[key] = val
	}
	return lhs
}

func mapSetToValueSet(lhs map[string]*sysl.Value) *sysl.Value {
	m := MakeValueSet()
	keys := make([]string, 0, len(lhs))
	for key := range lhs {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		AppendItemToValueList(m.GetSet(), lhs[key])
	}
	return m
}

func stringInNull(lhs, rhs *sysl.Value) *sysl.Value {
	return MakeValueBool(false)
}

func stringInList(lhs, rhs *sysl.Value) *sysl.Value {
	str := lhs.GetS()
	for _, v := range rhs.GetList().Value {
		if str == v.GetS() {
			return MakeValueBool(true)
		}
	}
	return MakeValueBool(false)
}

func whereSet(ee *exprEval, assign Scope, list *sysl.Value, scopeVar string, rhs *sysl.Expr) *sysl.Value {
	setResult := MakeValueSet()
	for _, l := range list.GetSet().Value {
		assign[scopeVar] = l
		predicate := Eval(ee, assign, rhs)
		if predicate.GetB() {
			AppendItemToValueList(setResult.GetSet(), l)
		}
	}
	return setResult
}

func whereList(ee *exprEval, assign Scope, list *sysl.Value, scopeVar string, rhs *sysl.Expr) *sysl.Value {
	listResult := MakeValueList()
	logrus.Printf("scope: %s, list len: %d", scopeVar, len(list.GetList().Value))
	for _, l := range list.GetList().Value {
		assign[scopeVar] = l
		predicate := Eval(ee, assign, rhs)
		if predicate.GetB() {
			AppendItemToValueList(listResult.GetList(), l)
		}
	}
	return listResult
}

func whereMap(ee *exprEval, assign Scope, m *sysl.Value, scopeVar string, rhs *sysl.Expr) *sysl.Value {
	mapResult := MakeValueMap()
	logrus.Printf("scope: %s, list len: %d", scopeVar, len(m.GetMap().Items))
	for key, val := range m.GetMap().Items {
		m := MakeValueMap()
		AddItemToValueMap(m, "key", MakeValueString(key))
		AddItemToValueMap(m, "value", val)
		assign[scopeVar] = m
		predicate := Eval(ee, assign, rhs)
		if predicate.GetB() {
			AddItemToValueMap(mapResult, key, val)
		}
	}
	return mapResult
}
