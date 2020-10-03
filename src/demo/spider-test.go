package main

import (
	"log"
)

import (
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	"strings"
	"net/url"
)

var img_dir = "/tmp/tmp/"
var host = ""

func findResourceAttr(url string, selector string, attr string) ([]string) {
	document, e := goquery.NewDocument(url)

	if nil != e {
		log.Println(e)
		return nil
	}

	var innerImg []string

	selections := document.Find(selector)

	selections.Each(func(i int, selection *goquery.Selection) {

		contentHref, exists := selection.Attr(attr)

		if !exists {
			return
		}

		//fmt.Println(contentHref)

		innerImg = append(innerImg, contentHref)

	})

	return innerImg
}

func checkExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func saveImages(img_url string) {
	log.Println(img_url)
	u, err := url.Parse(img_url)
	if err != nil {
		log.Println("parse url failed:", img_url, err)
		return
	}

	//去掉最左边的'/'
	tmp := strings.TrimLeft(u.Path, "/")
	filename := img_dir + strings.ToLower(strings.Replace(tmp, "/", "-", -1))

	exists := checkExists(filename)
	if exists {
		return
	}

	response, err := http.Get(img_url)
	if err != nil {
		log.Println("get img_url failed:", err)
		return
	}

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("read data failed:", img_url, err)
		return
	}

	image, err := os.Create(filename)
	if err != nil {
		log.Println("create file failed:", filename, err)
		return
	}

	defer image.Close()
	image.Write(data)
}

/**
"http://kail.xyz/forum-index-fid-8.htm",
"#threadlist .thread .subject .thread_icon",
"href"
".message img", "src"

"https://kail.xyz/taotu/",
".taotu-main>ul li a",
"href"
"#big-pic img", "src"


 */
func main() {

	contentHrefs := findResourceAttr(
		"http://kail.xyz/forum-index-fid-8.htm",
		"#threadlist .thread .subject .thread_icon",
		"href")

	fmt.Println()
	fmt.Println()

	// 经典的循环条件初始化/条件判断/循环后条件变化
	for i := 0; i < len(contentHrefs); i++ {
		//fmt.Println(host+contentHrefs[i])
		imageSrcs := findResourceAttr(host+contentHrefs[i], ".message img", "src")

		for j := 0; j < len(imageSrcs); j++ {
			saveImages(imageSrcs[j])
		}

	}

}
