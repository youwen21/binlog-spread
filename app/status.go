package app

const (
	SUCCESS = 200
	ERROR = 9999
)

var statusText = map[int]string{
	SUCCESS: "success",
	ERROR: "system error",
}

// StatusText returns a text for the HTTP status code. It returns the empty
// string if the code is unknown.
func StatusText(code int) string {
	return statusText[code]
}
