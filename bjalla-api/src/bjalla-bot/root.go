package bjalla_bot

import (
	"fmt"
	"strings"
	"time"

	"github.com/pocketbase/pocketbase/core"
)

// Debugging function to handle incoming messages and execute commands.
// This should be removed for production.
func HandleMessage(e *core.RecordEvent) error {
	content := e.Record.Get("content")

	// if content does not start with "!", ignore the message
	if content == nil || content.(string) == "" || content.(string)[0] != '!' {
		return e.Next()
	}

	// send incoming message
	if err := e.Next(); err != nil {
		return err
	}

	// fetch `messages` collection
	collection, err := e.App.FindCollectionByNameOrId("messages")
	if err != nil {
		return err
	}

	// find the command
	content_str := content.(string)[1:]
	spaceIdx := strings.Index(content_str, " ")
	if spaceIdx == -1 {
		spaceIdx = len(content_str)
	}
	command := content_str[:spaceIdx]

	switch command {
	// reply with "pong" when the command is "ping"
	case "ping":
		reply := core.NewRecord(collection)
		reply.Set("author", "pbbotbjalla0000")
		reply.Set("content", "pong")
		reply.Set("channel", e.Record.Get("channel"))

		if err := e.App.Save(reply); err != nil {
			return err
		}
	// flood N messages into the channel when the command is "flood N"
	case "flood":
		if spaceIdx == len(content_str) {
			return nil
		}

		n := 0
		if _, err := fmt.Sscanf(content_str[spaceIdx+1:], "%d", &n); err != nil {
			return nil
		}

		if err := e.Next(); err != nil {
			return nil
		}

		for i := 0; i < n; i++ {
			reply := core.NewRecord(collection)
			reply.Set("author", "pbbotbjalla0000")
			reply.Set("content", fmt.Sprintf("flood %d/%d", i+1, n))
			reply.Set("channel", e.Record.Get("channel"))

			if err := e.App.Save(reply); err != nil {
				return err
			}

			// sleep for 100ms to avoid flooding too fast
			time.Sleep(100 * time.Millisecond)
		}

		return nil
	}
	
	return nil
}
