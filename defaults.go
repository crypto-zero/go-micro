package micro

import (
	"github.com/crypto-zero/go-micro/v2/client"
	"github.com/crypto-zero/go-micro/v2/debug/trace"
	"github.com/crypto-zero/go-micro/v2/server"
	"github.com/crypto-zero/go-micro/v2/store"

	// set defaults
	gcli "github.com/crypto-zero/go-micro/v2/client/grpc"
	memTrace "github.com/crypto-zero/go-micro/v2/debug/trace/memory"
	gsrv "github.com/crypto-zero/go-micro/v2/server/grpc"
	memoryStore "github.com/crypto-zero/go-micro/v2/store/memory"
)

func init() {
	// default client
	client.DefaultClient = gcli.NewClient()
	// default server
	server.DefaultServer = gsrv.NewServer()
	// default store
	store.DefaultStore = memoryStore.NewStore()
	// set default trace
	trace.DefaultTracer = memTrace.NewTracer()
}
