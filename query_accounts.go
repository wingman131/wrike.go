package wrike

import (
	"github.com/google/go-querystring/query"
	"github.com/wingman131/wrike.go/parameters"
	types "github.com/wingman131/wrike.go/types"
)

// QueryAccounts fetches a Accounts by id.
// For details refer to https://developers.wrike.com/documentation/api/methods/query-accounts
func (api *API) QueryAccounts(params *parameters.QueryAccounts) (*types.Accounts, error) {
	path := "/account"

	qparams, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	data, err := api.get(path, &qparams)
	if err != nil {
		return nil, err
	}

	return types.NewAccountsFromJSON(data)
}
