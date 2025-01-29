package customerrors

import (
	"fmt"
	apperror "ladipage_server/core/app_error"
	"net/http"
)

type CustomError struct {
	error
	Status  int    `json:"status,omitempty"`
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func NewError(err error, status, code int, message string) *CustomError {
	return &CustomError{
		error:   err,
		Status:  status,
		Code:    code,
		Message: message,
	}
}

func (u *CustomError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", u.Code, u.Message)
}

func (u *CustomError) String() string {
	return u.Error()
}

var (
	ErrDB = NewError(
		apperror.ErrorDataBase,
		http.StatusInternalServerError,
		apperror.ErrorDataBaseCode,
		apperror.ErrorDataBase.Error(),
	)

	ErrAuth = NewError(
		apperror.ErrorAuthentication,
		http.StatusUnauthorized,
		apperror.ErrorAuthCode,
		apperror.ErrorAuthentication.Error(),
	)

	ErrNotFound = NewError(
		apperror.ErrorNotFound,
		http.StatusConflict,
		apperror.ErrorNotFoundCode,
		apperror.ErrorNotFound.Error(),
	)

	ErrUserExists = NewError(
		apperror.ErrorUserExists,
		http.StatusConflict,
		apperror.ErrorUserExistsCode,
		apperror.ErrorUserExists.Error(),
	)

	ErrHashPassword = NewError(
		apperror.ErrorHashPassword,
		http.StatusInternalServerError,
		apperror.ErrorHashPasswordCode,
		apperror.ErrorHashPassword.Error(),
	)

	ErrGenToken = NewError(
		apperror.ErrorGenToken,
		http.StatusInternalServerError,
		apperror.ErrorGenTokenCode,
		apperror.ErrorGenToken.Error(),
	)

	ErrVerifyToken = NewError(
		apperror.VerifyToken,
		http.StatusUnauthorized,
		apperror.ErrorAuthCode,
		apperror.ErrorAuthentication.Error(),
	)

	ErrVerifyTokenEmail = NewError(
		apperror.ErrorVerifyTokenEmail,
		http.StatusUnauthorized,
		apperror.ErrorVerifyTokenEmailCode,
		apperror.ErrorVerifyTokenEmail.Error(),
	)

	ErrorSendEmail = NewError(
		apperror.ErrorSendEmail,
		http.StatusInternalServerError,
		apperror.ErrorSendEmailCode,
		apperror.ErrorSendEmail.Error(),
	)

	ErrCategoryExists = NewError(
		apperror.ErrorCategoryExists,
		http.StatusConflict,
		apperror.ErrorCategoryExistsCode,
		apperror.ErrorCategoryExists.Error(),
	)
)
