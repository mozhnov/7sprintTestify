package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerCod400(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=ryazan", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки
	assert.Equal(t, responseRecorder.Code, http.StatusBadRequest, "wrong city value")

}
