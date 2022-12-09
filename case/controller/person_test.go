package controller

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/frtatmaca/case/middleware"
	"github.com/frtatmaca/case/pkg/logging"
	"github.com/frtatmaca/case/service"
	"github.com/frtatmaca/case/test"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_Person_Success(t *testing.T) {
	personApiClient := new(test.MockPersonClient)
	personApiSuccessResponse := http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(bytes.NewReader([]byte("body"))),
	}

	personApiClient.On("GetBaseUrl").Return("randomUrl")
	personApiClient.On("GetPerson", mock.Anything, mock.Anything, mock.Anything).Return(&personApiSuccessResponse, nil)
	personApiClient.On("CreatePerson", mock.Anything, mock.Anything, mock.Anything).Return(&personApiSuccessResponse, nil)

	personService := service.NewPersonService(personApiClient)

	personController := NewPersonController(personService)

	w := httptest.NewRecorder()
	ctx, r := gin.CreateTestContext(w)
	r.Use(middleware.TracingHandler(), middleware.ErrorHandler(logging.NewLoggerWithLevel("/tmp/case.log", "fatal")))

	r.GET("/persons", personController.Person)
	ctx.Request = httptest.NewRequest(http.MethodGet, "/persons", bytes.NewBuffer([]byte("")))

	r.ServeHTTP(w, ctx.Request)

	var body map[string]string
	json.Unmarshal(w.Body.Bytes(), &body)

	assert.Equal(t, http.StatusOK, w.Code)
}
