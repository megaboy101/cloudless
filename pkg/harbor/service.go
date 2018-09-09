package harbor

type Servicer interface {
	Cluster
}

type Cluster interface {
	Connect() (Cluster, error)
	CreateNamespace(name string) error
	InstallTiller() error
}

type Service struct {
	Cluster
}
