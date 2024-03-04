package main_test

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/josenymad/one-two-one-two/controller"
	"github.com/stretchr/testify/assert"
)

func TestGetRequestData(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	c.Request = &http.Request{
		URL: &url.URL{
			Path:     "/test",
			RawQuery: "status=300",
		},
	}

	jasonStr := []byte(`{"test": "testing"}`)
	c.Request.Body = io.NopCloser(bytes.NewBuffer(jasonStr))

	if c.Request.Header == nil {
		c.Request.Header = make(http.Header)
	}
	c.Request.Header.Set("Content-Type", "application/json")

	response := controller.GetRequestData(c)

	assert.NotNil(t, response)
	assert.NotNil(t, response.Data)
	assert.Equal(t, "300", response.Status)

	responseDataJSON, err := json.Marshal(response.Data)
	if err != nil {
		t.Fatal("failed to marshal response data to JSON")
		log.Println(err)
	}

	expectedData := map[string]interface{}{
		"status": "300",
		"test":   "testing",
	}
	expectedDataJSON, err := json.Marshal(expectedData)
	if err != nil {
		t.Fatal("failed to marshal expected data to JSON")
		log.Println(err)
	}

	assert.Equal(t, expectedDataJSON, responseDataJSON)
}
