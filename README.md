# Deploy a Go Web App to Azure Container Apps

Learn how to use Docker, Azure Container Registry, and Azure Container Apps to deploy a Go web ap to Azure.

## Getting Started

### Prerequisites

- An Azure account with an active subscription.
- Docker Desktop or other runtime installed.
- Go 1.18 or higher installed.

### Download the sample 

1. git clone https://github.com/Azure-Samples/msdocs-go-webapp-quickstart.git
2. cd msdocs-go-webapp-quickstart


### Create an Azure Container Registry

1. Create an Azure resource group.
    ```bash
    az group create --name <resourceGroupName> --location eastus
    ```
2. Create an Azure Container Registry.
    ```bash
    az acr create --resource-group <resourceGroupName> --name <acrName> --sku basic --admin-enabled true
    ```
3. Log into the Azure container instance.
    ```bash
    az acr login --name <acrName>  
    ```

### Build and push the Docker image


1. Build and push the docker image to ACR.
    ```bash
    docker build -t <loginServer>/<imageName>:latest
    ```
2. Push the docker image to ACR.
    ```bash
    docker push <loginServer>/<imageName>:latest
    ```
3. Verify the image was deployed to ACR.
    ```bash
    az acr repository list --name <acrName> --output table
    ```

### Deploy the Azure Container App

1. Get the ARC admin password.
    ```bash
    password=$(az acr credential show -n <acrName> --query 'passwords[0].value' --out tsv)
    ```
2. Create a container app environment.
    ```bash
    az containerapp env create \
    --name <containerAppEnvName> \
    --resource-group <resourceGroupName> \
    --location "East US"
    ```
3. Deploy the container app.
    ```bash
    az containerapp create \
    --name <containerAppName> \
    --resource-group <resourceGroupName> \
    --environment <goWebAppContainerAppEnv> \
    --image "<loginServer>/<imageName>:latest" \
    --registry-server "<loginServer>" \
    --registry-username "<acrName" \
    --registry-password "$password" \
    --target-port 8080 \
    --ingress 'external'
    ```

## Resources

(Any additional resources or related projects)

- Link to supporting information
- Link to similar sample
- ...
