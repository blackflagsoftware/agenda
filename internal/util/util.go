package util

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"

	"github.com/google/uuid"
)

type (
	Param struct {
		Limit  string
		Number int
		Size   int
		Sort   string
		Search []ParamSearch `json:"search"`
	}

	ParamSearch struct {
		Column  string      `json:"column"`
		Compare string      `json:"compare"`
		Value   interface{} `json:"value"`
	}
)

func GenerateUUID() string {
	return uuid.New().String()
}

func GetTypeCount(i interface{}) int {
	switch reflect.ValueOf(i).Kind() {
	case reflect.Map:
		return reflect.ValueOf(i).Len()
	case reflect.Array:
		return reflect.ValueOf(i).Len()
	case reflect.Slice:
		return reflect.ValueOf(i).Len()
	default:
		return 1
	}
}

func ValidJson(jsonValue json.RawMessage) bool {
	bValue, err := jsonValue.MarshalJSON()
	if err != nil {
		return false
	}
	check := make(map[string]interface{}, 0)
	if errCheck := json.Unmarshal(bValue, &check); errCheck != nil {
		return false
	}
	return true
}

func (p *Param) CalculateParam(primarySort string, availableSort map[string]string) (err error) {
	// calculate the limit
	if p.Size > 0 {
		if p.Number == 0 {
			// should not be empty, default to first page
			p.Number = 1
		}
		offset := p.Number - 1
		offset *= p.Size
		p.Limit = fmt.Sprintf("LIMIT %d, %d", offset, p.Size)
	}
	// calculate the sort
	if primarySort == "" {
		return
	}
	if p.Sort == "" {
		p.Sort = primarySort
	}
	sorted := []string{}
	sortParts := strings.Split(p.Sort, ":")
	for _, s := range sortParts {
		direction := "ASC"
		name := s
		if string(name[0]) == "-" {
			direction = "DESC"
			name = string(name[1:])
		}
		if _, ok := availableSort[name]; !ok {
			// if the name is not in the available sort list, you could return and error here
			continue
		}
		sorted = append(sorted, fmt.Sprintf("%s %s", availableSort[name], direction))
	}
	p.Sort = strings.Join(sorted, ", ")
	return
}

func CopyFileWithOverride(srcStr, dstStr string) error {
	srcStat, err := os.Stat(srcStr)
	if err != nil {
		return err
	}
	if !srcStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file.", srcStr)
	}
	src, err := os.Open(srcStr)
	if err != nil {
		return err
	}
	defer src.Close()
	dst, err := os.Create(dstStr)
	if err != nil {
		return err
	}

	buf := make([]byte, 1024)
	for {
		n, err := src.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		if _, err := dst.Write(buf[:n]); err != nil {
			return fmt.Errorf("error in writing: %s", err)
		}
	}
	return err
}
