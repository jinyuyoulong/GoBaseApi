package controller

type APIJson struct {
	Status bool        `json:"status"`
	Msg    interface{} `json:"msg"`
	Data   interface{} `json:"data"`
}

// APIResult 构造json返回
func APIResult(status bool, object interface{}, msg string) (apijson *APIJson) {
	apijson = &APIJson{Status: status, Data: object, Msg: msg}
	return apijson
}
