package dbs

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

// ModifyStorageStrategy invokes the dbs.ModifyStorageStrategy API synchronously
func (client *Client) ModifyStorageStrategy(request *ModifyStorageStrategyRequest) (response *ModifyStorageStrategyResponse, err error) {
	response = CreateModifyStorageStrategyResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyStorageStrategyWithChan invokes the dbs.ModifyStorageStrategy API asynchronously
func (client *Client) ModifyStorageStrategyWithChan(request *ModifyStorageStrategyRequest) (<-chan *ModifyStorageStrategyResponse, <-chan error) {
	responseChan := make(chan *ModifyStorageStrategyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyStorageStrategy(request)
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

// ModifyStorageStrategyWithCallback invokes the dbs.ModifyStorageStrategy API asynchronously
func (client *Client) ModifyStorageStrategyWithCallback(request *ModifyStorageStrategyRequest, callback func(response *ModifyStorageStrategyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyStorageStrategyResponse
		var err error
		defer close(result)
		response, err = client.ModifyStorageStrategy(request)
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

// ModifyStorageStrategyRequest is the request struct for api ModifyStorageStrategy
type ModifyStorageStrategyRequest struct {
	*requests.RpcRequest
	DuplicationArchivePeriod          requests.Integer `position:"Query" name:"DuplicationArchivePeriod"`
	ClientToken                       string           `position:"Query" name:"ClientToken"`
	BackupPlanId                      string           `position:"Query" name:"BackupPlanId"`
	OwnerId                           string           `position:"Query" name:"OwnerId"`
	BackupRetentionPeriod             requests.Integer `position:"Query" name:"BackupRetentionPeriod"`
	DuplicationInfrequentAccessPeriod requests.Integer `position:"Query" name:"DuplicationInfrequentAccessPeriod"`
}

// ModifyStorageStrategyResponse is the response struct for api ModifyStorageStrategy
type ModifyStorageStrategyResponse struct {
	*responses.BaseResponse
	Success        bool   `json:"Success" xml:"Success"`
	ErrCode        string `json:"ErrCode" xml:"ErrCode"`
	ErrMessage     string `json:"ErrMessage" xml:"ErrMessage"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	BackupPlanId   string `json:"BackupPlanId" xml:"BackupPlanId"`
	NeedPrecheck   bool   `json:"NeedPrecheck" xml:"NeedPrecheck"`
}

// CreateModifyStorageStrategyRequest creates a request to invoke ModifyStorageStrategy API
func CreateModifyStorageStrategyRequest() (request *ModifyStorageStrategyRequest) {
	request = &ModifyStorageStrategyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dbs", "2019-03-06", "ModifyStorageStrategy", "cbs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyStorageStrategyResponse creates a response to parse from ModifyStorageStrategy response
func CreateModifyStorageStrategyResponse() (response *ModifyStorageStrategyResponse) {
	response = &ModifyStorageStrategyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}