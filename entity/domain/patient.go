package domain

type PatientEntity struct {
	Id       uint   `json:"id"`
	Name     string `json:"name" validate:"required"`
	Age      int    `json:"age" validate:"required"`
	Gender   string `json:"gender" validate:"oneof=male female"`
	Status   string `json:"status" validate:"oneof=positive recovered dead"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}
