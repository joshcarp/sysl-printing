package cmdutils

import (
	"reflect"
	"testing"

	sysl "github.com/joshcarp/sysl-printing/pkg/sysl"
	"github.com/stretchr/testify/assert"
)

func TestTransformBlackBoxes(t *testing.T) {
	t.Parallel()

	type args struct {
		blackboxes []*sysl.Attribute
	}

	eltFirst := []*sysl.Attribute{
		{Attribute: &sysl.Attribute_S{S: "Value A"}},
		{Attribute: &sysl.Attribute_S{S: "Value B"}},
	}
	attrFirst := &sysl.Attribute{
		Attribute: &sysl.Attribute_A{A: &sysl.Attribute_Array{Elt: eltFirst}},
	}
	eltSecond := []*sysl.Attribute{
		{Attribute: &sysl.Attribute_S{S: "Value C"}},
		{Attribute: &sysl.Attribute_S{S: "Value D"}},
	}
	attrSecond := &sysl.Attribute{
		Attribute: &sysl.Attribute_A{A: &sysl.Attribute_Array{Elt: eltSecond}},
	}

	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "Case-Null",
			args: args{blackboxes: []*sysl.Attribute{}},
			want: [][]string{},
		},
		{
			name: "Case-ConvertSuccess",
			args: args{blackboxes: []*sysl.Attribute{attrFirst, attrSecond}},
			want: [][]string{{"Value A", "Value B"}, {"Value C", "Value D"}},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := TransformBlackBoxes(tt.args.blackboxes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransformBlackBoxes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseBlackBoxesFromArgument(t *testing.T) {
	t.Parallel()

	type args struct {
		blackboxFlags map[string]string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "Case-Null",
			args: args{map[string]string{}},
			want: [][]string{},
		},
		{
			name: "Case-ConvertSuccess",
			args: args{
				map[string]string{"Value A": "Value B", "Value C": "Value D"},
			},
			want: [][]string{{"Value A", "Value B"}, {"Value C", "Value D"}},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := ParseBlackBoxesFromArgument(tt.args.blackboxFlags)
			assert.ElementsMatch(t, got, tt.want)
		})
	}
}

func TestMergeAttributes(t *testing.T) {
	t.Parallel()

	type args struct {
		app   map[string]*sysl.Attribute
		edpnt map[string]*sysl.Attribute
	}

	appAttr := &sysl.Attribute{
		Attribute: &sysl.Attribute_S{
			S: "Value A",
		},
	}
	appMap := map[string]*sysl.Attribute{
		"app": appAttr,
	}
	epAttr := &sysl.Attribute{
		Attribute: &sysl.Attribute_S{
			S: "Value B",
		},
	}
	epMap := map[string]*sysl.Attribute{
		"ep": epAttr,
	}
	tests := []struct {
		name string
		args args
		want map[string]*sysl.Attribute
	}{
		{
			"Case-Null",
			args{},
			map[string]*sysl.Attribute{},
		},
		{
			"Case-Merge app",
			args{appMap, map[string]*sysl.Attribute{}},
			map[string]*sysl.Attribute{
				"app": appAttr,
			},
		},
		{
			"Case-Merge ep",
			args{map[string]*sysl.Attribute{}, epMap},
			map[string]*sysl.Attribute{
				"ep": epAttr,
			},
		},
		{
			"Case-Merge app and ep",
			args{appMap, epMap},
			map[string]*sysl.Attribute{
				"app": appAttr,
				"ep":  epAttr,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeAttributes(tt.args.app, tt.args.edpnt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeAttributes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransformBlackboxesToUptos(t *testing.T) {
	t.Parallel()

	// given
	bbs := map[string]*Upto{}
	m := [][]string{
		{"keyA", "value A"},
	}

	// when
	TransformBlackboxesToUptos(bbs, m, BBApplication)

	// then
	assert.Equal(t, m[0][1], bbs[m[0][0]].Comment)
}

func TestTransformBlackboxesToUptosByNil(t *testing.T) {
	t.Parallel()

	// given
	bbs := map[string]*Upto{}
	var m [][]string

	// when
	TransformBlackboxesToUptos(bbs, m, BBApplication)

	// then
	assert.Empty(t, bbs)
}

func TestGetAppAttr(t *testing.T) {
	t.Parallel()

	// given
	attr := map[string]*sysl.Attribute{
		"attr1": {},
	}
	m := &sysl.Module{
		Apps: map[string]*sysl.Application{
			"test": {Attrs: attr},
		},
	}

	// when
	actual := GetApplicationAttrs(m, "test")

	// then
	assert.Equal(t, attr, actual)
}

func TestGetAppAttrWhenAppNotExist(t *testing.T) {
	t.Parallel()

	// given
	m := &sysl.Module{
		Apps: make(map[string]*sysl.Application),
	}

	// when
	actual := GetApplicationAttrs(m, "test")

	// then
	assert.Nil(t, actual)
}

func TestSortedISOCtrlSlice(t *testing.T) {
	t.Parallel()

	// given
	attrs := map[string]*sysl.Attribute{
		"iso_ctrl_11_txt": {},
		"iso_ctrl_12_txt": {},
		"iso_ctrl_5_txt":  {},
	}

	// when
	actual := GetSortedISOCtrlSlice(attrs)

	// then
	assert.Equal(t, []string{"11", "12", "5"}, actual)
}

func TestSortedISOCtrlSliceEmpty(t *testing.T) {
	t.Parallel()

	// given
	attrs := make(map[string]*sysl.Attribute)

	// when
	actual := GetSortedISOCtrlSlice(attrs)

	// then
	assert.Equal(t, []string{}, actual)
}

func TestSortedISOCtrlStr(t *testing.T) {
	t.Parallel()

	// given
	attrs := map[string]*sysl.Attribute{
		"iso_ctrl_11_txt": {},
		"iso_ctrl_12_txt": {},
		"iso_ctrl_5_txt":  {},
	}

	// when
	actual := GetSortedISOCtrlStr(attrs)

	// then
	assert.Equal(t, "11, 12, 5", actual)
}

func TestSortedISOCtrlStrEmpty(t *testing.T) {
	t.Parallel()

	// given
	attrs := make(map[string]*sysl.Attribute)

	// when
	actual := GetSortedISOCtrlStr(attrs)

	// then
	assert.Equal(t, "", actual)
}

func TestFormatArgs(t *testing.T) {
	t.Parallel()

	// given
	m := &sysl.Module{
		Apps: map[string]*sysl.Application{
			"test": {
				Types: map[string]*sysl.Type{
					"User": {
						Attrs: map[string]*sysl.Attribute{
							"iso_conf": {
								Attribute: &sysl.Attribute_S{
									S: "Red",
								},
							},
							"iso_integ": {
								Attribute: &sysl.Attribute_S{
									S: "I",
								},
							},
						},
					},
				},
			},
		},
	}

	// when
	actual := FormatArgs(m, "test", "User")

	assert.Equal(t, "<color blue>test.User</color> <<color red>R, I</color>>", actual)
}

func TestFormatArgsWithoutIsoInteg(t *testing.T) {
	t.Parallel()

	// given
	m := &sysl.Module{
		Apps: map[string]*sysl.Application{
			"test": {
				Types: map[string]*sysl.Type{
					"User": {
						Attrs: map[string]*sysl.Attribute{
							"iso_conf": {
								Attribute: &sysl.Attribute_S{
									S: "Red",
								},
							},
						},
					},
				},
			},
		},
	}

	// when
	actual := FormatArgs(m, "test", "User")

	assert.Equal(t, "<color blue>test.User</color> <<color red>R, ?</color>>", actual)
}

func TestFormatArgsWithoutIsoConf(t *testing.T) {
	t.Parallel()

	// given
	m := &sysl.Module{
		Apps: map[string]*sysl.Application{
			"test": {
				Types: map[string]*sysl.Type{
					"User": {
						Attrs: map[string]*sysl.Attribute{
							"iso_integ": {
								Attribute: &sysl.Attribute_S{
									S: "I",
								},
							},
						},
					},
				},
			},
		},
	}

	// when
	actual := FormatArgs(m, "test", "User")

	assert.Equal(t, "<color blue>test.User</color> <<color green>?, I</color>>", actual)
}

func TestFormatArgsWithoutAttrs(t *testing.T) {
	t.Parallel()

	// given
	m := &sysl.Module{
		Apps: map[string]*sysl.Application{
			"test": {
				Types: map[string]*sysl.Type{
					"User": {
						Attrs: make(map[string]*sysl.Attribute),
					},
				},
			},
		},
	}

	// when
	actual := FormatArgs(m, "test", "User")

	assert.Equal(t, "<color blue>test.User</color> <<color green>?, ?</color>>", actual)
}

func TestFormatArgsWithoutParameterTypeName(t *testing.T) {
	t.Parallel()

	// given
	m := &sysl.Module{
		Apps: map[string]*sysl.Application{
			"test": {
				Types: map[string]*sysl.Type{
					"User": {
						Attrs: make(map[string]*sysl.Attribute),
					},
				},
			},
		},
	}

	// when
	actual := FormatArgs(m, "test", "")

	assert.Equal(t, "<color blue>test.</color> <<color green>?, ?</color>>", actual)
}

func TestFormatArgsWithoutAppName(t *testing.T) {
	t.Parallel()

	// given
	m := &sysl.Module{
		Apps: map[string]*sysl.Application{
			"test": {
				Types: map[string]*sysl.Type{
					"User": {
						Attrs: make(map[string]*sysl.Attribute),
					},
				},
			},
		},
	}

	// when
	actual := FormatArgs(m, "", "User")

	assert.Equal(t, "<color blue>.User</color> <<color green>?, ?</color>>", actual)
}

func TestFormatArgsWithoutAppNameAndParameterTypeName(t *testing.T) {
	t.Parallel()

	// given
	m := &sysl.Module{
		Apps: map[string]*sysl.Application{
			"test": {
				Types: map[string]*sysl.Type{
					"User": {
						Attrs: make(map[string]*sysl.Attribute),
					},
				},
			},
		},
	}

	// when
	actual := FormatArgs(m, "", "")

	assert.Equal(t, "<color blue>.</color> <<color green>?, ?</color>>", actual)
}

func TestFormatReturnParam(t *testing.T) {
	t.Parallel()

	// given
	m := &sysl.Module{
		Apps: map[string]*sysl.Application{
			"test": {
				Types: map[string]*sysl.Type{
					"User": {
						Attrs: make(map[string]*sysl.Attribute),
					},
				},
			},
		},
	}

	// when
	actual := FormatReturnParam(m, "test.User")

	assert.Equal(t, []string{"<color blue>test.User</color> <<color green>?, ?</color>>"}, actual)
}

func TestFormatReturnParamSplit(t *testing.T) {
	t.Parallel()

	// given
	m := &sysl.Module{
		Apps: map[string]*sysl.Application{
			"test": {
				Types: map[string]*sysl.Type{
					"User": {
						Attrs: make(map[string]*sysl.Attribute),
					},
				},
			},
		},
	}

	// when
	actual := FormatReturnParam(m, "test.User,profile<:test.User,Bool,set of test.User, one of {test.User,ab}")

	expected := []string{
		"<color blue>test.User</color> <<color green>?, ?</color>>",
		"<color blue>test.User</color> <<color green>?, ?</color>>",
		"<color blue>test.User</color> <<color green>?, ?</color>>",
		"<color blue>test.User</color> <<color green>?, ?</color>>",
		"ab",
	}

	assert.Equal(t, expected, actual)
}

func TestGetReturnPayload(t *testing.T) {
	t.Parallel()

	stmts := []*sysl.Statement{
		{
			Stmt: &sysl.Statement_Call{},
		},
		{
			Stmt: &sysl.Statement_Action{},
		},
		{
			Stmt: &sysl.Statement_Ret{
				Ret: &sysl.Return{
					Payload: "test",
				},
			},
		},
	}

	actual := GetReturnPayload(stmts)

	assert.Equal(t, "test", actual)
}

func TestGetReturnPayloadWithAlt(t *testing.T) {
	t.Parallel()

	stmts := []*sysl.Statement{
		{
			Stmt: &sysl.Statement_Alt{
				Alt: &sysl.Alt{
					Choice: []*sysl.Alt_Choice{
						{
							Cond: "cond 1",
							Stmt: []*sysl.Statement{},
						},
						{
							Cond: "cond 2",
							Stmt: []*sysl.Statement{
								{
									Stmt: &sysl.Statement_Ret{
										Ret: &sysl.Return{
											Payload: "test",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	actual := GetReturnPayload(stmts)

	assert.Equal(t, "test", actual)
}

func TestGetReturnPayloadWithCond(t *testing.T) {
	t.Parallel()

	stmts := []*sysl.Statement{
		{
			Stmt: &sysl.Statement_Cond{
				Cond: &sysl.Cond{
					Test: "cond 1",
					Stmt: []*sysl.Statement{
						{
							Stmt: &sysl.Statement_Ret{
								Ret: &sysl.Return{
									Payload: "test",
								},
							},
						},
					},
				},
			},
		},
	}

	actual := GetReturnPayload(stmts)

	assert.Equal(t, "test", actual)
}

func TestGetReturnPayloadWithLoop(t *testing.T) {
	t.Parallel()

	stmts := []*sysl.Statement{
		{
			Stmt: &sysl.Statement_Loop{
				Loop: &sysl.Loop{
					Mode:      sysl.Loop_WHILE,
					Criterion: "criterion",
					Stmt: []*sysl.Statement{
						{
							Stmt: &sysl.Statement_Ret{
								Ret: &sysl.Return{
									Payload: "test",
								},
							},
						},
					},
				},
			},
		},
	}

	actual := GetReturnPayload(stmts)

	assert.Equal(t, "test", actual)
}

func TestGetReturnPayloadWithLoopN(t *testing.T) {
	t.Parallel()

	stmts := []*sysl.Statement{
		{
			Stmt: &sysl.Statement_LoopN{
				LoopN: &sysl.LoopN{
					Count: 10,
					Stmt: []*sysl.Statement{
						{
							Stmt: &sysl.Statement_Ret{
								Ret: &sysl.Return{
									Payload: "test",
								},
							},
						},
					},
				},
			},
		},
	}

	actual := GetReturnPayload(stmts)

	assert.Equal(t, "test", actual)
}

func TestGetReturnPayloadWithForeach(t *testing.T) {
	t.Parallel()

	stmts := []*sysl.Statement{
		{
			Stmt: &sysl.Statement_Foreach{
				Foreach: &sysl.Foreach{
					Collection: "collection 1",
					Stmt: []*sysl.Statement{
						{
							Stmt: &sysl.Statement_Ret{
								Ret: &sysl.Return{
									Payload: "test",
								},
							},
						},
					},
				},
			},
		},
	}

	actual := GetReturnPayload(stmts)

	assert.Equal(t, "test", actual)
}

func TestGetReturnPayloadWithGroup(t *testing.T) {
	t.Parallel()

	stmts := []*sysl.Statement{
		{
			Stmt: &sysl.Statement_Group{
				Group: &sysl.Group{
					Title: "group 1",
					Stmt: []*sysl.Statement{
						{
							Stmt: &sysl.Statement_Ret{
								Ret: &sysl.Return{
									Payload: "test",
								},
							},
						},
					},
				},
			},
		},
	}

	actual := GetReturnPayload(stmts)

	assert.Equal(t, "test", actual)
}

func TestGetAndFmtParam(t *testing.T) {
	t.Parallel()

	m := &sysl.Module{
		Apps: map[string]*sysl.Application{
			"test": {
				Types: map[string]*sysl.Type{
					"User": {
						Attrs: map[string]*sysl.Attribute{
							"iso_conf": {
								Attribute: &sysl.Attribute_S{
									S: "Red",
								},
							},
							"iso_integ": {
								Attribute: &sysl.Attribute_S{
									S: "I",
								},
							},
						},
					},
				},
			},
		},
	}

	p := []*sysl.Param{
		{
			Name: "profile",
			Type: &sysl.Type{
				Type: &sysl.Type_TypeRef{
					TypeRef: &sysl.ScopedRef{
						Ref: &sysl.Scope{
							Appname: &sysl.AppName{
								Part: []string{"test"},
							},
							Path: []string{"User"},
						},
					},
				},
			},
		},
	}

	actual := GetAndFmtParam(m, p)

	assert.Equal(t, []string{"<color blue>test.User</color> <<color red>R, I</color>>"}, actual)
}

func TestNormalizeEndpointName(t *testing.T) {
	t.Parallel()

	actual := NormalizeEndpointName("a -> b")

	assert.Equal(t, " ⬄ b", actual)
}
