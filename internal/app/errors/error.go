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
	ErrNotFoundUser          = NewError(404, "NotFoundUser", "user not found")
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
	ErrNotFoundSite      = NewError(404, "NotFoundSite", "site not found")
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
)
