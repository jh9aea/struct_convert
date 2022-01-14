# struct_convert
golang struct convert func



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
