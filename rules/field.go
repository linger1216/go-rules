package rules

import (
	"strconv"
	"strings"
)

type FieldType int

const (
	TypeUnknown FieldType = iota
	TypeAttr

	TypeBoolean

	TypeLogicalAnd
	TypeLogicalOr
	TypeLogicalNot

	TypeOpEQ
	TypeOpNE
	TypeOpGT
	TypeOpLT
	TypeOpGE
	TypeOpLE
	TypeOpContain
	TypeOpPrefix
	TypeOpSuffix
	TypeOpRegex
	TypeOpIn

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
	return TypeBoolean
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

type OpIn struct {
	text string
}

func NewOpIn(text string) *OpIn {
	return &OpIn{text: text}
}

func (o *OpIn) Type() FieldType {
	switch strings.ToLower(o.text) {
	case "in":
		return TypeOpIn
	}
	return TypeUnknown
}

func (o *OpIn) Raw() string {
	return o.text
}

func (o *OpIn) Value() interface{} {
	return "in"
}

type OpRegex struct {
	text string
}

func NewOpRegex(text string) *OpRegex {
	return &OpRegex{text: text}
}

func (o *OpRegex) Type() FieldType {
	switch strings.ToLower(o.text) {
	case "regex":
		return TypeOpRegex
	}
	return TypeUnknown
}

func (o *OpRegex) Raw() string {
	return o.text
}

func (o *OpRegex) Value() interface{} {
	return "regex"
}

type OpSuffix struct {
	text string
}

func NewOpSuffix(text string) *OpSuffix {
	return &OpSuffix{text: text}
}

func (o *OpSuffix) Type() FieldType {
	switch strings.ToLower(o.text) {
	case "suffix":
		return TypeOpSuffix
	}
	return TypeUnknown
}

func (o *OpSuffix) Raw() string {
	return o.text
}

func (o *OpSuffix) Value() interface{} {
	return "suffix"
}

type OpPrefix struct {
	text string
}

func NewOpPrefix(text string) *OpPrefix {
	return &OpPrefix{text: text}
}

func (o *OpPrefix) Type() FieldType {
	switch strings.ToLower(o.text) {
	case "prefix":
		return TypeOpPrefix
	}
	return TypeUnknown
}

func (o *OpPrefix) Raw() string {
	return o.text
}

func (o *OpPrefix) Value() interface{} {
	return "prefix"
}

type OpContain struct {
	text string
}

func NewOpContain(text string) *OpContain {
	return &OpContain{text: text}
}

func (o *OpContain) Type() FieldType {
	switch strings.ToLower(o.text) {
	case "contain":
		return TypeOpContain
	}
	return TypeUnknown
}

func (o *OpContain) Raw() string {
	return o.text
}

func (o *OpContain) Value() interface{} {
	return "contain"
}

type OpLE struct {
	text string
}

func NewOpLE(text string) *OpLE {
	return &OpLE{text: text}
}

func (o *OpLE) Type() FieldType {
	switch strings.ToLower(o.text) {
	case "le", "<=":
		return TypeOpLE
	}
	return TypeUnknown
}

func (o *OpLE) Raw() string {
	return o.text
}

func (o *OpLE) Value() interface{} {
	return "<="
}

type OpGE struct {
	text string
}

func NewOpGE(text string) *OpGE {
	return &OpGE{text: text}
}

func (o *OpGE) Type() FieldType {
	switch strings.ToLower(o.text) {
	case "ge", ">=":
		return TypeOpGE
	}
	return TypeUnknown
}

func (o *OpGE) Raw() string {
	return o.text
}

func (o *OpGE) Value() interface{} {
	return ">="
}

type OpLT struct {
	text string
}

func NewOpLT(text string) *OpLT {
	return &OpLT{text: text}
}

func (o *OpLT) Type() FieldType {
	switch strings.ToLower(o.text) {
	case "gt", "<":
		return TypeOpLT
	}
	return TypeUnknown
}

func (o *OpLT) Raw() string {
	return o.text
}

func (o *OpLT) Value() interface{} {
	return "<"
}

type OpGT struct {
	text string
}

func NewOpGT(text string) *OpGT {
	return &OpGT{text: text}
}

func (o *OpGT) Type() FieldType {
	switch strings.ToLower(o.text) {
	case "gt", ">":
		return TypeOpGT
	}
	return TypeUnknown
}

func (o *OpGT) Raw() string {
	return o.text
}

func (o *OpGT) Value() interface{} {
	return ">"
}

type OpNE struct {
	text string
}

func NewOpNE(text string) *OpNE {
	return &OpNE{text: text}
}

func (o *OpNE) Type() FieldType {
	switch strings.ToLower(o.text) {
	case "ne", "!=":
		return TypeOpNE
	}
	return TypeUnknown
}

func (o *OpNE) Raw() string {
	return o.text
}

func (o *OpNE) Value() interface{} {
	return "!="
}

type OpEQ struct {
	text string
}

func NewOpEQ(text string) *OpEQ {
	return &OpEQ{text: text}
}

func (o *OpEQ) Type() FieldType {
	switch strings.ToLower(o.text) {
	case "eq", "==":
		return TypeOpEQ
	}
	return TypeUnknown
}

func (o *OpEQ) Raw() string {
	return o.text
}

func (o *OpEQ) Value() interface{} {
	return "=="
}

type Logical struct {
	text string
}

func NewLogical(text string) *Logical {
	return &Logical{text: text}
}

func (o *Logical) Type() FieldType {
	switch strings.ToLower(o.text) {
	case "and":
		return TypeLogicalAnd
	case "or":
		return TypeLogicalOr
	case "not":
		return TypeLogicalNot
	}
	return TypeUnknown
}

func (o *Logical) Raw() string {
	return o.text
}

func (o *Logical) Value() interface{} {
	return o.text
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

type ValueListDoubles struct {
	text string
}

func NewValueListDoubles(text string) *ValueListDoubles {
	return &ValueListDoubles{text: text}
}

func (o *ValueListDoubles) Type() FieldType {
	return TypeValueListDoubles
}

func (o *ValueListDoubles) Raw() string {
	return o.text
}

func (o *ValueListDoubles) Value() interface{} {
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

type ValueDouble struct {
	text string
}

func NewValueDouble(text string) *ValueDouble {
	return &ValueDouble{text: text}
}

func (o *ValueDouble) Type() FieldType {
	return TypeValueDouble
}

func (o *ValueDouble) Raw() string {
	return o.text
}

func (o *ValueDouble) Value() interface{} {
	ret, _ := strconv.ParseFloat(o.text, 64)
	return ret
}

type ValueNull struct {
	text string
}

func NewValueNull(text string) *ValueNull {
	return &ValueNull{text: text}
}

func (o *ValueNull) Type() FieldType {
	return TypeValueNull
}

func (o *ValueNull) Raw() string {
	return o.text
}

func (o *ValueNull) Value() interface{} {
	return o.text
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
