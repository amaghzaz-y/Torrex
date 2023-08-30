package chat

import (
	"container/list"

	nanoid "github.com/matoous/go-nanoid"
)

type Chat struct {
	rooms map[string]*ChatRoom
}
type Message struct {
	Sender  string
	Content string
}

type ChatRoom struct {
	Id       string
	Messages *list.List
}

func New() *Chat {
	return &Chat{
		rooms: make(map[string]*ChatRoom),
	}
}

func (c *Chat) NewChatRoom() *ChatRoom {
	id, _ := nanoid.ID(21)
	messages := list.New()
	cr := &ChatRoom{
		Id:       id,
		Messages: messages,
	}
	c.rooms[id] = cr
	return cr
}

func (c *Chat) DeleteChatRoom(id string) {
	delete(c.rooms, id)
}

func (cr *ChatRoom) PushMessage(sender string, content string) {
	message := Message{
		sender,
		content,
	}
	if cr.Messages.Len() < 50 {
		cr.Messages.PushFront(message)
	} else {
		first := cr.Messages.Back()
		cr.Messages.Remove(first)
		cr.Messages.PushFront(message)
	}
}

func (cr *ChatRoom) List() []Message {
	var m []Message
	for e := cr.Messages.Front(); e != nil; e = e.Next() {
		if msg, ok := e.Value.(Message); ok {
			m = append(m, msg)
		}
	}
	return m
}
