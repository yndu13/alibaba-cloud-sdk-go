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

// UpdateEngineNamespace invokes the mse.UpdateEngineNamespace API synchronously
func (client *Client) UpdateEngineNamespace(request *UpdateEngineNamespaceRequest) (response *UpdateEngineNamespaceResponse, err error) {
	response = CreateUpdateEngineNamespaceResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateEngineNamespaceWithChan invokes the mse.UpdateEngineNamespace API asynchronously
func (client *Client) UpdateEngineNamespaceWithChan(request *UpdateEngineNamespaceRequest) (<-chan *UpdateEngineNamespaceResponse, <-chan error) {
	responseChan := make(chan *UpdateEngineNamespaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateEngineNamespace(request)
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

// UpdateEngineNamespaceWithCallback invokes the mse.UpdateEngineNamespace API asynchronously
func (client *Client) UpdateEngineNamespaceWithCallback(request *UpdateEngineNamespaceRequest, callback func(response *UpdateEngineNamespaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateEngineNamespaceResponse
		var err error
		defer close(result)
		response, err = client.UpdateEngineNamespace(request)
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

// UpdateEngineNamespaceRequest is the request struct for api UpdateEngineNamespace
type UpdateEngineNamespaceRequest struct {
	*requests.RpcRequest
	ClusterId      string           `position:"Query" name:"ClusterId"`
	InstanceId     string           `position:"Query" name:"InstanceId"`
	ServiceCount   requests.Integer `position:"Query" name:"ServiceCount"`
	Name           string           `position:"Query" name:"Name"`
	AcceptLanguage string           `position:"Query" name:"AcceptLanguage"`
	Id             string           `position:"Query" name:"Id"`
	Desc           string           `position:"Query" name:"Desc"`
}

// UpdateEngineNamespaceResponse is the response struct for api UpdateEngineNamespace
type UpdateEngineNamespaceResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateUpdateEngineNamespaceRequest creates a request to invoke UpdateEngineNamespace API
func CreateUpdateEngineNamespaceRequest() (request *UpdateEngineNamespaceRequest) {
	request = &UpdateEngineNamespaceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "UpdateEngineNamespace", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateEngineNamespaceResponse creates a response to parse from UpdateEngineNamespace response
func CreateUpdateEngineNamespaceResponse() (response *UpdateEngineNamespaceResponse) {
	response = &UpdateEngineNamespaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
