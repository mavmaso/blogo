package main

import (
	"blogo/controllers"
	"blogo/models"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHome(t *testing.T) {
	router := setup()
	router.GET("/", controllers.GetHome)

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	refuse(err, nil, t)

	expected := httptest.NewRecorder()
	router.ServeHTTP(expected, req)
	fmt.Println(expected.Body)

	refuse(expected.Code, http.StatusOK, t)
}

func TestIndex(t *testing.T) {
	router := setup()
	router.GET("/posts", controllers.Index)

	req, err := http.NewRequest(http.MethodGet, "/posts", nil)
	refuse(err, nil, t)

	expected := httptest.NewRecorder()
	router.ServeHTTP(expected, req)
	fmt.Println(expected.Body)

	refuse(expected.Code, http.StatusOK, t)
}

func setup() *gin.Engine {
	models.ConnectDatabase()
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	return router
}

func refuse(x interface{}, y interface{}, t *testing.T) {
	if x != y {
		t.Fatalf("Couldn't create request: %v\n", x)
	}
}
