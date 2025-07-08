package src

import "time"

type Transaction struct {
	amount    float32
	timestamp time.Time
	from      string
	to        string
	// TODO create operations enum
	operation string
}
