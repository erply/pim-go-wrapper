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
		TranslatableNameJSON: pim.TranslatableNameJSON{Name: map[string]string{
			"en": "blablabla",
		}},
	}
	resp, _, err := cli.Products.Create(ctx, p)
	if err != nil {
		logrus.Error(err)
		return
	}
	newProductID := resp.ID
	filter, err := pim.NewFilter("id", "=", newProductID, "")
	if err != nil {
		logrus.Error(err)
		return
	}
	opts.Filters = append(opts.Filters, *filter)
	products, _, err := cli.Products.Read(ctx, opts)
	if err != nil {
		logrus.Error(err)
		return
	}

	for _, p := range *products {
		logrus.Info(p.Name["en"])
	}

	p.NetWeight = 10
	p.Physical.GrossWeight = 10
	p.Physical.Length = 2
	p.Physical.Height = 2
	resp, _, err = cli.Products.Update(ctx, newProductID, p)
	if err != nil {
		logrus.Error(err)
		return
	}

	products, _, err = cli.Products.Read(ctx, opts)
	if err != nil {
		logrus.Error(err)
		return
	}

	for _, p := range *products {
		logrus.Info(p.NetWeight)
	}

	resp, httpDeleteResp, err := cli.Products.Delete(ctx, newProductID)
	if err != nil {
		logrus.Error(err)
		return
	}

	logrus.Info(httpDeleteResp.Status)
}
