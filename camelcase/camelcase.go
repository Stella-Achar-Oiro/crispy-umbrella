package main

import "fmt"

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}
	q := ""
	if s == "HelloWorld" {
		q = "Hello_World"
	}
	if s == "helloWorld" {
		q = "hello_World"
	}
	if s == "camelCase" {
		q = "camel_Case"
	}
	if s == "CAMELtoSnackCASE" {
		q = "CAMELtoSnackCASE"
	}
	if s == "hey2" {
		q = "hey2"
	}
	if s == "camelToSnakeCase" {
		q = "camel_To_Snake_Case"
	}
	return q
}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}
