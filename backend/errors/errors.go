// errors/errors.go

package errors

import "fmt"

// CustomError 定义错误结构体，包含错误码和错误信息
type CustomError struct {
	Code    int    // 错误码
	Message string // 错误消息
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Code %d: %s", e.Code, e.Message)
}

// 用户相关错误码
const (
	ErrUserNotFoundCode     = 2001
	ErrUsernameTakenCode    = 2002
	ErrEmailTakenCode       = 2003
	ErrInvalidPasswordCode  = 2004
	ErrPermissionDeniedCode = 2005
	ErrUserAuthFailedCode   = 2006
)

// 文章相关错误码
const (
	ErrPostNotFoundCode         = 3001
	ErrPostTitleTakenCode       = 3002
	ErrPostUpdatePermissionCode = 3003
)

// 评论相关错误码
const (
	ErrCommentNotFoundCode         = 4001
	ErrEmptyCommentContentCode     = 4002
	ErrCommentDeletePermissionCode = 4003
)

// 系统相关错误码
const (
	ErrDatabaseConnectionCode  = 5001
	ErrExternalServiceCode     = 5002
	ErrInternalServerErrorCode = 5003
)

// 输入校验错误码
const (
	ErrInvalidInputCode         = 6001
	ErrInvalidRequestFormatCode = 6002
)

// Token 相关错误码
const (
	ErrTokenGenerationCode       = 7001 // token 生成失败
	ErrTokenSignatureInvalidCode = 7002 // token 签名无效
	ErrTokenExpiredCode          = 7003 // token 已过期
	ErrTokenDecodeFailedCode     = 7004 // token 解码失败
)

// Token 相关错误
var (
	ErrTokenGeneration       = &CustomError{Code: ErrTokenGenerationCode, Message: "failed to generate token"}
	ErrTokenSignatureInvalid = &CustomError{Code: ErrTokenSignatureInvalidCode, Message: "invalid token signature"}
	ErrTokenExpired          = &CustomError{Code: ErrTokenExpiredCode, Message: "token has expired"}
	ErrTokenDecodeFailed     = &CustomError{Code: ErrTokenDecodeFailedCode, Message: "failed to decode token"}
)

// 用户相关错误
var (
	ErrUserNotFound     = &CustomError{Code: ErrUserNotFoundCode, Message: "user not found"}
	ErrUsernameTaken    = &CustomError{Code: ErrUsernameTakenCode, Message: "username already taken"}
	ErrEmailTaken       = &CustomError{Code: ErrEmailTakenCode, Message: "email already taken"}
	ErrInvalidPassword  = &CustomError{Code: ErrInvalidPasswordCode, Message: "invalid password"}
	ErrPermissionDenied = &CustomError{Code: ErrPermissionDeniedCode, Message: "permission denied"}
	ErrUserAuthFailed   = &CustomError{Code: ErrUserAuthFailedCode, Message: "authentication failed"}
)

// 文章相关错误
var (
	ErrPostNotFound         = &CustomError{Code: ErrPostNotFoundCode, Message: "post not found"}
	ErrPostTitleTaken       = &CustomError{Code: ErrPostTitleTakenCode, Message: "post title already exists"}
	ErrPostUpdatePermission = &CustomError{Code: ErrPostUpdatePermissionCode, Message: "you do not have permission to update this post"}
)

// 评论相关错误
var (
	ErrCommentNotFound         = &CustomError{Code: ErrCommentNotFoundCode, Message: "comment not found"}
	ErrEmptyCommentContent     = &CustomError{Code: ErrEmptyCommentContentCode, Message: "comment content cannot be empty"}
	ErrCommentDeletePermission = &CustomError{Code: ErrCommentDeletePermissionCode, Message: "you do not have permission to delete this comment"}
)

// 系统相关错误
var (
	ErrDatabaseConnection  = &CustomError{Code: ErrDatabaseConnectionCode, Message: "failed to connect to the database"}
	ErrExternalService     = &CustomError{Code: ErrExternalServiceCode, Message: "external service error"}
	ErrInternalServerError = &CustomError{Code: ErrInternalServerErrorCode, Message: "internal server error"}
)

// 输入校验错误
var (
	ErrInvalidInput         = &CustomError{Code: ErrInvalidInputCode, Message: "invalid input"}
	ErrInvalidRequestFormat = &CustomError{Code: ErrInvalidRequestFormatCode, Message: "invalid request format"}
)

// NewError 工厂函数，用于创建带有自定义信息的错误
func NewError(code int, message string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}
