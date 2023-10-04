package models

const (
	DAO_STATE_PENDING = iota
	DAO_STATE_APPROVED
	DAO_STATE_DENIED
)

type DAO struct {
	Address            string `json:"addr" bson:"addr"`
	Contract           string `json:"contract" bson:"contract,omitempty"`
	Name               string `json:"name" bson:"name,omitempty"`
	Description        string `json:"description" bson:"description,omitempty"`
	Framework          string `json:"framework" bson:"framwork,omitempty"`
	MembersUri         string `json:"members_uri" bson:"membersUri,omitempty"`
	ProposalsUri       string `json:"proposals_uri" bson:"proposalsUri,omitempty"`
	IssuersUri         string `json:"issuers_uri" bson:"issuersUri,omitempty"`
	ContractsRegUri    string `json:"contracts_reg_uri" bson:"contractsRegUri,omitempty"`
	ManagerAddress     string `json:"manager_addr" bson:"managerAddress,omitempty"`
	GovernanceDocument string `json:"governance_doc" bson:"governanceDocument,omitempty"`
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
