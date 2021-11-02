package vod

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

// ListAIImageInfo invokes the vod.ListAIImageInfo API synchronously
func (client *Client) ListAIImageInfo(request *ListAIImageInfoRequest) (response *ListAIImageInfoResponse, err error) {
	response = CreateListAIImageInfoResponse()
	err = client.DoAction(request, response)
	return
}

// ListAIImageInfoWithChan invokes the vod.ListAIImageInfo API asynchronously
func (client *Client) ListAIImageInfoWithChan(request *ListAIImageInfoRequest) (<-chan *ListAIImageInfoResponse, <-chan error) {
	responseChan := make(chan *ListAIImageInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAIImageInfo(request)
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

// ListAIImageInfoWithCallback invokes the vod.ListAIImageInfo API asynchronously
func (client *Client) ListAIImageInfoWithCallback(request *ListAIImageInfoRequest, callback func(response *ListAIImageInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAIImageInfoResponse
		var err error
		defer close(result)
		response, err = client.ListAIImageInfo(request)
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

// ListAIImageInfoRequest is the request struct for api ListAIImageInfo
type ListAIImageInfoRequest struct {
	*requests.RpcRequest
	VideoId string `position:"Query" name:"VideoId"`
}

// ListAIImageInfoResponse is the response struct for api ListAIImageInfo
type ListAIImageInfoResponse struct {
	*responses.BaseResponse
	RequestId       string        `json:"RequestId" xml:"RequestId"`
	AIImageInfoList []AIImageInfo `json:"AIImageInfoList" xml:"AIImageInfoList"`
}

// CreateListAIImageInfoRequest creates a request to invoke ListAIImageInfo API
func CreateListAIImageInfoRequest() (request *ListAIImageInfoRequest) {
	request = &ListAIImageInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "ListAIImageInfo", "", "")
	request.Method = requests.POST
	return
}

// CreateListAIImageInfoResponse creates a response to parse from ListAIImageInfo response
func CreateListAIImageInfoResponse() (response *ListAIImageInfoResponse) {
	response = &ListAIImageInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
