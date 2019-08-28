package middlewares

import (
    "github.com/gin-gonic/gin"
    "kjlive-service/utils"
    "encoding/json"
    "kjlive-service/logs"
    "strconv"
)

type returnResult struct {
    Message map[string]interface{} `json:"message"`
    Code string `json:"code"`
    Data interface{} `json:"data"`
    Version float32 `json:"version"`
    Err interface{} `json:"err"`
}


func ApiInterfaceMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        var wb *utils.ResponseBuffer
        if c.Writer.Status() == 404 {
            c.Next()
            return
        } else if w, ok := c.Writer.(gin.ResponseWriter); ok {
            wb = utils.NewResponseBuffer(w)
            c.Writer = wb
            c.Next()
        } else {
            c.Next()
            return
        }

        result := &returnResult{}
        result.Version = 2.0
        body := wb.Body.Bytes()
        var mapJsonBody map[string]interface{}
        json.Unmarshal(body, &mapJsonBody)
        if wb.Status() == 200 {
            if code, ok := mapJsonBody["code"]; ok {
                result.Code = code.(string)
                result.Message = mapJsonBody["message"].(map[string]interface{})
                result.Err = mapJsonBody["err"]
            } else {
                result.Code = "ok"
                result.Data = mapJsonBody
            }
            wb.Body.Reset()
            jsonResult, err := json.Marshal(result)
            if err != nil {
                logs.Error(err)
            }else {
                wb.Body.Write(jsonResult)
            }
            wb.Header().Set("Content-Type", "application/json")
            wb.Header().Set("Content-Length", strconv.Itoa(wb.Body.Len()))
            wb.Flush()
        } else if wb.Status() == 403 {
            wb.Flush()
        }
    }
}
