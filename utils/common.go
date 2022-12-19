package utils

import (
	"fmt"
	"math"
	"reflect"
	"strings"
	"time"
)

func ToPointer[T any](in T) *T {
	return &in
}

func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func GetValueAndCallType(in string) (string, string) {
	callType := strings.Split(in, "-")
	if len(callType) == 2 {
		return callType[0], callType[1]
	}
	return "", ""
}

func Assign(maps ...map[int]string) map[int]string {
	out := map[int]string{}

	for _, m := range maps {
		for k, v := range m {
			out[k] = v
		}
	}

	return out
}

func ChunkSlice[T any](items []T, chunkSize int) (chunks [][]T) {
	var _chunks = make([][]T, 0, (len(items)/chunkSize)+1)
	for chunkSize < len(items) {
		items, _chunks = items[chunkSize:], append(_chunks, items[0:chunkSize:chunkSize])
	}
	return append(_chunks, items)
}

func GetTimeAt(dayAgo int) time.Time {
	return time.Now().Add(-1 * time.Hour * 24 * time.Duration(dayAgo))
}

func GetAllDatesBetween(from, to time.Time) []string {
	var listDate []string
	days := to.Sub(from).Hours() / 24
	for i := 0; i <= int(days); i++ {
		year, month, day := from.Add(time.Hour * 24 * time.Duration(i)).Date()
		date := fmt.Sprintf("%v-%v-%v", year, int(month), day)
		listDate = append(listDate, date)
	}
	return listDate
}
func GetDateFromTime(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%v-%v-%v", year, int(month), day)
}

func CheckValidFetchTime(t int64, validDay time.Duration) bool {
	return time.Now().Sub(time.Unix(t, 0)) < validDay
}

func TypeCast(src []interface{}, dst interface{}) interface{} {
	dstType := reflect.TypeOf(dst)
	dstValue := reflect.New(dstType)
	for i := 0; i < dstType.NumField(); i++ {
		dstValue.Elem().FieldByName(dstType.Field(i).Name).Set(reflect.ValueOf(src[i]))
	}
	return dstValue.Interface()
}

func StringDefault(input string, byDefault string) string {
	if len(input) == 0 {
		return byDefault
	}

	return input
}
