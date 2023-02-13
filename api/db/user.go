package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/minhtam3010/sql-injection/db/entity"
)

func (q *Querier) CreateUser(data entity.User) (id int64, err error) {
	var res sql.Result
	date := ConvertDateToMySQLTime(time.Now().Unix())
	res, err = q.DB.Exec("INSERT INTO users (username, password, date_created) VALUES (?, ?, ?)", data.UserName, data.Password, date)
	if err != nil {
		return
	}
	id, err = res.LastInsertId()
	return
}

func (q *Querier) Login(username, password string) (check bool, err error) {
	var res entity.User
	var date_created string
	var rows *sql.Rows
	var id int64

	rows, err = q.DB.Query(fmt.Sprintf("SELECT * FROM users WHERE username = '%s' AND password = '%s'", username, password))
	if err != nil {
		return
	}
	for rows.Next() {
		err = rows.Scan(&id, &res.UserName, &res.Password, &date_created)
		if err != nil {
			return
		}
	}
	if id != 0 {
		return true, err
	}
	return false, err
}

func ConvertDateToMySQLTime(unixTime int64) string {
	return time.Unix(unixTime, 0).Format("2006-01-02 15:04:05")
}
