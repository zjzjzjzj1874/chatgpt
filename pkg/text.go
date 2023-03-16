package pkg

func (c *Client) Send() {
	var body []byte
	var response []byte
	c.R().SetBody(body).SetSuccessResult(response).MustPost(GPT_URL)
}
