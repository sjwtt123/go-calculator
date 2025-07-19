package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var s string = "2*3+3.15+3*2"

	fmt.Println("请输入要计算的表达式：", s)
	as, err := xunhuan(s)
	if err != nil {
		fmt.Println("计算错误")
	} else {
		fmt.Println("最终结果为：", as)
	}

}

// 循环处理括号内的算术式，以及整个表达式
func xunhuan(s string) (string, error) {
	var err error
	for {
		if strings.Contains(s, "(") {
			s, err = qu(s)
			if err != nil {
				return "", err
			}
			continue
		} else if strings.Contains(s, "*") {
			s, err = OnlyChen(s)
			if err != nil {
				return "", err
			}
			continue
		} else if strings.Contains(s, "+") {
			s, err = OnlyAdd(s)
			if err != nil {
				return "", err
			}
			continue
		} else {
			if _, err := strconv.Atoi(s); err != nil {
				return "", fmt.Errorf("表达式包含非法字符: %s", s)
			}
			break
		}

	}
	return s, err
}

// 去括号时，可以获得第一个），最后一个（，循环得到没括号的式子，也可以，获得第一个（，最后一个），直接得到没括号的式子
func qu(s string) (string, error) {
	lastIndex := strings.Index(s, ")")
	firstIndex := strings.LastIndex(s, "(")
	s1 := s[firstIndex+1 : lastIndex]
	finn, err := xunhuan(s1)
	if err != nil {
		return "", err
	}
	s = strings.Replace(s, s[firstIndex:lastIndex+1], finn, 1)
	fmt.Println("replace():", s)
	return s, nil
}

// OnlyAdd 处理加法
func OnlyAdd(s string) (string, error) {
	index := strings.Index(s, "+")
	one, two := FindNumber(s, index)
	fmt.Println("one:", two, "two:", one)
	atoi1, err01 := strconv.Atoi(one)
	if err01 != nil {
		return s, fmt.Errorf("无法转换数字 '%s': %v", one, err01)

	}
	atoi2, error2 := strconv.Atoi(two)
	if error2 != nil {
		return s, fmt.Errorf("无法转换数字 '%s': %v", one, error2)

	}
	atoi2 = atoi2 + atoi1
	itoa := strconv.Itoa(atoi2)
	s = strings.Replace(s, s[index-len(two):index+len(one)+1], itoa, 1)
	fmt.Println("replace+:", s)
	return s, nil
}

// OnlyChen 处理乘法
func OnlyChen(s string) (string, error) {
	index := strings.Index(s, "*")
	one, two := FindNumber(s, index)
	fmt.Println("one:", two, "two:", one)
	atoi1, error1 := strconv.Atoi(one)
	if error1 != nil {
		return s, fmt.Errorf("无法转换数字 '%s': %v", one, error1)
	}
	atoi2, error2 := strconv.Atoi(two)
	if error2 != nil {
		return s, fmt.Errorf("无法转换数字 '%s': %v", one, error2)
	}
	atoi2 = atoi2 * atoi1
	itoa := strconv.Itoa(atoi2)
	s = strings.Replace(s, s[index-len(two):index+len(one)+1], itoa, 1)
	fmt.Println("replace*:", s)
	return s, nil
}

// FindNumber 得到符号两旁的数字
func FindNumber(s string, index int) (string, string) {
	var s1, s2 string
	for j := index + 1; j < len(s); j++ {
		if j == len(s)-1 {
			s1 = s[index+1:]
			break
		}
		if s[j] >= 48 && s[j] <= 57 && j != len(s)-1 {
			continue
		} else {
			s1 = s[index+1 : j]
			break
		}
	}
	for j := index - 1; j >= 0; j-- {
		if j == 0 {
			s2 = s[:index]
			break
		}
		if s[j] >= 48 && s[j] <= 57 {
			continue
		} else {
			s2 = s[j+1 : index]
			break
		}
	}
	return s1, s2
}
