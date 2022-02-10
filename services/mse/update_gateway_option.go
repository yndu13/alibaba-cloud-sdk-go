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

// UpdateGatewayOption invokes the mse.UpdateGatewayOption API synchronously
func (client *Client) UpdateGatewayOption(request *UpdateGatewayOptionRequest) (response *UpdateGatewayOptionResponse, err error) {
	response = CreateUpdateGatewayOptionResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateGatewayOptionWithChan invokes the mse.UpdateGatewayOption API asynchronously
func (client *Client) UpdateGatewayOptionWithChan(request *UpdateGatewayOptionRequest) (<-chan *UpdateGatewayOptionResponse, <-chan error) {
	responseChan := make(chan *UpdateGatewayOptionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateGatewayOption(request)
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

// UpdateGatewayOptionWithCallback invokes the mse.UpdateGatewayOption API asynchronously
func (client *Client) UpdateGatewayOptionWithCallback(request *UpdateGatewayOptionRequest, callback func(response *UpdateGatewayOptionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateGatewayOptionResponse
		var err error
		defer close(result)
		response, err = client.UpdateGatewayOption(request)
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

// UpdateGatewayOptionRequest is the request struct for api UpdateGatewayOption
type UpdateGatewayOptionRequest struct {
	*requests.RpcRequest
	GatewayUniqueId string                           `position:"Query" name:"GatewayUniqueId"`
	GatewayOption   UpdateGatewayOptionGatewayOption `position:"Query" name:"GatewayOption"  type:"Struct"`
	GatewayId       requests.Integer                 `position:"Query" name:"GatewayId"`
	AcceptLanguage  string                           `position:"Query" name:"AcceptLanguage"`
}

// UpdateGatewayOptionGatewayOption is a repeated param struct in UpdateGatewayOptionRequest
type UpdateGatewayOptionGatewayOption struct {
	EnableHardwareAcceleration string                                           `name:"EnableHardwareAcceleration"`
	DisableHttp2Alpn           string                                           `name:"DisableHttp2Alpn"`
	LogConfigDetails           UpdateGatewayOptionGatewayOptionLogConfigDetails `name:"LogConfigDetails" type:"Struct"`
	TraceDetails               UpdateGatewayOptionGatewayOptionTraceDetails     `name:"TraceDetails" type:"Struct"`
}

// UpdateGatewayOptionGatewayOptionLogConfigDetails is a repeated param struct in UpdateGatewayOptionRequest
type UpdateGatewayOptionGatewayOptionLogConfigDetails struct {
	ProjectName  string `name:"ProjectName"`
	LogStoreName string `name:"LogStoreName"`
	LogEnabled   string `name:"LogEnabled"`
}

// UpdateGatewayOptionGatewayOptionTraceDetails is a repeated param struct in UpdateGatewayOptionRequest
type UpdateGatewayOptionGatewayOptionTraceDetails struct {
	TraceEnabled string `name:"TraceEnabled"`
	Sample       string `name:"Sample"`
}

// UpdateGatewayOptionResponse is the response struct for api UpdateGatewayOption
type UpdateGatewayOptionResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string `json:"Message" xml:"Message"`
	Code           int    `json:"Code" xml:"Code"`
	Success        bool   `json:"Success" xml:"Success"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateUpdateGatewayOptionRequest creates a request to invoke UpdateGatewayOption API
func CreateUpdateGatewayOptionRequest() (request *UpdateGatewayOptionRequest) {
	request = &UpdateGatewayOptionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "UpdateGatewayOption", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateGatewayOptionResponse creates a response to parse from UpdateGatewayOption response
func CreateUpdateGatewayOptionResponse() (response *UpdateGatewayOptionResponse) {
	response = &UpdateGatewayOptionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
