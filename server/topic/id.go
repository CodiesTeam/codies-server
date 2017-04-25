package topic

import (
	"codies-server/skeleton/common"
	"fmt"
	"strconv"
	"strings"
)

func newTopicID() string {
	return "topic-" + common.NewUniqueID()
}

func newReplyID(topicID string, replyID int) string {
	return fmt.Sprintf("%s_r%d", topicID, replyID)
}

func newCommentID(replyID string, commentID int) string {
	return fmt.Sprintf("%s_c%d", replyID, commentID)
}

// getReplyID parse comment id, get replyID, and reply numberic id
func getReplyID(commentID string) (string, int) {
	parts := strings.Split(commentID, "_")
	replyID := strings.Join(parts[:2], "_")
	replyNumberic, err := strconv.Atoi(parts[1][1:])
	if err != nil {
		panic(err)
	}
	return replyID, replyNumberic
}
