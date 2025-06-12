package constant

import "fmt"

// Err 错误信息
//
//	@author centonhuang
//	@update 2025-06-12 21:01:52
type Err string

const (

	// ErrRequiredArgumentNotFound ErrorMessage 缺少必填参数
	//	@update 2025-06-12 21:01:00
	ErrRequiredArgumentNotFound Err = "required argument not found: %s"
)

// Errorf 格式化错误信息
//
//	@receiver e ErrorMessage
//	@param args
//	@return error
//	@author centonhuang
//	@update 2025-06-12 21:01:43
func (e Err) Errorf(args ...any) error {
	return fmt.Errorf(string(e), args...)
}
