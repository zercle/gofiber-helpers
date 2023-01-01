package helpers

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"hash/crc64"
	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/influxdata/influxdb/v2/pkg/snowflake"
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
	return fmt.Sprintf("Function: %s \nFile: %s:%d", chopPath(runtime.FuncForPC(function).Name(), "."), file, line)
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
func HashPasswordString(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// HashPassword with bcrypt
func HashPassword(password []byte) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	return bytes, err
}

// CheckPasswordHash with bcrypt
func CheckPasswordHashString(password, hash string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return
}

// CheckPasswordHash with bcrypt
func CheckPasswordHash(password, hash []byte) (err error) {
	err = bcrypt.CompareHashAndPassword(hash, password)
	return
}

// RandomKey with crc64(UUIDv4)
func RandomKey() (string, error) {
	seed, err := UUIDv4()
	if err != nil {
		return "", err
	}

	hasher := crc64.New(crc64.MakeTable(crc64.ECMA))
	_, err = hasher.Write([]byte(seed))
	if err != nil {
		return "", err
	}
	bs := hasher.Sum(nil)

	return hex.EncodeToString(bs), nil
}

// RandomHash with sha3-256
func RandomHash() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	hasher := sha3.New256()
	_, err = hasher.Write(b)
	if err != nil {
		return "", err
	}

	bs := hasher.Sum(nil)

	return hex.EncodeToString(bs), nil
}

// RemoveIndex from array
func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func UUIDv4() (uuidStr string, err error) {
	uuid4, err := uuid.NewRandom()
	uuidStr = uuid4.String()
	return
}

func StructSumSha256(obj interface{}) string {
	hash := sha3.Sum256([]byte(fmt.Sprintf("%v", obj)))

	return hex.EncodeToString(hash[:])
}

func StructCheckSumSha256(obj interface{}, checksum string) error {
	hash := sha3.Sum256([]byte(fmt.Sprintf("%v", obj)))
	cs, err := hex.DecodeString(checksum)
	if err != nil {
		return err
	}

	if !bytes.Equal(hash[:], cs) {
		return fiber.NewError(http.StatusBadRequest, fmt.Sprintf("Expect: %v\nResult: %x", checksum, hash[:]))
	}
	return nil
}

// InitSnowflake machine id base on process id
func InitSnowflake(machineIDs ...int) (snowflakeGen *snowflake.Generator) {
	machineID := (os.Getpid() % 1023)
	if len(machineIDs) == 1 {
		machineID = machineIDs[0]
	}
	snowflakeGen = snowflake.New(machineID)
	return
}

func ExtractAuthString(authStr string) (auth HttpAuth, err error) {
	authSlice := strings.Split(authStr, " ")
	if len(authSlice) != 2 {
		err = fiber.NewError(http.StatusUnauthorized, "invalid authorize format")
		return
	}
	switch strings.ToLower(authSlice[0]) {
	case string(BasicAuth):
		basicBytes, deErr := base64.StdEncoding.DecodeString(authSlice[1])
		if deErr != nil {
			err = fiber.NewError(http.StatusUnauthorized, err.Error())
			return
		}
		basicSlice := bytes.Split(basicBytes, []byte(":"))
		auth = HttpAuth{
			Type:     BasicAuth,
			Username: string(basicSlice[0]),
			Password: string(basicSlice[1]),
		}
	case string(BearerToken):
		auth = HttpAuth{
			Type:  BearerToken,
			Token: authSlice[1],
		}
	default:
		err = fiber.NewError(http.StatusUnauthorized, fmt.Sprintf("invalid authorize type: %s", authSlice[0]))
	}
	return
}
