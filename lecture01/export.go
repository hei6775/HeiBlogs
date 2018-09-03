package lecture01

type Method interface {
	Isempty() bool
	Size() int
	Pop() interface{}
	Push() interface{}
}
