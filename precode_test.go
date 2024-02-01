package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	TotalCount := "10"
	req := httptest.NewRequest("GET", "/cafe?count="+TotalCount+"&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	selectedCity := "moscow"
	cafesInCity := strings.Join(cafeList[selectedCity], ",")
	assert.Equal(t, cafesInCity, responseRecorder.Body.String())

	cafesInCityCount := len(cafeList[selectedCity])
	ResponseCafes := strings.Split(responseRecorder.Body.String(), ",")
	ResponseCafesCount := len(ResponseCafes)
	assert.Equal(t, cafesInCityCount, ResponseCafesCount)
}

func TestMainHandlerWrongCity(t *testing.T) {
	nonExistentCity := "orgrimmar"
	req := httptest.NewRequest("GET", "/cafe?count=1984&city="+nonExistentCity, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code
	assert.Equal(t, http.StatusBadRequest, status)

	expectedResponseMsg := "wrong city value"
	assert.Equal(t, expectedResponseMsg, responseRecorder.Body.String())
}

func TestMainHandlerValuesIsOk(t *testing.T) {
	city := "moscow"
	count := "2"
	req := httptest.NewRequest("GET", "/cafe?count="+count+"&city="+city, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code
	require.Equal(t, http.StatusOK, status)
	require.NotEmpty(t, responseRecorder.Body)
}
