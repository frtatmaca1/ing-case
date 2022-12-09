package proxy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/frtatmaca/case/model/request"
)

const (
	Endpoint = "/person"
)

type Client interface {
	CreatePerson(personGreeting request.PersonGreeting) (*http.Response, error)
	GetPerson() (*http.Response, error)
	GetBaseUrl() string
}

type client struct {
	baseUrl string
	client  *http.Client
}

func NewClient(baseUrl string, httpClient *http.Client) Client {
	return &client{baseUrl: baseUrl, client: httpClient}
}

func (c *client) CreatePerson(personGreeting request.PersonGreeting) (*http.Response, error) {
	url := c.baseUrl + fmt.Sprintf(Endpoint)

	reqBody, err := json.Marshal(personGreeting)
	httpReq, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(reqBody)))
	if err != nil {
		return nil, err
	}

	res, err := c.client.Do(httpReq)
	return res, err
}

func (c *client) GetPerson() (*http.Response, error) {
	url := c.baseUrl + fmt.Sprintf(Endpoint)

	httpReq, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	return c.client.Do(httpReq)
}

func (c *client) GetBaseUrl() string {
	return c.baseUrl
}
