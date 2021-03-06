package gateway

type CommonResponse struct {
	Status int           `json:"status"`
	Error *ErrorResponse `json:"error"`
	Result interface{}   `json:"result"`
}

// バリデートのエラーレスポンスを作成する
func (re *CommonResponse) CreateValidateErrorResponse(error string) *CommonResponse {
	re.Status = StatusError
	// ErrorResponseの定義
	errorResponse := &ErrorResponse{}
	errorResponse.Code = ErrorCodeRequestValidate
	errorResponse.Error = error
	re.Error = errorResponse

	return re
}

// SQLエラーのレスポンスを生成する
func (re *CommonResponse) CreateSQLErrorResponse(error string) *CommonResponse {
	re.Status = StatusError
	// ErrorResponseの定義
	errorResponse := &ErrorResponse{}
	errorResponse.Code = ErrorCodeSQL
	errorResponse.Error = error
	re.Error = errorResponse

	return re
}

// 正常時のレスポンスを作成する
func (re *CommonResponse) CreateSuccessResponse(result interface{}) *CommonResponse {
	re.Status = StatusSuccess
	re.Result = result
	re.Error = nil

	return re
}

// レスポンスステータス
const (
	StatusSuccess = 200 // 正常時
	StatusError = 500   // エラー時
)
