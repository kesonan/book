package logic

import (
	"book/shared"
)

var (
	errUserNotFound = shared.NewDefaultError("用户不存在")
	errBookNotFound = shared.NewDefaultError("书籍不存在")
	errInvalidParam = shared.NewDefaultError("参数错误")
	errUserReturn   = shared.NewDefaultError("没有查询到该用户的借书记录")
	errBookBorrowed = shared.NewDefaultError("该书籍已被借阅")
	errBookReturn   = shared.NewDefaultError("该书籍已归还")
)
