// Code generated by scripts/gengraphql.go. DO NOT EDIT.

package schema

import (
	errors "errors"
	graphql1 "github.com/graphql-go/graphql"
	graphql "github.com/sensu/sensu-go/graphql"
)

// OffsetPageInfoHasNextPageFieldResolver implement to resolve requests for the OffsetPageInfo's hasNextPage field.
type OffsetPageInfoHasNextPageFieldResolver interface {
	// HasNextPage implements response to request for hasNextPage field.
	HasNextPage(p graphql.ResolveParams) (bool, error)
}

// OffsetPageInfoHasPreviousPageFieldResolver implement to resolve requests for the OffsetPageInfo's hasPreviousPage field.
type OffsetPageInfoHasPreviousPageFieldResolver interface {
	// HasPreviousPage implements response to request for hasPreviousPage field.
	HasPreviousPage(p graphql.ResolveParams) (bool, error)
}

// OffsetPageInfoNextOffsetFieldResolver implement to resolve requests for the OffsetPageInfo's nextOffset field.
type OffsetPageInfoNextOffsetFieldResolver interface {
	// NextOffset implements response to request for nextOffset field.
	NextOffset(p graphql.ResolveParams) (int, error)
}

// OffsetPageInfoPreviousOffsetFieldResolver implement to resolve requests for the OffsetPageInfo's previousOffset field.
type OffsetPageInfoPreviousOffsetFieldResolver interface {
	// PreviousOffset implements response to request for previousOffset field.
	PreviousOffset(p graphql.ResolveParams) (int, error)
}

// OffsetPageInfoTotalCountFieldResolver implement to resolve requests for the OffsetPageInfo's totalCount field.
type OffsetPageInfoTotalCountFieldResolver interface {
	// TotalCount implements response to request for totalCount field.
	TotalCount(p graphql.ResolveParams) (int, error)
}

//
// OffsetPageInfoFieldResolvers represents a collection of methods whose products represent the
// response values of the 'OffsetPageInfo' type.
type OffsetPageInfoFieldResolvers interface {
	OffsetPageInfoHasNextPageFieldResolver
	OffsetPageInfoHasPreviousPageFieldResolver
	OffsetPageInfoNextOffsetFieldResolver
	OffsetPageInfoPreviousOffsetFieldResolver
	OffsetPageInfoTotalCountFieldResolver
}

// OffsetPageInfoAliases implements all methods on OffsetPageInfoFieldResolvers interface by using reflection to
// match name of field to a field on the given value. Intent is reduce friction
// of writing new resolvers by removing all the instances where you would simply
// have the resolvers method return a field.
type OffsetPageInfoAliases struct{}

// HasNextPage implements response to request for 'hasNextPage' field.
func (_ OffsetPageInfoAliases) HasNextPage(p graphql.ResolveParams) (bool, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := val.(bool)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'hasNextPage'")
	}
	return ret, err
}

// HasPreviousPage implements response to request for 'hasPreviousPage' field.
func (_ OffsetPageInfoAliases) HasPreviousPage(p graphql.ResolveParams) (bool, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := val.(bool)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'hasPreviousPage'")
	}
	return ret, err
}

// NextOffset implements response to request for 'nextOffset' field.
func (_ OffsetPageInfoAliases) NextOffset(p graphql.ResolveParams) (int, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := graphql1.Int.ParseValue(val).(int)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'nextOffset'")
	}
	return ret, err
}

// PreviousOffset implements response to request for 'previousOffset' field.
func (_ OffsetPageInfoAliases) PreviousOffset(p graphql.ResolveParams) (int, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := graphql1.Int.ParseValue(val).(int)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'previousOffset'")
	}
	return ret, err
}

// TotalCount implements response to request for 'totalCount' field.
func (_ OffsetPageInfoAliases) TotalCount(p graphql.ResolveParams) (int, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := graphql1.Int.ParseValue(val).(int)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'totalCount'")
	}
	return ret, err
}

// OffsetPageInfoType Information about the current page.
var OffsetPageInfoType = graphql.NewType("OffsetPageInfo", graphql.ObjectKind)

// RegisterOffsetPageInfo registers OffsetPageInfo object type with given service.
func RegisterOffsetPageInfo(svc *graphql.Service, impl OffsetPageInfoFieldResolvers) {
	svc.RegisterObject(_ObjectTypeOffsetPageInfoDesc, impl)
}
func _ObjTypeOffsetPageInfoHasNextPageHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(OffsetPageInfoHasNextPageFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.HasNextPage(frp)
	}
}

func _ObjTypeOffsetPageInfoHasPreviousPageHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(OffsetPageInfoHasPreviousPageFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.HasPreviousPage(frp)
	}
}

func _ObjTypeOffsetPageInfoNextOffsetHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(OffsetPageInfoNextOffsetFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.NextOffset(frp)
	}
}

func _ObjTypeOffsetPageInfoPreviousOffsetHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(OffsetPageInfoPreviousOffsetFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.PreviousOffset(frp)
	}
}

func _ObjTypeOffsetPageInfoTotalCountHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(OffsetPageInfoTotalCountFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.TotalCount(frp)
	}
}

func _ObjectTypeOffsetPageInfoConfigFn() graphql1.ObjectConfig {
	return graphql1.ObjectConfig{
		Description: "Information about the current page.",
		Fields: graphql1.Fields{
			"hasNextPage": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "When paginating forward, are there more items?",
				Name:              "hasNextPage",
				Type:              graphql1.NewNonNull(graphql1.Boolean),
			},
			"hasPreviousPage": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "When paginating backward, are there more items?",
				Name:              "hasPreviousPage",
				Type:              graphql1.NewNonNull(graphql1.Boolean),
			},
			"nextOffset": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Next offset to use when paginating forward; null if there are no more items.",
				Name:              "nextOffset",
				Type:              graphql1.Int,
			},
			"previousOffset": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Previous offset to use when paginating backward; null if at terminal.",
				Name:              "previousOffset",
				Type:              graphql1.Int,
			},
			"totalCount": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Total count of records in relationship.",
				Name:              "totalCount",
				Type:              graphql1.NewNonNull(graphql1.Int),
			},
		},
		Interfaces: []*graphql1.Interface{},
		IsTypeOf: func(_ graphql1.IsTypeOfParams) bool {
			// NOTE:
			// Panic by default. Intent is that when Service is invoked, values of
			// these fields are updated with instantiated resolvers. If these
			// defaults are called it is most certainly programmer err.
			// If you're see this comment then: 'Whoops! Sorry, my bad.'
			panic("Unimplemented; see OffsetPageInfoFieldResolvers.")
		},
		Name: "OffsetPageInfo",
	}
}

// describe OffsetPageInfo's configuration; kept private to avoid unintentional tampering of configuration at runtime.
var _ObjectTypeOffsetPageInfoDesc = graphql.ObjectDesc{
	Config: _ObjectTypeOffsetPageInfoConfigFn,
	FieldHandlers: map[string]graphql.FieldHandler{
		"hasNextPage":     _ObjTypeOffsetPageInfoHasNextPageHandler,
		"hasPreviousPage": _ObjTypeOffsetPageInfoHasPreviousPageHandler,
		"nextOffset":      _ObjTypeOffsetPageInfoNextOffsetHandler,
		"previousOffset":  _ObjTypeOffsetPageInfoPreviousOffsetHandler,
		"totalCount":      _ObjTypeOffsetPageInfoTotalCountHandler,
	},
}
