package fb

import (
	"encoding/json"
	"fmt"
	"log"
)

type Page struct {
	ID string
	Token string
}

type Paging struct {
	Previous string `json:"previous"`
	Next string `json:"next"`
}

type ConversationResult struct {
	Data []Conversation `json:"data"`
	Paging Paging `json:"paging"`
}

type MessageResult struct {
	Data []Message `json:"data"`
	Paging Paging `json:"paging"`
}

func (p *Page) Get100Conversations() ([]Conversation, error) {
	url := fmt.Sprintf("%s/%s/conversations?limit=100", GraphAPIEndpoint, p.ID)
	conversations, _, err := p.extractConversations(url)
	if err != nil {
		log.Println("Getting 100 conversations failed:", err.Error())
		return nil, err
	}
	return conversations, nil
}

func (p *Page) GetAllConversations() ([]Conversation, error) {
	var result, tmp []Conversation
	var next string = fmt.Sprintf("%s/%s/conversations?limit=50", GraphAPIEndpoint, p.ID)
	var err error

	for next != "" {
		tmp, next, err = p.extractConversations(next)
		if err != nil {
			log.Println("Getting conversations failed:", err.Error())
			return nil, err
		}
		result = append(result, tmp...)
		fmt.Printf("Got %d conversations\n", len(result))
	}

	return result, nil
}

func (p *Page) FetchMessagesTo(c *Conversation) error {
	var messages []Message 
	var next string = fmt.Sprintf("%s/%s/messages?limit=50", GraphAPIEndpoint, c.ID)
	var err error 
	for next != "" {
		messages, next, err = p.extractMessages(next)
		if err != nil {
			log.Println("Getting messages failed:", err.Error())
			return err
		}
		c.Messages = append(c.Messages, messages...)
	}

	for k, message := range c.Messages {
		resp, err := httpget(fmt.Sprintf("%s/%s?fields=id,from,to,message,created_time&access_token=%s", GraphAPIEndpoint, message.ID, p.Token))
		if err != nil {
			log.Println("Getting message failed:", err.Error())
			return err
		}
		err = json.Unmarshal(resp, &c.Messages[k])
		if err != nil {
			log.Println("Unmarshalling message failed:", err.Error())
			return err
		}
	} 

	return nil
}

func (p *Page) extractConversations(url string) ([]Conversation, string, error) {
	resp, err := httpget(fmt.Sprintf("%s&access_token=%s", url, p.Token))
	if err != nil {
		log.Println("Getting conversations failed:", err.Error())
		return nil, "", err
	}
	var cr ConversationResult
	err = json.Unmarshal(resp, &cr)
	if err != nil {
		log.Println("Unmarshalling conversation result failed:", err.Error())
		return nil, "", err
	}

	return cr.Data, cr.Paging.Next, nil
}

func (p *Page) extractMessages(url string) ([]Message, string, error) {
	resp, err := httpget(fmt.Sprintf("%s&access_token=%s", url, p.Token))
	if err != nil {
		log.Println("Getting messages failed:", err.Error())
		return nil, "", err
	}
	var mr MessageResult
	err = json.Unmarshal(resp, &mr)
	if err != nil {
		log.Println("Unmarshalling message result failed:", err.Error())
		return nil, "", err
	}

	return mr.Data, mr.Paging.Next, nil
}