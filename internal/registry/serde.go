package registry

type Serde interface{
	Register(v Registrable, options ...BuildOption)error
	RegisterKey(key string, v any, options ...BuildOption)error
	RegisterFactory(key string, fn func()any, options ...BuildOption)error
}