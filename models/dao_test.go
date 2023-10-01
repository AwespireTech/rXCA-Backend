package models

import (
	"encoding/hex"
	"math/rand"
	"testing"
)

func RandomDAO(t *testing.T) DAO {
	t.Helper()

	var dao DAO
	dao.Address = RandomAddress(t)
	dao.Contract = RandomAddress(t)
	dao.Name = "Test DAO"
	dao.Description = "Test DAO Description"
	dao.Framwork = "Test DAO Framework"
	dao.MembersUri = "Test DAO Members URI"
	dao.ProposalsUri = "Test DAO Proposals URI"
	dao.IssuersUri = "Test DAO Issuers URI"
	dao.ContractsRegUri = "Test DAO Contracts Reg URI"
	dao.ManagerAddress = RandomAddress(t)
	dao.GovernanceDocument = "Test DAO Governance Document"
	dao.State = rand.Intn(3)
	return dao

}
func RandomPandingDAO(t *testing.T) DAO {
	t.Helper()
	dao := RandomDAO(t)
	dao.State = DAO_STATE_PENDING
	return dao
}
func RandomApprovedDAO(t *testing.T) DAO {
	t.Helper()
	dao := RandomDAO(t)
	dao.State = DAO_STATE_APPROVED
	return dao
}
func RandomDeniedDAO(t *testing.T) DAO {
	t.Helper()
	dao := RandomDAO(t)
	dao.State = DAO_STATE_DENIED
	return dao
}
func RandomAddress(t *testing.T) string {
	t.Helper()
	addr := "0x"
	addr += hex.EncodeToString([]byte{byte(rand.Intn(255))})
	return addr
}
