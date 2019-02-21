package controllers

import (
	"fmt"
	"github.com/TimeBye/amap/models"
	"github.com/astaxie/beego"
	"github.com/parnurzeal/gorequest"
	"strconv"
	"sync"
)

const (
	offset = 20
)

var (
	amap = gorequest.New()
	pois = make([]models.Poi, 0)
)

type MainController struct {
	beego.Controller
}

type AroundController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["APIKey"] = beego.AppConfig.String("apikey")
	c.TplName = "index.html"
}

func GetResult(p models.Parameters, page int, wg *sync.WaitGroup) {
	r := models.Result{}
	response, bytes, errors := amap.Get(fmt.Sprintf("https://restapi.amap.com/v3/place/around?"+
		"key=%s&location=%s&keywords=%s&types=%s&radius=%d&offset=%d&page=%d&extensions=all",
		beego.AppConfig.String("apikey"),
		p.Location,
		p.Keywords,
		p.Types,
		p.Radius,
		offset,
		page,
	)).EndStruct(&r)
	if errors != nil {
		beego.Error(errors, string(bytes), response)
	}
	for _, poi := range r.Pois {
		if v, ok := poi.Tel.(string); ok && len(v) > 3 {
			pois = append(pois, poi)
		}
	}
	wg.Done()
}

func GetResults(p models.Parameters) []models.Poi {
	pois = make([]models.Poi, 0)
	r := models.Result{}
	response, bytes, errors := amap.Get(fmt.Sprintf("https://restapi.amap.com/v3/place/around?"+
		"key=%s&location=%s&keywords=%s&types=%s&radius=%d&offset=%d&page=%d&extensions=all",
		beego.AppConfig.String("apikey"),
		p.Location,
		p.Keywords,
		p.Types,
		p.Radius,
		offset,
		1,
	)).EndStruct(&r)
	if errors != nil {
		beego.Error(errors, string(bytes), response)
	}
	for _, poi := range r.Pois {
		if v, ok := poi.Tel.(string); ok && len(v) > 3 {
			pois = append(pois, poi)
		}
	}
	rC, _ := strconv.Atoi(r.Count)
	pages := rC/offset + 1
	wg := sync.WaitGroup{}
	for page := 2; page < pages; page = page + 1 {
		wg.Add(1)
		go GetResult(p, page, &wg)
	}
	wg.Wait()
	return pois
}

func (c *AroundController) Get() {
	p := models.Parameters{}
	if err := c.ParseForm(&p); err == nil {
		r := GetResults(p)
		c.Data["Results"] = r
		c.TplName = "table.html"
	} else {
		c.Ctx.WriteString(err.Error())
	}
}
