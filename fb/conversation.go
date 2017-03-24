package fb 

import (
	"log"
	"fmt"
	"os"
)

type Conversation struct {
	ID string `json:"id"`
	Link string `json:"link"`
	UpdatedTime string `json:"updated_time"`
	Messages []Message
}

func (c *Conversation) String() string {
	var conversation string 
	for k := len(c.Messages)-1; k >= 0; k-- {
		conversation = fmt.Sprintf("%s%s\t%s\t%s\n", conversation,c.Messages[k].CreatedTime, c.Messages[k].From.Name, c.Messages[k].Text)
	}
	return fmt.Sprintf(template, c.ID, c.Link, c.UpdatedTime, conversation)
}

var template string = `ID: %s
Link: http://facebook.com%s
Updated time: %s

--- Conversation ---

%s
`

func (c *Conversation) WriteTo(dir string) error {
	f, err := os.Create(fmt.Sprintf("%s/%s", dir, c.ID))
	if err != nil {
		log.Println("Writting file failed:", err.Error())
		return err
	}
	defer f.Close()

	_, err = f.WriteString(c.String())
	if err != nil {
		log.Println("Writting file failed:", err.Error())
		return err
	}

	return nil
}