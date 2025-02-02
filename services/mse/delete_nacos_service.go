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

// DeleteNacosService invokes the mse.DeleteNacosService API synchronously
func (client *Client) DeleteNacosService(request *DeleteNacosServiceRequest) (response *DeleteNacosServiceResponse, err error) {
	response = CreateDeleteNacosServiceResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteNacosServiceWithChan invokes the mse.DeleteNacosService API asynchronously
func (client *Client) DeleteNacosServiceWithChan(request *DeleteNacosServiceRequest) (<-chan *DeleteNacosServiceResponse, <-chan error) {
	responseChan := make(chan *DeleteNacosServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteNacosService(request)
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

// DeleteNacosServiceWithCallback invokes the mse.DeleteNacosService API asynchronously
func (client *Client) DeleteNacosServiceWithCallback(request *DeleteNacosServiceRequest, callback func(response *DeleteNacosServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteNacosServiceResponse
		var err error
		defer close(result)
		response, err = client.DeleteNacosService(request)
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

// DeleteNacosServiceRequest is the request struct for api DeleteNacosService
type DeleteNacosServiceRequest struct {
	*requests.RpcRequest
	GroupName      string `position:"Query" name:"GroupName"`
	InstanceId     string `position:"Query" name:"InstanceId"`
	NamespaceId    string `position:"Query" name:"NamespaceId"`
	AcceptLanguage string `position:"Query" name:"AcceptLanguage"`
	ServiceName    string `position:"Query" name:"ServiceName"`
}

// DeleteNacosServiceResponse is the response struct for api DeleteNacosService
type DeleteNacosServiceResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Message        string `json:"Message" xml:"Message"`
	Code           int    `json:"Code" xml:"Code"`
	Success        bool   `json:"Success" xml:"Success"`
	Data           string `json:"Data" xml:"Data"`
}

// CreateDeleteNacosServiceRequest creates a request to invoke DeleteNacosService API
func CreateDeleteNacosServiceRequest() (request *DeleteNacosServiceRequest) {
	request = &DeleteNacosServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "DeleteNacosService", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteNacosServiceResponse creates a response to parse from DeleteNacosService response
func CreateDeleteNacosServiceResponse() (response *DeleteNacosServiceResponse) {
	response = &DeleteNacosServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
