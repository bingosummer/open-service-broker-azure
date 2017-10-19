package aci

import "github.com/Azure/azure-service-broker/pkg/service"

// ProvisioningParameters encapsulates aci-specific provisioning options
type ProvisioningParameters struct {
	Location                 string            `json:"location"`
	ResourceGroup            string            `json:"resourceGroup"`
	Tags                     map[string]string `json:"tags"`
	ImageName                string            `json:"image"`
	NumberCores              string            `json:"cpuCores"`
	Memory                   string            `json:"memoryInGb"`
	Port                     string            `json:"port"`
	Sharename                string            `json:"sharename"`
	Storageaccountname       string            `json:"storageaccountname"`
	Storageaccountkey        string            `json:"storageaccountkey"`
	Volumename               string            `json:"volumename"`
	Mountpoint               string            `json:"mountpoint"`
	ImageRegistryLoginServer string            `json:"imageRegistryLoginServer"`
	ImageUsername            string            `json:"imageUsername"`
	ImagePassword            string            `json:"imagePassword"`
}

type aciProvisioningContext struct {
	ResourceGroupName string `json:"resourceGroup"`
	ARMDeploymentName string `json:"armDeployment"`
	ContainerName     string `json:"name"`
	IPAddress         string `json:"containerIPv4Address"`
}

// UpdatingParameters encapsulates aci-specific updating options
type UpdatingParameters struct {
}

// BindingParameters encapsulates aci-specific binding options
type BindingParameters struct {
}

type aciBindingContext struct {
}

type aciCredentials struct {
	IPAddress string `json:"containerIPv4Address"`
}

func (k *aciProvisioningContext) GetResourceGroupName() string {
	return k.ResourceGroupName
}

func (
	m *module,
) GetEmptyProvisioningParameters() service.ProvisioningParameters {
	return &ProvisioningParameters{}
}

func (
	m *module,
) GetEmptyUpdatingParameters() service.UpdatingParameters {
	return &UpdatingParameters{}
}

func (m *module) GetEmptyProvisioningContext() service.ProvisioningContext {
	return &aciProvisioningContext{}
}

func (m *module) GetEmptyBindingParameters() service.BindingParameters {
	return &BindingParameters{}
}

func (m *module) GetEmptyBindingContext() service.BindingContext {
	return &aciBindingContext{}
}

func (m *module) GetEmptyCredentials() service.Credentials {
	return &aciCredentials{}
}
