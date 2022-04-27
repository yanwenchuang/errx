package errx

import "fmt"


type CodeErr struct {
	errCode uint32
	errMsg  string
}

func (e *CodeErr) GetErrCode() uint32 {
	return e.errCode
}

func (e *CodeErr) GetErrMsg() string {
	return e.errMsg
}

func (e *CodeErr) Error() string {
	return fmt.Sprintf("ErrCode:%d,ErrMsg:%s", e.errCode, e.errMsg)
}

func NewErrCodeMsg(errCode uint32, errMsg string) *CodeErr {
	return &CodeErr{errCode: errCode, errMsg: errMsg}
}

func NewErrCode(errCode uint32) *CodeErr {
	return &CodeErr{errCode: errCode, errMsg: MapErrMsg(errCode)}
}

func NewErrMsg(errMsg string) *CodeErr {
	return &CodeErr{errCode: 100001, errMsg: errMsg}
}
