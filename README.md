Project Deliverables:
- Account creation
- Account login
- Creation of new application stack
- Creation of new application

Delieverable Specification:
- Account Creation
  - Takes credentials (email, username, password)
  - Creates new user in Firebase
  - Sends access keys back to client
  - Creates namespace for user in cluster
  - Installs application builder ServiceAccount to users namespace
  - Automatically deploys default stack with default application
- Account Login
  - Take credentials (email/username, password)
  - Sends access keys back to client
- Creation of new application stack
  - Creates Helm Chart
- Creation of new application
  - 