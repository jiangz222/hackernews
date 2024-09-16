# go-hackernews-client
Golang client for hackerNews https://news.ycombinator.com/news

Based on [Official API](https://github.com/HackerNews/API)

# Usage
```golang
	c := NewClient()

	storyIds, err := c.GetTopStoryIds(3)

    story, err := c.GetStory(storyIds.Ids[0])

    comment, err := c.GetComment(story.Kids[0])

    allComments, err := c.GetAllComments(story.ID, 3, 3)
```