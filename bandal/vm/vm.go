package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"strconv"
	"reflect"
	"regexp"
	"math"
)

var errCantReadFile = errors.New("Cannot Read Current File")
var errRuntimeError = errors.New("RunTime Error")
var fileData string = ""
var errNothingToPop = errors.New("Stack has been popped without anything in it")
var errStackEmpty = errors.New("Stack is Empty")
var DEBUG bool = true

// Struct Data Type
type Stack struct {
	stack []interface{}
}

// main
func main() {
	ReadFile()
	eval()
}

// eval
func eval() error {
	var stack Stack
	var PTR int = 0
	const MAX = 65535 //2^16 - 1\
	var err error = nil
	content := strings.Split(fileData, "\n")

	for _, str := range content {
		if PTR > MAX {
			return errRuntimeError
		}

		op := strings.Fields(string(str))
		if len(op) == 0 {
			continue
		}
		//operand = make([]string, 10)

		switch op[0] {
		case "add":
			last, _ := stack.Pop()
			first, _ := stack.Pop()
			if reflect.TypeOf(first) == reflect.TypeOf(last) {
				switch reflect.TypeOf(first).Kind() {
				case reflect.Slice:
					stack.Push(append(first.([]interface{}), last.([]interface{})...))
				default:
					switch reflect.TypeOf(first) {
					case reflect.TypeOf(int8(1)): stack.Push(first.(int8) + last.(int8))
					case reflect.TypeOf(int16(1)): stack.Push(first.(int16) + last.(int16))
					case reflect.TypeOf(int32(1)): stack.Push(first.(int32) + last.(int32))
					case reflect.TypeOf(int64(1)): stack.Push(first.(int64) + last.(int64))
					case reflect.TypeOf(uint8(1)): stack.Push(first.(uint8) + last.(uint8))
					case reflect.TypeOf(uint16(1)): stack.Push(first.(uint16) + last.(uint16))
					case reflect.TypeOf(uint32(1)): stack.Push(first.(uint32) + last.(uint32))
					case reflect.TypeOf(uint64(1)): stack.Push(first.(uint64) + last.(uint64))
					case reflect.TypeOf(float32(1)): stack.Push(first.(float32) + last.(float32))
					case reflect.TypeOf(float64(1)): stack.Push(first.(float64) + last.(float64))
					case reflect.TypeOf(string("")):  stack.Push(first.(string) + last.(string))
					}
				}
			} else {
				switch reflect.TypeOf(first) {
				case reflect.TypeOf(int8(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) + last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) + last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						if reflect.TypeOf(first.(int8)) != reflect.TypeOf(last.([]interface{})[0]) {
							fmt.Println(errors.New("TypeError: cannot eval '+' between different type(s)"))
						}
						stack.Push(append([]interface{}{first.(int8)}, last.([]interface{})...))
					}
				case reflect.TypeOf(int16(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) + last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) + last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						if reflect.TypeOf(first.(int16)) != reflect.TypeOf(last.([]interface{})[0]) {
							fmt.Println(errors.New("TypeError: cannot eval '+' between different type(s)"))
						}
						stack.Push(append([]interface{}{first.(int16)}, last.([]interface{})...))
					}
				case reflect.TypeOf(int32(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) + last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) + last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						if reflect.TypeOf(first.(int32)) != reflect.TypeOf(last.([]interface{})[0]) {
							fmt.Println(errors.New("TypeError: cannot eval '+' between different type(s)"))
						}
						stack.Push(append([]interface{}{first.(int32)}, last.([]interface{})...))
					}
				case reflect.TypeOf(int64(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) + last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) + last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						if reflect.TypeOf(first.(int64)) != reflect.TypeOf(last.([]interface{})[0]) {
							fmt.Println(errors.New("TypeError: cannot eval '+' between different type(s)"))
						}
						stack.Push(append([]interface{}{first.(int64)}, last.([]interface{})...))
					}
				case reflect.TypeOf(uint8(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) + last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) + last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						if reflect.TypeOf(first.(uint8)) != reflect.TypeOf(last.([]interface{})[0]) {
							fmt.Println(errors.New("TypeError: cannot eval '+' between different type(s)"))
						}
						stack.Push(append([]interface{}{first.(uint8)}, last.([]interface{})...))
					}
				case reflect.TypeOf(uint16(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) + last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) + last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						if reflect.TypeOf(first.(uint16)) != reflect.TypeOf(last.([]interface{})[0]) {
							fmt.Println(errors.New("TypeError: cannot eval '+' between different type(s)"))
						}
						stack.Push(append([]interface{}{first.(uint16)}, last.([]interface{})...))
					}
				case reflect.TypeOf(uint32(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) + last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) + last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						if reflect.TypeOf(first.(uint32)) != reflect.TypeOf(last.([]interface{})[0]) {
							fmt.Println(errors.New("TypeError: cannot eval '+' between different type(s)"))
						}
						stack.Push(append([]interface{}{first.(uint32)}, last.([]interface{})...))
					}
				case reflect.TypeOf(uint64(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) + last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) + last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						if reflect.TypeOf(first.(uint64)) != reflect.TypeOf(last.([]interface{})[0]) {
							fmt.Println(errors.New("TypeError: cannot eval '+' between different type(s)"))
						}
						stack.Push(append([]interface{}{first.(uint64)}, last.([]interface{})...))
					}
				case reflect.TypeOf(float32(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) + last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						if reflect.TypeOf(first.(float32)) != reflect.TypeOf(last.([]interface{})[0]) {
							fmt.Println(errors.New("TypeError: cannot eval '+' between different type(s)"))
						}
						stack.Push(append([]interface{}{first.(float32)}, last.([]interface{})...))
					}
				case reflect.TypeOf(float64(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float64) + last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						if reflect.TypeOf(first.(float64)) != reflect.TypeOf(last.([]interface{})[0]) {
							fmt.Println(errors.New("TypeError: cannot eval '+' between different type(s)"))
						}
						stack.Push(append([]interface{}{first.(float64)}, last.([]interface{})...))
					}
				case reflect.TypeOf(string("")):
					if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						if reflect.TypeOf(first.(string)) != reflect.TypeOf(last.([]interface{})[0]) {
							fmt.Println(errors.New("TypeError: cannot eval '+' between different type(s)"))
						}
						stack.Push(append([]interface{}{first.(string)}, last.([]interface{})...))
					}
				case reflect.TypeOf([]interface{}{}):
					if len(first.([]interface{})) > 0 {
						if reflect.TypeOf(first.([]interface{})[0]) != reflect.TypeOf(last) {
							fmt.Println(errors.New("TypeError: cannot eval '+' between different type(s)"))
						} 
					}
					switch reflect.TypeOf(last) {
					case reflect.TypeOf(int8(1)): stack.Push(append(first.([]interface{}), last.(int8)))
					case reflect.TypeOf(int16(1)): stack.Push(append(first.([]interface{}), last.(int16)))
					case reflect.TypeOf(int32(1)): stack.Push(append(first.([]interface{}), last.(int32)))
					case reflect.TypeOf(int64(1)): stack.Push(append(first.([]interface{}), last.(int64)))
					case reflect.TypeOf(uint8(1)): stack.Push(append(first.([]interface{}), last.(uint8)))
					case reflect.TypeOf(uint16(1)): stack.Push(append(first.([]interface{}), last.(uint16)))
					case reflect.TypeOf(uint32(1)): stack.Push(append(first.([]interface{}), last.(uint32)))
					case reflect.TypeOf(uint64(1)): stack.Push(append(first.([]interface{}), last.(uint64)))
					case reflect.TypeOf(float32(1)): stack.Push(append(first.([]interface{}), last.(float32)))
					case reflect.TypeOf(float64(1)): stack.Push(append(first.([]interface{}), last.(float64)))
					case reflect.TypeOf(string("")):  stack.Push(append(first.([]interface{}), last.(string)))
					}
				}
			}
		case "sub":
			last, _ := stack.Pop()
			first, _ := stack.Pop()
			if reflect.TypeOf(first) == reflect.TypeOf(last) {
				switch reflect.TypeOf(first).Kind() {
				case reflect.Slice:
					fmt.Println(errors.New("TypeError: unsupported operand type(s) for -: 'array' and 'array'"))
				default:
					switch reflect.TypeOf(first) {
					case reflect.TypeOf(int8(1)): stack.Push(first.(int8) - last.(int8))
					case reflect.TypeOf(int16(1)): stack.Push(first.(int16) - last.(int16))
					case reflect.TypeOf(int32(1)): stack.Push(first.(int32) - last.(int32))
					case reflect.TypeOf(int64(1)): stack.Push(first.(int64) - last.(int64))
					case reflect.TypeOf(uint8(1)): stack.Push(first.(uint8) - last.(uint8))
					case reflect.TypeOf(uint16(1)): stack.Push(first.(uint16) - last.(uint16))
					case reflect.TypeOf(uint32(1)): stack.Push(first.(uint32) - last.(uint32))
					case reflect.TypeOf(uint64(1)): stack.Push(first.(uint64) - last.(uint64))
					case reflect.TypeOf(float32(1)): stack.Push(first.(float32) - last.(float32))
					case reflect.TypeOf(float64(1)): stack.Push(first.(float64) - last.(float64))
					}
				}
			} else {
				switch reflect.TypeOf(first) {
				case reflect.TypeOf(int8(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) - last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) - last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for -: 'int8' and 'array'"))
					}
				case reflect.TypeOf(int16(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) - last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) - last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for -: 'int16' and 'array'"))
					}
				case reflect.TypeOf(int32(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) - last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) - last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for -: 'int32' and 'array'"))
					}
				case reflect.TypeOf(int64(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) - last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) - last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for -: 'int64' and 'array'"))
					}
				case reflect.TypeOf(uint8(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) - last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) - last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for -: 'uint8' and 'array'"))
					}
				case reflect.TypeOf(uint16(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) - last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) - last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for -: 'uint16' and 'array'"))
					}
				case reflect.TypeOf(uint32(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) - last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) - last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for -: 'uint32' and 'array'"))
					}
				case reflect.TypeOf(uint64(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) - last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) - last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for -: 'uint64' and 'array'"))
					}
				case reflect.TypeOf(float32(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) - last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for -: 'float32' and 'array'"))
					}
				case reflect.TypeOf(float64(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float64) - last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for -: 'float64' and 'array'"))
					}
				case reflect.TypeOf(string("")):
					if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for -: 'string' and 'array'"))
					}
				case reflect.TypeOf([]interface{}{}):
					if len(first.([]interface{})) > 0 {
						if reflect.TypeOf(first.([]interface{})[0]) != reflect.TypeOf(last) {
							fmt.Println(errors.New("TypeError: cannot eval '-' between different type(s)"))
						} 
					}
					switch reflect.TypeOf(last) {
					case reflect.TypeOf(int8(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for -: 'array' and 'int8'"))
					case reflect.TypeOf(int16(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for -: 'array' and 'int16'"))
					case reflect.TypeOf(int32(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for -: 'array' and 'int32'"))
					case reflect.TypeOf(int64(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for -: 'array' and 'int64'"))
					case reflect.TypeOf(uint8(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for -: 'array' and 'uint8'"))
					case reflect.TypeOf(uint16(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for -: 'array' and 'uint16'"))
					case reflect.TypeOf(uint32(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for -: 'array' and 'uint32'"))
					case reflect.TypeOf(uint64(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for -: 'array' and 'uint64'"))
					case reflect.TypeOf(float32(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for -: 'array' and 'float32'"))
					case reflect.TypeOf(float64(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for -: 'array' and 'float64'"))
					case reflect.TypeOf(string("")):  fmt.Println(errors.New("TypeError: unsupported operand type(s) for -: 'array' and 'string'"))
					}
				}
			}
		case "mul":
			last, _ := stack.Pop()
			first, _ := stack.Pop()
			if reflect.TypeOf(first) == reflect.TypeOf(last) {
				switch reflect.TypeOf(first).Kind() {
				case reflect.Slice:
					fmt.Println(errors.New("TypeError: unsupported operand type(s) for *: 'array' and 'array'"))
				default:
					switch reflect.TypeOf(first) {
					case reflect.TypeOf(int8(1)): stack.Push(first.(int8) * last.(int8))
					case reflect.TypeOf(int16(1)): stack.Push(first.(int16) * last.(int16))
					case reflect.TypeOf(int32(1)): stack.Push(first.(int32) * last.(int32))
					case reflect.TypeOf(int64(1)): stack.Push(first.(int64) * last.(int64))
					case reflect.TypeOf(uint8(1)): stack.Push(first.(uint8) * last.(uint8))
					case reflect.TypeOf(uint16(1)): stack.Push(first.(uint16) * last.(uint16))
					case reflect.TypeOf(uint32(1)): stack.Push(first.(uint32) * last.(uint32))
					case reflect.TypeOf(uint64(1)): stack.Push(first.(uint64) * last.(uint64))
					case reflect.TypeOf(float32(1)): stack.Push(first.(float32) * last.(float32))
					case reflect.TypeOf(float64(1)): stack.Push(first.(float64) * last.(float64))
					}
				}
			} else {
				switch reflect.TypeOf(first) {
				case reflect.TypeOf(int8(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) * last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) * last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for *: 'int8' and 'array'"))
					}
				case reflect.TypeOf(int16(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) * last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) * last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for *: 'int16' and 'array'"))
					}
				case reflect.TypeOf(int32(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) * last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) * last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for *: 'int32' and 'array'"))
					}
				case reflect.TypeOf(int64(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) * last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) * last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for *: 'int64' and 'array'"))
					}
				case reflect.TypeOf(uint8(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) * last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) * last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for *: 'uint8' and 'array'"))
					}
				case reflect.TypeOf(uint16(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) * last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) * last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for *: 'uint16' and 'array'"))
					}
				case reflect.TypeOf(uint32(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) * last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) * last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for *: 'uint32' and 'array'"))
					}
				case reflect.TypeOf(uint64(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) * last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) * last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for *: 'uint64' and 'array'"))
					}
				case reflect.TypeOf(float32(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) * last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for *: 'float32' and 'array'"))
					}
				case reflect.TypeOf(float64(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float64) * last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for *: 'float64' and 'array'"))
					}
				case reflect.TypeOf(string("")):
					if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for *: 'string' and 'array'"))
					}
				case reflect.TypeOf([]interface{}{}):
					if len(first.([]interface{})) > 0 {
						if reflect.TypeOf(first.([]interface{})[0]) != reflect.TypeOf(last) {
							fmt.Println(errors.New("TypeError: cannot eval '*' between different type(s)"))
						} 
					}
					switch reflect.TypeOf(last) {
					case reflect.TypeOf(int8(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for *: 'array' and 'int8'"))
					case reflect.TypeOf(int16(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for *: 'array' and 'int16'"))
					case reflect.TypeOf(int32(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for *: 'array' and 'int32'"))
					case reflect.TypeOf(int64(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for *: 'array' and 'int64'"))
					case reflect.TypeOf(uint8(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for *: 'array' and 'uint8'"))
					case reflect.TypeOf(uint16(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for *: 'array' and 'uint16'"))
					case reflect.TypeOf(uint32(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for *: 'array' and 'uint32'"))
					case reflect.TypeOf(uint64(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for *: 'array' and 'uint64'"))
					case reflect.TypeOf(float32(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for *: 'array' and 'float32'"))
					case reflect.TypeOf(float64(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for *: 'array' and 'float64'"))
					case reflect.TypeOf(string("")):  fmt.Println(errors.New("TypeError: unsupported operand type(s) for *: 'array' and 'string'"))
					}
				}
			}
		case "div":
			last, _ := stack.Pop()
			first, _ := stack.Pop()
			if reflect.TypeOf(first) == reflect.TypeOf(last) {
				switch reflect.TypeOf(first).Kind() {
				case reflect.Slice:
					fmt.Println(errors.New("TypeError: unsupported operand type(s) for /: 'array' and 'array'"))
				default:
					switch reflect.TypeOf(first) {
					case reflect.TypeOf(int8(1)):
						if first.(int8) < last.(int8) {
							stack.Push(float64(first.(int8)) / float64(last.(int8)))
						}
					case reflect.TypeOf(int16(1)):
						if first.(int16) < last.(int16) {
							stack.Push(float64(first.(int16)) / float64(last.(int16)))
						}
					case reflect.TypeOf(int32(1)):
						if first.(int32) < last.(int32) {
							stack.Push(float64(first.(int32)) / float64(last.(int32)))
						}
					case reflect.TypeOf(int64(1)):
						if first.(int64) < last.(int64) {
							stack.Push(float64(first.(int64)) / float64(last.(int64)))
						}
					case reflect.TypeOf(uint8(1)):
						if first.(uint8) < last.(uint8) {
							stack.Push(float64(first.(uint8)) / float64(last.(uint8)))
						}
					case reflect.TypeOf(uint16(1)):
						if first.(uint16) < last.(uint16) {
							stack.Push(float64(first.(uint16)) / float64(last.(uint16)))
						}
					case reflect.TypeOf(uint32(1)):
						if first.(uint32) < last.(uint32) {
							stack.Push(float64(first.(uint32)) / float64(last.(uint32)))
						}
					case reflect.TypeOf(uint64(1)):
						if first.(uint64) < last.(uint64) {
							stack.Push(float64(first.(uint64)) / float64(last.(uint64)))
						}
					case reflect.TypeOf(float32(1)): stack.Push(first.(float32) / last.(float32))
					case reflect.TypeOf(float64(1)): stack.Push(first.(float64) / last.(float64))
					}
				}
			} else {
				switch reflect.TypeOf(first) {
				case reflect.TypeOf(int8(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) / last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) / last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for /: 'int8' and 'array'"))
					}
				case reflect.TypeOf(int16(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) / last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) / last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for /: 'int16' and 'array'"))
					}
				case reflect.TypeOf(int32(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) / last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) / last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for /: 'int32' and 'array'"))
					}
				case reflect.TypeOf(int64(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) / last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) / last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for /: 'int64' and 'array'"))
					}
				case reflect.TypeOf(uint8(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) / last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) / last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for /: 'uint8' and 'array'"))
					}
				case reflect.TypeOf(uint16(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) / last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) / last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for /: 'uint16' and 'array'"))
					}
				case reflect.TypeOf(uint32(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) / last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) / last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for /: 'uint32' and 'array'"))
					}
				case reflect.TypeOf(uint64(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) / last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) / last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for /: 'uint64' and 'array'"))
					}
				case reflect.TypeOf(float32(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) / last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for /: 'float32' and 'array'"))
					}
				case reflect.TypeOf(float64(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float64) / last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for /: 'float64' and 'array'"))
					}
				case reflect.TypeOf(string("")):
					if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for /: 'string' and 'array'"))
					}
				case reflect.TypeOf([]interface{}{}):
					if len(first.([]interface{})) > 0 {
						if reflect.TypeOf(first.([]interface{})[0]) != reflect.TypeOf(last) {
							fmt.Println(errors.New("TypeError: cannot eval '/' between different type(s)"))
						} 
					}
					switch reflect.TypeOf(last) {
					case reflect.TypeOf(int8(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for /: 'array' and 'int8'"))
					case reflect.TypeOf(int16(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for /: 'array' and 'int16'"))
					case reflect.TypeOf(int32(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for /: 'array' and 'int32'"))
					case reflect.TypeOf(int64(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for /: 'array' and 'int64'"))
					case reflect.TypeOf(uint8(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for /: 'array' and 'uint8'"))
					case reflect.TypeOf(uint16(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for /: 'array' and 'uint16'"))
					case reflect.TypeOf(uint32(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for /: 'array' and 'uint32'"))
					case reflect.TypeOf(uint64(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for /: 'array' and 'uint64'"))
					case reflect.TypeOf(float32(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for /: 'array' and 'float32'"))
					case reflect.TypeOf(float64(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for /: 'array' and 'float64'"))
					case reflect.TypeOf(string("")):  fmt.Println(errors.New("TypeError: unsupported operand type(s) for /: 'array' and 'string'"))
					}
				}
			}
		case "push":
			rawValue := strings.Join(op[1:len(op)], " ")
			flag := false
			for _, s := range rawValue {
				if string(s) == `"` { //string
					stack.Push(regexp.MustCompile(`"`).ReplaceAllString(rawValue, ``))
					flag = true
					break
				} else if string(s) == `[` { //array
					newReplacer := strings.NewReplacer("[", "", "]", "")
					replaced := newReplacer.Replace(rawValue)
					arr := strings.Split(replaced, ", ")
					Value := []interface{}{}
					for _, elem := range arr {
						for _, elemStr := range elem {
							if string(elemStr) == `"` { //string
								Value = append(Value, regexp.MustCompile(`"`).ReplaceAllString(elem, ``))
								break
							} else if string(elemStr) == `.` { //float
								tmp, _ := strconv.ParseFloat(elem, 64)
								Value = append(Value, tmp)
								break
							} else { //int
								tmp, _ := strconv.ParseInt(elem, 10, 32) //always int64
								if -2147483648 <= tmp <= 2147483647 {
									Value = append(Value, int32(tmp)) //cvt int32
								} else {
									Value = append(Value, tmp) //int64
									break
								}
							}
						}
					}
					stack.Push(Value)
					flag = true
					break
				} else if string(s) == `.` { //float
					tmp, _ := strconv.ParseFloat(rawValue, 64)
					stack.Push(tmp)
					flag = true
					break
				}
			}
			if flag == false {
				tmp, _ := strconv.ParseInt(rawValue, 10, 32) //always int64
				if -2147483648 <= tmp <= 2147483647 {
					stack.Push(int32(tmp)) //cvt int32
				} else {
					stack.Push(tmp) //int64
				}
			}
		case "rem":
			last, _ := stack.Pop()
			first, _ := stack.Pop()
			if reflect.TypeOf(first) == reflect.TypeOf(last) {
				switch reflect.TypeOf(first).Kind() {
				case reflect.Slice:
					fmt.Println(errors.New("TypeError: unsupported operand type(s) for %: 'array' and 'array'"))
				default:
					switch reflect.TypeOf(first) {
					case reflect.TypeOf(int8(1)): stack.Push(first.(int8) % last.(int8))
					case reflect.TypeOf(int16(1)): stack.Push(first.(int16) % last.(int16))
					case reflect.TypeOf(int32(1)): stack.Push(first.(int32) % last.(int32))
					case reflect.TypeOf(int64(1)): stack.Push(first.(int64) % last.(int64))
					case reflect.TypeOf(uint8(1)): stack.Push(first.(uint8) % last.(uint8))
					case reflect.TypeOf(uint16(1)): stack.Push(first.(uint16) % last.(uint16))
					case reflect.TypeOf(uint32(1)): stack.Push(first.(uint32) % last.(uint32))
					case reflect.TypeOf(uint64(1)): stack.Push(first.(uint64) % last.(uint64))
					case reflect.TypeOf(float32(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for %: 'float32' and 'float32'"))
					case reflect.TypeOf(float64(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for %: 'float64' and 'float64'"))
					}
				}
			} else {
				switch reflect.TypeOf(first) {
				case reflect.TypeOf(int8(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for %: 'float32' and 'float32'"))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for %: 'float64' and 'float64'"))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for %: 'int8' and 'array'"))
					}
				case reflect.TypeOf(int16(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for %: 'float32' and 'float32'"))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for %: 'float64' and 'float64'"))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for %: 'int16' and 'array'"))
					}
				case reflect.TypeOf(int32(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for %: 'float32' and 'float32'"))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for %: 'float64' and 'float64'"))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for %: 'int32' and 'array'"))
					}
				case reflect.TypeOf(int64(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for %: 'float32' and 'float32'"))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for %: 'float64' and 'float64'"))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for %: 'int64' and 'array'"))
					}
				case reflect.TypeOf(uint8(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for %: 'float32' and 'float32'"))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for %: 'float64' and 'float64'"))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for %: 'uint8' and 'array'"))
					}
				case reflect.TypeOf(uint16(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for %: 'float32' and 'float32'"))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for %: 'float64' and 'float64'"))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for %: 'uint16' and 'array'"))
					}
				case reflect.TypeOf(uint32(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for %: 'float32' and 'float32'"))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for %: 'float64' and 'float64'"))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for %: 'uint32' and 'array'"))
					}
				case reflect.TypeOf(uint64(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for %: 'float32' and 'float32'"))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for %: 'float64' and 'float64'"))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for %: 'uint64' and 'array'"))
					}
				case reflect.TypeOf(float32(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for %: 'float64' and 'float64'"))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for %: 'float32' and 'array'"))
					}
				case reflect.TypeOf(float64(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for %: 'float64' and 'float64'"))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for %: 'float64' and 'array'"))
					}
				case reflect.TypeOf(string("")):
					if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for %: 'string' and 'array'"))
					}
				case reflect.TypeOf([]interface{}{}):
					if len(first.([]interface{})) > 0 {
						if reflect.TypeOf(first.([]interface{})[0]) != reflect.TypeOf(last) {
							fmt.Println(errors.New("TypeError: cannot eval '%' between different type(s)"))
						} 
					}
					switch reflect.TypeOf(last) {
					case reflect.TypeOf(int8(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for %: 'array' and 'int8'"))
					case reflect.TypeOf(int16(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for %: 'array' and 'int16'"))
					case reflect.TypeOf(int32(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for %: 'array' and 'int32'"))
					case reflect.TypeOf(int64(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for %: 'array' and 'int64'"))
					case reflect.TypeOf(uint8(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for %: 'array' and 'uint8'"))
					case reflect.TypeOf(uint16(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for %: 'array' and 'uint16'"))
					case reflect.TypeOf(uint32(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for %: 'array' and 'uint32'"))
					case reflect.TypeOf(uint64(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for %: 'array' and 'uint64'"))
					case reflect.TypeOf(float32(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for %: 'array' and 'float32'"))
					case reflect.TypeOf(float64(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for %: 'array' and 'float64'"))
					case reflect.TypeOf(string("")):  fmt.Println(errors.New("TypeError: unsupported operand type(s) for %: 'array' and 'string'"))
					}
				}
			}
		case "fix":
			last, _ := stack.Pop()
			first, _ := stack.Pop()
			if reflect.TypeOf(first) == reflect.TypeOf(last) {
				switch reflect.TypeOf(first).Kind() {
				case reflect.Slice:
					fmt.Println(errors.New("TypeError: unsupported operand type(s) for //: 'array' and 'array'"))
				default:
					switch reflect.TypeOf(first) {
					case reflect.TypeOf(int8(1)): stack.Push(math.Floor(float64(first.(int8) / last.(int8))))
					case reflect.TypeOf(int16(1)): stack.Push(math.Floor(float64(first.(int16) / last.(int16))))
					case reflect.TypeOf(int32(1)): stack.Push(math.Floor(float64(first.(int32) / last.(int32))))
					case reflect.TypeOf(int64(1)): stack.Push(math.Floor(float64(first.(int64) / last.(int64))))
					case reflect.TypeOf(uint8(1)): stack.Push(math.Floor(float64(first.(uint8) / last.(uint8))))
					case reflect.TypeOf(uint16(1)): stack.Push(math.Floor(float64(first.(uint16) / last.(uint16))))
					case reflect.TypeOf(uint32(1)): stack.Push(math.Floor(float64(first.(uint32) / last.(uint32))))
					case reflect.TypeOf(uint64(1)): stack.Push(math.Floor(float64(first.(uint64) / last.(uint64))))
					case reflect.TypeOf(float32(1)): stack.Push(math.Floor(float64(first.(float32) / last.(float32))))
					case reflect.TypeOf(float64(1)): stack.Push(math.Floor(first.(float64) / last.(float64)))
					}
				}
			} else {
				switch reflect.TypeOf(first) {
				case reflect.TypeOf(int8(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(math.Floor(float64(first.(float32) / last.(float32))))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(math.Floor(first.(float64) / last.(float64)))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for //: 'int8' and 'array'"))
					}
				case reflect.TypeOf(int16(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(math.Floor(float64(first.(float32) / last.(float32))))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(math.Floor(first.(float64) / last.(float64)))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for //: 'int16' and 'array'"))
					}
				case reflect.TypeOf(int32(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(math.Floor(float64(first.(float32) / last.(float32))))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(math.Floor(first.(float64) / last.(float64)))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for //: 'int32' and 'array'"))
					}
				case reflect.TypeOf(int64(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(math.Floor(float64(first.(float32) / last.(float32))))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(math.Floor(first.(float64) / last.(float64)))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for //: 'int64' and 'array'"))
					}
				case reflect.TypeOf(uint8(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(math.Floor(float64(first.(float32) / last.(float32))))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(math.Floor(first.(float64) / last.(float64)))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for //: 'uint8' and 'array'"))
					}
				case reflect.TypeOf(uint16(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(math.Floor(float64(first.(float32) / last.(float32))))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(math.Floor(first.(float64) / last.(float64)))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for //: 'uint16' and 'array'"))
					}
				case reflect.TypeOf(uint32(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(math.Floor(float64(first.(float32) / last.(float32))))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(math.Floor(first.(float64) / last.(float64)))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for //: 'uint32' and 'array'"))
					}
				case reflect.TypeOf(uint64(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(math.Floor(float64(first.(float32) / last.(float32))))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(math.Floor(first.(float64) / last.(float64)))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for //: 'uint64' and 'array'"))
					}
				case reflect.TypeOf(float32(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(math.Floor(first.(float64) / last.(float64)))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for //: 'float32' and 'array'"))
					}
				case reflect.TypeOf(float64(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(math.Floor(first.(float64) / last.(float64)))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for //: 'float64' and 'array'"))
					}
				case reflect.TypeOf(string("")):
					if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for //: 'string' and 'array'"))
					}
				case reflect.TypeOf([]interface{}{}):
					if len(first.([]interface{})) > 0 {
						if reflect.TypeOf(first.([]interface{})[0]) != reflect.TypeOf(last) {
							fmt.Println(errors.New("TypeError: cannot eval '//' between different type(s)"))
						} 
					}
					switch reflect.TypeOf(last) {
					case reflect.TypeOf(int8(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for //: 'array' and 'int8'"))
					case reflect.TypeOf(int16(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for //: 'array' and 'int16'"))
					case reflect.TypeOf(int32(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for //: 'array' and 'int32'"))
					case reflect.TypeOf(int64(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for //: 'array' and 'int64'"))
					case reflect.TypeOf(uint8(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for //: 'array' and 'uint8'"))
					case reflect.TypeOf(uint16(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for //: 'array' and 'uint16'"))
					case reflect.TypeOf(uint32(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for //: 'array' and 'uint32'"))
					case reflect.TypeOf(uint64(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for //: 'array' and 'uint64'"))
					case reflect.TypeOf(float32(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for //: 'array' and 'float32'"))
					case reflect.TypeOf(float64(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for //: 'array' and 'float64'"))
					case reflect.TypeOf(string("")):  fmt.Println(errors.New("TypeError: unsupported operand type(s) for //: 'array' and 'string'"))
					}
				}
			}
		case "sqr":
			last, _ := stack.Pop()
			first, _ := stack.Pop()
			if reflect.TypeOf(first) == reflect.TypeOf(last) {
				switch reflect.TypeOf(first).Kind() {
				case reflect.Slice:
					fmt.Println(errors.New("TypeError: unsupported operand type(s) for ^: 'array' and 'array'"))
				default:
					switch reflect.TypeOf(first) {
					case reflect.TypeOf(int8(1)): stack.Push(math.Pow(float64(first.(int8)), float64(last.(int8))))
					case reflect.TypeOf(int16(1)): stack.Push(math.Pow(float64(first.(int16)), float64(last.(int16))))
					case reflect.TypeOf(int32(1)): stack.Push(math.Pow(float64(first.(int32)), float64(last.(int32))))
					case reflect.TypeOf(int64(1)): stack.Push(math.Pow(float64(first.(int64)), float64(last.(int64))))
					case reflect.TypeOf(uint8(1)): stack.Push(math.Pow(float64(first.(uint8)), float64(last.(uint8))))
					case reflect.TypeOf(uint16(1)): stack.Push(math.Pow(float64(first.(uint16)), float64(last.(uint16))))
					case reflect.TypeOf(uint32(1)): stack.Push(math.Pow(float64(first.(uint32)), float64(last.(uint32))))
					case reflect.TypeOf(uint64(1)): stack.Push(math.Pow(float64(first.(uint64)), float64(last.(uint64))))
					case reflect.TypeOf(float32(1)): stack.Push(math.Pow(float64(first.(float32)), float64(last.(float32))))
					case reflect.TypeOf(float64(1)): stack.Push(math.Pow(first.(float64), last.(float64)))
					}
				}
			} else {
				switch reflect.TypeOf(first) {
				case reflect.TypeOf(int8(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(math.Pow(float64(first.(float32)), float64(last.(float32))))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(math.Pow(first.(float64), last.(float64)))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for ^: 'int8' and 'array'"))
					}
				case reflect.TypeOf(int16(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(math.Pow(float64(first.(float32)), float64(last.(float32))))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(math.Pow(first.(float64), last.(float64)))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for ^: 'int16' and 'array'"))
					}
				case reflect.TypeOf(int32(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(math.Pow(float64(first.(float32)), float64(last.(float32))))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(math.Pow(first.(float64), last.(float64)))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for ^: 'int32' and 'array'"))
					}
				case reflect.TypeOf(int64(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(math.Pow(float64(first.(float32)), float64(last.(float32))))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(math.Pow(first.(float64), last.(float64)))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for ^: 'int64' and 'array'"))
					}
				case reflect.TypeOf(uint8(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(math.Pow(float64(first.(float32)), float64(last.(float32))))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(math.Pow(first.(float64), last.(float64)))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for ^: 'uint8' and 'array'"))
					}
				case reflect.TypeOf(uint16(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(math.Pow(float64(first.(float32)), float64(last.(float32))))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(math.Pow(first.(float64), last.(float64)))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for ^: 'uint16' and 'array'"))
					}
				case reflect.TypeOf(uint32(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(math.Pow(float64(first.(float32)), float64(last.(float32))))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(math.Pow(first.(float64), last.(float64)))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for ^: 'uint32' and 'array'"))
					}
				case reflect.TypeOf(uint64(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(math.Pow(float64(first.(float32)), float64(last.(float32))))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(math.Pow(first.(float64), last.(float64)))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for ^: 'uint64' and 'array'"))
					}
				case reflect.TypeOf(float32(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(math.Pow(first.(float64), last.(float64)))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for ^: 'float32' and 'array'"))
					}
				case reflect.TypeOf(float64(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(math.Pow(first.(float64), last.(float64)))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for ^: 'float64' and 'array'"))
					}
				case reflect.TypeOf(string("")):
					if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for ^: 'string' and 'array'"))
					}
				case reflect.TypeOf([]interface{}{}):
					if len(first.([]interface{})) > 0 {
						if reflect.TypeOf(first.([]interface{})[0]) != reflect.TypeOf(last) {
							fmt.Println(errors.New("TypeError: cannot eval '^' between different type(s)"))
						} 
					}
					switch reflect.TypeOf(last) {
					case reflect.TypeOf(int8(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for ^: 'array' and 'int8'"))
					case reflect.TypeOf(int16(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for ^: 'array' and 'int16'"))
					case reflect.TypeOf(int32(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for ^: 'array' and 'int32'"))
					case reflect.TypeOf(int64(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for ^: 'array' and 'int64'"))
					case reflect.TypeOf(uint8(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for ^: 'array' and 'uint8'"))
					case reflect.TypeOf(uint16(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for ^: 'array' and 'uint16'"))
					case reflect.TypeOf(uint32(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for ^: 'array' and 'uint32'"))
					case reflect.TypeOf(uint64(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for ^: 'array' and 'uint64'"))
					case reflect.TypeOf(float32(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for ^: 'array' and 'float32'"))
					case reflect.TypeOf(float64(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for ^: 'array' and 'float64'"))
					case reflect.TypeOf(string("")):  fmt.Println(errors.New("TypeError: unsupported operand type(s) for ^: 'array' and 'string'"))
					}
				}
			}
		case "gt":
			last, _ := stack.Pop()
			first, _ := stack.Pop()
			if reflect.TypeOf(first) == reflect.TypeOf(last) {
				switch reflect.TypeOf(first).Kind() {
				case reflect.Slice:
					fmt.Println(errors.New("TypeError: unsupported operand type(s) for >: 'array' and 'array'"))
				default:
					switch reflect.TypeOf(first) {
					case reflect.TypeOf(int8(1)): stack.Push(first.(int8) > last.(int8))
					case reflect.TypeOf(int16(1)): stack.Push(first.(int16) > last.(int16))
					case reflect.TypeOf(int32(1)): stack.Push(first.(int32) > last.(int32))
					case reflect.TypeOf(int64(1)): stack.Push(first.(int64) > last.(int64))
					case reflect.TypeOf(uint8(1)): stack.Push(first.(uint8) > last.(uint8))
					case reflect.TypeOf(uint16(1)): stack.Push(first.(uint16) > last.(uint16))
					case reflect.TypeOf(uint32(1)): stack.Push(first.(uint32) > last.(uint32))
					case reflect.TypeOf(uint64(1)): stack.Push(first.(uint64) > last.(uint64))
					case reflect.TypeOf(float32(1)): stack.Push(first.(float32) > last.(float32))
					case reflect.TypeOf(float64(1)): stack.Push(first.(float64) > last.(float64))
					case reflect.TypeOf(""): stack.Push(first.(string) > last.(string))
					}
				}
			} else {
				switch reflect.TypeOf(first) {
				case reflect.TypeOf(int8(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) > last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) > last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for >: 'int8' and 'array'"))
					}
				case reflect.TypeOf(int16(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) > last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) > last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for >: 'int16' and 'array'"))
					}
				case reflect.TypeOf(int32(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) > last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) > last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for >: 'int32' and 'array'"))
					}
				case reflect.TypeOf(int64(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) > last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) > last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for >: 'int64' and 'array'"))
					}
				case reflect.TypeOf(uint8(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) > last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) > last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for >: 'uint8' and 'array'"))
					}
				case reflect.TypeOf(uint16(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) > last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) > last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for >: 'uint16' and 'array'"))
					}
				case reflect.TypeOf(uint32(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) > last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) > last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for >: 'uint32' and 'array'"))
					}
				case reflect.TypeOf(uint64(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) > last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) > last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for >: 'uint64' and 'array'"))
					}
				case reflect.TypeOf(float32(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) > last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for >: 'float32' and 'array'"))
					}
				case reflect.TypeOf(float64(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float64) > last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for >: 'float64' and 'array'"))
					}
				case reflect.TypeOf(string("")):
					if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for >: 'string' and 'array'"))
					}
				case reflect.TypeOf([]interface{}{}):
					if len(first.([]interface{})) > 0 {
						if reflect.TypeOf(first.([]interface{})[0]) != reflect.TypeOf(last) {
							fmt.Println(errors.New("TypeError: cannot eval '>' between different type(s)"))
						} 
					}
					switch reflect.TypeOf(last) {
					case reflect.TypeOf(int8(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for >: 'array' and 'int8'"))
					case reflect.TypeOf(int16(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for >: 'array' and 'int16'"))
					case reflect.TypeOf(int32(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for >: 'array' and 'int32'"))
					case reflect.TypeOf(int64(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for >: 'array' and 'int64'"))
					case reflect.TypeOf(uint8(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for >: 'array' and 'uint8'"))
					case reflect.TypeOf(uint16(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for >: 'array' and 'uint16'"))
					case reflect.TypeOf(uint32(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for >: 'array' and 'uint32'"))
					case reflect.TypeOf(uint64(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for >: 'array' and 'uint64'"))
					case reflect.TypeOf(float32(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for >: 'array' and 'float32'"))
					case reflect.TypeOf(float64(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for >: 'array' and 'float64'"))
					case reflect.TypeOf(string("")):  fmt.Println(errors.New("TypeError: unsupported operand type(s) for >: 'array' and 'string'"))
					}
				}
			}
		case "ge":
			last, _ := stack.Pop()
			first, _ := stack.Pop()
			if reflect.TypeOf(first) == reflect.TypeOf(last) {
				switch reflect.TypeOf(first).Kind() {
				case reflect.Slice:
					fmt.Println(errors.New("TypeError: unsupported operand type(s) for >=: 'array' and 'array'"))
				default:
					switch reflect.TypeOf(first) {
					case reflect.TypeOf(int8(1)): stack.Push(first.(int8) >= last.(int8))
					case reflect.TypeOf(int16(1)): stack.Push(first.(int16) >= last.(int16))
					case reflect.TypeOf(int32(1)): stack.Push(first.(int32) >= last.(int32))
					case reflect.TypeOf(int64(1)): stack.Push(first.(int64) >= last.(int64))
					case reflect.TypeOf(uint8(1)): stack.Push(first.(uint8) >= last.(uint8))
					case reflect.TypeOf(uint16(1)): stack.Push(first.(uint16) >= last.(uint16))
					case reflect.TypeOf(uint32(1)): stack.Push(first.(uint32) >= last.(uint32))
					case reflect.TypeOf(uint64(1)): stack.Push(first.(uint64) >= last.(uint64))
					case reflect.TypeOf(float32(1)): stack.Push(first.(float32) >= last.(float32))
					case reflect.TypeOf(float64(1)): stack.Push(first.(float64) >= last.(float64))
					case reflect.TypeOf(""): stack.Push(first.(string) >= last.(string))
					}
				}
			} else {
				switch reflect.TypeOf(first) {
				case reflect.TypeOf(int8(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) >= last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) >= last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for >=: 'int8' and 'array'"))
					}
				case reflect.TypeOf(int16(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) >= last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) >= last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for >=: 'int16' and 'array'"))
					}
				case reflect.TypeOf(int32(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) >= last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) >= last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for >=: 'int32' and 'array'"))
					}
				case reflect.TypeOf(int64(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) >= last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) >= last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for >=: 'int64' and 'array'"))
					}
				case reflect.TypeOf(uint8(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) >= last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) >= last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for >=: 'uint8' and 'array'"))
					}
				case reflect.TypeOf(uint16(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) >= last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) >= last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for >=: 'uint16' and 'array'"))
					}
				case reflect.TypeOf(uint32(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) >= last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) >= last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for >=: 'uint32' and 'array'"))
					}
				case reflect.TypeOf(uint64(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) >= last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) >= last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for >=: 'uint64' and 'array'"))
					}
				case reflect.TypeOf(float32(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) >= last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for >=: 'float32' and 'array'"))
					}
				case reflect.TypeOf(float64(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float64) >= last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for >=: 'float64' and 'array'"))
					}
				case reflect.TypeOf(string("")):
					if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for >=: 'string' and 'array'"))
					}
				case reflect.TypeOf([]interface{}{}):
					if len(first.([]interface{})) >= 0 {
						if reflect.TypeOf(first.([]interface{})[0]) != reflect.TypeOf(last) {
							fmt.Println(errors.New("TypeError: cannot eval '>=' between different type(s)"))
						} 
					}
					switch reflect.TypeOf(last) {
					case reflect.TypeOf(int8(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for >=: 'array' and 'int8'"))
					case reflect.TypeOf(int16(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for >=: 'array' and 'int16'"))
					case reflect.TypeOf(int32(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for >=: 'array' and 'int32'"))
					case reflect.TypeOf(int64(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for >=: 'array' and 'int64'"))
					case reflect.TypeOf(uint8(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for >=: 'array' and 'uint8'"))
					case reflect.TypeOf(uint16(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for >=: 'array' and 'uint16'"))
					case reflect.TypeOf(uint32(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for >=: 'array' and 'uint32'"))
					case reflect.TypeOf(uint64(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for >=: 'array' and 'uint64'"))
					case reflect.TypeOf(float32(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for >=: 'array' and 'float32'"))
					case reflect.TypeOf(float64(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for >=: 'array' and 'float64'"))
					case reflect.TypeOf(string("")):  fmt.Println(errors.New("TypeError: unsupported operand type(s) for >=: 'array' and 'string'"))
					}
				}
			}
		case "lt":
			last, _ := stack.Pop()
			first, _ := stack.Pop()
			if reflect.TypeOf(first) == reflect.TypeOf(last) {
				switch reflect.TypeOf(first).Kind() {
				case reflect.Slice:
					fmt.Println(errors.New("TypeError: unsupported operand type(s) for <: 'array' and 'array'"))
				default:
					switch reflect.TypeOf(first) {
					case reflect.TypeOf(int8(1)): stack.Push(first.(int8) < last.(int8))
					case reflect.TypeOf(int16(1)): stack.Push(first.(int16) < last.(int16))
					case reflect.TypeOf(int32(1)): stack.Push(first.(int32) < last.(int32))
					case reflect.TypeOf(int64(1)): stack.Push(first.(int64) < last.(int64))
					case reflect.TypeOf(uint8(1)): stack.Push(first.(uint8) < last.(uint8))
					case reflect.TypeOf(uint16(1)): stack.Push(first.(uint16) < last.(uint16))
					case reflect.TypeOf(uint32(1)): stack.Push(first.(uint32) < last.(uint32))
					case reflect.TypeOf(uint64(1)): stack.Push(first.(uint64) < last.(uint64))
					case reflect.TypeOf(float32(1)): stack.Push(first.(float32) < last.(float32))
					case reflect.TypeOf(float64(1)): stack.Push(first.(float64) < last.(float64))
					case reflect.TypeOf(""): stack.Push(first.(string) < last.(string))
					}
				}
			} else {
				switch reflect.TypeOf(first) {
				case reflect.TypeOf(int8(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) < last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) < last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for <: 'int8' and 'array'"))
					}
				case reflect.TypeOf(int16(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) < last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) < last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for <: 'int16' and 'array'"))
					}
				case reflect.TypeOf(int32(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) < last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) < last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for <: 'int32' and 'array'"))
					}
				case reflect.TypeOf(int64(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) < last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) < last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for <: 'int64' and 'array'"))
					}
				case reflect.TypeOf(uint8(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) < last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) < last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for <: 'uint8' and 'array'"))
					}
				case reflect.TypeOf(uint16(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) < last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) < last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for <: 'uint16' and 'array'"))
					}
				case reflect.TypeOf(uint32(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) < last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) < last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for <: 'uint32' and 'array'"))
					}
				case reflect.TypeOf(uint64(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) < last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) < last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for <: 'uint64' and 'array'"))
					}
				case reflect.TypeOf(float32(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) < last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for <: 'float32' and 'array'"))
					}
				case reflect.TypeOf(float64(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float64) < last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for <: 'float64' and 'array'"))
					}
				case reflect.TypeOf(string("")):
					if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for <: 'string' and 'array'"))
					}
				case reflect.TypeOf([]interface{}{}):
					if len(first.([]interface{})) < 0 {
						if reflect.TypeOf(first.([]interface{})[0]) != reflect.TypeOf(last) {
							fmt.Println(errors.New("TypeError: cannot eval '<' between different type(s)"))
						} 
					}
					switch reflect.TypeOf(last) {
					case reflect.TypeOf(int8(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for <: 'array' and 'int8'"))
					case reflect.TypeOf(int16(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for <: 'array' and 'int16'"))
					case reflect.TypeOf(int32(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for <: 'array' and 'int32'"))
					case reflect.TypeOf(int64(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for <: 'array' and 'int64'"))
					case reflect.TypeOf(uint8(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for <: 'array' and 'uint8'"))
					case reflect.TypeOf(uint16(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for <: 'array' and 'uint16'"))
					case reflect.TypeOf(uint32(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for <: 'array' and 'uint32'"))
					case reflect.TypeOf(uint64(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for <: 'array' and 'uint64'"))
					case reflect.TypeOf(float32(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for <: 'array' and 'float32'"))
					case reflect.TypeOf(float64(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for <: 'array' and 'float64'"))
					case reflect.TypeOf(string("")):  fmt.Println(errors.New("TypeError: unsupported operand type(s) for <: 'array' and 'string'"))
					}
				}
			}
		case "le":
			last, _ := stack.Pop()
			first, _ := stack.Pop()
			if reflect.TypeOf(first) == reflect.TypeOf(last) {
				switch reflect.TypeOf(first).Kind() {
				case reflect.Slice:
					fmt.Println(errors.New("TypeError: unsupported operand type(s) for <=: 'array' and 'array'"))
				default:
					switch reflect.TypeOf(first) {
					case reflect.TypeOf(int8(1)): stack.Push(first.(int8) <= last.(int8))
					case reflect.TypeOf(int16(1)): stack.Push(first.(int16) <= last.(int16))
					case reflect.TypeOf(int32(1)): stack.Push(first.(int32) <= last.(int32))
					case reflect.TypeOf(int64(1)): stack.Push(first.(int64) <= last.(int64))
					case reflect.TypeOf(uint8(1)): stack.Push(first.(uint8) <= last.(uint8))
					case reflect.TypeOf(uint16(1)): stack.Push(first.(uint16) <= last.(uint16))
					case reflect.TypeOf(uint32(1)): stack.Push(first.(uint32) <= last.(uint32))
					case reflect.TypeOf(uint64(1)): stack.Push(first.(uint64) <= last.(uint64))
					case reflect.TypeOf(float32(1)): stack.Push(first.(float32) <= last.(float32))
					case reflect.TypeOf(float64(1)): stack.Push(first.(float64) <= last.(float64))
					case reflect.TypeOf(""): stack.Push(first.(string) <= last.(string))
					}
				}
			} else {
				switch reflect.TypeOf(first) {
				case reflect.TypeOf(int8(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) <= last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) <= last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for <=: 'int8' and 'array'"))
					}
				case reflect.TypeOf(int16(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) <= last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) <= last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for <=: 'int16' and 'array'"))
					}
				case reflect.TypeOf(int32(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) <= last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) <= last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for <=: 'int32' and 'array'"))
					}
				case reflect.TypeOf(int64(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) <= last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) <= last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for <=: 'int64' and 'array'"))
					}
				case reflect.TypeOf(uint8(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) <= last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) <= last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for <=: 'uint8' and 'array'"))
					}
				case reflect.TypeOf(uint16(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) <= last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) <= last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for <=: 'uint16' and 'array'"))
					}
				case reflect.TypeOf(uint32(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) <= last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) <= last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for <=: 'uint32' and 'array'"))
					}
				case reflect.TypeOf(uint64(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) <= last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) <= last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for <=: 'uint64' and 'array'"))
					}
				case reflect.TypeOf(float32(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) <= last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for <=: 'float32' and 'array'"))
					}
				case reflect.TypeOf(float64(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float64) <= last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for <=: 'float64' and 'array'"))
					}
				case reflect.TypeOf(string("")):
					if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						fmt.Println(errors.New("TypeError: unsupported operand type(s) for <=: 'string' and 'array'"))
					}
				case reflect.TypeOf([]interface{}{}):
					if len(first.([]interface{})) <= 0 {
						if reflect.TypeOf(first.([]interface{})[0]) != reflect.TypeOf(last) {
							fmt.Println(errors.New("TypeError: cannot eval '<=' between different type(s)"))
						} 
					}
					switch reflect.TypeOf(last) {
					case reflect.TypeOf(int8(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for <=: 'array' and 'int8'"))
					case reflect.TypeOf(int16(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for <=: 'array' and 'int16'"))
					case reflect.TypeOf(int32(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for <=: 'array' and 'int32'"))
					case reflect.TypeOf(int64(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for <=: 'array' and 'int64'"))
					case reflect.TypeOf(uint8(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for <=: 'array' and 'uint8'"))
					case reflect.TypeOf(uint16(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for <=: 'array' and 'uint16'"))
					case reflect.TypeOf(uint32(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for <=: 'array' and 'uint32'"))
					case reflect.TypeOf(uint64(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for <=: 'array' and 'uint64'"))
					case reflect.TypeOf(float32(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for <=: 'array' and 'float32'"))
					case reflect.TypeOf(float64(1)): fmt.Println(errors.New("TypeError: unsupported operand type(s) for <=: 'array' and 'float64'"))
					case reflect.TypeOf(string("")):  fmt.Println(errors.New("TypeError: unsupported operand type(s) for <=: 'array' and 'string'"))
					}
				}
			}
		case "eq":
			last, _ := stack.Pop()
			first, _ := stack.Pop()
			if reflect.TypeOf(first) == reflect.TypeOf(last) {
				switch reflect.TypeOf(first).Kind() {
				case reflect.Slice:
					stack.Push(reflect.DeepEqual(first.([]interface{}), last.([]interface{})))
				default:
					switch reflect.TypeOf(first) {
					case reflect.TypeOf(int8(1)): stack.Push(first.(int8) == last.(int8))
					case reflect.TypeOf(int16(1)): stack.Push(first.(int16) == last.(int16))
					case reflect.TypeOf(int32(1)): stack.Push(first.(int32) == last.(int32))
					case reflect.TypeOf(int64(1)): stack.Push(first.(int64) == last.(int64))
					case reflect.TypeOf(uint8(1)): stack.Push(first.(uint8) == last.(uint8))
					case reflect.TypeOf(uint16(1)): stack.Push(first.(uint16) == last.(uint16))
					case reflect.TypeOf(uint32(1)): stack.Push(first.(uint32) == last.(uint32))
					case reflect.TypeOf(uint64(1)): stack.Push(first.(uint64) == last.(uint64))
					case reflect.TypeOf(float32(1)): stack.Push(first.(float32) == last.(float32))
					case reflect.TypeOf(float64(1)): stack.Push(first.(float64) == last.(float64))
					case reflect.TypeOf(""): stack.Push(first.(string) == last.(string))
					}
				}
			} else {
				switch reflect.TypeOf(first) {
				case reflect.TypeOf(int8(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) == last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) == last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						stack.Push(false)
					}
				case reflect.TypeOf(int16(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) == last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) == last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						stack.Push(false)
					}
				case reflect.TypeOf(int32(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) == last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) == last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						stack.Push(false)
					}
				case reflect.TypeOf(int64(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) == last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) == last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						stack.Push(false)
					}
				case reflect.TypeOf(uint8(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) == last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) == last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						stack.Push(false)
					}
				case reflect.TypeOf(uint16(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) == last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) == last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						stack.Push(false)
					}
				case reflect.TypeOf(uint32(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) == last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) == last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						stack.Push(false)
					}
				case reflect.TypeOf(uint64(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) == last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) == last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						stack.Push(false)
					}
				case reflect.TypeOf(float32(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) == last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						stack.Push(false)
					}
				case reflect.TypeOf(float64(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float64) == last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						stack.Push(false)
					}
				case reflect.TypeOf(string("")):
					if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						stack.Push(false)
					}
				case reflect.TypeOf([]interface{}{}):
					stack.Push(false)
				}
			}
		case "neq":
			last, _ := stack.Pop()
			first, _ := stack.Pop()
			if reflect.TypeOf(first) == reflect.TypeOf(last) {
				switch reflect.TypeOf(first).Kind() {
				case reflect.Slice:
					stack.Push(!reflect.DeepEqual(first.([]interface{}), last.([]interface{})))
				default:
					switch reflect.TypeOf(first) {
					case reflect.TypeOf(int8(1)): stack.Push(first.(int8) == last.(int8))
					case reflect.TypeOf(int16(1)): stack.Push(first.(int16) == last.(int16))
					case reflect.TypeOf(int32(1)): stack.Push(first.(int32) == last.(int32))
					case reflect.TypeOf(int64(1)): stack.Push(first.(int64) == last.(int64))
					case reflect.TypeOf(uint8(1)): stack.Push(first.(uint8) == last.(uint8))
					case reflect.TypeOf(uint16(1)): stack.Push(first.(uint16) == last.(uint16))
					case reflect.TypeOf(uint32(1)): stack.Push(first.(uint32) == last.(uint32))
					case reflect.TypeOf(uint64(1)): stack.Push(first.(uint64) == last.(uint64))
					case reflect.TypeOf(float32(1)): stack.Push(first.(float32) == last.(float32))
					case reflect.TypeOf(float64(1)): stack.Push(first.(float64) == last.(float64))
					case reflect.TypeOf(""): stack.Push(first.(string) == last.(string))
					}
				}
			} else {
				switch reflect.TypeOf(first) {
				case reflect.TypeOf(int8(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) == last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) == last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						stack.Push(true)
					}
				case reflect.TypeOf(int16(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) == last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) == last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						stack.Push(true)
					}
				case reflect.TypeOf(int32(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) == last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) == last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						stack.Push(true)
					}
				case reflect.TypeOf(int64(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) == last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) == last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						stack.Push(true)
					}
				case reflect.TypeOf(uint8(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) == last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) == last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						stack.Push(true)
					}
				case reflect.TypeOf(uint16(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) == last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) == last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						stack.Push(true)
					}
				case reflect.TypeOf(uint32(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) == last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) == last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						stack.Push(true)
					}
				case reflect.TypeOf(uint64(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) == last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) == last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						stack.Push(true)
					}
				case reflect.TypeOf(float32(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) == last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						stack.Push(true)
					}
				case reflect.TypeOf(float64(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float64) == last.(float64))
					} else if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						stack.Push(true)
					}
				case reflect.TypeOf(string("")):
					if reflect.TypeOf(last) == reflect.TypeOf([]interface{}{}) {
						stack.Push(true)
					}
				case reflect.TypeOf([]interface{}{}):
					stack.Push(true)
				}
			}
		case "lds":
			popped := stack.Pop()
			
		case "upd":
			popped := stack.Pop()
			switch reflect.TypeOf(popped) {
			case reflect.TypeOf(int8(1)): stack.stack[op[1]] = popped.(int8)
			case reflect.TypeOf(int16(1)): stack.stack[op[1]] = popped.(int16)
			case reflect.TypeOf(int32(1)): stack.stack[op[1]] = popped.(int32)
			case reflect.TypeOf(int64(1)): stack.stack[op[1]] = popped.(int64)
			case reflect.TypeOf(uint8(1)): stack.stack[op[1]] = popped.(uint8)
			case reflect.TypeOf(uint16(1)): stack.stack[op[1]] = popped.(uint16)
			case reflect.TypeOf(uint32(1)): stack.stack[op[1]] = popped.(uint32)
			case reflect.TypeOf(uint64(1)): stack.stack[op[1]] = popped.(uint64)
			case reflect.TypeOf(float32(1)): stack.stack[op[1]] = popped.(float32)
			case reflect.TypeOf(float64(1)): stack.stack[op[1]] = popped.(float64)
			case reflect.TypeOf(string("")): stack.stack[op[1]] = popped.(string)
			case reflect.TypeOf([]interface{}{}): stack.stack[op[1]] = popped.([]interface{}{})
			}
		case "load":
			//
		case "store":
			popped := stack.Pop()
			switch reflect.TypeOf(popped) {
			case reflect.TypeOf(int8(1)): stack.Push(popped.(int8))
			case reflect.TypeOf(int16(1)): stack.Push(popped.(int16))
			case reflect.TypeOf(int32(1)): stack.Push(popped.(int32))
			case reflect.TypeOf(int64(1)): stack.Push(popped.(int64))
			case reflect.TypeOf(uint8(1)): stack.Push(popped.(uint8))
			case reflect.TypeOf(uint16(1)): stack.Push(popped.(uint16))
			case reflect.TypeOf(uint32(1)): stack.Push(popped.(uint32))
			case reflect.TypeOf(uint64(1)): stack.Push(popped.(uint64))
			case reflect.TypeOf(float32(1)): stack.Push(popped.(float32))
			case reflect.TypeOf(float64(1)): stack.Push(popped.(float64))
			case reflect.TypeOf(string("")): stack.Push(popped.(string))
			case reflect.TypeOf([]interface{}{}): stack.Push(popped.([]interface{}{}))
			}
		case "call":
			func_name := op[1]
			strings.Contains(func_name, "&")
		}

	}
	if DEBUG {
		fmt.Println(stack.stack)
	}
	return err
}

//ReadFile : 파일 처리(읽기)!
func ReadFile() error {
	file, err := os.Open("src/src.basm")
	if err != nil {
		return errCantReadFile
	}
	defer file.Close()
	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		fileData += string(fileScanner.Text()) + "\n"
	}
	return nil
}

func (s Stack) IsEmpty() bool {
	return len(s.stack) == 0
}

func (s *Stack) Push(data interface{}) {
	s.stack = append(s.stack, data)
}

func (s *Stack) Pop() (interface{}, error) {
	var err error = nil
	if s.IsEmpty() {
		err = errNothingToPop
	}

	poppedValue := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return poppedValue, err

}

func (s Stack) Top() (interface{}, error) {
	var err error = nil
	if s.IsEmpty() {
		err = errStackEmpty
	}
	top := len(s.stack) - 1
	return s.stack[top], err
}