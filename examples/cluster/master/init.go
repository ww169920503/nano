package master

import (
	"github.com/ww169920503/nano/component"
	"github.com/ww169920503/nano/session"
)

var (
	// All services in master server
	Services = &component.Components{}

	// Topic service
	topicService = newTopicService()
	// ... other services
)

func init() {
	Services.Register(topicService)
}

func OnSessionClosed(s *session.Session) {
	topicService.userDisconnected(s)
}
