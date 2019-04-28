package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/stretchr/testify/assert"
	"github.com/labstack/echo/v4"
)

func TestHandler_Error(t *testing.T) {
	assert := assert.New(t)
	e := echo.New()

	s, err := miniredis.Run()
	assert.NoError(err)
	defer s.Close()

	h, err := NewHandler([]string{s.Addr()}, 10)
	assert.NoError(err)

	request := httptest.NewRequest("GET", "http://localhost", nil)
	w := httptest.NewRecorder()
	ectx := e.NewContext(request, w)

	h.Error(fmt.Errorf("[err] test"), ectx)

	resp := w.Result()
	assert.Equal(http.StatusInternalServerError, resp.StatusCode)
}
