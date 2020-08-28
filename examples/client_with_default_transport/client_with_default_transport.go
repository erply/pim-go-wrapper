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
		opts       = pim.NewListOptions(nil, nil, nil)
	)

	locations, httpResp, err := cli.WarehouseLocations.Get(ctx, opts)
	if err != nil {
		logrus.Error(err)
		return
	}
	logrus.Info(locations)
	logrus.Info(httpResp.Status)
}
