package code

import "encoding/json"

type FieldMap map[FieldName]FieldValue

//FieldName 字段名称
type FieldName string

//FieldValue 字段值
type FieldValue interface{}

type EstimateID string

//QuotationSetReq 添加 & 更新报价单请求结构
type QuotationSetReq struct {
	//Fields 更新字段的元祖
	Params FieldMap `json:"params"`
	//EstimateID 冒泡ID
	EstimateID string `json:"estimate_id"`
}

type SetReq struct {
	//Fields 更新字段的元祖
	Params map[string]interface{} `json:"params"`
	//EstimateID 冒泡ID
	EstimateID string `json:"estimate_id"`
}

func LoopMap(req *QuotationSetReq) *SetReq {
	setReq := SetReq{}
	setReq.Params = map[string]interface{}{}
	setReq.EstimateID = string(req.EstimateID)
	for k, v := range req.Params {
		setReq.Params[string(k)] = interface{}(v)
	}

	return &setReq
}

func JsonMap(req *QuotationSetReq) *SetReq {
	setReq := SetReq{}
	data, _ := json.Marshal(req)
	json.Unmarshal(data, setReq)
	return &setReq
}

type InReq struct {
	req1  string
	req2  string
	req3  string
	req4  string
	req5  string
	req6  int
	req7  int
	req8  int
	req9  int
	req10 int
}

type OutReq struct {
	req1  string
	req2  string
	req3  string
	req4  string
	req5  string
	req6  int
	req7  int
	req8  int
	req9  int
	req10 int
}

func LoopStruct(req *InReq) *OutReq {
	out := &OutReq{}
	out.req1 = req.req1
	out.req2 = req.req2
	out.req3 = req.req3
	out.req4 = req.req4
	out.req5 = req.req5
	out.req6 = req.req6
	out.req7 = req.req7
	out.req8 = req.req8
	out.req9 = req.req9
	out.req10 = req.req10
	return out
}

func JsonStruct(req *InReq) *OutReq {
	out := &OutReq{}
	data, _ := json.Marshal(req)
	json.Unmarshal(data, out)
	return out
}
