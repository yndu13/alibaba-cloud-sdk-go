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

// UpdateBlackWhiteList invokes the mse.UpdateBlackWhiteList API synchronously
func (client *Client) UpdateBlackWhiteList(request *UpdateBlackWhiteListRequest) (response *UpdateBlackWhiteListResponse, err error) {
	response = CreateUpdateBlackWhiteListResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateBlackWhiteListWithChan invokes the mse.UpdateBlackWhiteList API asynchronously
func (client *Client) UpdateBlackWhiteListWithChan(request *UpdateBlackWhiteListRequest) (<-chan *UpdateBlackWhiteListResponse, <-chan error) {
	responseChan := make(chan *UpdateBlackWhiteListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateBlackWhiteList(request)
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

// UpdateBlackWhiteListWithCallback invokes the mse.UpdateBlackWhiteList API asynchronously
func (client *Client) UpdateBlackWhiteListWithCallback(request *UpdateBlackWhiteListRequest, callback func(response *UpdateBlackWhiteListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateBlackWhiteListResponse
		var err error
		defer close(result)
		response, err = client.UpdateBlackWhiteList(request)
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

// UpdateBlackWhiteListRequest is the request struct for api UpdateBlackWhiteList
type UpdateBlackWhiteListRequest struct {
	*requests.RpcRequest
	GatewayUniqueId string           `position:"Query" name:"GatewayUniqueId"`
	Type            string           `position:"Query" name:"Type"`
	Content         string           `position:"Query" name:"Content"`
	Id              requests.Integer `position:"Query" name:"Id"`
	ResourceType    string           `position:"Query" name:"ResourceType"`
	AcceptLanguage  string           `position:"Query" name:"AcceptLanguage"`
	Status          string           `position:"Query" name:"Status"`
}

// UpdateBlackWhiteListResponse is the response struct for api UpdateBlackWhiteList
type UpdateBlackWhiteListResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string `json:"Message" xml:"Message"`
	Code           int    `json:"Code" xml:"Code"`
	Success        bool   `json:"Success" xml:"Success"`
	Data           int64  `json:"Data" xml:"Data"`
}

// CreateUpdateBlackWhiteListRequest creates a request to invoke UpdateBlackWhiteList API
func CreateUpdateBlackWhiteListRequest() (request *UpdateBlackWhiteListRequest) {
	request = &UpdateBlackWhiteListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "UpdateBlackWhiteList", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateBlackWhiteListResponse creates a response to parse from UpdateBlackWhiteList response
func CreateUpdateBlackWhiteListResponse() (response *UpdateBlackWhiteListResponse) {
	response = &UpdateBlackWhiteListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
