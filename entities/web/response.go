package web

type Response struct {
	Code    int         `json:"status_code"`
	Message string      `json:"message"`
	Remark  interface{} `json:"remark,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// Out is interface ...
func (r *Response) Out(code int, message string, remark, data interface{}) {
	r.Code = code
	r.Message = message
	r.Remark = remark
	r.Data = data
}

// InternalServerError is method for internal server error
func (r *Response) InternalServerError(message string, remark interface{}) {
	//r.Code = 500
	//r.Message = message
	r.Out(500, message, remark, nil)
}

// Success is method for succeed
func (r *Response) Success(message string, data interface{}) {
	//r.Code = 200
	//r.Message = message
	//r.Data = data
	r.Out(200, message, nil, data)
}

// Success is method for succeed
func (r *Response) SuccessCreate(message string, data interface{}) {
	//r.Code = 200
	//r.Message = message
	//r.Data = data
	r.Out(201, message, nil, data)
}

// BadRequest is method for bad request
func (r *Response) BadRequest(message string, remark interface{}) {
	//r.Code = 400
	//r.Message = message
	r.Out(400, message, remark, nil)
}
