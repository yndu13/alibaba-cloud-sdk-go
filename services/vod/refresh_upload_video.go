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

// RefreshUploadVideo invokes the vod.RefreshUploadVideo API synchronously
func (client *Client) RefreshUploadVideo(request *RefreshUploadVideoRequest) (response *RefreshUploadVideoResponse, err error) {
	response = CreateRefreshUploadVideoResponse()
	err = client.DoAction(request, response)
	return
}

// RefreshUploadVideoWithChan invokes the vod.RefreshUploadVideo API asynchronously
func (client *Client) RefreshUploadVideoWithChan(request *RefreshUploadVideoRequest) (<-chan *RefreshUploadVideoResponse, <-chan error) {
	responseChan := make(chan *RefreshUploadVideoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RefreshUploadVideo(request)
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

// RefreshUploadVideoWithCallback invokes the vod.RefreshUploadVideo API asynchronously
func (client *Client) RefreshUploadVideoWithCallback(request *RefreshUploadVideoRequest, callback func(response *RefreshUploadVideoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RefreshUploadVideoResponse
		var err error
		defer close(result)
		response, err = client.RefreshUploadVideo(request)
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

// RefreshUploadVideoRequest is the request struct for api RefreshUploadVideo
type RefreshUploadVideoRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	VideoId              string           `position:"Query" name:"VideoId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// RefreshUploadVideoResponse is the response struct for api RefreshUploadVideo
type RefreshUploadVideoResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	UploadAuth    string `json:"UploadAuth" xml:"UploadAuth"`
	UploadAddress string `json:"UploadAddress" xml:"UploadAddress"`
	VideoId       string `json:"VideoId" xml:"VideoId"`
}

// CreateRefreshUploadVideoRequest creates a request to invoke RefreshUploadVideo API
func CreateRefreshUploadVideoRequest() (request *RefreshUploadVideoRequest) {
	request = &RefreshUploadVideoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "RefreshUploadVideo", "", "")
	request.Method = requests.POST
	return
}

// CreateRefreshUploadVideoResponse creates a response to parse from RefreshUploadVideo response
func CreateRefreshUploadVideoResponse() (response *RefreshUploadVideoResponse) {
	response = &RefreshUploadVideoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
