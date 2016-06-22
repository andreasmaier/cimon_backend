package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"io"
	"bytes"
	"golang.org/x/net/context"
	"google.golang.org/grpc/credentials"
	"log"
	"crypto/x509"
	"google.golang.org/grpc"
	pb "github.com/andreasmaier/cimon_jobs/jobs"
	"github.com/philips/grpc-gateway-example/insecure"
)

type JobBuild struct {
	Number int `json:"number"`
	QueueId int `json:"queue_id"`
	Phase string `json:"phase"`
	Status string `json:"status"`
	Url string `json:"url"`
	Log string `json:"log"`
}

type JobUpdate struct {
	Name string `json:"name"`
	Url string `json:"url"`
	Build JobBuild `json:"build"`
}

const (
	grpcPort = 10000
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
	demoAddr = fmt.Sprintf("localhost:%d", grpcPort)

	var opts []grpc.DialOption
	creds := credentials.NewClientTLSFromCert(demoCertPool, demoAddr)
	opts = append(opts, grpc.WithTransportCredentials(creds))

	conn, err := grpc.Dial(demoAddr, opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewJobsApiClient(conn)

	r, err := client.GetAllJobs(context.Background(), &pb.Empty{})

	if err != nil {
		panic(err)
	}
	log.Printf("Watched Jobs %d", len(r.Jobs))
}

func JobUpdates(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", body)

	decoder := json.NewDecoder(bytes.NewReader(body))
	var update JobUpdate
	if err := decoder.Decode(&update); err != nil {
		panic(err)
	}

	if update.Build.Phase == "COMPLETED" {
		updateRequest := pb.UpdateStatusRequest{
			Path: update.Url,
			Status: update.Build.Status,
		}

		demoCertPool = x509.NewCertPool()
		ok := demoCertPool.AppendCertsFromPEM([]byte(insecure.Cert))
		if !ok {
			panic("bad certs")
		}
		demoAddr = fmt.Sprintf("localhost:%d", grpcPort)

		var opts []grpc.DialOption
		creds := credentials.NewClientTLSFromCert(demoCertPool, demoAddr)
		opts = append(opts, grpc.WithTransportCredentials(creds))

		conn, err := grpc.Dial(demoAddr, opts...)
		if err != nil {
			panic(err)
		}
		defer conn.Close()

		client := pb.NewJobsApiClient(conn)

		_, err = client.UpdateJobStatus(context.Background(), &updateRequest)

		if err != nil {
			fmt.Printf("Error updating job status: %s\n", err)
		}
	}

	h.broadcast <- body
}

func WSTest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ws test\n")

	h.broadcast <- []byte("Hey this is a message")
}