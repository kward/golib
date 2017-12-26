/*
Package sysexits provides exit status codes for system programs.

This package is based on errors provided in `/usr/include/sysexits.h`.

The following example calls os.Exit with an OK (0) exit code.

```go
package main

import (
	"os"
	"github.com/kward/golib/os/sysexits"
)

func main() {
	fmt.Println("Hello, world!")
  os.Exit(sysexits.OK.Int())
}
```

The following code snippet exits with a Software (70) error code. Note that
os.Exit is implicitly called by Exitf.

```go
package main

import (
	"github.com/kward/golib/os/sysexits"
)

func main() {
	sysexits.Exitf("some error", sysexits.Software)
}
```
*/
package sysexits

import (
	"fmt"
	"os"
)

// Code is an error code for os.Exit().
type Code int

//go:generate stringer -type=Code

const (
	Usage        Code = iota + 64 // Command-line usage error.
	DataError                     // Data format error.
	NoInput                       // Cannot open input.
	NoUser                        // Addressee unknown.
	NoHost                        // Host name unknown.
	Unavailable                   // Service unavailable.
	Software                      // Internal software error.
	OSError                       // System error (e.g., can't fork).
	OSFile                        // Critical OS file missing.
	CantCreate                    // Can't create (user) output file.
	ToError                       // Input/output error.
	TempFailure                   // temp failure; user is invited to retry
	Protocol                      // remote error in protocol
	NoPermission                  // permission denied
	Config                        // configuration error
)

const (
	OK   Code = 0      // successful termination
	Base      = Usage  // base value for error messages
	Max       = Config // maximum listed value

	// Aliases for compatibility.
	DataErr   = DataError
	OSErr     = OSError
	CantCreat = CantCreate
	ToErr     = ToError
	TempFail  = TempFailure
	NoPerm    = NoPermission
)

// Exit calls os.Exit with the given code.
func Exit(code int) { os.Exit(code) }

// Exitf prints a message to STDERR and calls Exit with the given code.
func Exitf(code int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	Exit(code)
}

// Int returns the int value of the code.
func (c Code) Int() int { return int(c) }
