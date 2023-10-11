package models

const (
	DAO_STATE_PENDING = iota
	DAO_STATE_APPROVED
	DAO_STATE_DENIED
)

type DAO struct {
	Address            string `json:"address" bson:"addr"`
	Contract           string `json:"network" bson:"contract,omitempty"`
	Name               string `json:"name" bson:"name,omitempty"`
	Description        string `json:"description" bson:"description,omitempty"`
	Framework          string `json:"framework" bson:"framwork,omitempty"`
	MembersUri         string `json:"membersUri" bson:"membersUri,omitempty"`
	ProposalsUri       string `json:"proposalsUri" bson:"proposalsUri,omitempty"`
	IssuersUri         string `json:"issuersUri" bson:"issuersUri,omitempty"`
	ContractsRegUri    string `json:"contractsRegUri" bson:"contractsRegUri,omitempty"`
	ManagerAddress     string `json:"managerAddress" bson:"managerAddress,omitempty"`
	GovernanceDocument string `json:"governanceDocument" bson:"governanceDocument,omitempty"`
	State              int    `json:"state" bson:"state,omitempty"`
}
type DAOFilter struct {
	Address string `bson:"addr,omitempty"`
	Name    string `bson:"name,omitempty"`
	State   int    `bson:"state,omitempty"`
}
type DAOsResponse struct {
	DAOs  []DAO `json:"daos"`
	Count int   `json:"count"`
}
type DAOVerifyRequest struct {
	Validate bool `json:"validate"`
}
 