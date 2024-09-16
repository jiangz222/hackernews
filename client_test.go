package hackernews

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	ast := assert.New(t)

	c := NewClient()
	stories, err := c.GetTopStoryIds(3)
	ast.Nil(err)
	ast.Len(stories.Ids, 3)

	story, err := c.GetStory(stories.Ids[0])
	ast.Nil(err)
	ast.Equal(stories.Ids[0], story.ID)

	comment, err := c.GetComment(story.Kids[0])
	ast.Nil(err)
	ast.Equal(story.ID, comment.Parent)

	allComments, err := c.GetAllComments(story.ID, 3, 3)
	ast.Nil(err)
	fmt.Println(allComments.Kids)
	//for _, v := range allComments.Comments {
	//	fmt.Println("1st level: ", v.ID, v.Text, v.Kids)
	//	for _, v2 := range v.Comments {
	//		fmt.Println("2nd level: ", v2.ID, v2.Text, v2.Kids)
	//		for _, v3 := range v2.Comments {
	//			fmt.Println("3rd level: ", v3.ID, v3.Text, v3.Kids)
	//			for _, v4 := range v3.Comments {
	//				fmt.Println("4th level: ", v4.ID, v4.Text, v4.Kids)
	//			}
	//		}
	//	}
	//}
}
