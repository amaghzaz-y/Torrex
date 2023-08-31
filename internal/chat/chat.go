package chat

import (
	"container/list"

	"github.com/amaghzaz-y/torrex/internal/nanoid"
)

type Chat struct {
	rooms map[string]*ChatRoom
}
type Message struct {
	Sender  string `json:"sender"`
	Content string `json:"content"`
}

type ChatRoom struct {
	id       string
	messages *list.List
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
		id,
		messages,
	}
	c.rooms[id] = cr
	return cr
}

// returns an existing chatroom, or creates a new one if it doesnt exists
func (c *Chat) ChatRoom(id string) *ChatRoom {
	if r, e := c.rooms[id]; e {
		return r
	}
	return c.NewChatRoom()
}

func (c *Chat) DeleteChatRoom(id string) {
	delete(c.rooms, id)
}

func (cr *ChatRoom) PushMessage(sender string, content string) {
	message := Message{
		sender,
		content,
	}
	if cr.messages.Len() < 50 {
		cr.messages.PushFront(message)
	} else {
		first := cr.messages.Back()
		cr.messages.Remove(first)
		cr.messages.PushFront(message)
	}
}

func (cr *ChatRoom) Messages() []Message {
	var m []Message
	for e := cr.messages.Front(); e != nil; e = e.Next() {
		if msg, ok := e.Value.(Message); ok {
			m = append(m, msg)
		}
	}
	return m
}
