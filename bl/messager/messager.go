package messager

import (
	"github.com/datreeio/datree/pkg/cliClient"
)

type MessagesClient interface {
	GetVersionMessage(cliVersion string, timeout int) (*cliClient.VersionMessage, error)
}

type Messager struct {
	defaultTimeout int
	messagesClient MessagesClient
}

func New(c MessagesClient) *Messager {
	return &Messager{
		defaultTimeout: 900,
		messagesClient: c,
	}
}

type VersionMessage struct {
	CliVersion   string
	MessageText  string
	MessageColor string
}

func (m *Messager) LoadVersionMessages(messages chan *VersionMessage, cliVersion string) {
	go func() {
		msg, _ := m.messagesClient.GetVersionMessage(cliVersion, 900)
		if msg != nil {
			messages <- m.toVersionMessage(msg)
		}
	}()

}

func (m *Messager) toVersionMessage(msg *cliClient.VersionMessage) *VersionMessage {
	return &VersionMessage{
		CliVersion:   msg.CliVersion,
		MessageText:  msg.MessageText,
		MessageColor: msg.MessageColor,
	}
}
