package ast

import (
	"fmt"
	"strconv"
)

func _internal_NotImplemented(method string) {
	panic(fmt.Errorf("Method not implemented; %s\n", method))
}

func _internal_CreateBuilder() *Builder {
	return &Builder{}
}

func _internal_CreateObject() *AstObject {
	return &AstObject{
		children: make([]*AstNode, 0),
	}
}

func _internal_CreateArray() *AstArray {
	return &AstArray{
		children: make([]*AstNode, 0),
	}
}

func (serializer *Serializer) _internal_DumpValue(object *AstNode, silent bool) {
	if silent == false {
		serializer.buffer += fmt.Sprintf("\"%s\":", object.name)
	}

	switch object.value.astValueType {
	case ValueBool:
		serializer.buffer += strconv.FormatBool(*object.value._bool)
	case ValueNumber:
		serializer.buffer += strconv.FormatFloat(*object.value._number, 'g', 4, 64)
	case ValueString:
		serializer.buffer += fmt.Sprintf("\"%s\"", *object.value._string)
	case ValueNull:
		serializer.buffer += "null"
	default:
		serializer.buffer += "null"
	}
}

func (serializer *Serializer) _internal_DumpObject(object *AstNode, silent bool) {
	if silent == false {
		serializer.buffer += fmt.Sprintf("\"%s\":{", object.name)
	} else {
		serializer.buffer += "{"
	}

	for index, element := range object.object.children {
		if element == nil {
			continue
		}
		switch element.astNodeType {
		case NodeValue:
			serializer._internal_DumpValue(element, false)
		case NodeObject:
			serializer._internal_DumpObject(element, false)
		case NodeArray:
			serializer._internal_DumpArray(element, false)

		}

		if index < len(object.object.children)-1 {
			serializer.buffer += ","
		}
	}

	serializer.buffer += "}"
}

func (serializer *Serializer) _internal_DumpArray(array *AstNode, silent bool) {
	if silent == false {
		serializer.buffer += fmt.Sprintf("\"%s\":[", array.name)
	} else {
		serializer.buffer += "["
	}

	for index, element := range array.array.children {
		if element == nil {
			continue
		}

		switch element.astNodeType {
		case NodeValue:
			serializer._internal_DumpValue(element, true)
		case NodeObject:
			serializer._internal_DumpObject(element, true)
		case NodeArray:
			serializer._internal_DumpArray(element, true)
		}

		if index < len(array.array.children)-1 {
			serializer.buffer += ","
		}
	}

	serializer.buffer += "]"
}
