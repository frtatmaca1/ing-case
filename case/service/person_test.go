package service

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/frtatmaca/case/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_Person_Success(t *testing.T) {
	personApiClient := test.NewMockPersonApiClient("randomUrl", http.DefaultClient)
	personService := NewPersonService(personApiClient)

	personApiSuccessResponse := http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(bytes.NewReader([]byte("body"))),
	}

	personApiClient.On("GetBaseUrl").Return("randomUrl")
	personApiClient.On("GetPerson", mock.Anything, mock.Anything, mock.Anything).Return(&personApiSuccessResponse, nil)
	personApiClient.On("CreatePerson", mock.Anything, mock.Anything, mock.Anything).Return(&personApiSuccessResponse, nil)

	_, err := personService.Person()
	assert.Nil(t, err)
}
