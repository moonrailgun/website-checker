package main

import (
  "log"
  "flag"
  "net/http"
  "time"
)

func main(){
  url := flag.String("url", "", "测试网址的地址")
  space := flag.Int("space", 1, "测试网址的间隔")
  flag.Parse()
  log.Println("测试的地址 :", *url)
  log.Println("请求的间隔 :", *space, "s")
  if *url == "" {
    log.Fatalf("测试地址不能为空, 请使用 --url 方式输入测试地址")
  }
  interval := time.Duration(*space) * time.Second

  for {
    sendTest(*url)
    time.Sleep(interval)
  }
}

func sendTest(url string) {
  client := &http.Client{}
  request, err := http.NewRequest("GET", url, nil)
  if err != nil {
		log.Println("请求错误:", err)
	}
  response, _ := client.Do(request)
  log.Println("请求结果:", response.StatusCode)
}
