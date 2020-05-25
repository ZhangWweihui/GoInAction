package search

import (
	"log"
	"sync"
)

//用于搜索的匹配器的集合
var matchers = make(map[string]Matcher)

//执行搜索
func Run(searchTerm string) {

	//取回要搜索的种子列表
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("feeds length : %d\n", len(feeds))

	//创建一个无缓冲的channel来接收要展示的匹配结果
	results := make(chan *Result)

	//设置一个wait group，这样我们就可以处理所有的种子
	var waitGroup sync.WaitGroup

	//设置需要等待的goroutine的数量
	waitGroup.Add(len(feeds))

	//为每一个种子启动一个goroutine来查询结果
	for _, feed := range feeds {
		//根据种子的类型获取一个匹配器
		log.Printf("feed type : %s \n", feed.Type)
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	//运行一个goroutine来监控是否所有的工作已完成
	go func() {
		waitGroup.Wait()
		//关闭channel，以向Display函数发送信号
		close(results)
	}()

	Display(results)
}

//注册匹配器
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "Matcher")
	matchers[feedType] = matcher
}
