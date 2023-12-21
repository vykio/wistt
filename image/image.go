package image

import (
	"bytes"
	"context"
	"embed"
	"github.com/chromedp/chromedp"
	"html/template"
	"time"
)

//go:embed templates/*
var content embed.FS

type Image struct {
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
		chromedp.Navigate("data:text/html,"+b.String()),
		chromedp.WaitReady("pre"),
		chromedp.Screenshot("#output", &buf),
	); err != nil {
		return nil, err
	}

	return buf, nil
}
