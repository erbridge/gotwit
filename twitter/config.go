package twitter

type (
	ConsumerConfig struct {
		key    string
		secret string
	}

	AccessConfig struct {
		token       string
		tokenSecret string
	}
)

func NewConsumerConfig(key string, secret string) ConsumerConfig {
	return ConsumerConfig{
		key:    key,
		secret: secret,
	}
}

func NewAccessConfig(token string, tokenSecret string) AccessConfig {
	return AccessConfig{
		token:       token,
		tokenSecret: tokenSecret,
	}
}
