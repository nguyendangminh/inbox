package fb 

import (
	"encoding/json"
	"fmt"
	"log"
)

type Facebook struct {
	user_token string
}

func New(user_token string) *Facebook {
	return &Facebook{user_token: user_token}
}

func (f *Facebook) getPageToken(page_id string) (string, error) {
	resp, err := httpget(fmt.Sprintf("%s/%s?fields=access_token&access_token=%s", GraphAPIEndpoint, page_id, f.user_token))
	if err != nil {
		log.Println("Getting page token failed:", err.Error())
		return "", err
	}

	var result map[string]string
	err = json.Unmarshal(resp, &result)
	if err != nil {
		log.Println("Unmashalling response failed:", err.Error())
		return "", err
	}	

	return result["access_token"], nil
}

func (f *Facebook) NewPage(page_id string) (*Page, error) {
	var p Page
	page_token, err := f.getPageToken(page_id)
	if err != nil {
		log.Println("Init page failed:", err.Error())
		return nil, err
	}
	p.ID = page_id
	p.Token = page_token
	return &p, nil
}