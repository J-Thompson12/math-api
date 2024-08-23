package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newContextWithBody(body []byte) (echo.Context, *httptest.ResponseRecorder) {
	req := httptest.NewRequest("GET", "/", bytes.NewBuffer(body))
	rec := httptest.NewRecorder()
	e := echo.New()
	e.Validator = NewValidator()
	context := e.NewContext(req, rec)
	context.Request().Header.Set("Content-Type", "application/json")
	return context, rec
}

func assertHTTPError(t *testing.T, err error, code int) {
	httpError, ok := err.(*echo.HTTPError)
	require.True(t, ok)
	assert.Equal(t, code, httpError.Code)
}

func TestRequest(t *testing.T) {
	t.Run("valid request", func(t *testing.T) {
		// Create a test request
		req := new(Request)
		req.List = []int{1, 2, 3}
		req.Quantifier = 2

		// Marshal the request to JSON
		jsonReq, err := json.Marshal(req)
		require.NoError(t, err)

		c, _ := newContextWithBody(jsonReq)
		result, err := newRequest(c)
		require.NoError(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("bad validator request", func(t *testing.T) {
		// Create a test request
		req := new(Request)
		req.List = []int{}
		req.Quantifier = 2

		// Marshal the request to JSON
		jsonReq, err := json.Marshal(req)
		require.NoError(t, err)

		c, _ := newContextWithBody(jsonReq)
		_, err = newRequest(c)
		require.ErrorContains(t, err, "validation")
	})

	t.Run("bad json request", func(t *testing.T) {
		// Create a test request
		req := new(Request)
		req.List = []int{}
		req.Quantifier = 2

		// Marshal the request to JSON
		jsonReq, err := json.Marshal(req)
		require.NoError(t, err)

		c, _ := newContextWithBody(jsonReq)
		c.Request().Header.Set("Content-Type", "text")
		_, err = newRequest(c)
		require.ErrorContains(t, err, "invalid request")
	})

}

func TestHandleMin(t *testing.T) {
	t.Run("valid request", func(t *testing.T) {
		// Create a test request
		req := new(Request)
		req.List = []int{1, 2, 3}
		req.Quantifier = 2

		// Marshal the request to JSON
		jsonReq, err := json.Marshal(req)
		require.NoError(t, err)

		c, rec := newContextWithBody(jsonReq)

		// Call the handleMin function
		err = handleMin(c)
		require.NoError(t, err)

		// Check the response status code
		assert.Equal(t, http.StatusOK, rec.Result().StatusCode)

		// Unmarshal the response JSON
		var resp []int
		err = json.Unmarshal(rec.Body.Bytes(), &resp)
		require.NoError(t, err)

		// Check the response values
		assert.Equal(t, []int{1, 2}, resp)
	})

	t.Run("invalid request", func(t *testing.T) {
		// Create a test request
		req := new(Request)
		req.List = []int{1, 2, 3}
		req.Quantifier = 4

		// Marshal the request to JSON
		jsonReq, err := json.Marshal(req)
		require.NoError(t, err)

		c, _ := newContextWithBody(jsonReq)

		// Call the handleMin function
		err = handleMin(c)
		assertHTTPError(t, err, http.StatusBadRequest)
	})
}

func TestHandleMax(t *testing.T) {
	t.Run("valid request", func(t *testing.T) {
		// Create a test request
		req := new(Request)
		req.List = []int{1, 2, 3}
		req.Quantifier = 2

		// Marshal the request to JSON
		jsonReq, err := json.Marshal(req)
		require.NoError(t, err)

		c, rec := newContextWithBody(jsonReq)

		// Call the handleMax function
		err = handleMax(c)
		require.NoError(t, err)

		// Check the response status code
		assert.Equal(t, http.StatusOK, rec.Result().StatusCode)

		// Unmarshal the response JSON
		var resp []int
		err = json.Unmarshal(rec.Body.Bytes(), &resp)
		require.NoError(t, err)

		// Check the response values
		assert.Equal(t, []int{3, 2}, resp)
	})

	t.Run("invalid request", func(t *testing.T) {
		// Create a test request
		req := new(Request)
		req.List = []int{1, 2, 3}
		req.Quantifier = 0

		// Marshal the request to JSON
		jsonReq, err := json.Marshal(req)
		require.NoError(t, err)

		c, _ := newContextWithBody(jsonReq)

		// Call the handleMax function
		err = handleMax(c)
		assertHTTPError(t, err, http.StatusBadRequest)
	})
}

func TestHandleAverage(t *testing.T) {
	t.Run("valid request", func(t *testing.T) {
		// Create a test request
		req := new(Request)
		req.List = []int{1, 2, 3}

		// Marshal the request to JSON
		jsonReq, err := json.Marshal(req)
		require.NoError(t, err)

		c, rec := newContextWithBody(jsonReq)

		// Call the handleMax function
		err = handleAverage(c)
		require.NoError(t, err)

		// Check the response status code
		assert.Equal(t, http.StatusOK, rec.Result().StatusCode)

		// Unmarshal the response JSON
		var resp int
		err = json.Unmarshal(rec.Body.Bytes(), &resp)
		require.NoError(t, err)

		// Check the response values
		assert.Equal(t, 2, resp)
	})

	t.Run("invalid request", func(t *testing.T) {
		// Create a test request
		req := new(Request)
		req.List = []int{}

		// Marshal the request to JSON
		jsonReq, err := json.Marshal(req)
		require.NoError(t, err)

		c, _ := newContextWithBody(jsonReq)

		// Call the handleMax function
		err = handleAverage(c)
		assertHTTPError(t, err, http.StatusBadRequest)
	})
}

func TestHandleMedian(t *testing.T) {
	t.Run("valid request", func(t *testing.T) {
		// Create a test request
		req := new(Request)
		req.List = []int{1, 2, 3}

		// Marshal the request to JSON
		jsonReq, err := json.Marshal(req)
		require.NoError(t, err)

		c, rec := newContextWithBody(jsonReq)

		// Call the handleMax function
		err = handleMedian(c)
		require.NoError(t, err)

		// Check the response status code
		assert.Equal(t, http.StatusOK, rec.Result().StatusCode)

		// Unmarshal the response JSON
		var resp int
		err = json.Unmarshal(rec.Body.Bytes(), &resp)
		require.NoError(t, err)

		// Check the response values
		assert.Equal(t, 2, resp)
	})

	t.Run("invalid request", func(t *testing.T) {
		// Create a test request
		req := new(Request)
		req.List = []int{}

		// Marshal the request to JSON
		jsonReq, err := json.Marshal(req)
		require.NoError(t, err)

		c, _ := newContextWithBody(jsonReq)

		// Call the handleMax function
		err = handleMedian(c)
		assertHTTPError(t, err, http.StatusBadRequest)
	})
}

func TestHandlePercentile(t *testing.T) {
	t.Run("valid request", func(t *testing.T) {
		// Create a test request
		req := new(Request)
		req.List = []int{1, 2, 3}
		req.Quantifier = 2

		// Marshal the request to JSON
		jsonReq, err := json.Marshal(req)
		require.NoError(t, err)

		c, rec := newContextWithBody(jsonReq)

		// Call the handleMax function
		err = handlePercentile(c)
		require.NoError(t, err)

		// Check the response status code
		assert.Equal(t, http.StatusOK, rec.Result().StatusCode)

		// Unmarshal the response JSON
		var resp int
		err = json.Unmarshal(rec.Body.Bytes(), &resp)
		require.NoError(t, err)

		// Check the response values
		assert.Equal(t, 1, resp)
	})

	t.Run("invalid request", func(t *testing.T) {
		// Create a test request
		req := new(Request)
		req.List = []int{1, 2, 3}
		req.Quantifier = 0

		// Marshal the request to JSON
		jsonReq, err := json.Marshal(req)
		require.NoError(t, err)

		c, _ := newContextWithBody(jsonReq)

		// Call the handleMax function
		err = handlePercentile(c)
		assertHTTPError(t, err, http.StatusBadRequest)
	})
}
