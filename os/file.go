/*
Package os implements OS utility functions.
*/
package os

import (
	"io"
	"os"

	"github.com/kward/golib/errors"
	"google.golang.org/grpc/codes"
)

// Copy a file, returning the number of bytes copied.
func Copy(oldpath, newpath string) (int64, error) {
	src, err := os.Open(oldpath)
	if err != nil {
		return -1, errors.Errorf(codes.InvalidArgument, "failure opening source; %s", err)
	}
	defer src.Close()

	dest, err := os.Create(newpath)
	if err != nil {
		return -1, errors.Errorf(codes.InvalidArgument, "failure creating destination; %s", err)
	}
	defer dest.Close()

	written, err := io.Copy(dest, src)
	if err != nil {
		return written, errors.Errorf(codes.FailedPrecondition, "failure during copy; %s", err)
	}
	if err := dest.Sync(); err != nil {
		return written, errors.Errorf(codes.FailedPrecondition, "failure during sync; %s", err)
	}

	return written, nil
}
