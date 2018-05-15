package utils

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path"
	"time"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
)

func IsEmpty(val string) bool {
	if val == "" {
		return true
	} else {
		return false
	}
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func IsNotEmpty(val string) bool {
	return !IsEmpty(val)
}

func Equals(val1 string, val2 string) bool {
	return val1 == val2
}

//
// Convert Timestamp to yyyy-MM-dd HH:mm:sss
//
func TimeToString(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func TimeToYyyyMmDd(t time.Time) string {
	return t.Format("2006-01-02")
}

func GetTimeStamp() string {
	now := time.Now().Local()
	return now.Format("20060102150405")
}

func ToTimeStamp(t time.Time) string {
	now := t.Local()
	return now.Format("20060102150405")
}

// GeneratePassword hash password then set to model
func GeneratePassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

// ComparePassword verify password and hashed
func ComparePassword(hashPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}

func DayToDuration(day time.Duration) time.Duration {
	return time.Hour * 24 * day
}

func Decode(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, traditionalchinese.Big5.NewDecoder())
	d, e := ioutil.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func Base64Decode(str string) (string, bool) {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", true
	}
	return string(data), false
}

func RandStringBytes(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func Sha1(msg string) string {
	h := sha1.New()
	h.Write([]byte(msg))
	hash := h.Sum(nil)
	key := hex.EncodeToString(hash)
	return key
}

func InfInt64ToInt64(inf interface{}) int64 {
	return inf.(int64)
}

func InfInt32ToInt32(inf interface{}) int32 {
	return inf.(int32)
}

func InfInt16ToInt16(inf interface{}) int16 {
	return inf.(int16)
}

func InfIntToInt(inf interface{}) int {
	return inf.(int)
}

func InfFloat64ToFloat64(inf interface{}) float64 {
	return inf.(float64)
}

func InfFloat32ToFloat32(inf interface{}) float32 {
	return inf.(float32)
}

/**
 * https://github.com/labstack/echo/issues/871
 * https://github.com/kysnm/echo_log_practice/blob/master/main.go
 */
func LoggerConfig(pathname string, filename string) *os.File {
	logfile := path.Join(pathname, filename)
	f, err := os.Open(logfile) // For read access.
	defer f.Close()
	if err != nil {
		fmt.Println("Create logs directory.")
		os.MkdirAll(pathname, 0755)
	}
	file, err := os.OpenFile(logfile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Create", logfile)
	}
	log.SetOutput(file)

	return file
}
