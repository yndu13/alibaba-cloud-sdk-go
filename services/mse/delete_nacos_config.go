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

// DeleteNacosConfig invokes the mse.DeleteNacosConfig API synchronously
func (client *Client) DeleteNacosConfig(request *DeleteNacosConfigRequest) (response *DeleteNacosConfigResponse, err error) {
	response = CreateDeleteNacosConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteNacosConfigWithChan invokes the mse.DeleteNacosConfig API asynchronously
func (client *Client) DeleteNacosConfigWithChan(request *DeleteNacosConfigRequest) (<-chan *DeleteNacosConfigResponse, <-chan error) {
	responseChan := make(chan *DeleteNacosConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteNacosConfig(request)
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

// DeleteNacosConfigWithCallback invokes the mse.DeleteNacosConfig API asynchronously
func (client *Client) DeleteNacosConfigWithCallback(request *DeleteNacosConfigRequest, callback func(response *DeleteNacosConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteNacosConfigResponse
		var err error
		defer close(result)
		response, err = client.DeleteNacosConfig(request)
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

// DeleteNacosConfigRequest is the request struct for api DeleteNacosConfig
type DeleteNacosConfigRequest struct {
	*requests.RpcRequest
	InstanceId     string           `position:"Query" name:"InstanceId"`
	DataId         string           `position:"Query" name:"DataId"`
	NamespaceId    string           `position:"Query" name:"NamespaceId"`
	AcceptLanguage string           `position:"Query" name:"AcceptLanguage"`
	Beta           requests.Boolean `position:"Query" name:"Beta"`
	Group          string           `position:"Query" name:"Group"`
}

// DeleteNacosConfigResponse is the response struct for api DeleteNacosConfig
type DeleteNacosConfigResponse struct {
	*responses.BaseResponse
	HttpCode  string `json:"HttpCode" xml:"HttpCode"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateDeleteNacosConfigRequest creates a request to invoke DeleteNacosConfig API
func CreateDeleteNacosConfigRequest() (request *DeleteNacosConfigRequest) {
	request = &DeleteNacosConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "DeleteNacosConfig", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteNacosConfigResponse creates a response to parse from DeleteNacosConfig response
func CreateDeleteNacosConfigResponse() (response *DeleteNacosConfigResponse) {
	response = &DeleteNacosConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
