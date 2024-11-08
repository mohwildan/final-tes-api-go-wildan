package sql

import "app/domain/model/sql"


type UserRepository interface {
    CreateUser(user *sql.User) error
    DetailUserByEmail(email string) (*sql.User, error)
    FindUserByID(id int) (*sql.User, error)
    FindAllUser() ([]sql.User, error)
}

func (r *repository) CreateUser(user *sql.User) error {
    return r.db.Create(user).Error
}

func (r *repository) DetailUserByEmail(email string) (*sql.User, error) {
    var user sql.User
    err := r.db.Where("email = ?", email).First(&user).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *repository) FindUserByID(id int) (*sql.User, error) {
    var user sql.User
    err := r.db.Where("id = ?", id).First(&user).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *repository) FindAllUser() ([]sql.User, error) {
    var users []sql.User
    err := r.db.Find(&users).Error
    if err != nil {
        return nil, err
    }
    return users, nil
}

