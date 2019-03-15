package Lecture01
//接口
type Method interface {
	Isempty() bool
	Size() int
	Pop() interface{}
	Push() interface{}
}
