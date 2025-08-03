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
			_internal_NotImplemented("ast::Dump::NodeArray")
		}

		if index < len(builder.root.object.children)-1 {
			buffer += ","
		}
	}

	buffer += "}"

	return buffer
}
