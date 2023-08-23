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
	// ServerChassisTypePrefix is the prefix for all ServerChassisType nodes
	ServerChassisTypePrefix = ApplicationPrefix + "sct"
	// ServerChassisPrefix is the prefix for all ServerChassis nodes
	ServerChassisPrefix = ApplicationPrefix + "sch"
	// ServerCPUTypePrefix is the prefix for all ServerCPUType nodes
	ServerCPUTypePrefix = ApplicationPrefix + "cpt"
	// ServerCPUPrefix is the prefix for all ServerCPU nodes
	ServerCPUPrefix = ApplicationPrefix + "cpu"
	// ServerMotherboardTypePrefix is the prefix for all ServerMotherboardType nodes
	ServerMotherboardTypePrefix = ApplicationPrefix + "mbt"
	// ServerMotherboardPrefix is the prefix for all ServerMotherboard nodes
	ServerMotherboardPrefix = ApplicationPrefix + "mbd"
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
