package ssr

import (
	"context"

	"log"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/cdproto/fetch"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

type Renderer struct {
}

func New() *Renderer {
	return &Renderer{}
}

func (r *Renderer) Render(url string) ([]byte, error) {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	chromedp.ListenTarget(ctx, DisableImageLoad(ctx))

	// capture screenshot of an element
	var out string
	if err := chromedp.Run(ctx, fetch.Enable(), scrapIt(url, &out)); err != nil {
		log.Fatal(err)
	}

	return []byte(out), nil
}

func scrapIt(url string, str *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.ActionFunc(func(ctx context.Context) error {
			node, err := dom.GetDocument().Do(ctx)
			if err != nil {
				return err
			}
			*str, err = dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)
			return err
		}),
	}
}

func DisableImageLoad(ctx context.Context) func(event interface{}) {
	return func(event interface{}) {
		switch ev := event.(type) {
		case *fetch.EventRequestPaused:
			go func() {
				c := chromedp.FromContext(ctx)
				ctx := cdp.WithExecutor(ctx, c.Target)

				if ev.ResourceType == network.ResourceTypeImage {
					fetch.FailRequest(ev.RequestID, network.ErrorReasonBlockedByClient).Do(ctx)
				} else {
					fetch.ContinueRequest(ev.RequestID).Do(ctx)
				}
			}()
		}
	}
}
