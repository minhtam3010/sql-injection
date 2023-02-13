package middleware

import (
	"github.com/minhtam3010/sql-injection/db"
	"github.com/minhtam3010/sql-injection/db/entity"
)

type UserUsecase struct {
	Querier *db.Querier
}

func NewUserUsecase() *UserUsecase {
	return &UserUsecase{
		Querier: db.NewQuerier(db.Connect()),
	}
}

func (u *UserUsecase) CreateUser(data entity.User) (id int64, err error) {
	id, err = u.Querier.CreateUser(data)
	if err != nil {
		u.Querier.DB.Rollback()
	}
	return id, u.Querier.DB.Commit()
}

func (u *UserUsecase) Login(username, password string) (check bool, err error) {
	return u.Querier.Login(username, password)
}
