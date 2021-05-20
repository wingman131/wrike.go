package wrike

import (
	"fmt"

	"github.com/google/go-querystring/query"
	parameters "github.com/wingman131/wrike.go/parameters"
	types "github.com/wingman131/wrike.go/types"
)

// ModifyCustomField modifies a custom field with given parameters.
// For details refer to https://developers.wrike.com/documentation/api/methods/modify-custom-field
func (api *API) ModifyCustomField(id string, params *parameters.ModifyCustomField) (*types.CustomFields, error) {
	path := fmt.Sprintf("/customfields/%s", id)

	body, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	data, err := api.put(path, &body)
	if err != nil {
		return nil, err
	}

	return types.NewCustomFieldsFromJSON(data)
}
