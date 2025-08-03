package ast

type Serializer struct {
	buffer string
}

func (serializer Serializer) Dump(builder *Builder) string {
	serializer.buffer += "{"
	for index, element := range builder.root.object.children {
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

		if index < len(builder.root.object.children)-1 {
			serializer.buffer += ","
		}
	}

	serializer.buffer += "}"

	return serializer.buffer
}
