package client

import (
	"encoding/json"
	"errors"
	"github.com/huobirdcenter/huobi_golang/internal"
	"github.com/huobirdcenter/huobi_golang/internal/requestbuilder"
	"github.com/huobirdcenter/huobi_golang/pkg/model"
	"github.com/huobirdcenter/huobi_golang/pkg/model/linearswap"
)

// Responsible to operate linear swap
type LinearSwapClient struct {
	privateUrlBuilder *requestbuilder.PrivateUrlBuilder
}

// Initializer
func (p *LinearSwapClient) Init(accessKey string, secretKey string, host string) *LinearSwapClient {
	p.privateUrlBuilder = new(requestbuilder.PrivateUrlBuilder).Init(accessKey, secretKey, host)
	return p
}

// Returns linear swap positions info.
func (p *LinearSwapClient) QueryUserPositionInfo(request *linearswap.CrossLinearSwapPositionInfoRequest) ([]linearswap.CrossLinearSwapPositionInfo, error) {
	postBody, jsonErr := model.ToJson(request)
	if jsonErr != nil {
		return nil, jsonErr
	}
	url := p.privateUrlBuilder.Build("POST", "/linear-swap-api/v1/swap_cross_position_info", nil)
	postResp, postErr := internal.HttpPost(url, postBody)
	if postErr != nil {
		return nil, postErr
	}
	result := linearswap.CrossLinearSwapPositionInfoResponse{}
	if jsonErr = json.Unmarshal([]byte(postResp), &result); jsonErr != nil {
		return nil, jsonErr
	}
	if result.Status == "ok" && result.Data != nil {
		return result.Data, nil
	}
	return nil, errors.New(postResp)
}

// Returns user's linear swap account info.
func (p *LinearSwapClient) QueryUserAccountInfo(request *linearswap.CrossLinearSwapUserAccountInfoRequest) ([]linearswap.CrossLinearSwapUserAccountInfo, error) {
	postBody, jsonErr := model.ToJson(request)
	if jsonErr != nil {
		return nil, jsonErr
	}
	url := p.privateUrlBuilder.Build("POST", "/linear-swap-api/v1/swap_cross_account_info", nil)
	postResp, postErr := internal.HttpPost(url, postBody)
	if postErr != nil {
		return nil, postErr
	}
	result := linearswap.CrossLinearSwapUserAccountInfoResponse{}
	if jsonErr = json.Unmarshal([]byte(postResp), &result); jsonErr != nil {
		return nil, jsonErr
	}
	if result.Status == "ok" && result.Data != nil {
		return result.Data, nil
	}
	return nil, errors.New(postResp)
}

// Returns user's linear financial records.
func (p *LinearSwapClient) QueryAccountFinancialRecords(request *linearswap.CrossLinearSwapAccountFinancialRecordsRequest) (resp *linearswap.CrossLinearSwapAccountFinancialRecordsData, err error) {
	var postBody string
	if postBody, err = model.ToJson(request); err != nil {
		return
	}
	url := p.privateUrlBuilder.Build("POST", "/linear-swap-api/v1/swap_financial_record", nil)
	var postResp string
	if postResp, err = internal.HttpPost(url, postBody); err != nil {
		return
	}
	result := new(linearswap.CrossLinearSwapAccountFinancialRecordsResponse)
	if err = json.Unmarshal([]byte(postResp), result); err != nil {
		return
	}
	if result.Status == "ok" && result.Data != nil {
		return result.Data, nil
	}
	return nil, errors.New(postResp)
}

func (p *LinearSwapClient) QueryAccountFinancialRecordsViaMultipleFields(request *linearswap.QueryAccountFinancialRecordsViaMultipleFieldsRequest) (resp *linearswap.CrossLinearSwapAccountFinancialRecordsViaMultipleFieldsData, err error) {
	var postBody string
	if postBody, err = model.ToJson(request); err != nil {
		return
	}
	url := p.privateUrlBuilder.Build("POST", "/linear-swap-api/v1/swap_financial_record_exact", nil)
	var postResp string
	if postResp, err = internal.HttpPost(url, postBody); err != nil {
		return
	}
	result := new(linearswap.CrossLinearSwapAccountFinancialRecordsViaMultipleFieldsResponse)
	if err = json.Unmarshal([]byte(postResp), result); err != nil {
		return
	}
	if result.Status == "ok" && result.Data != nil {
		return result.Data, nil
	}
	return nil, errors.New(postResp)
}

// Returns user's linear assets and positions.
func (p *LinearSwapClient) QueryAssetsAndPositions(request *linearswap.CrossLinearAssetsAndPositionsRequest) (resp *linearswap.CrossLinearAssetsAndPositionsInfo, err error) {
	var postBody string
	if postBody, err = model.ToJson(request); err != nil {
		return
	}
	url := p.privateUrlBuilder.Build("POST", "/linear-swap-api/v1/swap_cross_account_position_info", nil)
	var postResp string
	if postResp, err = internal.HttpPost(url, postBody); err != nil {
		return
	}
	result := new(linearswap.CrossLinearAssetsAndPositionsResponse)
	if err = json.Unmarshal([]byte(postResp), result); err != nil {
		return
	}
	if result.Status == "ok" && result.Data != nil {
		return result.Data, nil
	}
	return nil, errors.New(postResp)
}

// Switch leverage of linear positions.
func (p *LinearSwapClient) SwitchLeverage(request *linearswap.CrossLinearSwapSwitchLeverageRequest) (resp *linearswap.CrossLinearSwapSwitchLeverageResponseData, err error) {
	var postBody string
	if postBody, err = model.ToJson(request); err != nil {
		return
	}
	url := p.privateUrlBuilder.Build("POST", "/linear-swap-api/v1/swap_cross_switch_lever_rate", nil)
	var postResp string
	if postResp, err = internal.HttpPost(url, postBody); err != nil {
		return
	}
	result := new(linearswap.CrossLinearSwapSwitchLeverageResponse)
	if err = json.Unmarshal([]byte(postResp), result); err != nil {
		return
	}
	if result.Status == "ok" && result.Data != nil {
		return result.Data, nil
	}
	return nil, errors.New(postResp)
}

// Returns information of user's linear order
func (p *LinearSwapClient) GetInformationOfOrder(request *linearswap.GetInformationOfOrderRequest) (order []*linearswap.Order, err error) {
	var body string
	if body, err = model.ToJson(request); err != nil {
		return
	}
	url := p.privateUrlBuilder.Build("POST", "/linear-swap-api/v1/swap_cross_order_info", nil)
	if body, err = internal.HttpPost(url, body); err != nil {
		return
	}
	result := new(linearswap.GetInformationOfOrderResponse)
	if err = json.Unmarshal([]byte(body), result); err != nil {
		return
	}
	if result.Status == "ok" && result.Data != nil {
		return result.Data, nil
	}
	return nil, errors.New(body)
}

// Returns detail information of user's linear order
func (p *LinearSwapClient) GetDetailInformationOfOrder(request *linearswap.GetDetailInformationOfOrderRequest) (resp *linearswap.DetailInformationOfOrder, err error) {
	var body string
	if body, err = model.ToJson(request); err != nil {
		return
	}
	url := p.privateUrlBuilder.Build("POST", "/linear-swap-api/v1/swap_cross_order_detail", nil)
	if body, err = internal.HttpPost(url, body); err != nil {
		return
	}
	result := new(linearswap.GetDetailInformationOfOrderResponse)
	if err = json.Unmarshal([]byte(body), result); err != nil {
		return
	}
	if result.Status == "ok" && result.Data != nil {
		return result.Data, nil
	}
	return nil, errors.New(body)
}

// Returns user's linear history order
func (p *LinearSwapClient) GetHistoryOrders(request *linearswap.GetHistoryOrdersRequest) (resp *linearswap.GetHistoryOrdersData, err error) {
	var body string
	if body, err = model.ToJson(request); err != nil {
		return
	}
	url := p.privateUrlBuilder.Build("POST", "/linear-swap-api/v1/swap_cross_hisorders", nil)
	if body, err = internal.HttpPost(url, body); err != nil {
		return
	}
	result := new(linearswap.GetHistoryOrdersResponse)
	if err = json.Unmarshal([]byte(body), result); err != nil {
		return
	}
	if result.Status == "ok" && result.Data != nil {
		return result.Data, nil
	}
	return nil, errors.New(body)
}

// Returns user's linear history order via multiple fields
func (p *LinearSwapClient) GetHistoryOrdersViaMultipleFields(request *linearswap.GetHistoryOrdersViaMultipleFieldsRequest) (resp *linearswap.GetHistoryOrdersViaMultipleFieldsData, err error) {
	var body string
	if body, err = model.ToJson(request); err != nil {
		return
	}
	url := p.privateUrlBuilder.Build("POST", "/linear-swap-api/v1/swap_cross_hisorders_exact", nil)
	if body, err = internal.HttpPost(url, body); err != nil {
		return
	}
	result := new(linearswap.GetHistoryOrdersViaMultipleFieldsResponse)
	if err = json.Unmarshal([]byte(body), result); err != nil {
		return
	}
	if result.Status == "ok" && result.Data != nil {
		return result.Data, nil
	}
	return nil, errors.New(body)
}

// Cancel user's linear swap order
func (p *LinearSwapClient) CancelOrder(request *linearswap.CancelOrderRequest) (resp *linearswap.CancelOrderResult, err error) {
	var body string
	if body, err = model.ToJson(request); err != nil {
		return
	}
	url := p.privateUrlBuilder.Build("POST", "/linear-swap-api/v1/swap_cross_cancel", nil)
	if body, err = internal.HttpPost(url, body); err != nil {
		return
	}
	result := new(linearswap.CancelOrderResponse)
	if err = json.Unmarshal([]byte(body), result); err != nil {
		return
	}
	if result.Status == "ok" && result.Data != nil {
		return result.Data, nil
	}
	return nil, errors.New(body)
}

func (p *LinearSwapClient) PlaceOrder(request *linearswap.PlaceOrderRequest) (resp *linearswap.PlaceOrderData, err error) {
	var body string
	if body, err = model.ToJson(request); err != nil {
		return
	}
	url := p.privateUrlBuilder.Build("POST", "/linear-swap-api/v1/swap_cross_order", nil)
	if body, err = internal.HttpPost(url, body); err != nil {
		return
	}
	result := new(linearswap.PlaceOrderResponse)
	if err = json.Unmarshal([]byte(body), result); err != nil {
		return
	}
	if result.Status == "ok" && result.Data != nil {
		return result.Data, nil
	}
	return nil, errors.New(body)
}

func (p *LinearSwapClient) SwitchPositionMode(request *linearswap.SwitchPositionModeRequest) (resp []*linearswap.PositionMode, err error) {
	var body string
	if body, err = model.ToJson(request); err != nil {
		return
	}
	url := p.privateUrlBuilder.Build("POST", "/linear-swap-api/v1/swap_cross_switch_position_mode", nil)
	if body, err = internal.HttpPost(url, body); err != nil {
		return
	}
	result := new(linearswap.SwitchPositionModeResponse)
	if err = json.Unmarshal([]byte(body), result); err != nil {
		return
	}
	if result.Status == "ok" && result.Data != nil {
		return result.Data, nil
	}
	return nil, errors.New(body)
}
