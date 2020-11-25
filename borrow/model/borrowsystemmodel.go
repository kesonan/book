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

const (
	_ = iota
	Borrowing
	Return
)

var (
	borrowSystemFieldNames          = builderx.FieldNames(&BorrowSystem{})
	borrowSystemRows                = strings.Join(borrowSystemFieldNames, ",")
	borrowSystemRowsExpectAutoSet   = strings.Join(stringx.Remove(borrowSystemFieldNames, "id", "create_time", "update_time"), ",")
	borrowSystemRowsWithPlaceHolder = strings.Join(stringx.Remove(borrowSystemFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	BorrowSystemModel interface {
		Insert(data BorrowSystem) (sql.Result, error)
		FindOne(id int64) (*BorrowSystem, error)
		Update(data BorrowSystem) error
		Delete(id int64) error
		FindOneByUserAndBookNo(userId int64, bookNo string) (*BorrowSystem, error)
		FindOneByBookNo(bookNo string, status int) (*BorrowSystem, error)
	}

	defaultBorrowSystemModel struct {
		conn  sqlx.SqlConn
		table string
	}

	BorrowSystem struct {
		BookNo         string    `db:"book_no"` // 书籍号
		UserId         int64     `db:"user_id"` // 借书人
		CreateTime     time.Time `db:"create_time"`
		ReturnDate     int64     `db:"return_date"`      // 实际还书时间
		ReturnPlanDate time.Time `db:"return_plan_date"` // 预计还书时间
		Status         int64     `db:"status"`           // 书籍状态，0-未归还，1-已归还
		UpdateTime     time.Time `db:"update_time"`
		Id             int64     `db:"id"`
	}
)

func NewBorrowSystemModel(conn sqlx.SqlConn) BorrowSystemModel {
	return &defaultBorrowSystemModel{
		conn:  conn,
		table: "borrow_system",
	}
}

func (m *defaultBorrowSystemModel) Insert(data BorrowSystem) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, borrowSystemRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.BookNo, data.UserId, data.ReturnDate, data.ReturnPlanDate, data.Status)
	return ret, err
}

func (m *defaultBorrowSystemModel) FindOne(id int64) (*BorrowSystem, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", borrowSystemRows, m.table)
	var resp BorrowSystem
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

func (m *defaultBorrowSystemModel) Update(data BorrowSystem) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, borrowSystemRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.BookNo, data.UserId, data.ReturnDate, data.ReturnPlanDate, data.Status, data.Id)
	return err
}

func (m *defaultBorrowSystemModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}

func (m *defaultBorrowSystemModel) FindOneByUserAndBookNo(userId int64, bookNo string) (*BorrowSystem, error) {
	query := `select ` + borrowSystemRows + ` from ` + m.table + ` where user_id = ? and book_no = ? limit 1`
	var resp BorrowSystem
	err := m.conn.QueryRow(&resp, query, userId, bookNo)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultBorrowSystemModel) FindOneByBookNo(bookNo string, status int) (*BorrowSystem, error) {
	query := `select ` + borrowSystemRows + ` from ` + m.table + ` where book_no = ? and status = ? limit 1`
	var resp BorrowSystem
	err := m.conn.QueryRow(&resp, query, bookNo, status)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
