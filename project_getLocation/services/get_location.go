package services

import (
	"encoding/json"
	"getLocation/utils"
	"io/ioutil"
	
	"github.com/zxysilent/logs"
	"net/http"
)

func GetLocation(args string) utils.PCALocation {

	return func(str string) utils.PCALocation {
		args = str
		urlStr := "http://api.map.baidu.com/place/v2/search?query=" +
			args + "&region=中国&output=json&ak=ww0usZ7UsbDdfhIpQPdRtixP"
		// url:="http://api.map.baidu.com/place/v2/search?query=中山大学&region=中国&output=json&ak=ww0usZ7UsbDdfhIpQPdRtixP"
		resp, err := http.Get(urlStr)
		if err != nil {
			logs.Fatal(err)
		}

		//关闭流
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logs.Fatal(err)
		}

		//接收百度地图查询结果
		var mapRes utils.MapRes
		err = json.Unmarshal(body, &mapRes)
		if err != nil {
			logs.Fatalf("unmarshal err=%v\n", err)
		}

		// fmt.Println("mapRes.Results",len(mapRes.Results))
		var pcaLocation utils.PCALocation
		if len(mapRes.Results)!=0{
			pcaLocation = utils.PCALocation{
				Province: mapRes.Results[0].Province,
				City:     mapRes.Results[0].City,
				Area:     mapRes.Results[0].Area,
			}
		} else{
			pcaLocation = *new(utils.PCALocation)
		}
		return pcaLocation
	}(args)
}
