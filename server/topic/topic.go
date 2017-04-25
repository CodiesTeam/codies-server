package topic

import "codies-server/skeleton/common"

func AddTopic(title, content, authorID string) (string, error) {
	topicID := newTopicID()
	return addPost(TopicType, topicID, title, content, authorID)
}

func AddReply(topicID, content, authorID string) (string, error) {
	count, err := CountPost(topicID, ReplyType)
	if err != nil {
		return "", err
	}
	replyID := newReplyID(topicID, count+1)
	return addPost(ReplyType, replyID, "", content, authorID)
}

func AddComment(replyID, content, authorID string) (string, error) {
	count, err := CountPost(replyID, CommentType)
	if err != nil {
		return "", err
	}
	commentID := newCommentID(replyID, count+1)
	return addPost(CommentType, commentID, "", content, authorID)
}

func FullTopicByID(id string) (*Topic, error) {
	posts, err := fullPost(id)
	if err != nil {
		return nil, err
	}
	if len(posts) == 0 {
		return nil, common.NotFoundError("no post found by id: %s", id)
	}
	topic := newTopic(posts[0], nil)
	for _, p := range posts[1:] {
		// reply
		if p.Type == ReplyType {
			topic.Replies = append(topic.Replies, newReply(p, nil))
			continue
		}
		// comment
		_, replyNumbericID := getReplyID(p.ID)
		index := replyNumbericID - 1
		topic.Replies[index].Comments = append(topic.Replies[index].Comments, p)
	}
	return &topic, nil
}
