package wrike

import (
	"github.com/google/go-querystring/query"
	parameters "github.com/wingman131/wrike.go/parameters"
	types "github.com/wingman131/wrike.go/types"
)

// CreateWorkflow creates a new workflow with given parameters.
// For details refer to https://developers.wrike.com/documentation/api/methods/create-workflows
func (api *API) CreateWorkflow(params parameters.CreateWorkflow) (*types.Workflows, error) {
	path := "/workflows"

	body, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	data, err := api.post(path, &body)
	if err != nil {
		return nil, err
	}

	return types.NewWorkflowsFromJSON(data)
}
