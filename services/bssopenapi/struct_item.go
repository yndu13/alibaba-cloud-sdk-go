package bssopenapi

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

// Item is a nested struct in bssopenapi response
type Item struct {
	DeductQuantity            float64 `json:"DeductQuantity" xml:"DeductQuantity"`
	Tax                       float64 `json:"Tax" xml:"Tax"`
	PayerAccount              string  `json:"PayerAccount" xml:"PayerAccount"`
	CoveragePercentage        float64 `json:"CoveragePercentage" xml:"CoveragePercentage"`
	TotalQuantity             float64 `json:"TotalQuantity" xml:"TotalQuantity"`
	BizType                   string  `json:"BizType" xml:"BizType"`
	ResourceInstanceId        string  `json:"ResourceInstanceId" xml:"ResourceInstanceId"`
	CommodityCode             string  `json:"CommodityCode" xml:"CommodityCode"`
	UsageStartTime            string  `json:"UsageStartTime" xml:"UsageStartTime"`
	SolutionCode              string  `json:"SolutionCode" xml:"SolutionCode"`
	ProductDetail             string  `json:"ProductDetail" xml:"ProductDetail"`
	Usage                     string  `json:"Usage" xml:"Usage"`
	ServicePeriodUnit         string  `json:"ServicePeriodUnit" xml:"ServicePeriodUnit"`
	PretaxAmountLocal         float64 `json:"PretaxAmountLocal" xml:"PretaxAmountLocal"`
	OwnerName                 string  `json:"OwnerName" xml:"OwnerName"`
	OutstandingAmount         float64 `json:"OutstandingAmount" xml:"OutstandingAmount"`
	DeductedByResourcePackage string  `json:"DeductedByResourcePackage" xml:"DeductedByResourcePackage"`
	PipCode                   string  `json:"PipCode" xml:"PipCode"`
	SplitProductDetail        string  `json:"SplitProductDetail" xml:"SplitProductDetail"`
	SplitAccountID            string  `json:"SplitAccountID" xml:"SplitAccountID"`
	ProductCode               string  `json:"ProductCode" xml:"ProductCode"`
	ListPrice                 string  `json:"ListPrice" xml:"ListPrice"`
	BillAccountID             string  `json:"BillAccountID" xml:"BillAccountID"`
	Period                    string  `json:"Period" xml:"Period"`
	InvoiceDiscount           float64 `json:"InvoiceDiscount" xml:"InvoiceDiscount"`
	StartTime                 string  `json:"StartTime" xml:"StartTime"`
	EndTime                   string  `json:"EndTime" xml:"EndTime"`
	BucketOwnerId             int64   `json:"BucketOwnerId" xml:"BucketOwnerId"`
	InstanceID                string  `json:"InstanceID" xml:"InstanceID"`
	PretaxGrossAmount         float64 `json:"PretaxGrossAmount" xml:"PretaxGrossAmount"`
	SplitItemName             string  `json:"SplitItemName" xml:"SplitItemName"`
	ImageType                 string  `json:"ImageType" xml:"ImageType"`
	SubscribeBucket           string  `json:"SubscribeBucket" xml:"SubscribeBucket"`
	RecordID                  string  `json:"RecordID" xml:"RecordID"`
	Status                    string  `json:"Status" xml:"Status"`
	Item                      string  `json:"Item" xml:"Item"`
	ProductName               string  `json:"ProductName" xml:"ProductName"`
	Region                    string  `json:"Region" xml:"Region"`
	PaymentAmount             float64 `json:"PaymentAmount" xml:"PaymentAmount"`
	PostpaidCost              string  `json:"PostpaidCost" xml:"PostpaidCost"`
	ReservationCost           string  `json:"ReservationCost" xml:"ReservationCost"`
	UsageEndTime              string  `json:"UsageEndTime" xml:"UsageEndTime"`
	RoundDownDiscount         string  `json:"RoundDownDiscount" xml:"RoundDownDiscount"`
	InternetIP                string  `json:"InternetIP" xml:"InternetIP"`
	CommodityName             string  `json:"CommodityName" xml:"CommodityName"`
	PaymentTime               string  `json:"PaymentTime" xml:"PaymentTime"`
	CostUnit                  string  `json:"CostUnit" xml:"CostUnit"`
	SplitBillingCycle         string  `json:"SplitBillingCycle" xml:"SplitBillingCycle"`
	AfterTaxAmount            float64 `json:"AfterTaxAmount" xml:"AfterTaxAmount"`
	ResourceGroup             string  `json:"ResourceGroup" xml:"ResourceGroup"`
	SplitAccountName          string  `json:"SplitAccountName" xml:"SplitAccountName"`
	SubOrderId                string  `json:"SubOrderId" xml:"SubOrderId"`
	CapacityUnit              string  `json:"CapacityUnit" xml:"CapacityUnit"`
	UsagePercentage           float64 `json:"UsagePercentage" xml:"UsagePercentage"`
	BillingType               string  `json:"BillingType" xml:"BillingType"`
	InstanceSpec              string  `json:"InstanceSpec" xml:"InstanceSpec"`
	SubscribeTime             string  `json:"SubscribeTime" xml:"SubscribeTime"`
	Tag                       string  `json:"Tag" xml:"Tag"`
	RegionNo                  string  `json:"RegionNo" xml:"RegionNo"`
	OwnerID                   string  `json:"OwnerID" xml:"OwnerID"`
	UserId                    string  `json:"UserId" xml:"UserId"`
	InstanceId                string  `json:"InstanceId" xml:"InstanceId"`
	SolutionName              string  `json:"SolutionName" xml:"SolutionName"`
	NickName                  string  `json:"NickName" xml:"NickName"`
	SubscriptionType          string  `json:"SubscriptionType" xml:"SubscriptionType"`
	InstanceConfig            string  `json:"InstanceConfig" xml:"InstanceConfig"`
	ZoneName                  string  `json:"ZoneName" xml:"ZoneName"`
	PotentialSavedCost        string  `json:"PotentialSavedCost" xml:"PotentialSavedCost"`
	DeductedByCashCoupons     float64 `json:"DeductedByCashCoupons" xml:"DeductedByCashCoupons"`
	ServicePeriod             string  `json:"ServicePeriod" xml:"ServicePeriod"`
	MultAccountRelSubscribe   string  `json:"MultAccountRelSubscribe" xml:"MultAccountRelSubscribe"`
	DiscountAmount            float64 `json:"DiscountAmount" xml:"DiscountAmount"`
	ListPriceUnit             string  `json:"ListPriceUnit" xml:"ListPriceUnit"`
	SplitCommodityCode        string  `json:"SplitCommodityCode" xml:"SplitCommodityCode"`
	UsageUnit                 string  `json:"UsageUnit" xml:"UsageUnit"`
	SplitItemID               string  `json:"SplitItemID" xml:"SplitItemID"`
	PaymentCurrency           string  `json:"PaymentCurrency" xml:"PaymentCurrency"`
	SubscribeType             string  `json:"SubscribeType" xml:"SubscribeType"`
	SubscribeLanguage         string  `json:"SubscribeLanguage" xml:"SubscribeLanguage"`
	ProductType               string  `json:"ProductType" xml:"ProductType"`
	SplitBillingDate          string  `json:"SplitBillingDate" xml:"SplitBillingDate"`
	Currency                  string  `json:"Currency" xml:"Currency"`
	DeductedByPrepaidCard     float64 `json:"DeductedByPrepaidCard" xml:"DeductedByPrepaidCard"`
	BillAccountName           string  `json:"BillAccountName" xml:"BillAccountName"`
	PaymentTransactionID      string  `json:"PaymentTransactionID" xml:"PaymentTransactionID"`
	BillingDate               string  `json:"BillingDate" xml:"BillingDate"`
	StatusName                string  `json:"StatusName" xml:"StatusName"`
	BillingItem               string  `json:"BillingItem" xml:"BillingItem"`
	DeductedByCoupons         float64 `json:"DeductedByCoupons" xml:"DeductedByCoupons"`
	Zone                      string  `json:"Zone" xml:"Zone"`
	Quantity                  int64   `json:"Quantity" xml:"Quantity"`
	SavedCost                 string  `json:"SavedCost" xml:"SavedCost"`
	UserName                  string  `json:"UserName" xml:"UserName"`
	PretaxAmount              float64 `json:"PretaxAmount" xml:"PretaxAmount"`
	IntranetIP                string  `json:"IntranetIP" xml:"IntranetIP"`
}
