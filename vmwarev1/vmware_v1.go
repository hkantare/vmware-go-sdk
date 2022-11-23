/**
 * (C) Copyright IBM Corp. 2022.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/*
 * IBM OpenAPI SDK Code Generator Version: 3.62.0-a2a22f95-20221115-162524
 */

// Package vmwarev1 : Operations and models for the VmwareV1 service
package vmwarev1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	common "github.com/hkantare/vmware-go-sdk/common"
)

// VmwareV1 : IBM Cloud for VMware as a Service API
//
// API Version: 1.0
type VmwareV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://vmware.cloud.ibm.com/v1"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "vmware"

// VmwareV1Options : Service options
type VmwareV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewVmwareV1UsingExternalConfig : constructs an instance of VmwareV1 with passed in options and external configuration.
func NewVmwareV1UsingExternalConfig(options *VmwareV1Options) (vmware *VmwareV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	vmware, err = NewVmwareV1(options)
	if err != nil {
		return
	}

	err = vmware.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = vmware.Service.SetServiceURL(options.URL)
	}
	return
}

// NewVmwareV1 : constructs an instance of VmwareV1 with passed in options.
func NewVmwareV1(options *VmwareV1Options) (service *VmwareV1, err error) {
	serviceOptions := &core.ServiceOptions{
		URL:           DefaultServiceURL,
		Authenticator: options.Authenticator,
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			return
		}
	}

	service = &VmwareV1{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "vmware" suitable for processing requests.
func (vmware *VmwareV1) Clone() *VmwareV1 {
	if core.IsNil(vmware) {
		return nil
	}
	clone := *vmware
	clone.Service = vmware.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (vmware *VmwareV1) SetServiceURL(url string) error {
	return vmware.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (vmware *VmwareV1) GetServiceURL() string {
	return vmware.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (vmware *VmwareV1) SetDefaultHeaders(headers http.Header) {
	vmware.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (vmware *VmwareV1) SetEnableGzipCompression(enableGzip bool) {
	vmware.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (vmware *VmwareV1) GetEnableGzipCompression() bool {
	return vmware.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (vmware *VmwareV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	vmware.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (vmware *VmwareV1) DisableRetries() {
	vmware.Service.DisableRetries()
}

// CreateWorkloadDomain : Create a director site instance
// Create a new instance of a director site with specified configurations. The director site instance is the
// infrastructure and associated VMware software stack consisting of vCenter, NSX-T, and VMware Cloud Director. VMware
// platform management and operations are performed with VMware Cloud Director. The minimum initial order size is 2
// hosts (2-Socket 32 Cores, 192 GB RAM) with 24 TB of 2.0 IOPS/GB storage.
func (vmware *VmwareV1) CreateWorkloadDomain(createWorkloadDomainOptions *CreateWorkloadDomainOptions) (result *DirectorSite, response *core.DetailedResponse, err error) {
	return vmware.CreateWorkloadDomainWithContext(context.Background(), createWorkloadDomainOptions)
}

// CreateWorkloadDomainWithContext is an alternate form of the CreateWorkloadDomain method which supports a Context parameter
func (vmware *VmwareV1) CreateWorkloadDomainWithContext(ctx context.Context, createWorkloadDomainOptions *CreateWorkloadDomainOptions) (result *DirectorSite, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createWorkloadDomainOptions, "createWorkloadDomainOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createWorkloadDomainOptions, "createWorkloadDomainOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = vmware.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(vmware.Service.Options.URL, `/director_sites`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createWorkloadDomainOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("vmware", "V1", "CreateWorkloadDomain")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createWorkloadDomainOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*createWorkloadDomainOptions.AcceptLanguage))
	}
	if createWorkloadDomainOptions.XGlobalTransactionID != nil {
		builder.AddHeader("X-Global-Transaction-ID", fmt.Sprint(*createWorkloadDomainOptions.XGlobalTransactionID))
	}

	body := make(map[string]interface{})
	if createWorkloadDomainOptions.Name != nil {
		body["name"] = createWorkloadDomainOptions.Name
	}
	if createWorkloadDomainOptions.ResourceGroup != nil {
		body["resource_group"] = createWorkloadDomainOptions.ResourceGroup
	}
	if createWorkloadDomainOptions.Clusters != nil {
		body["clusters"] = createWorkloadDomainOptions.Clusters
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = vmware.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDirectorSite)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListWorkloadDomainInstances : List director site instances
// List all VMware director site instances that the user can access in the cloud account.
func (vmware *VmwareV1) ListWorkloadDomainInstances(listWorkloadDomainInstancesOptions *ListWorkloadDomainInstancesOptions) (result *ListDirectorSites, response *core.DetailedResponse, err error) {
	return vmware.ListWorkloadDomainInstancesWithContext(context.Background(), listWorkloadDomainInstancesOptions)
}

// ListWorkloadDomainInstancesWithContext is an alternate form of the ListWorkloadDomainInstances method which supports a Context parameter
func (vmware *VmwareV1) ListWorkloadDomainInstancesWithContext(ctx context.Context, listWorkloadDomainInstancesOptions *ListWorkloadDomainInstancesOptions) (result *ListDirectorSites, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listWorkloadDomainInstancesOptions, "listWorkloadDomainInstancesOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = vmware.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(vmware.Service.Options.URL, `/director_sites`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listWorkloadDomainInstancesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("vmware", "V1", "ListWorkloadDomainInstances")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listWorkloadDomainInstancesOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*listWorkloadDomainInstancesOptions.AcceptLanguage))
	}
	if listWorkloadDomainInstancesOptions.XGlobalTransactionID != nil {
		builder.AddHeader("X-Global-Transaction-ID", fmt.Sprint(*listWorkloadDomainInstancesOptions.XGlobalTransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = vmware.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListDirectorSites)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetSpecificWorkloadDomainInstance : Get a director site instance
// Get a director site instance by specifying the instance ID.
func (vmware *VmwareV1) GetSpecificWorkloadDomainInstance(getSpecificWorkloadDomainInstanceOptions *GetSpecificWorkloadDomainInstanceOptions) (result *DirectorSite, response *core.DetailedResponse, err error) {
	return vmware.GetSpecificWorkloadDomainInstanceWithContext(context.Background(), getSpecificWorkloadDomainInstanceOptions)
}

// GetSpecificWorkloadDomainInstanceWithContext is an alternate form of the GetSpecificWorkloadDomainInstance method which supports a Context parameter
func (vmware *VmwareV1) GetSpecificWorkloadDomainInstanceWithContext(ctx context.Context, getSpecificWorkloadDomainInstanceOptions *GetSpecificWorkloadDomainInstanceOptions) (result *DirectorSite, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getSpecificWorkloadDomainInstanceOptions, "getSpecificWorkloadDomainInstanceOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getSpecificWorkloadDomainInstanceOptions, "getSpecificWorkloadDomainInstanceOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"site_id": *getSpecificWorkloadDomainInstanceOptions.SiteID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = vmware.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(vmware.Service.Options.URL, `/director_sites/{site_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getSpecificWorkloadDomainInstanceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("vmware", "V1", "GetSpecificWorkloadDomainInstance")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getSpecificWorkloadDomainInstanceOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*getSpecificWorkloadDomainInstanceOptions.AcceptLanguage))
	}
	if getSpecificWorkloadDomainInstanceOptions.XGlobalTransactionID != nil {
		builder.AddHeader("X-Global-Transaction-ID", fmt.Sprint(*getSpecificWorkloadDomainInstanceOptions.XGlobalTransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = vmware.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDirectorSite)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteWorkloadDomain : Delete a director site instance
// Delete a director site instance by specifying the instance ID.
func (vmware *VmwareV1) DeleteWorkloadDomain(deleteWorkloadDomainOptions *DeleteWorkloadDomainOptions) (result *DirectorSite, response *core.DetailedResponse, err error) {
	return vmware.DeleteWorkloadDomainWithContext(context.Background(), deleteWorkloadDomainOptions)
}

// DeleteWorkloadDomainWithContext is an alternate form of the DeleteWorkloadDomain method which supports a Context parameter
func (vmware *VmwareV1) DeleteWorkloadDomainWithContext(ctx context.Context, deleteWorkloadDomainOptions *DeleteWorkloadDomainOptions) (result *DirectorSite, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteWorkloadDomainOptions, "deleteWorkloadDomainOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteWorkloadDomainOptions, "deleteWorkloadDomainOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"site_id": *deleteWorkloadDomainOptions.SiteID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = vmware.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(vmware.Service.Options.URL, `/director_sites/{site_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteWorkloadDomainOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("vmware", "V1", "DeleteWorkloadDomain")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if deleteWorkloadDomainOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*deleteWorkloadDomainOptions.AcceptLanguage))
	}
	if deleteWorkloadDomainOptions.XGlobalTransactionID != nil {
		builder.AddHeader("X-Global-Transaction-ID", fmt.Sprint(*deleteWorkloadDomainOptions.XGlobalTransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = vmware.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDirectorSite)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListClusterInstances : List clusters
// List all VMware clusters of a director site instance by specifying the ID of the instance.
func (vmware *VmwareV1) ListClusterInstances(listClusterInstancesOptions *ListClusterInstancesOptions) (result *ListClusters, response *core.DetailedResponse, err error) {
	return vmware.ListClusterInstancesWithContext(context.Background(), listClusterInstancesOptions)
}

// ListClusterInstancesWithContext is an alternate form of the ListClusterInstances method which supports a Context parameter
func (vmware *VmwareV1) ListClusterInstancesWithContext(ctx context.Context, listClusterInstancesOptions *ListClusterInstancesOptions) (result *ListClusters, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listClusterInstancesOptions, "listClusterInstancesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listClusterInstancesOptions, "listClusterInstancesOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"site_id": *listClusterInstancesOptions.SiteID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = vmware.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(vmware.Service.Options.URL, `/director_sites/{site_id}/clusters`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listClusterInstancesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("vmware", "V1", "ListClusterInstances")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listClusterInstancesOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*listClusterInstancesOptions.AcceptLanguage))
	}
	if listClusterInstancesOptions.XGlobalTransactionID != nil {
		builder.AddHeader("X-Global-Transaction-ID", fmt.Sprint(*listClusterInstancesOptions.XGlobalTransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = vmware.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListClusters)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetSpecificClusterInstance : Get a cluster
// Get a specific VMware cluster from the director site instance.
func (vmware *VmwareV1) GetSpecificClusterInstance(getSpecificClusterInstanceOptions *GetSpecificClusterInstanceOptions) (result *Cluster, response *core.DetailedResponse, err error) {
	return vmware.GetSpecificClusterInstanceWithContext(context.Background(), getSpecificClusterInstanceOptions)
}

// GetSpecificClusterInstanceWithContext is an alternate form of the GetSpecificClusterInstance method which supports a Context parameter
func (vmware *VmwareV1) GetSpecificClusterInstanceWithContext(ctx context.Context, getSpecificClusterInstanceOptions *GetSpecificClusterInstanceOptions) (result *Cluster, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getSpecificClusterInstanceOptions, "getSpecificClusterInstanceOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getSpecificClusterInstanceOptions, "getSpecificClusterInstanceOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"site_id":    *getSpecificClusterInstanceOptions.SiteID,
		"cluster_id": *getSpecificClusterInstanceOptions.ClusterID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = vmware.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(vmware.Service.Options.URL, `/director_sites/{site_id}/clusters/{cluster_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getSpecificClusterInstanceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("vmware", "V1", "GetSpecificClusterInstance")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getSpecificClusterInstanceOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*getSpecificClusterInstanceOptions.AcceptLanguage))
	}
	if getSpecificClusterInstanceOptions.XGlobalTransactionID != nil {
		builder.AddHeader("X-Global-Transaction-ID", fmt.Sprint(*getSpecificClusterInstanceOptions.XGlobalTransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = vmware.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCluster)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// SetHostsCount : Update the number of hosts of a cluster
// Update the number of hosts of a specific cluster in a specific director site instance. VMware clusters must have
// between [2-25] hosts.
func (vmware *VmwareV1) SetHostsCount(setHostsCountOptions *SetHostsCountOptions) (result *SetHostsCountResponse, response *core.DetailedResponse, err error) {
	return vmware.SetHostsCountWithContext(context.Background(), setHostsCountOptions)
}

// SetHostsCountWithContext is an alternate form of the SetHostsCount method which supports a Context parameter
func (vmware *VmwareV1) SetHostsCountWithContext(ctx context.Context, setHostsCountOptions *SetHostsCountOptions) (result *SetHostsCountResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(setHostsCountOptions, "setHostsCountOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(setHostsCountOptions, "setHostsCountOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"site_id":    *setHostsCountOptions.SiteID,
		"cluster_id": *setHostsCountOptions.ClusterID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = vmware.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(vmware.Service.Options.URL, `/director_sites/{site_id}/clusters/{cluster_id}/hosts_count`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range setHostsCountOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("vmware", "V1", "SetHostsCount")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if setHostsCountOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*setHostsCountOptions.AcceptLanguage))
	}
	if setHostsCountOptions.XGlobalTransactionID != nil {
		builder.AddHeader("X-Global-Transaction-ID", fmt.Sprint(*setHostsCountOptions.XGlobalTransactionID))
	}

	body := make(map[string]interface{})
	if setHostsCountOptions.Count != nil {
		body["count"] = setHostsCountOptions.Count
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = vmware.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSetHostsCountResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// SetFileShares : Update the file storage shares of a cluster
// Update the file_shares of a specific cluster in a specific director site instance.
func (vmware *VmwareV1) SetFileShares(setFileSharesOptions *SetFileSharesOptions) (result *FileShares, response *core.DetailedResponse, err error) {
	return vmware.SetFileSharesWithContext(context.Background(), setFileSharesOptions)
}

// SetFileSharesWithContext is an alternate form of the SetFileShares method which supports a Context parameter
func (vmware *VmwareV1) SetFileSharesWithContext(ctx context.Context, setFileSharesOptions *SetFileSharesOptions) (result *FileShares, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(setFileSharesOptions, "setFileSharesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(setFileSharesOptions, "setFileSharesOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"site_id":    *setFileSharesOptions.SiteID,
		"cluster_id": *setFileSharesOptions.ClusterID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = vmware.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(vmware.Service.Options.URL, `/director_sites/{site_id}/clusters/{cluster_id}/file_shares`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range setFileSharesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("vmware", "V1", "SetFileShares")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if setFileSharesOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*setFileSharesOptions.AcceptLanguage))
	}
	if setFileSharesOptions.XGlobalTransactionID != nil {
		builder.AddHeader("X-Global-Transaction-ID", fmt.Sprint(*setFileSharesOptions.XGlobalTransactionID))
	}

	body := make(map[string]interface{})
	if setFileSharesOptions.STORAGEPOINTTWOFIVEIOPSGB != nil {
		body["STORAGE_POINT_TWO_FIVE_IOPS_GB"] = setFileSharesOptions.STORAGEPOINTTWOFIVEIOPSGB
	}
	if setFileSharesOptions.STORAGETWOIOPSGB != nil {
		body["STORAGE_TWO_IOPS_GB"] = setFileSharesOptions.STORAGETWOIOPSGB
	}
	if setFileSharesOptions.STORAGEFOURIOPSGB != nil {
		body["STORAGE_FOUR_IOPS_GB"] = setFileSharesOptions.STORAGEFOURIOPSGB
	}
	if setFileSharesOptions.STORAGETENIOPSGB != nil {
		body["STORAGE_TEN_IOPS_GB"] = setFileSharesOptions.STORAGETENIOPSGB
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = vmware.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalFileShares)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetRegions : List regions
// List all IBM Cloud regions enabled for users to create a new director site instance.
func (vmware *VmwareV1) GetRegions(getRegionsOptions *GetRegionsOptions) (result *DirectorSiteRegions, response *core.DetailedResponse, err error) {
	return vmware.GetRegionsWithContext(context.Background(), getRegionsOptions)
}

// GetRegionsWithContext is an alternate form of the GetRegions method which supports a Context parameter
func (vmware *VmwareV1) GetRegionsWithContext(ctx context.Context, getRegionsOptions *GetRegionsOptions) (result *DirectorSiteRegions, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getRegionsOptions, "getRegionsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = vmware.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(vmware.Service.Options.URL, `/director_site_regions`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range getRegionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("vmware", "V1", "GetRegions")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getRegionsOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*getRegionsOptions.AcceptLanguage))
	}
	if getRegionsOptions.XGlobalTransactionID != nil {
		builder.AddHeader("X-Global-Transaction-ID", fmt.Sprint(*getRegionsOptions.XGlobalTransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = vmware.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDirectorSiteRegions)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ViewInstance : List host profiles
// List available host profiles that could be used when creating a director site instance. IBM Cloud offers several
// different host types. Typically, the host type is selected based on the properties of the workload to be run in the
// VMware cluster.
func (vmware *VmwareV1) ViewInstance(viewInstanceOptions *ViewInstanceOptions) (result *ListHostProfiles, response *core.DetailedResponse, err error) {
	return vmware.ViewInstanceWithContext(context.Background(), viewInstanceOptions)
}

// ViewInstanceWithContext is an alternate form of the ViewInstance method which supports a Context parameter
func (vmware *VmwareV1) ViewInstanceWithContext(ctx context.Context, viewInstanceOptions *ViewInstanceOptions) (result *ListHostProfiles, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(viewInstanceOptions, "viewInstanceOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = vmware.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(vmware.Service.Options.URL, `/director_site_host_profiles`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range viewInstanceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("vmware", "V1", "ViewInstance")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if viewInstanceOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*viewInstanceOptions.AcceptLanguage))
	}
	if viewInstanceOptions.XGlobalTransactionID != nil {
		builder.AddHeader("X-Global-Transaction-ID", fmt.Sprint(*viewInstanceOptions.XGlobalTransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = vmware.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListHostProfiles)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ReplaceOrgAdminPassword : Replace the password of VMware Cloud Director tenant portal
// Replace the admin password used to log on to the VMware Cloud Director tenant portal and return the new value. VMware
// Cloud Director has its own authentication and authorization model. The first time that you access the VMware Cloud
// Director console you must set the admin credentials to generate an initial, complex, and random password. After the
// first admin password is generated, the VMware Cloud Director console option is enabled on the VDC details page. IBM
// Cloud does not capture the password. If the password is lost it needs to be reset.
func (vmware *VmwareV1) ReplaceOrgAdminPassword(replaceOrgAdminPasswordOptions *ReplaceOrgAdminPasswordOptions) (result *NewPassword, response *core.DetailedResponse, err error) {
	return vmware.ReplaceOrgAdminPasswordWithContext(context.Background(), replaceOrgAdminPasswordOptions)
}

// ReplaceOrgAdminPasswordWithContext is an alternate form of the ReplaceOrgAdminPassword method which supports a Context parameter
func (vmware *VmwareV1) ReplaceOrgAdminPasswordWithContext(ctx context.Context, replaceOrgAdminPasswordOptions *ReplaceOrgAdminPasswordOptions) (result *NewPassword, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(replaceOrgAdminPasswordOptions, "replaceOrgAdminPasswordOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(replaceOrgAdminPasswordOptions, "replaceOrgAdminPasswordOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = vmware.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(vmware.Service.Options.URL, `/director_site_password`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range replaceOrgAdminPasswordOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("vmware", "V1", "ReplaceOrgAdminPassword")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("site_id", fmt.Sprint(*replaceOrgAdminPasswordOptions.SiteID))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = vmware.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalNewPassword)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListPrices : List billing metrics
// List all billing metrics and associated prices.
func (vmware *VmwareV1) ListPrices(listPricesOptions *ListPricesOptions) (result *DirectorSitePricingInfo, response *core.DetailedResponse, err error) {
	return vmware.ListPricesWithContext(context.Background(), listPricesOptions)
}

// ListPricesWithContext is an alternate form of the ListPrices method which supports a Context parameter
func (vmware *VmwareV1) ListPricesWithContext(ctx context.Context, listPricesOptions *ListPricesOptions) (result *DirectorSitePricingInfo, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listPricesOptions, "listPricesOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = vmware.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(vmware.Service.Options.URL, `/director_site_pricing`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listPricesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("vmware", "V1", "ListPrices")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listPricesOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*listPricesOptions.AcceptLanguage))
	}
	if listPricesOptions.XGlobalTransactionID != nil {
		builder.AddHeader("X-Global-Transaction-ID", fmt.Sprint(*listPricesOptions.XGlobalTransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = vmware.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDirectorSitePricingInfo)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetVcddPrice : Quote price
// Quote price for a specific director site instance configuration.
func (vmware *VmwareV1) GetVcddPrice(getVcddPriceOptions *GetVcddPriceOptions) (result *DirectorSitePriceQuoteResponse, response *core.DetailedResponse, err error) {
	return vmware.GetVcddPriceWithContext(context.Background(), getVcddPriceOptions)
}

// GetVcddPriceWithContext is an alternate form of the GetVcddPrice method which supports a Context parameter
func (vmware *VmwareV1) GetVcddPriceWithContext(ctx context.Context, getVcddPriceOptions *GetVcddPriceOptions) (result *DirectorSitePriceQuoteResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getVcddPriceOptions, "getVcddPriceOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getVcddPriceOptions, "getVcddPriceOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = vmware.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(vmware.Service.Options.URL, `/director_site_price_quote`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range getVcddPriceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("vmware", "V1", "GetVcddPrice")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if getVcddPriceOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*getVcddPriceOptions.AcceptLanguage))
	}
	if getVcddPriceOptions.XGlobalTransactionID != nil {
		builder.AddHeader("X-Global-Transaction-ID", fmt.Sprint(*getVcddPriceOptions.XGlobalTransactionID))
	}

	body := make(map[string]interface{})
	if getVcddPriceOptions.Country != nil {
		body["country"] = getVcddPriceOptions.Country
	}
	if getVcddPriceOptions.Clusters != nil {
		body["clusters"] = getVcddPriceOptions.Clusters
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = vmware.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDirectorSitePriceQuoteResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListVdcs : List Virtual Data Centers
// List all Virtual Data Centers that user has access to in the cloud account.
func (vmware *VmwareV1) ListVdcs(listVdcsOptions *ListVdcsOptions) (result *ListVDCs, response *core.DetailedResponse, err error) {
	return vmware.ListVdcsWithContext(context.Background(), listVdcsOptions)
}

// ListVdcsWithContext is an alternate form of the ListVdcs method which supports a Context parameter
func (vmware *VmwareV1) ListVdcsWithContext(ctx context.Context, listVdcsOptions *ListVdcsOptions) (result *ListVDCs, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listVdcsOptions, "listVdcsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = vmware.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(vmware.Service.Options.URL, `/vdcs`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listVdcsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("vmware", "V1", "ListVdcs")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listVdcsOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*listVdcsOptions.AcceptLanguage))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = vmware.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListVDCs)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateVdc : Create a Virtual Data Center
// Create a new Virtual Data Center with specified configurations.
func (vmware *VmwareV1) CreateVdc(createVdcOptions *CreateVdcOptions) (result *VDC, response *core.DetailedResponse, err error) {
	return vmware.CreateVdcWithContext(context.Background(), createVdcOptions)
}

// CreateVdcWithContext is an alternate form of the CreateVdc method which supports a Context parameter
func (vmware *VmwareV1) CreateVdcWithContext(ctx context.Context, createVdcOptions *CreateVdcOptions) (result *VDC, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createVdcOptions, "createVdcOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createVdcOptions, "createVdcOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = vmware.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(vmware.Service.Options.URL, `/vdcs`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createVdcOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("vmware", "V1", "CreateVdc")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createVdcOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*createVdcOptions.AcceptLanguage))
	}

	body := make(map[string]interface{})
	if createVdcOptions.Name != nil {
		body["name"] = createVdcOptions.Name
	}
	if createVdcOptions.DirectorSite != nil {
		body["director_site"] = createVdcOptions.DirectorSite
	}
	if createVdcOptions.Edge != nil {
		body["edge"] = createVdcOptions.Edge
	}
	if createVdcOptions.ResourceGroup != nil {
		body["resource_group"] = createVdcOptions.ResourceGroup
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = vmware.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalVDC)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetVdc : Get a Virtual Data Center
// Get details about a Virtual Data Center by specifying the VDC ID.
func (vmware *VmwareV1) GetVdc(getVdcOptions *GetVdcOptions) (result *VDC, response *core.DetailedResponse, err error) {
	return vmware.GetVdcWithContext(context.Background(), getVdcOptions)
}

// GetVdcWithContext is an alternate form of the GetVdc method which supports a Context parameter
func (vmware *VmwareV1) GetVdcWithContext(ctx context.Context, getVdcOptions *GetVdcOptions) (result *VDC, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getVdcOptions, "getVdcOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getVdcOptions, "getVdcOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"vdc_id": *getVdcOptions.VdcID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = vmware.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(vmware.Service.Options.URL, `/vdcs/{vdc_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getVdcOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("vmware", "V1", "GetVdc")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getVdcOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*getVdcOptions.AcceptLanguage))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = vmware.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalVDC)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteVdc : Delete a Virtual Data Center
// Delete a Virtual Data Center by specifying the VDC ID.
func (vmware *VmwareV1) DeleteVdc(deleteVdcOptions *DeleteVdcOptions) (result *VDC, response *core.DetailedResponse, err error) {
	return vmware.DeleteVdcWithContext(context.Background(), deleteVdcOptions)
}

// DeleteVdcWithContext is an alternate form of the DeleteVdc method which supports a Context parameter
func (vmware *VmwareV1) DeleteVdcWithContext(ctx context.Context, deleteVdcOptions *DeleteVdcOptions) (result *VDC, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteVdcOptions, "deleteVdcOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteVdcOptions, "deleteVdcOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"vdc_id": *deleteVdcOptions.VdcID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = vmware.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(vmware.Service.Options.URL, `/vdcs/{vdc_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteVdcOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("vmware", "V1", "DeleteVdc")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if deleteVdcOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*deleteVdcOptions.AcceptLanguage))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = vmware.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalVDC)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// Cluster : A cluster resource.
type Cluster struct {
	// The cluster ID.
	ID *string `json:"id,omitempty"`

	// The cluster name.
	Name *string `json:"name,omitempty"`

	// The time that the instance is ordered.
	InstanceOrdered *string `json:"instance_ordered,omitempty"`

	// The time that the instance is created.
	InstanceCreated *string `json:"instance_created,omitempty"`

	// The time that the instance is deleted.
	InstanceDeleted *string `json:"instance_deleted,omitempty"`

	// The location of deployed cluster.
	Location *string `json:"location,omitempty"`

	// The number of hosts in the cluster.
	HostCount *int64 `json:"host_count,omitempty"`

	// The status of the director site cluster.
	Status *string `json:"status,omitempty"`

	// The ID of the director site.
	SiteID *string `json:"site_id,omitempty"`

	// The name of the host profile.
	HostProfile *string `json:"host_profile,omitempty"`

	// The storage type of the cluster.
	StorageType *string `json:"storage_type,omitempty"`

	// The billing plan for the cluster.
	BillingPlan *string `json:"billing_plan,omitempty"`

	// The chosen storage policies and their sizes.
	FileShares map[string]interface{} `json:"file_shares,omitempty"`
}

// Constants associated with the Cluster.StorageType property.
// The storage type of the cluster.
const (
	Cluster_StorageType_Nfs = "nfs"
)

// Constants associated with the Cluster.BillingPlan property.
// The billing plan for the cluster.
const (
	Cluster_BillingPlan_Monthly = "monthly"
)

// UnmarshalCluster unmarshals an instance of Cluster from the specified map of raw messages.
func UnmarshalCluster(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Cluster)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "instance_ordered", &obj.InstanceOrdered)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "instance_created", &obj.InstanceCreated)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "instance_deleted", &obj.InstanceDeleted)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "location", &obj.Location)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "host_count", &obj.HostCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "site_id", &obj.SiteID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "host_profile", &obj.HostProfile)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "storage_type", &obj.StorageType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "billing_plan", &obj.BillingPlan)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "file_shares", &obj.FileShares)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ClusterOrderInfo : VMware Cluster order information. Clusters form VMware workload availibility boundaries.
type ClusterOrderInfo struct {
	// Name of the VMware cluster. Cluster names must be unique per director site instance. Cluster names cannot be changed
	// after creation.
	Name *string `json:"name" validate:"required"`

	// Data center location to deploy the cluster. See `GET /director_site_regions` for supported data center locations.
	Location *string `json:"location" validate:"required"`

	// Number of hosts in the VMware cluster.
	HostCount *int64 `json:"host_count" validate:"required"`

	// Chosen storage policies and their sizes.
	FileShares *FileShares `json:"file_shares" validate:"required"`

	// The host type. IBM Cloud offers several different host types. Typically, the host type is selected based on the
	// properties of the workload to be run in the VMware cluster.
	HostProfile *string `json:"host_profile" validate:"required"`
}

// NewClusterOrderInfo : Instantiate ClusterOrderInfo (Generic Model Constructor)
func (*VmwareV1) NewClusterOrderInfo(name string, location string, hostCount int64, fileShares *FileShares, hostProfile string) (_model *ClusterOrderInfo, err error) {
	_model = &ClusterOrderInfo{
		Name:        core.StringPtr(name),
		Location:    core.StringPtr(location),
		HostCount:   core.Int64Ptr(hostCount),
		FileShares:  fileShares,
		HostProfile: core.StringPtr(hostProfile),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalClusterOrderInfo unmarshals an instance of ClusterOrderInfo from the specified map of raw messages.
func UnmarshalClusterOrderInfo(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ClusterOrderInfo)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "location", &obj.Location)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "host_count", &obj.HostCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "file_shares", &obj.FileShares, UnmarshalFileShares)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "host_profile", &obj.HostProfile)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ClusterSummary : The list of VMware clusters deployed on the director site.
type ClusterSummary struct {
	// The cluster ID.
	ID *string `json:"id,omitempty"`

	// The cluster name.
	Name *string `json:"name,omitempty"`

	// The location of the deployed cluster.
	Location *string `json:"location,omitempty"`

	// The number of hosts in the cluster.
	HostCount *int64 `json:"host_count,omitempty"`

	// THe cluster status.
	Status *string `json:"status,omitempty"`

	// The cluster's internal name.
	ClusterName *string `json:"cluster_name,omitempty"`

	// The name of the host profile.
	HostProfile *string `json:"host_profile,omitempty"`

	// The chosen storage policies and their sizes.
	FileShares map[string]interface{} `json:"file_shares,omitempty"`
}

// UnmarshalClusterSummary unmarshals an instance of ClusterSummary from the specified map of raw messages.
func UnmarshalClusterSummary(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ClusterSummary)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "location", &obj.Location)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "host_count", &obj.HostCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cluster_name", &obj.ClusterName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "host_profile", &obj.HostProfile)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "file_shares", &obj.FileShares)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateVdcOptions : The CreateVdc options.
type CreateVdcOptions struct {
	// A human readable identifier for the Virtual Data Center. Use a name that is unique to your region.
	Name *string `json:"name" validate:"required"`

	// The director site in which to deploy the Virtual Data Center.
	DirectorSite *NewVDCDirectorSite `json:"director_site" validate:"required"`

	// The networking Edge to be deployed on the Virtual Data Center.
	Edge *NewVDCEdge `json:"edge,omitempty"`

	// The resource group to associate with the Virtual Data Center.
	// If not specified, the default resource group in the account is used.
	ResourceGroup *NewVDCResourceGroup `json:"resource_group,omitempty"`

	// Language.
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateVdcOptions : Instantiate CreateVdcOptions
func (*VmwareV1) NewCreateVdcOptions(name string, directorSite *NewVDCDirectorSite) *CreateVdcOptions {
	return &CreateVdcOptions{
		Name:         core.StringPtr(name),
		DirectorSite: directorSite,
	}
}

// SetName : Allow user to set Name
func (_options *CreateVdcOptions) SetName(name string) *CreateVdcOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDirectorSite : Allow user to set DirectorSite
func (_options *CreateVdcOptions) SetDirectorSite(directorSite *NewVDCDirectorSite) *CreateVdcOptions {
	_options.DirectorSite = directorSite
	return _options
}

// SetEdge : Allow user to set Edge
func (_options *CreateVdcOptions) SetEdge(edge *NewVDCEdge) *CreateVdcOptions {
	_options.Edge = edge
	return _options
}

// SetResourceGroup : Allow user to set ResourceGroup
func (_options *CreateVdcOptions) SetResourceGroup(resourceGroup *NewVDCResourceGroup) *CreateVdcOptions {
	_options.ResourceGroup = resourceGroup
	return _options
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (_options *CreateVdcOptions) SetAcceptLanguage(acceptLanguage string) *CreateVdcOptions {
	_options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateVdcOptions) SetHeaders(param map[string]string) *CreateVdcOptions {
	options.Headers = param
	return options
}

// CreateWorkloadDomainOptions : The CreateWorkloadDomain options.
type CreateWorkloadDomainOptions struct {
	// Name of the director site instance. Use a name that is unique to your region and meaningful. Names cannot be changed
	// after initial creation.
	Name *string `json:"name" validate:"required"`

	// The name or ID of the IBM resource group where the instance is deployed.
	ResourceGroup *string `json:"resource_group" validate:"required"`

	// List of VMware clusters to deploy on the instance. Clusters form VMware workload availibility boundaries.
	Clusters []ClusterOrderInfo `json:"clusters" validate:"required"`

	// Language.
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// Transaction id.
	XGlobalTransactionID *string `json:"X-Global-Transaction-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateWorkloadDomainOptions : Instantiate CreateWorkloadDomainOptions
func (*VmwareV1) NewCreateWorkloadDomainOptions(name string, resourceGroup string, clusters []ClusterOrderInfo) *CreateWorkloadDomainOptions {
	return &CreateWorkloadDomainOptions{
		Name:          core.StringPtr(name),
		ResourceGroup: core.StringPtr(resourceGroup),
		Clusters:      clusters,
	}
}

// SetName : Allow user to set Name
func (_options *CreateWorkloadDomainOptions) SetName(name string) *CreateWorkloadDomainOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetResourceGroup : Allow user to set ResourceGroup
func (_options *CreateWorkloadDomainOptions) SetResourceGroup(resourceGroup string) *CreateWorkloadDomainOptions {
	_options.ResourceGroup = core.StringPtr(resourceGroup)
	return _options
}

// SetClusters : Allow user to set Clusters
func (_options *CreateWorkloadDomainOptions) SetClusters(clusters []ClusterOrderInfo) *CreateWorkloadDomainOptions {
	_options.Clusters = clusters
	return _options
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (_options *CreateWorkloadDomainOptions) SetAcceptLanguage(acceptLanguage string) *CreateWorkloadDomainOptions {
	_options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return _options
}

// SetXGlobalTransactionID : Allow user to set XGlobalTransactionID
func (_options *CreateWorkloadDomainOptions) SetXGlobalTransactionID(xGlobalTransactionID string) *CreateWorkloadDomainOptions {
	_options.XGlobalTransactionID = core.StringPtr(xGlobalTransactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateWorkloadDomainOptions) SetHeaders(param map[string]string) *CreateWorkloadDomainOptions {
	options.Headers = param
	return options
}

// DataCenterInfo : Details of the data center.
type DataCenterInfo struct {
	// The display name of the data center.
	DisplayName *string `json:"display_name,omitempty"`

	// The name of the data center.
	Name *string `json:"name,omitempty"`

	// The speed available per data center.
	UplinkSpeed *string `json:"uplink_speed,omitempty"`
}

// UnmarshalDataCenterInfo unmarshals an instance of DataCenterInfo from the specified map of raw messages.
func UnmarshalDataCenterInfo(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataCenterInfo)
	err = core.UnmarshalPrimitive(m, "display_name", &obj.DisplayName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "uplink_speed", &obj.UplinkSpeed)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeleteVdcOptions : The DeleteVdc options.
type DeleteVdcOptions struct {
	// A unique identifier for a given Virtual Data Center.
	VdcID *string `json:"vdc_id" validate:"required,ne="`

	// Language.
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteVdcOptions : Instantiate DeleteVdcOptions
func (*VmwareV1) NewDeleteVdcOptions(vdcID string) *DeleteVdcOptions {
	return &DeleteVdcOptions{
		VdcID: core.StringPtr(vdcID),
	}
}

// SetVdcID : Allow user to set VdcID
func (_options *DeleteVdcOptions) SetVdcID(vdcID string) *DeleteVdcOptions {
	_options.VdcID = core.StringPtr(vdcID)
	return _options
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (_options *DeleteVdcOptions) SetAcceptLanguage(acceptLanguage string) *DeleteVdcOptions {
	_options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteVdcOptions) SetHeaders(param map[string]string) *DeleteVdcOptions {
	options.Headers = param
	return options
}

// DeleteWorkloadDomainOptions : The DeleteWorkloadDomain options.
type DeleteWorkloadDomainOptions struct {
	// A unique identifier for the director site in which the Virtual Data Center was created.
	SiteID *string `json:"site_id" validate:"required,ne="`

	// Language.
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// Transaction id.
	XGlobalTransactionID *string `json:"X-Global-Transaction-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteWorkloadDomainOptions : Instantiate DeleteWorkloadDomainOptions
func (*VmwareV1) NewDeleteWorkloadDomainOptions(siteID string) *DeleteWorkloadDomainOptions {
	return &DeleteWorkloadDomainOptions{
		SiteID: core.StringPtr(siteID),
	}
}

// SetSiteID : Allow user to set SiteID
func (_options *DeleteWorkloadDomainOptions) SetSiteID(siteID string) *DeleteWorkloadDomainOptions {
	_options.SiteID = core.StringPtr(siteID)
	return _options
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (_options *DeleteWorkloadDomainOptions) SetAcceptLanguage(acceptLanguage string) *DeleteWorkloadDomainOptions {
	_options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return _options
}

// SetXGlobalTransactionID : Allow user to set XGlobalTransactionID
func (_options *DeleteWorkloadDomainOptions) SetXGlobalTransactionID(xGlobalTransactionID string) *DeleteWorkloadDomainOptions {
	_options.XGlobalTransactionID = core.StringPtr(xGlobalTransactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteWorkloadDomainOptions) SetHeaders(param map[string]string) *DeleteWorkloadDomainOptions {
	options.Headers = param
	return options
}

// DirectorSite : A director site resource. The director site instance is the infrastructure and associated VMware software stack
// consisting of vCenter, NSX-T and VMware Cloud Director.
type DirectorSite struct {
	// ID of the director site.
	ID *string `json:"id,omitempty"`

	// The time that the instance is ordered.
	InstanceOrdered *string `json:"instance_ordered,omitempty"`

	// The time that the instance is created and available to use.
	InstanceCreated *string `json:"instance_created,omitempty"`

	// The name of director site. The name of the director site cannot be changed after creation.
	Name *string `json:"name,omitempty"`

	// The status of director site.
	Status *string `json:"status,omitempty"`

	// The name of the IBM resource group.
	ResourceGroup *string `json:"resource_group,omitempty"`

	// The email identity of the user that ordered the VMware as a Service director site instance.
	Requester *string `json:"requester,omitempty"`

	// The ID of the resource group.
	ResourceGroupID *string `json:"resource_group_id,omitempty"`

	// The CRN of the resource group.
	ResourceGroupCrn *string `json:"resource_group_crn,omitempty"`

	// The VMware clusters on the director site.
	Clusters []ClusterSummary `json:"clusters,omitempty"`
}

// Constants associated with the DirectorSite.Status property.
// The status of director site.
const (
	DirectorSite_Status_Creating   = "Creating"
	DirectorSite_Status_Deleted    = "Deleted"
	DirectorSite_Status_Deleting   = "Deleting"
	DirectorSite_Status_Readytouse = "ReadyToUse"
	DirectorSite_Status_Updating   = "Updating"
)

// UnmarshalDirectorSite unmarshals an instance of DirectorSite from the specified map of raw messages.
func UnmarshalDirectorSite(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DirectorSite)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "instance_ordered", &obj.InstanceOrdered)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "instance_created", &obj.InstanceCreated)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_group", &obj.ResourceGroup)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "requester", &obj.Requester)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_group_id", &obj.ResourceGroupID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_group_crn", &obj.ResourceGroupCrn)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "clusters", &obj.Clusters, UnmarshalClusterSummary)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DirectorSitePriceItem : sub items for a metric and associated prices.
type DirectorSitePriceItem struct {
	// The price for the metric.
	Price *float64 `json:"price,omitempty"`

	// Quantity tier.
	QuantityTier *int64 `json:"quantity_tier,omitempty"`
}

// UnmarshalDirectorSitePriceItem unmarshals an instance of DirectorSitePriceItem from the specified map of raw messages.
func UnmarshalDirectorSitePriceItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DirectorSitePriceItem)
	err = core.UnmarshalPrimitive(m, "price", &obj.Price)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "quantity_tier", &obj.QuantityTier)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DirectorSitePriceListItem : items for a metric and associated prices.
type DirectorSitePriceListItem struct {
	// The country for which this price applies.
	Country *string `json:"country,omitempty"`

	// The unit of currency for this price.
	Currency *string `json:"currency,omitempty"`

	// A list of prices.
	Prices []DirectorSitePriceItem `json:"prices,omitempty"`
}

// UnmarshalDirectorSitePriceListItem unmarshals an instance of DirectorSitePriceListItem from the specified map of raw messages.
func UnmarshalDirectorSitePriceListItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DirectorSitePriceListItem)
	err = core.UnmarshalPrimitive(m, "country", &obj.Country)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "currency", &obj.Currency)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "prices", &obj.Prices, UnmarshalDirectorSitePriceItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DirectorSitePriceMetric : A metric and associated prices.
type DirectorSitePriceMetric struct {
	// The metric name.
	Metric *string `json:"metric,omitempty"`

	// The metric description.
	Description *string `json:"description,omitempty"`

	// A list of prices for each country.
	PriceList []DirectorSitePriceListItem `json:"price_list,omitempty"`
}

// UnmarshalDirectorSitePriceMetric unmarshals an instance of DirectorSitePriceMetric from the specified map of raw messages.
func UnmarshalDirectorSitePriceMetric(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DirectorSitePriceMetric)
	err = core.UnmarshalPrimitive(m, "metric", &obj.Metric)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "price_list", &obj.PriceList, UnmarshalDirectorSitePriceListItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DirectorSitePriceQuoteClusterInfo : VCDD instance cluster information for price calculation.
type DirectorSitePriceQuoteClusterInfo struct {
	// The name of the cluster.
	Name *string `json:"name,omitempty"`

	// The metric for the host configuration.
	HostProfile *string `json:"host_profile,omitempty"`

	// The number of hosts.
	HostCount *int64 `json:"host_count,omitempty"`

	// Chosen storage policies and their sizes.
	FileShares *FileShares `json:"file_shares,omitempty"`
}

// UnmarshalDirectorSitePriceQuoteClusterInfo unmarshals an instance of DirectorSitePriceQuoteClusterInfo from the specified map of raw messages.
func UnmarshalDirectorSitePriceQuoteClusterInfo(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DirectorSitePriceQuoteClusterInfo)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "host_profile", &obj.HostProfile)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "host_count", &obj.HostCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "file_shares", &obj.FileShares, UnmarshalFileShares)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DirectorSitePriceQuoteResponse : Return price information for a VCDD instance.
type DirectorSitePriceQuoteResponse struct {
	// Details of the instance base charge.
	BaseCharge *PriceInfoBaseCharge `json:"base_charge,omitempty"`

	// A list of the clusters with price information.
	Clusters []PriceInfoClusterCharge `json:"clusters,omitempty"`

	// The currency unit for this price.
	Currency *string `json:"currency,omitempty"`

	// The total price for the instance.
	Total *float64 `json:"total,omitempty"`
}

// UnmarshalDirectorSitePriceQuoteResponse unmarshals an instance of DirectorSitePriceQuoteResponse from the specified map of raw messages.
func UnmarshalDirectorSitePriceQuoteResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DirectorSitePriceQuoteResponse)
	err = core.UnmarshalModel(m, "base_charge", &obj.BaseCharge, UnmarshalPriceInfoBaseCharge)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "clusters", &obj.Clusters, UnmarshalPriceInfoClusterCharge)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "currency", &obj.Currency)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total", &obj.Total)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DirectorSitePricingInfo : Return all metrics with associate prices.
type DirectorSitePricingInfo struct {
	// A list of metrics and associated prices.
	DirectorSitePricing []DirectorSitePriceMetric `json:"director_site_pricing,omitempty"`
}

// UnmarshalDirectorSitePricingInfo unmarshals an instance of DirectorSitePricingInfo from the specified map of raw messages.
func UnmarshalDirectorSitePricingInfo(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DirectorSitePricingInfo)
	err = core.UnmarshalModel(m, "director_site_pricing", &obj.DirectorSitePricing, UnmarshalDirectorSitePriceMetric)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DirectorSiteRegions : Success. The request was successfully processed.
type DirectorSiteRegions struct {
	// available region.
	DirectorSiteRegions map[string]RegionDetail `json:"director_site_regions,omitempty"`
}

// UnmarshalDirectorSiteRegions unmarshals an instance of DirectorSiteRegions from the specified map of raw messages.
func UnmarshalDirectorSiteRegions(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DirectorSiteRegions)
	err = core.UnmarshalModel(m, "director_site_regions", &obj.DirectorSiteRegions, UnmarshalRegionDetail)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Edge : A networking Edge deployed on a Virtual Data Center. Networking edges are based on NSX-T and used for bridging
// virtualize networking to the physical public-internet and IBM private networking.
type Edge struct {
	// A unique identifier for the Edge.
	ID *string `json:"id" validate:"required"`

	// The public IP addresses assigned to the Edge.
	PublicIps []string `json:"public_ips" validate:"required"`

	// The size of the Edge.
	//
	// The size can only be specified for dedicated Edges. Larger sizes require more capacity from the director site in
	// which the Virtual Data Center was created to be deployed.
	Size *string `json:"size,omitempty"`

	// The type of Edge to be deployed.
	//
	// Shared Edges allow for multiple VDCs to share some Edge resources. Dedicated Edges do not share resources between
	// VDCs.
	Type *string `json:"type" validate:"required"`
}

// Constants associated with the Edge.Size property.
// The size of the Edge.
//
// The size can only be specified for dedicated Edges. Larger sizes require more capacity from the director site in
// which the Virtual Data Center was created to be deployed.
const (
	Edge_Size_ExtraLarge = "extra_large"
	Edge_Size_Large      = "large"
	Edge_Size_Medium     = "medium"
)

// Constants associated with the Edge.Type property.
// The type of Edge to be deployed.
//
// Shared Edges allow for multiple VDCs to share some Edge resources. Dedicated Edges do not share resources between
// VDCs.
const (
	Edge_Type_Dedicated = "dedicated"
	Edge_Type_Shared    = "shared"
)

// UnmarshalEdge unmarshals an instance of Edge from the specified map of raw messages.
func UnmarshalEdge(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Edge)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "public_ips", &obj.PublicIps)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "size", &obj.Size)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Error : Information about why a request cannot be completed or why a resource could not be created.
type Error struct {
	// An error code specific to the error encountered.
	Code *string `json:"code" validate:"required"`

	// A message describing why the error ocurred.
	Message *string `json:"message" validate:"required"`

	// A URL that links to a page with more information about this error.
	MoreInfo *string `json:"more_info,omitempty"`
}

// UnmarshalError unmarshals an instance of Error from the specified map of raw messages.
func UnmarshalError(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Error)
	err = core.UnmarshalPrimitive(m, "code", &obj.Code)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "more_info", &obj.MoreInfo)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// FileShares : Chosen storage policies and their sizes.
type FileShares struct {
	// The amount of 0.25 IOPS/GB storage in GB (1024^3 bytes).
	STORAGEPOINTTWOFIVEIOPSGB *int64 `json:"STORAGE_POINT_TWO_FIVE_IOPS_GB,omitempty"`

	// The amount of 2 IOPS/GB storage in GB (1024^3 bytes).
	STORAGETWOIOPSGB *int64 `json:"STORAGE_TWO_IOPS_GB,omitempty"`

	// The amount of 4 IOPS/GB storage in GB (1024^3 bytes).
	STORAGEFOURIOPSGB *int64 `json:"STORAGE_FOUR_IOPS_GB,omitempty"`

	// The amount of 10 IOPS/GB storage in GB (1024^3 bytes).
	STORAGETENIOPSGB *int64 `json:"STORAGE_TEN_IOPS_GB,omitempty"`
}

// UnmarshalFileShares unmarshals an instance of FileShares from the specified map of raw messages.
func UnmarshalFileShares(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FileShares)
	err = core.UnmarshalPrimitive(m, "STORAGE_POINT_TWO_FIVE_IOPS_GB", &obj.STORAGEPOINTTWOFIVEIOPSGB)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "STORAGE_TWO_IOPS_GB", &obj.STORAGETWOIOPSGB)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "STORAGE_FOUR_IOPS_GB", &obj.STORAGEFOURIOPSGB)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "STORAGE_TEN_IOPS_GB", &obj.STORAGETENIOPSGB)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetRegionsOptions : The GetRegions options.
type GetRegionsOptions struct {
	// Language.
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// Transaction id.
	XGlobalTransactionID *string `json:"X-Global-Transaction-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetRegionsOptions : Instantiate GetRegionsOptions
func (*VmwareV1) NewGetRegionsOptions() *GetRegionsOptions {
	return &GetRegionsOptions{}
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (_options *GetRegionsOptions) SetAcceptLanguage(acceptLanguage string) *GetRegionsOptions {
	_options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return _options
}

// SetXGlobalTransactionID : Allow user to set XGlobalTransactionID
func (_options *GetRegionsOptions) SetXGlobalTransactionID(xGlobalTransactionID string) *GetRegionsOptions {
	_options.XGlobalTransactionID = core.StringPtr(xGlobalTransactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetRegionsOptions) SetHeaders(param map[string]string) *GetRegionsOptions {
	options.Headers = param
	return options
}

// GetSpecificClusterInstanceOptions : The GetSpecificClusterInstance options.
type GetSpecificClusterInstanceOptions struct {
	// A unique identifier for the director site in which the Virtual Data Center was created.
	SiteID *string `json:"site_id" validate:"required,ne="`

	// The cluster to query.
	ClusterID *string `json:"cluster_id" validate:"required,ne="`

	// Language.
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// Transaction id.
	XGlobalTransactionID *string `json:"X-Global-Transaction-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetSpecificClusterInstanceOptions : Instantiate GetSpecificClusterInstanceOptions
func (*VmwareV1) NewGetSpecificClusterInstanceOptions(siteID string, clusterID string) *GetSpecificClusterInstanceOptions {
	return &GetSpecificClusterInstanceOptions{
		SiteID:    core.StringPtr(siteID),
		ClusterID: core.StringPtr(clusterID),
	}
}

// SetSiteID : Allow user to set SiteID
func (_options *GetSpecificClusterInstanceOptions) SetSiteID(siteID string) *GetSpecificClusterInstanceOptions {
	_options.SiteID = core.StringPtr(siteID)
	return _options
}

// SetClusterID : Allow user to set ClusterID
func (_options *GetSpecificClusterInstanceOptions) SetClusterID(clusterID string) *GetSpecificClusterInstanceOptions {
	_options.ClusterID = core.StringPtr(clusterID)
	return _options
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (_options *GetSpecificClusterInstanceOptions) SetAcceptLanguage(acceptLanguage string) *GetSpecificClusterInstanceOptions {
	_options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return _options
}

// SetXGlobalTransactionID : Allow user to set XGlobalTransactionID
func (_options *GetSpecificClusterInstanceOptions) SetXGlobalTransactionID(xGlobalTransactionID string) *GetSpecificClusterInstanceOptions {
	_options.XGlobalTransactionID = core.StringPtr(xGlobalTransactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetSpecificClusterInstanceOptions) SetHeaders(param map[string]string) *GetSpecificClusterInstanceOptions {
	options.Headers = param
	return options
}

// GetSpecificWorkloadDomainInstanceOptions : The GetSpecificWorkloadDomainInstance options.
type GetSpecificWorkloadDomainInstanceOptions struct {
	// A unique identifier for the director site in which the Virtual Data Center was created.
	SiteID *string `json:"site_id" validate:"required,ne="`

	// Language.
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// Transaction id.
	XGlobalTransactionID *string `json:"X-Global-Transaction-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetSpecificWorkloadDomainInstanceOptions : Instantiate GetSpecificWorkloadDomainInstanceOptions
func (*VmwareV1) NewGetSpecificWorkloadDomainInstanceOptions(siteID string) *GetSpecificWorkloadDomainInstanceOptions {
	return &GetSpecificWorkloadDomainInstanceOptions{
		SiteID: core.StringPtr(siteID),
	}
}

// SetSiteID : Allow user to set SiteID
func (_options *GetSpecificWorkloadDomainInstanceOptions) SetSiteID(siteID string) *GetSpecificWorkloadDomainInstanceOptions {
	_options.SiteID = core.StringPtr(siteID)
	return _options
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (_options *GetSpecificWorkloadDomainInstanceOptions) SetAcceptLanguage(acceptLanguage string) *GetSpecificWorkloadDomainInstanceOptions {
	_options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return _options
}

// SetXGlobalTransactionID : Allow user to set XGlobalTransactionID
func (_options *GetSpecificWorkloadDomainInstanceOptions) SetXGlobalTransactionID(xGlobalTransactionID string) *GetSpecificWorkloadDomainInstanceOptions {
	_options.XGlobalTransactionID = core.StringPtr(xGlobalTransactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetSpecificWorkloadDomainInstanceOptions) SetHeaders(param map[string]string) *GetSpecificWorkloadDomainInstanceOptions {
	options.Headers = param
	return options
}

// GetVcddPriceOptions : The GetVcddPrice options.
type GetVcddPriceOptions struct {
	// String representing the billing country.
	Country *string `json:"country,omitempty"`

	// The list of clusters.
	Clusters []DirectorSitePriceQuoteClusterInfo `json:"clusters,omitempty"`

	// Language.
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// Transaction id.
	XGlobalTransactionID *string `json:"X-Global-Transaction-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetVcddPriceOptions : Instantiate GetVcddPriceOptions
func (*VmwareV1) NewGetVcddPriceOptions() *GetVcddPriceOptions {
	return &GetVcddPriceOptions{}
}

// SetCountry : Allow user to set Country
func (_options *GetVcddPriceOptions) SetCountry(country string) *GetVcddPriceOptions {
	_options.Country = core.StringPtr(country)
	return _options
}

// SetClusters : Allow user to set Clusters
func (_options *GetVcddPriceOptions) SetClusters(clusters []DirectorSitePriceQuoteClusterInfo) *GetVcddPriceOptions {
	_options.Clusters = clusters
	return _options
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (_options *GetVcddPriceOptions) SetAcceptLanguage(acceptLanguage string) *GetVcddPriceOptions {
	_options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return _options
}

// SetXGlobalTransactionID : Allow user to set XGlobalTransactionID
func (_options *GetVcddPriceOptions) SetXGlobalTransactionID(xGlobalTransactionID string) *GetVcddPriceOptions {
	_options.XGlobalTransactionID = core.StringPtr(xGlobalTransactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetVcddPriceOptions) SetHeaders(param map[string]string) *GetVcddPriceOptions {
	options.Headers = param
	return options
}

// GetVdcOptions : The GetVdc options.
type GetVdcOptions struct {
	// A unique identifier for a given Virtual Data Center.
	VdcID *string `json:"vdc_id" validate:"required,ne="`

	// Language.
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetVdcOptions : Instantiate GetVdcOptions
func (*VmwareV1) NewGetVdcOptions(vdcID string) *GetVdcOptions {
	return &GetVdcOptions{
		VdcID: core.StringPtr(vdcID),
	}
}

// SetVdcID : Allow user to set VdcID
func (_options *GetVdcOptions) SetVdcID(vdcID string) *GetVdcOptions {
	_options.VdcID = core.StringPtr(vdcID)
	return _options
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (_options *GetVdcOptions) SetAcceptLanguage(acceptLanguage string) *GetVdcOptions {
	_options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetVdcOptions) SetHeaders(param map[string]string) *GetVdcOptions {
	options.Headers = param
	return options
}

// HostProfile : Host profile template.
type HostProfile struct {
	// The name for this host profile.
	ProfileName *string `json:"profile_name,omitempty"`

	// The CPU type for this host profile.
	CpuType *string `json:"cpu_type,omitempty"`

	// The number of CPUs for this host profile.
	CpuCount *int64 `json:"cpu_count,omitempty"`

	// The RAM for this host profile in GB (1024^3 bytes).
	Ram *int64 `json:"ram,omitempty"`

	// The collection of the host profile disks.
	LocalDisks []HostProfileDisk `json:"local_disks,omitempty"`
}

// UnmarshalHostProfile unmarshals an instance of HostProfile from the specified map of raw messages.
func UnmarshalHostProfile(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(HostProfile)
	err = core.UnmarshalPrimitive(m, "profile_name", &obj.ProfileName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cpu_type", &obj.CpuType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cpu_count", &obj.CpuCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ram", &obj.Ram)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "local_disks", &obj.LocalDisks, UnmarshalHostProfileDisk)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// HostProfileDisk : The host profile disk description. The host profile disk defines the properties of local disks in the host.
type HostProfileDisk struct {
	// The number of disks of this configuration for an instance with this profile.
	Quantity *int64 `json:"quantity,omitempty"`

	// The disk size in GB (1024^3 bytes).
	Size *int64 `json:"size,omitempty"`

	// The disk type and specifications.
	Type *string `json:"type,omitempty"`
}

// UnmarshalHostProfileDisk unmarshals an instance of HostProfileDisk from the specified map of raw messages.
func UnmarshalHostProfileDisk(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(HostProfileDisk)
	err = core.UnmarshalPrimitive(m, "quantity", &obj.Quantity)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "size", &obj.Size)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListClusterInstancesOptions : The ListClusterInstances options.
type ListClusterInstancesOptions struct {
	// A unique identifier for the director site in which the Virtual Data Center was created.
	SiteID *string `json:"site_id" validate:"required,ne="`

	// Language.
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// Transaction id.
	XGlobalTransactionID *string `json:"X-Global-Transaction-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListClusterInstancesOptions : Instantiate ListClusterInstancesOptions
func (*VmwareV1) NewListClusterInstancesOptions(siteID string) *ListClusterInstancesOptions {
	return &ListClusterInstancesOptions{
		SiteID: core.StringPtr(siteID),
	}
}

// SetSiteID : Allow user to set SiteID
func (_options *ListClusterInstancesOptions) SetSiteID(siteID string) *ListClusterInstancesOptions {
	_options.SiteID = core.StringPtr(siteID)
	return _options
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (_options *ListClusterInstancesOptions) SetAcceptLanguage(acceptLanguage string) *ListClusterInstancesOptions {
	_options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return _options
}

// SetXGlobalTransactionID : Allow user to set XGlobalTransactionID
func (_options *ListClusterInstancesOptions) SetXGlobalTransactionID(xGlobalTransactionID string) *ListClusterInstancesOptions {
	_options.XGlobalTransactionID = core.StringPtr(xGlobalTransactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListClusterInstancesOptions) SetHeaders(param map[string]string) *ListClusterInstancesOptions {
	options.Headers = param
	return options
}

// ListClusters : Return all clusters instances.
type ListClusters struct {
	// list of cluster objects.
	Clusters []Cluster `json:"clusters,omitempty"`
}

// UnmarshalListClusters unmarshals an instance of ListClusters from the specified map of raw messages.
func UnmarshalListClusters(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListClusters)
	err = core.UnmarshalModel(m, "clusters", &obj.Clusters, UnmarshalCluster)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListDirectorSites : Return all director site instances.
type ListDirectorSites struct {
	// List of director site instances.
	DirectorSites []DirectorSite `json:"director_sites,omitempty"`
}

// UnmarshalListDirectorSites unmarshals an instance of ListDirectorSites from the specified map of raw messages.
func UnmarshalListDirectorSites(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListDirectorSites)
	err = core.UnmarshalModel(m, "director_sites", &obj.DirectorSites, UnmarshalDirectorSite)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListHostProfiles : Success. The request was successfully processed.
type ListHostProfiles struct {
	// The list of available host profiles.
	DirectorSiteHostProfiles []HostProfile `json:"director_site_host_profiles,omitempty"`
}

// UnmarshalListHostProfiles unmarshals an instance of ListHostProfiles from the specified map of raw messages.
func UnmarshalListHostProfiles(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListHostProfiles)
	err = core.UnmarshalModel(m, "director_site_host_profiles", &obj.DirectorSiteHostProfiles, UnmarshalHostProfile)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListPricesOptions : The ListPrices options.
type ListPricesOptions struct {
	// Language.
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// Transaction id.
	XGlobalTransactionID *string `json:"X-Global-Transaction-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListPricesOptions : Instantiate ListPricesOptions
func (*VmwareV1) NewListPricesOptions() *ListPricesOptions {
	return &ListPricesOptions{}
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (_options *ListPricesOptions) SetAcceptLanguage(acceptLanguage string) *ListPricesOptions {
	_options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return _options
}

// SetXGlobalTransactionID : Allow user to set XGlobalTransactionID
func (_options *ListPricesOptions) SetXGlobalTransactionID(xGlobalTransactionID string) *ListPricesOptions {
	_options.XGlobalTransactionID = core.StringPtr(xGlobalTransactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListPricesOptions) SetHeaders(param map[string]string) *ListPricesOptions {
	options.Headers = param
	return options
}

// ListVDCs : A list of Virtual Data Centers.
type ListVDCs struct {
	// A List of Virtual Data Centers.
	Vdcs []VDC `json:"vdcs" validate:"required"`
}

// UnmarshalListVDCs unmarshals an instance of ListVDCs from the specified map of raw messages.
func UnmarshalListVDCs(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListVDCs)
	err = core.UnmarshalModel(m, "vdcs", &obj.Vdcs, UnmarshalVDC)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListVdcsOptions : The ListVdcs options.
type ListVdcsOptions struct {
	// Language.
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListVdcsOptions : Instantiate ListVdcsOptions
func (*VmwareV1) NewListVdcsOptions() *ListVdcsOptions {
	return &ListVdcsOptions{}
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (_options *ListVdcsOptions) SetAcceptLanguage(acceptLanguage string) *ListVdcsOptions {
	_options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListVdcsOptions) SetHeaders(param map[string]string) *ListVdcsOptions {
	options.Headers = param
	return options
}

// ListWorkloadDomainInstancesOptions : The ListWorkloadDomainInstances options.
type ListWorkloadDomainInstancesOptions struct {
	// Language.
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// Transaction id.
	XGlobalTransactionID *string `json:"X-Global-Transaction-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListWorkloadDomainInstancesOptions : Instantiate ListWorkloadDomainInstancesOptions
func (*VmwareV1) NewListWorkloadDomainInstancesOptions() *ListWorkloadDomainInstancesOptions {
	return &ListWorkloadDomainInstancesOptions{}
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (_options *ListWorkloadDomainInstancesOptions) SetAcceptLanguage(acceptLanguage string) *ListWorkloadDomainInstancesOptions {
	_options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return _options
}

// SetXGlobalTransactionID : Allow user to set XGlobalTransactionID
func (_options *ListWorkloadDomainInstancesOptions) SetXGlobalTransactionID(xGlobalTransactionID string) *ListWorkloadDomainInstancesOptions {
	_options.XGlobalTransactionID = core.StringPtr(xGlobalTransactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListWorkloadDomainInstancesOptions) SetHeaders(param map[string]string) *ListWorkloadDomainInstancesOptions {
	options.Headers = param
	return options
}

// NewPassword : The new admin password used to log in to the VMware Cloud Director tenant portal. VMware Cloud Director has its own
// internal authentication and authorization model. The previous Director admin password is reset to a newly generated
// random password.
type NewPassword struct {
	// The password used to log in to the VMware Cloud Director tenant portal.
	Password *string `json:"password" validate:"required"`
}

// UnmarshalNewPassword unmarshals an instance of NewPassword from the specified map of raw messages.
func UnmarshalNewPassword(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(NewPassword)
	err = core.UnmarshalPrimitive(m, "password", &obj.Password)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// NewVDCDirectorSite : The director site in which to deploy the Virtual Data Center.
type NewVDCDirectorSite struct {
	// A unique identifier for the director site.
	ID *string `json:"id" validate:"required"`

	// The cluster within the director site in which to deploy the Virtual Data Center.
	Cluster *VDCDirectorSiteCluster `json:"cluster" validate:"required"`
}

// NewNewVDCDirectorSite : Instantiate NewVDCDirectorSite (Generic Model Constructor)
func (*VmwareV1) NewNewVDCDirectorSite(id string, cluster *VDCDirectorSiteCluster) (_model *NewVDCDirectorSite, err error) {
	_model = &NewVDCDirectorSite{
		ID:      core.StringPtr(id),
		Cluster: cluster,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalNewVDCDirectorSite unmarshals an instance of NewVDCDirectorSite from the specified map of raw messages.
func UnmarshalNewVDCDirectorSite(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(NewVDCDirectorSite)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "cluster", &obj.Cluster, UnmarshalVDCDirectorSiteCluster)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// NewVDCEdge : The networking Edge to be deployed on the Virtual Data Center.
type NewVDCEdge struct {
	// The size of the Edge. Only used for Edges of type dedicated.
	Size *string `json:"size,omitempty"`

	// The type of Edge to be deployed on the Virtual Data Center.
	Type *string `json:"type" validate:"required"`
}

// Constants associated with the NewVDCEdge.Size property.
// The size of the Edge. Only used for Edges of type dedicated.
const (
	NewVDCEdge_Size_ExtraLarge = "extra_large"
	NewVDCEdge_Size_Large      = "large"
	NewVDCEdge_Size_Medium     = "medium"
)

// Constants associated with the NewVDCEdge.Type property.
// The type of Edge to be deployed on the Virtual Data Center.
const (
	NewVDCEdge_Type_Dedicated = "dedicated"
	NewVDCEdge_Type_Shared    = "shared"
)

// NewNewVDCEdge : Instantiate NewVDCEdge (Generic Model Constructor)
func (*VmwareV1) NewNewVDCEdge(typeVar string) (_model *NewVDCEdge, err error) {
	_model = &NewVDCEdge{
		Type: core.StringPtr(typeVar),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalNewVDCEdge unmarshals an instance of NewVDCEdge from the specified map of raw messages.
func UnmarshalNewVDCEdge(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(NewVDCEdge)
	err = core.UnmarshalPrimitive(m, "size", &obj.Size)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// NewVDCResourceGroup : The resource group to associate with the Virtual Data Center. If not specified, the default resource group in the
// account is used.
type NewVDCResourceGroup struct {
	// A unique identifier for the resource group.
	ID *string `json:"id" validate:"required"`
}

// NewNewVDCResourceGroup : Instantiate NewVDCResourceGroup (Generic Model Constructor)
func (*VmwareV1) NewNewVDCResourceGroup(id string) (_model *NewVDCResourceGroup, err error) {
	_model = &NewVDCResourceGroup{
		ID: core.StringPtr(id),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalNewVDCResourceGroup unmarshals an instance of NewVDCResourceGroup from the specified map of raw messages.
func UnmarshalNewVDCResourceGroup(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(NewVDCResourceGroup)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PriceInfoBaseCharge : Details of the instance base charge.
type PriceInfoBaseCharge struct {
	// The name of the metric that is being charged.
	Name *string `json:"name,omitempty"`

	// The unit of currency for this pric.
	Currency *string `json:"currency,omitempty"`

	// The price for this metric.
	Price *float64 `json:"price,omitempty"`
}

// UnmarshalPriceInfoBaseCharge unmarshals an instance of PriceInfoBaseCharge from the specified map of raw messages.
func UnmarshalPriceInfoBaseCharge(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PriceInfoBaseCharge)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "currency", &obj.Currency)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "price", &obj.Price)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PriceInfoClusterCharge : A cluster for the instance and its price information.
type PriceInfoClusterCharge struct {
	// The cluster name.
	Name *string `json:"name,omitempty"`

	// The unit of currency for this price.
	Currency *string `json:"currency,omitempty"`

	// The total price for this cluster.
	Price *float64 `json:"price,omitempty"`

	// A list of items that make up the cluster and their price information.
	Items []PriceInfoClusterItem `json:"items,omitempty"`
}

// UnmarshalPriceInfoClusterCharge unmarshals an instance of PriceInfoClusterCharge from the specified map of raw messages.
func UnmarshalPriceInfoClusterCharge(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PriceInfoClusterCharge)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "currency", &obj.Currency)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "price", &obj.Price)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "items", &obj.Items, UnmarshalPriceInfoClusterItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PriceInfoClusterItem : items for VCDD instance cluster price information.
type PriceInfoClusterItem struct {
	// The item name.
	Name *string `json:"name,omitempty"`

	// The unit of currency for this price.
	Currency *string `json:"currency,omitempty"`

	// The total price for this item.
	Price *float64 `json:"price,omitempty"`

	// A list of subitems and their price information.
	Items []PriceInfoClusterSubItem `json:"items,omitempty"`
}

// UnmarshalPriceInfoClusterItem unmarshals an instance of PriceInfoClusterItem from the specified map of raw messages.
func UnmarshalPriceInfoClusterItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PriceInfoClusterItem)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "currency", &obj.Currency)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "price", &obj.Price)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "items", &obj.Items, UnmarshalPriceInfoClusterSubItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PriceInfoClusterSubItem : sub items for VCDD instance cluster price information.
type PriceInfoClusterSubItem struct {
	// The metric that is being charged.
	Name *string `json:"name,omitempty"`

	// The number of items that this metric will be charged.
	Count *int64 `json:"count,omitempty"`

	// The unit of currency for this price.
	Currency *string `json:"currency,omitempty"`

	// The price for a single charge of this metric.
	Price *float64 `json:"price,omitempty"`
}

// UnmarshalPriceInfoClusterSubItem unmarshals an instance of PriceInfoClusterSubItem from the specified map of raw messages.
func UnmarshalPriceInfoClusterSubItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PriceInfoClusterSubItem)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "count", &obj.Count)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "currency", &obj.Currency)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "price", &obj.Price)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RegionDetail : The region details.
type RegionDetail struct {
	// The datacenter details.
	Datacenters []DataCenterInfo `json:"datacenters,omitempty"`

	// Accessable endpoint of the region.
	Endpoint *string `json:"endpoint,omitempty"`
}

// UnmarshalRegionDetail unmarshals an instance of RegionDetail from the specified map of raw messages.
func UnmarshalRegionDetail(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RegionDetail)
	err = core.UnmarshalModel(m, "datacenters", &obj.Datacenters, UnmarshalDataCenterInfo)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "endpoint", &obj.Endpoint)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ReplaceOrgAdminPasswordOptions : The ReplaceOrgAdminPassword options.
type ReplaceOrgAdminPasswordOptions struct {
	// A unique identifier for the director site.
	SiteID *string `json:"site_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewReplaceOrgAdminPasswordOptions : Instantiate ReplaceOrgAdminPasswordOptions
func (*VmwareV1) NewReplaceOrgAdminPasswordOptions(siteID string) *ReplaceOrgAdminPasswordOptions {
	return &ReplaceOrgAdminPasswordOptions{
		SiteID: core.StringPtr(siteID),
	}
}

// SetSiteID : Allow user to set SiteID
func (_options *ReplaceOrgAdminPasswordOptions) SetSiteID(siteID string) *ReplaceOrgAdminPasswordOptions {
	_options.SiteID = core.StringPtr(siteID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ReplaceOrgAdminPasswordOptions) SetHeaders(param map[string]string) *ReplaceOrgAdminPasswordOptions {
	options.Headers = param
	return options
}

// SetFileSharesOptions : The SetFileShares options.
type SetFileSharesOptions struct {
	// A unique identifier for the director site in which the Virtual Data Center was created.
	SiteID *string `json:"site_id" validate:"required,ne="`

	// The cluster to query.
	ClusterID *string `json:"cluster_id" validate:"required,ne="`

	// The amount of 0.25 IOPS/GB storage in GB (1024^3 bytes).
	STORAGEPOINTTWOFIVEIOPSGB *int64 `json:"STORAGE_POINT_TWO_FIVE_IOPS_GB,omitempty"`

	// The amount of 2 IOPS/GB storage in GB (1024^3 bytes).
	STORAGETWOIOPSGB *int64 `json:"STORAGE_TWO_IOPS_GB,omitempty"`

	// The amount of 4 IOPS/GB storage in GB (1024^3 bytes).
	STORAGEFOURIOPSGB *int64 `json:"STORAGE_FOUR_IOPS_GB,omitempty"`

	// The amount of 10 IOPS/GB storage in GB (1024^3 bytes).
	STORAGETENIOPSGB *int64 `json:"STORAGE_TEN_IOPS_GB,omitempty"`

	// Language.
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// Transaction id.
	XGlobalTransactionID *string `json:"X-Global-Transaction-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewSetFileSharesOptions : Instantiate SetFileSharesOptions
func (*VmwareV1) NewSetFileSharesOptions(siteID string, clusterID string) *SetFileSharesOptions {
	return &SetFileSharesOptions{
		SiteID:    core.StringPtr(siteID),
		ClusterID: core.StringPtr(clusterID),
	}
}

// SetSiteID : Allow user to set SiteID
func (_options *SetFileSharesOptions) SetSiteID(siteID string) *SetFileSharesOptions {
	_options.SiteID = core.StringPtr(siteID)
	return _options
}

// SetClusterID : Allow user to set ClusterID
func (_options *SetFileSharesOptions) SetClusterID(clusterID string) *SetFileSharesOptions {
	_options.ClusterID = core.StringPtr(clusterID)
	return _options
}

// SetSTORAGEPOINTTWOFIVEIOPSGB : Allow user to set STORAGEPOINTTWOFIVEIOPSGB
func (_options *SetFileSharesOptions) SetSTORAGEPOINTTWOFIVEIOPSGB(sTORAGEPOINTTWOFIVEIOPSGB int64) *SetFileSharesOptions {
	_options.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(sTORAGEPOINTTWOFIVEIOPSGB)
	return _options
}

// SetSTORAGETWOIOPSGB : Allow user to set STORAGETWOIOPSGB
func (_options *SetFileSharesOptions) SetSTORAGETWOIOPSGB(sTORAGETWOIOPSGB int64) *SetFileSharesOptions {
	_options.STORAGETWOIOPSGB = core.Int64Ptr(sTORAGETWOIOPSGB)
	return _options
}

// SetSTORAGEFOURIOPSGB : Allow user to set STORAGEFOURIOPSGB
func (_options *SetFileSharesOptions) SetSTORAGEFOURIOPSGB(sTORAGEFOURIOPSGB int64) *SetFileSharesOptions {
	_options.STORAGEFOURIOPSGB = core.Int64Ptr(sTORAGEFOURIOPSGB)
	return _options
}

// SetSTORAGETENIOPSGB : Allow user to set STORAGETENIOPSGB
func (_options *SetFileSharesOptions) SetSTORAGETENIOPSGB(sTORAGETENIOPSGB int64) *SetFileSharesOptions {
	_options.STORAGETENIOPSGB = core.Int64Ptr(sTORAGETENIOPSGB)
	return _options
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (_options *SetFileSharesOptions) SetAcceptLanguage(acceptLanguage string) *SetFileSharesOptions {
	_options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return _options
}

// SetXGlobalTransactionID : Allow user to set XGlobalTransactionID
func (_options *SetFileSharesOptions) SetXGlobalTransactionID(xGlobalTransactionID string) *SetFileSharesOptions {
	_options.XGlobalTransactionID = core.StringPtr(xGlobalTransactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *SetFileSharesOptions) SetHeaders(param map[string]string) *SetFileSharesOptions {
	options.Headers = param
	return options
}

// SetHostsCountOptions : The SetHostsCount options.
type SetHostsCountOptions struct {
	// A unique identifier for the director site in which the Virtual Data Center was created.
	SiteID *string `json:"site_id" validate:"required,ne="`

	// The cluster to query.
	ClusterID *string `json:"cluster_id" validate:"required,ne="`

	// count of hosts to add or remove on cluster.
	Count *int64 `json:"count" validate:"required"`

	// Language.
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// Transaction id.
	XGlobalTransactionID *string `json:"X-Global-Transaction-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewSetHostsCountOptions : Instantiate SetHostsCountOptions
func (*VmwareV1) NewSetHostsCountOptions(siteID string, clusterID string, count int64) *SetHostsCountOptions {
	return &SetHostsCountOptions{
		SiteID:    core.StringPtr(siteID),
		ClusterID: core.StringPtr(clusterID),
		Count:     core.Int64Ptr(count),
	}
}

// SetSiteID : Allow user to set SiteID
func (_options *SetHostsCountOptions) SetSiteID(siteID string) *SetHostsCountOptions {
	_options.SiteID = core.StringPtr(siteID)
	return _options
}

// SetClusterID : Allow user to set ClusterID
func (_options *SetHostsCountOptions) SetClusterID(clusterID string) *SetHostsCountOptions {
	_options.ClusterID = core.StringPtr(clusterID)
	return _options
}

// SetCount : Allow user to set Count
func (_options *SetHostsCountOptions) SetCount(count int64) *SetHostsCountOptions {
	_options.Count = core.Int64Ptr(count)
	return _options
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (_options *SetHostsCountOptions) SetAcceptLanguage(acceptLanguage string) *SetHostsCountOptions {
	_options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return _options
}

// SetXGlobalTransactionID : Allow user to set XGlobalTransactionID
func (_options *SetHostsCountOptions) SetXGlobalTransactionID(xGlobalTransactionID string) *SetHostsCountOptions {
	_options.XGlobalTransactionID = core.StringPtr(xGlobalTransactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *SetHostsCountOptions) SetHeaders(param map[string]string) *SetHostsCountOptions {
	options.Headers = param
	return options
}

// SetHostsCountResponse : Response of set hosts count.
type SetHostsCountResponse struct {
	// Information of request accepted.
	Message *string `json:"message,omitempty"`
}

// UnmarshalSetHostsCountResponse unmarshals an instance of SetHostsCountResponse from the specified map of raw messages.
func UnmarshalSetHostsCountResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SetHostsCountResponse)
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// VDC : A VMware Virtual Data Center (VDC). VMware VDCs are used to deploy and run VMware virtualized networking and run
// VMware workloads. VMware VDCs form loose boundaries of networking and workload where networking and workload can be
// shared or optionally isolated between VDCs. You can deploy one or more VDCs in an instance except when using the
// minimal instance configuration consisting of 2 hosts (2-Socket 32 Cores, 192 GB RAM). With the minimal instance
// configuration you can start with just one VDC and a performance network edge of medium size until additional hosts
// are added to the cluster.
type VDC struct {
	// A unique identifier for the Virtual Data Center.
	ID *string `json:"id" validate:"required"`

	// Determines how resources are made available to the Virtual Data Center. VMware as a Services uses the VMware Cloud
	// Director Pay-As-You-Go (paygo) allocation model. With paygo, resources are committed as they are allocated by VMware
	// vApps and VMs. IaaS resources are not reserved until vApps and VMs are specifically defined to VMware Cloud
	// Director. The paygo model supports an optimal use of resources where resources are allocated on-demand as needed
	// rather than prereserved without use.
	AllocationModel *string `json:"allocation_model" validate:"required"`

	// The time after which the Virtual Data Center is considered usable.
	CreatedTime *strfmt.DateTime `json:"created_time" validate:"required"`

	// A unique identifier for the Virtual Data Center in IBM Cloud.
	Crn *string `json:"crn" validate:"required"`

	// The time after which the Virtual Data Center is no longer considered usable.
	DeletedTime *strfmt.DateTime `json:"deleted_time" validate:"required"`

	// The director site in which to deploy the Virtual Data Center.
	DirectorSite *VDCDirectorSite `json:"director_site" validate:"required"`

	// The VMware NSX-T networking Edges deployed on the Virtual Data Center. NSX-T edges are used for bridging virtualize
	// networking to the physical public-internet and IBM private networking.
	Edges []Edge `json:"edges" validate:"required"`

	// Information about why the request to create the Virtual Data Center cannot be completed.
	Errors []Error `json:"errors" validate:"required"`

	// A human readable identifier for the Virtual Data Center.
	Name *string `json:"name" validate:"required"`

	// The time at which the request to create the Virtual Data Center was made.
	OrderedTime *strfmt.DateTime `json:"ordered_time" validate:"required"`

	// The name of the VMware Cloud Director organization containing this Virtual Data Center. VMware Cloud Director
	// organizations are used to create strong boundaries between virtual data centers. There is a complete isolation of
	// user administration, networking, workloads and VMware Cloud Director catalogs between different Director
	// organizations.
	OrgName *string `json:"org_name" validate:"required"`

	// Determines the state the Virtual Data Center is currently in.
	Status *string `json:"status" validate:"required"`

	// Determines if this Virtual Data Center is in a single-tenant or multi-tenant director site.
	Type *string `json:"type" validate:"required"`
}

// Constants associated with the VDC.AllocationModel property.
// Determines how resources are made available to the Virtual Data Center. VMware as a Services uses the VMware Cloud
// Director Pay-As-You-Go (paygo) allocation model. With paygo, resources are committed as they are allocated by VMware
// vApps and VMs. IaaS resources are not reserved until vApps and VMs are specifically defined to VMware Cloud Director.
// The paygo model supports an optimal use of resources where resources are allocated on-demand as needed rather than
// prereserved without use.
const (
	VDC_AllocationModel_Paygo = "paygo"
)

// Constants associated with the VDC.Status property.
// Determines the state the Virtual Data Center is currently in.
const (
	VDC_Status_Creating   = "Creating"
	VDC_Status_Deleted    = "Deleted"
	VDC_Status_Deleting   = "Deleting"
	VDC_Status_Failed     = "Failed"
	VDC_Status_Modifying  = "Modifying"
	VDC_Status_Readytouse = "ReadyToUse"
)

// Constants associated with the VDC.Type property.
// Determines if this Virtual Data Center is in a single-tenant or multi-tenant director site.
const (
	VDC_Type_Dedicated = "dedicated"
)

// UnmarshalVDC unmarshals an instance of VDC from the specified map of raw messages.
func UnmarshalVDC(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(VDC)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "allocation_model", &obj.AllocationModel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_time", &obj.CreatedTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "crn", &obj.Crn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "deleted_time", &obj.DeletedTime)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "director_site", &obj.DirectorSite, UnmarshalVDCDirectorSite)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "edges", &obj.Edges, UnmarshalEdge)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "errors", &obj.Errors, UnmarshalError)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ordered_time", &obj.OrderedTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "org_name", &obj.OrgName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// VDCDirectorSite : The director site in which to deploy the Virtual Data Center.
type VDCDirectorSite struct {
	// A unique identifier for the director site.
	ID *string `json:"id" validate:"required"`

	// The cluster within the director site in which to deploy the Virtual Data Center.
	Cluster *VDCDirectorSiteCluster `json:"cluster" validate:"required"`

	// The URL of the VMware Cloud Director tenant portal where this Virtual Data Center can be managed.
	URL *string `json:"url" validate:"required"`
}

// UnmarshalVDCDirectorSite unmarshals an instance of VDCDirectorSite from the specified map of raw messages.
func UnmarshalVDCDirectorSite(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(VDCDirectorSite)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "cluster", &obj.Cluster, UnmarshalVDCDirectorSiteCluster)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// VDCDirectorSiteCluster : The cluster within the director site in which to deploy the Virtual Data Center.
type VDCDirectorSiteCluster struct {
	// A unique identifier for the cluster.
	ID *string `json:"id" validate:"required"`
}

// NewVDCDirectorSiteCluster : Instantiate VDCDirectorSiteCluster (Generic Model Constructor)
func (*VmwareV1) NewVDCDirectorSiteCluster(id string) (_model *VDCDirectorSiteCluster, err error) {
	_model = &VDCDirectorSiteCluster{
		ID: core.StringPtr(id),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalVDCDirectorSiteCluster unmarshals an instance of VDCDirectorSiteCluster from the specified map of raw messages.
func UnmarshalVDCDirectorSiteCluster(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(VDCDirectorSiteCluster)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ViewInstanceOptions : The ViewInstance options.
type ViewInstanceOptions struct {
	// Language.
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// Transaction id.
	XGlobalTransactionID *string `json:"X-Global-Transaction-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewViewInstanceOptions : Instantiate ViewInstanceOptions
func (*VmwareV1) NewViewInstanceOptions() *ViewInstanceOptions {
	return &ViewInstanceOptions{}
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (_options *ViewInstanceOptions) SetAcceptLanguage(acceptLanguage string) *ViewInstanceOptions {
	_options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return _options
}

// SetXGlobalTransactionID : Allow user to set XGlobalTransactionID
func (_options *ViewInstanceOptions) SetXGlobalTransactionID(xGlobalTransactionID string) *ViewInstanceOptions {
	_options.XGlobalTransactionID = core.StringPtr(xGlobalTransactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ViewInstanceOptions) SetHeaders(param map[string]string) *ViewInstanceOptions {
	options.Headers = param
	return options
}
