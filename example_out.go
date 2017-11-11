package foo

import "io"

func bar3(w io.Writer) error {
	_, err := w.Write([]byte("foo"))
	if err != nil {
		return err
	}
	_, err = w.Write([]byte("bar"))
	if err != nil {
		return err
	}
	_, err = w.Write([]byte("baz"))
	if err != nil {
		return err
	}
	_, err = w.Write([]byte("quz"))
	if err != nil {
		return err
	}
	return nil
}
