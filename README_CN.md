# go-hackernews-client
[HackerNews](https://news.ycombinator.com/news) 的 Golang client

基于[官方的API文档](https://github.com/HackerNews/API)

# Usage
```golang
	c := NewClient()
	
	stories, err := c.GetTopStoryIds(3)

	story, err := c.GetStory(stories.Ids[0])

	comment, err := c.GetComment(story.Kids[0])

	allComments, err := c.GetAllComments(story.ID, 3, 3)
```