// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package inputtext

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
	opts        []OptFn
}

var styleguideItems = []styleguideItem{
	{
		title: "Default",
		usage: `
@inputtext.New(
    inputtext.WithID("zing"),
    inputtext.WithClasses("foo"),
    inputtext.WithAttributes(map[string]any{"data-zed": "zung"}),
    // inputtext.WithDisabled(),
    inputtext.WithRequired(),
    inputtext.WithName("curr-salary"),
    inputtext.WithValue(500),
    inputtext.WithPlaceholder("Set your salary"),
    func(c *inputtext.InputText) {
        c.ClassNames = append(c.ClassNames, "bar")
    },
)
    `,
		opts: []OptFn{
			WithID("zing"),
			WithClasses("foo"),
			WithAttributes(map[string]any{"data-zed": "zung"}),
			// WithDisabled(),
			WithRequired(),
			WithName("curr-salary"),
			WithValue(500),
			WithPlaceholder("Set your salary"),
			func(c *InputText) {
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
