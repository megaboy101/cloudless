API Abstract Structures:
  - Stack
    - The stack is the basic structure of Cloudless
    - A stack comprises of one or more applications, along with needed dependancies
    - Stacks use Helm Charts behind the scenes
    - Users can create multiple charts
    - Users can add other users to charts, and control access rights
    - Stacks can be custom made, or installed
  - Account
    - Accounts control a series of Stacks
    - Accounts are linked to a Kubernetes namespace
  - Application
    - An application represents a singular deployment
    - Can be given a live URL
    - Deployed using Cloud Foundry Build Packs
    - Can be linked to a stack via Git


