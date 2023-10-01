package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/AwespireTech/dXCA-Backend/database"
	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	err := database.Init("mongodb://localhost:27017")
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	gin.SetMode(gin.TestMode)
	m.Run()
}
func getRouter(t *testing.T) *gin.Engine {
	t.Helper()
	return createRouter()
}
func TestDefaultRoute(t *testing.T) {
	recorder := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("failed to create mock request: %v", err)
	}
	router := getRouter(t)
	router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, recorder.Code)
	}

}
func TestGetAllDAOs(t *testing.T) {
	recorder := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/api/dao", nil)
	if err != nil {
		t.Fatalf("failed to create mock request: %v", err)
	}
	router := getRouter(t)
	router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, recorder.Code)
	}
}
func TestGetDAOByAddr(t *testing.T) {
	address := "0xEDaf4083F29753753d0Cd6c3C50ACEb08c87b5BD"
	recorder := httptest.NewRecorder()
	url := fmt.Sprintf("/api/dao/%s", url.QueryEscape(address))
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatalf("failed to create mock request: %v", err)
	}
	router := getRouter(t)
	router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, recorder.Code)
	}
}

func TestCreateDAO(t *testing.T) {
	recorder := httptest.NewRecorder()
	request, err := http.NewRequest("POST", "/api/dao", nil)
	if err != nil {
		t.Fatalf("failed to create mock request: %v", err)
	}
	router := getRouter(t)
	router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusCreated {
		t.Errorf("expected status code %d, got %d", http.StatusCreated, recorder.Code)
	}
}
func TestCancelDAO(t *testing.T) {
	//TODO Vaildation
	recorder := httptest.NewRecorder()
	request, err := http.NewRequest("DELETE", "/api/dao/12", nil)
	if err != nil {
		t.Fatalf("failed to create mock request: %v", err)
	}
	router := getRouter(t)
	router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, recorder.Code)
	}
}
func TestValidateDAO(t *testing.T) {
	//TODO Vaildation
	body := []byte(`{"validate":true}`)
	recorder := httptest.NewRecorder()
	request, err := http.NewRequest("POST", "/api/dao/12", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("failed to create mock request: %v", err)
	}
	router := getRouter(t)
	router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, recorder.Code)
	}
}
