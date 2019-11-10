package models

import "fmt"

type ServerPort string

func (p ServerPort) String() string {
	return fmt.Sprintf(":%s", string(p))
}
