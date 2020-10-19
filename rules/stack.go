package rules

type objStack struct {
	items []IFiled
}

func (o *objStack) empty() bool {
	return len(o.items) == 0
}

func (o *objStack) peek() IFiled {
	return o.items[len(o.items)-1]
}

func (o *objStack) pop() IFiled {
	val := o.peek()
	o.items = o.items[:len(o.items)-1]
	return val
}

func (o *objStack) push(item IFiled) {
	o.items = append(o.items, item)
}

func (o *objStack) clear() {
	o.items = nil
}
