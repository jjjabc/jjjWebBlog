package article

type ArticleSortSlice []JJJarticle

func (articles ArticleSortSlice) Len() int {
	return len(articles)
}
func (articles ArticleSortSlice) Less(i, j int) bool {
	if articles[i].Priority < 0 {
		if articles[j].Priority < 0 {
			return articles[i].PublishedTime.Before(articles[j].PublishedTime)
		} else {
			return false
		}
	}
	if articles[j].Priority < 0 {
		return true
	}
	if articles[i].Priority == articles[j].Priority {
		return articles[i].PublishedTime.Before(articles[j].PublishedTime)
	} else {
		return articles[i].Priority < articles[j].Priority
	}
}
func (articles ArticleSortSlice) Swap(i, j int) {
	articles[i], articles[j] = articles[j], articles[i]
}
