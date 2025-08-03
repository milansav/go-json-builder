package ast

func (builder *Builder) Dump() string {
	buffer := "{"

	for index, element := range builder.root.object.children {
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

		if index < len(builder.root.object.children)-1 {
			buffer += ","
		}
	}

	buffer += "}"

	return buffer
}
