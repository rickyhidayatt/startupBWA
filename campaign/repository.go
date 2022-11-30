package campaign

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Campaign, error)
	FindByUserID() ([]Campaign, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Campaign, error) {

	var campagins []Campaign

	err := r.db.Preload("CampaignImages", "campaigm_images.is_primary = 1").Find(&campagins).Error
	if err != nil {
		return campagins, err
	}
	return campagins, nil
}

func (r *repository) FindByUserID(userID int) ([]Campaign, error) {
	var campagins []Campaign

	err := r.db.Where("user_id = ?", userID).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campagins).Error //"Campaign adalah field dari entity go (struct), campaigm_images di dapat dari nama tabel di database dan .is_primary maksudnya kita cuma mau ngambil data is primary yang datanya =1"

	if err != nil {
		return campagins, err
	}

	return campagins, nil
}
