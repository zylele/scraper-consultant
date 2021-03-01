package main

import (
    "fmt"
    "github.com/bitly/go-simplejson"
    "github.com/gocolly/colly"
    "time"
)


func main() {
    // Instantiate default collector
    c := colly.NewCollector()
    c.OnResponse(func(response *colly.Response) {
        newJson, err := simplejson.NewJson(response.Body)
        if err != nil{

        }
        array := newJson.Get("rows").MustArray()
        for _, a := range array {
            stock,_ := a.(map[string]interface{})
            if stock["id"].(string) == "163406" {
                cell := stock["cell"].(map[string]interface{})
                fmt.Println(cell["fund_nm"])
                fmt.Println(cell["discount_rt"])
            }
        }
    })
    url := fmt.Sprintf("https://www.jisilu.cn/data/lof/stock_lof_list/?___jsl=LST___t=%s&rp=25&page=1", time.Now().UnixNano())
    err := c.Visit(url)
    if err != nil {
        fmt.Println(err)
    }
}