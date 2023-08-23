package torrex

import (
	"log"

	"github.com/anacrolix/torrent"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Torrex struct {
	client   *torrent.Client
	torrents []*torrent.Torrent
	files    []*torrent.File
}

func NewClient() *Torrex {
	client, err := torrent.NewClient(nil)
	if err != nil {
		log.Fatalln(err)
	}
	minio.New("", &minio.Options{
		Creds: credentials.New(&credentials.CustomTokenIdentity{
			
		}),
	})
	return &Torrex{
		client,
		nil,
		nil,
	}
}
