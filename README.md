# struct_convert
golang struct convert func

## install
`go get github.com/jh9aea/struct_convert`

or just copy main.go

## use
```go
type aStruct struct {
	A string
	B string
}

type bStruct struct {
	B string
}

a := aStruct{A:"a", B: "b"}
b, err := Convert[bStruct](a)

```
