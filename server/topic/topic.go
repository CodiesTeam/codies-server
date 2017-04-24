package topic

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
