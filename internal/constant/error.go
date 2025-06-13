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

	// ErrRequireCredentials 需要认证
	//	@update 2025-06-13 16:52:01
	ErrRequireCredentials Err = "require credentials"

	// ErrInvalidCredentials 认证失败
	//	@update 2025-06-13 16:48:01
	ErrInvalidCredentials Err = "invalid credentials"
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
