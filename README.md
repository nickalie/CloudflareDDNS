[![Build Status](https://travis-ci.org/nickalie/CloudflareDDNS.svg?branch=Config-file-support)](https://travis-ci.org/nickalie/CloudflareDDNS)
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
./cloudflareddns -token=api_key_from_account_settings -email=email_from_account_settings -domains=mysubdomain.domain.com
```

That's it. Now your local web server will be available via web browser by http://mysubdomain.domain.com
With "-ipv4only=true" option cloudflareddns will update only IPv4 address.
"domains" could contain several comma-separated domains.
Or you can create config.json file near cloudflareddns and specify all arguments there

```
{
  "token":"api_key_from_account_settings",
  "email":"email_from_account_settings",
  "domains":[
    "sub1.domain.com",
    "sub2.anotherdomain.com"
  ]
}
```

Use cron to run cloudflareddns periodicaly. For example, every hour:
```
crontab -e
0 * * * * /path/to/cloudflareddns arguments
```
To use cloudflareddns with [runit][5], [supervisord][6] or another process control system there is "interval" argument:

```
./cloudflareddns -token=api_key_from_account_settings -email=email_from_account_settings -domains=mysubdomain.domain.com -interval=30
```

or

```
{
  "token":"api_key_from_account_settings",
  "email":"email_from_account_settings",
  "domains":[
    "sub1.domain.com",
    "sub2.anotherdomain.com"
  ],
  "interval":30
}
```

It means cloudflareddns will run "infinitely" and update IPs every 30 minutes.


###Building from source
1. Install [golang][4]
2. ```git clone https://github.com/nickalie/CloudflareDDNS.git```
3. ```cd CloudflareDDNS/src```
4. ```go build -o cloudflareddns main.go```
5. cloudflareddns ready to use

Step 4 for Windows:
```
go build -o cloudflareddns.exe main.go
```


  [1]: https://www.cloudflare.com
  [2]: http://en.wikipedia.org/wiki/Teredo_tunneling
  [3]: https://github.com/nickalie/CloudflareDDNS/releases
  [4]: http://golang.org/
  [5]: http://smarden.org/runit/
  [6]: http://supervisord.org/
