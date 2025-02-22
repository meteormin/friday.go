package app

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

// Commons
var (
	ErrForbidden = NewError(403, "Forbidden", "forbidden")
	ErrNotFound  = NewError(404, "NotFound", "not found")
)

// User
var (
	ErrInvalidUserName       = NewError(400, "InvalidUserName", "invalid name")
	ErrInvalidUserUsername   = NewError(400, "InvalidUserUsername", "invalid username")
	ErrInvalidUserPassword   = NewError(400, "InvalidUserPassword", "invalid password")
	ErrDuplicateUserUsername = NewError(409, "DuplicateUser", "username already exists")
)

// File
var (
	ErrInvalidFileName = NewError(400, "InvalidFileName", "invalid file name")
	ErrInvalidFile     = NewError(400, "InvalidFile", "invalid file")
)

// Site
var (
	ErrInvalidSiteHost   = NewError(400, "InvalidSiteHost", "invalid site host")
	ErrInvalidSiteName   = NewError(400, "InvalidSiteName", "invalid site name")
	ErrDuplicateSiteHost = NewError(409, "DuplicateSite", "site host already exists")
	ErrDuplicateSiteName = NewError(409, "DuplicateSite", "site name already exists")
)

// Post
var (
	ErrInvalidPostTitle   = NewError(400, "InvalidPostTitle", "invalid post title")
	ErrInvalidPostContent = NewError(400, "InvalidPostContent", "invalid post content")
	ErrInvalidPostFile    = NewError(400, "InvalidFileId", "invalid file id")
	ErrInvalidPostSite    = NewError(400, "InvalidTags", "invalid tags")
	ErrInvalidPostTags    = NewError(400, "InvalidTags", "invalid tags")
	ErrInvalidPostPath    = NewError(400, "InvalidPostPath", "invalid post path")
	ErrDuplicatePostPath  = NewError(409, "DuplicatePostPath", "post path already exists")
)
