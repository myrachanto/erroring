package httperors

////////////errors ////////////////////////
type HttpError struct {
	message string
	code    int
	errors  string
	results bool
}
