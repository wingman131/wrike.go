package wrike

import (
	"fmt"

	"github.com/google/go-querystring/query"
	"github.com/wingman131/wrike.go/parameters"
	"github.com/wingman131/wrike.go/types"
)

// CreateFolder creates a new folder with given parameters.
// For details refer to https://developers.wrike.com/documentation/api/methods/create-folder
func (api *API) CreateFolder(id types.FolderID, params parameters.CreateFolder) (*types.Folders, error) {
	path := fmt.Sprintf("/folders/%s/folders", id)

	body, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	data, err := api.post(path, &body)
	if err != nil {
		return nil, err
	}

	return types.NewFoldersFromJSON(data)
}
