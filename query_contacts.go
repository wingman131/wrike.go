package wrike

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"
	parameters "github.com/wingman131/wrike.go/parameters"
	types "github.com/wingman131/wrike.go/types"
)

// QueryContacts fetches a list of contacts.
// For details refer to https://developers.wrike.com/documentation/api/methods/query-contacts
func (api *API) QueryContacts(params *parameters.QueryContacts) (*types.Contacts, error) {
	return api.queryContacts("/contacts", params)
}

// QueryContactsByIDs fetches contacts by IDs.
// For details refer to https://developers.wrike.com/documentation/api/methods/query-contacts
func (api *API) QueryContactsByIDs(ids parameters.ContactIDSet, params *parameters.QueryContacts) (*types.Contacts, error) {
	baseURL := fmt.Sprintf("/contacts/%s", strings.Join(ids.ToSlice(), ","))

	return api.queryContacts(baseURL, params)
}

func (api *API) queryContacts(path string, params *parameters.QueryContacts) (*types.Contacts, error) {
	qparams, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	data, err := api.get(path, &qparams)
	if err != nil {
		return nil, err
	}

	return types.NewContactsFromJSON(data)
}
