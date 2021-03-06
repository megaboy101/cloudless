Application CLI Architecture:
  - Create Account
    - Registers user with Google Cloud AIM role
    - Initializes Kubernetes cluster to users' AIM role in Google Cloud
    - Installs custom extensions (istio, Knative, GitKube)
  - Create Stack
    - Registers a cluster namespace as the stack name
    - Reserves an ip for the namespace
    - Creates a dns record for the ip, and sends it to the user
  - Deploy Application to stack
    - Sends git project to private git repo
    - Builds code from repo into Knative container
    - Sends container to Google Cloud registry
    - Pulls container, and deploys to Kubernetes
