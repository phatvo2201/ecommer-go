package common

import "time"

type SQLModel struct {
	Id        int        `json:"-" gorm:"column:id"` // khong dua ra json ma dung fake id
	FakeId    *UID       `json:"id" gorm:"-"`        //db ko co cai nay
	Status    int        `json:"status" gorm:"column:status;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdateAt  *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

//mask id thnah object fakeid
func (sqlModel *SQLModel) Mask(dbType DbType) {
	uid := NewUID(uint32(sqlModel.Id), int(dbType), 1)
	sqlModel.FakeId = &uid
}

func (sqlModel *SQLModel) PrepareForInsert() {
	now := time.Now().UTC()
	sqlModel.Id = 0
	sqlModel.Status = 1
	sqlModel.CreatedAt = &now
	sqlModel.UpdateAt = &now
}
