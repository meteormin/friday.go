package errors

type Error struct {
	Code    int
	Title   string
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

func NewError(code int, title, message string) *Error {
	return &Error{
		Code:    code,
		Title:   title,
		Message: message,
	}
}

// User
var (
	ErrNotFoundUser        = NewError(404, "NotFoundUser", "user not found")
	ErrInvalidUserName     = NewError(400, "InvalidUserName", "invalid name")
	ErrInvalidUserUsername = NewError(400, "InvalidUserUsername", "invalid username")
	ErrInvalidUserPassword = NewError(400, "InvalidUserPassword", "invalid password")
)

// File
var (
	ErrInvalidFileName = NewError(400, "InvalidFileName", "invalid file name")
	ErrInvalidFile     = NewError(400, "InvalidFile", "invalid file")
)
