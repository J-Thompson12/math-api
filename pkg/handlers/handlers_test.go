package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/J-Thompson12/math-api/pkg/equations"
	"github.com/J-Thompson12/math-api/pkg/middleware"
	"github.com/J-Thompson12/math-api/pkg/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newContext(m equations.Equation) (echo.Context, *httptest.ResponseRecorder) {
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()
	e := echo.New()
	e.Validator = validator.NewValidator()
	context := e.NewContext(req, rec)
	context.Request().Header.Set("Content-Type", "application/json")
	middleware.SetEquation(context, m)
	return context, rec
}

func assertHTTPError(t *testing.T, err error, code int) {
	httpError, ok := err.(*echo.HTTPError)
	require.True(t, ok)
	assert.Equal(t, code, httpError.Code)
}

type Request struct {
	List       []int
	Quantifier int
}

func TestHandleMin(t *testing.T) {
	t.Run("valid request", func(t *testing.T) {
		m := equations.NewMockEquation()
		c, rec := newContext(m)

		err := handleMin(c)
		require.NoError(t, err)

		// Check the response status code
		assert.Equal(t, http.StatusOK, rec.Result().StatusCode)

		// Unmarshal the response JSON
		var resp []int
		err = json.Unmarshal(rec.Body.Bytes(), &resp)
		require.NoError(t, err)
	})

	t.Run("invalid request", func(t *testing.T) {
		m := equations.NewMockEquation()
		m.Error = equations.ErrBadRequest("bad request")
		c, _ := newContext(m)

		err := handleMin(c)
		assertHTTPError(t, err, http.StatusBadRequest)
	})

	t.Run("server errpr", func(t *testing.T) {
		m := equations.NewMockEquation()
		m.Error = fmt.Errorf("server error")
		c, _ := newContext(m)

		err := handleMin(c)
		assertHTTPError(t, err, http.StatusInternalServerError)
	})
}

func TestHandleMax(t *testing.T) {
	t.Run("valid request", func(t *testing.T) {
		m := equations.NewMockEquation()
		c, rec := newContext(m)

		err := handleMax(c)
		require.NoError(t, err)

		// Check the response status code
		assert.Equal(t, http.StatusOK, rec.Result().StatusCode)

		// Unmarshal the response JSON
		var resp []int
		err = json.Unmarshal(rec.Body.Bytes(), &resp)
		require.NoError(t, err)
	})

	t.Run("invalid request", func(t *testing.T) {
		m := equations.NewMockEquation()
		m.Error = equations.ErrBadRequest("bad request")
		c, _ := newContext(m)

		err := handleMax(c)
		assertHTTPError(t, err, http.StatusBadRequest)
	})

	t.Run("server errpr", func(t *testing.T) {
		m := equations.NewMockEquation()
		m.Error = fmt.Errorf("server error")
		c, _ := newContext(m)

		err := handleMax(c)
		assertHTTPError(t, err, http.StatusInternalServerError)
	})
}

func TestHandleAverage(t *testing.T) {
	t.Run("valid request", func(t *testing.T) {
		m := equations.NewMockEquation()
		c, rec := newContext(m)

		err := handleAverage(c)
		require.NoError(t, err)

		// Check the response status code
		assert.Equal(t, http.StatusOK, rec.Result().StatusCode)

		// Unmarshal the response JSON
		var resp float64
		err = json.Unmarshal(rec.Body.Bytes(), &resp)
		require.NoError(t, err)
	})

	t.Run("invalid request", func(t *testing.T) {
		m := equations.NewMockEquation()
		m.Error = equations.ErrBadRequest("bad request")
		c, _ := newContext(m)

		err := handleAverage(c)
		assertHTTPError(t, err, http.StatusBadRequest)
	})

	t.Run("server errpr", func(t *testing.T) {
		m := equations.NewMockEquation()
		m.Error = fmt.Errorf("server error")
		c, _ := newContext(m)

		err := handleAverage(c)
		assertHTTPError(t, err, http.StatusInternalServerError)
	})
}

func TestHandleMedian(t *testing.T) {
	t.Run("valid request", func(t *testing.T) {
		m := equations.NewMockEquation()
		c, rec := newContext(m)

		err := handleMedian(c)
		require.NoError(t, err)

		// Check the response status code
		assert.Equal(t, http.StatusOK, rec.Result().StatusCode)

		// Unmarshal the response JSON
		var resp int
		err = json.Unmarshal(rec.Body.Bytes(), &resp)
		require.NoError(t, err)
	})

	t.Run("invalid request", func(t *testing.T) {
		m := equations.NewMockEquation()
		m.Error = equations.ErrBadRequest("bad request")
		c, _ := newContext(m)

		err := handleMedian(c)
		assertHTTPError(t, err, http.StatusBadRequest)
	})

	t.Run("server errpr", func(t *testing.T) {
		m := equations.NewMockEquation()
		m.Error = fmt.Errorf("server error")
		c, _ := newContext(m)

		err := handleMedian(c)
		assertHTTPError(t, err, http.StatusInternalServerError)
	})
}

func TestHandlePercentile(t *testing.T) {
	t.Run("valid request", func(t *testing.T) {
		m := equations.NewMockEquation()
		c, rec := newContext(m)

		err := handlePercentile(c)
		require.NoError(t, err)

		// Check the response status code
		assert.Equal(t, http.StatusOK, rec.Result().StatusCode)

		// Unmarshal the response JSON
		var resp int
		err = json.Unmarshal(rec.Body.Bytes(), &resp)
		require.NoError(t, err)
	})

	t.Run("invalid request", func(t *testing.T) {
		m := equations.NewMockEquation()
		m.Error = equations.ErrBadRequest("bad request")
		c, _ := newContext(m)

		err := handlePercentile(c)
		assertHTTPError(t, err, http.StatusBadRequest)
	})

	t.Run("server errpr", func(t *testing.T) {
		m := equations.NewMockEquation()
		m.Error = fmt.Errorf("server error")
		c, _ := newContext(m)

		err := handlePercentile(c)
		assertHTTPError(t, err, http.StatusInternalServerError)
	})
}
