package main

import (
	"net/http"
	"log"
	"google.golang.org/grpc"

	pb "github.com/andreasmaier/cimon_jobs/jobs"
	"golang.org/x/net/context"
	"fmt"
	"google.golang.org/grpc/credentials"
	"github.com/philips/grpc-gateway-example/insecure"
	"crypto/x509"
)

const (
	port = 10000
)

var (
	demoCertPool *x509.CertPool
	demoAddr     string
)

func init() {
	demoCertPool = x509.NewCertPool()
	ok := demoCertPool.AppendCertsFromPEM([]byte(insecure.Cert))
	if !ok {
		panic("bad certs")
	}
	demoAddr = fmt.Sprintf("localhost:%d", port)
}

func main() {
	go h.run()

	router := NewRouter()

	var opts []grpc.DialOption
	creds := credentials.NewClientTLSFromCert(demoCertPool, "localhost:10000")
	opts = append(opts, grpc.WithTransportCredentials(creds))

	conn, err := grpc.Dial(demoAddr, opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := pb.NewJobsApiClient(conn)

	r, err := c.GetAllJobs(context.Background(), &pb.Empty{})

	if err != nil {
		panic(err)
	}
	log.Printf("Watched Jobs %d", len(r.Jobs))

	log.Fatal(http.ListenAndServe(":3000", router))
}