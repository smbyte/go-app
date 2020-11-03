package main

import (
	"path/filepath"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

func newStart() app.UI {
	return newPage().
		Path("/web/documents/start.md").
		TableOfContents(
			"Getting started",
			"Prerequisite",
			"Install",
			"User interface",
			"Server",
			"Build and run",
			"Tips",
			"Next",
		)
}

func newArchitecture() app.UI {
	return newPage().
		Path("/web/documents/architecture.md").
		TableOfContents(
			"Architecture",
			"Web browser",
			"Server",
			"App",
			"Static resources",
			"Next",
		)
}

func newCompo() app.UI {
	return newPage().
		Path("/web/documents/components.md").
		TableOfContents(
			"Create",
			"Customize",
			"Update",
			"Lifecycle",
		)
}

func newSyntax() app.UI {
	return newPage().
		Path("/web/documents/syntax.md").
		TableOfContents(
			"HTML elements",
			"    Create",
			"    Standard elements",
			"    Self closing elements",
			"    Style",
			"    Attributes",
			"    Event handlers",
			"Text",
			"Raw elements",
			"Nested components",
			"Condition",
			"    If",
			"    ElseIf",
			"    Else",
			"Range",
			"    Slice",
			"    Map",
		)
}

type page struct {
	app.Compo

	path  string
	links []string
}

func newPage() *page {
	return &page{}
}

func (p *page) Path(v string) *page {
	p.path = v
	return p
}

func (p *page) TableOfContents(v ...string) *page {
	p.links = v
	return p
}

func (p *page) Render() app.UI {
	return app.Shell().
		Class("app-background").
		Menu(Menu()).
		Submenu(
			newTableOfContents().
				Links(p.links...),
		).
		OverlayMenu(Menu()).
		Content(
			newDocument(p.path).
				Description(filepath.Base(p.path)),
		)
}