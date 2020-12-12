// Code generated by scripts/gengraphql.go. DO NOT EDIT.

package schema

import (
	errors "errors"
	graphql1 "github.com/graphql-go/graphql"
	graphql "github.com/sensu/sensu-go/graphql"
)

//
// RuleFieldResolvers represents a collection of methods whose products represent the
// response values of the 'Rule' type.
type RuleFieldResolvers interface {
	// Verbs implements response to request for 'verbs' field.
	Verbs(p graphql.ResolveParams) ([]string, error)

	// Resources implements response to request for 'resources' field.
	Resources(p graphql.ResolveParams) ([]string, error)

	// ResourceNames implements response to request for 'resourceNames' field.
	ResourceNames(p graphql.ResolveParams) ([]string, error)
}

// RuleAliases implements all methods on RuleFieldResolvers interface by using reflection to
// match name of field to a field on the given value. Intent is reduce friction
// of writing new resolvers by removing all the instances where you would simply
// have the resolvers method return a field.
type RuleAliases struct{}

// Verbs implements response to request for 'verbs' field.
func (_ RuleAliases) Verbs(p graphql.ResolveParams) ([]string, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := val.([]string)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'verbs'")
	}
	return ret, err
}

// Resources implements response to request for 'resources' field.
func (_ RuleAliases) Resources(p graphql.ResolveParams) ([]string, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := val.([]string)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'resources'")
	}
	return ret, err
}

// ResourceNames implements response to request for 'resourceNames' field.
func (_ RuleAliases) ResourceNames(p graphql.ResolveParams) ([]string, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := val.([]string)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'resourceNames'")
	}
	return ret, err
}

// RuleType Rule holds information that describes an action that can be taken
var RuleType = graphql.NewType("Rule", graphql.ObjectKind)

// RegisterRule registers Rule object type with given service.
func RegisterRule(svc *graphql.Service, impl RuleFieldResolvers) {
	svc.RegisterObject(_ObjectTypeRuleDesc, impl)
}
func _ObjTypeRuleVerbsHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(RuleVerbsFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Verbs(frp)
	}
}

func _ObjTypeRuleResourcesHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(RuleResourcesFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Resources(frp)
	}
}

func _ObjTypeRuleResourceNamesHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(RuleResourceNamesFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.ResourceNames(frp)
	}
}

func _ObjectTypeRuleConfigFn() graphql1.ObjectConfig {
	return graphql1.ObjectConfig{
		Description: "Rule holds information that describes an action that can be taken",
		Fields: graphql1.Fields{
			"resourceNames": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "ResourceNames is an optional list of resource names that the rule applies\nto.",
				Name:              "resourceNames",
				Type:              graphql1.NewNonNull(graphql1.NewList(graphql1.NewNonNull(graphql1.String))),
			},
			"resources": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Resources is a list of resources that this rule applies to. \"*\" represents\nall resources.",
				Name:              "resources",
				Type:              graphql1.NewNonNull(graphql1.NewList(graphql1.NewNonNull(graphql1.String))),
			},
			"verbs": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Verbs is a list of verbs that apply to all of the listed resources for this\nrule. These include \"get\", \"list\", \"watch\", \"create\", \"update\", \"delete\".\nTODO: add support for \"patch\" (this is expensive and should be delayed\nuntil a further release). TODO: add support for \"watch\" (via websockets)",
				Name:              "verbs",
				Type:              graphql1.NewNonNull(graphql1.NewList(graphql1.NewNonNull(graphql1.String))),
			},
		},
		Interfaces: []*graphql1.Interface{},
		IsTypeOf: func(_ graphql1.IsTypeOfParams) bool {
			// NOTE:
			// Panic by default. Intent is that when Service is invoked, values of
			// these fields are updated with instantiated resolvers. If these
			// defaults are called it is most certainly programmer err.
			// If you're see this comment then: 'Whoops! Sorry, my bad.'
			panic("Unimplemented; see RuleFieldResolvers.")
		},
		Name: "Rule",
	}
}

// describe Rule's configuration; kept private to avoid unintentional tampering of configuration at runtime.
var _ObjectTypeRuleDesc = graphql.ObjectDesc{
	Config: _ObjectTypeRuleConfigFn,
	FieldHandlers: map[string]graphql.FieldHandler{
		"resourceNames": _ObjTypeRuleResourceNamesHandler,
		"resources":     _ObjTypeRuleResourcesHandler,
		"verbs":         _ObjTypeRuleVerbsHandler,
	},
}

//
// ClusterRoleFieldResolvers represents a collection of methods whose products represent the
// response values of the 'ClusterRole' type.
type ClusterRoleFieldResolvers interface {
	// Rules implements response to request for 'rules' field.
	Rules(p graphql.ResolveParams) (interface{}, error)

	// Name implements response to request for 'name' field.
	Name(p graphql.ResolveParams) (string, error)
}

// ClusterRoleAliases implements all methods on ClusterRoleFieldResolvers interface by using reflection to
// match name of field to a field on the given value. Intent is reduce friction
// of writing new resolvers by removing all the instances where you would simply
// have the resolvers method return a field.
type ClusterRoleAliases struct{}

// Rules implements response to request for 'rules' field.
func (_ ClusterRoleAliases) Rules(p graphql.ResolveParams) (interface{}, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	return val, err
}

// Name implements response to request for 'name' field.
func (_ ClusterRoleAliases) Name(p graphql.ResolveParams) (string, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := val.(string)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'name'")
	}
	return ret, err
}

// ClusterRoleType ClusterRole applies to all namespaces within a cluster.
var ClusterRoleType = graphql.NewType("ClusterRole", graphql.ObjectKind)

// RegisterClusterRole registers ClusterRole object type with given service.
func RegisterClusterRole(svc *graphql.Service, impl ClusterRoleFieldResolvers) {
	svc.RegisterObject(_ObjectTypeClusterRoleDesc, impl)
}
func _ObjTypeClusterRoleRulesHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(ClusterRoleRulesFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Rules(frp)
	}
}

func _ObjTypeClusterRoleNameHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(ClusterRoleNameFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Name(frp)
	}
}

func _ObjectTypeClusterRoleConfigFn() graphql1.ObjectConfig {
	return graphql1.ObjectConfig{
		Description: "ClusterRole applies to all namespaces within a cluster.",
		Fields: graphql1.Fields{
			"name": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Name of the ClusterRole",
				Name:              "name",
				Type:              graphql1.NewNonNull(graphql1.String),
			},
			"rules": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "self descriptive",
				Name:              "rules",
				Type:              graphql1.NewList(graphql1.NewNonNull(graphql.OutputType("Rule"))),
			},
		},
		Interfaces: []*graphql1.Interface{},
		IsTypeOf: func(_ graphql1.IsTypeOfParams) bool {
			// NOTE:
			// Panic by default. Intent is that when Service is invoked, values of
			// these fields are updated with instantiated resolvers. If these
			// defaults are called it is most certainly programmer err.
			// If you're see this comment then: 'Whoops! Sorry, my bad.'
			panic("Unimplemented; see ClusterRoleFieldResolvers.")
		},
		Name: "ClusterRole",
	}
}

// describe ClusterRole's configuration; kept private to avoid unintentional tampering of configuration at runtime.
var _ObjectTypeClusterRoleDesc = graphql.ObjectDesc{
	Config: _ObjectTypeClusterRoleConfigFn,
	FieldHandlers: map[string]graphql.FieldHandler{
		"name":  _ObjTypeClusterRoleNameHandler,
		"rules": _ObjTypeClusterRoleRulesHandler,
	},
}

//
// RoleFieldResolvers represents a collection of methods whose products represent the
// response values of the 'Role' type.
type RoleFieldResolvers interface {
	// Rules implements response to request for 'rules' field.
	Rules(p graphql.ResolveParams) (interface{}, error)

	// Namespace implements response to request for 'namespace' field.
	Namespace(p graphql.ResolveParams) (string, error)

	// Name implements response to request for 'name' field.
	Name(p graphql.ResolveParams) (string, error)
}

// RoleAliases implements all methods on RoleFieldResolvers interface by using reflection to
// match name of field to a field on the given value. Intent is reduce friction
// of writing new resolvers by removing all the instances where you would simply
// have the resolvers method return a field.
type RoleAliases struct{}

// Rules implements response to request for 'rules' field.
func (_ RoleAliases) Rules(p graphql.ResolveParams) (interface{}, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	return val, err
}

// Namespace implements response to request for 'namespace' field.
func (_ RoleAliases) Namespace(p graphql.ResolveParams) (string, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := val.(string)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'namespace'")
	}
	return ret, err
}

// Name implements response to request for 'name' field.
func (_ RoleAliases) Name(p graphql.ResolveParams) (string, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := val.(string)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'name'")
	}
	return ret, err
}

// RoleType Role applies only to a single namespace.
var RoleType = graphql.NewType("Role", graphql.ObjectKind)

// RegisterRole registers Role object type with given service.
func RegisterRole(svc *graphql.Service, impl RoleFieldResolvers) {
	svc.RegisterObject(_ObjectTypeRoleDesc, impl)
}
func _ObjTypeRoleRulesHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(RoleRulesFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Rules(frp)
	}
}

func _ObjTypeRoleNamespaceHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(RoleNamespaceFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Namespace(frp)
	}
}

func _ObjTypeRoleNameHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(RoleNameFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Name(frp)
	}
}

func _ObjectTypeRoleConfigFn() graphql1.ObjectConfig {
	return graphql1.ObjectConfig{
		Description: "Role applies only to a single namespace.",
		Fields: graphql1.Fields{
			"name": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Name of the Role",
				Name:              "name",
				Type:              graphql1.NewNonNull(graphql1.String),
			},
			"namespace": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Namespace of the Role",
				Name:              "namespace",
				Type:              graphql1.NewNonNull(graphql1.String),
			},
			"rules": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "self descriptive",
				Name:              "rules",
				Type:              graphql1.NewList(graphql1.NewNonNull(graphql.OutputType("Rule"))),
			},
		},
		Interfaces: []*graphql1.Interface{},
		IsTypeOf: func(_ graphql1.IsTypeOfParams) bool {
			// NOTE:
			// Panic by default. Intent is that when Service is invoked, values of
			// these fields are updated with instantiated resolvers. If these
			// defaults are called it is most certainly programmer err.
			// If you're see this comment then: 'Whoops! Sorry, my bad.'
			panic("Unimplemented; see RoleFieldResolvers.")
		},
		Name: "Role",
	}
}

// describe Role's configuration; kept private to avoid unintentional tampering of configuration at runtime.
var _ObjectTypeRoleDesc = graphql.ObjectDesc{
	Config: _ObjectTypeRoleConfigFn,
	FieldHandlers: map[string]graphql.FieldHandler{
		"name":      _ObjTypeRoleNameHandler,
		"namespace": _ObjTypeRoleNamespaceHandler,
		"rules":     _ObjTypeRoleRulesHandler,
	},
}

//
// RoleRefFieldResolvers represents a collection of methods whose products represent the
// response values of the 'RoleRef' type.
type RoleRefFieldResolvers interface {
	// Type implements response to request for 'type' field.
	Type(p graphql.ResolveParams) (string, error)

	// Name implements response to request for 'name' field.
	Name(p graphql.ResolveParams) (string, error)
}

// RoleRefAliases implements all methods on RoleRefFieldResolvers interface by using reflection to
// match name of field to a field on the given value. Intent is reduce friction
// of writing new resolvers by removing all the instances where you would simply
// have the resolvers method return a field.
type RoleRefAliases struct{}

// Type implements response to request for 'type' field.
func (_ RoleRefAliases) Type(p graphql.ResolveParams) (string, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := val.(string)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'type'")
	}
	return ret, err
}

// Name implements response to request for 'name' field.
func (_ RoleRefAliases) Name(p graphql.ResolveParams) (string, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := val.(string)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'name'")
	}
	return ret, err
}

// RoleRefType RoleRef maps groups to Roles or ClusterRoles.
var RoleRefType = graphql.NewType("RoleRef", graphql.ObjectKind)

// RegisterRoleRef registers RoleRef object type with given service.
func RegisterRoleRef(svc *graphql.Service, impl RoleRefFieldResolvers) {
	svc.RegisterObject(_ObjectTypeRoleRefDesc, impl)
}
func _ObjTypeRoleRefTypeHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(RoleRefTypeFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Type(frp)
	}
}

func _ObjTypeRoleRefNameHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(RoleRefNameFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Name(frp)
	}
}

func _ObjectTypeRoleRefConfigFn() graphql1.ObjectConfig {
	return graphql1.ObjectConfig{
		Description: "RoleRef maps groups to Roles or ClusterRoles.",
		Fields: graphql1.Fields{
			"name": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Name of the resource being referenced",
				Name:              "name",
				Type:              graphql1.NewNonNull(graphql1.String),
			},
			"type": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Type of role being referenced.",
				Name:              "type",
				Type:              graphql1.NewNonNull(graphql1.String),
			},
		},
		Interfaces: []*graphql1.Interface{},
		IsTypeOf: func(_ graphql1.IsTypeOfParams) bool {
			// NOTE:
			// Panic by default. Intent is that when Service is invoked, values of
			// these fields are updated with instantiated resolvers. If these
			// defaults are called it is most certainly programmer err.
			// If you're see this comment then: 'Whoops! Sorry, my bad.'
			panic("Unimplemented; see RoleRefFieldResolvers.")
		},
		Name: "RoleRef",
	}
}

// describe RoleRef's configuration; kept private to avoid unintentional tampering of configuration at runtime.
var _ObjectTypeRoleRefDesc = graphql.ObjectDesc{
	Config: _ObjectTypeRoleRefConfigFn,
	FieldHandlers: map[string]graphql.FieldHandler{
		"name": _ObjTypeRoleRefNameHandler,
		"type": _ObjTypeRoleRefTypeHandler,
	},
}

//
// SubjectFieldResolvers represents a collection of methods whose products represent the
// response values of the 'Subject' type.
type SubjectFieldResolvers interface {
	// Kind implements response to request for 'kind' field.
	Kind(p graphql.ResolveParams) (string, error)

	// Name implements response to request for 'name' field.
	Name(p graphql.ResolveParams) (string, error)
}

// SubjectAliases implements all methods on SubjectFieldResolvers interface by using reflection to
// match name of field to a field on the given value. Intent is reduce friction
// of writing new resolvers by removing all the instances where you would simply
// have the resolvers method return a field.
type SubjectAliases struct{}

// Kind implements response to request for 'kind' field.
func (_ SubjectAliases) Kind(p graphql.ResolveParams) (string, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := val.(string)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'kind'")
	}
	return ret, err
}

// Name implements response to request for 'name' field.
func (_ SubjectAliases) Name(p graphql.ResolveParams) (string, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := val.(string)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'name'")
	}
	return ret, err
}

// SubjectType self descriptive
var SubjectType = graphql.NewType("Subject", graphql.ObjectKind)

// RegisterSubject registers Subject object type with given service.
func RegisterSubject(svc *graphql.Service, impl SubjectFieldResolvers) {
	svc.RegisterObject(_ObjectTypeSubjectDesc, impl)
}
func _ObjTypeSubjectKindHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(SubjectKindFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Kind(frp)
	}
}

func _ObjTypeSubjectNameHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(SubjectNameFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Name(frp)
	}
}

func _ObjectTypeSubjectConfigFn() graphql1.ObjectConfig {
	return graphql1.ObjectConfig{
		Description: "self descriptive",
		Fields: graphql1.Fields{
			"kind": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Kind of object referenced (user or group)",
				Name:              "kind",
				Type:              graphql1.NewNonNull(graphql1.String),
			},
			"name": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Name of the referenced object",
				Name:              "name",
				Type:              graphql1.NewNonNull(graphql1.String),
			},
		},
		Interfaces: []*graphql1.Interface{},
		IsTypeOf: func(_ graphql1.IsTypeOfParams) bool {
			// NOTE:
			// Panic by default. Intent is that when Service is invoked, values of
			// these fields are updated with instantiated resolvers. If these
			// defaults are called it is most certainly programmer err.
			// If you're see this comment then: 'Whoops! Sorry, my bad.'
			panic("Unimplemented; see SubjectFieldResolvers.")
		},
		Name: "Subject",
	}
}

// describe Subject's configuration; kept private to avoid unintentional tampering of configuration at runtime.
var _ObjectTypeSubjectDesc = graphql.ObjectDesc{
	Config: _ObjectTypeSubjectConfigFn,
	FieldHandlers: map[string]graphql.FieldHandler{
		"kind": _ObjTypeSubjectKindHandler,
		"name": _ObjTypeSubjectNameHandler,
	},
}

//
// ClusterRoleBindingFieldResolvers represents a collection of methods whose products represent the
// response values of the 'ClusterRoleBinding' type.
type ClusterRoleBindingFieldResolvers interface {
	// Subjects implements response to request for 'subjects' field.
	Subjects(p graphql.ResolveParams) (interface{}, error)

	// RoleRef implements response to request for 'roleRef' field.
	RoleRef(p graphql.ResolveParams) (interface{}, error)

	// Name implements response to request for 'name' field.
	Name(p graphql.ResolveParams) (string, error)
}

// ClusterRoleBindingAliases implements all methods on ClusterRoleBindingFieldResolvers interface by using reflection to
// match name of field to a field on the given value. Intent is reduce friction
// of writing new resolvers by removing all the instances where you would simply
// have the resolvers method return a field.
type ClusterRoleBindingAliases struct{}

// Subjects implements response to request for 'subjects' field.
func (_ ClusterRoleBindingAliases) Subjects(p graphql.ResolveParams) (interface{}, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	return val, err
}

// RoleRef implements response to request for 'roleRef' field.
func (_ ClusterRoleBindingAliases) RoleRef(p graphql.ResolveParams) (interface{}, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	return val, err
}

// Name implements response to request for 'name' field.
func (_ ClusterRoleBindingAliases) Name(p graphql.ResolveParams) (string, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := val.(string)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'name'")
	}
	return ret, err
}

/*
ClusterRoleBindingType ClusterRoleBinding grants the permissions defined in a ClusterRole referenced
to a user or a set of users
*/
var ClusterRoleBindingType = graphql.NewType("ClusterRoleBinding", graphql.ObjectKind)

// RegisterClusterRoleBinding registers ClusterRoleBinding object type with given service.
func RegisterClusterRoleBinding(svc *graphql.Service, impl ClusterRoleBindingFieldResolvers) {
	svc.RegisterObject(_ObjectTypeClusterRoleBindingDesc, impl)
}
func _ObjTypeClusterRoleBindingSubjectsHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(ClusterRoleBindingSubjectsFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Subjects(frp)
	}
}

func _ObjTypeClusterRoleBindingRoleRefHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(ClusterRoleBindingRoleRefFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.RoleRef(frp)
	}
}

func _ObjTypeClusterRoleBindingNameHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(ClusterRoleBindingNameFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Name(frp)
	}
}

func _ObjectTypeClusterRoleBindingConfigFn() graphql1.ObjectConfig {
	return graphql1.ObjectConfig{
		Description: "ClusterRoleBinding grants the permissions defined in a ClusterRole referenced\nto a user or a set of users",
		Fields: graphql1.Fields{
			"name": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Name of the ClusterRoleBinding",
				Name:              "name",
				Type:              graphql1.NewNonNull(graphql1.String),
			},
			"roleRef": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "RoleRef references a ClusterRole in the current namespace",
				Name:              "roleRef",
				Type:              graphql.OutputType("RoleRef"),
			},
			"subjects": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Subjects holds references to the objects the ClusterRole applies to",
				Name:              "subjects",
				Type:              graphql1.NewList(graphql1.NewNonNull(graphql.OutputType("Subject"))),
			},
		},
		Interfaces: []*graphql1.Interface{},
		IsTypeOf: func(_ graphql1.IsTypeOfParams) bool {
			// NOTE:
			// Panic by default. Intent is that when Service is invoked, values of
			// these fields are updated with instantiated resolvers. If these
			// defaults are called it is most certainly programmer err.
			// If you're see this comment then: 'Whoops! Sorry, my bad.'
			panic("Unimplemented; see ClusterRoleBindingFieldResolvers.")
		},
		Name: "ClusterRoleBinding",
	}
}

// describe ClusterRoleBinding's configuration; kept private to avoid unintentional tampering of configuration at runtime.
var _ObjectTypeClusterRoleBindingDesc = graphql.ObjectDesc{
	Config: _ObjectTypeClusterRoleBindingConfigFn,
	FieldHandlers: map[string]graphql.FieldHandler{
		"name":     _ObjTypeClusterRoleBindingNameHandler,
		"roleRef":  _ObjTypeClusterRoleBindingRoleRefHandler,
		"subjects": _ObjTypeClusterRoleBindingSubjectsHandler,
	},
}

//
// RoleBindingFieldResolvers represents a collection of methods whose products represent the
// response values of the 'RoleBinding' type.
type RoleBindingFieldResolvers interface {
	// Subjects implements response to request for 'subjects' field.
	Subjects(p graphql.ResolveParams) (interface{}, error)

	// RoleRef implements response to request for 'roleRef' field.
	RoleRef(p graphql.ResolveParams) (interface{}, error)

	// Namespace implements response to request for 'namespace' field.
	Namespace(p graphql.ResolveParams) (string, error)

	// Name implements response to request for 'name' field.
	Name(p graphql.ResolveParams) (string, error)
}

// RoleBindingAliases implements all methods on RoleBindingFieldResolvers interface by using reflection to
// match name of field to a field on the given value. Intent is reduce friction
// of writing new resolvers by removing all the instances where you would simply
// have the resolvers method return a field.
type RoleBindingAliases struct{}

// Subjects implements response to request for 'subjects' field.
func (_ RoleBindingAliases) Subjects(p graphql.ResolveParams) (interface{}, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	return val, err
}

// RoleRef implements response to request for 'roleRef' field.
func (_ RoleBindingAliases) RoleRef(p graphql.ResolveParams) (interface{}, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	return val, err
}

// Namespace implements response to request for 'namespace' field.
func (_ RoleBindingAliases) Namespace(p graphql.ResolveParams) (string, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := val.(string)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'namespace'")
	}
	return ret, err
}

// Name implements response to request for 'name' field.
func (_ RoleBindingAliases) Name(p graphql.ResolveParams) (string, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := val.(string)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'name'")
	}
	return ret, err
}

/*
RoleBindingType RoleBinding grants the permissions defined in a Role referenced to a user or
a set of users
*/
var RoleBindingType = graphql.NewType("RoleBinding", graphql.ObjectKind)

// RegisterRoleBinding registers RoleBinding object type with given service.
func RegisterRoleBinding(svc *graphql.Service, impl RoleBindingFieldResolvers) {
	svc.RegisterObject(_ObjectTypeRoleBindingDesc, impl)
}
func _ObjTypeRoleBindingSubjectsHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(RoleBindingSubjectsFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Subjects(frp)
	}
}

func _ObjTypeRoleBindingRoleRefHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(RoleBindingRoleRefFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.RoleRef(frp)
	}
}

func _ObjTypeRoleBindingNamespaceHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(RoleBindingNamespaceFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Namespace(frp)
	}
}

func _ObjTypeRoleBindingNameHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(RoleBindingNameFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Name(frp)
	}
}

func _ObjectTypeRoleBindingConfigFn() graphql1.ObjectConfig {
	return graphql1.ObjectConfig{
		Description: "RoleBinding grants the permissions defined in a Role referenced to a user or\na set of users",
		Fields: graphql1.Fields{
			"name": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Name of the RoleBinding",
				Name:              "name",
				Type:              graphql1.NewNonNull(graphql1.String),
			},
			"namespace": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Namespace of the RoleBinding",
				Name:              "namespace",
				Type:              graphql1.NewNonNull(graphql1.String),
			},
			"roleRef": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "RoleRef references a Role in the current namespace",
				Name:              "roleRef",
				Type:              graphql.OutputType("RoleRef"),
			},
			"subjects": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Subjects holds references to the objects the Role applies to",
				Name:              "subjects",
				Type:              graphql1.NewList(graphql1.NewNonNull(graphql.OutputType("Subject"))),
			},
		},
		Interfaces: []*graphql1.Interface{},
		IsTypeOf: func(_ graphql1.IsTypeOfParams) bool {
			// NOTE:
			// Panic by default. Intent is that when Service is invoked, values of
			// these fields are updated with instantiated resolvers. If these
			// defaults are called it is most certainly programmer err.
			// If you're see this comment then: 'Whoops! Sorry, my bad.'
			panic("Unimplemented; see RoleBindingFieldResolvers.")
		},
		Name: "RoleBinding",
	}
}

// describe RoleBinding's configuration; kept private to avoid unintentional tampering of configuration at runtime.
var _ObjectTypeRoleBindingDesc = graphql.ObjectDesc{
	Config: _ObjectTypeRoleBindingConfigFn,
	FieldHandlers: map[string]graphql.FieldHandler{
		"name":      _ObjTypeRoleBindingNameHandler,
		"namespace": _ObjTypeRoleBindingNamespaceHandler,
		"roleRef":   _ObjTypeRoleBindingRoleRefHandler,
		"subjects":  _ObjTypeRoleBindingSubjectsHandler,
	},
}
