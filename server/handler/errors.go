package handler

const (
	GENERIC_ERROR          = "oops, something went wrong"
	SERVICE_ERROR          = "the service is not available, please try again later"
	NO_INT_ERROR           = "message id must be an integer"
	NOT_FOUND_ERROR        = "message with this id is not found"
	BAD_REQUEST_ERROR      = "request format is not correct"
	MESSAGE_TOO_LONG_ERROR = "message content must be less than 255 characters"
)
