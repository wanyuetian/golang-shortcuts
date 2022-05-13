package main

// sentinel error 预定义错误类型
// https://github.com/golang/go/blob/master/src/bufio/bufio.go
/*
	var (
		ErrInvalidUnreadByte = errors.New("bufio: invalid use of UnreadByte")
		ErrInvalidUnreadRune = errors.New("bufio: invalid use of UnreadRune")
		ErrBufferFull        = errors.New("bufio: buffer full")
		ErrNegativeCount     = errors.New("bufio: negative count")
	)
*/

// Error types
// https://github.com/golang/go/blob/master/src/os/error.go
/*
	type PathError struct {
		Op   string
		Path string
		Err  error
	}
*/

// Opaque errors 不透明错误处理
// https://github.com/golang/go/blob/master/src/net/net.go
/*
	type Error interface {
		error
		Timeout() bool // Is the error a timeout?

		// Deprecated: Temporary errors are not well-defined.
		// Most "temporary" errors are timeouts, and the few exceptions are surprising.
		// Do not use this method.
		Temporary() bool
	}
*/

func main() {
	//errors.Wrap()
	
}
