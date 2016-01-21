package btrfs

import "fmt"

func unexpectedResult(data []byte) (u *UnexpectedResponseError) {
	u = new(UnexpectedResponseError)
	u.Data = data
	return u
}

type UnexpectedResponseError struct {
	Data []byte
}

func (e *UnexpectedResponseError) Error() string {
	return fmt.Sprintf("Unexpected Response: %s\n", e.Data)
}
