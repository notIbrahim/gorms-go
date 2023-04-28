package handler

import (
	Commons "api-go/handler/common/response"
	"errors"
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

type ApplicationError struct {
	Message string
	Code    string
	Cause   error
}

func (Error *ApplicationError) getErrorMessage() string {
	return Error.Message
}

func (Error *ApplicationError) getErrorCode() string {
	return Error.Code
}

func (Error *ApplicationError) Error() string {
	EntryMessage := Error.getErrorMessage()
	if Error.Message != "" {
		EntryMessage = fmt.Sprintf("%s (type: %s, retryable: %v)", EntryMessage, Error.Code, Error.Cause)
	}

	if Error.Cause != nil {
		EntryMessage = fmt.Sprintf("%s: %v", EntryMessage, Error.Cause)
	}

	return EntryMessage
}

func ApplicationHandler(MessageReceive string, CodeReceive string, CauseReceive error) error {
	AppHandler := &ApplicationError{
		Message: MessageReceive,
		Code:    CodeReceive,
		Cause:   CauseReceive,
	}

	return AppHandler
}

// Application Handler if Error are handled
func SetApplicationHandler(ErrMessage string, ErrCode string, ErrCause error) error {
	return ApplicationHandler(ErrMessage, ErrCode, ErrCause)
}

// Application Handler if Error being unhandled
func SetApplicationHandlerUnhandled(ErrMessage string, ErrCode string) error {
	return ApplicationHandler(ErrMessage, ErrCode, nil)
}

// Error Baseline implementation
func BaseError(Err error, ErrCode int, ErrMessage string) error {
	if _, file, line, ok := runtime.Caller(1); ok {
		path := strings.Split(file, "/")
		ErrMessage = fmt.Sprintf("%s/%s:%d", ErrMessage, path[len(path)-1], line)
	}

	if Err == nil {
		return SetApplicationHandlerUnhandled(ErrMessage, ErrorCodeTypecast(ErrCode))
	}
	return SetApplicationHandler(ErrMessage, ErrorCodeTypecast(ErrCode), Err)
}

// Typecast Error String to Int

func ErrorCodeTypecast(ErrCode int) string {
	return fmt.Sprintf("%d", ErrCode)
}

// ApplicationError Struct : Function Message Handler
func ErrorMessages(Err error) string {
	return Err.Error()
}

// ApplicationError Struct : Function Code Handler
func ErrorCode(Err error) int {
	var ApplicationErrors *ApplicationError
	if errors.As(Err, &ApplicationErrors) {
		LineCode, _ := strconv.Atoi(ApplicationErrors.getErrorCode())
		return LineCode
	}
	return Commons.Success
}
