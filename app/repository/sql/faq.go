package sql

import "app/domain/model/sql"

// FaqRepository defines methods for managing FAQs
type FaqRepository interface {
    CreateFaq(faq *sql.Faq) error
    GetFaqByQuestion(question string) (*sql.Faq, error)
    FindFaqByID(id string) (*sql.Faq, error)
    FindAllFaqs() ([]sql.Faq, error)
}

// CreateFaq inserts a new FAQ record
func (r *repository) CreateFaq(faq *sql.Faq) error {
    return r.db.Create(faq).Error
}

// GetFaqByQuestion retrieves a FAQ by its question text
func (r *repository) GetFaqByQuestion(question string) (*sql.Faq, error) {
    var faq sql.Faq
    err := r.db.Where("question = ?", question).First(&faq).Error
    if err != nil {
        return nil, err
    }
    return &faq, nil
}

// FindFaqByID retrieves a FAQ by its ID
func (r *repository) FindFaqByID(id string) (*sql.Faq, error) {
    var faq sql.Faq
    err := r.db.Where("id = ?", id).First(&faq).Error
    if err != nil {
        return nil, err
    }
    return &faq, nil
}

// FindAllFaqs retrieves all FAQs
func (r *repository) FindAllFaqs() ([]sql.Faq, error) {
    var faqs []sql.Faq
    err := r.db.Find(&faqs).Error
    if err != nil {
        return nil, err
    }
    return faqs, nil
}
