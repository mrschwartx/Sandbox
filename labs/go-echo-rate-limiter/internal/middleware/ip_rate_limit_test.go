package middleware

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestIPRateLimiter(t *testing.T) {
	ipStore = sync.Map{}
	e := echo.New()

	handler := func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	}

	limited := IPRateLimiter(2, time.Minute)(handler)

	req1 := httptest.NewRequest(http.MethodGet, "/", nil)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	err := limited(c1)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec1.Code)

	req2 := httptest.NewRequest(http.MethodGet, "/", nil)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)
	err = limited(c2)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec2.Code)

	req3 := httptest.NewRequest(http.MethodGet, "/", nil)
	rec3 := httptest.NewRecorder()
	c3 := e.NewContext(req3, rec3)
	err = limited(c3)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusTooManyRequests, rec3.Code)
}
