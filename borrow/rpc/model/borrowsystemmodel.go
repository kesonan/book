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
	BorrowSystemModel struct {
		conn  sqlx.SqlConn
		table string
	}

	BorrowSystem struct {
		Id             int64     `db:"id"`
		BookNo         string    `db:"book_no"`          // 书籍号
		UserId         int64     `db:"user_id"`          // 借书人
		Status         int64     `db:"status"`           // 书籍状态，0-未归还，1-已归还
		ReturnPlanDate time.Time `db:"return_plan_date"` // 预计还书时间
		ReturnDate     time.Time `db:"return_date"`      // 实际还书时间
		CreateTime     time.Time `db:"create_time"`
		UpdateTime     time.Time `db:"update_time"`
	}
)

func NewBorrowSystemModel(conn sqlx.SqlConn) *BorrowSystemModel {
	return &BorrowSystemModel{
		conn:  conn,
		table: "borrow_system",
	}
}

func (m *BorrowSystemModel) Insert(data BorrowSystem) (sql.Result, error) {
	query := `insert into ` + m.table + ` (` + borrowSystemRowsExpectAutoSet + `) values (?, ?, ?, ?, ?)`
	ret, err := m.conn.Exec(query, data.BookNo, data.UserId, data.Status, data.ReturnPlanDate, data.ReturnDate)
	return ret, err
}

func (m *BorrowSystemModel) FindOne(id int64) (*BorrowSystem, error) {
	query := `select ` + borrowSystemRows + ` from ` + m.table + ` where id = ? limit 1`
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

func (m *BorrowSystemModel) FindOneByBookNoAndStatus(bookNo string, status int) (*BorrowSystem, error) {
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
func (m *BorrowSystemModel) FindOneByBookNo(bookNo string) (*BorrowSystem, error) {
	query := `select ` + borrowSystemRows + ` from ` + m.table + ` where book_no = ? limit 1`
	var resp BorrowSystem
	err := m.conn.QueryRow(&resp, query, bookNo)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *BorrowSystemModel) Update(data BorrowSystem) (sql.Result, error) {
	query := `update ` + m.table + ` set ` + borrowSystemRowsWithPlaceHolder + ` where id = ?`
	ret, err := m.conn.Exec(query, data.BookNo, data.UserId, data.Status, data.ReturnPlanDate, data.ReturnDate, data.Id)
	return ret, err
}

func (m *BorrowSystemModel) Delete(id int64) error {
	query := `delete from ` + m.table + ` where id = ?`
	_, err := m.conn.Exec(query, id)
	return err
}
