package service

func ErrToJSON(err error) map[string]error {
	return map[string]error{
		"error": err,
	}
}
