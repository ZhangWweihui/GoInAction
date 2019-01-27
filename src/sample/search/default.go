package search

//一个默认的matcher实现
type defaultMatcher struct{}

//注册默认matcher到程序中
func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

//默认matcher的行为实现
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}
