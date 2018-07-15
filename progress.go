package bootstrap4

import (
	"github.com/gowasm/vecty"
	"github.com/gowasm/vecty/elem"
	"github.com/gowasm/vecty/prop"
)

// Progress ...
type Progress struct {
	vecty.Core
	ID       string                `vecty:"prop"`
	Height   string                `vecty:"prop"`
	Markup   vecty.MarkupList      `vecty:"prop"`
	Children vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *Progress) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.MarkupIf(len(c.ID) > 0, prop.ID(c.ID)),
			vecty.Class("progress"),
			vecty.MarkupIf(len(c.Height) > 0, vecty.Style("height", c.Height)),
		),
		c.Markup,
		c.Children,
	)
}

// ProgressBar ...
type ProgressBar struct {
	vecty.Core
	ID       string                `vecty:"prop"`
	Kind     Kind                  `vecty:"prop"`
	Width    string                `vecty:"prop"`
	Striped  bool                  `vecty:"prop"`
	Animated bool                  `vecty:"prop"`
	Markup   vecty.MarkupList      `vecty:"prop"`
	Children vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *ProgressBar) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.MarkupIf(len(c.ID) > 0, prop.ID(c.ID)),
			vecty.ClassMap{
				"progress-bar":          true,
				"progress-bar-striped":  c.Striped,
				"progress-bar-animated": c.Animated,
				"bg-" + c.Kind.String(): len(c.Kind) > 0,
			},
			vecty.Attribute("role", "progress-bar"),
			vecty.MarkupIf(len(c.Width) > 0, vecty.Style("width", c.Width)),
		),
		c.Markup,
		c.Children,
	)
}
