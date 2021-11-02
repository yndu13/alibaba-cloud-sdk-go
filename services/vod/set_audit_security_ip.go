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

// SetAuditSecurityIp invokes the vod.SetAuditSecurityIp API synchronously
func (client *Client) SetAuditSecurityIp(request *SetAuditSecurityIpRequest) (response *SetAuditSecurityIpResponse, err error) {
	response = CreateSetAuditSecurityIpResponse()
	err = client.DoAction(request, response)
	return
}

// SetAuditSecurityIpWithChan invokes the vod.SetAuditSecurityIp API asynchronously
func (client *Client) SetAuditSecurityIpWithChan(request *SetAuditSecurityIpRequest) (<-chan *SetAuditSecurityIpResponse, <-chan error) {
	responseChan := make(chan *SetAuditSecurityIpResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetAuditSecurityIp(request)
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

// SetAuditSecurityIpWithCallback invokes the vod.SetAuditSecurityIp API asynchronously
func (client *Client) SetAuditSecurityIpWithCallback(request *SetAuditSecurityIpRequest, callback func(response *SetAuditSecurityIpResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetAuditSecurityIpResponse
		var err error
		defer close(result)
		response, err = client.SetAuditSecurityIp(request)
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

// SetAuditSecurityIpRequest is the request struct for api SetAuditSecurityIp
type SetAuditSecurityIpRequest struct {
	*requests.RpcRequest
	SecurityGroupName string `position:"Query" name:"SecurityGroupName"`
	OperateMode       string `position:"Query" name:"OperateMode"`
	Ips               string `position:"Query" name:"Ips"`
}

// SetAuditSecurityIpResponse is the response struct for api SetAuditSecurityIp
type SetAuditSecurityIpResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetAuditSecurityIpRequest creates a request to invoke SetAuditSecurityIp API
func CreateSetAuditSecurityIpRequest() (request *SetAuditSecurityIpRequest) {
	request = &SetAuditSecurityIpRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "SetAuditSecurityIp", "", "")
	request.Method = requests.POST
	return
}

// CreateSetAuditSecurityIpResponse creates a response to parse from SetAuditSecurityIp response
func CreateSetAuditSecurityIpResponse() (response *SetAuditSecurityIpResponse) {
	response = &SetAuditSecurityIpResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
