package search

import "log"

//Result包含搜索的结果
type Result struct {
	Field   string
	Content string
}

//定义了想要实现一个新的搜索类型时必须有的行为
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

//为每一个独立的feed执行goroutine并发搜索
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	//由指定的matcher对象执行搜索
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	//将结果写入到channel中
	for _, result := range searchResults {
		results <- result
	}
}

//由独立的goroutine执行，收到结果时打印到控制台
func Display(results chan *Result) {
	//channel会一直阻塞，直到有结果写入
	//一旦channel关闭，循环就会结束
	for result := range results {
		log.Printf("%s : %s\n\n", result.Field, result.Content)
	}
}
