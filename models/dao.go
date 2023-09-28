package models

const (
	DAO_STATE_PENDING = iota
	DAO_STATE_APPROVED
	DAO_STATE_DENIED
)

type DAO struct {
	Address            string `json:"addr" bson:"addr"`
	Contract           string `json:"contract" bson:"contract"`
	Name               string `json:"name" bson:"name"`
	Description        string `json:"description" bson:"description"`
	Framwork           string `json:"framwork" bson:"framwork"`
	MembersUri         string `json:"members_uri" bson:"membersUri"`
	ProposalsUri       string `json:"proposals_uri" bson:"proposalsUri"`
	IssuersUri         string `json:"issuers_uri" bson:"issuersUri"`
	ContractsRegUri    string `json:"contracts_reg_uri" bson:"contractsRegUri"`
	ManagerAddress     string `json:"manager_addr" bson:"managerAddress"`
	GovernanceDocument string `json:"governance_doc" bson:"governanceDocument"`
	State              int    `json:"state" bson:"state"`
}
type DAOFilter struct {
	Address string `bson:"addr,omitempty"`
	Name    string `bson:"name,omitempty"`
	State   int    `bson:"state,omitempty"`
}
