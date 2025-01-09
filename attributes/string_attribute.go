package attributes

type String struct {
	key   string
	value string
}

func NewString(key, value string) *String {
	return &String{key, value}
}

func (attribute *String) Key() string {
	return attribute.key
}

func (attribute *String) Value() string {
	return attribute.value
}
