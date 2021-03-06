package ui

// Code generated by webgen. DO NOT EDIT.

import (
	"github.com/gowebapi/webapi"
	"github.com/gowebapi/webapi/dom"
	"github.com/gowebapi/webapi/html"
	"github.com/gowebapi/webapi/html/canvas"
	"github.com/gowebapi/webapi/html/media"
)

type (
	_ *webapi.Document // prevent unused import errors
	_ *dom.Element
	_ *html.HTMLDivElement
	_ *canvas.HTMLCanvasElement
	_ *media.HTMLAudioElement
)

var (
	_document = webapi.GetDocument()
)

// source: testdata/standalone/ref.html

type ref struct {
	readme *html.HTMLAnchorElement
	roots  []*dom.Element
}

func newRef() *ref {
	a0 := _document.CreateElement("a", nil)
	stringliteral0 := "README"
	a0.SetTextContent(&stringliteral0)
	return &ref{
		readme: html.HTMLAnchorElementFromJS(a0),
		roots:  []*dom.Element{a0},
	}
}

func (v *ref) Roots() []*dom.Element {
	return v.roots
}
