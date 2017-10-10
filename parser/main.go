package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"time"

	item "github.com/techsysfr/paastek-item"
)

type ret struct {
	V   item.AWSBillingItem
	Err error
}

// parse a csv file and return an array of resources
func parse(r io.Reader) chan ret {
	c := make(chan ret, 0)
	go func() {
		defer close(c)
		rd := csv.NewReader(r)
		var header []string
		header, err := rd.Read()
		if err != nil {
			c <- ret{item.AWSBillingItem{}, err}
		}

		e := item.AWSBillingItem{}
		et := reflect.TypeOf(e)
		var headers = make(map[string]int, et.NumField())
		for i := 0; i < et.NumField(); i++ {
			headers[et.Field(i).Name] = func(element string, array []string) int {
				for k, v := range array {
					if v == element {
						return k
					}
				}
				return -1
			}(et.Field(i).Tag.Get("csv"), header)
		}
		for {
			var e = item.AWSBillingItem{}
			record, err := rd.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				c <- ret{item.AWSBillingItem{}, err}
			}
			for h, i := range headers {
				if i == -1 {
					continue
				}
				elem := reflect.ValueOf(&e).Elem()
				field := elem.FieldByName(h)
				if field.CanSet() {
					switch field.Type().Name() {
					case "float64":
						a, _ := strconv.ParseFloat(record[i], 64)
						field.Set(reflect.ValueOf(a))
					case "Time":
						a, _ := time.Parse("2006-01-02T00:00:00Z", record[i])
						field.Set(reflect.ValueOf(a))
					default:
						field.Set(reflect.ValueOf(record[i]))
					}
				}
			}
			c <- ret{e, nil}
		}
	}()
	return c
}

func main() {
	c := parse(os.Stdin)
	for val := range c {
		fmt.Println(val.V.ProductFamily)
	}
}
