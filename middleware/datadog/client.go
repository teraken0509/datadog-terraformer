package middleware

import (
	"errors"

	datadog "github.com/zorkian/go-datadog-api"
)

// Credential ...
type Credential struct {
	APIKey string
	AppKey string
	Client *datadog.Client
}

// NewCredential ...
func NewCredential(apiKey, appKey string) (*Credential, error) {
	client := datadog.NewClient(apiKey, appKey)

	return &Credential{
		APIKey: apiKey,
		AppKey: appKey,
		Client: client,
	}, nil
}

// validate ...
func (c *Credential) validate() error {
	ok, err := c.Client.Validate()
	if err != nil {
		return err
	}

	if !ok {
		return errors.New("Datadog Client is invalid")
	}

	return nil
}
