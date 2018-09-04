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


API Module Structure:
- Each api module consists of 3 parts:
  - View: This represents anything that consumes the module.
    These are implemented seperately from the core module to allow
    for consumer-agnostic code
  - Core: The core logic of the module. This section cannot interface
    with the outside world, it merely exposes function endpoints. This
    module is meant to be pure, with any external resources injected as
    dependancies
    - Consists of structs to represent entities, a servicer to manage
      those entities, and interfaces for injectors to implement
  - Dependancies: Represent external services that can be used by services.
    This can be anything that a service needs to do that cant be handled
    internally (database read/writes, external api calls)