package fail

import "net/http"

const (
	Undefined int = iota
	InvalidRequestJSON
	AccountIsExists
	AccountIsNotExists
	IncorrectUserPassword
	InvalidUserPassword
)

type UserErrorData struct {
	httpStatusCode int
	message        string
}

func MessageByID(id int) *UserErrorData {
	list := map[int]*UserErrorData{
		InvalidRequestJSON: {
			httpStatusCode: http.StatusBadRequest,
			message:        "Invalid reqoest body",
		},
		AccountIsExists: {
			httpStatusCode: http.StatusInternalServerError,
			message:        "Account alredy exist",
		},
		AccountIsNotExists: {
			httpStatusCode: http.StatusInternalServerError,
			message:        "Account not found",
		},
		IncorrectUserPassword: {
			httpStatusCode: http.StatusBadRequest,
			message:        "Incorrect user password",
		},
		InvalidUserPassword: {
			httpStatusCode: http.StatusBadRequest,
			message:        "Invalid user password",
		},
	}

	if key, ok := list[id]; ok {
		return key
	}

	return &UserErrorData {
		httpStatusCode: http.StatusInternalServerError,
		message: 		"Internal error"
	}
}
