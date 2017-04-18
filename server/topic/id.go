package topic

import (
	"codies-server/skeleton/common"
	"fmt"
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
