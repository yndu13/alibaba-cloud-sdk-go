package cms

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

// DescribeAlertHistoryList invokes the cms.DescribeAlertHistoryList API synchronously
func (client *Client) DescribeAlertHistoryList(request *DescribeAlertHistoryListRequest) (response *DescribeAlertHistoryListResponse, err error) {
	response = CreateDescribeAlertHistoryListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAlertHistoryListWithChan invokes the cms.DescribeAlertHistoryList API asynchronously
func (client *Client) DescribeAlertHistoryListWithChan(request *DescribeAlertHistoryListRequest) (<-chan *DescribeAlertHistoryListResponse, <-chan error) {
	responseChan := make(chan *DescribeAlertHistoryListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAlertHistoryList(request)
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

// DescribeAlertHistoryListWithCallback invokes the cms.DescribeAlertHistoryList API asynchronously
func (client *Client) DescribeAlertHistoryListWithCallback(request *DescribeAlertHistoryListRequest, callback func(response *DescribeAlertHistoryListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAlertHistoryListResponse
		var err error
		defer close(result)
		response, err = client.DescribeAlertHistoryList(request)
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

// DescribeAlertHistoryListRequest is the request struct for api DescribeAlertHistoryList
type DescribeAlertHistoryListRequest struct {
	*requests.RpcRequest
	RuleName   string           `position:"Query" name:"RuleName"`
	StartTime  string           `position:"Query" name:"StartTime"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	State      string           `position:"Query" name:"State"`
	MetricName string           `position:"Query" name:"MetricName"`
	GroupId    string           `position:"Query" name:"GroupId"`
	EndTime    string           `position:"Query" name:"EndTime"`
	Ascending  requests.Boolean `position:"Query" name:"Ascending"`
	Namespace  string           `position:"Query" name:"Namespace"`
	Page       requests.Integer `position:"Query" name:"Page"`
	RuleId     string           `position:"Query" name:"RuleId"`
	Status     string           `position:"Query" name:"Status"`
}

// DescribeAlertHistoryListResponse is the response struct for api DescribeAlertHistoryList
type DescribeAlertHistoryListResponse struct {
	*responses.BaseResponse
	Code             string           `json:"Code" xml:"Code"`
	Message          string           `json:"Message" xml:"Message"`
	RequestId        string           `json:"RequestId" xml:"RequestId"`
	Total            string           `json:"Total" xml:"Total"`
	Success          bool             `json:"Success" xml:"Success"`
	AlarmHistoryList AlarmHistoryList `json:"AlarmHistoryList" xml:"AlarmHistoryList"`
}

// CreateDescribeAlertHistoryListRequest creates a request to invoke DescribeAlertHistoryList API
func CreateDescribeAlertHistoryListRequest() (request *DescribeAlertHistoryListRequest) {
	request = &DescribeAlertHistoryListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DescribeAlertHistoryList", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeAlertHistoryListResponse creates a response to parse from DescribeAlertHistoryList response
func CreateDescribeAlertHistoryListResponse() (response *DescribeAlertHistoryListResponse) {
	response = &DescribeAlertHistoryListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
