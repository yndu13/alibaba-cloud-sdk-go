package scsp

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

// GetTicketTemplateSchema invokes the scsp.GetTicketTemplateSchema API synchronously
func (client *Client) GetTicketTemplateSchema(request *GetTicketTemplateSchemaRequest) (response *GetTicketTemplateSchemaResponse, err error) {
	response = CreateGetTicketTemplateSchemaResponse()
	err = client.DoAction(request, response)
	return
}

// GetTicketTemplateSchemaWithChan invokes the scsp.GetTicketTemplateSchema API asynchronously
func (client *Client) GetTicketTemplateSchemaWithChan(request *GetTicketTemplateSchemaRequest) (<-chan *GetTicketTemplateSchemaResponse, <-chan error) {
	responseChan := make(chan *GetTicketTemplateSchemaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetTicketTemplateSchema(request)
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

// GetTicketTemplateSchemaWithCallback invokes the scsp.GetTicketTemplateSchema API asynchronously
func (client *Client) GetTicketTemplateSchemaWithCallback(request *GetTicketTemplateSchemaRequest, callback func(response *GetTicketTemplateSchemaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetTicketTemplateSchemaResponse
		var err error
		defer close(result)
		response, err = client.GetTicketTemplateSchema(request)
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

// GetTicketTemplateSchemaRequest is the request struct for api GetTicketTemplateSchema
type GetTicketTemplateSchemaRequest struct {
	*requests.RpcRequest
	ClientToken string           `position:"Query"`
	InstanceId  string           `position:"Query"`
	TemplateId  requests.Integer `position:"Query"`
}

// GetTicketTemplateSchemaResponse is the response struct for api GetTicketTemplateSchema
type GetTicketTemplateSchemaResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateGetTicketTemplateSchemaRequest creates a request to invoke GetTicketTemplateSchema API
func CreateGetTicketTemplateSchemaRequest() (request *GetTicketTemplateSchemaRequest) {
	request = &GetTicketTemplateSchemaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scsp", "2020-07-02", "GetTicketTemplateSchema", "", "")
	request.Method = requests.GET
	return
}

// CreateGetTicketTemplateSchemaResponse creates a response to parse from GetTicketTemplateSchema response
func CreateGetTicketTemplateSchemaResponse() (response *GetTicketTemplateSchemaResponse) {
	response = &GetTicketTemplateSchemaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}