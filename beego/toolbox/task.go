package toolbox

import "time"

//bounds provides a range of acceptable values (plus a map of name to value).
type bounds struct {
	min,max uint
	names map[string]uint
}

const (
	// Set the top bit if a star was included in the expression.
	starBit = 1 << 63
)

// Schedule time taks schedule
type Schedule struct {
	Second uint64
	Minute uint64
	Hour   uint64
	Day    uint64
	Month  uint64
	Week   uint64
}

//TaskFunc task func type
type TaskFunc func()error

//Tasker task interface
type Tasker interface {
	GetSpec() string
	GetStatus() string
	Run() error
	SetNext(time.Time)
	GetNext()time.Time
	SetPrev()time.Time
	GetPrev()time.Time
}

//The bounds for each field
var (
	
)

