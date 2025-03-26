#!/bin/bash

# Define the project name
PROJECT_NAME="azure-web-forum"

# Create the folder structure
mkdir -p ./{frontend,backend,infra,.github,docs}

# Initialize AngularJS project
cd frontend
npx -y @angular/cli@latest new angular-app --defaults
mv angular-app/* . && mv angular-app/.* . 2>/dev/null
rmdir angular-app
echo "node_modules/" >> .gitignore
cd ..

# Initialize Golang project
cd backend
go mod init github.com/example/$PROJECT_NAME
mkdir cmd internal pkg
echo -e "bin/\n*.exe\n*.out\n*.test" >> .gitignore
cd ..

# Initialize Terraform project
cd infra
touch main.tf variables.tf outputs.tf
terraform init
echo -e ".terraform/\n*.tfstate\n*.tfstate.backup\nterraform.tfvars" >> .gitignore
cd ..

# Create global .gitignore
echo -e "node_modules/\n.DS_Store\nterraform.tfstate*\n.terraform/\n*.log\nbin/\n*.exe\n*.out\n*.test" > .gitignore

# Initialize Git repository
git init
git add .
git commit -m "Initial project setup"

echo "Project structure created and initialized successfully!"

