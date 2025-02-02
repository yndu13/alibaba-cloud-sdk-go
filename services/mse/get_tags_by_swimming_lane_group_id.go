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

// GetTagsBySwimmingLaneGroupId invokes the mse.GetTagsBySwimmingLaneGroupId API synchronously
func (client *Client) GetTagsBySwimmingLaneGroupId(request *GetTagsBySwimmingLaneGroupIdRequest) (response *GetTagsBySwimmingLaneGroupIdResponse, err error) {
	response = CreateGetTagsBySwimmingLaneGroupIdResponse()
	err = client.DoAction(request, response)
	return
}

// GetTagsBySwimmingLaneGroupIdWithChan invokes the mse.GetTagsBySwimmingLaneGroupId API asynchronously
func (client *Client) GetTagsBySwimmingLaneGroupIdWithChan(request *GetTagsBySwimmingLaneGroupIdRequest) (<-chan *GetTagsBySwimmingLaneGroupIdResponse, <-chan error) {
	responseChan := make(chan *GetTagsBySwimmingLaneGroupIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetTagsBySwimmingLaneGroupId(request)
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

// GetTagsBySwimmingLaneGroupIdWithCallback invokes the mse.GetTagsBySwimmingLaneGroupId API asynchronously
func (client *Client) GetTagsBySwimmingLaneGroupIdWithCallback(request *GetTagsBySwimmingLaneGroupIdRequest, callback func(response *GetTagsBySwimmingLaneGroupIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetTagsBySwimmingLaneGroupIdResponse
		var err error
		defer close(result)
		response, err = client.GetTagsBySwimmingLaneGroupId(request)
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

// GetTagsBySwimmingLaneGroupIdRequest is the request struct for api GetTagsBySwimmingLaneGroupId
type GetTagsBySwimmingLaneGroupIdRequest struct {
	*requests.RpcRequest
	GroupId        requests.Integer `position:"Query" name:"GroupId"`
	AcceptLanguage string           `position:"Query" name:"AcceptLanguage"`
}

// GetTagsBySwimmingLaneGroupIdResponse is the response struct for api GetTagsBySwimmingLaneGroupId
type GetTagsBySwimmingLaneGroupIdResponse struct {
	*responses.BaseResponse
}

// CreateGetTagsBySwimmingLaneGroupIdRequest creates a request to invoke GetTagsBySwimmingLaneGroupId API
func CreateGetTagsBySwimmingLaneGroupIdRequest() (request *GetTagsBySwimmingLaneGroupIdRequest) {
	request = &GetTagsBySwimmingLaneGroupIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "GetTagsBySwimmingLaneGroupId", "", "")
	request.Method = requests.POST
	return
}

// CreateGetTagsBySwimmingLaneGroupIdResponse creates a response to parse from GetTagsBySwimmingLaneGroupId response
func CreateGetTagsBySwimmingLaneGroupIdResponse() (response *GetTagsBySwimmingLaneGroupIdResponse) {
	response = &GetTagsBySwimmingLaneGroupIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
