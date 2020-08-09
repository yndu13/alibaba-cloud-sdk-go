package vcs

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

// GetDeviceConfig invokes the vcs.GetDeviceConfig API synchronously
// api document: https://help.aliyun.com/api/vcs/getdeviceconfig.html
func (client *Client) GetDeviceConfig(request *GetDeviceConfigRequest) (response *GetDeviceConfigResponse, err error) {
	response = CreateGetDeviceConfigResponse()
	err = client.DoAction(request, response)
	return
}

// GetDeviceConfigWithChan invokes the vcs.GetDeviceConfig API asynchronously
// api document: https://help.aliyun.com/api/vcs/getdeviceconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetDeviceConfigWithChan(request *GetDeviceConfigRequest) (<-chan *GetDeviceConfigResponse, <-chan error) {
	responseChan := make(chan *GetDeviceConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDeviceConfig(request)
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

// GetDeviceConfigWithCallback invokes the vcs.GetDeviceConfig API asynchronously
// api document: https://help.aliyun.com/api/vcs/getdeviceconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetDeviceConfigWithCallback(request *GetDeviceConfigRequest, callback func(response *GetDeviceConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDeviceConfigResponse
		var err error
		defer close(result)
		response, err = client.GetDeviceConfig(request)
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

// GetDeviceConfigRequest is the request struct for api GetDeviceConfig
type GetDeviceConfigRequest struct {
	*requests.RpcRequest
	DeviceTimeStamp string `position:"Body" name:"DeviceTimeStamp"`
	DeviceSn        string `position:"Body" name:"DeviceSn"`
}

// GetDeviceConfigResponse is the response struct for api GetDeviceConfig
type GetDeviceConfigResponse struct {
	*responses.BaseResponse
	AudioEnable   bool          `json:"AudioEnable" xml:"AudioEnable"`
	AudioFormat   string        `json:"AudioFormat" xml:"AudioFormat"`
	BitRate       string        `json:"BitRate" xml:"BitRate"`
	Code          string        `json:"Code" xml:"Code"`
	DeviceAddress string        `json:"DeviceAddress" xml:"DeviceAddress"`
	DeviceName    string        `json:"DeviceName" xml:"DeviceName"`
	EncodeFormat  string        `json:"EncodeFormat" xml:"EncodeFormat"`
	FrameRate     string        `json:"FrameRate" xml:"FrameRate"`
	GovLength     int           `json:"GovLength" xml:"GovLength"`
	Latitude      string        `json:"Latitude" xml:"Latitude"`
	Longitude     string        `json:"Longitude" xml:"Longitude"`
	Message       string        `json:"Message" xml:"Message"`
	OSDTimeEnable string        `json:"OSDTimeEnable" xml:"OSDTimeEnable"`
	OSDTimeType   string        `json:"OSDTimeType" xml:"OSDTimeType"`
	OSDTimeX      string        `json:"OSDTimeX" xml:"OSDTimeX"`
	OSDTimeY      string        `json:"OSDTimeY" xml:"OSDTimeY"`
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	Resolution    string        `json:"Resolution" xml:"Resolution"`
	RetryInterval string        `json:"RetryInterval" xml:"RetryInterval"`
	DeviceId      string        `json:"DeviceId" xml:"DeviceId"`
	UserName      string        `json:"UserName" xml:"UserName"`
	PassWord      string        `json:"PassWord" xml:"PassWord"`
	Protocol      string        `json:"Protocol" xml:"Protocol"`
	ServerId      string        `json:"ServerId" xml:"ServerId"`
	ServerPort    string        `json:"ServerPort" xml:"ServerPort"`
	ServerIp      string        `json:"ServerIp" xml:"ServerIp"`
	OSDList       []OSDListItem `json:"OSDList" xml:"OSDList"`
}

// CreateGetDeviceConfigRequest creates a request to invoke GetDeviceConfig API
func CreateGetDeviceConfigRequest() (request *GetDeviceConfigRequest) {
	request = &GetDeviceConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vcs", "2020-05-15", "GetDeviceConfig", "vcs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetDeviceConfigResponse creates a response to parse from GetDeviceConfig response
func CreateGetDeviceConfigResponse() (response *GetDeviceConfigResponse) {
	response = &GetDeviceConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
