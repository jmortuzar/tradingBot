package context

var GlobalContext *Context

type Context interface {
	GetValue(key string) (interface{}, bool)
	PutValue(key string, value interface{}) bool
}

type myContext map[string]interface{}

func (c *myContext) GetValue(key string) (interface{}, bool) {
	value, ok := (*c)[key]
	return value, ok
}

func (c *myContext) PutValue(key string, value interface{}) bool {
	(*c)[key] = value
	return true
}

func NewContext() *Context {
	var ctx Context = &myContext{}
	return &ctx
}

func InitContext() *Context {
	GlobalContext = NewContext()
	return GlobalContext
}
