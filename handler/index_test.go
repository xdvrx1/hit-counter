package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/stretchr/testify/assert"
	"github.com/labstack/echo/v4"
)

func TestHandler_Index(t *testing.T) {
	assert := assert.New(t)

	e := echo.New()

	s, err := miniredis.Run()
	assert.NoError(err)
	defer s.Close()

	h, err := NewHandler([]string{s.Addr()})
	assert.NoError(err)

	r := httptest.NewRequest("GET", "http://localhost:8080", nil)
	w := httptest.NewRecorder()

	hctx := &HitCounterContext{Context: e.NewContext(r, w)}
	err = h.Index(hctx)
	assert.NoError(err)

	resp := w.Result()
	assert.Equal(http.StatusOK, resp.StatusCode)
}
