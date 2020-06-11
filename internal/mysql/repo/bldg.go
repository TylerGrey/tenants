package repo

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Bldg 건물
type Bldg struct {
	ID          uint64 `gorm:"primary_key"`
	Addr        string
	Lat         float64
	Lng         float64
	Rating      float64
	Address     string
	RoadAddress string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

// BldgRepository ...
type BldgRepository interface {
	Create(bldg Bldg) (*Bldg, error)
	FindByLatLng(lat float64, lng float64) (*Bldg, error)
	List(lat float64, lng float64) ([]*Bldg, error)
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

func (r bldgRepository) List(lat float64, lng float64) ([]*Bldg, error) {
	var bldgs []*Bldg
	if err := r.replica.Table("bldg").
		Select("*, ST_DISTANCE_SPHERE(POINT(?, ?), POINT(lng, lat)) AS dist", lng, lat).
		Having("dist < 500").
		Order("dist").
		Find(&bldgs).Error; gorm.IsRecordNotFoundError(err) {
		return bldgs, err
	} else if err != nil {
		return nil, err
	}

	return bldgs, nil
}
