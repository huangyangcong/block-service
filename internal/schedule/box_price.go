package schedule

import (
	"fmt"
)

func NewBoxPrice() *Job {
	return &Job{"@every 1s", func() {
		fmt.Errorf("aaaa")
	}}
}
