package image

import (
	"bytes"
	"context"
	"embed"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"html/template"
	"time"
)

//go:embed templates/*
var content embed.FS

type Image struct {
	Input  string
	Output string
}

func GenerateBuffer(image Image) ([]byte, error) {
	var buf []byte

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, time.Second)
	defer cancel()

	templateContent, err := content.ReadFile("templates/simple.tpl")
	if err != nil {
		return nil, err
	}

	tmpl, err := template.New("image").Parse(string(templateContent))
	if err != nil {
		return nil, err
	}

	var b bytes.Buffer
	if err := tmpl.Execute(&b, image); err != nil {
		return nil, err
	}

	if err := chromedp.Run(ctx,
		chromedp.Navigate("about:blank"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			frameTree, err := page.GetFrameTree().Do(ctx)
			if err != nil {
				return err
			}
			return page.SetDocumentContent(frameTree.Frame.ID, b.String()).Do(ctx)
		}),
		chromedp.WaitReady("pre"),
		chromedp.Screenshot("#container", &buf),
	); err != nil {
		return nil, err
	}

	return buf, nil
}
