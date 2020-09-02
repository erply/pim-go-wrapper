package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/erply/pim-go-wrapper/pkg/pim"
	"github.com/sirupsen/logrus"
	"net/url"
)

func main() {
	var (
		sess    = flag.String("sessionKey", "123123sdasd123", "session key")
		cc      = flag.String("clientCode", "123456", "client code")
		baseUrl = flag.String("baseUrl", "https://xyz/v1/", "base URL with version and slash")
	)
	flag.Parse()
	var (
		tp         = pim.NewDefaultAuthTransport(*sess, *cc, nil)
		baseURL, _ = url.Parse(*baseUrl)
		cli        = pim.NewClient(baseURL, tp.Client())
		ctx        = context.Background()
	)

	p := &pim.Product{
		Type:                        "PRODUCT",
		GroupID:                     3,
		TranslatableNameJSON:        pim.TranslatableNameJSON{Name: map[string]string{}},
		TranslatableDescriptionJSON: pim.TranslatableDescriptionJSON{Description: map[string]pim.ProductDescription{}},
		Attributes:                  &pim.Attributes{},
	}

	var ps []pim.Product
	for i := 0; i < 10*100; i++ {
		name := "name"
		desc := "desc"
		p.TranslatableNameJSON.Name["el"] = name
		p.TranslatableNameJSON.Name["en"] = name
		p.TranslatableNameJSON.Name["et"] = name
		//add descriptions
		p.TranslatableDescriptionJSON.Description["el"] = pim.ProductDescription{
			PlainText: desc,
			HTML:      desc,
		}
		p.TranslatableDescriptionJSON.Description["en"] = pim.ProductDescription{
			PlainText: desc,
			HTML:      desc,
		}
		p.TranslatableDescriptionJSON.Description["et"] = pim.ProductDescription{
			PlainText: desc,
			HTML:      desc,
		}
		//add some attributes
		p.PackagingType = "t"
		p.TaxFree = 1
		p.RewardPointsNotAllowed = 1
		ps = append(ps, *p)
	}

	resp, httpResp, err := cli.Products.CreateBulk(ctx, ps)
	if err != nil {
		logrus.Error(err)
		return
	}
	fmt.Println(httpResp.Status)
	fmt.Println(resp.IDs)

}
