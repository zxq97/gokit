package errcode

import "fmt"

type BizErr struct {
	Code    int32
	Message string
	Detail  string
}

func (e *BizErr) Error() string {
	return fmt.Sprintf("biz error code:%d message:%s detail:%s", e.Code, e.Message, e.Detail)
}

func NewBizErr(code int32, message, detail string) *BizErr {
	return &BizErr{
		Code:    code,
		Message: message,
		Detail:  detail,
	}
}
