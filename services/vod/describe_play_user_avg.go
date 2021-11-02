package vod

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

// DescribePlayUserAvg invokes the vod.DescribePlayUserAvg API synchronously
func (client *Client) DescribePlayUserAvg(request *DescribePlayUserAvgRequest) (response *DescribePlayUserAvgResponse, err error) {
	response = CreateDescribePlayUserAvgResponse()
	err = client.DoAction(request, response)
	return
}

// DescribePlayUserAvgWithChan invokes the vod.DescribePlayUserAvg API asynchronously
func (client *Client) DescribePlayUserAvgWithChan(request *DescribePlayUserAvgRequest) (<-chan *DescribePlayUserAvgResponse, <-chan error) {
	responseChan := make(chan *DescribePlayUserAvgResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribePlayUserAvg(request)
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

// DescribePlayUserAvgWithCallback invokes the vod.DescribePlayUserAvg API asynchronously
func (client *Client) DescribePlayUserAvgWithCallback(request *DescribePlayUserAvgRequest, callback func(response *DescribePlayUserAvgResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribePlayUserAvgResponse
		var err error
		defer close(result)
		response, err = client.DescribePlayUserAvg(request)
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

// DescribePlayUserAvgRequest is the request struct for api DescribePlayUserAvg
type DescribePlayUserAvgRequest struct {
	*requests.RpcRequest
	StartTime string           `position:"Query" name:"StartTime"`
	EndTime   string           `position:"Query" name:"EndTime"`
	OwnerId   requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribePlayUserAvgResponse is the response struct for api DescribePlayUserAvg
type DescribePlayUserAvgResponse struct {
	*responses.BaseResponse
	RequestId          string             `json:"RequestId" xml:"RequestId"`
	UserPlayStatisAvgs UserPlayStatisAvgs `json:"UserPlayStatisAvgs" xml:"UserPlayStatisAvgs"`
}

// CreateDescribePlayUserAvgRequest creates a request to invoke DescribePlayUserAvg API
func CreateDescribePlayUserAvgRequest() (request *DescribePlayUserAvgRequest) {
	request = &DescribePlayUserAvgRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "DescribePlayUserAvg", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribePlayUserAvgResponse creates a response to parse from DescribePlayUserAvg response
func CreateDescribePlayUserAvgResponse() (response *DescribePlayUserAvgResponse) {
	response = &DescribePlayUserAvgResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
