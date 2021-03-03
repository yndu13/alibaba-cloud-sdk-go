package linkwan

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

// ListNodeGroupTransferFlowStats invokes the linkwan.ListNodeGroupTransferFlowStats API synchronously
func (client *Client) ListNodeGroupTransferFlowStats(request *ListNodeGroupTransferFlowStatsRequest) (response *ListNodeGroupTransferFlowStatsResponse, err error) {
	response = CreateListNodeGroupTransferFlowStatsResponse()
	err = client.DoAction(request, response)
	return
}

// ListNodeGroupTransferFlowStatsWithChan invokes the linkwan.ListNodeGroupTransferFlowStats API asynchronously
func (client *Client) ListNodeGroupTransferFlowStatsWithChan(request *ListNodeGroupTransferFlowStatsRequest) (<-chan *ListNodeGroupTransferFlowStatsResponse, <-chan error) {
	responseChan := make(chan *ListNodeGroupTransferFlowStatsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListNodeGroupTransferFlowStats(request)
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

// ListNodeGroupTransferFlowStatsWithCallback invokes the linkwan.ListNodeGroupTransferFlowStats API asynchronously
func (client *Client) ListNodeGroupTransferFlowStatsWithCallback(request *ListNodeGroupTransferFlowStatsRequest, callback func(response *ListNodeGroupTransferFlowStatsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListNodeGroupTransferFlowStatsResponse
		var err error
		defer close(result)
		response, err = client.ListNodeGroupTransferFlowStats(request)
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

// ListNodeGroupTransferFlowStatsRequest is the request struct for api ListNodeGroupTransferFlowStats
type ListNodeGroupTransferFlowStatsRequest struct {
	*requests.RpcRequest
	EndMillis        requests.Integer `position:"Query" name:"EndMillis"`
	IotInstanceId    string           `position:"Query" name:"IotInstanceId"`
	TimeIntervalUnit string           `position:"Query" name:"TimeIntervalUnit"`
	NodeGroupId      string           `position:"Query" name:"NodeGroupId"`
	ApiProduct       string           `position:"Body" name:"ApiProduct"`
	ApiRevision      string           `position:"Body" name:"ApiRevision"`
	BeginMillis      requests.Integer `position:"Query" name:"BeginMillis"`
}

// ListNodeGroupTransferFlowStatsResponse is the response struct for api ListNodeGroupTransferFlowStats
type ListNodeGroupTransferFlowStatsResponse struct {
	*responses.BaseResponse
	RequestId string                                     `json:"RequestId" xml:"RequestId"`
	Success   bool                                       `json:"Success" xml:"Success"`
	Data      []FlowStatInListNodeGroupTransferFlowStats `json:"Data" xml:"Data"`
}

// CreateListNodeGroupTransferFlowStatsRequest creates a request to invoke ListNodeGroupTransferFlowStats API
func CreateListNodeGroupTransferFlowStatsRequest() (request *ListNodeGroupTransferFlowStatsRequest) {
	request = &ListNodeGroupTransferFlowStatsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("LinkWAN", "2019-03-01", "ListNodeGroupTransferFlowStats", "linkwan", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListNodeGroupTransferFlowStatsResponse creates a response to parse from ListNodeGroupTransferFlowStats response
func CreateListNodeGroupTransferFlowStatsResponse() (response *ListNodeGroupTransferFlowStatsResponse) {
	response = &ListNodeGroupTransferFlowStatsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}