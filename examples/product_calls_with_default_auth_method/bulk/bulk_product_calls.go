package main

import (
	"context"
	"encoding/json"
	"flag"
	"github.com/erply/pim-go-wrapper/pkg/pim"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"time"
)

func main() {
	var (
		sess             = flag.String("sessionKey", "sess", "session key")
		cc               = flag.String("clientCode", "123456", "client code")
		baseUrl          = flag.String("baseUrl", "https://xyz/v1/", "base URL with version and slash")
		productsToCreate = flag.Int("amount-of-products", 0, "amount of products to create")
	)
	flag.Parse()
	var (
		tp         = pim.NewDefaultAuthTransport(*sess, *cc, nil)
		baseURL, _ = url.Parse(*baseUrl)
		cli        = pim.NewAPIClient(baseURL, tp.Client(), "bulk_caller")
		ctx        = context.Background()
	)

	p := &pim.Product{}

	p.Type = "PRODUCT"
	p.GroupID = 3
	p.TranslatableNameJSON = pim.TranslatableNameJSON{Name: map[string]string{}}
	p.TranslatableDescriptionJSON = pim.TranslatableDescriptionJSON{Description: map[string]pim.ProductDescription{}}
	p.ProductAttributes = &pim.ProductAttributes{}

	for i := 0; i < *productsToCreate/100; i++ {
		products := []pim.Product{}

		for i := 0; i < 100; i++ {
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
			p.GroupID = 3

			p.TaxFree = 1
			p.RewardPointsNotAllowed = 1
			products = append(products, *p)
		}

		logrus.Info("creating in bulk")
		start := time.Now()
		resp, httpResp, err := cli.Products.CreateBulk(ctx, products)
		if err != nil {
			logrus.Error("create ", err)
			return
		}
		logrus.Infof("created %d products in %d ms", len(products), time.Now().Sub(start).Milliseconds())
		if httpResp.StatusCode != http.StatusOK {
			logrus.Error("create ", httpResp.Status)
			var resp []byte
			if _, err := httpResp.Body.Read(resp); err == nil {
				logrus.Error(string(resp))
			}
			return
		}
		j, _ := json.Marshal(resp.Results)
		logrus.Info(string(j))
	}
}
