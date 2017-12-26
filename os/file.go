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

// Copy a file.
func Copy(oldpath, newpath string) error {
	src, err := os.Open(oldpath)
	if err != nil {
		return errors.Errorf(codes.InvalidArgument, "failure opening source; %s", err)
	}
	defer src.Close()

	dest, err := os.Create(newpath)
	if err != nil {
		return errors.Errorf(codes.InvalidArgument, "failure creating destination; %s", err)
	}
	defer dest.Close()

	if _, err := io.Copy(dest, src); err != nil {
		return errors.Errorf(codes.FailedPrecondition, "failure during copy; %s", err)
	}
	if err := dest.Sync(); err != nil {
		return errors.Errorf(codes.FailedPrecondition, "failure during sync; %s", err)
	}

	return nil
}
