package ng

import (
	"net/http"
	"io/ioutil"
)

func GetIpv6() string{
	return getIp("http://v6.ipv6-test.com/api/myip.php")
}

func GetIpv4() string{
	return getIp("http://v4.ipv6-test.com/api/myip.php")
}

func getIp(url string) string{
	resp,err:=http.Get(url)

	if (err != nil){
		return ""
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return ""
	}

	return string(body)
}
