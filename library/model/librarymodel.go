package model

import (
	"database/sql"
	"fmt"
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
	LibraryModel interface {
		Insert(data Library) (sql.Result, error)
		FindOne(id string) (*Library, error)
		FindOneByName(name string) (*Library, error)
		Update(data Library) error
		Delete(id string) error
	}

	defaultLibraryModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Library struct {
		Name        string    `db:"name"`   // 书籍名称
		Author      string    `db:"author"` // 书籍作者
		CreateTime  time.Time `db:"create_time"`
		PublishDate time.Time `db:"publish_date"`
		UpdateTime  time.Time `db:"update_time"`
		Id          string    `db:"id"` // 书籍序列号
	}
)

func NewLibraryModel(conn sqlx.SqlConn) LibraryModel {
	return &defaultLibraryModel{
		conn:  conn,
		table: "library",
	}
}

func (m *defaultLibraryModel) Insert(data Library) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, libraryRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Name, data.Author, data.PublishDate, data.Id)
	return ret, err
}

func (m *defaultLibraryModel) FindOne(id string) (*Library, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", libraryRows, m.table)
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

func (m *defaultLibraryModel) FindOneByName(name string) (*Library, error) {
	var resp Library
	query := fmt.Sprintf("select %s from %s where name = ? limit 1", libraryRows, m.table)
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

func (m *defaultLibraryModel) Update(data Library) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, libraryRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Name, data.Author, data.PublishDate, data.Id)
	return err
}

func (m *defaultLibraryModel) Delete(id string) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
