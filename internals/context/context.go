package context

var GlobalContext *Context

type Context map[string]interface{}

func (c Context) GetValue(key string) (interface{}, bool) {
	value, ok := c[key]
	return value, ok
}

func (c *Context) PutValue(key string, value interface{}) bool {
	(*c)[key] = value
	return true
}

func NewContext() *Context {
	return &Context{}
}

func InitContext() *Context {
	GlobalContext = NewContext()
	return GlobalContext
}
