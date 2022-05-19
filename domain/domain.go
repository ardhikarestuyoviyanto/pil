package domain

import "pil/internal/model"

type AllRepository interface {
	CreateActivity(activity model.Activity)
	UpdateActivity(id int, activity model.Activity)
	DeleteActivity(id int)
	GetAllActivity() []model.APIResponseActivity
	GetByIdActivity(id int) model.APIResponseActivity
}

type AllServices interface {
	CreateActivity(activity model.Activity)
	UpdateActivity(id int, activity model.Activity)
	DeleteActivity(id int)
	GetAllActivity() []model.APIResponseActivity
	GetByIdActivity(id int) model.APIResponseActivity
}
