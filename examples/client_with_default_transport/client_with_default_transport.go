package client_with_default_transport

import (
	"context"
	"github.com/sirupsen/logrus"
	"my_workspace/sdks/pim-go-wrapper/pkg/pim"
	"net/http"
	"net/url"
)

func main() {
	tp := pim.NewDefaultAuthTransport("session", "123456", nil)
	baseURL, _ := url.Parse("https://pim-example.erply.com/")
	cli := pim.NewClient(baseURL, tp.Client())
	ctx := context.Background()

	resp := new(interface{})
	req, _ := cli.NewRequest(http.MethodGet, "v1/product", resp)
	httpResp, _ := cli.Do(ctx, req, resp)
	logrus.Info(httpResp)
}
