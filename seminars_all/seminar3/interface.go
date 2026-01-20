package main

type Fooer interface {
	Foo()
}

type Foo struct{}

func (f Foo) Foo() {}

type Bar struct{}

func (b Bar) Bar() {}


func example2() {
	var (
		foo Foo
		bar Bar
	)
	var _ Fooer = foo
	var _ Fooer = bar //будект ошибка тк в Fooer есть Foo, a не Bar
}

type Reader interface {
	Read(p []byte) (n int, err error)
}
type Closer interface {
	Close() error
}

type ReadCloser1 interface {
	Reader
	Closer
}
type ReadCloser2 interface { //тоже самое что и выше
	Read(p []byte) (n int, err error)
	Close() error
}