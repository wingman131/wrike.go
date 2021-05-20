package wrike

import (
	"fmt"

	"github.com/google/go-querystring/query"
	parameters "github.com/wingman131/wrike.go/parameters"
	types "github.com/wingman131/wrike.go/types"
)

// CreateTask creates a new task with given parameters.
// For details refer to https://developers.wrike.com/documentation/api/methods/create-task
func (api *API) CreateTask(folderID types.FolderID, params parameters.CreateTask) (*types.Tasks, error) {
	path := fmt.Sprintf("/folders/%s/tasks", folderID)

	body, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	data, err := api.post(path, &body)
	if err != nil {
		return nil, err
	}

	return types.NewTasksFromJSON(data)
}
