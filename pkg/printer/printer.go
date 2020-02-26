package printer

import "C"
import (
	"fmt"
	"strings"

	"github.com/anz-bank/sysl/pkg/sysl"
)

func PrintModule(mod *sysl.Module) {
	for _, A := range mod.Apps {
		PrintApplication(A)
	}
}

func PrintParam(param []*sysl.Param) {
	params := "("
	for i, p := range param {
		params += p.Name + " <: " + ParamType(p)
		if i != len(param)-1 {
			params += ","
		}
	}
	params += ")"
	fmt.Print(params)
}

func PrintApplication(A *sysl.Application) {
	fmt.Printf("%s:\n", A.Name.GetPart()[0])
	for key, val := range A.Attrs {
		PrintAttrs(key, val)
	}
	for _, e := range A.Endpoints {
		PrintEndpoint(e)

	}
}

func PrintEndpoint(E *sysl.Endpoint) {
	fmt.Printf("    %s", E.Name)

	if len(E.Param) != 0 {
		PrintParam(E.Param)
	}
	fmt.Printf(":\n")

	for _, stmnt := range E.Stmt {
		PrintStatement(stmnt)
	}
}

func PrintStatement(S *sysl.Statement) {
	if call := S.GetCall(); call != nil {
		PrintCall(call)
	}
	if action := S.GetAction(); action != nil {
		PrintAction(action)
	}
}

func PrintAction(A *sysl.Action) {
	fmt.Println(A.GetAction())
}

func PrintAttrs(key string, A *sysl.Attribute) {
	fmt.Printf(`    @%s="%s"`, key, A.GetS())
	fmt.Println()
}

func ParamType(P *sysl.Param) string {
	return strings.Join(P.Type.GetTypeRef().Ref.Appname.Part, "")
}

func PrintType(T *sysl.Type) string {
	return strings.Join(T.GetTypeRef().Ref.Appname.Part, "")
}

func PrintCall(C *sysl.Call) {
	fmt.Printf("        %s <- %s\n", C.Target.Part[0], C.Endpoint)
}
