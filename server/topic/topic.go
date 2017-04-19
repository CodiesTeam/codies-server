package topic

func AddTopic(title, content, authorID string) (string, error) {
	topicID := newTopicID()
	return addPost(TopicType, topicID, title, content, authorID)
}

func AddReply(topicID, content, authorID string) (string, error) {
	replyCount, err := CountReply(topicID)
	if err != nil {
		return "", err
	}
	replyID := newReplyID(topicID, replyCount+1)
	return addPost(ReplyType, replyID, "", content, authorID)
}
