package main

import (
	"context"
	"flag"
	"github.com/erply/pim-go-wrapper/pkg/pim"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"time"
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
		cli        = pim.NewAPIClient(baseURL, tp.Client(), "bulk_caller")
		ctx        = context.Background()
	)

	p := &pim.Product{}

	p.Type = "PRODUCT"
	p.GroupID = 3
	p.TranslatableNameJSON = pim.TranslatableNameJSON{Name: map[string]string{}}
	p.TranslatableDescriptionJSON = pim.TranslatableDescriptionJSON{Description: map[string]pim.ProductDescription{}}
	p.ProductAttributes = &pim.ProductAttributes{}

	var products []pim.Product
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
		//let the request fail half of the times
		if i%2 == 0 {
			p.GroupID = 2
		} else {
			p.GroupID = 3
		}

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
		logrus.Error("create ", httpResp.StatusCode)
		var resp []byte
		if _, err := httpResp.Body.Read(resp); err == nil {
			logrus.Error(string(resp))
		}
		return
	}
	var createdProducts []pim.Product
	for i, res := range resp.Results {
		if res.ResourceID != 0 {
			products[i].ID = res.ResourceID
			createdProducts = append(createdProducts, products[i])
		}
	}
	var (
		updateTypeReqs []pim.UpdateProductTypeBulkRequest
		updateReqs     []pim.BulkUpdateProductRequestItem
	)

	for i, createdProduct := range createdProducts {
		id := uint(createdProduct.ID)
		logrus.Info("created product ID: ", id)
		if i%2 == 0 {
			//changing a variation into a matrix product is not allowed
			u := pim.BulkUpdateProductRequestItem{
				ResourceID: id,
			}
			u.Type = "MATRIX"
			updateReqs = append(updateReqs, u)
			updateTypeReqs = append(updateTypeReqs, pim.UpdateProductTypeBulkRequest{
				ResourceID:               id,
				UpdateProductTypeRequest: pim.UpdateProductTypeRequest{Type: u.Type},
			})
		}
	}

	logrus.Info("updating in bulk")
	resp, httpResp, err = cli.Products.UpdateBulk(ctx, updateReqs)
	if err != nil {
		logrus.Error("update", err)
		return
	}
	if httpResp.StatusCode != http.StatusOK {
		logrus.Error("update", httpResp.StatusCode)
		return
	}

	logrus.Info("updating types in bulk")
	resp, httpResp, err = cli.Products.UpdateTypeBulk(ctx, updateTypeReqs)
	if err != nil {
		logrus.Error("update type", err)
		return
	}
	if httpResp.StatusCode != http.StatusOK {
		logrus.Error("update type", httpResp.StatusCode)
		return
	}

	typeFilter, err := pim.NewFilter("type", "=", "PRODUCT", "")
	if err != nil {
		logrus.Error(err)
		return
	}

	logrus.Info("reading in bulk")
	bulkReadResponse, httpResp, err := cli.Products.ReadBulk(ctx, []pim.ListOptions{
		{
			Filters: []pim.Filter{
				*typeFilter,
			},
			PaginationParameters: nil,
			SortingParameter:     nil,
			WithTotalCount:       true,
		},
		{},
	})
	if err != nil {
		logrus.Error(err)
	}
	if httpResp.StatusCode != http.StatusOK {
		logrus.Error(httpResp.Status)
		return
	}

	for _, item := range bulkReadResponse.Results {
		logrus.Info("result ID: ", item.ResultID, " , total count: ", item.TotalCount, ", total product records: ", len(item.Products))
	}

	logrus.Info("deleting the items")
	deleteStart := time.Now()
	for _, p := range createdProducts {
		_, httpResp, err := cli.Products.Delete(ctx, p.ID)
		if err != nil {
			logrus.Error("delete", err)
			return
		}
		if httpResp.StatusCode != http.StatusOK {
			logrus.Error("delete", httpResp.StatusCode)
			return
		}
	}
	logrus.Infof("deleted %d products in %d s", len(products), time.Now().Sub(deleteStart).Seconds())

}
