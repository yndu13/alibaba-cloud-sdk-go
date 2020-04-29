package lrg

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

// ChangeStatus invokes the lrg.ChangeStatus API synchronously
// api document: https://help.aliyun.com/api/lrg/changestatus.html
func (client *Client) ChangeStatus(request *ChangeStatusRequest) (response *ChangeStatusResponse, err error) {
	response = CreateChangeStatusResponse()
	err = client.DoAction(request, response)
	return
}

// ChangeStatusWithChan invokes the lrg.ChangeStatus API asynchronously
// api document: https://help.aliyun.com/api/lrg/changestatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ChangeStatusWithChan(request *ChangeStatusRequest) (<-chan *ChangeStatusResponse, <-chan error) {
	responseChan := make(chan *ChangeStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ChangeStatus(request)
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

// ChangeStatusWithCallback invokes the lrg.ChangeStatus API asynchronously
// api document: https://help.aliyun.com/api/lrg/changestatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ChangeStatusWithCallback(request *ChangeStatusRequest, callback func(response *ChangeStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ChangeStatusResponse
		var err error
		defer close(result)
		response, err = client.ChangeStatus(request)
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

// ChangeStatusRequest is the request struct for api ChangeStatus
type ChangeStatusRequest struct {
	*requests.RoaRequest
	ProcessId requests.Integer `position:"Path" name:"processId"`
}

// ChangeStatusResponse is the response struct for api ChangeStatus
type ChangeStatusResponse struct {
	*responses.BaseResponse
	Code    int      `json:"code" xml:"code"`
	Message string   `json:"message" xml:"message"`
	Success bool     `json:"success" xml:"success"`
	Data    []string `json:"data" xml:"data"`
}

// CreateChangeStatusRequest creates a request to invoke ChangeStatus API
func CreateChangeStatusRequest() (request *ChangeStatusRequest) {
	request = &ChangeStatusRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("LRG", "2019-10-10", "ChangeStatus", "/api/v2/process/[processId]/status", "", "")
	request.Method = requests.POST
	return
}

// CreateChangeStatusResponse creates a response to parse from ChangeStatus response
func CreateChangeStatusResponse() (response *ChangeStatusResponse) {
	response = &ChangeStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}