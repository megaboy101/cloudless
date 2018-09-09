package harbor

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
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
	ns := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: name}}

	_, err := k.clientset.CoreV1().Namespaces().Create(ns)
	if err != nil {
		panic(err.Error())
	}

	return nil
}

func (k *KubeCluster) InstallTiller() error {
	// Deployment labels
	labels := map[string]string{
		"name": "tiller",
		"app":  "cloudless",
	}

	// Deployment Configuration
	d := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "tiller",
			Name:      "tiller",
			Labels:    labels,
		},
		Spec: appsv1.DeploymentSpec{
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					ServiceAccountName: "default",
					Containers: []corev1.Container{
						{
							Name:            "tiller",
							Image:           "gcr.io/kubernetes-helm/tiller",
							ImagePullPolicy: corev1.PullIfNotPresent,
							Ports: []corev1.ContainerPort{
								{ContainerPort: 44134, Name: "tiller"},
								{ContainerPort: 44135, Name: "http"},
							},
							Env: []corev1.EnvVar{
								{Name: "TILLER_NAMESPACE", Value: "tiller"},
								{Name: "TILLER_HISTORY_MAX", Value: "0"},
							},
							LivenessProbe: &corev1.Probe{
								Handler: corev1.Handler{
									HTTPGet: &corev1.HTTPGetAction{
										Path: "/liveness",
										Port: intstr.FromInt(44135),
									},
								},
								InitialDelaySeconds: 1,
								TimeoutSeconds:      1,
							},
							ReadinessProbe: &corev1.Probe{
								Handler: corev1.Handler{
									HTTPGet: &corev1.HTTPGetAction{
										Path: "/readiness",
										Port: intstr.FromInt(44135),
									},
								},
								InitialDelaySeconds: 1,
								TimeoutSeconds:      1,
							},
						},
					},
					HostNetwork:  false,
					NodeSelector: map[string]string{},
				},
			},
		},
	}

	// Create deployment in Kubernetes
	_, err := k.clientset.AppsV1().Deployments(d.Namespace).Create(d)
	if err != nil {
		return err
	}

	s := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "tiller",
			Name:      "tiller-service",
			Labels:    labels,
		},
		Spec: corev1.ServiceSpec{
			Type: corev1.ServiceTypeClusterIP,
			Ports: []corev1.ServicePort{
				{
					Name:       "tiller",
					Port:       44134,
					TargetPort: intstr.FromString("tiller"),
				},
			},
			Selector: labels,
		},
	}

	_, err = k.clientset.CoreV1().Services(s.Namespace).Create(s)
	if err != nil {
		return err
	}

	return nil
}
