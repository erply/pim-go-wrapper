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
		cli        = pim.NewAPIClient(baseURL, tp.Client(), "bulk_caller")
		ctx        = context.Background()
	)

	names := []string{"attr1", "attr2", "attr3", "attr4", "attr5", "attr6", "attr7", "attr8", "attr9", "attr10"}
	var attrReqs []pim.AttributeRequest
	startingProductID := 545265
	for id := startingProductID; id < startingProductID+99999; id++ {
		for _, name := range names {
			attr := pim.AttributeRequest{
				RecordID: id,
				Entity:   "product",
				Type:     "text",
				Name:     name,
				Value:    "blablabla",
			}
			attrReqs = append(attrReqs, attr)
		}
	}

	for _, a := range attrReqs {
		temp := a
		//start := time.Now()
		_, httpResp, err := cli.Attributes.Attach(ctx, &temp)
		if err != nil {
			logrus.Error("attach ", err)
			return
		}
		//logrus.Infof("created %d products in %d ms", len(products), time.Now().Sub(start).Milliseconds())
		if httpResp.StatusCode != http.StatusCreated {
			logrus.Error("attach ", httpResp.Status)
			var resp []byte
			if _, err := httpResp.Body.Read(resp); err == nil {
				logrus.Error(string(resp))
			}
		}
		//j, _ := json.Marshal(resp.ID)
		if a.Name == "attr10" {
			logrus.Infof("created attrs for product ID %d", a.RecordID)
		}
	}
}
