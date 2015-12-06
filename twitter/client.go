package twitter

import (
	"net/url"
	"strconv"

	"github.com/ChimeraCoder/anaconda"
)

type (
	Client struct {
		api      *anaconda.TwitterApi
		stream   *anaconda.Stream
		callback func(t Tweet)
	}
)

func NewClient(accessConfig AccessConfig, callback func(t Tweet)) Client {
	return Client{
		api:      anaconda.NewTwitterApi(accessConfig.token, accessConfig.tokenSecret),
		callback: callback,
	}
}

func (c *Client) handleStream() {
	for t := range c.stream.C {
		if t, ok := t.(Tweet); ok {
			c.callback(t)
		}
	}
}

func (c *Client) Start() error {
	if ok, err := c.api.VerifyCredentials(); !ok {
		return err
	}

	v := url.Values{
		"replies": {"all"},
	}

	c.stream = c.api.UserStream(v)
	c.handleStream()

	return nil
}

func (c *Client) Stop() (err error) {
	close(c.stream.C)

	return
}

func (c *Client) post(message string, v url.Values) error {
	_, err := c.api.PostTweet(message, v)
	return err
}

func (c *Client) Post(message string, nsfw bool) error {
	v := url.Values{
		"possibly_sensitive": {strconv.FormatBool(nsfw)},
	}
	return c.post(message, v)
}

func (c *Client) Reply(tweet Tweet, message string, nsfw bool) error {
	v := url.Values{
		"possibly_sensitive":    {strconv.FormatBool(nsfw)},
		"in_reply_to_status_id": {tweet.IdStr},
	}
	return c.post(message, v)
}
