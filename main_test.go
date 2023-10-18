package main

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/AwespireTech/dXCA-Backend/blockchain"
	"github.com/AwespireTech/dXCA-Backend/database"
	"github.com/AwespireTech/dXCA-Backend/models"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	err := database.Init("mongodb://localhost:27017")
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	gin.SetMode(gin.TestMode)
	err = blockchain.Init("https://eth-sepolia.g.alchemy.com/v2/s65QiLyN-74IJZtUJgtWJiZ9gGzUfxOm")
	if err != nil {
		log.Fatalf("failed to connect to eth: %v", err)
	}
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
	//Check if cors is enabled
	if recorder.Header().Get("Access-Control-Allow-Origin") != "*" {
		t.Errorf("cors is not enabled, expected %s, got %s", "*", recorder.Header().Get("Access-Control-Allow-Origin"))
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
	t.Log(recorder.Body)

}
func TestGetDAOByAddr(t *testing.T) {
	address := InsertRandomDAO(t).Address
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
	t.Log(recorder.Body)
}

func TestCreateDAO(t *testing.T) {
	dao := RandomDAO(t)
	body, err := json.Marshal(dao)
	if err != nil {
		t.Fatalf("failed to marshal dao: %v", err)
	}
	recorder := httptest.NewRecorder()
	request, err := http.NewRequest("POST", "/api/dao", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("failed to create mock request: %v", err)
	}
	router := getRouter(t)
	router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusCreated {
		t.Errorf("expected status code %d, got %d", http.StatusCreated, recorder.Code)
	}

	t.Log(recorder.Body)

}
func TestCancelDAO(t *testing.T) {
	dao := RandomPandingDAO(t)
	err := database.InsertDAO(dao)
	if err != nil {
		t.Fatalf("failed to insert dao: %v", err)
	}
	address := dao.Address
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/api/dao/%s", url.QueryEscape(address))
	request, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		t.Fatalf("failed to create mock request: %v", err)
	}

	router := getRouter(t)
	router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, recorder.Code)
	}
	_, err = database.GetDAOByAddress(address)
	if err.Error() != mongo.ErrNoDocuments.Error() {
		t.Errorf("DAO not deleted")
	}
	t.Log(recorder.Body)

}
func TestValidateDAO(t *testing.T) {
	//TODO Vaildation
	dao := RandomPandingDAO(t)
	err := database.InsertDAO(dao)
	if err != nil {
		t.Fatalf("failed to insert dao: %v", err)
	}
	address := dao.Address
	body := []byte(`{"validate":true}`)
	recorder := httptest.NewRecorder()
	url := fmt.Sprintf("/api/dao/%s", url.QueryEscape(address))
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("failed to create mock request: %v", err)
	}
	router := getRouter(t)
	router.ServeHTTP(recorder, request)
	// No hash, failed approval
	if recorder.Code != http.StatusBadRequest {
		t.Errorf("expected status code %d, got %d", http.StatusBadRequest, recorder.Code)
	}
	t.Log("No hash, failed approval")
	t.Log(recorder.Body)
	// No hash, success denial
	body = []byte(`{"validate":false}`)
	recorder = httptest.NewRecorder()
	request, err = http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("failed to create mock request: %v", err)
	}
	router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, recorder.Code)
	}
	t.Log("No hash, success denial")
	t.Log(recorder.Body)

}
func InsertRandomDAO(t *testing.T) models.DAO {
	t.Helper()
	dao := RandomDAO(t)
	err := database.InsertDAO(dao)
	if err != nil {
		t.Errorf("failed to insert dao: %v", err)
	}
	return dao
}
func RandomDAO(t *testing.T) models.DAO {
	t.Helper()

	var dao models.DAO
	dao.Address = RandomAddress(t)
	dao.Contract = RandomAddress(t)
	dao.Name = "Test DAO"
	dao.Description = "Test DAO Description"
	dao.Framework = "Test DAO Framework"
	dao.MembersUri = "Test DAO Members URI"
	dao.ProposalsUri = "Test DAO Proposals URI"
	dao.IssuersUri = "Test DAO Issuers URI"
	dao.ContractsRegUri = "Test DAO Contracts Reg URI"
	dao.ManagerAddress = RandomAddress(t)
	dao.GovernanceDocument = "Test DAO Governance Document"
	dao.State = rand.Intn(3)
	return dao

}
func RandomPandingDAO(t *testing.T) models.DAO {
	t.Helper()
	dao := RandomDAO(t)
	dao.State = models.DAO_STATE_PENDING
	return dao
}
func RandomApprovedDAO(t *testing.T) models.DAO {
	t.Helper()
	dao := RandomDAO(t)
	dao.State = models.DAO_STATE_APPROVED
	return dao
}
func RandomDeniedDAO(t *testing.T) models.DAO {
	t.Helper()
	dao := RandomDAO(t)
	dao.State = models.DAO_STATE_DENIED
	return dao
}
func RandomAddress(t *testing.T) string {
	t.Helper()
	addr := "0x"
	addr += hex.EncodeToString([]byte{byte(rand.Intn(2147483647))})
	return addr
}
