package ng

import (
	"net/url"
	"net/http"
	"io/ioutil"
	"errors"
	"encoding/json"
)

type Api struct {
	config Config
}

func NewApi(config Config) Api {
	return Api{config}
}

func (this *Api) RecLoadAll(domain string) (RecLoadAllResponse, error) {
	result, err := this.callMethod("rec_load_all", map[string]string{"z":domain})

	if err != nil {
		return nil, err
	} else {
		return result["response"].(map[string]interface{})["recs"].(map[string]interface{}), nil
	}
}

func (this *Api) RecNew(domain string, name string, ip string, ipType string) (ObjVO, error) {
	result, err := this.callMethod("rec_new", map[string]string{"z":domain, "name":name, "content":ip, "type":ipType, "ttl":"1"})

	if err == nil {
		return result["response"].(map[string]interface{})["rec"].(map[string]interface{})["obj"].(map[string]interface{}), nil
	} else {
		return nil, err
	}
}

func (this *Api) RecEdit(domain string, name string, id string, ip string, ipType string) error {
	_, err := this.callMethod("rec_edit", map[string]string{"z":domain, "name":name, "id":id, "content":ip, "type":ipType, "ttl":"1", "service_mode":"1"})
	return err
}

func (this *Api) RecDelete(domain string, id string) error {
	_, err := this.callMethod("rec_delete", map[string]string{"z":domain, "id":id})
	return err
}

func (this *Api) callMethod(method string, params map[string]string) (map[string]interface{}, error) {
	methodParams := this.getParams(params)
	methodParams.Set("a", method)
	resp, err := http.PostForm("https://www.cloudflare.com/api_json.html", methodParams)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	json.Unmarshal(body, &result)
	if result["result"].(string) == "success" {
		return result, nil
	} else {
		return nil, errors.New(result["msg"].(string))
	}
}

func (this *Api) getParams(params map[string]string) url.Values {
	result := url.Values{"tkn": {this.config.Token}, "email": {this.config.Email}}

	for key, value := range params {
		result.Set(key, value)
	}

	return result
}
