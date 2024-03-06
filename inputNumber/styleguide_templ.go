// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package inputNumber

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"github.com/iamajoe/templui/utils"
)

type styleguideItem struct {
	title       string
	description string
	usage       string
	opts        []OptsFn
}

var styleguideItems = []styleguideItem{
	{
		title: "Default",
		usage: `
@inputNumber.New(
    inputNumber.WithID("zing"),
    inputNumber.WithClasses("foo"),
    inputNumber.WithAttributes(map[string]any{"data-zed": "zung"}),
    // inputNumber.WithDisabled(),
    inputNumber.WithRequired(),
    inputNumber.WithName("curr-salary"),
    inputNumber.WithValue(500),
    inputNumber.WithPlaceholder("Set your salary"),
    inputNumber.WithMin(100),
    inputNumber.WithMax(8000),
    inputNumber.WithStep(10),
    func(c *inputNumber.InputNumber) {
        c.ClassNames = append(c.ClassNames, "bar")
    },
)
    `,
		opts: []OptsFn{
			WithID("zing"),
			WithClasses("foo"),
			WithAttributes(map[string]any{"data-zed": "zung"}),
			// WithDisabled(),
			WithRequired(),
			WithName("curr-salary"),
			WithValue(500),
			WithPlaceholder("Set your salary"),
			WithMin(100),
			WithMax(8000),
			WithStep(10),
			func(c *InputNumber) {
				c.ClassNames = append(c.ClassNames, "bar")
			},
		},
	},
}

func Styleguide() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		for _, component := range styleguideItems {
			templ_7745c5c3_Var2 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
				templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
				if !templ_7745c5c3_IsBuffer {
					templ_7745c5c3_Buffer = templ.GetBuffer()
					defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
				}
				templ_7745c5c3_Err = New(component.opts...).Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				if !templ_7745c5c3_IsBuffer {
					_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
				}
				return templ_7745c5c3_Err
			})
			templ_7745c5c3_Err = utils.Item(component.title, component.description, component.usage).Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
