package arms

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

// SendMseIncident invokes the arms.SendMseIncident API synchronously
func (client *Client) SendMseIncident(request *SendMseIncidentRequest) (response *SendMseIncidentResponse, err error) {
	response = CreateSendMseIncidentResponse()
	err = client.DoAction(request, response)
	return
}

// SendMseIncidentWithChan invokes the arms.SendMseIncident API asynchronously
func (client *Client) SendMseIncidentWithChan(request *SendMseIncidentRequest) (<-chan *SendMseIncidentResponse, <-chan error) {
	responseChan := make(chan *SendMseIncidentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SendMseIncident(request)
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

// SendMseIncidentWithCallback invokes the arms.SendMseIncident API asynchronously
func (client *Client) SendMseIncidentWithCallback(request *SendMseIncidentRequest, callback func(response *SendMseIncidentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SendMseIncidentResponse
		var err error
		defer close(result)
		response, err = client.SendMseIncident(request)
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

// SendMseIncidentRequest is the request struct for api SendMseIncident
type SendMseIncidentRequest struct {
	*requests.RpcRequest
	Incidents   string `position:"Query" name:"Incidents"`
	ProxyUserId string `position:"Query" name:"ProxyUserId"`
}

// SendMseIncidentResponse is the response struct for api SendMseIncident
type SendMseIncidentResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateSendMseIncidentRequest creates a request to invoke SendMseIncident API
func CreateSendMseIncidentRequest() (request *SendMseIncidentRequest) {
	request = &SendMseIncidentRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "SendMseIncident", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSendMseIncidentResponse creates a response to parse from SendMseIncident response
func CreateSendMseIncidentResponse() (response *SendMseIncidentResponse) {
	response = &SendMseIncidentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}