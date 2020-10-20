package rules

import (
	"strconv"
	"strings"
)

type FieldType int

const (
	TypeDummy FieldType = iota
	TypeAttr
	TypeResultBoolean

	TypeValueString
	TypeValueBoolean
	TypeValueNull
	TypeValueDouble
	TypeValueInteger
	TypeValueListStrings
	TypeValueListDoubles
	TypeValueListIntegers
)

type IFiled interface {
	Type() FieldType
	Raw() string
	Value() interface{}
}

type ResultBoolean struct {
	res bool
}

func NewResultBoolean(res bool) *ResultBoolean {
	return &ResultBoolean{res: res}
}

func (r *ResultBoolean) Type() FieldType {
	return TypeResultBoolean
}

func (r *ResultBoolean) Raw() string {
	if r.res {
		return "true"
	}
	return "false"
}

func (r *ResultBoolean) Value() interface{} {
	return r.res
}

func NewAttr(text string) *Attr {
	return &Attr{text: text}
}

type Attr struct {
	text string
}

func (a *Attr) Value() interface{} {
	return a.text
}

func (a *Attr) Type() FieldType {
	return TypeAttr
}

func (a *Attr) Raw() string {
	return a.text
}

type ValueLisIntegers struct {
	text string
}

func NewValueListIntegers(text string) *ValueLisIntegers {
	return &ValueLisIntegers{text: text}
}

func (o *ValueLisIntegers) Type() FieldType {
	return TypeValueListIntegers
}

func (o *ValueLisIntegers) Raw() string {
	return o.text
}

func (o *ValueLisIntegers) Value() interface{} {
	tmp := o.text[1 : len(o.text)-1]
	arr := strings.Split(tmp, ",")
	integers := make([]int64, len(arr))
	for i := range arr {
		integers[i], _ = strconv.ParseInt(arr[i], 10, 64)
	}
	return integers
}

type ValueListFloats struct {
	text string
}

func NewValueListFloats(text string) *ValueListFloats {
	return &ValueListFloats{text: text}
}

func (o *ValueListFloats) Type() FieldType {
	return TypeValueListDoubles
}

func (o *ValueListFloats) Raw() string {
	return o.text
}

func (o *ValueListFloats) Value() interface{} {
	tmp := o.text[1 : len(o.text)-1]
	arr := strings.Split(tmp, ",")
	doubles := make([]float64, len(arr))
	for i := range arr {
		doubles[i], _ = strconv.ParseFloat(arr[i], 64)
	}
	return doubles
}

type ValueListStrings struct {
	text string
}

func NewValueListStrings(text string) *ValueListStrings {
	return &ValueListStrings{text: text}
}

func (o *ValueListStrings) Type() FieldType {
	return TypeValueListStrings
}

func (o *ValueListStrings) Raw() string {
	return o.text
}

func (o *ValueListStrings) Value() interface{} {
	tmp := o.text[1 : len(o.text)-1]
	return strings.Split(tmp, ",")
}

type ValueInteger struct {
	text string
}

func NewValueInteger(text string) *ValueInteger {
	return &ValueInteger{text: text}
}

func (o *ValueInteger) Type() FieldType {
	return TypeValueInteger
}

func (o *ValueInteger) Raw() string {
	return o.text
}

func (o *ValueInteger) Value() interface{} {
	ret, _ := strconv.ParseInt(o.text, 10, 64)
	return ret
}

type ValueFloat struct {
	text string
}

func NewValueFloat(text string) *ValueFloat {
	return &ValueFloat{text: text}
}

func (o *ValueFloat) Type() FieldType {
	return TypeValueDouble
}

func (o *ValueFloat) Raw() string {
	return o.text
}

func (o *ValueFloat) Value() interface{} {
	ret, _ := strconv.ParseFloat(o.text, 64)
	return ret
}

type ValueNull struct {
}

func NewValueNull() *ValueNull {
	return &ValueNull{}
}

func (o *ValueNull) Type() FieldType {
	return TypeValueNull
}

func (o *ValueNull) Raw() string {
	return ""
}

func (o *ValueNull) Value() interface{} {
	return ""
}

type ValueBoolean struct {
	text string
}

func NewValueBoolean(text string) *ValueBoolean {
	return &ValueBoolean{text: text}
}

func (o *ValueBoolean) Type() FieldType {
	return TypeValueBoolean
}

func (o *ValueBoolean) Raw() string {
	return o.text
}

func (o *ValueBoolean) Value() interface{} {
	switch strings.ToLower(o.text) {
	case "true":
		return true
	default:
		return false
	}
}

type ValueString struct {
	text string
}

func NewValueString(text string) *ValueString {
	return &ValueString{text: text}
}

func (o *ValueString) Type() FieldType {
	return TypeValueString
}

func (o *ValueString) Raw() string {
	return o.text
}

func (o *ValueString) Value() interface{} {
	if len(o.text) >= 2 && o.text[0] == '"' && o.text[len(o.text)-1] == '"' {
		return o.text[1 : len(o.text)-1]
	}
	return o.text
}
