package iot

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

// QueryConsumerGroupList invokes the iot.QueryConsumerGroupList API synchronously
// api document: https://help.aliyun.com/api/iot/queryconsumergrouplist.html
func (client *Client) QueryConsumerGroupList(request *QueryConsumerGroupListRequest) (response *QueryConsumerGroupListResponse, err error) {
	response = CreateQueryConsumerGroupListResponse()
	err = client.DoAction(request, response)
	return
}

// QueryConsumerGroupListWithChan invokes the iot.QueryConsumerGroupList API asynchronously
// api document: https://help.aliyun.com/api/iot/queryconsumergrouplist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryConsumerGroupListWithChan(request *QueryConsumerGroupListRequest) (<-chan *QueryConsumerGroupListResponse, <-chan error) {
	responseChan := make(chan *QueryConsumerGroupListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryConsumerGroupList(request)
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

// QueryConsumerGroupListWithCallback invokes the iot.QueryConsumerGroupList API asynchronously
// api document: https://help.aliyun.com/api/iot/queryconsumergrouplist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryConsumerGroupListWithCallback(request *QueryConsumerGroupListRequest, callback func(response *QueryConsumerGroupListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryConsumerGroupListResponse
		var err error
		defer close(result)
		response, err = client.QueryConsumerGroupList(request)
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

// QueryConsumerGroupListRequest is the request struct for api QueryConsumerGroupList
type QueryConsumerGroupListRequest struct {
	*requests.RpcRequest
	IotInstanceId string           `position:"Query" name:"IotInstanceId"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	Fuzzy         requests.Boolean `position:"Query" name:"Fuzzy"`
	CurrentPage   requests.Integer `position:"Query" name:"CurrentPage"`
	GroupName     string           `position:"Query" name:"GroupName"`
	ApiProduct    string           `position:"Body" name:"ApiProduct"`
	ApiRevision   string           `position:"Body" name:"ApiRevision"`
}

// QueryConsumerGroupListResponse is the response struct for api QueryConsumerGroupList
type QueryConsumerGroupListResponse struct {
	*responses.BaseResponse
	RequestId    string                       `json:"RequestId" xml:"RequestId"`
	Success      bool                         `json:"Success" xml:"Success"`
	ErrorMessage string                       `json:"ErrorMessage" xml:"ErrorMessage"`
	PageSize     int                          `json:"PageSize" xml:"PageSize"`
	PageCount    int                          `json:"PageCount" xml:"PageCount"`
	Total        int                          `json:"Total" xml:"Total"`
	CurrentPage  int                          `json:"CurrentPage" xml:"CurrentPage"`
	Code         string                       `json:"Code" xml:"Code"`
	Data         DataInQueryConsumerGroupList `json:"Data" xml:"Data"`
}

// CreateQueryConsumerGroupListRequest creates a request to invoke QueryConsumerGroupList API
func CreateQueryConsumerGroupListRequest() (request *QueryConsumerGroupListRequest) {
	request = &QueryConsumerGroupListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryConsumerGroupList", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryConsumerGroupListResponse creates a response to parse from QueryConsumerGroupList response
func CreateQueryConsumerGroupListResponse() (response *QueryConsumerGroupListResponse) {
	response = &QueryConsumerGroupListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}