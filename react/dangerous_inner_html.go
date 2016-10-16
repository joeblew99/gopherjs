package react

import "github.com/gopherjs/gopherjs/js"

type DangerousInnerHTMLDef struct {
	o *js.Object
}

func DangerousInnerHTML(s string) *DangerousInnerHTMLDef {
	o := object.New()
	o.Set("__html", s)

	res := &DangerousInnerHTMLDef{o: o}

	return res
}

func (d *DangerousInnerHTMLDef) reactElement() {}
