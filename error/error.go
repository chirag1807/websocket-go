package errorhandling

import "net/http"

type CustomError struct {
	StatusCode int
	Message    string
}

// here i implemented error interface's Error() method.
func (c CustomError) Error() string {
	return c.Message
}

var (
	ReadBodyError           = CreateCustomError("Could not Read Request Body, Please Provide Valid Body.", http.StatusBadRequest)
	ReadDataError           = CreateCustomError("Could not Decode the Data, Please Provide Valid Data.", http.StatusBadRequest)
	EmailvalidationError    = CreateCustomError("Email Validation Failed, Please Provide Valid Email.", http.StatusBadRequest)
	DuplicateEmailFound     = CreateCustomError("Duplicate Email Found.", http.StatusConflict)
	RegistrationFailedError = CreateCustomError("User Registration Failed.", http.StatusInternalServerError)
	LoginFailedError        = CreateCustomError("User Login Failed.", http.StatusUnauthorized)
	AccessTokenExpired      = CreateCustomError("Access Token is Expired, Please Regenrate It.", http.StatusUnauthorized)
	RefreshTokenExpired     = CreateCustomError("Access Token is Expired, Please Do Login Again.", http.StatusUnauthorized)
	RefreshTokenError       = CreateCustomError("Access Token Can't be Regenerated, Please Do Login Again.", http.StatusUnauthorized)
	UnauthorizedError       = CreateCustomError("You are Not Authorized to Perform this Action.", http.StatusUnauthorized)
	NoUserFound             = CreateCustomError("No User Found for This Request.", http.StatusNotFound)
	ChatNotAdded            = CreateCustomError("This Chat Can not Be Added to Database.", http.StatusBadRequest)
	RoomNotCreated          = CreateCustomError("This Room Can not be Created.", http.StatusBadRequest)
	MemberNotAdded          = CreateCustomError("This member Can not be Added to Room.", http.StatusBadRequest)
	MemberNotRemoved        = CreateCustomError("This member Can not be Removed from Room.", http.StatusBadRequest)
)

func CreateCustomError(Message string, StatusCode int) error {
	return CustomError{
		Message:    Message,
		StatusCode: StatusCode,
	}
}
