package live

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

// DescribeLiveTagResources invokes the live.DescribeLiveTagResources API synchronously
// api document: https://help.aliyun.com/api/live/describelivetagresources.html
func (client *Client) DescribeLiveTagResources(request *DescribeLiveTagResourcesRequest) (response *DescribeLiveTagResourcesResponse, err error) {
	response = CreateDescribeLiveTagResourcesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLiveTagResourcesWithChan invokes the live.DescribeLiveTagResources API asynchronously
// api document: https://help.aliyun.com/api/live/describelivetagresources.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLiveTagResourcesWithChan(request *DescribeLiveTagResourcesRequest) (<-chan *DescribeLiveTagResourcesResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveTagResourcesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveTagResources(request)
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

// DescribeLiveTagResourcesWithCallback invokes the live.DescribeLiveTagResources API asynchronously
// api document: https://help.aliyun.com/api/live/describelivetagresources.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLiveTagResourcesWithCallback(request *DescribeLiveTagResourcesRequest, callback func(response *DescribeLiveTagResourcesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveTagResourcesResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveTagResources(request)
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

// DescribeLiveTagResourcesRequest is the request struct for api DescribeLiveTagResources
type DescribeLiveTagResourcesRequest struct {
	*requests.RpcRequest
	Scope        string                         `position:"Query" name:"Scope"`
	Tag          *[]DescribeLiveTagResourcesTag `position:"Query" name:"Tag"  type:"Repeated"`
	ResourceId   *[]string                      `position:"Query" name:"ResourceId"  type:"Repeated"`
	OwnerId      requests.Integer               `position:"Query" name:"OwnerId"`
	ResourceType string                         `position:"Query" name:"ResourceType"`
}

// DescribeLiveTagResourcesTag is a repeated param struct in DescribeLiveTagResourcesRequest
type DescribeLiveTagResourcesTag struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

// DescribeLiveTagResourcesResponse is the response struct for api DescribeLiveTagResources
type DescribeLiveTagResourcesResponse struct {
	*responses.BaseResponse
	RequestId    string        `json:"RequestId" xml:"RequestId"`
	TagResources []TagResource `json:"TagResources" xml:"TagResources"`
}

// CreateDescribeLiveTagResourcesRequest creates a request to invoke DescribeLiveTagResources API
func CreateDescribeLiveTagResourcesRequest() (request *DescribeLiveTagResourcesRequest) {
	request = &DescribeLiveTagResourcesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeLiveTagResources", "live", "openAPI")
	return
}

// CreateDescribeLiveTagResourcesResponse creates a response to parse from DescribeLiveTagResources response
func CreateDescribeLiveTagResourcesResponse() (response *DescribeLiveTagResourcesResponse) {
	response = &DescribeLiveTagResourcesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}