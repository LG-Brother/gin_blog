package model

type Model struct {
	ID        uint32 `gorm:"primary_key" json:"id"`
	CreatedOn uint32 `json:"created_on"`
	CreatedBy string `json:"created_by"`
	UpdatedOn uint32 `json:"update_on"`
	UpdatedBy string `json:"updated_by"`
	DeletedOn uint32 `json:"deleted_on"`
	IsDel     uint8  `json:"is_del"`
}
