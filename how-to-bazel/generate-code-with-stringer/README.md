# How to generate and compile code with Go Stringer tool

Ready-to-go example base on [Stack overflow](https://stackoverflow.com/questions/66499040/bazel-build-gogenerate-stringer-stringer-cant-happen-constant-is-not-an-i) question.


> Stringer is a tool to automate the creation of methods that satisfy the
> fmt.Stringer interface. Given the name of a (signed or unsigned) integer type
> T that has constants defined, stringer will create a new self-contained Go
> source file implementing
```golang
func (t T) String() string
```

See [docs](https://pkg.go.dev/golang.org/x/tools/cmd/stringer) for more details.

