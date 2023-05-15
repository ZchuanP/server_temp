package zeebe

import (
	"github.com/camunda/zeebe/clients/go/v8/pkg/zbc"
	"log"
	"os"
)

func InitZeebeClient() zbc.Client {
	credentials, err := zbc.NewOAuthCredentialsProvider(&zbc.OAuthProviderConfig{
		ClientID:               os.Getenv("CLIENT_ID"),
		ClientSecret:           os.Getenv("CLIENT_SECRET"),
		AuthorizationServerURL: AuthorizationServerURL,
		Audience:               Audience,
	})
	if err != nil {
		panic(err)
	}

	config := zbc.ClientConfig{
		GatewayAddress:      os.Getenv("Zeebe_Addr"),
		CredentialsProvider: credentials,
	}

	client, err := zbc.NewClient(&config)
	if err != nil {
		panic(err)
	}

	return client
}

func MustCloseClient(client zbc.Client) {
	log.Println("closing client")
	_ = client.Close()
}
