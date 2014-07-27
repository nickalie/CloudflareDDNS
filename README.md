Cloudflare Dynamic DNS Client
==============


Classic Dynamic DNS clients based on IPv4. Currently even white dynamic IPv4 is quite rare thing. More and more providers use NAT and it's impossible to use classic dynamic dns to expose your home, testing or virtual server to the internet. Fortunately there is a lot of free IPv6 addresses (about 300 000 000 for each person on Earth). It's a good idea to use IPv6 instead of IPv4. But IPv4 is still most popular technology. The problem is IPv4 has no direct access to IPv6 and vice versa. [Cloudflare][1] provides transparent bridge from IPv4 to IPv6 for free. So IPv4 clients able to access IPv6 resources. The easiest way to get "white" IPv6 is [teredo][2]. Installation for Ubuntu:

```
sudo apt-get install miredo
```

That's all. Now your Ubuntu system has access to IPv6 resources over the internet. Now we want to expose local webserver globaly.

###Requirements
* Cloudflare account (it's free)
* Own domain that uses Cloudflare as DNS service
* Enabled IPv6 support in domain settings (Settings->Cloudflate settings->Automatic IPv6=Full
* API Key and Email from Cloudflare account settings

###Usage

Grab latest release from [here][3] or [build cloudflareddns from sources](#building-from-source).

```
./cloudflareddns -token="api_key_from_account_settings" -email="email_from_account_settings" -domain="mysubdomain.domain.com"
```

Thats it. Now your local web server will be available via web browser by http://mysubdomain.domain.com
With "-ipv4only=true" option cloudflareddns will update only IPv4 address.

###Building from source
1. Install [golang][4]
2. ```git clone https://github.com/nickalie/CloudflareDDNS.git```
3. ```cd CloudflareDDNS/src```
4. ```go build -o cloudflareddns main.go```
5. cloudflareddns ready to use

If run on Windows then use
```
go build -o cloudflareddns.exe main.go
```


  [1]: https://www.cloudflare.com
  [2]: http://en.wikipedia.org/wiki/Teredo_tunneling
  [3]: https://github.com/nickalie/CloudflareDDNS/releases
  [4]: http://golang.org/
