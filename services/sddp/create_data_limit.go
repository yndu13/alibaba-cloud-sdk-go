package sddp

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

// CreateDataLimit invokes the sddp.CreateDataLimit API synchronously
func (client *Client) CreateDataLimit(request *CreateDataLimitRequest) (response *CreateDataLimitResponse, err error) {
	response = CreateCreateDataLimitResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDataLimitWithChan invokes the sddp.CreateDataLimit API asynchronously
func (client *Client) CreateDataLimitWithChan(request *CreateDataLimitRequest) (<-chan *CreateDataLimitResponse, <-chan error) {
	responseChan := make(chan *CreateDataLimitResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDataLimit(request)
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

// CreateDataLimitWithCallback invokes the sddp.CreateDataLimit API asynchronously
func (client *Client) CreateDataLimitWithCallback(request *CreateDataLimitRequest, callback func(response *CreateDataLimitResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDataLimitResponse
		var err error
		defer close(result)
		response, err = client.CreateDataLimit(request)
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

// CreateDataLimitRequest is the request struct for api CreateDataLimit
type CreateDataLimitRequest struct {
	*requests.RpcRequest
	OcrStatus       requests.Integer `position:"Query" name:"OcrStatus"`
	ParentId        string           `position:"Query" name:"ParentId"`
	Password        string           `position:"Query" name:"Password"`
	BatchCreate     requests.Boolean `position:"Query" name:"BatchCreate"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	DataLimitList   string           `position:"Query" name:"DataLimitList"`
	Lang            string           `position:"Query" name:"Lang"`
	ServiceRegionId string           `position:"Query" name:"ServiceRegionId"`
	EngineType      string           `position:"Query" name:"EngineType"`
	AuditStatus     requests.Integer `position:"Query" name:"AuditStatus"`
	AutoScan        requests.Integer `position:"Query" name:"AutoScan"`
	LogStoreDay     requests.Integer `position:"Query" name:"LogStoreDay"`
	ResourceType    requests.Integer `position:"Query" name:"ResourceType"`
	Connector       string           `position:"Query" name:"Connector"`
	Port            requests.Integer `position:"Query" name:"Port"`
	EventStatus     requests.Integer `position:"Query" name:"EventStatus"`
	UserName        string           `position:"Query" name:"UserName"`
}

// CreateDataLimitResponse is the response struct for api CreateDataLimit
type CreateDataLimitResponse struct {
	*responses.BaseResponse
	Id        int    `json:"Id" xml:"Id"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateDataLimitRequest creates a request to invoke CreateDataLimit API
func CreateCreateDataLimitRequest() (request *CreateDataLimitRequest) {
	request = &CreateDataLimitRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sddp", "2019-01-03", "CreateDataLimit", "sddp", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateDataLimitResponse creates a response to parse from CreateDataLimit response
func CreateCreateDataLimitResponse() (response *CreateDataLimitResponse) {
	response = &CreateDataLimitResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
