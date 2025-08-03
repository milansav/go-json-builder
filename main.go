package main

import (
	"fmt"
	"json-builder/ast"
)

func main() {
	root := ast.Object(func(_root *ast.Builder) {
		_root.Property("foo", ast.Object(func(_foo *ast.Builder) {
			_foo.Property("baz", ast.Value(true))
			_foo.Property("bar", ast.Value(10.1))
			_foo.Property("test_array", ast.Array(func(_test_array *ast.Builder) {

			}))
		}))
	})

	fmt.Printf("%s\n", root.Dump())
}
