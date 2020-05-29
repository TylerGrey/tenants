package repo

import (
	"fmt"
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
	Cursor               string `gorm:"-"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            *time.Time
}

// ReviewRepository ...
type ReviewRepository interface {
	Create(review Review) (*Review, error)
	Update(review Review) (*Review, error)
	Delete(id uint64) (bool, error)

	FindByID(id uint64) (*Review, error)
	List(bldgID uint64, args ListArgs) ([]*Review, PageInfo, error)
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

// FindByID ID로 Hub 조회
func (r reviewRepository) FindByID(id uint64) (*Review, error) {
	review := &Review{}

	if err := r.replica.Table("review").Where("id = ?", id).Scan(review).Error; gorm.IsRecordNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return review, nil
}

// List ...
func (r reviewRepository) List(bldgID uint64, args ListArgs) ([]*Review, PageInfo, error) {
	var (
		reviews                 []*Review
		count                   int32
		baseCursor              = getBaseCursor(args)
		cursor, cursorDirection = getCursor(args)
		field, direction        = getOrderBy(args)
		limit                   = getLimit(args)
	)

	tx := r.replica.Table("review")
	tx = tx.Select(fmt.Sprintf("*, TO_BASE64(%s) as `cursor`", baseCursor))
	tx = tx.Where("bldg_id = ?", bldgID)
	tx.Where("deleted_at IS NULL").Count(&count)

	if len(cursor) > 0 {
		tx = tx.Where(fmt.Sprintf("%s %s FROM_BASE64(?)", baseCursor, cursorDirection), cursor)
	}

	tx = tx.Order(fmt.Sprintf("%s %s", field, direction))
	tx = tx.Order(fmt.Sprintf("id %s", direction))
	tx = tx.Limit(limit + 1)

	if err := tx.Find(&reviews).Error; err != nil {
		return nil, PageInfo{}, err
	}

	list, pageInfo := getPageInfo(args, limit, &reviews)
	pageInfo.Total = count

	return list.([]*Review), pageInfo, nil
}
