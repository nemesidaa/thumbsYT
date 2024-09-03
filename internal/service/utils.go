package service

import "fmt"

func Port(addr int) string {
	return fmt.Sprintf(":%d", addr)
}
