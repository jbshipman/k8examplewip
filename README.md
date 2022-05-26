# Cloud Native Example Project
Hello there!

More details in this readme to be added as I build the project.

## Problem Statement
Using a programming language of my choosing expose a REST endpoint that returns the following JSON with a current timestamp:

```json
{
    "message" : "Automate all the things!",
    "timestamp" : <current timestamp>
}
```

Deloy in a public clould using a Kubernetes cluser. The cluster provisioning and application deployment is be all via code.

## Requirements
- All code in a public git repo.
- README.md to include details instructions on how to run application, what is running, how to clean up after running.
- Provide the single command to launch the evironment and deploy the application.
- Launch and deploy should be able to be done on any public cloud account (end user's choice basically).
- Automated testing to ensure environment is properly configured and validated.

*Wizbangs and Nernits*

## Project Notes

### API Structure
Option 1:
- /things/all     (to be the route to the required payload)
- /things         (location of that which I can add to the project)

Option 2:
- root            (route to required payload)
- /things         (additional payload )