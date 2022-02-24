package eipanycast

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

// UnassociateAnycastEipAddress invokes the eipanycast.UnassociateAnycastEipAddress API synchronously
func (client *Client) UnassociateAnycastEipAddress(request *UnassociateAnycastEipAddressRequest) (response *UnassociateAnycastEipAddressResponse, err error) {
	response = CreateUnassociateAnycastEipAddressResponse()
	err = client.DoAction(request, response)
	return
}

// UnassociateAnycastEipAddressWithChan invokes the eipanycast.UnassociateAnycastEipAddress API asynchronously
func (client *Client) UnassociateAnycastEipAddressWithChan(request *UnassociateAnycastEipAddressRequest) (<-chan *UnassociateAnycastEipAddressResponse, <-chan error) {
	responseChan := make(chan *UnassociateAnycastEipAddressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnassociateAnycastEipAddress(request)
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

// UnassociateAnycastEipAddressWithCallback invokes the eipanycast.UnassociateAnycastEipAddress API asynchronously
func (client *Client) UnassociateAnycastEipAddressWithCallback(request *UnassociateAnycastEipAddressRequest, callback func(response *UnassociateAnycastEipAddressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnassociateAnycastEipAddressResponse
		var err error
		defer close(result)
		response, err = client.UnassociateAnycastEipAddress(request)
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

// UnassociateAnycastEipAddressRequest is the request struct for api UnassociateAnycastEipAddress
type UnassociateAnycastEipAddressRequest struct {
	*requests.RpcRequest
	DryRun               string `position:"Query" name:"DryRun"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	BindInstanceType     string `position:"Query" name:"BindInstanceType"`
	BindInstanceRegionId string `position:"Query" name:"BindInstanceRegionId"`
	PrivateIpAddress     string `position:"Query" name:"PrivateIpAddress"`
	AnycastId            string `position:"Query" name:"AnycastId"`
	BindInstanceId       string `position:"Query" name:"BindInstanceId"`
}

// UnassociateAnycastEipAddressResponse is the response struct for api UnassociateAnycastEipAddress
type UnassociateAnycastEipAddressResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUnassociateAnycastEipAddressRequest creates a request to invoke UnassociateAnycastEipAddress API
func CreateUnassociateAnycastEipAddressRequest() (request *UnassociateAnycastEipAddressRequest) {
	request = &UnassociateAnycastEipAddressRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Eipanycast", "2020-03-09", "UnassociateAnycastEipAddress", "eipanycast", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUnassociateAnycastEipAddressResponse creates a response to parse from UnassociateAnycastEipAddress response
func CreateUnassociateAnycastEipAddressResponse() (response *UnassociateAnycastEipAddressResponse) {
	response = &UnassociateAnycastEipAddressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
