package common

type ApiResponse struct {
	Code string `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(data interface{})*ApiResponse{
	ar := &ApiResponse{
		Code: "200",
		Msg: "success",
		Data: data,
	}
	return ar
}

func Error(code,msg string)*ApiResponse{
	ar := &ApiResponse{
		Code: code,
		Msg: msg,
	}
	return ar
}
