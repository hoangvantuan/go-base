package hhttp

const (
	tokenNotFound    = "token not found"
	tokenInvalid     = "token is invalid"
	tokenExpired     = "token is expired"
	parameterInvalid = "parameter is invalid"
	dataNotfound     = "data not found"
	systemError      = "system error"
)

type cerrors struct {
	message string
}

func (ce *cerrors) Error() string {
	return ce.message
}

func newTokenNotFound() error {
	return &cerrors{
		message: tokenNotFound,
	}
}

func newTokenInvalid() error {
	return &cerrors{
		message: tokenInvalid,
	}
}

func newTokenExpired() error {
	return &cerrors{
		message: tokenExpired,
	}
}

func newParameterInvalid() error {
	return &cerrors{
		message: parameterInvalid,
	}
}

func newDataNotFound() error {
	return &cerrors{
		message: dataNotfound,
	}
}

func newSystemError() error {
	return &cerrors{
		message: systemError,
	}
}
