package dict_query

import (
	"fmt"
	"gin-blog/pkg/e"
	"log"
	"net/http"

	"gin-blog/routers/base_routers"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

type search_result struct {
	Word   string            `json:word`
	Result map[string]string `json:result`
	From   string            `json:结果来自`
}

func ExecuteSearch(bs *base_routers.BaseRouterService) {
	word := bs.Context.Param("word")
	fmt.Println("QueryWord:" + word)
	//TODO check word
	query_result := getResultFromBing(word)
	bs.SetSuccessStatus(e.SUCCESS, "")
	bs.SetRespData(base_routers.Struct2Map(query_result))
}
func QueryWord(c *gin.Context) {
	bs := base_routers.DefaultBaseRouterService(c, ExecuteSearch)
	bs.Execute()
	//word := c.Param("word")
	//fmt.Println("QueryWord:" + word)
	//code := e.SUCCESS
	//data := make(map[string]interface{})
	//TODO check word
	//query_result := getResultFromBing(word)
	//TODO check search result,may be return empty map
	//data["result"] = query_result

	//c.JSON(http.StatusOK, gin.H{
	//"code": code,
	//"msg":  e.GetMsg(code),
	//"data": query_result,
	//})
}

//TODO move to miro services
func getResultFromBing(word string) search_result {
	result := make(map[string]string)
	search_url := "https://cn.bing.com/dict/search?q="
	//res, err := http.Get("https://cn.bing.com/dict/search?q=word")
	fmt.Println(search_url + word)
	res, err := http.Get(search_url + word)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatal("status code error :%d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	rs := doc.Find(".qdef")
	rs.Each(func(i int, s *goquery.Selection) {
		s.Find("ul").Each(func(i int, s *goquery.Selection) {
			post_dict := make(map[int]string)
			s.Find(".pos").Each(func(i int, s *goquery.Selection) {
				post_dict[i] = s.Text()
			})
			s.Find(".def").Each(func(i int, s *goquery.Selection) {
				if pos, ok := post_dict[i]; ok {
					result[pos] = s.Text()
				}
				//TODO if pos not find
			})
		})
	})
	//TODO log search result
	//TODO handle search result empty
	return search_result{word, result, search_url + word}
}
