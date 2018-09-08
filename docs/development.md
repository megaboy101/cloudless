Starting a local Minikube Cluster
==========

To run Harbor locally, you can run it through minikube with the following steps:

1. Ensure you have a running virtualization (libvirtd) and Docker daemon
2. Start Minikube
3. Run `eval $(minikube docker-env)` to set the correct docker settings for minikube
4. Compile Harbor to the `docker` directory
5. Build the docker image using the Dockerfile in `docker`
6. Use `kubectl run` to create a running harbor deployment
7. User `kubectl expose deployment` to expose Harbor as a service