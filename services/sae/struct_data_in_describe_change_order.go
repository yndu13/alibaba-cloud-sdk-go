package sae

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

// DataInDescribeChangeOrder is a nested struct in sae response
type DataInDescribeChangeOrder struct {
	Status             int        `json:"Status" xml:"Status"`
	Description        string     `json:"Description" xml:"Description"`
	SupportAbortFreeze bool       `json:"SupportAbortFreeze" xml:"SupportAbortFreeze"`
	CreateTime         string     `json:"CreateTime" xml:"CreateTime"`
	ChangeOrderId      string     `json:"ChangeOrderId" xml:"ChangeOrderId"`
	BatchType          string     `json:"BatchType" xml:"BatchType"`
	AppName            string     `json:"AppName" xml:"AppName"`
	Auto               bool       `json:"Auto" xml:"Auto"`
	CurrentPipelineId  string     `json:"CurrentPipelineId" xml:"CurrentPipelineId"`
	CoTypeCode         string     `json:"CoTypeCode" xml:"CoTypeCode"`
	SupportRollback    bool       `json:"SupportRollback" xml:"SupportRollback"`
	BatchWaitTime      int        `json:"BatchWaitTime" xml:"BatchWaitTime"`
	ErrorMessage       string     `json:"ErrorMessage" xml:"ErrorMessage"`
	CoType             string     `json:"CoType" xml:"CoType"`
	BatchCount         int        `json:"BatchCount" xml:"BatchCount"`
	CoTargets          []string   `json:"CoTargets" xml:"CoTargets"`
	Pipelines          []Pipeline `json:"Pipelines" xml:"Pipelines"`
}