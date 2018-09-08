package harbor

import "log"

func New() (*Service, error) {
	k := &KubeCluster{}

	cluster, err := k.Connect()
	if err != nil {
		log.Fatalf("Could not connect to Kubernetes cluster")
		return nil, err
	}

	return &Service{cluster}, nil
}
