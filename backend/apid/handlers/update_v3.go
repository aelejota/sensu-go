package handlers

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
	corev3 "github.com/sensu/sensu-go/api/core/v3"
	"github.com/sensu/sensu-go/backend/apid/actions"
	"github.com/sensu/sensu-go/backend/authentication/jwt"
	"github.com/sensu/sensu-go/backend/store"
	storev2 "github.com/sensu/sensu-go/backend/store/v2"
	"github.com/sensu/sensu-go/backend/store/v2/wrap"
)

// CreateOrUpdateResource creates or updates the resource given in the request
// body, regardless of whether it already exists or not
func (h Handlers) CreateOrUpdateV3Resource(r *http.Request) (interface{}, error) {
	payload := reflect.New(reflect.TypeOf(h.V3Resource).Elem())
	if err := json.NewDecoder(r.Body).Decode(payload.Interface()); err != nil {
		return nil, actions.NewError(actions.InvalidArgument, err)
	}

	if err := CheckV3Meta(payload.Interface(), mux.Vars(r), "id"); err != nil {
		return nil, actions.NewError(actions.InvalidArgument, err)
	}

	resource, ok := payload.Interface().(corev3.Resource)
	if !ok {
		return nil, actions.NewErrorf(actions.InvalidArgument)
	}

	req := storev2.NewResourceRequestFromResource(r.Context(), resource)
	meta := resource.GetMetadata()

	if claims := jwt.GetClaimsFromContext(r.Context()); claims != nil {
		meta.CreatedBy = claims.StandardClaims.Subject
	}

	wrapper, err := wrap.Resource(resource)
	if err != nil {
		return nil, actions.NewError(actions.InvalidArgument, err)
	}

	if err := h.StoreV2.CreateOrUpdate(req, wrapper); err != nil {
		switch err := err.(type) {
		case *store.ErrNotValid:
			return nil, actions.NewError(actions.InvalidArgument, err)
		default:
			return nil, actions.NewError(actions.InternalErr, err)
		}
	}

	return nil, nil
}
