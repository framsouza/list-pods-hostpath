package kubectl

import (
	"context"
	"encoding/base64"
	"fmt"

	container "google.golang.org/api/container/v1"

	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/tools/clientcmd/api"
)

func GetK8sClusterConfigs(ctx context.Context, projectID string) (*api.Config, error) {
	svc, err := container.NewService(ctx)
	if err != nil {
		return nil, fmt.Errorf("container.NewService: %w", err)
	}

	// Basic config structure
	ret := api.Config{
		APIVersion: "v1",
		Kind:       "Config",
		Clusters:   map[string]*api.Cluster{},  // Clusters is a map of referencable names to cluster configs
		AuthInfos:  map[string]*api.AuthInfo{}, // AuthInfos is a map of referencable names to user configs
		Contexts:   map[string]*api.Context{},  // Contexts is a map of referencable names to context configs
	}

	// Ask Google for a list of all kube clusters in the given project.
	resp, err := svc.Projects.Zones.Clusters.List(projectID, "-").Context(ctx).Do()
	if err != nil {
		return nil, fmt.Errorf("clusters list project=%s: %w", projectID, err)
	}

	for _, f := range resp.Clusters {
		name := fmt.Sprintf(f.Name)
		cert, err := base64.StdEncoding.DecodeString(f.MasterAuth.ClusterCaCertificate)
		if err != nil {
			return nil, fmt.Errorf("invalid certificate cluster=%s cert=%s: %w", name, f.MasterAuth.ClusterCaCertificate, err)
		}

		ret.Clusters[name] = &api.Cluster{
			CertificateAuthorityData: cert,
			Server:                   "https://" + f.Endpoint,
		}

		ret.Contexts[name] = &api.Context{
			Cluster:  name,
			AuthInfo: name,
		}
		ret.AuthInfos[name] = &api.AuthInfo{
			AuthProvider: &api.AuthProviderConfig{
				Name: "gcp",
				Config: map[string]string{
					"scopes": "https://www.googleapis.com/auth/cloud-platform",
				},
			},
		}
	}

	return &ret, nil
}
