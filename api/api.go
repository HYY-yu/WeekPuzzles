package api

import (
	"net/http"
	"net/url"
	"WeekPuzzles/saturday"
	"strings"
	"io/ioutil"
	"github.com/tidwall/gjson"
	"log"
)

func ApiToServer(sa *saturday.Saturday) {
	apiUrl := "http://lanfly.vicp.io/api/holiday/info/$date"

	client := &http.Client{
	}

	newDate := sa.Date[:4] + "-" + sa.Date[4:6] + "-" + sa.Date[6:]

	apiUrl = strings.Replace(apiUrl, "$date", newDate, -1)
	u, _ := url.Parse(apiUrl)
	req, _ := http.NewRequest("GET", u.String(), nil)
	req.Header.Add("User-Agent", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36`)
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.8")

	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		log.Println("网络连接失败", err)
		return
	}

	robots, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Println("Body解析错误", err)
		return
	}

	respJson := gjson.ParseBytes(robots)
	code := respJson.Get("code").Int()

	if code != 0 {
		log.Println("服务器出错")
		return
	}

	holidayExist := respJson.Get("holiday.holiday").Exists()
	if holidayExist {
		holiday := respJson.Get("holiday.holiday").Bool()
		sa.Low = !holiday
		sa.Relax = holiday
	}
	return
}
