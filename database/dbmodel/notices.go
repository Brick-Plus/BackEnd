package dbmodel

import "gorm.io/gorm"

type Notice struct {
	gorm.Model
	IdUser   int    `json:"id_user"`
	User     User   `gorm:"foreignKey:IdUser"`
	IdProduct int   `json:"id_product"`
	Product  Product `gorm:"foreignKey:IdProduct"`
	Content  string `json:"content"`
}

type NoticesRepository interface {
	Create(newNotice *Notice) (*Notice, error)
	FindNoticesByProductId(productId string) ([]*Notice, error)
	FindNoticesByUserId(userId string) ([]*Notice, error)
	Delete(noticeToDelete *Notice) error
	Update(noticeToUpdate *Notice, noticeId string) error
	FindNoticeById(noticeId string) ([]*Notice, error)
}

type noticesRepository struct {
	db *gorm.DB
}

func NewNoticesRepository(db *gorm.DB) NoticesRepository {
	return &noticesRepository{db: db}
}

func (r *noticesRepository) Create(notice *Notice) (*Notice, error) {
	if err := r.db.Create(notice).Error; err != nil {
		return nil, err
	}
	return notice, nil
}

func (r *noticesRepository) FindNoticesByProductId(productId string) ([]*Notice, error) {
	var entries []*Notice
	if err := r.db.Where("id_product = ?", productId).Find(&entries).Error; err != nil {
		return nil, err
	}
	return entries, nil
}

func (r *noticesRepository) FindNoticesByUserId(userId string) ([]*Notice, error) {
	var entries []*Notice
	if err := r.db.Where("id_user = ?", userId).Find(&entries).Error; err != nil {
		return nil, err
	}
	return entries, nil
}

func (r *noticesRepository) FindNoticeById(noticeId string) ([]*Notice, error) {
	var entries []*Notice
	if err := r.db.Where("id = ?", noticeId).Find(&entries).Error; err != nil {
		return nil, err
	}
	return entries, nil
}

func (r *noticesRepository) Delete(noticeToDelete *Notice) error {
	if err := r.db.Delete(noticeToDelete).Error; err != nil {
		return err
	}
	return nil
}

func (r *noticesRepository) Update(noticeToUpdate *Notice, noticeId string) error {
	if err := r.db.Where("id = ?", noticeId).Updates(noticeToUpdate).Error; err != nil {
		return err
	}
	return nil
}