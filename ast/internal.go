package ast

import (
	"fmt"
	"strconv"
)

func _internal_NotImplemented(method string) string {
	panic(fmt.Errorf("Method not implemented; %s\n", method))
}

func _internal_CreateBuilder() *Builder {
	return &Builder{}
}

func _internal_CreateObject() *AstObject {
	return &AstObject{
		children: make([]*AstNode, NODE_DEFAULT_CHILDREN_COUNT),
	}
}

func _internal_CreateArray() *AstArray {
	return &AstArray{
		children: make([]*AstNode, NODE_DEFAULT_CHILDREN_COUNT),
	}
}

func (object *AstNode) _internal_DumpValue() string {
	buffer := fmt.Sprintf("\"%s\":", object.name)

	switch object.value.astValueType {
	case ValueBool:
		buffer += strconv.FormatBool(*object.value._bool)
	case ValueNumber:
		buffer += strconv.FormatFloat(*object.value._number, 'g', 4, 64)
	case ValueString:
		buffer += fmt.Sprintf("\"%s\"", *object.value._string)
	case ValueNull:
		buffer += "null"
	default:
		buffer += "null"
	}

	return buffer
}

func (object *AstNode) _internal_DumpObject() string {
	buffer := fmt.Sprintf("\"%s\":{", object.name)
	for index, element := range object.object.children {
		if element == nil {
			continue
		}
		switch element.astNodeType {
		case NodeValue:
			buffer += element._internal_DumpValue()
		case NodeObject:
			buffer += element._internal_DumpObject()
		case NodeArray:
			buffer += element._internal_DumpArray()

		}

		if index < len(object.object.children)-1 {
			buffer += ","
		}
	}

	buffer += "}"

	return buffer
}

func (array *AstNode) _internal_DumpArray() string {
	buffer := fmt.Sprintf("\"%s\":[", array.name)

	for index, element := range array.array.children {
		if element == nil {
			continue
		}

		switch element.astNodeType {
		case NodeValue:
			buffer += element._internal_DumpValue()
		case NodeObject:
			buffer += element._internal_DumpObject()
		case NodeArray:
			buffer += element._internal_DumpArray()
		}

		if index < len(array.object.children)-1 {
			buffer += ","
		}
	}

	buffer += "]"

	return buffer
}
