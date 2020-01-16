package testing

import (
	"fmt"
	"os"

	"github.com/varunpurohit76/crawler/base"
	"github.com/varunpurohit76/crawler/sitemap"
)

func TestServicesInit(algorithm int, view int) {
	base.LoadConfig("./../config-test.json")
	if err := base.ConnectDb(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	base.InitLog()
	sitemap.Service.Init(algorithm, view)
}
