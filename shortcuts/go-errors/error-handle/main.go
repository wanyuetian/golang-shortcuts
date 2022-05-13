package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
)

// 原则
// 1 You should only handle error once

// CountLines 简化错误处理
func CountLines(r io.Reader) (int, error) {
	var (
		br    = bufio.NewReader(r)
		lines int
		err   error
	)

	for {
		_, err = br.ReadString('\n')
		lines++
		if err != nil {
			break
		}
	}
	if err != io.EOF {
		return 0, err
	}
	return lines, nil
}

// CountLines1 改进版本
func CountLines1(r io.Reader) (int, error) {
	sc := bufio.NewScanner(r)
	lines := 0
	for sc.Scan() {
		lines++
	}
	return lines, sc.Err()
}

// Wrap errors

type Request struct {
	*http.Request
	User string
}

func AuthenticateRequest(r *Request) error {
	err := authenticate(r.User)
	if err != nil {
		return fmt.Errorf("authenticate failed: %v", err)
	}
	return nil
}

func authenticate(user string) error {
	return nil
}

//
// 第三方库 https://github.com/pkg/errors
