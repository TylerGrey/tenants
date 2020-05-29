package repo

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Bldg 건물
type Bldg struct {
	ID        uint64 `gorm:"primary_key"`
	Addr      string
	Lat       float64
	Lng       float64
	Rating    float32
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// BldgRepository ...
type BldgRepository interface {
	Create(bldg Bldg) (*Bldg, error)
	FindByLatLng(lat float64, lng float64) (*Bldg, error)
}

// bldgRepository 인터페이스 구조체
type bldgRepository struct {
	master  *gorm.DB
	replica *gorm.DB
}

// NewBldgRepository ...
func NewBldgRepository(master *gorm.DB, replica *gorm.DB) BldgRepository {
	return &bldgRepository{
		master:  master,
		replica: replica,
	}
}

// Create 리뷰 생성
func (r bldgRepository) Create(bldg Bldg) (*Bldg, error) {
	if err := r.master.Table("bldg").Create(&bldg).Error; err != nil {
		return nil, err
	}

	return &bldg, nil
}

func (r bldgRepository) FindByLatLng(lat float64, lng float64) (*Bldg, error) {
	bldg := &Bldg{}

	if err := r.replica.Table("bldg").Where("lat = ? AND lng = ?", lat, lng).Find(bldg).Error; gorm.IsRecordNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return bldg, nil
}
