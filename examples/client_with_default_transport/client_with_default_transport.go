package main

import (
	"context"
	"github.com/erply/pim-go-wrapper/pkg/pim"
	"github.com/sirupsen/logrus"
	"net/url"
)

func main() {
	var (
		tp         = pim.NewDefaultAuthTransport("", "", nil)
		baseURL, _ = url.Parse("https://xyz.erply.com/v1/")
		cli        = pim.NewClient(baseURL, tp.Client())
		ctx        = context.Background()
		opts       = pim.NewListOptions(nil, nil, nil, false)
	)

	p := &pim.Product{
		Type:    "PRODUCT",
		GroupID: 3,
	}
	newID, _, err := cli.Products.Create(ctx, p)
	if err != nil {
		logrus.Error(err)
		return
	}

	columnFilter := [3]interface{}{"id", "=", newID}
	opts.Filters = append(opts.Filters, *pim.NewFilter(columnFilter, ""))
	products, _, err := cli.Products.Read(ctx, opts)
	if err != nil {
		logrus.Error(err)
		return
	}
	logrus.Info(products)
}
