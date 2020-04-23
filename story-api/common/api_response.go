package common

type ApiResponse struct {
	Code string `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

func (ar *ApiResponse) Success(data interface{}){
	ar.Data = data
	ar.Code = "200"
}
