package common

type ApiResponse struct {
	Code string `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

func (ar *ApiResponse) Error(code,msg string)*ApiResponse{
	ar.Msg = msg
	ar.Code = code
	return ar
}

func (ar *ApiResponse) Success(data interface{})*ApiResponse{
	ar.Data = data
	ar.Code = "200"
	return ar
}
