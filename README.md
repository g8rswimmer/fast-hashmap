# fast-hashmap
Fast implementation of a hashmap

## Getting Started 

### Installing
To start using `hashmap`, install `GO` and run `go get`
```
go get -u github.com/g8rswimmer/fast-hashmap
```
This will retrive the package

## Usage 

```
    hm := hashmap.New()
    hm.Put(3, "this is an example")
    value := hm.Get(3)
    fmt.Printf("%v\n", value)
```

## Benchmarking

```
goos: darwin
goarch: amd64
pkg: github.com/g8rswimmer/fast-hashmap
BenchmarkPut-8   	   20000	     63940 ns/op	   69632 B/op	     769 allocs/op

BenchmarkGet-8   	10000000	       207 ns/op	       0 B/op	       0 allocs/op
```
