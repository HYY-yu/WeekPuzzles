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

func ApiToServer(saturdays saturday.Saturdays) {
	apiUrl := "http://lanfly.vicp.io/api/holiday/batch?d=$date"

	client := &http.Client{
	}

	newDate, dateMap := getDateStringAndMap(saturdays)

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

	for k, v := range dateMap {
		pathStr := "holiday." + k
		holi := respJson.Get(pathStr)
		holidayExist := holi.IsObject()
		if holidayExist {
			holiday := holi.Get("holiday").Bool()
			saturdays[v].Low = !holiday
			saturdays[v].Relax = holiday
		}
	}
	return
}

func getDateStringAndMap(sas saturday.Saturdays) (string, map[string]int) {
	temp := make([]string, 0, len(sas))
	tempMap := make(map[string]int)

	for i, sa := range sas {
		date := sa.Date[:4] + "-" + sa.Date[4:6] + "-" + sa.Date[6:]
		temp = append(temp, date)
		tempMap[date] = i
	}
	return strings.Join(temp, ","), tempMap
}
