package wrike

import (
	"fmt"

	"github.com/google/go-querystring/query"
	"github.com/wingman131/wrike.go/parameters"
	types "github.com/wingman131/wrike.go/types"
)

// ModifyContact modifies a contact with given parameters.
// For details refer to https://developers.wrike.com/documentation/api/methods/modify-contact
func (api *API) ModifyContact(id string, params *parameters.ModifyContact) (*types.Contacts, error) {
	path := fmt.Sprintf("/contacts/%s", id)

	body, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	data, err := api.put(path, &body)
	if err != nil {
		return nil, err
	}

	return types.NewContactsFromJSON(data)
}
