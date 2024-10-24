package xerrors

type XError string

func (e XError) Error() string {
	return string(e)
}
