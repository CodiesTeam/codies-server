package topic

import (
	"codies-server/skeleton/common"
	"codies-server/skeleton/context"
	"codies-server/skeleton/reply"
	"codies-server/skeleton/route"
)

func NewRoute() []*route.Route {
	return []*route.Route{
		route.NewRoute(
			"/topic/add",
			"POST",
			addTopic,
		),
		route.NewRoute(
			"/topic/:id",
			"GET",
			topicByID,
		),
		route.NewRoute(
			"/topic/:topicID/reply/add",
			"POST",
			addReply,
		),
		route.NewRoute(
			"/reply/:replyID/comment/add",
			"POST",
			addComment,
		),
	}
}

func addComment(ctx *context.Context) reply.Replyer {
	var replyID string
	if err := ctx.Input.Var("replyID", &replyID).Error(); err != nil {
		return reply.Err(err)
	}
	p := Post{}
	if err := ctx.Input.JSONBody(&p).Error(); err != nil {
		return reply.Err(err)
	}

	commentID, err := AddComment(replyID, p.Content, p.AuthorID)
	if err != nil {
		return reply.Err(err)
	}
	return reply.JSON(map[string]interface{}{
		"id": commentID,
	})
}

func addReply(ctx *context.Context) reply.Replyer {
	var topicID string
	if err := ctx.Input.Var("topicID", &topicID).Error(); err != nil {
		return reply.Err(err)
	}
	p := Post{}
	if err := ctx.Input.JSONBody(&p).Error(); err != nil {
		return reply.Err(err)
	}

	replyID, err := AddReply(topicID, p.Content, p.AuthorID)
	if err != nil {
		return reply.Err(err)
	}
	return reply.JSON(map[string]interface{}{
		"id": replyID,
	})
}

func topicByID(ctx *context.Context) reply.Replyer {
	var topicID string
	if err := ctx.Input.Var("id", &topicID).Error(); err != nil {
		return reply.Err(err)
	}
	p, err := FullTopicByID(topicID)
	if err != nil {
		return reply.Err(err)
	}
	ret := common.M{"topic": p}
	return reply.JSON(ret)
}

func addTopic(ctx *context.Context) reply.Replyer {
	p := Post{}
	if err := ctx.Input.JSONBody(&p).Error(); err != nil {
		return reply.Err(err)
	}

	topicID, err := AddTopic(p.Title, p.Content, p.AuthorID)
	if err != nil {
		return reply.Err(err)
	}
	return reply.JSON(map[string]interface{}{
		"id": topicID,
	})
}
