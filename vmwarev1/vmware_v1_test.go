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

package vmwarev1_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	"github.com/hkantare/vmware-go-sdk/vmwarev1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`VmwareV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(vmwareService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(vmwareService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
				URL: "https://vmwarev1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(vmwareService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"VMWARE_URL": "https://vmwarev1/api",
				"VMWARE_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				vmwareService, serviceErr := vmwarev1.NewVmwareV1UsingExternalConfig(&vmwarev1.VmwareV1Options{
				})
				Expect(vmwareService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := vmwareService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != vmwareService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(vmwareService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(vmwareService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				vmwareService, serviceErr := vmwarev1.NewVmwareV1UsingExternalConfig(&vmwarev1.VmwareV1Options{
					URL: "https://testService/api",
				})
				Expect(vmwareService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := vmwareService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != vmwareService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(vmwareService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(vmwareService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				vmwareService, serviceErr := vmwarev1.NewVmwareV1UsingExternalConfig(&vmwarev1.VmwareV1Options{
				})
				err := vmwareService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := vmwareService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != vmwareService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(vmwareService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(vmwareService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"VMWARE_URL": "https://vmwarev1/api",
				"VMWARE_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			vmwareService, serviceErr := vmwarev1.NewVmwareV1UsingExternalConfig(&vmwarev1.VmwareV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(vmwareService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"VMWARE_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			vmwareService, serviceErr := vmwarev1.NewVmwareV1UsingExternalConfig(&vmwarev1.VmwareV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(vmwareService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = vmwarev1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`CreateWorkloadDomain(createWorkloadDomainOptions *CreateWorkloadDomainOptions) - Operation response error`, func() {
		createWorkloadDomainPath := "/director_sites"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createWorkloadDomainPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateWorkloadDomain with error: Operation response processing error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the FileShares model
				fileSharesModel := new(vmwarev1.FileShares)
				fileSharesModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))

				// Construct an instance of the ClusterOrderInfo model
				clusterOrderInfoModel := new(vmwarev1.ClusterOrderInfo)
				clusterOrderInfoModel.Name = core.StringPtr("testString")
				clusterOrderInfoModel.Location = core.StringPtr("testString")
				clusterOrderInfoModel.HostCount = core.Int64Ptr(int64(2))
				clusterOrderInfoModel.FileShares = fileSharesModel
				clusterOrderInfoModel.HostProfile = core.StringPtr("testString")

				// Construct an instance of the CreateWorkloadDomainOptions model
				createWorkloadDomainOptionsModel := new(vmwarev1.CreateWorkloadDomainOptions)
				createWorkloadDomainOptionsModel.Name = core.StringPtr("testString")
				createWorkloadDomainOptionsModel.ResourceGroup = core.StringPtr("testString")
				createWorkloadDomainOptionsModel.Clusters = []vmwarev1.ClusterOrderInfo{*clusterOrderInfoModel}
				createWorkloadDomainOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createWorkloadDomainOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				createWorkloadDomainOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vmwareService.CreateWorkloadDomain(createWorkloadDomainOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vmwareService.EnableRetries(0, 0)
				result, response, operationErr = vmwareService.CreateWorkloadDomain(createWorkloadDomainOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateWorkloadDomain(createWorkloadDomainOptions *CreateWorkloadDomainOptions)`, func() {
		createWorkloadDomainPath := "/director_sites"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createWorkloadDomainPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"id": "ID", "instance_ordered": "InstanceOrdered", "instance_created": "InstanceCreated", "name": "Name", "status": "Creating", "resource_group": "ResourceGroup", "requester": "Requester", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCrn", "clusters": [{"id": "ID", "name": "Name", "location": "Location", "host_count": 9, "status": "Status", "cluster_name": "ClusterName", "host_profile": "HostProfile", "file_shares": {"anyKey": "anyValue"}}]}`)
				}))
			})
			It(`Invoke CreateWorkloadDomain successfully with retries`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())
				vmwareService.EnableRetries(0, 0)

				// Construct an instance of the FileShares model
				fileSharesModel := new(vmwarev1.FileShares)
				fileSharesModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))

				// Construct an instance of the ClusterOrderInfo model
				clusterOrderInfoModel := new(vmwarev1.ClusterOrderInfo)
				clusterOrderInfoModel.Name = core.StringPtr("testString")
				clusterOrderInfoModel.Location = core.StringPtr("testString")
				clusterOrderInfoModel.HostCount = core.Int64Ptr(int64(2))
				clusterOrderInfoModel.FileShares = fileSharesModel
				clusterOrderInfoModel.HostProfile = core.StringPtr("testString")

				// Construct an instance of the CreateWorkloadDomainOptions model
				createWorkloadDomainOptionsModel := new(vmwarev1.CreateWorkloadDomainOptions)
				createWorkloadDomainOptionsModel.Name = core.StringPtr("testString")
				createWorkloadDomainOptionsModel.ResourceGroup = core.StringPtr("testString")
				createWorkloadDomainOptionsModel.Clusters = []vmwarev1.ClusterOrderInfo{*clusterOrderInfoModel}
				createWorkloadDomainOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createWorkloadDomainOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				createWorkloadDomainOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vmwareService.CreateWorkloadDomainWithContext(ctx, createWorkloadDomainOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vmwareService.DisableRetries()
				result, response, operationErr := vmwareService.CreateWorkloadDomain(createWorkloadDomainOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vmwareService.CreateWorkloadDomainWithContext(ctx, createWorkloadDomainOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createWorkloadDomainPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"id": "ID", "instance_ordered": "InstanceOrdered", "instance_created": "InstanceCreated", "name": "Name", "status": "Creating", "resource_group": "ResourceGroup", "requester": "Requester", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCrn", "clusters": [{"id": "ID", "name": "Name", "location": "Location", "host_count": 9, "status": "Status", "cluster_name": "ClusterName", "host_profile": "HostProfile", "file_shares": {"anyKey": "anyValue"}}]}`)
				}))
			})
			It(`Invoke CreateWorkloadDomain successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vmwareService.CreateWorkloadDomain(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the FileShares model
				fileSharesModel := new(vmwarev1.FileShares)
				fileSharesModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))

				// Construct an instance of the ClusterOrderInfo model
				clusterOrderInfoModel := new(vmwarev1.ClusterOrderInfo)
				clusterOrderInfoModel.Name = core.StringPtr("testString")
				clusterOrderInfoModel.Location = core.StringPtr("testString")
				clusterOrderInfoModel.HostCount = core.Int64Ptr(int64(2))
				clusterOrderInfoModel.FileShares = fileSharesModel
				clusterOrderInfoModel.HostProfile = core.StringPtr("testString")

				// Construct an instance of the CreateWorkloadDomainOptions model
				createWorkloadDomainOptionsModel := new(vmwarev1.CreateWorkloadDomainOptions)
				createWorkloadDomainOptionsModel.Name = core.StringPtr("testString")
				createWorkloadDomainOptionsModel.ResourceGroup = core.StringPtr("testString")
				createWorkloadDomainOptionsModel.Clusters = []vmwarev1.ClusterOrderInfo{*clusterOrderInfoModel}
				createWorkloadDomainOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createWorkloadDomainOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				createWorkloadDomainOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vmwareService.CreateWorkloadDomain(createWorkloadDomainOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateWorkloadDomain with error: Operation validation and request error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the FileShares model
				fileSharesModel := new(vmwarev1.FileShares)
				fileSharesModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))

				// Construct an instance of the ClusterOrderInfo model
				clusterOrderInfoModel := new(vmwarev1.ClusterOrderInfo)
				clusterOrderInfoModel.Name = core.StringPtr("testString")
				clusterOrderInfoModel.Location = core.StringPtr("testString")
				clusterOrderInfoModel.HostCount = core.Int64Ptr(int64(2))
				clusterOrderInfoModel.FileShares = fileSharesModel
				clusterOrderInfoModel.HostProfile = core.StringPtr("testString")

				// Construct an instance of the CreateWorkloadDomainOptions model
				createWorkloadDomainOptionsModel := new(vmwarev1.CreateWorkloadDomainOptions)
				createWorkloadDomainOptionsModel.Name = core.StringPtr("testString")
				createWorkloadDomainOptionsModel.ResourceGroup = core.StringPtr("testString")
				createWorkloadDomainOptionsModel.Clusters = []vmwarev1.ClusterOrderInfo{*clusterOrderInfoModel}
				createWorkloadDomainOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createWorkloadDomainOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				createWorkloadDomainOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vmwareService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vmwareService.CreateWorkloadDomain(createWorkloadDomainOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateWorkloadDomainOptions model with no property values
				createWorkloadDomainOptionsModelNew := new(vmwarev1.CreateWorkloadDomainOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = vmwareService.CreateWorkloadDomain(createWorkloadDomainOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(202)
				}))
			})
			It(`Invoke CreateWorkloadDomain successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the FileShares model
				fileSharesModel := new(vmwarev1.FileShares)
				fileSharesModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))

				// Construct an instance of the ClusterOrderInfo model
				clusterOrderInfoModel := new(vmwarev1.ClusterOrderInfo)
				clusterOrderInfoModel.Name = core.StringPtr("testString")
				clusterOrderInfoModel.Location = core.StringPtr("testString")
				clusterOrderInfoModel.HostCount = core.Int64Ptr(int64(2))
				clusterOrderInfoModel.FileShares = fileSharesModel
				clusterOrderInfoModel.HostProfile = core.StringPtr("testString")

				// Construct an instance of the CreateWorkloadDomainOptions model
				createWorkloadDomainOptionsModel := new(vmwarev1.CreateWorkloadDomainOptions)
				createWorkloadDomainOptionsModel.Name = core.StringPtr("testString")
				createWorkloadDomainOptionsModel.ResourceGroup = core.StringPtr("testString")
				createWorkloadDomainOptionsModel.Clusters = []vmwarev1.ClusterOrderInfo{*clusterOrderInfoModel}
				createWorkloadDomainOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createWorkloadDomainOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				createWorkloadDomainOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vmwareService.CreateWorkloadDomain(createWorkloadDomainOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListWorkloadDomainInstances(listWorkloadDomainInstancesOptions *ListWorkloadDomainInstancesOptions) - Operation response error`, func() {
		listWorkloadDomainInstancesPath := "/director_sites"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listWorkloadDomainInstancesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListWorkloadDomainInstances with error: Operation response processing error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the ListWorkloadDomainInstancesOptions model
				listWorkloadDomainInstancesOptionsModel := new(vmwarev1.ListWorkloadDomainInstancesOptions)
				listWorkloadDomainInstancesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listWorkloadDomainInstancesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listWorkloadDomainInstancesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vmwareService.ListWorkloadDomainInstances(listWorkloadDomainInstancesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vmwareService.EnableRetries(0, 0)
				result, response, operationErr = vmwareService.ListWorkloadDomainInstances(listWorkloadDomainInstancesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListWorkloadDomainInstances(listWorkloadDomainInstancesOptions *ListWorkloadDomainInstancesOptions)`, func() {
		listWorkloadDomainInstancesPath := "/director_sites"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listWorkloadDomainInstancesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"director_sites": [{"id": "ID", "instance_ordered": "InstanceOrdered", "instance_created": "InstanceCreated", "name": "Name", "status": "Creating", "resource_group": "ResourceGroup", "requester": "Requester", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCrn", "clusters": [{"id": "ID", "name": "Name", "location": "Location", "host_count": 9, "status": "Status", "cluster_name": "ClusterName", "host_profile": "HostProfile", "file_shares": {"anyKey": "anyValue"}}]}]}`)
				}))
			})
			It(`Invoke ListWorkloadDomainInstances successfully with retries`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())
				vmwareService.EnableRetries(0, 0)

				// Construct an instance of the ListWorkloadDomainInstancesOptions model
				listWorkloadDomainInstancesOptionsModel := new(vmwarev1.ListWorkloadDomainInstancesOptions)
				listWorkloadDomainInstancesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listWorkloadDomainInstancesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listWorkloadDomainInstancesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vmwareService.ListWorkloadDomainInstancesWithContext(ctx, listWorkloadDomainInstancesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vmwareService.DisableRetries()
				result, response, operationErr := vmwareService.ListWorkloadDomainInstances(listWorkloadDomainInstancesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vmwareService.ListWorkloadDomainInstancesWithContext(ctx, listWorkloadDomainInstancesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listWorkloadDomainInstancesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"director_sites": [{"id": "ID", "instance_ordered": "InstanceOrdered", "instance_created": "InstanceCreated", "name": "Name", "status": "Creating", "resource_group": "ResourceGroup", "requester": "Requester", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCrn", "clusters": [{"id": "ID", "name": "Name", "location": "Location", "host_count": 9, "status": "Status", "cluster_name": "ClusterName", "host_profile": "HostProfile", "file_shares": {"anyKey": "anyValue"}}]}]}`)
				}))
			})
			It(`Invoke ListWorkloadDomainInstances successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vmwareService.ListWorkloadDomainInstances(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListWorkloadDomainInstancesOptions model
				listWorkloadDomainInstancesOptionsModel := new(vmwarev1.ListWorkloadDomainInstancesOptions)
				listWorkloadDomainInstancesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listWorkloadDomainInstancesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listWorkloadDomainInstancesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vmwareService.ListWorkloadDomainInstances(listWorkloadDomainInstancesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListWorkloadDomainInstances with error: Operation request error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the ListWorkloadDomainInstancesOptions model
				listWorkloadDomainInstancesOptionsModel := new(vmwarev1.ListWorkloadDomainInstancesOptions)
				listWorkloadDomainInstancesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listWorkloadDomainInstancesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listWorkloadDomainInstancesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vmwareService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vmwareService.ListWorkloadDomainInstances(listWorkloadDomainInstancesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListWorkloadDomainInstances successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the ListWorkloadDomainInstancesOptions model
				listWorkloadDomainInstancesOptionsModel := new(vmwarev1.ListWorkloadDomainInstancesOptions)
				listWorkloadDomainInstancesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listWorkloadDomainInstancesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listWorkloadDomainInstancesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vmwareService.ListWorkloadDomainInstances(listWorkloadDomainInstancesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSpecificWorkloadDomainInstance(getSpecificWorkloadDomainInstanceOptions *GetSpecificWorkloadDomainInstanceOptions) - Operation response error`, func() {
		getSpecificWorkloadDomainInstancePath := "/director_sites/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSpecificWorkloadDomainInstancePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetSpecificWorkloadDomainInstance with error: Operation response processing error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the GetSpecificWorkloadDomainInstanceOptions model
				getSpecificWorkloadDomainInstanceOptionsModel := new(vmwarev1.GetSpecificWorkloadDomainInstanceOptions)
				getSpecificWorkloadDomainInstanceOptionsModel.SiteID = core.StringPtr("testString")
				getSpecificWorkloadDomainInstanceOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getSpecificWorkloadDomainInstanceOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getSpecificWorkloadDomainInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vmwareService.GetSpecificWorkloadDomainInstance(getSpecificWorkloadDomainInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vmwareService.EnableRetries(0, 0)
				result, response, operationErr = vmwareService.GetSpecificWorkloadDomainInstance(getSpecificWorkloadDomainInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSpecificWorkloadDomainInstance(getSpecificWorkloadDomainInstanceOptions *GetSpecificWorkloadDomainInstanceOptions)`, func() {
		getSpecificWorkloadDomainInstancePath := "/director_sites/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSpecificWorkloadDomainInstancePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "instance_ordered": "InstanceOrdered", "instance_created": "InstanceCreated", "name": "Name", "status": "Creating", "resource_group": "ResourceGroup", "requester": "Requester", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCrn", "clusters": [{"id": "ID", "name": "Name", "location": "Location", "host_count": 9, "status": "Status", "cluster_name": "ClusterName", "host_profile": "HostProfile", "file_shares": {"anyKey": "anyValue"}}]}`)
				}))
			})
			It(`Invoke GetSpecificWorkloadDomainInstance successfully with retries`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())
				vmwareService.EnableRetries(0, 0)

				// Construct an instance of the GetSpecificWorkloadDomainInstanceOptions model
				getSpecificWorkloadDomainInstanceOptionsModel := new(vmwarev1.GetSpecificWorkloadDomainInstanceOptions)
				getSpecificWorkloadDomainInstanceOptionsModel.SiteID = core.StringPtr("testString")
				getSpecificWorkloadDomainInstanceOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getSpecificWorkloadDomainInstanceOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getSpecificWorkloadDomainInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vmwareService.GetSpecificWorkloadDomainInstanceWithContext(ctx, getSpecificWorkloadDomainInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vmwareService.DisableRetries()
				result, response, operationErr := vmwareService.GetSpecificWorkloadDomainInstance(getSpecificWorkloadDomainInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vmwareService.GetSpecificWorkloadDomainInstanceWithContext(ctx, getSpecificWorkloadDomainInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSpecificWorkloadDomainInstancePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "instance_ordered": "InstanceOrdered", "instance_created": "InstanceCreated", "name": "Name", "status": "Creating", "resource_group": "ResourceGroup", "requester": "Requester", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCrn", "clusters": [{"id": "ID", "name": "Name", "location": "Location", "host_count": 9, "status": "Status", "cluster_name": "ClusterName", "host_profile": "HostProfile", "file_shares": {"anyKey": "anyValue"}}]}`)
				}))
			})
			It(`Invoke GetSpecificWorkloadDomainInstance successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vmwareService.GetSpecificWorkloadDomainInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSpecificWorkloadDomainInstanceOptions model
				getSpecificWorkloadDomainInstanceOptionsModel := new(vmwarev1.GetSpecificWorkloadDomainInstanceOptions)
				getSpecificWorkloadDomainInstanceOptionsModel.SiteID = core.StringPtr("testString")
				getSpecificWorkloadDomainInstanceOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getSpecificWorkloadDomainInstanceOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getSpecificWorkloadDomainInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vmwareService.GetSpecificWorkloadDomainInstance(getSpecificWorkloadDomainInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetSpecificWorkloadDomainInstance with error: Operation validation and request error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the GetSpecificWorkloadDomainInstanceOptions model
				getSpecificWorkloadDomainInstanceOptionsModel := new(vmwarev1.GetSpecificWorkloadDomainInstanceOptions)
				getSpecificWorkloadDomainInstanceOptionsModel.SiteID = core.StringPtr("testString")
				getSpecificWorkloadDomainInstanceOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getSpecificWorkloadDomainInstanceOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getSpecificWorkloadDomainInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vmwareService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vmwareService.GetSpecificWorkloadDomainInstance(getSpecificWorkloadDomainInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetSpecificWorkloadDomainInstanceOptions model with no property values
				getSpecificWorkloadDomainInstanceOptionsModelNew := new(vmwarev1.GetSpecificWorkloadDomainInstanceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = vmwareService.GetSpecificWorkloadDomainInstance(getSpecificWorkloadDomainInstanceOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetSpecificWorkloadDomainInstance successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the GetSpecificWorkloadDomainInstanceOptions model
				getSpecificWorkloadDomainInstanceOptionsModel := new(vmwarev1.GetSpecificWorkloadDomainInstanceOptions)
				getSpecificWorkloadDomainInstanceOptionsModel.SiteID = core.StringPtr("testString")
				getSpecificWorkloadDomainInstanceOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getSpecificWorkloadDomainInstanceOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getSpecificWorkloadDomainInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vmwareService.GetSpecificWorkloadDomainInstance(getSpecificWorkloadDomainInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteWorkloadDomain(deleteWorkloadDomainOptions *DeleteWorkloadDomainOptions) - Operation response error`, func() {
		deleteWorkloadDomainPath := "/director_sites/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteWorkloadDomainPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteWorkloadDomain with error: Operation response processing error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the DeleteWorkloadDomainOptions model
				deleteWorkloadDomainOptionsModel := new(vmwarev1.DeleteWorkloadDomainOptions)
				deleteWorkloadDomainOptionsModel.SiteID = core.StringPtr("testString")
				deleteWorkloadDomainOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteWorkloadDomainOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				deleteWorkloadDomainOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vmwareService.DeleteWorkloadDomain(deleteWorkloadDomainOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vmwareService.EnableRetries(0, 0)
				result, response, operationErr = vmwareService.DeleteWorkloadDomain(deleteWorkloadDomainOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteWorkloadDomain(deleteWorkloadDomainOptions *DeleteWorkloadDomainOptions)`, func() {
		deleteWorkloadDomainPath := "/director_sites/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteWorkloadDomainPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"id": "ID", "instance_ordered": "InstanceOrdered", "instance_created": "InstanceCreated", "name": "Name", "status": "Creating", "resource_group": "ResourceGroup", "requester": "Requester", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCrn", "clusters": [{"id": "ID", "name": "Name", "location": "Location", "host_count": 9, "status": "Status", "cluster_name": "ClusterName", "host_profile": "HostProfile", "file_shares": {"anyKey": "anyValue"}}]}`)
				}))
			})
			It(`Invoke DeleteWorkloadDomain successfully with retries`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())
				vmwareService.EnableRetries(0, 0)

				// Construct an instance of the DeleteWorkloadDomainOptions model
				deleteWorkloadDomainOptionsModel := new(vmwarev1.DeleteWorkloadDomainOptions)
				deleteWorkloadDomainOptionsModel.SiteID = core.StringPtr("testString")
				deleteWorkloadDomainOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteWorkloadDomainOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				deleteWorkloadDomainOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vmwareService.DeleteWorkloadDomainWithContext(ctx, deleteWorkloadDomainOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vmwareService.DisableRetries()
				result, response, operationErr := vmwareService.DeleteWorkloadDomain(deleteWorkloadDomainOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vmwareService.DeleteWorkloadDomainWithContext(ctx, deleteWorkloadDomainOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteWorkloadDomainPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"id": "ID", "instance_ordered": "InstanceOrdered", "instance_created": "InstanceCreated", "name": "Name", "status": "Creating", "resource_group": "ResourceGroup", "requester": "Requester", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCrn", "clusters": [{"id": "ID", "name": "Name", "location": "Location", "host_count": 9, "status": "Status", "cluster_name": "ClusterName", "host_profile": "HostProfile", "file_shares": {"anyKey": "anyValue"}}]}`)
				}))
			})
			It(`Invoke DeleteWorkloadDomain successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vmwareService.DeleteWorkloadDomain(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteWorkloadDomainOptions model
				deleteWorkloadDomainOptionsModel := new(vmwarev1.DeleteWorkloadDomainOptions)
				deleteWorkloadDomainOptionsModel.SiteID = core.StringPtr("testString")
				deleteWorkloadDomainOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteWorkloadDomainOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				deleteWorkloadDomainOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vmwareService.DeleteWorkloadDomain(deleteWorkloadDomainOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteWorkloadDomain with error: Operation validation and request error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the DeleteWorkloadDomainOptions model
				deleteWorkloadDomainOptionsModel := new(vmwarev1.DeleteWorkloadDomainOptions)
				deleteWorkloadDomainOptionsModel.SiteID = core.StringPtr("testString")
				deleteWorkloadDomainOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteWorkloadDomainOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				deleteWorkloadDomainOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vmwareService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vmwareService.DeleteWorkloadDomain(deleteWorkloadDomainOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteWorkloadDomainOptions model with no property values
				deleteWorkloadDomainOptionsModelNew := new(vmwarev1.DeleteWorkloadDomainOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = vmwareService.DeleteWorkloadDomain(deleteWorkloadDomainOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(202)
				}))
			})
			It(`Invoke DeleteWorkloadDomain successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the DeleteWorkloadDomainOptions model
				deleteWorkloadDomainOptionsModel := new(vmwarev1.DeleteWorkloadDomainOptions)
				deleteWorkloadDomainOptionsModel.SiteID = core.StringPtr("testString")
				deleteWorkloadDomainOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteWorkloadDomainOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				deleteWorkloadDomainOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vmwareService.DeleteWorkloadDomain(deleteWorkloadDomainOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListClusterInstances(listClusterInstancesOptions *ListClusterInstancesOptions) - Operation response error`, func() {
		listClusterInstancesPath := "/director_sites/testString/clusters"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listClusterInstancesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListClusterInstances with error: Operation response processing error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the ListClusterInstancesOptions model
				listClusterInstancesOptionsModel := new(vmwarev1.ListClusterInstancesOptions)
				listClusterInstancesOptionsModel.SiteID = core.StringPtr("testString")
				listClusterInstancesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listClusterInstancesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listClusterInstancesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vmwareService.ListClusterInstances(listClusterInstancesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vmwareService.EnableRetries(0, 0)
				result, response, operationErr = vmwareService.ListClusterInstances(listClusterInstancesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListClusterInstances(listClusterInstancesOptions *ListClusterInstancesOptions)`, func() {
		listClusterInstancesPath := "/director_sites/testString/clusters"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listClusterInstancesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"clusters": [{"id": "ID", "name": "Name", "instance_ordered": "InstanceOrdered", "instance_created": "InstanceCreated", "instance_deleted": "InstanceDeleted", "location": "Location", "host_count": 9, "status": "Status", "site_id": "SiteID", "host_profile": "HostProfile", "storage_type": "nfs", "billing_plan": "monthly", "file_shares": {"anyKey": "anyValue"}}]}`)
				}))
			})
			It(`Invoke ListClusterInstances successfully with retries`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())
				vmwareService.EnableRetries(0, 0)

				// Construct an instance of the ListClusterInstancesOptions model
				listClusterInstancesOptionsModel := new(vmwarev1.ListClusterInstancesOptions)
				listClusterInstancesOptionsModel.SiteID = core.StringPtr("testString")
				listClusterInstancesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listClusterInstancesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listClusterInstancesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vmwareService.ListClusterInstancesWithContext(ctx, listClusterInstancesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vmwareService.DisableRetries()
				result, response, operationErr := vmwareService.ListClusterInstances(listClusterInstancesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vmwareService.ListClusterInstancesWithContext(ctx, listClusterInstancesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listClusterInstancesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"clusters": [{"id": "ID", "name": "Name", "instance_ordered": "InstanceOrdered", "instance_created": "InstanceCreated", "instance_deleted": "InstanceDeleted", "location": "Location", "host_count": 9, "status": "Status", "site_id": "SiteID", "host_profile": "HostProfile", "storage_type": "nfs", "billing_plan": "monthly", "file_shares": {"anyKey": "anyValue"}}]}`)
				}))
			})
			It(`Invoke ListClusterInstances successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vmwareService.ListClusterInstances(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListClusterInstancesOptions model
				listClusterInstancesOptionsModel := new(vmwarev1.ListClusterInstancesOptions)
				listClusterInstancesOptionsModel.SiteID = core.StringPtr("testString")
				listClusterInstancesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listClusterInstancesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listClusterInstancesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vmwareService.ListClusterInstances(listClusterInstancesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListClusterInstances with error: Operation validation and request error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the ListClusterInstancesOptions model
				listClusterInstancesOptionsModel := new(vmwarev1.ListClusterInstancesOptions)
				listClusterInstancesOptionsModel.SiteID = core.StringPtr("testString")
				listClusterInstancesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listClusterInstancesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listClusterInstancesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vmwareService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vmwareService.ListClusterInstances(listClusterInstancesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListClusterInstancesOptions model with no property values
				listClusterInstancesOptionsModelNew := new(vmwarev1.ListClusterInstancesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = vmwareService.ListClusterInstances(listClusterInstancesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListClusterInstances successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the ListClusterInstancesOptions model
				listClusterInstancesOptionsModel := new(vmwarev1.ListClusterInstancesOptions)
				listClusterInstancesOptionsModel.SiteID = core.StringPtr("testString")
				listClusterInstancesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listClusterInstancesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listClusterInstancesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vmwareService.ListClusterInstances(listClusterInstancesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSpecificClusterInstance(getSpecificClusterInstanceOptions *GetSpecificClusterInstanceOptions) - Operation response error`, func() {
		getSpecificClusterInstancePath := "/director_sites/testString/clusters/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSpecificClusterInstancePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetSpecificClusterInstance with error: Operation response processing error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the GetSpecificClusterInstanceOptions model
				getSpecificClusterInstanceOptionsModel := new(vmwarev1.GetSpecificClusterInstanceOptions)
				getSpecificClusterInstanceOptionsModel.SiteID = core.StringPtr("testString")
				getSpecificClusterInstanceOptionsModel.ClusterID = core.StringPtr("testString")
				getSpecificClusterInstanceOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getSpecificClusterInstanceOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getSpecificClusterInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vmwareService.GetSpecificClusterInstance(getSpecificClusterInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vmwareService.EnableRetries(0, 0)
				result, response, operationErr = vmwareService.GetSpecificClusterInstance(getSpecificClusterInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSpecificClusterInstance(getSpecificClusterInstanceOptions *GetSpecificClusterInstanceOptions)`, func() {
		getSpecificClusterInstancePath := "/director_sites/testString/clusters/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSpecificClusterInstancePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "instance_ordered": "InstanceOrdered", "instance_created": "InstanceCreated", "instance_deleted": "InstanceDeleted", "location": "Location", "host_count": 9, "status": "Status", "site_id": "SiteID", "host_profile": "HostProfile", "storage_type": "nfs", "billing_plan": "monthly", "file_shares": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke GetSpecificClusterInstance successfully with retries`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())
				vmwareService.EnableRetries(0, 0)

				// Construct an instance of the GetSpecificClusterInstanceOptions model
				getSpecificClusterInstanceOptionsModel := new(vmwarev1.GetSpecificClusterInstanceOptions)
				getSpecificClusterInstanceOptionsModel.SiteID = core.StringPtr("testString")
				getSpecificClusterInstanceOptionsModel.ClusterID = core.StringPtr("testString")
				getSpecificClusterInstanceOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getSpecificClusterInstanceOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getSpecificClusterInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vmwareService.GetSpecificClusterInstanceWithContext(ctx, getSpecificClusterInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vmwareService.DisableRetries()
				result, response, operationErr := vmwareService.GetSpecificClusterInstance(getSpecificClusterInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vmwareService.GetSpecificClusterInstanceWithContext(ctx, getSpecificClusterInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSpecificClusterInstancePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "instance_ordered": "InstanceOrdered", "instance_created": "InstanceCreated", "instance_deleted": "InstanceDeleted", "location": "Location", "host_count": 9, "status": "Status", "site_id": "SiteID", "host_profile": "HostProfile", "storage_type": "nfs", "billing_plan": "monthly", "file_shares": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke GetSpecificClusterInstance successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vmwareService.GetSpecificClusterInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSpecificClusterInstanceOptions model
				getSpecificClusterInstanceOptionsModel := new(vmwarev1.GetSpecificClusterInstanceOptions)
				getSpecificClusterInstanceOptionsModel.SiteID = core.StringPtr("testString")
				getSpecificClusterInstanceOptionsModel.ClusterID = core.StringPtr("testString")
				getSpecificClusterInstanceOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getSpecificClusterInstanceOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getSpecificClusterInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vmwareService.GetSpecificClusterInstance(getSpecificClusterInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetSpecificClusterInstance with error: Operation validation and request error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the GetSpecificClusterInstanceOptions model
				getSpecificClusterInstanceOptionsModel := new(vmwarev1.GetSpecificClusterInstanceOptions)
				getSpecificClusterInstanceOptionsModel.SiteID = core.StringPtr("testString")
				getSpecificClusterInstanceOptionsModel.ClusterID = core.StringPtr("testString")
				getSpecificClusterInstanceOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getSpecificClusterInstanceOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getSpecificClusterInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vmwareService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vmwareService.GetSpecificClusterInstance(getSpecificClusterInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetSpecificClusterInstanceOptions model with no property values
				getSpecificClusterInstanceOptionsModelNew := new(vmwarev1.GetSpecificClusterInstanceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = vmwareService.GetSpecificClusterInstance(getSpecificClusterInstanceOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetSpecificClusterInstance successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the GetSpecificClusterInstanceOptions model
				getSpecificClusterInstanceOptionsModel := new(vmwarev1.GetSpecificClusterInstanceOptions)
				getSpecificClusterInstanceOptionsModel.SiteID = core.StringPtr("testString")
				getSpecificClusterInstanceOptionsModel.ClusterID = core.StringPtr("testString")
				getSpecificClusterInstanceOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getSpecificClusterInstanceOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getSpecificClusterInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vmwareService.GetSpecificClusterInstance(getSpecificClusterInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SetHostsCount(setHostsCountOptions *SetHostsCountOptions) - Operation response error`, func() {
		setHostsCountPath := "/director_sites/testString/clusters/testString/hosts_count"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setHostsCountPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SetHostsCount with error: Operation response processing error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the SetHostsCountOptions model
				setHostsCountOptionsModel := new(vmwarev1.SetHostsCountOptions)
				setHostsCountOptionsModel.SiteID = core.StringPtr("testString")
				setHostsCountOptionsModel.ClusterID = core.StringPtr("testString")
				setHostsCountOptionsModel.Count = core.Int64Ptr(int64(2))
				setHostsCountOptionsModel.AcceptLanguage = core.StringPtr("testString")
				setHostsCountOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				setHostsCountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vmwareService.SetHostsCount(setHostsCountOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vmwareService.EnableRetries(0, 0)
				result, response, operationErr = vmwareService.SetHostsCount(setHostsCountOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SetHostsCount(setHostsCountOptions *SetHostsCountOptions)`, func() {
		setHostsCountPath := "/director_sites/testString/clusters/testString/hosts_count"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setHostsCountPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"message": "The request has been accepted."}`)
				}))
			})
			It(`Invoke SetHostsCount successfully with retries`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())
				vmwareService.EnableRetries(0, 0)

				// Construct an instance of the SetHostsCountOptions model
				setHostsCountOptionsModel := new(vmwarev1.SetHostsCountOptions)
				setHostsCountOptionsModel.SiteID = core.StringPtr("testString")
				setHostsCountOptionsModel.ClusterID = core.StringPtr("testString")
				setHostsCountOptionsModel.Count = core.Int64Ptr(int64(2))
				setHostsCountOptionsModel.AcceptLanguage = core.StringPtr("testString")
				setHostsCountOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				setHostsCountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vmwareService.SetHostsCountWithContext(ctx, setHostsCountOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vmwareService.DisableRetries()
				result, response, operationErr := vmwareService.SetHostsCount(setHostsCountOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vmwareService.SetHostsCountWithContext(ctx, setHostsCountOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setHostsCountPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"message": "The request has been accepted."}`)
				}))
			})
			It(`Invoke SetHostsCount successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vmwareService.SetHostsCount(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SetHostsCountOptions model
				setHostsCountOptionsModel := new(vmwarev1.SetHostsCountOptions)
				setHostsCountOptionsModel.SiteID = core.StringPtr("testString")
				setHostsCountOptionsModel.ClusterID = core.StringPtr("testString")
				setHostsCountOptionsModel.Count = core.Int64Ptr(int64(2))
				setHostsCountOptionsModel.AcceptLanguage = core.StringPtr("testString")
				setHostsCountOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				setHostsCountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vmwareService.SetHostsCount(setHostsCountOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke SetHostsCount with error: Operation validation and request error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the SetHostsCountOptions model
				setHostsCountOptionsModel := new(vmwarev1.SetHostsCountOptions)
				setHostsCountOptionsModel.SiteID = core.StringPtr("testString")
				setHostsCountOptionsModel.ClusterID = core.StringPtr("testString")
				setHostsCountOptionsModel.Count = core.Int64Ptr(int64(2))
				setHostsCountOptionsModel.AcceptLanguage = core.StringPtr("testString")
				setHostsCountOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				setHostsCountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vmwareService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vmwareService.SetHostsCount(setHostsCountOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the SetHostsCountOptions model with no property values
				setHostsCountOptionsModelNew := new(vmwarev1.SetHostsCountOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = vmwareService.SetHostsCount(setHostsCountOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke SetHostsCount successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the SetHostsCountOptions model
				setHostsCountOptionsModel := new(vmwarev1.SetHostsCountOptions)
				setHostsCountOptionsModel.SiteID = core.StringPtr("testString")
				setHostsCountOptionsModel.ClusterID = core.StringPtr("testString")
				setHostsCountOptionsModel.Count = core.Int64Ptr(int64(2))
				setHostsCountOptionsModel.AcceptLanguage = core.StringPtr("testString")
				setHostsCountOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				setHostsCountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vmwareService.SetHostsCount(setHostsCountOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SetFileShares(setFileSharesOptions *SetFileSharesOptions) - Operation response error`, func() {
		setFileSharesPath := "/director_sites/testString/clusters/testString/file_shares"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setFileSharesPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SetFileShares with error: Operation response processing error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the SetFileSharesOptions model
				setFileSharesOptionsModel := new(vmwarev1.SetFileSharesOptions)
				setFileSharesOptionsModel.SiteID = core.StringPtr("testString")
				setFileSharesOptionsModel.ClusterID = core.StringPtr("testString")
				setFileSharesOptionsModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				setFileSharesOptionsModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				setFileSharesOptionsModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				setFileSharesOptionsModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))
				setFileSharesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				setFileSharesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				setFileSharesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vmwareService.SetFileShares(setFileSharesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vmwareService.EnableRetries(0, 0)
				result, response, operationErr = vmwareService.SetFileShares(setFileSharesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SetFileShares(setFileSharesOptions *SetFileSharesOptions)`, func() {
		setFileSharesPath := "/director_sites/testString/clusters/testString/file_shares"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setFileSharesPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"STORAGE_POINT_TWO_FIVE_IOPS_GB": 0, "STORAGE_TWO_IOPS_GB": 0, "STORAGE_FOUR_IOPS_GB": 0, "STORAGE_TEN_IOPS_GB": 0}`)
				}))
			})
			It(`Invoke SetFileShares successfully with retries`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())
				vmwareService.EnableRetries(0, 0)

				// Construct an instance of the SetFileSharesOptions model
				setFileSharesOptionsModel := new(vmwarev1.SetFileSharesOptions)
				setFileSharesOptionsModel.SiteID = core.StringPtr("testString")
				setFileSharesOptionsModel.ClusterID = core.StringPtr("testString")
				setFileSharesOptionsModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				setFileSharesOptionsModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				setFileSharesOptionsModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				setFileSharesOptionsModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))
				setFileSharesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				setFileSharesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				setFileSharesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vmwareService.SetFileSharesWithContext(ctx, setFileSharesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vmwareService.DisableRetries()
				result, response, operationErr := vmwareService.SetFileShares(setFileSharesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vmwareService.SetFileSharesWithContext(ctx, setFileSharesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setFileSharesPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"STORAGE_POINT_TWO_FIVE_IOPS_GB": 0, "STORAGE_TWO_IOPS_GB": 0, "STORAGE_FOUR_IOPS_GB": 0, "STORAGE_TEN_IOPS_GB": 0}`)
				}))
			})
			It(`Invoke SetFileShares successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vmwareService.SetFileShares(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SetFileSharesOptions model
				setFileSharesOptionsModel := new(vmwarev1.SetFileSharesOptions)
				setFileSharesOptionsModel.SiteID = core.StringPtr("testString")
				setFileSharesOptionsModel.ClusterID = core.StringPtr("testString")
				setFileSharesOptionsModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				setFileSharesOptionsModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				setFileSharesOptionsModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				setFileSharesOptionsModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))
				setFileSharesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				setFileSharesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				setFileSharesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vmwareService.SetFileShares(setFileSharesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke SetFileShares with error: Operation validation and request error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the SetFileSharesOptions model
				setFileSharesOptionsModel := new(vmwarev1.SetFileSharesOptions)
				setFileSharesOptionsModel.SiteID = core.StringPtr("testString")
				setFileSharesOptionsModel.ClusterID = core.StringPtr("testString")
				setFileSharesOptionsModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				setFileSharesOptionsModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				setFileSharesOptionsModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				setFileSharesOptionsModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))
				setFileSharesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				setFileSharesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				setFileSharesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vmwareService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vmwareService.SetFileShares(setFileSharesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the SetFileSharesOptions model with no property values
				setFileSharesOptionsModelNew := new(vmwarev1.SetFileSharesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = vmwareService.SetFileShares(setFileSharesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke SetFileShares successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the SetFileSharesOptions model
				setFileSharesOptionsModel := new(vmwarev1.SetFileSharesOptions)
				setFileSharesOptionsModel.SiteID = core.StringPtr("testString")
				setFileSharesOptionsModel.ClusterID = core.StringPtr("testString")
				setFileSharesOptionsModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				setFileSharesOptionsModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				setFileSharesOptionsModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				setFileSharesOptionsModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))
				setFileSharesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				setFileSharesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				setFileSharesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vmwareService.SetFileShares(setFileSharesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetRegions(getRegionsOptions *GetRegionsOptions) - Operation response error`, func() {
		getRegionsPath := "/director_site_regions"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getRegionsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetRegions with error: Operation response processing error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the GetRegionsOptions model
				getRegionsOptionsModel := new(vmwarev1.GetRegionsOptions)
				getRegionsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getRegionsOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getRegionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vmwareService.GetRegions(getRegionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vmwareService.EnableRetries(0, 0)
				result, response, operationErr = vmwareService.GetRegions(getRegionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetRegions(getRegionsOptions *GetRegionsOptions)`, func() {
		getRegionsPath := "/director_site_regions"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getRegionsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"director_site_regions": {"mapKey": {"datacenters": [{"display_name": "DisplayName", "name": "Name", "uplink_speed": "UplinkSpeed"}], "endpoint": "Endpoint"}}}`)
				}))
			})
			It(`Invoke GetRegions successfully with retries`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())
				vmwareService.EnableRetries(0, 0)

				// Construct an instance of the GetRegionsOptions model
				getRegionsOptionsModel := new(vmwarev1.GetRegionsOptions)
				getRegionsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getRegionsOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getRegionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vmwareService.GetRegionsWithContext(ctx, getRegionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vmwareService.DisableRetries()
				result, response, operationErr := vmwareService.GetRegions(getRegionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vmwareService.GetRegionsWithContext(ctx, getRegionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getRegionsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"director_site_regions": {"mapKey": {"datacenters": [{"display_name": "DisplayName", "name": "Name", "uplink_speed": "UplinkSpeed"}], "endpoint": "Endpoint"}}}`)
				}))
			})
			It(`Invoke GetRegions successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vmwareService.GetRegions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetRegionsOptions model
				getRegionsOptionsModel := new(vmwarev1.GetRegionsOptions)
				getRegionsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getRegionsOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getRegionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vmwareService.GetRegions(getRegionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetRegions with error: Operation request error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the GetRegionsOptions model
				getRegionsOptionsModel := new(vmwarev1.GetRegionsOptions)
				getRegionsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getRegionsOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getRegionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vmwareService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vmwareService.GetRegions(getRegionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetRegions successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the GetRegionsOptions model
				getRegionsOptionsModel := new(vmwarev1.GetRegionsOptions)
				getRegionsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getRegionsOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getRegionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vmwareService.GetRegions(getRegionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ViewInstance(viewInstanceOptions *ViewInstanceOptions) - Operation response error`, func() {
		viewInstancePath := "/director_site_host_profiles"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(viewInstancePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ViewInstance with error: Operation response processing error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the ViewInstanceOptions model
				viewInstanceOptionsModel := new(vmwarev1.ViewInstanceOptions)
				viewInstanceOptionsModel.AcceptLanguage = core.StringPtr("testString")
				viewInstanceOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				viewInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vmwareService.ViewInstance(viewInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vmwareService.EnableRetries(0, 0)
				result, response, operationErr = vmwareService.ViewInstance(viewInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ViewInstance(viewInstanceOptions *ViewInstanceOptions)`, func() {
		viewInstancePath := "/director_site_host_profiles"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(viewInstancePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"director_site_host_profiles": [{"profile_name": "ProfileName", "cpu_type": "CpuType", "cpu_count": 8, "ram": 3, "local_disks": [{"quantity": 8, "size": 4, "type": "Type"}]}]}`)
				}))
			})
			It(`Invoke ViewInstance successfully with retries`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())
				vmwareService.EnableRetries(0, 0)

				// Construct an instance of the ViewInstanceOptions model
				viewInstanceOptionsModel := new(vmwarev1.ViewInstanceOptions)
				viewInstanceOptionsModel.AcceptLanguage = core.StringPtr("testString")
				viewInstanceOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				viewInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vmwareService.ViewInstanceWithContext(ctx, viewInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vmwareService.DisableRetries()
				result, response, operationErr := vmwareService.ViewInstance(viewInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vmwareService.ViewInstanceWithContext(ctx, viewInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(viewInstancePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"director_site_host_profiles": [{"profile_name": "ProfileName", "cpu_type": "CpuType", "cpu_count": 8, "ram": 3, "local_disks": [{"quantity": 8, "size": 4, "type": "Type"}]}]}`)
				}))
			})
			It(`Invoke ViewInstance successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vmwareService.ViewInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ViewInstanceOptions model
				viewInstanceOptionsModel := new(vmwarev1.ViewInstanceOptions)
				viewInstanceOptionsModel.AcceptLanguage = core.StringPtr("testString")
				viewInstanceOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				viewInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vmwareService.ViewInstance(viewInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ViewInstance with error: Operation request error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the ViewInstanceOptions model
				viewInstanceOptionsModel := new(vmwarev1.ViewInstanceOptions)
				viewInstanceOptionsModel.AcceptLanguage = core.StringPtr("testString")
				viewInstanceOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				viewInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vmwareService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vmwareService.ViewInstance(viewInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ViewInstance successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the ViewInstanceOptions model
				viewInstanceOptionsModel := new(vmwarev1.ViewInstanceOptions)
				viewInstanceOptionsModel.AcceptLanguage = core.StringPtr("testString")
				viewInstanceOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				viewInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vmwareService.ViewInstance(viewInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceOrgAdminPassword(replaceOrgAdminPasswordOptions *ReplaceOrgAdminPasswordOptions) - Operation response error`, func() {
		replaceOrgAdminPasswordPath := "/director_site_password"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceOrgAdminPasswordPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.URL.Query()["site_id"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceOrgAdminPassword with error: Operation response processing error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the ReplaceOrgAdminPasswordOptions model
				replaceOrgAdminPasswordOptionsModel := new(vmwarev1.ReplaceOrgAdminPasswordOptions)
				replaceOrgAdminPasswordOptionsModel.SiteID = core.StringPtr("testString")
				replaceOrgAdminPasswordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vmwareService.ReplaceOrgAdminPassword(replaceOrgAdminPasswordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vmwareService.EnableRetries(0, 0)
				result, response, operationErr = vmwareService.ReplaceOrgAdminPassword(replaceOrgAdminPasswordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceOrgAdminPassword(replaceOrgAdminPasswordOptions *ReplaceOrgAdminPasswordOptions)`, func() {
		replaceOrgAdminPasswordPath := "/director_site_password"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceOrgAdminPasswordPath))
					Expect(req.Method).To(Equal("PUT"))

					Expect(req.URL.Query()["site_id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"password": "Password"}`)
				}))
			})
			It(`Invoke ReplaceOrgAdminPassword successfully with retries`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())
				vmwareService.EnableRetries(0, 0)

				// Construct an instance of the ReplaceOrgAdminPasswordOptions model
				replaceOrgAdminPasswordOptionsModel := new(vmwarev1.ReplaceOrgAdminPasswordOptions)
				replaceOrgAdminPasswordOptionsModel.SiteID = core.StringPtr("testString")
				replaceOrgAdminPasswordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vmwareService.ReplaceOrgAdminPasswordWithContext(ctx, replaceOrgAdminPasswordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vmwareService.DisableRetries()
				result, response, operationErr := vmwareService.ReplaceOrgAdminPassword(replaceOrgAdminPasswordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vmwareService.ReplaceOrgAdminPasswordWithContext(ctx, replaceOrgAdminPasswordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceOrgAdminPasswordPath))
					Expect(req.Method).To(Equal("PUT"))

					Expect(req.URL.Query()["site_id"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"password": "Password"}`)
				}))
			})
			It(`Invoke ReplaceOrgAdminPassword successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vmwareService.ReplaceOrgAdminPassword(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ReplaceOrgAdminPasswordOptions model
				replaceOrgAdminPasswordOptionsModel := new(vmwarev1.ReplaceOrgAdminPasswordOptions)
				replaceOrgAdminPasswordOptionsModel.SiteID = core.StringPtr("testString")
				replaceOrgAdminPasswordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vmwareService.ReplaceOrgAdminPassword(replaceOrgAdminPasswordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceOrgAdminPassword with error: Operation validation and request error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the ReplaceOrgAdminPasswordOptions model
				replaceOrgAdminPasswordOptionsModel := new(vmwarev1.ReplaceOrgAdminPasswordOptions)
				replaceOrgAdminPasswordOptionsModel.SiteID = core.StringPtr("testString")
				replaceOrgAdminPasswordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vmwareService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vmwareService.ReplaceOrgAdminPassword(replaceOrgAdminPasswordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceOrgAdminPasswordOptions model with no property values
				replaceOrgAdminPasswordOptionsModelNew := new(vmwarev1.ReplaceOrgAdminPasswordOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = vmwareService.ReplaceOrgAdminPassword(replaceOrgAdminPasswordOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ReplaceOrgAdminPassword successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the ReplaceOrgAdminPasswordOptions model
				replaceOrgAdminPasswordOptionsModel := new(vmwarev1.ReplaceOrgAdminPasswordOptions)
				replaceOrgAdminPasswordOptionsModel.SiteID = core.StringPtr("testString")
				replaceOrgAdminPasswordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vmwareService.ReplaceOrgAdminPassword(replaceOrgAdminPasswordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListPrices(listPricesOptions *ListPricesOptions) - Operation response error`, func() {
		listPricesPath := "/director_site_pricing"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listPricesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListPrices with error: Operation response processing error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the ListPricesOptions model
				listPricesOptionsModel := new(vmwarev1.ListPricesOptions)
				listPricesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listPricesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listPricesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vmwareService.ListPrices(listPricesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vmwareService.EnableRetries(0, 0)
				result, response, operationErr = vmwareService.ListPrices(listPricesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListPrices(listPricesOptions *ListPricesOptions)`, func() {
		listPricesPath := "/director_site_pricing"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listPricesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"director_site_pricing": [{"metric": "Metric", "description": "Description", "price_list": [{"country": "Country", "currency": "Currency", "prices": [{"price": 5, "quantity_tier": 12}]}]}]}`)
				}))
			})
			It(`Invoke ListPrices successfully with retries`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())
				vmwareService.EnableRetries(0, 0)

				// Construct an instance of the ListPricesOptions model
				listPricesOptionsModel := new(vmwarev1.ListPricesOptions)
				listPricesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listPricesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listPricesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vmwareService.ListPricesWithContext(ctx, listPricesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vmwareService.DisableRetries()
				result, response, operationErr := vmwareService.ListPrices(listPricesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vmwareService.ListPricesWithContext(ctx, listPricesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listPricesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"director_site_pricing": [{"metric": "Metric", "description": "Description", "price_list": [{"country": "Country", "currency": "Currency", "prices": [{"price": 5, "quantity_tier": 12}]}]}]}`)
				}))
			})
			It(`Invoke ListPrices successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vmwareService.ListPrices(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListPricesOptions model
				listPricesOptionsModel := new(vmwarev1.ListPricesOptions)
				listPricesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listPricesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listPricesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vmwareService.ListPrices(listPricesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListPrices with error: Operation request error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the ListPricesOptions model
				listPricesOptionsModel := new(vmwarev1.ListPricesOptions)
				listPricesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listPricesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listPricesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vmwareService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vmwareService.ListPrices(listPricesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListPrices successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the ListPricesOptions model
				listPricesOptionsModel := new(vmwarev1.ListPricesOptions)
				listPricesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listPricesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listPricesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vmwareService.ListPrices(listPricesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetVcddPrice(getVcddPriceOptions *GetVcddPriceOptions) - Operation response error`, func() {
		getVcddPricePath := "/director_site_price_quote"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getVcddPricePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetVcddPrice with error: Operation response processing error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the FileShares model
				fileSharesModel := new(vmwarev1.FileShares)
				fileSharesModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))

				// Construct an instance of the DirectorSitePriceQuoteClusterInfo model
				directorSitePriceQuoteClusterInfoModel := new(vmwarev1.DirectorSitePriceQuoteClusterInfo)
				directorSitePriceQuoteClusterInfoModel.Name = core.StringPtr("testString")
				directorSitePriceQuoteClusterInfoModel.HostProfile = core.StringPtr("testString")
				directorSitePriceQuoteClusterInfoModel.HostCount = core.Int64Ptr(int64(2))
				directorSitePriceQuoteClusterInfoModel.FileShares = fileSharesModel

				// Construct an instance of the GetVcddPriceOptions model
				getVcddPriceOptionsModel := new(vmwarev1.GetVcddPriceOptions)
				getVcddPriceOptionsModel.Country = core.StringPtr("USA")
				getVcddPriceOptionsModel.Clusters = []vmwarev1.DirectorSitePriceQuoteClusterInfo{*directorSitePriceQuoteClusterInfoModel}
				getVcddPriceOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getVcddPriceOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getVcddPriceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vmwareService.GetVcddPrice(getVcddPriceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vmwareService.EnableRetries(0, 0)
				result, response, operationErr = vmwareService.GetVcddPrice(getVcddPriceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetVcddPrice(getVcddPriceOptions *GetVcddPriceOptions)`, func() {
		getVcddPricePath := "/director_site_price_quote"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getVcddPricePath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"base_charge": {"name": "Name", "currency": "Currency", "price": 5}, "clusters": [{"name": "Name", "currency": "Currency", "price": 5, "items": [{"name": "Name", "currency": "Currency", "price": 5, "items": [{"name": "Name", "count": 5, "currency": "Currency", "price": 5}]}]}], "currency": "Currency", "total": 5}`)
				}))
			})
			It(`Invoke GetVcddPrice successfully with retries`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())
				vmwareService.EnableRetries(0, 0)

				// Construct an instance of the FileShares model
				fileSharesModel := new(vmwarev1.FileShares)
				fileSharesModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))

				// Construct an instance of the DirectorSitePriceQuoteClusterInfo model
				directorSitePriceQuoteClusterInfoModel := new(vmwarev1.DirectorSitePriceQuoteClusterInfo)
				directorSitePriceQuoteClusterInfoModel.Name = core.StringPtr("testString")
				directorSitePriceQuoteClusterInfoModel.HostProfile = core.StringPtr("testString")
				directorSitePriceQuoteClusterInfoModel.HostCount = core.Int64Ptr(int64(2))
				directorSitePriceQuoteClusterInfoModel.FileShares = fileSharesModel

				// Construct an instance of the GetVcddPriceOptions model
				getVcddPriceOptionsModel := new(vmwarev1.GetVcddPriceOptions)
				getVcddPriceOptionsModel.Country = core.StringPtr("USA")
				getVcddPriceOptionsModel.Clusters = []vmwarev1.DirectorSitePriceQuoteClusterInfo{*directorSitePriceQuoteClusterInfoModel}
				getVcddPriceOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getVcddPriceOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getVcddPriceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vmwareService.GetVcddPriceWithContext(ctx, getVcddPriceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vmwareService.DisableRetries()
				result, response, operationErr := vmwareService.GetVcddPrice(getVcddPriceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vmwareService.GetVcddPriceWithContext(ctx, getVcddPriceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getVcddPricePath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"base_charge": {"name": "Name", "currency": "Currency", "price": 5}, "clusters": [{"name": "Name", "currency": "Currency", "price": 5, "items": [{"name": "Name", "currency": "Currency", "price": 5, "items": [{"name": "Name", "count": 5, "currency": "Currency", "price": 5}]}]}], "currency": "Currency", "total": 5}`)
				}))
			})
			It(`Invoke GetVcddPrice successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vmwareService.GetVcddPrice(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the FileShares model
				fileSharesModel := new(vmwarev1.FileShares)
				fileSharesModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))

				// Construct an instance of the DirectorSitePriceQuoteClusterInfo model
				directorSitePriceQuoteClusterInfoModel := new(vmwarev1.DirectorSitePriceQuoteClusterInfo)
				directorSitePriceQuoteClusterInfoModel.Name = core.StringPtr("testString")
				directorSitePriceQuoteClusterInfoModel.HostProfile = core.StringPtr("testString")
				directorSitePriceQuoteClusterInfoModel.HostCount = core.Int64Ptr(int64(2))
				directorSitePriceQuoteClusterInfoModel.FileShares = fileSharesModel

				// Construct an instance of the GetVcddPriceOptions model
				getVcddPriceOptionsModel := new(vmwarev1.GetVcddPriceOptions)
				getVcddPriceOptionsModel.Country = core.StringPtr("USA")
				getVcddPriceOptionsModel.Clusters = []vmwarev1.DirectorSitePriceQuoteClusterInfo{*directorSitePriceQuoteClusterInfoModel}
				getVcddPriceOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getVcddPriceOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getVcddPriceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vmwareService.GetVcddPrice(getVcddPriceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetVcddPrice with error: Operation request error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the FileShares model
				fileSharesModel := new(vmwarev1.FileShares)
				fileSharesModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))

				// Construct an instance of the DirectorSitePriceQuoteClusterInfo model
				directorSitePriceQuoteClusterInfoModel := new(vmwarev1.DirectorSitePriceQuoteClusterInfo)
				directorSitePriceQuoteClusterInfoModel.Name = core.StringPtr("testString")
				directorSitePriceQuoteClusterInfoModel.HostProfile = core.StringPtr("testString")
				directorSitePriceQuoteClusterInfoModel.HostCount = core.Int64Ptr(int64(2))
				directorSitePriceQuoteClusterInfoModel.FileShares = fileSharesModel

				// Construct an instance of the GetVcddPriceOptions model
				getVcddPriceOptionsModel := new(vmwarev1.GetVcddPriceOptions)
				getVcddPriceOptionsModel.Country = core.StringPtr("USA")
				getVcddPriceOptionsModel.Clusters = []vmwarev1.DirectorSitePriceQuoteClusterInfo{*directorSitePriceQuoteClusterInfoModel}
				getVcddPriceOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getVcddPriceOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getVcddPriceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vmwareService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vmwareService.GetVcddPrice(getVcddPriceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke GetVcddPrice successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the FileShares model
				fileSharesModel := new(vmwarev1.FileShares)
				fileSharesModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))

				// Construct an instance of the DirectorSitePriceQuoteClusterInfo model
				directorSitePriceQuoteClusterInfoModel := new(vmwarev1.DirectorSitePriceQuoteClusterInfo)
				directorSitePriceQuoteClusterInfoModel.Name = core.StringPtr("testString")
				directorSitePriceQuoteClusterInfoModel.HostProfile = core.StringPtr("testString")
				directorSitePriceQuoteClusterInfoModel.HostCount = core.Int64Ptr(int64(2))
				directorSitePriceQuoteClusterInfoModel.FileShares = fileSharesModel

				// Construct an instance of the GetVcddPriceOptions model
				getVcddPriceOptionsModel := new(vmwarev1.GetVcddPriceOptions)
				getVcddPriceOptionsModel.Country = core.StringPtr("USA")
				getVcddPriceOptionsModel.Clusters = []vmwarev1.DirectorSitePriceQuoteClusterInfo{*directorSitePriceQuoteClusterInfoModel}
				getVcddPriceOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getVcddPriceOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getVcddPriceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vmwareService.GetVcddPrice(getVcddPriceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListVdcs(listVdcsOptions *ListVdcsOptions) - Operation response error`, func() {
		listVdcsPath := "/vdcs"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listVdcsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListVdcs with error: Operation response processing error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the ListVdcsOptions model
				listVdcsOptionsModel := new(vmwarev1.ListVdcsOptions)
				listVdcsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listVdcsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vmwareService.ListVdcs(listVdcsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vmwareService.EnableRetries(0, 0)
				result, response, operationErr = vmwareService.ListVdcs(listVdcsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListVdcs(listVdcsOptions *ListVdcsOptions)`, func() {
		listVdcsPath := "/vdcs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listVdcsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"vdcs": [{"id": "ID", "allocation_model": "paygo", "created_time": "2019-01-01T12:00:00.000Z", "crn": "Crn", "deleted_time": "2019-01-01T12:00:00.000Z", "director_site": {"id": "ID", "cluster": {"id": "ID"}, "url": "URL"}, "edges": [{"id": "ID", "public_ips": ["PublicIps"], "size": "medium", "type": "dedicated"}], "errors": [{"code": "Code", "message": "Message", "more_info": "MoreInfo"}], "name": "Name", "ordered_time": "2019-01-01T12:00:00.000Z", "org_name": "OrgName", "status": "Creating", "type": "dedicated"}]}`)
				}))
			})
			It(`Invoke ListVdcs successfully with retries`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())
				vmwareService.EnableRetries(0, 0)

				// Construct an instance of the ListVdcsOptions model
				listVdcsOptionsModel := new(vmwarev1.ListVdcsOptions)
				listVdcsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listVdcsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vmwareService.ListVdcsWithContext(ctx, listVdcsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vmwareService.DisableRetries()
				result, response, operationErr := vmwareService.ListVdcs(listVdcsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vmwareService.ListVdcsWithContext(ctx, listVdcsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listVdcsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"vdcs": [{"id": "ID", "allocation_model": "paygo", "created_time": "2019-01-01T12:00:00.000Z", "crn": "Crn", "deleted_time": "2019-01-01T12:00:00.000Z", "director_site": {"id": "ID", "cluster": {"id": "ID"}, "url": "URL"}, "edges": [{"id": "ID", "public_ips": ["PublicIps"], "size": "medium", "type": "dedicated"}], "errors": [{"code": "Code", "message": "Message", "more_info": "MoreInfo"}], "name": "Name", "ordered_time": "2019-01-01T12:00:00.000Z", "org_name": "OrgName", "status": "Creating", "type": "dedicated"}]}`)
				}))
			})
			It(`Invoke ListVdcs successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vmwareService.ListVdcs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListVdcsOptions model
				listVdcsOptionsModel := new(vmwarev1.ListVdcsOptions)
				listVdcsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listVdcsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vmwareService.ListVdcs(listVdcsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListVdcs with error: Operation request error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the ListVdcsOptions model
				listVdcsOptionsModel := new(vmwarev1.ListVdcsOptions)
				listVdcsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listVdcsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vmwareService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vmwareService.ListVdcs(listVdcsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListVdcs successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the ListVdcsOptions model
				listVdcsOptionsModel := new(vmwarev1.ListVdcsOptions)
				listVdcsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listVdcsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vmwareService.ListVdcs(listVdcsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateVdc(createVdcOptions *CreateVdcOptions) - Operation response error`, func() {
		createVdcPath := "/vdcs"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createVdcPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateVdc with error: Operation response processing error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the VDCDirectorSiteCluster model
				vdcDirectorSiteClusterModel := new(vmwarev1.VDCDirectorSiteCluster)
				vdcDirectorSiteClusterModel.ID = core.StringPtr("testString")

				// Construct an instance of the NewVDCDirectorSite model
				newVdcDirectorSiteModel := new(vmwarev1.NewVDCDirectorSite)
				newVdcDirectorSiteModel.ID = core.StringPtr("testString")
				newVdcDirectorSiteModel.Cluster = vdcDirectorSiteClusterModel

				// Construct an instance of the NewVDCEdge model
				newVdcEdgeModel := new(vmwarev1.NewVDCEdge)
				newVdcEdgeModel.Size = core.StringPtr("medium")
				newVdcEdgeModel.Type = core.StringPtr("dedicated")

				// Construct an instance of the NewVDCResourceGroup model
				newVdcResourceGroupModel := new(vmwarev1.NewVDCResourceGroup)
				newVdcResourceGroupModel.ID = core.StringPtr("testString")

				// Construct an instance of the CreateVdcOptions model
				createVdcOptionsModel := new(vmwarev1.CreateVdcOptions)
				createVdcOptionsModel.Name = core.StringPtr("testString")
				createVdcOptionsModel.DirectorSite = newVdcDirectorSiteModel
				createVdcOptionsModel.Edge = newVdcEdgeModel
				createVdcOptionsModel.ResourceGroup = newVdcResourceGroupModel
				createVdcOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createVdcOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vmwareService.CreateVdc(createVdcOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vmwareService.EnableRetries(0, 0)
				result, response, operationErr = vmwareService.CreateVdc(createVdcOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateVdc(createVdcOptions *CreateVdcOptions)`, func() {
		createVdcPath := "/vdcs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createVdcPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"id": "ID", "allocation_model": "paygo", "created_time": "2019-01-01T12:00:00.000Z", "crn": "Crn", "deleted_time": "2019-01-01T12:00:00.000Z", "director_site": {"id": "ID", "cluster": {"id": "ID"}, "url": "URL"}, "edges": [{"id": "ID", "public_ips": ["PublicIps"], "size": "medium", "type": "dedicated"}], "errors": [{"code": "Code", "message": "Message", "more_info": "MoreInfo"}], "name": "Name", "ordered_time": "2019-01-01T12:00:00.000Z", "org_name": "OrgName", "status": "Creating", "type": "dedicated"}`)
				}))
			})
			It(`Invoke CreateVdc successfully with retries`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())
				vmwareService.EnableRetries(0, 0)

				// Construct an instance of the VDCDirectorSiteCluster model
				vdcDirectorSiteClusterModel := new(vmwarev1.VDCDirectorSiteCluster)
				vdcDirectorSiteClusterModel.ID = core.StringPtr("testString")

				// Construct an instance of the NewVDCDirectorSite model
				newVdcDirectorSiteModel := new(vmwarev1.NewVDCDirectorSite)
				newVdcDirectorSiteModel.ID = core.StringPtr("testString")
				newVdcDirectorSiteModel.Cluster = vdcDirectorSiteClusterModel

				// Construct an instance of the NewVDCEdge model
				newVdcEdgeModel := new(vmwarev1.NewVDCEdge)
				newVdcEdgeModel.Size = core.StringPtr("medium")
				newVdcEdgeModel.Type = core.StringPtr("dedicated")

				// Construct an instance of the NewVDCResourceGroup model
				newVdcResourceGroupModel := new(vmwarev1.NewVDCResourceGroup)
				newVdcResourceGroupModel.ID = core.StringPtr("testString")

				// Construct an instance of the CreateVdcOptions model
				createVdcOptionsModel := new(vmwarev1.CreateVdcOptions)
				createVdcOptionsModel.Name = core.StringPtr("testString")
				createVdcOptionsModel.DirectorSite = newVdcDirectorSiteModel
				createVdcOptionsModel.Edge = newVdcEdgeModel
				createVdcOptionsModel.ResourceGroup = newVdcResourceGroupModel
				createVdcOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createVdcOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vmwareService.CreateVdcWithContext(ctx, createVdcOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vmwareService.DisableRetries()
				result, response, operationErr := vmwareService.CreateVdc(createVdcOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vmwareService.CreateVdcWithContext(ctx, createVdcOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createVdcPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"id": "ID", "allocation_model": "paygo", "created_time": "2019-01-01T12:00:00.000Z", "crn": "Crn", "deleted_time": "2019-01-01T12:00:00.000Z", "director_site": {"id": "ID", "cluster": {"id": "ID"}, "url": "URL"}, "edges": [{"id": "ID", "public_ips": ["PublicIps"], "size": "medium", "type": "dedicated"}], "errors": [{"code": "Code", "message": "Message", "more_info": "MoreInfo"}], "name": "Name", "ordered_time": "2019-01-01T12:00:00.000Z", "org_name": "OrgName", "status": "Creating", "type": "dedicated"}`)
				}))
			})
			It(`Invoke CreateVdc successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vmwareService.CreateVdc(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the VDCDirectorSiteCluster model
				vdcDirectorSiteClusterModel := new(vmwarev1.VDCDirectorSiteCluster)
				vdcDirectorSiteClusterModel.ID = core.StringPtr("testString")

				// Construct an instance of the NewVDCDirectorSite model
				newVdcDirectorSiteModel := new(vmwarev1.NewVDCDirectorSite)
				newVdcDirectorSiteModel.ID = core.StringPtr("testString")
				newVdcDirectorSiteModel.Cluster = vdcDirectorSiteClusterModel

				// Construct an instance of the NewVDCEdge model
				newVdcEdgeModel := new(vmwarev1.NewVDCEdge)
				newVdcEdgeModel.Size = core.StringPtr("medium")
				newVdcEdgeModel.Type = core.StringPtr("dedicated")

				// Construct an instance of the NewVDCResourceGroup model
				newVdcResourceGroupModel := new(vmwarev1.NewVDCResourceGroup)
				newVdcResourceGroupModel.ID = core.StringPtr("testString")

				// Construct an instance of the CreateVdcOptions model
				createVdcOptionsModel := new(vmwarev1.CreateVdcOptions)
				createVdcOptionsModel.Name = core.StringPtr("testString")
				createVdcOptionsModel.DirectorSite = newVdcDirectorSiteModel
				createVdcOptionsModel.Edge = newVdcEdgeModel
				createVdcOptionsModel.ResourceGroup = newVdcResourceGroupModel
				createVdcOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createVdcOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vmwareService.CreateVdc(createVdcOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateVdc with error: Operation validation and request error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the VDCDirectorSiteCluster model
				vdcDirectorSiteClusterModel := new(vmwarev1.VDCDirectorSiteCluster)
				vdcDirectorSiteClusterModel.ID = core.StringPtr("testString")

				// Construct an instance of the NewVDCDirectorSite model
				newVdcDirectorSiteModel := new(vmwarev1.NewVDCDirectorSite)
				newVdcDirectorSiteModel.ID = core.StringPtr("testString")
				newVdcDirectorSiteModel.Cluster = vdcDirectorSiteClusterModel

				// Construct an instance of the NewVDCEdge model
				newVdcEdgeModel := new(vmwarev1.NewVDCEdge)
				newVdcEdgeModel.Size = core.StringPtr("medium")
				newVdcEdgeModel.Type = core.StringPtr("dedicated")

				// Construct an instance of the NewVDCResourceGroup model
				newVdcResourceGroupModel := new(vmwarev1.NewVDCResourceGroup)
				newVdcResourceGroupModel.ID = core.StringPtr("testString")

				// Construct an instance of the CreateVdcOptions model
				createVdcOptionsModel := new(vmwarev1.CreateVdcOptions)
				createVdcOptionsModel.Name = core.StringPtr("testString")
				createVdcOptionsModel.DirectorSite = newVdcDirectorSiteModel
				createVdcOptionsModel.Edge = newVdcEdgeModel
				createVdcOptionsModel.ResourceGroup = newVdcResourceGroupModel
				createVdcOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createVdcOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vmwareService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vmwareService.CreateVdc(createVdcOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateVdcOptions model with no property values
				createVdcOptionsModelNew := new(vmwarev1.CreateVdcOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = vmwareService.CreateVdc(createVdcOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(202)
				}))
			})
			It(`Invoke CreateVdc successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the VDCDirectorSiteCluster model
				vdcDirectorSiteClusterModel := new(vmwarev1.VDCDirectorSiteCluster)
				vdcDirectorSiteClusterModel.ID = core.StringPtr("testString")

				// Construct an instance of the NewVDCDirectorSite model
				newVdcDirectorSiteModel := new(vmwarev1.NewVDCDirectorSite)
				newVdcDirectorSiteModel.ID = core.StringPtr("testString")
				newVdcDirectorSiteModel.Cluster = vdcDirectorSiteClusterModel

				// Construct an instance of the NewVDCEdge model
				newVdcEdgeModel := new(vmwarev1.NewVDCEdge)
				newVdcEdgeModel.Size = core.StringPtr("medium")
				newVdcEdgeModel.Type = core.StringPtr("dedicated")

				// Construct an instance of the NewVDCResourceGroup model
				newVdcResourceGroupModel := new(vmwarev1.NewVDCResourceGroup)
				newVdcResourceGroupModel.ID = core.StringPtr("testString")

				// Construct an instance of the CreateVdcOptions model
				createVdcOptionsModel := new(vmwarev1.CreateVdcOptions)
				createVdcOptionsModel.Name = core.StringPtr("testString")
				createVdcOptionsModel.DirectorSite = newVdcDirectorSiteModel
				createVdcOptionsModel.Edge = newVdcEdgeModel
				createVdcOptionsModel.ResourceGroup = newVdcResourceGroupModel
				createVdcOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createVdcOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vmwareService.CreateVdc(createVdcOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetVdc(getVdcOptions *GetVdcOptions) - Operation response error`, func() {
		getVdcPath := "/vdcs/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getVdcPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetVdc with error: Operation response processing error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the GetVdcOptions model
				getVdcOptionsModel := new(vmwarev1.GetVdcOptions)
				getVdcOptionsModel.VdcID = core.StringPtr("testString")
				getVdcOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getVdcOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vmwareService.GetVdc(getVdcOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vmwareService.EnableRetries(0, 0)
				result, response, operationErr = vmwareService.GetVdc(getVdcOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetVdc(getVdcOptions *GetVdcOptions)`, func() {
		getVdcPath := "/vdcs/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getVdcPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "allocation_model": "paygo", "created_time": "2019-01-01T12:00:00.000Z", "crn": "Crn", "deleted_time": "2019-01-01T12:00:00.000Z", "director_site": {"id": "ID", "cluster": {"id": "ID"}, "url": "URL"}, "edges": [{"id": "ID", "public_ips": ["PublicIps"], "size": "medium", "type": "dedicated"}], "errors": [{"code": "Code", "message": "Message", "more_info": "MoreInfo"}], "name": "Name", "ordered_time": "2019-01-01T12:00:00.000Z", "org_name": "OrgName", "status": "Creating", "type": "dedicated"}`)
				}))
			})
			It(`Invoke GetVdc successfully with retries`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())
				vmwareService.EnableRetries(0, 0)

				// Construct an instance of the GetVdcOptions model
				getVdcOptionsModel := new(vmwarev1.GetVdcOptions)
				getVdcOptionsModel.VdcID = core.StringPtr("testString")
				getVdcOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getVdcOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vmwareService.GetVdcWithContext(ctx, getVdcOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vmwareService.DisableRetries()
				result, response, operationErr := vmwareService.GetVdc(getVdcOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vmwareService.GetVdcWithContext(ctx, getVdcOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getVdcPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "allocation_model": "paygo", "created_time": "2019-01-01T12:00:00.000Z", "crn": "Crn", "deleted_time": "2019-01-01T12:00:00.000Z", "director_site": {"id": "ID", "cluster": {"id": "ID"}, "url": "URL"}, "edges": [{"id": "ID", "public_ips": ["PublicIps"], "size": "medium", "type": "dedicated"}], "errors": [{"code": "Code", "message": "Message", "more_info": "MoreInfo"}], "name": "Name", "ordered_time": "2019-01-01T12:00:00.000Z", "org_name": "OrgName", "status": "Creating", "type": "dedicated"}`)
				}))
			})
			It(`Invoke GetVdc successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vmwareService.GetVdc(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetVdcOptions model
				getVdcOptionsModel := new(vmwarev1.GetVdcOptions)
				getVdcOptionsModel.VdcID = core.StringPtr("testString")
				getVdcOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getVdcOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vmwareService.GetVdc(getVdcOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetVdc with error: Operation validation and request error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the GetVdcOptions model
				getVdcOptionsModel := new(vmwarev1.GetVdcOptions)
				getVdcOptionsModel.VdcID = core.StringPtr("testString")
				getVdcOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getVdcOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vmwareService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vmwareService.GetVdc(getVdcOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetVdcOptions model with no property values
				getVdcOptionsModelNew := new(vmwarev1.GetVdcOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = vmwareService.GetVdc(getVdcOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetVdc successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the GetVdcOptions model
				getVdcOptionsModel := new(vmwarev1.GetVdcOptions)
				getVdcOptionsModel.VdcID = core.StringPtr("testString")
				getVdcOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getVdcOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vmwareService.GetVdc(getVdcOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteVdc(deleteVdcOptions *DeleteVdcOptions) - Operation response error`, func() {
		deleteVdcPath := "/vdcs/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteVdcPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteVdc with error: Operation response processing error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the DeleteVdcOptions model
				deleteVdcOptionsModel := new(vmwarev1.DeleteVdcOptions)
				deleteVdcOptionsModel.VdcID = core.StringPtr("testString")
				deleteVdcOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteVdcOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vmwareService.DeleteVdc(deleteVdcOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vmwareService.EnableRetries(0, 0)
				result, response, operationErr = vmwareService.DeleteVdc(deleteVdcOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteVdc(deleteVdcOptions *DeleteVdcOptions)`, func() {
		deleteVdcPath := "/vdcs/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteVdcPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"id": "ID", "allocation_model": "paygo", "created_time": "2019-01-01T12:00:00.000Z", "crn": "Crn", "deleted_time": "2019-01-01T12:00:00.000Z", "director_site": {"id": "ID", "cluster": {"id": "ID"}, "url": "URL"}, "edges": [{"id": "ID", "public_ips": ["PublicIps"], "size": "medium", "type": "dedicated"}], "errors": [{"code": "Code", "message": "Message", "more_info": "MoreInfo"}], "name": "Name", "ordered_time": "2019-01-01T12:00:00.000Z", "org_name": "OrgName", "status": "Creating", "type": "dedicated"}`)
				}))
			})
			It(`Invoke DeleteVdc successfully with retries`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())
				vmwareService.EnableRetries(0, 0)

				// Construct an instance of the DeleteVdcOptions model
				deleteVdcOptionsModel := new(vmwarev1.DeleteVdcOptions)
				deleteVdcOptionsModel.VdcID = core.StringPtr("testString")
				deleteVdcOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteVdcOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vmwareService.DeleteVdcWithContext(ctx, deleteVdcOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vmwareService.DisableRetries()
				result, response, operationErr := vmwareService.DeleteVdc(deleteVdcOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vmwareService.DeleteVdcWithContext(ctx, deleteVdcOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteVdcPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"id": "ID", "allocation_model": "paygo", "created_time": "2019-01-01T12:00:00.000Z", "crn": "Crn", "deleted_time": "2019-01-01T12:00:00.000Z", "director_site": {"id": "ID", "cluster": {"id": "ID"}, "url": "URL"}, "edges": [{"id": "ID", "public_ips": ["PublicIps"], "size": "medium", "type": "dedicated"}], "errors": [{"code": "Code", "message": "Message", "more_info": "MoreInfo"}], "name": "Name", "ordered_time": "2019-01-01T12:00:00.000Z", "org_name": "OrgName", "status": "Creating", "type": "dedicated"}`)
				}))
			})
			It(`Invoke DeleteVdc successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vmwareService.DeleteVdc(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteVdcOptions model
				deleteVdcOptionsModel := new(vmwarev1.DeleteVdcOptions)
				deleteVdcOptionsModel.VdcID = core.StringPtr("testString")
				deleteVdcOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteVdcOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vmwareService.DeleteVdc(deleteVdcOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteVdc with error: Operation validation and request error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the DeleteVdcOptions model
				deleteVdcOptionsModel := new(vmwarev1.DeleteVdcOptions)
				deleteVdcOptionsModel.VdcID = core.StringPtr("testString")
				deleteVdcOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteVdcOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vmwareService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vmwareService.DeleteVdc(deleteVdcOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteVdcOptions model with no property values
				deleteVdcOptionsModelNew := new(vmwarev1.DeleteVdcOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = vmwareService.DeleteVdc(deleteVdcOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(202)
				}))
			})
			It(`Invoke DeleteVdc successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the DeleteVdcOptions model
				deleteVdcOptionsModel := new(vmwarev1.DeleteVdcOptions)
				deleteVdcOptionsModel.VdcID = core.StringPtr("testString")
				deleteVdcOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteVdcOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vmwareService.DeleteVdc(deleteVdcOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			vmwareService, _ := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
				URL:           "http://vmwarev1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewClusterOrderInfo successfully`, func() {
				name := "testString"
				location := "testString"
				hostCount := int64(2)
				var fileShares *vmwarev1.FileShares = nil
				hostProfile := "testString"
				_, err := vmwareService.NewClusterOrderInfo(name, location, hostCount, fileShares, hostProfile)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewCreateVdcOptions successfully`, func() {
				// Construct an instance of the VDCDirectorSiteCluster model
				vdcDirectorSiteClusterModel := new(vmwarev1.VDCDirectorSiteCluster)
				Expect(vdcDirectorSiteClusterModel).ToNot(BeNil())
				vdcDirectorSiteClusterModel.ID = core.StringPtr("testString")
				Expect(vdcDirectorSiteClusterModel.ID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the NewVDCDirectorSite model
				newVdcDirectorSiteModel := new(vmwarev1.NewVDCDirectorSite)
				Expect(newVdcDirectorSiteModel).ToNot(BeNil())
				newVdcDirectorSiteModel.ID = core.StringPtr("testString")
				newVdcDirectorSiteModel.Cluster = vdcDirectorSiteClusterModel
				Expect(newVdcDirectorSiteModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(newVdcDirectorSiteModel.Cluster).To(Equal(vdcDirectorSiteClusterModel))

				// Construct an instance of the NewVDCEdge model
				newVdcEdgeModel := new(vmwarev1.NewVDCEdge)
				Expect(newVdcEdgeModel).ToNot(BeNil())
				newVdcEdgeModel.Size = core.StringPtr("medium")
				newVdcEdgeModel.Type = core.StringPtr("dedicated")
				Expect(newVdcEdgeModel.Size).To(Equal(core.StringPtr("medium")))
				Expect(newVdcEdgeModel.Type).To(Equal(core.StringPtr("dedicated")))

				// Construct an instance of the NewVDCResourceGroup model
				newVdcResourceGroupModel := new(vmwarev1.NewVDCResourceGroup)
				Expect(newVdcResourceGroupModel).ToNot(BeNil())
				newVdcResourceGroupModel.ID = core.StringPtr("testString")
				Expect(newVdcResourceGroupModel.ID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the CreateVdcOptions model
				createVdcOptionsName := "testString"
				var createVdcOptionsDirectorSite *vmwarev1.NewVDCDirectorSite = nil
				createVdcOptionsModel := vmwareService.NewCreateVdcOptions(createVdcOptionsName, createVdcOptionsDirectorSite)
				createVdcOptionsModel.SetName("testString")
				createVdcOptionsModel.SetDirectorSite(newVdcDirectorSiteModel)
				createVdcOptionsModel.SetEdge(newVdcEdgeModel)
				createVdcOptionsModel.SetResourceGroup(newVdcResourceGroupModel)
				createVdcOptionsModel.SetAcceptLanguage("testString")
				createVdcOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createVdcOptionsModel).ToNot(BeNil())
				Expect(createVdcOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createVdcOptionsModel.DirectorSite).To(Equal(newVdcDirectorSiteModel))
				Expect(createVdcOptionsModel.Edge).To(Equal(newVdcEdgeModel))
				Expect(createVdcOptionsModel.ResourceGroup).To(Equal(newVdcResourceGroupModel))
				Expect(createVdcOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(createVdcOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateWorkloadDomainOptions successfully`, func() {
				// Construct an instance of the FileShares model
				fileSharesModel := new(vmwarev1.FileShares)
				Expect(fileSharesModel).ToNot(BeNil())
				fileSharesModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))
				Expect(fileSharesModel.STORAGEPOINTTWOFIVEIOPSGB).To(Equal(core.Int64Ptr(int64(0))))
				Expect(fileSharesModel.STORAGETWOIOPSGB).To(Equal(core.Int64Ptr(int64(0))))
				Expect(fileSharesModel.STORAGEFOURIOPSGB).To(Equal(core.Int64Ptr(int64(0))))
				Expect(fileSharesModel.STORAGETENIOPSGB).To(Equal(core.Int64Ptr(int64(0))))

				// Construct an instance of the ClusterOrderInfo model
				clusterOrderInfoModel := new(vmwarev1.ClusterOrderInfo)
				Expect(clusterOrderInfoModel).ToNot(BeNil())
				clusterOrderInfoModel.Name = core.StringPtr("testString")
				clusterOrderInfoModel.Location = core.StringPtr("testString")
				clusterOrderInfoModel.HostCount = core.Int64Ptr(int64(2))
				clusterOrderInfoModel.FileShares = fileSharesModel
				clusterOrderInfoModel.HostProfile = core.StringPtr("testString")
				Expect(clusterOrderInfoModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(clusterOrderInfoModel.Location).To(Equal(core.StringPtr("testString")))
				Expect(clusterOrderInfoModel.HostCount).To(Equal(core.Int64Ptr(int64(2))))
				Expect(clusterOrderInfoModel.FileShares).To(Equal(fileSharesModel))
				Expect(clusterOrderInfoModel.HostProfile).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the CreateWorkloadDomainOptions model
				createWorkloadDomainOptionsName := "testString"
				createWorkloadDomainOptionsResourceGroup := "testString"
				createWorkloadDomainOptionsClusters := []vmwarev1.ClusterOrderInfo{}
				createWorkloadDomainOptionsModel := vmwareService.NewCreateWorkloadDomainOptions(createWorkloadDomainOptionsName, createWorkloadDomainOptionsResourceGroup, createWorkloadDomainOptionsClusters)
				createWorkloadDomainOptionsModel.SetName("testString")
				createWorkloadDomainOptionsModel.SetResourceGroup("testString")
				createWorkloadDomainOptionsModel.SetClusters([]vmwarev1.ClusterOrderInfo{*clusterOrderInfoModel})
				createWorkloadDomainOptionsModel.SetAcceptLanguage("testString")
				createWorkloadDomainOptionsModel.SetXGlobalTransactionID("testString")
				createWorkloadDomainOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createWorkloadDomainOptionsModel).ToNot(BeNil())
				Expect(createWorkloadDomainOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createWorkloadDomainOptionsModel.ResourceGroup).To(Equal(core.StringPtr("testString")))
				Expect(createWorkloadDomainOptionsModel.Clusters).To(Equal([]vmwarev1.ClusterOrderInfo{*clusterOrderInfoModel}))
				Expect(createWorkloadDomainOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(createWorkloadDomainOptionsModel.XGlobalTransactionID).To(Equal(core.StringPtr("testString")))
				Expect(createWorkloadDomainOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteVdcOptions successfully`, func() {
				// Construct an instance of the DeleteVdcOptions model
				vdcID := "testString"
				deleteVdcOptionsModel := vmwareService.NewDeleteVdcOptions(vdcID)
				deleteVdcOptionsModel.SetVdcID("testString")
				deleteVdcOptionsModel.SetAcceptLanguage("testString")
				deleteVdcOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteVdcOptionsModel).ToNot(BeNil())
				Expect(deleteVdcOptionsModel.VdcID).To(Equal(core.StringPtr("testString")))
				Expect(deleteVdcOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(deleteVdcOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteWorkloadDomainOptions successfully`, func() {
				// Construct an instance of the DeleteWorkloadDomainOptions model
				siteID := "testString"
				deleteWorkloadDomainOptionsModel := vmwareService.NewDeleteWorkloadDomainOptions(siteID)
				deleteWorkloadDomainOptionsModel.SetSiteID("testString")
				deleteWorkloadDomainOptionsModel.SetAcceptLanguage("testString")
				deleteWorkloadDomainOptionsModel.SetXGlobalTransactionID("testString")
				deleteWorkloadDomainOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteWorkloadDomainOptionsModel).ToNot(BeNil())
				Expect(deleteWorkloadDomainOptionsModel.SiteID).To(Equal(core.StringPtr("testString")))
				Expect(deleteWorkloadDomainOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(deleteWorkloadDomainOptionsModel.XGlobalTransactionID).To(Equal(core.StringPtr("testString")))
				Expect(deleteWorkloadDomainOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetRegionsOptions successfully`, func() {
				// Construct an instance of the GetRegionsOptions model
				getRegionsOptionsModel := vmwareService.NewGetRegionsOptions()
				getRegionsOptionsModel.SetAcceptLanguage("testString")
				getRegionsOptionsModel.SetXGlobalTransactionID("testString")
				getRegionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getRegionsOptionsModel).ToNot(BeNil())
				Expect(getRegionsOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(getRegionsOptionsModel.XGlobalTransactionID).To(Equal(core.StringPtr("testString")))
				Expect(getRegionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSpecificClusterInstanceOptions successfully`, func() {
				// Construct an instance of the GetSpecificClusterInstanceOptions model
				siteID := "testString"
				clusterID := "testString"
				getSpecificClusterInstanceOptionsModel := vmwareService.NewGetSpecificClusterInstanceOptions(siteID, clusterID)
				getSpecificClusterInstanceOptionsModel.SetSiteID("testString")
				getSpecificClusterInstanceOptionsModel.SetClusterID("testString")
				getSpecificClusterInstanceOptionsModel.SetAcceptLanguage("testString")
				getSpecificClusterInstanceOptionsModel.SetXGlobalTransactionID("testString")
				getSpecificClusterInstanceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSpecificClusterInstanceOptionsModel).ToNot(BeNil())
				Expect(getSpecificClusterInstanceOptionsModel.SiteID).To(Equal(core.StringPtr("testString")))
				Expect(getSpecificClusterInstanceOptionsModel.ClusterID).To(Equal(core.StringPtr("testString")))
				Expect(getSpecificClusterInstanceOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(getSpecificClusterInstanceOptionsModel.XGlobalTransactionID).To(Equal(core.StringPtr("testString")))
				Expect(getSpecificClusterInstanceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSpecificWorkloadDomainInstanceOptions successfully`, func() {
				// Construct an instance of the GetSpecificWorkloadDomainInstanceOptions model
				siteID := "testString"
				getSpecificWorkloadDomainInstanceOptionsModel := vmwareService.NewGetSpecificWorkloadDomainInstanceOptions(siteID)
				getSpecificWorkloadDomainInstanceOptionsModel.SetSiteID("testString")
				getSpecificWorkloadDomainInstanceOptionsModel.SetAcceptLanguage("testString")
				getSpecificWorkloadDomainInstanceOptionsModel.SetXGlobalTransactionID("testString")
				getSpecificWorkloadDomainInstanceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSpecificWorkloadDomainInstanceOptionsModel).ToNot(BeNil())
				Expect(getSpecificWorkloadDomainInstanceOptionsModel.SiteID).To(Equal(core.StringPtr("testString")))
				Expect(getSpecificWorkloadDomainInstanceOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(getSpecificWorkloadDomainInstanceOptionsModel.XGlobalTransactionID).To(Equal(core.StringPtr("testString")))
				Expect(getSpecificWorkloadDomainInstanceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetVcddPriceOptions successfully`, func() {
				// Construct an instance of the FileShares model
				fileSharesModel := new(vmwarev1.FileShares)
				Expect(fileSharesModel).ToNot(BeNil())
				fileSharesModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))
				Expect(fileSharesModel.STORAGEPOINTTWOFIVEIOPSGB).To(Equal(core.Int64Ptr(int64(0))))
				Expect(fileSharesModel.STORAGETWOIOPSGB).To(Equal(core.Int64Ptr(int64(0))))
				Expect(fileSharesModel.STORAGEFOURIOPSGB).To(Equal(core.Int64Ptr(int64(0))))
				Expect(fileSharesModel.STORAGETENIOPSGB).To(Equal(core.Int64Ptr(int64(0))))

				// Construct an instance of the DirectorSitePriceQuoteClusterInfo model
				directorSitePriceQuoteClusterInfoModel := new(vmwarev1.DirectorSitePriceQuoteClusterInfo)
				Expect(directorSitePriceQuoteClusterInfoModel).ToNot(BeNil())
				directorSitePriceQuoteClusterInfoModel.Name = core.StringPtr("testString")
				directorSitePriceQuoteClusterInfoModel.HostProfile = core.StringPtr("testString")
				directorSitePriceQuoteClusterInfoModel.HostCount = core.Int64Ptr(int64(2))
				directorSitePriceQuoteClusterInfoModel.FileShares = fileSharesModel
				Expect(directorSitePriceQuoteClusterInfoModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(directorSitePriceQuoteClusterInfoModel.HostProfile).To(Equal(core.StringPtr("testString")))
				Expect(directorSitePriceQuoteClusterInfoModel.HostCount).To(Equal(core.Int64Ptr(int64(2))))
				Expect(directorSitePriceQuoteClusterInfoModel.FileShares).To(Equal(fileSharesModel))

				// Construct an instance of the GetVcddPriceOptions model
				getVcddPriceOptionsModel := vmwareService.NewGetVcddPriceOptions()
				getVcddPriceOptionsModel.SetCountry("USA")
				getVcddPriceOptionsModel.SetClusters([]vmwarev1.DirectorSitePriceQuoteClusterInfo{*directorSitePriceQuoteClusterInfoModel})
				getVcddPriceOptionsModel.SetAcceptLanguage("testString")
				getVcddPriceOptionsModel.SetXGlobalTransactionID("testString")
				getVcddPriceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getVcddPriceOptionsModel).ToNot(BeNil())
				Expect(getVcddPriceOptionsModel.Country).To(Equal(core.StringPtr("USA")))
				Expect(getVcddPriceOptionsModel.Clusters).To(Equal([]vmwarev1.DirectorSitePriceQuoteClusterInfo{*directorSitePriceQuoteClusterInfoModel}))
				Expect(getVcddPriceOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(getVcddPriceOptionsModel.XGlobalTransactionID).To(Equal(core.StringPtr("testString")))
				Expect(getVcddPriceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetVdcOptions successfully`, func() {
				// Construct an instance of the GetVdcOptions model
				vdcID := "testString"
				getVdcOptionsModel := vmwareService.NewGetVdcOptions(vdcID)
				getVdcOptionsModel.SetVdcID("testString")
				getVdcOptionsModel.SetAcceptLanguage("testString")
				getVdcOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getVdcOptionsModel).ToNot(BeNil())
				Expect(getVdcOptionsModel.VdcID).To(Equal(core.StringPtr("testString")))
				Expect(getVdcOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(getVdcOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListClusterInstancesOptions successfully`, func() {
				// Construct an instance of the ListClusterInstancesOptions model
				siteID := "testString"
				listClusterInstancesOptionsModel := vmwareService.NewListClusterInstancesOptions(siteID)
				listClusterInstancesOptionsModel.SetSiteID("testString")
				listClusterInstancesOptionsModel.SetAcceptLanguage("testString")
				listClusterInstancesOptionsModel.SetXGlobalTransactionID("testString")
				listClusterInstancesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listClusterInstancesOptionsModel).ToNot(BeNil())
				Expect(listClusterInstancesOptionsModel.SiteID).To(Equal(core.StringPtr("testString")))
				Expect(listClusterInstancesOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(listClusterInstancesOptionsModel.XGlobalTransactionID).To(Equal(core.StringPtr("testString")))
				Expect(listClusterInstancesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListPricesOptions successfully`, func() {
				// Construct an instance of the ListPricesOptions model
				listPricesOptionsModel := vmwareService.NewListPricesOptions()
				listPricesOptionsModel.SetAcceptLanguage("testString")
				listPricesOptionsModel.SetXGlobalTransactionID("testString")
				listPricesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listPricesOptionsModel).ToNot(BeNil())
				Expect(listPricesOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(listPricesOptionsModel.XGlobalTransactionID).To(Equal(core.StringPtr("testString")))
				Expect(listPricesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListVdcsOptions successfully`, func() {
				// Construct an instance of the ListVdcsOptions model
				listVdcsOptionsModel := vmwareService.NewListVdcsOptions()
				listVdcsOptionsModel.SetAcceptLanguage("testString")
				listVdcsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listVdcsOptionsModel).ToNot(BeNil())
				Expect(listVdcsOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(listVdcsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListWorkloadDomainInstancesOptions successfully`, func() {
				// Construct an instance of the ListWorkloadDomainInstancesOptions model
				listWorkloadDomainInstancesOptionsModel := vmwareService.NewListWorkloadDomainInstancesOptions()
				listWorkloadDomainInstancesOptionsModel.SetAcceptLanguage("testString")
				listWorkloadDomainInstancesOptionsModel.SetXGlobalTransactionID("testString")
				listWorkloadDomainInstancesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listWorkloadDomainInstancesOptionsModel).ToNot(BeNil())
				Expect(listWorkloadDomainInstancesOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(listWorkloadDomainInstancesOptionsModel.XGlobalTransactionID).To(Equal(core.StringPtr("testString")))
				Expect(listWorkloadDomainInstancesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewNewVDCDirectorSite successfully`, func() {
				id := "testString"
				var cluster *vmwarev1.VDCDirectorSiteCluster = nil
				_, err := vmwareService.NewNewVDCDirectorSite(id, cluster)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewNewVDCEdge successfully`, func() {
				typeVar := "dedicated"
				_model, err := vmwareService.NewNewVDCEdge(typeVar)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewNewVDCResourceGroup successfully`, func() {
				id := "testString"
				_model, err := vmwareService.NewNewVDCResourceGroup(id)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewReplaceOrgAdminPasswordOptions successfully`, func() {
				// Construct an instance of the ReplaceOrgAdminPasswordOptions model
				siteID := "testString"
				replaceOrgAdminPasswordOptionsModel := vmwareService.NewReplaceOrgAdminPasswordOptions(siteID)
				replaceOrgAdminPasswordOptionsModel.SetSiteID("testString")
				replaceOrgAdminPasswordOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceOrgAdminPasswordOptionsModel).ToNot(BeNil())
				Expect(replaceOrgAdminPasswordOptionsModel.SiteID).To(Equal(core.StringPtr("testString")))
				Expect(replaceOrgAdminPasswordOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSetFileSharesOptions successfully`, func() {
				// Construct an instance of the SetFileSharesOptions model
				siteID := "testString"
				clusterID := "testString"
				setFileSharesOptionsModel := vmwareService.NewSetFileSharesOptions(siteID, clusterID)
				setFileSharesOptionsModel.SetSiteID("testString")
				setFileSharesOptionsModel.SetClusterID("testString")
				setFileSharesOptionsModel.SetSTORAGEPOINTTWOFIVEIOPSGB(int64(0))
				setFileSharesOptionsModel.SetSTORAGETWOIOPSGB(int64(0))
				setFileSharesOptionsModel.SetSTORAGEFOURIOPSGB(int64(0))
				setFileSharesOptionsModel.SetSTORAGETENIOPSGB(int64(0))
				setFileSharesOptionsModel.SetAcceptLanguage("testString")
				setFileSharesOptionsModel.SetXGlobalTransactionID("testString")
				setFileSharesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(setFileSharesOptionsModel).ToNot(BeNil())
				Expect(setFileSharesOptionsModel.SiteID).To(Equal(core.StringPtr("testString")))
				Expect(setFileSharesOptionsModel.ClusterID).To(Equal(core.StringPtr("testString")))
				Expect(setFileSharesOptionsModel.STORAGEPOINTTWOFIVEIOPSGB).To(Equal(core.Int64Ptr(int64(0))))
				Expect(setFileSharesOptionsModel.STORAGETWOIOPSGB).To(Equal(core.Int64Ptr(int64(0))))
				Expect(setFileSharesOptionsModel.STORAGEFOURIOPSGB).To(Equal(core.Int64Ptr(int64(0))))
				Expect(setFileSharesOptionsModel.STORAGETENIOPSGB).To(Equal(core.Int64Ptr(int64(0))))
				Expect(setFileSharesOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(setFileSharesOptionsModel.XGlobalTransactionID).To(Equal(core.StringPtr("testString")))
				Expect(setFileSharesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSetHostsCountOptions successfully`, func() {
				// Construct an instance of the SetHostsCountOptions model
				siteID := "testString"
				clusterID := "testString"
				setHostsCountOptionsCount := int64(2)
				setHostsCountOptionsModel := vmwareService.NewSetHostsCountOptions(siteID, clusterID, setHostsCountOptionsCount)
				setHostsCountOptionsModel.SetSiteID("testString")
				setHostsCountOptionsModel.SetClusterID("testString")
				setHostsCountOptionsModel.SetCount(int64(2))
				setHostsCountOptionsModel.SetAcceptLanguage("testString")
				setHostsCountOptionsModel.SetXGlobalTransactionID("testString")
				setHostsCountOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(setHostsCountOptionsModel).ToNot(BeNil())
				Expect(setHostsCountOptionsModel.SiteID).To(Equal(core.StringPtr("testString")))
				Expect(setHostsCountOptionsModel.ClusterID).To(Equal(core.StringPtr("testString")))
				Expect(setHostsCountOptionsModel.Count).To(Equal(core.Int64Ptr(int64(2))))
				Expect(setHostsCountOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(setHostsCountOptionsModel.XGlobalTransactionID).To(Equal(core.StringPtr("testString")))
				Expect(setHostsCountOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewVDCDirectorSiteCluster successfully`, func() {
				id := "testString"
				_model, err := vmwareService.NewVDCDirectorSiteCluster(id)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewViewInstanceOptions successfully`, func() {
				// Construct an instance of the ViewInstanceOptions model
				viewInstanceOptionsModel := vmwareService.NewViewInstanceOptions()
				viewInstanceOptionsModel.SetAcceptLanguage("testString")
				viewInstanceOptionsModel.SetXGlobalTransactionID("testString")
				viewInstanceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(viewInstanceOptionsModel).ToNot(BeNil())
				Expect(viewInstanceOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(viewInstanceOptionsModel.XGlobalTransactionID).To(Equal(core.StringPtr("testString")))
				Expect(viewInstanceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("This is a test")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It(`Invoke CreateMockReader() successfully`, func() {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate("2019-01-01")
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime("2019-01-01T12:00:00.000Z")
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, 0)
	ba = append(ba, mockData...)
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return io.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate(mockData string) *strfmt.Date {
	d, err := core.ParseDate(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func CreateMockDateTime(mockData string) *strfmt.DateTime {
	d, err := core.ParseDateTime(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func SetTestEnvironment(testEnvironment map[string]string) {
	for key, value := range testEnvironment {
		os.Setenv(key, value)
	}
}

func ClearTestEnvironment(testEnvironment map[string]string) {
	for key := range testEnvironment {
		os.Unsetenv(key)
	}
}
