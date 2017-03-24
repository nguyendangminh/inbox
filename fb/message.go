package fb 

type Message struct {
	ID string `json:"id"`
	From User `json:"from"`
	To struct {
		Data []User `json:"data"`
	} `json:"to"`
	Text string `json:"message"`
	CreatedTime string `json:"created_time"`
}