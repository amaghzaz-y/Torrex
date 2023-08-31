package chat

import (
	"testing"

	nanoid "github.com/matoous/go-nanoid"
)

func TestChatFlow(t *testing.T) {
	chat := New()
	room := chat.NewChatRoom()
	room.PushMessage("Alice", "Hello There.")
	room.PushMessage("Bob", "Hey, How Are you ?")
	room.PushMessage("Alice", "I'm Good !")
	messages := room.Messages()
	msg0 := Message{"Alice", "I'm Good !"}
	msg1 := Message{"Bob", "Hey, How Are you ?"}
	msg2 := Message{"Alice", "Hello There."}
	if messages[0] != msg0 {
		t.Logf("Want %v but got %v", messages[0], msg0)
		t.Fail()
	}
	if messages[1] != msg1 {
		t.Logf("Want %v but got %v", messages[1], msg1)
		t.Fail()
	}
	if messages[2] != msg2 {
		t.Logf("Want %v but got %v", messages[2], msg2)
		t.Fail()
	}
	if len(messages) != 3 {
		t.Log("Want Len 3 but got ", len(messages))
		t.Fail()
	}
}

func TestChatOverFlow(t *testing.T) {
	chat := New()
	room := chat.NewChatRoom()
	for x := 0; x < 100; x++ {
		str, _ := nanoid.Nanoid(20)
		room.PushMessage(str, str)
	}
	if len(room.Messages()) != 50 {
		t.Log("want 50 but got", len(room.Messages()))
	}
}
