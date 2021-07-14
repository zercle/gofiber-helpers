package helpers

import (
	"crypto/rand"
	"fmt"
	"runtime"
	"strings"

	"golang.org/x/crypto/sha3"

	"golang.org/x/crypto/bcrypt"
)

// GetStackTrace prints the stack trace with given error by the formatter.
// If the error is not traceable, empty string is returned.
func GetStackTrace() (trace []string) {
	stackBuf := make([]uintptr, 10)
	length := runtime.Callers(1, stackBuf[:])
	if length == 0 {
		// No stackBuf available. Stop now.
		// This can happen if the first argument to runtime.Callers is large.
		return
	}
	// pass only valid pcs to runtime.CallersFrames
	stack := stackBuf[:length]
	frames := runtime.CallersFrames(stack)

	// Loop to get frames.
	// A fixed number of stackBuf can expand to an indefinite number of Frames.
	for {
		frame, more := frames.Next()
		// To keep this output stable
		// even if there are changes in the testing package,
		// stop unwinding when we leave package runtime.
		// if !strings.Contains(frame.File, "runtime/") {
		// continue
		// }
		trace = append(trace, fmt.Sprintf("File: %s, Line: %d. Function: %s", frame.File, frame.Line, frame.Function))
		if !more {
			break
		}
	}
	return
}

// WhereAmI return current path and line
func WhereAmI(skipList ...int) string {
	var skip int
	if skipList == nil {
		skip = 1
	} else {
		skip = skipList[0]
	}
	function, file, line, _ := runtime.Caller(skip)
	return fmt.Sprintf("File: %s, Line: %d. Function: %s", chopPath(file, "/"), line, chopPath(runtime.FuncForPC(function).Name(), "."))
}

// return the source filename after the last slash
func chopPath(original string, pathChar string) string {
	i := strings.LastIndex(original, pathChar)
	if i == -1 {
		return original
	}
	return original[i+1:]
}

// HashPassword with bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash with bcrypt
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// RandomHash with sha3-256
func RandomHash() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	h := sha3.New256()
	h.Write(b)

	bs := h.Sum(nil)

	return string(bs), nil
}

// RemoveIndex from array
func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
