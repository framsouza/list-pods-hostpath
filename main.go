package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	listPods "github.com/framsouza/list-pods-hostpath/pkg/listpods"
	"golang.org/x/oauth2/google"
	container "google.golang.org/api/container/v1"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

var (
	projectID = flag.String("project", "", "Enter the Project ID")
	zone      = flag.String("zone", "-", "Enter the Compute zone")
)

func main() {
	flag.Parse()

	if *projectID == "" {
		fmt.Fprintln(os.Stderr, "Missing project")
		flag.Usage()
		os.Exit(2)
	}
	if *zone == "" {
		fmt.Fprintln(os.Stderr, "Missing zone")
		flag.Usage()
		os.Exit(2)
	}

	ctx := context.Background()

	hc, err := google.DefaultClient(ctx, container.CloudPlatformScope)
	if err != nil {
		log.Fatalf("Could not get authenticated client: %v", err)
	}

	svc, err := container.New(hc)
	if err != nil {
		log.Fatalf("Could not initialize gke client: %v", err)
	}

	if err := listPods.ListPods(svc, *projectID, *zone); err != nil {
		log.Fatal(err)
	}
}
