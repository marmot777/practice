package main

import (
	"errors"
	"fmt"
	"strings"
)

func Hello() string {
	return "Hello, world"
}

func main() {
	fmt.Println(Hello())
}

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Division(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("被除数不能为0")
	}
	return a / b, nil
}

func Split(s, sep string) (result []string) {
	i := strings.Index(s, sep)

	for i > -1 {
		result = append(result, s[:i])
		s = s[i+len(sep):]
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}

/*
创建了一个author struct，它包含字段名、lastName和bio。添加了一个方法fullName()，将作者作为接收者类型，这将返回作者的全名。
*/
type author struct {
	firstName string
	lastName  string
	bio       string
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

/*
post struct有字段标题、内容。它还有一个嵌入式匿名字段作者。这个字段表示post struct是由author组成的。现在post struct可以访问作者结构的所有字段和方法。还在post struct中添加了details()方法
*/
type post struct {
	title   string
	content string
	author
}

func (p post) details() {
	fmt.Println("Title: ", p.title)
	fmt.Println("Content: ", p.content)
	fmt.Println("Author: ", p.fullName())
	fmt.Println("Bio: ", p.bio)
}
