package main

import (
	"context"
	"flag"
	"github.com/erply/pim-go-wrapper/pkg/pim"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
)

func main() {
	var (
		sess    = flag.String("sessionKey", "sess", "session key")
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
		ProductAttributes:           &pim.ProductAttributes{},
	}

	var ps []pim.Product
	for i := 0; i < 10; i++ {
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
		//let the request fail half of the times
		if i%2 == 0 {
			p.GroupID = 2
		} else {
			p.GroupID = 3
		}

		p.TaxFree = 1
		p.RewardPointsNotAllowed = 1
		ps = append(ps, *p)
	}

	resp, httpResp, err := cli.Products.CreateBulk(ctx, ps)
	if err != nil {
		logrus.Error(err)
		return
	}
	if httpResp.StatusCode != http.StatusCreated {
		logrus.Error(httpResp.StatusCode)
		return
	}
	var createdProducts []pim.Product
	for i, res := range resp.Results {
		if res.ID != 0 {
			ps[i].ID = res.ID
			createdProducts = append(createdProducts, ps[i])
		}
	}
	var updateTypeReqs []pim.UpdateProductTypeBulkRequest
	for i := range createdProducts {
		if i%2 == 0 {
			//changing a variation into a matrix product is not allowed
			createdProducts[i].Type = "MATRIX"
			updateTypeReqs = append(updateTypeReqs, pim.UpdateProductTypeBulkRequest{
				ID:                       uint(createdProducts[i].ID),
				UpdateProductTypeRequest: pim.UpdateProductTypeRequest{Type: "MATRIX"},
			})
		}
	}

	resp, httpResp, err = cli.Products.UpdateBulk(ctx, createdProducts)
	if err != nil {
		logrus.Error(err)
		return
	}
	if httpResp.StatusCode != http.StatusOK {
		logrus.Error(httpResp.StatusCode)
		return
	}
	logrus.Info(resp)

	resp, httpResp, err = cli.Products.UpdateTypeBulk(ctx, updateTypeReqs)
	if err != nil {
		logrus.Error(err)
		return
	}
	if httpResp.StatusCode != http.StatusOK {
		logrus.Error(httpResp.StatusCode)
		return
	}
	logrus.Info(resp)

	//clean up
	for _, p := range createdProducts {
		_, httpResp, err := cli.Products.Delete(ctx, p.ID)
		if err != nil {
			logrus.Error(err)
			return
		}
		if httpResp.StatusCode != http.StatusOK {
			logrus.Error(httpResp.StatusCode)
			return
		}
	}

}
