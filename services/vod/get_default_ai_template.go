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

// GetDefaultAITemplate invokes the vod.GetDefaultAITemplate API synchronously
func (client *Client) GetDefaultAITemplate(request *GetDefaultAITemplateRequest) (response *GetDefaultAITemplateResponse, err error) {
	response = CreateGetDefaultAITemplateResponse()
	err = client.DoAction(request, response)
	return
}

// GetDefaultAITemplateWithChan invokes the vod.GetDefaultAITemplate API asynchronously
func (client *Client) GetDefaultAITemplateWithChan(request *GetDefaultAITemplateRequest) (<-chan *GetDefaultAITemplateResponse, <-chan error) {
	responseChan := make(chan *GetDefaultAITemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDefaultAITemplate(request)
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

// GetDefaultAITemplateWithCallback invokes the vod.GetDefaultAITemplate API asynchronously
func (client *Client) GetDefaultAITemplateWithCallback(request *GetDefaultAITemplateRequest, callback func(response *GetDefaultAITemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDefaultAITemplateResponse
		var err error
		defer close(result)
		response, err = client.GetDefaultAITemplate(request)
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

// GetDefaultAITemplateRequest is the request struct for api GetDefaultAITemplate
type GetDefaultAITemplateRequest struct {
	*requests.RpcRequest
	TemplateType string `position:"Query" name:"TemplateType"`
}

// GetDefaultAITemplateResponse is the response struct for api GetDefaultAITemplate
type GetDefaultAITemplateResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	TemplateInfo TemplateInfo `json:"TemplateInfo" xml:"TemplateInfo"`
}

// CreateGetDefaultAITemplateRequest creates a request to invoke GetDefaultAITemplate API
func CreateGetDefaultAITemplateRequest() (request *GetDefaultAITemplateRequest) {
	request = &GetDefaultAITemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "GetDefaultAITemplate", "", "")
	request.Method = requests.POST
	return
}

// CreateGetDefaultAITemplateResponse creates a response to parse from GetDefaultAITemplate response
func CreateGetDefaultAITemplateResponse() (response *GetDefaultAITemplateResponse) {
	response = &GetDefaultAITemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
