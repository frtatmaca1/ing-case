package test

import (
	"net/http"

	"github.com/frtatmaca/case/model/request"
	"github.com/stretchr/testify/mock"
)

type MockPersonClient struct {
	mock.Mock
	baseUrl string
	client  *http.Client
}

func NewMockPersonApiClient(baseUrl string, client *http.Client) *MockPersonClient {
	return &MockPersonClient{baseUrl: baseUrl, client: client}
}

func (m *MockPersonClient) CreatePerson(personGreeting request.PersonGreeting) (*http.Response, error) {
	args := m.Called(personGreeting)
	return args.Get(0).(*http.Response), args.Error(1)
}

func (m *MockPersonClient) GetPerson() (*http.Response, error) {
	args := m.Called()
	return args.Get(0).(*http.Response), args.Error(1)
}

func (m *MockPersonClient) GetBaseUrl() string {
	args := m.Called()
	return args.Get(0).(string)
}
