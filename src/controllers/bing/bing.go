package bing

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const SUBSCRIPTION_KEY string = "d065001678af4760a0809e2d32810702"
const ENDPOINT string = "https://api.bing.microsoft.com/v7.0/images/search"

func GetImages(q string) Bing_res{
	req, _ := http.NewRequest("GET", ENDPOINT, nil)
	req.Header.Set("Ocp-Apim-Subscription-Key", SUBSCRIPTION_KEY)
	
	params := req.URL.Query()
	params.Add("q", q)
	params.Add("count", "12")
	req.URL.RawQuery = params.Encode()

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil{
		log.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	var ans Bing_res
	err = json.Unmarshal(body, &ans)
	if err != nil {
		log.Println(err)
	}
	return ans
}
