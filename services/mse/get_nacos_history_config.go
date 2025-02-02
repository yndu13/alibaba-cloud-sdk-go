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

// GetNacosHistoryConfig invokes the mse.GetNacosHistoryConfig API synchronously
func (client *Client) GetNacosHistoryConfig(request *GetNacosHistoryConfigRequest) (response *GetNacosHistoryConfigResponse, err error) {
	response = CreateGetNacosHistoryConfigResponse()
	err = client.DoAction(request, response)
	return
}

// GetNacosHistoryConfigWithChan invokes the mse.GetNacosHistoryConfig API asynchronously
func (client *Client) GetNacosHistoryConfigWithChan(request *GetNacosHistoryConfigRequest) (<-chan *GetNacosHistoryConfigResponse, <-chan error) {
	responseChan := make(chan *GetNacosHistoryConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetNacosHistoryConfig(request)
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

// GetNacosHistoryConfigWithCallback invokes the mse.GetNacosHistoryConfig API asynchronously
func (client *Client) GetNacosHistoryConfigWithCallback(request *GetNacosHistoryConfigRequest, callback func(response *GetNacosHistoryConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetNacosHistoryConfigResponse
		var err error
		defer close(result)
		response, err = client.GetNacosHistoryConfig(request)
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

// GetNacosHistoryConfigRequest is the request struct for api GetNacosHistoryConfig
type GetNacosHistoryConfigRequest struct {
	*requests.RpcRequest
	Nid            string `position:"Query" name:"Nid"`
	InstanceId     string `position:"Query" name:"InstanceId"`
	DataId         string `position:"Query" name:"DataId"`
	NamespaceId    string `position:"Query" name:"NamespaceId"`
	AcceptLanguage string `position:"Query" name:"AcceptLanguage"`
	Group          string `position:"Query" name:"Group"`
}

// GetNacosHistoryConfigResponse is the response struct for api GetNacosHistoryConfig
type GetNacosHistoryConfigResponse struct {
	*responses.BaseResponse
	Message       string        `json:"Message" xml:"Message"`
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	ErrorCode     string        `json:"ErrorCode" xml:"ErrorCode"`
	Success       bool          `json:"Success" xml:"Success"`
	Configuration Configuration `json:"Configuration" xml:"Configuration"`
}

// CreateGetNacosHistoryConfigRequest creates a request to invoke GetNacosHistoryConfig API
func CreateGetNacosHistoryConfigRequest() (request *GetNacosHistoryConfigRequest) {
	request = &GetNacosHistoryConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "GetNacosHistoryConfig", "", "")
	request.Method = requests.POST
	return
}

// CreateGetNacosHistoryConfigResponse creates a response to parse from GetNacosHistoryConfig response
func CreateGetNacosHistoryConfigResponse() (response *GetNacosHistoryConfigResponse) {
	response = &GetNacosHistoryConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
