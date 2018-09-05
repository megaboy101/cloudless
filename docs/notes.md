Project Notes:

Development Workflow:
- All Code must pass tests before being pushed to remote
- Commits follow GitFlow workflow for managing commits/versions
- Code kept in public Github Repo
- Code Managed via CircleCI
- All code must be tested before being pushed to remote
- Each feature should be self-sufficient, including
  - Proper error handling
  - Proper testing
  - Proper documentation
  - Follows style guide

Project Goals:
- Create a Kubernetes-native PaaS/FaaS for fast moving dev teams
- Allows you to harness the power of Kubernetes without needing
  all the technical experience it requires
- Allows for complex microservices and serverless functions to be
  built, and shared on one platform

Project Features:
- Full GitOps
- Support for stateful and stateless serverless containers
- Intuitive end-user web-interface
- Hosted or on-prem solution
- Application grouping/categorization with extras
  - Group authentication
  - Group rate limiting
  - Group scaling
  - Group pricing

CLI Application Deliverables:
- Create Account
  - Registers user via Firebase authentication
- Create Stack
  - Registers a cluster namespace as the stack name
  - Reserves an ip for the namespace
  - Creates a subdomain via a DNS record for the ip, and sends it to the user
- Create Application
  - Creates Git server for application
  - Responds with Git remote URL
- Deploy Application to Stack
  - Sends git project to private git repo
  - Reads app config file(s)
  - Builds code from repo into Knative container
  - Sends container to Google Cloud registry
  - Pulls container, and deploys it to Kubernetes

Kubernetes Stack Requirements:
- Multi-tenancy
  - Network/Resource policies?
  - namespaces
- Knative suite installed
- Custom Git integrator installed

Custom Software:
- Kubernetes Git server (custom GitKube)
- Application config protocol and interpreter (similar to Heroku app.json/Procfile)
- Stack packager into Helm package
- CLI Client
- Web Client

User Flow:
- User creates an account for the first time
- User is presented with first git remote url (default stack and application is created)
- User pushes code to remote and is given live url back
- User creates new stack
- User creates monorepo application

Kubernetes App Deployment Flow
- Build process is triggered
- Builder ServiceAccount downloads source, and processes build according
  to provided build template
  - Currently Projects are pulled via github urls, but there is support for
    private repos as well, provided they be uploaded to GCloud storage first
  - Current build template uses the Cloud Foundry buildpack system
- Build ServiceAccount sends built image to provided image registry using attached
  Secret credentials
  - If using GCloud, it is then handled by a GCloud ServiceAccount with permissions to
    write to the registry
- Service process pulls container image and deploys it as containers to cluster

Application Organizational Structure
- Organization inspired by Slack/Asana
  - Cloudless workspaces are equivalent to Slack Enterprise Grid or Asana
    workspaces
  - Cloudless projects are equivalent to Slack workspaces or Asana teams
  - Cloudless stacks are equivalent to Slack channels or Asana pages
  - Cloudless apps are equivalent to Slack chats or Asana items
- Absolute top-level structure is the project
  - First one is free (only one for most teams)
  - Creating additional workpspaces is an enterprise feature
    - Private workspace of projects seperate from personal projects
- Common organizational structure is the project
  - Each project has its own set of team members
- Inside a project is a series of stacks
  - Uses Helm charts behind the scenes
  - Stacks can be user-created or installed (ex: Gitlab)
  - Each stack can have its own environment
    - Multiple versions of same stack
  - Stacks can have addons (ex: MySql storage space)
- Inside a stack is a series of applications
