package ast

type AstNodeType = int
type AstPrimitiveType = int

type BuildFn = func(_builder *Builder)

const (
	NodeObject AstNodeType = iota
	NodeArray
	NodeValue
)

const (
	ValueNull AstPrimitiveType = iota
	ValueString
	ValueNumber
	ValueBool
)

var NODE_DEFAULT_CHILDREN_COUNT = 5

type AstObject struct {
	children []*AstNode
}

// Array is basically the same thing as AstObject, except the names can be undefined or index
type AstArray = AstObject

type AstValue struct {
	astValueType AstPrimitiveType

	_string *string
	_number *float64
	_bool   *bool
}

type AstNode struct {
	name        string
	astNodeType AstNodeType

	// astNodeType = NodeObject
	object *AstObject

	// astNodeType = NodeArray
	array *AstArray

	// astNodeType = NodeValue
	value *AstValue
}

type Builder struct {
	root *AstNode
}

func (builder *Builder) Property(name string, _builder *Builder) *Builder {

	_builder.root.name = name

	builder.root.object.children = append(builder.root.object.children, _builder.root)

	return builder
}

func Object(_buildfn BuildFn) *Builder {
	builder := _internal_CreateBuilder()

	_node := &AstNode{
		name:        "ast::Object[default name]",
		astNodeType: NodeObject,
		object:      _internal_CreateObject(),
	}

	builder.root = _node

	_buildfn(builder)

	return builder
}

func Array(_buildFn BuildFn) *Builder {
	builder := _internal_CreateBuilder()

	_node := &AstNode{
		name:        "ast::Array[default name]",
		astNodeType: NodeArray,
		array:       _internal_CreateArray(),
	}

	builder.root = _node

	_buildFn(builder)

	return builder
}

func Value[T string | float64 | bool](val T) *Builder {
	builder := _internal_CreateBuilder()

	_node := &AstNode{
		name:        "ast::Value[default name]",
		astNodeType: NodeValue,
		value:       &AstValue{},
	}

	if v, ok := any(val).(bool); ok {
		_node.value.astValueType = ValueBool
		_node.value._bool = new(bool)
		*_node.value._bool = v
	} else if v, ok := any(val).(float64); ok {
		_node.value.astValueType = ValueNumber
		_node.value._number = new(float64)
		*_node.value._number = v
	} else if v, ok := any(val).(string); ok {
		_node.value.astValueType = ValueString
		_node.value._string = new(string)
		*_node.value._string = v
	}

	builder.root = _node

	return builder
}
