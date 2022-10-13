package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Time struct {
	Hour   int
	Minute int
}

func ChangeToStandartTime(time interface{}) string {
	switch time.(type) {
	case string:
		return fromString(time)
	case []int:
		return fromSliceInt(time)
	case map[string]int:
		return fromMap(time)
	case Time:
		return fromStruct(time)
	default:
		return "Invalid input"
	}
}

func fromString(t interface{}) string {
	a := fmt.Sprintf("%v", t)
	s := strings.Split(a, ":")
	res := ""
	if s[len(s)-1] == "" || len(s) == 1 {
		return "Invalid input"
	} else {
		hour, _ := strconv.Atoi(s[0])
		minute, _ := strconv.Atoi(s[1])
		res = count(hour, minute)
	}

	return res
}

func fromSliceInt(t interface{}) string {
	conv := reflect.ValueOf(t)
	a := conv.Interface().([]int)
	res := " "
	if len(a) == 1 {
		return "Invalid input"
	} else {
		res = count(a[0], a[1])
	}
	return res
}

func fromMap(t interface{}) string {
	conv := reflect.ValueOf(t)
	a := conv.Interface().(map[string]int)
	values := []int{}
	res := " "
	for key, value := range a {
		if key == "second" {
			res = "Invalid input"
			break
		} else {
			values = append(values, value)
		}

	}
	if len(values) <= 1 {
		return "Invalid input"
	} else {
		res = count(values[0], values[1])
	}
	return res
}

func fromStruct(t interface{}) string {
	conv := reflect.ValueOf(t)
	a := conv.Interface().(Time)
	res := count(a.Hour, a.Minute)
	return res
}

func count(hour, minute int) string {
	if hour >= 24 {
		r := hour - 24
		if minute <= 59 {
			return fmt.Sprintf("0%d:%d", r, minute)
		}
	} else if hour == 12 {
		if minute < 10 {
			return fmt.Sprintf("%d:0%d PM", hour, minute)
		} else if minute >= 10 && minute < 59 {
			return fmt.Sprintf("%d:%d PM", hour, minute)
		}
	} else if hour > 12 {
		r := hour - 12
		if r < 10 {
			if minute < 10 {
				return fmt.Sprintf("0%d:0%d PM", r, minute)
			} else {
				return fmt.Sprintf("0%d:%d PM", r, minute)
			}
		} else if r >= 10 && minute >= 10 && minute < 59 {
			return fmt.Sprintf("%d:%d PM", r, minute)
		}
	} else if hour < 12 {
		if minute < 10 {
			return fmt.Sprintf("0%d:0%d AM", hour, minute)
		} else if minute >= 10 && minute < 59 {
			return fmt.Sprintf("0%d:%d AM", hour, minute)
		}
	}
	return ""
}

func main() {
	fmt.Println(ChangeToStandartTime("20:11"))
	fmt.Println(ChangeToStandartTime([]int{16, 0}))
	fmt.Println(ChangeToStandartTime(map[string]int{"hour": 23, "minute": 11}))
	fmt.Println(ChangeToStandartTime(map[string]int{"hour": 14, "minute": 0}))
	fmt.Println(ChangeToStandartTime(Time{20, 11}))
}
