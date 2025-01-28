package dtos

type (
	HTTPMsg struct {
		Message string `json:"message"`
	}
	HTTPErrMsg struct {
		Error string `json:"error"`
	}
	HTTPErrs struct {
		Errors map[string][]string `json:"errors"`
	}
)
