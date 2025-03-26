provider "azurerm" {
  features {}
  subscription_id = var.subscription_id
}

resource "azurerm_resource_group" "web_forum_rg" {
  name     = "azure-web-forum-rg"
  location = "East US"
}

resource "azurerm_kubernetes_cluster" "web_forum_aks" {
  name                = "web-forum-aks"
  location            = azurerm_resource_group.web_forum_rg.location
  resource_group_name = azurerm_resource_group.web_forum_rg.name
  dns_prefix          = "web-forum"

  default_node_pool {
    name       = "default"
    node_count = 2
    vm_size    = "Standard_D2s_v3"
  }

  identity {
    type = "SystemAssigned"
  }
}
