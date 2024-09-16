package hackernews

import "fmt"

type Client struct {
}

const (
	topStoriesUrl = "https://hacker-news.firebaseio.com/v0/topstories.json"
	storyUrl      = "https://hacker-news.firebaseio.com/v0/item/%d.json"
	commentUrl    = "https://hacker-news.firebaseio.com/v0/item/%d.json"
)

func NewClient() *Client {
	return &Client{}
}

// TopStoryIds Top story ids
type TopStoryIds struct {
	Ids []uint32 `json:"ids"` // ids of stories, ordered by API, same as website shows
}

// GetTopStoryIds get stories, ordered by official API,same as website shows
// limit: limit the length of ids, official API  return MAX 500 stories
func (c *Client) GetTopStoryIds(limit int) (*TopStoryIds, error) {
	ts := &TopStoryIds{}
	err := get(topStoriesUrl, &ts.Ids)
	if err != nil {
		return nil, err
	}
	if len(ts.Ids) > limit {
		ts.Ids = ts.Ids[:limit]
	}
	return ts, err
}

// Story basic information of Story
type Story struct {
	By          string     `json:"by"`
	Descendants int        `json:"descendants"` // numbers of comments
	ID          uint32     `json:"id"`
	Kids        []uint32   `json:"kids"`
	Score       int        `json:"score"`
	Time        int64      `json:"time"` // unix seconds
	Title       string     `json:"title"`
	Type        string     `json:"type"`
	URL         string     `json:"url"`      // url to original new
	Comments    []*Comment `json:"comments"` // unsupported by official API
}

func (c *Client) GetStory(id uint32) (*Story, error) {
	s := &Story{}
	err := get(fmt.Sprintf(storyUrl, id), s)
	return s, err
}

type Comment struct {
	By       string     `json:"by"`
	ID       uint32     `json:"id"`
	Kids     []uint32   `json:"kids"`
	Parent   uint32     `json:"parent"`
	Text     string     `json:"text"`
	Time     int        `json:"time"` // unix seconds
	Type     string     `json:"type"`
	Comments []*Comment `json:"comments"` // unsupported by official API
}

func (c *Client) GetComment(id uint32) (*Comment, error) {
	s := &Comment{}
	err := get(fmt.Sprintf(commentUrl, id), s)
	return s, err
}

// GetAllComments get all comments of storyId
// limit: limit the number of comments in same level
// depth: depth of comments should return
func (c *Client) GetAllComments(storyId uint32, limit int, depth int) (*Story, error) {
	story, err := c.GetStory(storyId)
	if err != nil {
		return nil, err
	}
	if len(story.Kids) < limit {
		limit = len(story.Kids)
	}
	story.Comments = c.getKidComments(story.Kids[:limit], limit, depth)
	return story, nil
}

// getKidComments get comments
func (c *Client) getKidComments(kids []uint32, limit, depth int) (comments []*Comment) {
	if depth == 0 || len(kids) == 0 {
		return nil
	}
	for _, cId := range kids {
		comment, err := c.GetComment(cId)
		if err != nil {
			fmt.Println("get comment err:", cId, err)
			continue
		}
		kids := comment.Kids
		if len(comment.Kids) > limit {
			kids = comment.Kids[:limit]
		}
		comment.Comments = c.getKidComments(kids, limit, depth-1)
		comments = append(comments, comment)
	}
	return
}
