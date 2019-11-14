package qualitycheck

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

// GetAccAsrResult invokes the qualitycheck.GetAccAsrResult API synchronously
// api document: https://help.aliyun.com/api/qualitycheck/getaccasrresult.html
func (client *Client) GetAccAsrResult(request *GetAccAsrResultRequest) (response *GetAccAsrResultResponse, err error) {
	response = CreateGetAccAsrResultResponse()
	err = client.DoAction(request, response)
	return
}

// GetAccAsrResultWithChan invokes the qualitycheck.GetAccAsrResult API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/getaccasrresult.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetAccAsrResultWithChan(request *GetAccAsrResultRequest) (<-chan *GetAccAsrResultResponse, <-chan error) {
	responseChan := make(chan *GetAccAsrResultResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAccAsrResult(request)
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

// GetAccAsrResultWithCallback invokes the qualitycheck.GetAccAsrResult API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/getaccasrresult.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetAccAsrResultWithCallback(request *GetAccAsrResultRequest, callback func(response *GetAccAsrResultResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAccAsrResultResponse
		var err error
		defer close(result)
		response, err = client.GetAccAsrResult(request)
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

// GetAccAsrResultRequest is the request struct for api GetAccAsrResult
type GetAccAsrResultRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
}

// GetAccAsrResultResponse is the response struct for api GetAccAsrResult
type GetAccAsrResultResponse struct {
	*responses.BaseResponse
	RequestId string                `json:"RequestId" xml:"RequestId"`
	Success   bool                  `json:"Success" xml:"Success"`
	Code      string                `json:"Code" xml:"Code"`
	Message   string                `json:"Message" xml:"Message"`
	Count     int                   `json:"Count" xml:"Count"`
	Data      DataInGetAccAsrResult `json:"Data" xml:"Data"`
}

// CreateGetAccAsrResultRequest creates a request to invoke GetAccAsrResult API
func CreateGetAccAsrResultRequest() (request *GetAccAsrResultRequest) {
	request = &GetAccAsrResultRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "GetAccAsrResult", "", "")
	return
}

// CreateGetAccAsrResultResponse creates a response to parse from GetAccAsrResult response
func CreateGetAccAsrResultResponse() (response *GetAccAsrResultResponse) {
	response = &GetAccAsrResultResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}