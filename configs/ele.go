package configs

var (
	EleDomain = "h5.ele.me"
	EleUrl = "https://" + EleDomain
	SearchUrl = "/restapi/shopping/v2/restaurants/search?offset=0&limit=15&keyword=%s&latitude=%s&longitude=%s&search_item_type=3&is_rewrite=1&extras[]=activities&extras[]=coupon&terminal=h5"
	UserAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.67 Safari/537.36"
	
	Headers = map[string]string{
		"User-agent":UserAgent,
		"Connection":"keep-alive",
		"Host":EleDomain,
		"Accept-Encoding":"",
		"Cookie":"ubt_ssid=a8h1z9r8wq4yeb56z991jj1al88q4hah_2018-10-23; _utrace=97c0db9c29f6280ccf4f8d98d9d75483_2018-10-23; cna=wSdWEx2OCFkCAWVR5QwzamQ0; perf_ssid=2h3rek1csxhkiuw2kamjnzw9zdx593c9_2018-10-23; isg=BMDAvXIxj00iG3NJOFIgEgOVkkeYQ4S_HAabMjpRnltutWDf4ll0o5bMyZ11Hlzr",
		"accept-language": "zh-CN,zh;q=0.9,en;q=0.8",
		"accept":"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8",
		"Upgrade-Insecure-Requests":"1",
		"Cache-Control":"max-age=0",
		
	}
	
)



//