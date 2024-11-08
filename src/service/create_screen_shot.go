package service

import (
	"context"
	"sync"

	"github.com/chromedp/chromedp"
)

var (
	ctx    context.Context
	cancel context.CancelFunc
	once   sync.Once
)

func initChromedpContext() {
	ctx, cancel = chromedp.NewContext(
		context.Background(),
	)
}

func CreateScreenShot(width, height int64, uri string) (htmlPng []byte, err error) {
	once.Do(initChromedpContext)

	err = chromedp.Run(ctx, chromedp.Tasks{
		chromedp.EmulateViewport(width, height),
		chromedp.Navigate(uri),
		chromedp.FullScreenshot(&htmlPng, 100),
	})
	if err != nil {
		cancel()
		initChromedpContext()
		return nil, err
	}

	return htmlPng, nil
}
