package zerror

import (
	"errors"
	"fmt"
)

type ZError struct {
	ErrCode ZErrorCode `json:"err_code,omitempty"`
	ErrMsg  string     `json:"err_msg,omitempty"`
	errs    []error
}

func (e *ZError) Error() string {
	if e == nil {
		return ""
	}
	if e.ErrMsg != "" {
		return fmt.Sprintf("ErrCode:%s; ErrMsg:%s;", e.ErrCode, e.ErrMsg)
	}
	res := ""
	if e.errs == nil || len(e.errs) == 0 {
		return res
	}
	var first = true
	for _, err := range e.errs {
		if first {
			res = err.Error()
			first = false
		} else {
			res += ";" + err.Error()
		}
	}
	return res
}
func (e *ZError) Errors() []error {
	if e == nil {
		return nil
	}
	return e.errs
}
func (e *ZError) Append(err error) {
	if e == nil || err == nil {
		return
	}
	ze, ok := err.(*ZError)
	if ok {
		e.errs = append(e.errs, ze.errs...)
	} else {
		e.errs = append(e.errs, err)
	}
}

func NewByErr(err ...error) error {
	res := &ZError{
		errs: make([]error, 0),
	}
	for i, e := range err {
		if e == nil {
			continue
		}
		ze, ok := err[i].(*ZError)
		if ok {
			res.errs = append(res.errs, ze.errs...)
		} else {
			res.errs = append(res.errs, err[i])
		}
	}
	if len(res.errs) > 0 {
		return res
	}
	return nil
}
func NewByCode(errCode ZErrorCode, errMsg ...string) error {
	msg := ""
	if len(errMsg) > 0 {
		msg = errMsg[0]
	} else {
		msg = getErrMsg(errCode)
	}
	return &ZError{
		ErrCode: errCode,
		ErrMsg:  msg,
	}
}
func NewByMsg(msg string) error {
	err := errors.New(msg)
	return NewByErr(err)
}
func Errors(err error) []error {
	if err == nil {
		return nil
	}
	ze, ok := err.(*ZError)
	if !ok {
		return []error{err}
	}
	//将一个空的[]error切片与ze.Errors()返回的错误切片合并在一起
	//返回一个新切片
	return append(([]error)(nil), ze.Errors()...)
}
