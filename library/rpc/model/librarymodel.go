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
	libraryFieldNames          = builderx.FieldNames(&Library{})
	libraryRows                = strings.Join(libraryFieldNames, ",")
	libraryRowsExpectAutoSet   = strings.Join(stringx.Remove(libraryFieldNames, "create_time", "update_time"), ",")
	libraryRowsWithPlaceHolder = strings.Join(stringx.Remove(libraryFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	LibraryModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Library struct {
		Id          string    `db:"id"`     // 书籍序列号
		Name        string    `db:"name"`   // 书籍名称
		Author      string    `db:"author"` // 书籍作者
		PublishDate time.Time `db:"publish_date"`
		CreateTime  time.Time `db:"create_time"`
		UpdateTime  time.Time `db:"update_time"`
	}
)

func NewLibraryModel(conn sqlx.SqlConn) *LibraryModel {
	return &LibraryModel{
		conn:  conn,
		table: "library",
	}
}

func (m *LibraryModel) Insert(data Library) (sql.Result, error) {
	query := `insert into ` + m.table + ` (` + libraryRowsExpectAutoSet + `) values (?, ?, ?, ?)`
	ret, err := m.conn.Exec(query, data.Id, data.Name, data.Author, data.PublishDate)
	return ret, err
}

func (m *LibraryModel) FindOne(id string) (*Library, error) {
	query := `select ` + libraryRows + ` from ` + m.table + ` where id = ? limit 1`
	var resp Library
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

func (m *LibraryModel) FindOneByName(name string) (*Library, error) {
	var resp Library
	query := `select ` + libraryRows + ` from ` + m.table + ` where name = ? limit 1`
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

func (m *LibraryModel) Update(data Library) (sql.Result, error) {
	query := `update ` + m.table + ` set ` + libraryRowsWithPlaceHolder + ` where id = ?`
	ret, err := m.conn.Exec(query, data.Name, data.Author, data.PublishDate, data.Id)
	return ret, err
}

func (m *LibraryModel) Delete(id string) error {
	query := `delete from ` + m.table + ` where id = ?`
	_, err := m.conn.Exec(query, id)
	return err
}
