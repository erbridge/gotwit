package callback

import (
	"github.com/erbridge/gotwit/twitter"
)

type (
	Callback func(twitter.Tweet)

	Type int
)

const (
	Retweet Type = iota
	Reply
	Mention
	Post
)
