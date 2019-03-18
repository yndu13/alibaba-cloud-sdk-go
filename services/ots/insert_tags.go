package ots

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

// InsertTags invokes the ots.InsertTags API synchronously
// api document: https://help.aliyun.com/api/ots/inserttags.html
func (client *Client) InsertTags(request *InsertTagsRequest) (response *InsertTagsResponse, err error) {
	response = CreateInsertTagsResponse()
	err = client.DoAction(request, response)
	return
}

// InsertTagsWithChan invokes the ots.InsertTags API asynchronously
// api document: https://help.aliyun.com/api/ots/inserttags.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) InsertTagsWithChan(request *InsertTagsRequest) (<-chan *InsertTagsResponse, <-chan error) {
	responseChan := make(chan *InsertTagsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.InsertTags(request)
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

// InsertTagsWithCallback invokes the ots.InsertTags API asynchronously
// api document: https://help.aliyun.com/api/ots/inserttags.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) InsertTagsWithCallback(request *InsertTagsRequest, callback func(response *InsertTagsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *InsertTagsResponse
		var err error
		defer close(result)
		response, err = client.InsertTags(request)
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

// InsertTagsRequest is the request struct for api InsertTags
type InsertTagsRequest struct {
	*requests.RpcRequest
	AccessKeyId     string               `position:"Query" name:"access_key_id"`
	ResourceOwnerId requests.Integer     `position:"Query" name:"ResourceOwnerId"`
	InstanceName    string               `position:"Query" name:"InstanceName"`
	TagInfo         *[]InsertTagsTagInfo `position:"Query" name:"TagInfo"  type:"Repeated"`
}

// InsertTagsTagInfo is a repeated param struct in InsertTagsRequest
type InsertTagsTagInfo struct {
	TagValue string `name:"TagValue"`
	TagKey   string `name:"TagKey"`
}

// InsertTagsResponse is the response struct for api InsertTags
type InsertTagsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateInsertTagsRequest creates a request to invoke InsertTags API
func CreateInsertTagsRequest() (request *InsertTagsRequest) {
	request = &InsertTagsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ots", "2016-06-20", "InsertTags", "ots", "openAPI")
	return
}

// CreateInsertTagsResponse creates a response to parse from InsertTags response
func CreateInsertTagsResponse() (response *InsertTagsResponse) {
	response = &InsertTagsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
