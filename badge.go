package bootstrap4

import (
	"github.com/gowasm/vecty"
	"github.com/gowasm/vecty/elem"
	"github.com/gowasm/vecty/prop"
)

// Badge ...
type Badge struct {
	vecty.Core
	ID       string              `vecty:"prop"`
	Kind     Kind                `vecty:"prop"`
	Pill     bool                `vecty:"prop"`
	Markup   vecty.MarkupList    `vecty:"prop"`
	Children vecty.MarkupOrChild `vecty:"prop"`
}

// Render ...
func (c *Badge) Render() vecty.ComponentOrHTML {
	return elem.Span(
		vecty.Markup(
			vecty.MarkupIf(len(c.ID) > 0, prop.ID(c.ID)),
			vecty.Class("badge"),
			vecty.MarkupIf(len(c.Kind) == 0, vecty.Class("badge-"+KindPrimary.String())),
			vecty.MarkupIf(len(c.Kind) > 0, vecty.Class("badge-"+c.Kind.String())),
			vecty.MarkupIf(c.Pill, vecty.Class("badge-pill")),
		),
		c.Markup,
		c.Children,
	)
}

// BadgeLinks ...
type BadgeLinks struct {
	vecty.Core
	ID       string              `vecty:"prop"`
	Kind     Kind                `vecty:"prop"`
	Href     string              `vecty:"prop"`
	Pill     bool                `vecty:"prop"`
	Markup   vecty.MarkupList    `vecty:"prop"`
	Children vecty.MarkupOrChild `vecty:"prop"`
}

// Render ...
func (c *BadgeLinks) Render() vecty.ComponentOrHTML {
	return elem.Anchor(
		vecty.Markup(
			vecty.MarkupIf(len(c.ID) > 0, prop.ID(c.ID)),
			vecty.Class("badge"),
			vecty.MarkupIf(len(c.Kind) == 0, vecty.Class("badge-"+KindPrimary.String())),
			vecty.MarkupIf(len(c.Kind) > 0, vecty.Class("badge-"+c.Kind.String())),
			vecty.MarkupIf(c.Pill, vecty.Class("badge-pill")),
			vecty.MarkupIf(len(c.Href) > 0, prop.Href(c.Href)),
		),
		c.Markup,
		c.Children,
	)
}
