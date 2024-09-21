package custom_models

import (
	"github.com/pocketbase/pocketbase/models"
)

var _ models.Model = (*TransactionHistory)(nil)

type TransactionHistory struct {
	models.BaseModel
	EntityType   string `db:"entityType" json:"entityType"`
	EntityId     string `db:"entityId" json:"entityId"`
	StatusCodeId string `db:"statusCodeId" json:"statusCodeId"`
	Note         string `db:"note" json:"note,omitempty"`
}

func (m *TransactionHistory) TableName() string {
	return "transaction_history"
}