package main

import (
	"fmt"
	"log"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

// StartChrome 启动谷歌浏览器headless模式

func StartChrome() {

	opts := []selenium.ServiceOption{}

	caps := selenium.Capabilities{
		"browserName": "chrome",
	}

	// 禁止加载图片，加快渲染速度

	imagCaps := map[string]interface{}{
		"profile.managed_default_content_settings.images": 2,
	}

	chromeCaps := chrome.Capabilities{
		Prefs: imagCaps,
		Path:  "",
		Args: []string{
			//"--headless", // 设置Chrome无头模式
			"--no-sandbox",
			"--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/604.4.7 (KHTML, like Gecko) Version/11.0.2 Safari/604.4.7", // 模拟user-agent，防反爬
		},
	}

	caps.AddChrome(chromeCaps)

	// 启动chromedriver，端口号可自定义

	service, err := selenium.NewChromeDriverService("./chromedriver", 9516, opts...)

	if err != nil {

		log.Printf("Error starting the ChromeDriver server: %v", err)
	}

	// 调起chrome浏览器

	webDriver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 9516))

	if err != nil {

		panic(err)

	}

	//目标网站

	//targeUrl := "https://www.toutiao.com/i6846744256028082696"
	targeUrl := "http://www.baidu.com"

	// 导航到目标网站

	err = webDriver.Get(targeUrl)

	if err != nil {
		panic(fmt.Sprintf("Failed to load page: %s\n", err))
	}

	log.Println(webDriver.GetCookies())
	defer service.Stop() // 停止chromedriver

	defer webDriver.Quit() // 关闭浏览器

}

func main() {
	StartChrome()
}
