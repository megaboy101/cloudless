package harbor

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type KubeCluster struct {
	clientset *kubernetes.Clientset
}

func (k *KubeCluster) Connect() (Cluster, error) {
	// Load config from Kubernetes
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
		return nil, err
	}

	// Create a client from the fetched configuration
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
		return nil, err
	}

	// // Test to see how many pods are in the cluster
	// pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	// if err != nil {
	// 	panic(err.Error())
	// 	return nil, err
	// }

	return &KubeCluster{clientset}, nil
}

func (k *KubeCluster) CreateNamespace(name string) error {
	ns := &v1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: name}}

	_, err := k.clientset.CoreV1().Namespaces().Create(ns)
	if err != nil {
		panic(err.Error())
	}

	return nil
}
