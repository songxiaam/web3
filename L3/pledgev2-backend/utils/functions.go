package utils

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"golang.org/x/crypto/bcrypt"
	"io"
	"math"
	mrand "math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"
)

func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func UniqueId() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	encryptedString := GetMd5String(base64.URLEncoding.EncodeToString(b))
	return encryptedString[0:16] + Int64ToString(time.Now().Unix()) + encryptedString[26:]
}

func JsonToMap(str string) map[string]interface{} {
	var tempMap map[string]interface{}
	err := json.Unmarshal([]byte(str), &tempMap)
	if err != nil {
		panic(err)
	}
	return tempMap
}

func GenerateCode(figures int) (randNum string) {
	startNum := math.Pow(10, float64(figures))
	number := mrand.New(mrand.NewSource(time.Now().UnixNano())).Int31n(int32(startNum))
	return fmt.Sprintf("%06d", number)
}

func IsPhone(phoneNo string) bool {
	if phoneNo != "" {
		if isOk, _ := regexp.MatchString("^1[0-9]{10}$", phoneNo); !isOk {
			return isOk
		}
	}
	return false
}

func IsNumber(num string) bool {
	if num != "" {
		if isOk, _ := regexp.MatchString("^[0-9]*$", num); !isOk {
			return isOk
		}
	}
	return false
}

func CheckAccountFormat(s string) bool {
	if s != "" {
		isOk, _ := regexp.MatchString(`^[A-Za-z][A-Za-z0-9]{5,19}$`, s)
		return isOk
	}
	return false
}

func IsPassword(pwd string) bool {
	if pwd != "" {
		if isOk, _ := regexp.MatchString(`^[a-zA-Z0-9!@#￥%^&*]{6,20}$`, pwd); isOk {
			return isOk
		}
	}
	return false
}

func IsEmail(email string) bool {
	if email != "" {
		if isOk, _ := regexp.MatchString(`^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`, email); isOk {
			return isOk
		}
	}
	return false
}

func StringToInt64(str string) int64 {
	int64Num, _ := strconv.ParseInt(str, 10, 64)
	return int64Num
}

func Int32ToString(n int32) string {
	i := int64(n)
	//FormatInt(i, 10) , i 转成10进制
	return strconv.FormatInt(i, 10)
}

func StringToInt32(str string) int32 {
	var j int32
	int10, _ := strconv.ParseInt(str, 10, 32)
	j = int32(int10)
	return j
}

func Int64ToString(n int64) string {
	i := int64(n)
	return strconv.FormatInt(i, 10)
}

func Int64ToInt(n int64) int {
	strInt64 := strconv.FormatInt(n, 10)
	// Atoi 是 ASCII to integer 的缩写, 字符串转int
	id16, _ := strconv.Atoi(strInt64)
	return id16
}

func Wrap(num float64, retain int) int64 {
	// Pow10(5) ==> 10^5
	return int64(num * math.Pow10(retain))
}

func Unwrap(num int64, retain int) float64 {
	return float64(num) / math.Pow10(retain)
}

func WrapToFloat64(num float64, retain int) float64 {
	return num * math.Pow10(retain)
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GetRandomString(n int) string {
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	bytes := []byte(str)
	result := []byte{}
	r := mrand.New(mrand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// 随机生成n位数字
func CreateCaptcha(n int) string {
	if n <= 0 {
		return ""
	}
	min := int64(math.Pow10(n - 1))
	max := int64(math.Pow10(n)) - 1
	r := mrand.New(mrand.NewSource(time.Now().UnixNano()))
	randomInt := r.Int63n(max-min+1) + min
	return fmt.Sprintf("%d", randomInt)
}

func HttpGet(url string, header map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(req.Body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("content-type", "application/json")
	for k, v := range header {
		req.Header.Add(k, v)
	}

	client := &http.Client{
		Timeout: time.Second * 5,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	return io.ReadAll(resp.Body)
}

func HttpPost(uri string, header map[string]string, data interface{}, args ...string) ([]byte, error) {
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(req.Body)

	req.Header.Add("content-type", "application/json")
	for k, v := range header {
		req.Header.Add(k, v)
	}

	client := &http.Client{Timeout: time.Second * 5}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	return io.ReadAll(resp.Body)
}

// 精度
func Float64AddToString(fa, fb float64) string {
	decimalA := decimal.NewFromFloat(fa)
	decimalB := decimal.NewFromFloat(fb)
	decimalC := decimalA.Add(decimalB)
	return decimalC.String()
}

// 两个float64相减
func Float64SubToString(fa, fb float64) string {
	decimalA := decimal.NewFromFloat(fa)
	decimalB := decimal.NewFromFloat(fb)
	decimalC := decimalA.Sub(decimalB)
	return decimalC.String()
}

func Float64MulToString(fa, fb float64) string {
	decimalA := decimal.NewFromFloat(fa)
	decimalB := decimal.NewFromFloat(fb)
	decimalC := decimalA.Mul(decimalB)
	return decimalC.String()
}

func Float64DivToString(fa, fb float64) string {
	decimalA := decimal.NewFromFloat(fa)
	decimalB := decimal.NewFromFloat(fb)
	decimalC := decimalA.Div(decimalB)
	return decimalC.String()
}

func Float64AddToFloat(fa, fb float64) float64 {
	decimalA := decimal.NewFromFloat(fa)
	decimalB := decimal.NewFromFloat(fb)
	decimalC := decimalA.Add(decimalB)
	res, _ := decimalC.Float64()
	return res
}

func Float64SubToFloat(fa, fb float64) float64 {
	decimalA := decimal.NewFromFloat(fa)
	decimalB := decimal.NewFromFloat(fb)
	decimalC := decimalA.Sub(decimalB)
	res, _ := decimalC.Float64()
	return res
}

func Float64MulToFloat(fa, fb float64) float64 {
	decimalA := decimal.NewFromFloat(fa)
	decimalB := decimal.NewFromFloat(fb)
	decimalC := decimalA.Mul(decimalB)
	res, _ := decimalC.Float64()
	return res
}
func Float64DivToFloat(fa, fb float64) float64 {
	decimalA := decimal.NewFromFloat(fa)
	decimalB := decimal.NewFromFloat(fb)
	decimalC := decimalA.Div(decimalB)
	res, _ := decimalC.Float64()
	return res
}

func Float64SubToFloat64s(args ...float64) float64 {
	total := decimal.NewFromFloat(0)
	for _, arg := range args {
		total = total.Sub(decimal.NewFromFloat(arg))
	}
	res, _ := total.Float64()
	return res
}

func StringAddToString(sa, sb string) (string, error) {
	decimalA, err := decimal.NewFromString(sa)
	if err != nil {
		return "", err
	}
	decimalB, err := decimal.NewFromString(sb)
	if err != nil {
		return "", err
	}
	decimalC := decimalA.Add(decimalB)
	return decimalC.String(), nil
}

func StringSubToString(sa, sb string) (string, error) {
	decimalA, err := decimal.NewFromString(sa)
	if err != nil {
		return "", err
	}
	decimalB, err := decimal.NewFromString(sb)
	if err != nil {
		return "", err
	}
	decimalC := decimalA.Sub(decimalB)
	return decimalC.String(), nil
}

func StringSubStrings(args ...string) (string, error) {
	total := decimal.NewFromFloat(0)
	for _, arg := range args {
		decimalA, err := decimal.NewFromString(arg)
		if err != nil {
			return "", err
		}
		total = total.Sub(decimalA)
	}
	return total.String(), nil
}

func StringMulToString(sa, sb string) (string, error) {
	decimalA, err := decimal.NewFromString(sa)
	if err != nil {
		return "", err
	}
	decimalB, err := decimal.NewFromString(sb)
	if err != nil {
		return "", err
	}
	decimalC := decimalA.Mul(decimalB)
	return decimalC.String(), nil
}
func StringDivToString(sa, sb string) (string, error) {
	decimalA, err := decimal.NewFromString(sa)
	if err != nil {
		return "", err
	}
	decimalB, err := decimal.NewFromString(sb)
	if err != nil {
		return "", err
	}
	decimalC := decimalA.Div(decimalB)
	return decimalC.String(), nil
}

func StringToFloat64(s string) float64 {
	decimalA, err := decimal.NewFromString(s)
	if err != nil {
		return 0
	}
	res, _ := decimalA.Float64()
	return res
}

func Float64ToString(fa float64) string {
	decimalA := decimal.NewFromFloat(fa)
	return decimalA.String()
}

func HashPassword(password string) (string, error) {
	// cost 成本因子, 越高越慢越安全
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// hash 加密后密码, password明码
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ToJsonString(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(b)
}
