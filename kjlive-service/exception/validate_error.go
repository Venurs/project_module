package exception

type ValidateError struct {
    Err  error `json:"err"`
    Code string `json:"code"`
    Message map[string]string `json:"message"`
}


func (e *ValidateError) Error() string {
    return e.Code
}
