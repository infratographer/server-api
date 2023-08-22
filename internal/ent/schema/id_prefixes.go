package schema

import (
	"entgo.io/contrib/entgql"
	"github.com/vektah/gqlparser/v2/ast"
)

const (
	// ApplicationPrefix is the prefix for all Server nodes
	ApplicationPrefix = "srvr"
	// ServerPrefix is the prefix for all Server nodes
	ServerPrefix = ApplicationPrefix + "srv"
	// ServerProviderPrefix is the prefix for all ServerProvider nodes
	ServerProviderPrefix = ApplicationPrefix + "prv"
	// ServerTypePrefix is the prefix for all ServerType nodes
	ServerTypePrefix = ApplicationPrefix + "typ"
	// ServerComponentTypePrefix is the prefix for all ServerComponentType nodes
	ServerComponentTypePrefix = ApplicationPrefix + "cpt"
	// ServerComponentPrefix is the prefix for all ServerComponent nodes
	ServerComponentPrefix = ApplicationPrefix + "cmp"
)

func prefixIDDirective(prefix string) entgql.Annotation {
	var args []*ast.Argument
	if prefix != "" {
		args = append(args, &ast.Argument{
			Name: "prefix",
			Value: &ast.Value{
				Raw:  prefix,
				Kind: ast.StringValue,
			},
		})
	}

	return entgql.Directives(entgql.NewDirective("prefixedID", args...))
}
