# Azure Web Forum

<p align="center">
  <a href="https://example.com/">
    <img src="https://via.placeholder.com/72" alt="Logo" width=72 height=72>
  </a>

  <h3 align="center">Azure Web Forum</h3>

  <p align="center">
    Simple full-stack web application hosted on Azure with full CI/CD

    <br>
    <a href="https://reponame/issues/new?template=bug.md">Report bug</a>
    ·
    <a href="https://reponame/issues/new?template=feature.md&labels=feature">Request feature</a>
  </p>
</p>

## Table of contents
- [Quick start](#quick-start)
- [Status](#status)
- [Infrastructure](#infrastructure)
- [Repo Structure](#repo-structure)
- [Bugs and feature requests](#bugs-and-feature-requests)
- [Contributing](#contributing)
- [Creators](#creators)
- [Thanks](#thanks)
- [Copyright and license](#copyright-and-license)

## Project steps
[ ] Infrastructure (Terraform)
Set up Azure Kubernetes Service (AKS), Azure Container Registry (ACR), and Storage using Terraform.

Ensure the Kubernetes cluster is running and can deploy workloads.

Configure networking (VNet, subnets, security groups) to ensure secure communication.

Why?: This ensures a solid foundation before deploying applications.

[ ] Backend (Golang API)
Define API endpoints (user authentication, posts, comments, etc.).

Implement database models and connect to a PostgreSQL instance.

Containerize the backend (Dockerfile) and push to ACR.

Why?: The backend provides the core functionality needed for the frontend.

[ ] Frontend (AngularJS)
Set up the basic structure and routing.

Implement authentication (JWT/OAuth).

Connect frontend to the backend API.

Why?: The frontend depends on the backend being available.

[ ] CI/CD Pipeline (GitHub Actions / Azure DevOps)
Automate builds, tests, and deployments for the frontend and backend.

Implement a deployment strategy (Blue-Green, Canary).

Why?: Automates deployment to reduce errors and improve efficiency.

[ ] Security & Monitoring
Set up Prometheus/Grafana for monitoring.

Implement RBAC, TLS, and logging for security.

Why?: Ensures observability and security before going live.


## Quick start
Follow these steps to set up and deploy the project:

- Install dependencies and tools.
- Set up Azure infrastructure using Terraform.
- Deploy application components.
- Configure monitoring and logging.

## Status
Project is under active development. The latest features and bug fixes are tracked in the issue board.

### Under Active Development
- frontend
- backend
- Infrastructure & hosting
- CI/CD
### Finished Steps
- Rough Project Outline
- Initaliziation of Repo


## Infrastructure
Infrastructure is deployed using Terraform, featuring:
- Azure Kubernetes Service (AKS)
- Azure Container Registry (ACR)
- Virtual Network and Storage
- CI/CD pipeline automation

## Repo Structure
Project directory structure:

```text
azure-web-forum/
├── frontend/        # AngularJS frontend
├── backend/         # Golang backend
├── infra/           # Terraform scripts - Typescript
│   ├── main.tf
│   ├── variables.tf
│   └── outputs.tf
├── .github/         # CI/CD pipeline configs
└── docs/            # Documentation
```

## Bugs and feature requests
Have a bug or a feature request? Please first read the [issue guidelines](https://reponame/blob/master/CONTRIBUTING.md) and search for existing and closed issues. If your problem or idea is not addressed yet, [please open a new issue](https://reponame/issues/new).

## Contributing
Please read through our [contributing guidelines](https://reponame/blob/master/CONTRIBUTING.md). Included are directions for opening issues, coding standards, and notes on development.

Moreover, all code should conform to the [Code Guide](https://github.com/mdo/code-guide), maintained by [Main author](https://github.com/usernamemainauthor).

Editor preferences are available in the [editor config](https://reponame/blob/master/.editorconfig) for easy use in common text editors. Read more and download plugins at <https://editorconfig.org/>.


## Creators
**Project Maintainer**
- <https://github.com/usernamecreator1>

## Thanks
Special thanks to contributors and the open-source community for their support.




