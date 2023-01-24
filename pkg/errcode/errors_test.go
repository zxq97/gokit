package errcode

import (
	"errors"
	"testing"
)

func TestNewBizErr(t *testing.T) {
	bizerr := &BizErr{}
	var err error
	err = NewBizErr(1, "msg", "xxx")
	if !errors.As(err, &bizerr) {
		t.Error("err should as BizErr")
	}
}
