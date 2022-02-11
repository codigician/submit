package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/codigician/submit/handler"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSubmit_ShouldFailWhenContentIsNull(t *testing.T) {
	srv := prepareHandlerAndStartServer(nil)
	defer srv.Close()

	rsp, err := http.Post(fmt.Sprintf("%s/submit/1", srv.URL), "application/json", bytes.NewBuffer([]byte("abc")))

	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, rsp.StatusCode)
}

func TestSubmit_InvalidRequestBody(t *testing.T) {
	srv := prepareHandlerAndStartServer(nil)
	defer srv.Close()
	testCases := []struct {
		scenario     string
		givenReqBody handler.SubmissionReq
	}{
		{
			scenario:     "empty language",
			givenReqBody: handler.SubmissionReq{Content: "asdfd"},
		},
		{
			scenario:     "empty content",
			givenReqBody: handler.SubmissionReq{Lang: "golang"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			reqBytes, _ := json.Marshal(tc.givenReqBody)
			rsp, err := http.Post(fmt.Sprintf("%s/submit/1", srv.URL), "application/json", bytes.NewBuffer(reqBytes))

			assert.Nil(t, err)
			assert.Equal(t, http.StatusBadRequest, rsp.StatusCode)
		})

	}
}

func TestSubmit_ServiceReturnsError_ExpectsInternalServerError(t *testing.T) {
	srv := prepareHandlerAndStartServer(&FakeSubmitService{})
	defer srv.Close()

	reqBody := handler.SubmissionReq{
		Content: "asdf",
		Lang:    "golang",
	}

	reqBytes, _ := json.Marshal(reqBody)

	rsp, err := http.Post(fmt.Sprintf("%s/submit/1", srv.URL), "application/json", bytes.NewBuffer(reqBytes))

	assert.Nil(t, err)
	assert.Equal(t, http.StatusInternalServerError, rsp.StatusCode)
}

func TestSubmit(t *testing.T) {
	srv := prepareHandlerAndStartServer(&FakeSubmitService{})
	defer srv.Close()

	reqBody := handler.SubmissionReq{
		Content: "asdf",
		Lang:    "golang",
	}
	reqBytes, _ := json.Marshal(reqBody)

	rsp, _ := http.Post(fmt.Sprintf("%s/submit/2", srv.URL), "application/json", bytes.NewBuffer(reqBytes))

	assert.Equal(t, http.StatusOK, rsp.StatusCode)
}

func prepareHandlerAndStartServer(submitService handler.SubmitService) *httptest.Server {
	submitHandler := handler.NewSubmit(submitService)
	e := echo.New()

	submitHandler.RegisterRoutes(e)
	return httptest.NewServer(e.Server.Handler)
}
