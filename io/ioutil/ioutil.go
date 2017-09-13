/*
Package ioutil implements IO utility functions.
*/
package ioutil

import (
	"io"
	"os"

	"github.com/kward/golib/errors"
	"google.golang.org/grpc/codes"
)

// CopyFile copies a file, returning the number of bytes copied.
func CopyFile(destName, srcName string) (int64, error) {
	src, err := os.Open(srcName)
	if err != nil {
		return -1, errors.Errorf(codes.InvalidArgument, "failure opening source; %s", err)
	}
	defer src.Close()

	dest, err := os.Create(destName)
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
