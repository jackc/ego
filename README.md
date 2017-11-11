# Ego

Ego is an experimental shorthand for Go error handling.

```
jack@happy:~/dev/go/src/github.com/jackc/ego$ cat example_in.go
package foo

import "io"

func bar3(w io.Writer) error {
  _, err := w.Write([]byte("foo")) //; if err != nil { return err }
  _, err = w.Write([]byte("bar"))  //; if err != nil { return err }
  _, err = w.Write([]byte("baz"))  //; if err != nil { return err }
  _, err = w.Write([]byte("quz"))  //; if err != nil { return err }
  return nil
}
jack@happy:~/dev/go/src/github.com/jackc/ego$ ruby ego_expand.rb < example_in.go | gofmt
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
```

```
jack@happy:~/dev/go/src/github.com/jackc/ego$ cat example_out.go
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
jack@happy:~/dev/go/src/github.com/jackc/ego$ ruby ego_contract.rb < example_out.go
package foo

import "io"

func bar3(w io.Writer) error {
  _, err := w.Write([]byte("foo")) //; if err != nil { return err }
  _, err = w.Write([]byte("bar")) //; if err != nil { return err }
  _, err = w.Write([]byte("baz")) //; if err != nil { return err }
  _, err = w.Write([]byte("quz")) //; if err != nil { return err }
  return nil
}
```
