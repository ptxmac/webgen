package ui

// Code generated by nausicaa. DO NOT EDIT.

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

// source: testdata/standalone/specificElement.html

type specificElement struct {
	roots []*dom.Element
}

func newSpecificElement() *specificElement {
	div0 := _document.CreateElement("div", nil)
	audio0 := _document.CreateElement("audio", nil)
	div0.AppendChild(&audio0.Node)
	return &specificElement{
		roots: []*dom.Element{div0},
	}
}

func (v *specificElement) Roots() []*dom.Element {
	return v.roots
}
