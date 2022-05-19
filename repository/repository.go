package repository

import (
	"pil/domain"
	"pil/internal/model"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

// CreateActivity implements domain.AllRepository
func (r *repository) CreateActivity(activity model.Activity) {
	r.DB.Create(&activity)
}

// DeleteActivity implements domain.AllRepository
func (r *repository) DeleteActivity(id int) {
	r.DB.Unscoped().Where("id", id).Delete(&model.Activity{})
}

// GetAllActivity implements domain.AllRepository
func (r *repository) GetAllActivity() []model.APIResponseActivity {
	var activity []model.APIResponseActivity
	r.DB.Table("activities").Order("id desc").Scan(&activity)
	return activity
}

// GetByIdActivity implements domain.AllRepository
func (r *repository) GetByIdActivity(id int) model.APIResponseActivity {
	var activity model.APIResponseActivity
	r.DB.Table("activities").Where("id", id).Scan(&activity)
	return activity
}

// UpdateActivity implements domain.AllRepository
func (r *repository) UpdateActivity(id int, activity model.Activity) {
	r.DB.Model(&model.Activity{}).Select("judul", "deskripsi", "link").Where("id", id).Updates(activity)
}

func NewRepository(db *gorm.DB) domain.AllRepository {
	return &repository{
		DB: db,
	}
}
