package mse

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// AddSSLCert invokes the mse.AddSSLCert API synchronously
func (client *Client) AddSSLCert(request *AddSSLCertRequest) (response *AddSSLCertResponse, err error) {
	response = CreateAddSSLCertResponse()
	err = client.DoAction(request, response)
	return
}

// AddSSLCertWithChan invokes the mse.AddSSLCert API asynchronously
func (client *Client) AddSSLCertWithChan(request *AddSSLCertRequest) (<-chan *AddSSLCertResponse, <-chan error) {
	responseChan := make(chan *AddSSLCertResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddSSLCert(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// AddSSLCertWithCallback invokes the mse.AddSSLCert API asynchronously
func (client *Client) AddSSLCertWithCallback(request *AddSSLCertRequest, callback func(response *AddSSLCertResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddSSLCertResponse
		var err error
		defer close(result)
		response, err = client.AddSSLCert(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// AddSSLCertRequest is the request struct for api AddSSLCert
type AddSSLCertRequest struct {
	*requests.RpcRequest
	GatewayUniqueId string           `position:"Query" name:"GatewayUniqueId"`
	DomainId        requests.Integer `position:"Query" name:"DomainId"`
	CertIdentifier  string           `position:"Query" name:"CertIdentifier"`
	AcceptLanguage  string           `position:"Query" name:"AcceptLanguage"`
}

// AddSSLCertResponse is the response struct for api AddSSLCert
type AddSSLCertResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string `json:"Message" xml:"Message"`
	Code           int    `json:"Code" xml:"Code"`
	Success        bool   `json:"Success" xml:"Success"`
	Data           bool   `json:"Data" xml:"Data"`
}

// CreateAddSSLCertRequest creates a request to invoke AddSSLCert API
func CreateAddSSLCertRequest() (request *AddSSLCertRequest) {
	request = &AddSSLCertRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "AddSSLCert", "", "")
	request.Method = requests.GET
	return
}

// CreateAddSSLCertResponse creates a response to parse from AddSSLCert response
func CreateAddSSLCertResponse() (response *AddSSLCertResponse) {
	response = &AddSSLCertResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
