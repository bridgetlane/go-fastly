package fastly

import (
	"fmt"
)

// BigQuery represents a BigQuery logging response from the Fastly API.
type BigQuery struct {
	ServiceID         string `mapstructure:"service_id"`
	Name              string `mapstructure:"name"`
	Format            string `mapstructure:"format"`
	User              string `mapstructure:"user"`
	ProjectID         string `mapstructure:"project_id"`
	Dataset           string `mapstructure:"dataset"`
	Table             string `mapstructure:"table"`
	SecretKey         string `mapstructure:"secret_key"`
	CreatedAt         string `mapstructure:"created_at"`
	UpdatedAt         string `mapstructure:"updated_at"`
	DeletedAt         string `mapstructure:"deleted_at"`
	ResponseCondition string `mapstructure:"response_condition"`
}

// GetBigQueryInput is used as input to the GetBQs function.
type GetBigQueryInput struct {
	// Service is the ID of the service. Version is the specific configuration
	// version. Both fields are required.
	Service string
	Version int
}

// GetBigQuery lists all BigQuerys associated with a service version.
func (c *Client) GetBigQuery(i *GetBigQueryInput) ([]*BigQuery, error) {
	if i.Service == "" {
		return nil, ErrMissingService
	}

	if i.Version == 0 {
		return nil, ErrMissingVersion
	}

	path := fmt.Sprintf("/service/%s/version/%d/logging/bigquery", i.Service, i.Version)
	resp, err := c.Get(path, nil)
	if err != nil {
		return nil, err
	}

	var bs []*BigQuery
	if err := decodeJSON(&bs, resp.Body); err != nil {
		return nil, err
	}
	return bs, nil
}

// CreateBigQueryInput is used as input to the CreateBigQuery function.
// All fields are required.
type CreateBigQueryInput struct {
	// Service is the ID of the service.
	Service string

	//Version is the specific configuration version.
	Version int

	// Name is the name if your bigquery logging endpoint.
	Name string

	// Project ID your GCP project ID.
	ProjectID string

	// Dataset is your BigQuery dataset.
	Dataset string

	// Table is your BigQuery table.
	Table string

	// User is the user with access to write to your BigQuery dataset.
	User string

	// Secret key is the user's secret key.
	SecretKey string
}

// CreateBigQuery creates a new Fastly BigQuery logging endpoint.
func (c *Client) CreateBigQuery(i *CreateBigQueryInput) (*BigQuery, error) {
	if i.Service == "" {
		return nil, ErrMissingService
	}

	if i.Version == 0 {
		return nil, ErrMissingVersion
	}

	if i.Name == "" {
		return nil, ErrMissingName
	}

	if i.ProjectID == "" {
		return nil, ErrMissingProjectID
	}

	if i.Dataset == "" {
		return nil, ErrMissingDataset
	}

	if i.Table == "" {
		return nil, ErrMissingTable
	}

	if i.User == "" {
		return nil, ErrMissingUser
	}

	if i.SecretKey == "" {
		return nil, ErrMissingSecretKey
	}

	params := make(map[string]string)
	params["name"] = i.Name
	params["project_id"] = i.ProjectID
	params["dataset"] = i.Dataset
	params["table"] = i.Table
	params["user"] = i.User
	params["secret_key"] = i.SecretKey

	path := fmt.Sprintf("/service/%s/version/%d/logging/gcs", i.Service, i.Version)
	resp, err := c.PostForm(path, i, &RequestOptions{
		Params: params,
	})
	if err != nil {
		return nil, err
	}

	var b *BigQuery
	if err := decodeJSON(&b, resp.Body); err != nil {
		return nil, err
	}
	return b, nil
}

// UpdateBigQueryInput is used as input to the UpdateBigQuery function.
// All fields are required.
type UpdateBigQueryInput struct {
	// Service is the ID of the service.
	Service string

	//Version is the specific configuration version.
	Version int

	// Name is the old name if your bigquery logging endpoint.
	// Used to identify the correct BigQuery logging endpoint if there
	// is a name change.
	Name string

	// NewName is the new name of your BigQuery logging endpoint.
	NewName string
}

// UpdateBigQuery updates a BigQuery logging endpoint.
func (c *Client) UpdateBigQuery(i *UpdateBigQueryInput) (*BigQuery, error) {
	if i.Service == "" {
		return nil, ErrMissingService
	}

	if i.Version == 0 {
		return nil, ErrMissingVersion
	}

	if i.Name == "" {
		return nil, ErrMissingName
	}

	if i.NewName == "" {
		return nil, ErrMissingNewName
	}

	params := make(map[string]string)
	params["name"] = i.NewName

	path := fmt.Sprintf("/service/%s/version/%d/logging/bigquery/%s", i.Service, i.Version, i.Name)
	resp, err := c.PutForm(path, i, &RequestOptions{
		Params: params,
	})
	if err != nil {
		return nil, err
	}

	var b *BigQuery
	if err := decodeJSON(&b, resp.Body); err != nil {
		return nil, err
	}
	return b, nil
}

// DeleteBigQueryInput is the input parameter to DeleteBigQuery.
// All fields are required.
type DeleteBigQueryInput struct {
	// Service is the ID of the service.
	Service string

	// Version is the specific configuration.
	Version int

	// Name is the name of the BigQuery logging endpoint to delete.
	Name string
}

// DeleteBigQuery deletes the given BigQuery logging endpoint.
func (c *Client) DeleteBigQuery(i *DeleteBigQueryInput) error {
	if i.Service == "" {
		return ErrMissingService
	}

	if i.Version == 0 {
		return ErrMissingVersion
	}

	if i.Name == "" {
		return ErrMissingName
	}

	path := fmt.Sprintf("/service/%s/version/%d/logging/bigquery/%s", i.Service, i.Version, i.Name)
	resp, err := c.Delete(path, nil)
	if err != nil {
		return err
	}

	var r *statusResp
	if err := decodeJSON(&r, resp.Body); err != nil {
		return err
	}
	if !r.Ok() {
		return fmt.Errorf("Not Ok")
	}
	return nil
}
