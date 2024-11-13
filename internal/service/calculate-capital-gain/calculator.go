package calculatecapitalgain

import "io"

// Calculator is an interface that defines the Calculate methods.
type Calculator interface {
	Calculate(Reader, Writer) error
}

// Reader is an interface that extends io.Reader.
type Reader interface {
	io.Reader
}

// Writer is an interface that extends io.Writer.
type Writer interface {
	io.Writer
}
