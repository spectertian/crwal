package main

import (
	"fmt"
	"regexp"
)

func main() {

	url := "https://www.domp4.cc/html/NrygxYDDDDDY.html"
	Regexp := regexp.MustCompile(`([^/]*?)\.html`)
	params := Regexp.FindStringSubmatch(url)

	fmt.Println(params)

	panic(23)
	reg, err := regexp.Compile(`.*?:`)

	//s := "magnet:?xt=urn:btih:dcdfddbb8c5c4e007a5b98d1494a791f0c6de952&amp;dn=[www.domp4.cc]团圆饭之小小少年.2021.HD1080p.国语中字.mp4&amp;tr=http://tracker.domp4.cc:9001/announce&amp;tr=http://1337.abcvg.info:80/announce&amp;tr=http://5rt.tace.ru:60889/announce&amp;tr=http://h4.trakx.nibba.trade:80/announce&amp;tr=http://open.acgnxtracker.com:80/announce&amp;tr=http://rt.tace.ru:80/announce&amp;tr=http://share.camoe.cn:8080/announce&amp;tr=http://t.nyaatracker.com:80/announce&amp;tr=https://tp.m-team.cc:443/announce.php&amp;tr=https://tracker.cyber-hub.net:443/announce&amp;tr=https://tracker.foreverpirates.co:443/announce&amp;tr=https://tracker.gbitt.info:443/announce&amp;tr=https://tracker.imgoingto.icu:443/announce&amp;tr=https://tracker.lilithraws.cf:443/announce&amp;tr=https://tracker.nitrix.me:443/announce&amp;tr=udp://10.rarbg.me:80/announce&amp;tr=udp://12.rarbg.me:80/announce&amp;tr=udp://3rt.tace.ru:60889/announce&amp;tr=udp://47.ip-51-68-199.eu:6969/announce&amp;tr=udp://61626c.net:6969/announce&amp;tr=udp://6rt.tace.ru:80/announce&amp;tr=udp://9.rarbg.me:2710/announce&amp;tr=udp://9.rarbg.to:2710/announce&amp;tr=udp://aaa.army:8866/announce"
	s := "rvrv:?xt=urn:btih:dcdfddbb8c5c4e007a5b98d1494a791f0c6de952&amp;dn=[www.domp4.cc]团圆饭之小小少年.2021.HD1080p.国语中字.mp4&amp;tr=http://tracker.domp4.cc:9001/announce&amp;tr=http://1337.abcvg.info:80/announce&amp;tr=http://5rt.tace.ru:60889/announce&amp;tr=http://h4.trakx.nibba.trade:80/announce&amp;tr=http://open.acgnxtracker.com:80/announce&amp;tr=http://rt.tace.ru:80/announce&amp;tr=http://share.camoe.cn:8080/announce&amp;tr=http://t.nyaatracker.com:80/announce&amp;tr=https://tp.m-team.cc:443/announce.php&amp;tr=https://tracker.cyber-hub.net:443/announce&amp;tr=https://tracker.foreverpirates.co:443/announce&amp;tr=https://tracker.gbitt.info:443/announce&amp;tr=https://tracker.imgoingto.icu:443/announce&amp;tr=https://tracker.lilithraws.cf:443/announce&amp;tr=https://tracker.nitrix.me:443/announce&amp;tr=udp://10.rarbg.me:80/announce&amp;tr=udp://12.rarbg.me:80/announce&amp;tr=udp://3rt.tace.ru:60889/announce&amp;tr=udp://47.ip-51-68-199.eu:6969/announce&amp;tr=udp://61626c.net:6969/announce&amp;tr=udp://6rt.tace.ru:80/announce&amp;tr=udp://9.rarbg.me:2710/announce&amp;tr=udp://9.rarbg.to:2710/announce&amp;tr=udp://aaa.army:8866/announce"
	fmt.Printf("%q,%v\n", reg.FindString(s), err)
	// "Hello",
}
