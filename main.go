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
			_foo.Property("null", ast.Null())
			_foo.Property("test_array", ast.Array(func(_test_array *ast.Builder) {
				_test_array.Item(ast.Value(10))
				_test_array.Item(ast.Object(func(_item *ast.Builder) {
					_item.Property("test_property", ast.Value("Hello World"))
				}))
				_test_array.Item(ast.Null())
			}))
		}))
	})

	fmt.Printf("%s\n", ast.Serializer{}.Dump(root))
}
