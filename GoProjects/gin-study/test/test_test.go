package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gin-study/models"
	"gin-study/routers"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLog(t *testing.T) {
	fmt.Println("test log")
}

func TestLog2(t *testing.T) {
	fmt.Println("test log2")
}

func TestPintRoute(t *testing.T) {
	router := routers.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user/ping", nil)
	router.ServeHTTP(w, req)

	fmt.Println(w.Body.String())

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestPostRoute(t *testing.T) {
	router := routers.SetupRouter()

	w := httptest.NewRecorder()

	exampleUser := models.User{
		ID:        1234,
		LoginName: "mike",
	}

	userJson, _ := json.Marshal(exampleUser)
	req, _ := http.NewRequest("POST", "/user/get", bytes.NewBuffer(userJson))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	fmt.Println(w.Body.String())
	fmt.Println(userJson)
	assert.Equal(t, string(userJson), w.Body.String())
}
