package register

type Register interface {
	Write(value int64) bool
	Read() int64
}
