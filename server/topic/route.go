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
		// route.NewRoute(
		// 	"/topic/:topicID/reply/add",
		// 	"POST",
		// 	addReply,
		// ),
	}
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

	topicID, err := AddReply(topicID, p.Content, p.AuthorID)
	if err != nil {
		return reply.Err(err)
	}
	return reply.JSON(map[string]interface{}{
		"id": topicID,
	})
}

func topicByID(ctx *context.Context) reply.Replyer {
	var topicID string
	if err := ctx.Input.Var("id", &topicID).Error(); err != nil {
		return reply.Err(err)
	}
	p, err := GetPostByID(topicID)
	if err != nil {
		return reply.Err(err)
	}
	return reply.JSON(common.M{"topic": p})
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
