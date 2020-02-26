package printer

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/joshcarp/sysl-printing/pkg/syslutil"

	"github.com/joshcarp/sysl-printing/pkg/sysl"
)

type Printer struct {
	io.Writer
}

func NewPrinter(buf *bytes.Buffer) *Printer {
	return &Printer{Writer: buf}
}

//PrintModule Prints a whole module, calling
func (p *Printer) PrintModule(mod *sysl.Module) {
	for _, A := range mod.Apps {
		p.PrintApplication(A)
	}
}

// Endpoint(This <: ParamHere):
func (p *Printer) PrintParam(params []*sysl.Param) {
	ans := "("
	for i, param := range params {
		ans += param.Name + " <: " + p.ParamType(param)
		if i != len(params)-1 {
			ans += ","
		}
	}
	ans += ")"
	fmt.Fprint(p.Writer, ans)
}

// App:
func (p *Printer) PrintApplication(A *sysl.Application) {
	fmt.Fprintf(p.Writer, "%s:\n", A.Name.GetPart()[0])
	for key, val := range A.Attrs {
		p.PrintAttrs(key, val)
	}
	for _, e := range A.Endpoints {
		p.PrintEndpoint(e)
	}
	for typeName, t := range A.Types {
		p.PrintTypeDecl(typeName, t)
	}
}

// Endpoint:
func (p *Printer) PrintEndpoint(E *sysl.Endpoint) {
	fmt.Fprintf(p.Writer, "    %s", E.Name)

	if len(E.Param) != 0 {
		p.PrintParam(E.Param)
	}
	fmt.Fprintf(p.Writer, ":\n")

	for _, stmnt := range E.Stmt {
		p.PrintStatement(stmnt)
	}
}

// !type Foo:
//     this <: string
func (p *Printer) PrintTypeDecl(key string, t *sysl.Type) {
	fmt.Fprintf(p.Writer, "    !type %s:\n", key)
	if tuple := t.GetTuple(); tuple != nil {
		for key, val := range tuple.AttrDefs {
			typeClass, typeIdent := syslutil.GetTypeDetail(val)
			if typeClass == "primitive" {
				typeIdent = strings.ToLower(typeIdent)
			}
			fmt.Fprintf(p.Writer, "        %s <: %s\n", key, typeIdent)
		}
	}

}

// return ret <: string
func (p *Printer) PrintStatement(S *sysl.Statement) {
	if call := S.GetCall(); call != nil {
		p.PrintCall(call)
	}
	if action := S.GetAction(); action != nil {
		p.PrintAction(action)
	}
	if ret := S.GetRet(); ret != nil {
		p.PrintReturn(ret)
	}
}

// return foo <: type
func (p *Printer) PrintReturn(R *sysl.Return) {
	fmt.Fprintf(p.Writer, "        return ret <: %s\n", R.Payload)
}

// lookup data
func (p *Printer) PrintAction(A *sysl.Action) {
	fmt.Fprintf(p.Writer, "        %s\n", A.GetAction())
}

// @owner="server"
func (p *Printer) PrintAttrs(key string, A *sysl.Attribute) {
	fmt.Fprintf(p.Writer, "    @%s=\"%s\"\n", key, A.GetS())
}

//foo(this <: <ParamType>):
func (p *Printer) ParamType(P *sysl.Param) string {
	if P.Type == nil {
		return ""
	}
	if P.Type.GetTypeRef() == nil {
		return ""
	}
	return strings.Join(P.Type.GetTypeRef().Ref.Appname.Part, "")
}

// AnApp <- AnEndpoint
func (p *Printer) PrintCall(c *sysl.Call) {
	fmt.Fprintf(p.Writer, "        %s <- %s\n", strings.Join(c.Target.GetPart(), ""), c.GetEndpoint())
}
