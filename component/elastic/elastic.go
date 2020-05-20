package elastic

import (
	"github.com/olivere/elastic/v7"
	"log"
)

// Connect, 返回连接与错误
func Connect(host string) (*elastic.Client, error) {
	client, err := elastic.NewClient(elastic.SetURL(host), elastic.SetSniff(false))
	if err != nil {
		log.Fatalf("ConnectES error:%v", err)
	}

	return client, err
}
