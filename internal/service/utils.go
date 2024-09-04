package service

import "fmt"

func Addr(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}
