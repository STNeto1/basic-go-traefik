package __errors

type ErrUserNotFound struct{}

func (e ErrUserNotFound) Error() string {
	return "user not found"
}
