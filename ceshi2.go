package main

import (
	"strconv"
	"log"
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"github.com/satori/go.uuid"
	"os"
	"net/http"
	"io/ioutil"
)

func getAllUrls()[]string{
	var urls []string
	var url string
	for i:=0;i<2 ;i++  {
		url ="http://www.meizitu.com/a/more_" +strconv.Itoa(i+1)+".html"
		urls = append(urls,url)
	}
	return urls
}

func parseHtml(url string){
	doc,err:=goquery.NewDocument(url) //获取将要爬取的html文档信息
	if err!=nil{
		log.Fatal(err)
	}
	p :=make(chan string) //新开的管道
	doc.Find(".pic > a > img").Each(func(i int, s *goquery.Selection) {//遍历整个文档
		img_url,_ := s.Attr("src")

		//启动携程下载图片

		go download(img_url,p)
		fmt.Println("src =" + <-p+"图片爬取完毕")

	})


}

func download(img_url string ,p chan string){
	uid ,_ :=uuid.NewV4() //随机成成四段文件名

	file_name := uid.String() + ".jpg"

	fmt.Println(file_name)
	f,err := os.Create(file_name)
	if err != nil {
		panic("文件创建失败")
	}

	defer f.Close()

	resp , err := http.Get(img_url)

	if err != nil{
		fmt.Println("http.get err",err)
	}

	body,err1 := ioutil.ReadAll(resp.Body)

	if err1 != nil{
		fmt.Println("读取数据失败")
	}

	defer resp.Body.Close() //结束关闭

	f.Write(body)
	p<-file_name

}



func main(){
	urls := getAllUrls()
	for _,url :=range urls{
		parseHtml(url)
	}
}
///D:\gowww