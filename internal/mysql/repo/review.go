package repo

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Review 리뷰
type Review struct {
	ID                   uint64 `gorm:"primary_key"`
	BldgID               uint64
	Title                string
	Content              string
	ScoreRent            int32
	ScoreMaintenanceFees int32
	ScorePublicTransport int32
	ScoreConvenience     int32
	ScoreLandlord        int32
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            *time.Time
}

// ReviewRepository ...
type ReviewRepository interface {
	Create(review Review) (*Review, error)
	Update(review Review) (*Review, error)
	Delete(id uint64) (bool, error)
}

// reviewRepository 인터페이스 구조체
type reviewRepository struct {
	master  *gorm.DB
	replica *gorm.DB
}

// NewReviewRepository ...
func NewReviewRepository(master *gorm.DB, replica *gorm.DB) ReviewRepository {
	return &reviewRepository{
		master:  master,
		replica: replica,
	}
}

// Create 리뷰 생성
func (r reviewRepository) Create(review Review) (*Review, error) {
	if err := r.master.Table("review").Create(&review).Error; err != nil {
		return nil, err
	}

	return &review, nil
}

// Update 리뷰 수정
func (r reviewRepository) Update(review Review) (*Review, error) {
	if err := r.master.Table("review").Model(&review).Update(&review).Error; err != nil {
		return nil, err
	}

	response := &Review{}
	if err := r.master.Table("review").First(&response, review.ID).Error; err != nil {
		return nil, err
	}

	return response, nil
}

// Delete 리뷰 삭제
func (r reviewRepository) Delete(id uint64) (bool, error) {
	query := r.master.Table("review").Delete(&Review{
		ID: id,
	})
	if query.Error != nil {
		return false, query.Error
	}

	return query.RowsAffected > 0, nil
}
