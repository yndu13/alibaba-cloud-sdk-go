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

// DescribeSiteMonitorData invokes the cms.DescribeSiteMonitorData API synchronously
func (client *Client) DescribeSiteMonitorData(request *DescribeSiteMonitorDataRequest) (response *DescribeSiteMonitorDataResponse, err error) {
	response = CreateDescribeSiteMonitorDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSiteMonitorDataWithChan invokes the cms.DescribeSiteMonitorData API asynchronously
func (client *Client) DescribeSiteMonitorDataWithChan(request *DescribeSiteMonitorDataRequest) (<-chan *DescribeSiteMonitorDataResponse, <-chan error) {
	responseChan := make(chan *DescribeSiteMonitorDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSiteMonitorData(request)
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

// DescribeSiteMonitorDataWithCallback invokes the cms.DescribeSiteMonitorData API asynchronously
func (client *Client) DescribeSiteMonitorDataWithCallback(request *DescribeSiteMonitorDataRequest, callback func(response *DescribeSiteMonitorDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSiteMonitorDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeSiteMonitorData(request)
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

// DescribeSiteMonitorDataRequest is the request struct for api DescribeSiteMonitorData
type DescribeSiteMonitorDataRequest struct {
	*requests.RpcRequest
	Period     string           `position:"Query" name:"Period"`
	Length     requests.Integer `position:"Query" name:"Length"`
	EndTime    string           `position:"Query" name:"EndTime"`
	StartTime  string           `position:"Query" name:"StartTime"`
	Type       string           `position:"Query" name:"Type"`
	NextToken  string           `position:"Query" name:"NextToken"`
	MetricName string           `position:"Query" name:"MetricName"`
	TaskId     string           `position:"Query" name:"TaskId"`
}

// DescribeSiteMonitorDataResponse is the response struct for api DescribeSiteMonitorData
type DescribeSiteMonitorDataResponse struct {
	*responses.BaseResponse
	NextToken string `json:"NextToken" xml:"NextToken"`
	Data      string `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   string `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateDescribeSiteMonitorDataRequest creates a request to invoke DescribeSiteMonitorData API
func CreateDescribeSiteMonitorDataRequest() (request *DescribeSiteMonitorDataRequest) {
	request = &DescribeSiteMonitorDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DescribeSiteMonitorData", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSiteMonitorDataResponse creates a response to parse from DescribeSiteMonitorData response
func CreateDescribeSiteMonitorDataResponse() (response *DescribeSiteMonitorDataResponse) {
	response = &DescribeSiteMonitorDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
