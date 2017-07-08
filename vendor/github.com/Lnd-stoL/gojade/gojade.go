package gojade

import (
	"bytes"
	"github.com/Lnd-stoL/gojade/jadeparser"
	"io"
	"reflect"
)

// Engine keeps configuration information and redirect calls to the jadeparser.
type Engine struct {
	ViewPath string
	Beautify bool
	extfunc  map[string]reflect.Value
}

// Creates a new instance of the jade instance struct.
func New() *Engine {
	gojade := new(Engine)
	gojade.extfunc = make(map[string]reflect.Value)
	return gojade
}

// RenderFile Renders a jade file to a bytes.Buffer.
// Example: jade.RenderFile("index.jade",nil).String()
func (this *Engine) RenderFile(filename string, data interface{}) *bytes.Buffer {
	buf := new(bytes.Buffer)
	eval := this.init(buf)
	eval.SetData(data)
	eval.RenderFile(filename)
	return buf
}

// REnderString Renders a jade string to html.
func (this *Engine) RenderString(template string, data interface{}) *bytes.Buffer {
	buf := new(bytes.Buffer)
	eval := this.init(buf)
	eval.SetData(data)
	eval.RenderString(template)
	return buf
}

// RenderFileW Render a jade file to a io.writer stream.
func (this *Engine) RenderFileW(wr io.Writer, template string, data interface{}) error {
	eval := this.init(wr)
	eval.SetData(data)
	eval.RenderFile(template)
	return nil
}

// RegisterFunction registers a function tobe called from your jade template.
func (this *Engine) RegisterFunction(name string, fn interface{}) {
	fnvalue := reflect.ValueOf(fn)
	switch fnvalue.Kind() {
	case reflect.Func:
		this.extfunc[name] = fnvalue
	default:
		panic("argument 'fn' is not a function. " + fnvalue.String())
	}
}

func (this *Engine) init(writer io.Writer) *jadeparser.EvalJade {
	eval := jadeparser.NewEvalJade(writer)
	eval.SetViewPath(this.ViewPath)
	eval.Beautify = this.Beautify
	eval.Extfunc = this.extfunc
	return eval
}
