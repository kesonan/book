package model

import (
	"database/sql"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	userFieldNames          = builderx.FieldNames(&User{})
	userRows                = strings.Join(userFieldNames, ",")
	userRowsExpectAutoSet   = strings.Join(stringx.Remove(userFieldNames, "id", "create_time", "update_time"), ",")
	userRowsWithPlaceHolder = strings.Join(stringx.Remove(userFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	UserModel struct {
		conn  sqlx.SqlConn
		table string
	}

	User struct {
		Id         int64     `db:"id"`
		Name       string    `db:"name"`     // 用户名称
		Password   string    `db:"password"` // 用户密码
		Mobile     string    `db:"mobile"`   // 手机号
		Gender     string    `db:"gender"`   // 男｜女｜未公开
		Nickname   string    `db:"nickname"` // 用户昵称
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
	}
)

func NewUserModel(conn sqlx.SqlConn) *UserModel {
	return &UserModel{
		conn:  conn,
		table: "user",
	}
}

func (m *UserModel) Insert(data User) (sql.Result, error) {
	query := `insert into ` + m.table + ` (` + userRowsExpectAutoSet + `) values (?, ?, ?, ?, ?)`
	ret, err := m.conn.Exec(query, data.Name, data.Password, data.Mobile, data.Gender, data.Nickname)
	return ret, err
}

func (m *UserModel) FindOne(id int64) (*User, error) {
	query := `select ` + userRows + ` from ` + m.table + ` where id = ? limit 1`
	var resp User
	err := m.conn.QueryRow(&resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *UserModel) FindOneByName(name string) (*User, error) {
	var resp User
	query := `select ` + userRows + ` from ` + m.table + ` where name = ? limit 1`
	err := m.conn.QueryRow(&resp, query, name)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *UserModel) FindOneByMobile(mobile string) (*User, error) {
	var resp User
	query := `select ` + userRows + ` from ` + m.table + ` where mobile = ? limit 1`
	err := m.conn.QueryRow(&resp, query, mobile)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *UserModel) Update(data User) (sql.Result, error) {
	query := `update ` + m.table + ` set ` + userRowsWithPlaceHolder + ` where id = ?`
	ret, err := m.conn.Exec(query, data.Name, data.Password, data.Mobile, data.Gender, data.Nickname, data.Id)
	return ret, err
}

func (m *UserModel) Delete(id int64) error {
	query := `delete from ` + m.table + ` where id = ?`
	_, err := m.conn.Exec(query, id)
	return err
}
