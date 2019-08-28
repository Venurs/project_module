package models

import (
    "kjlive-service/redis"
    "github.com/astaxie/beego/orm"
    "time"
)

func GetParameter(key string) (params map[string]interface{})  {
    parameters := redis.GetMap("COMMON_PARMETER_VALUE_GO_" + key)
    if parameters == nil {
        params = getParameterByKey(key)
    } else {
        params = parameters
    }
    
    return params
}

type CommonParameter struct {
    ParmKey string `json:"parm_key" orm:"column(parm_key)"`
    ParmValue string `json:"parm_value" orm:"column(parm_value)"`
}

func getParameterByKey(key string) map[string]interface{} {
    o := orm.NewOrm()
    sql := `
		SELECT
		  parm_key        parm_key,
          parm_value      parm_value
		FROM base_commonparmeter
		WHERE cls = ?
			  AND validity_time > ?
			  AND delete_flag = 0
	`
    var commonParameter []CommonParameter
    mapCommonParameter := make(map[string]interface{})
    num, err := o.Raw(sql, key, time.Now()).QueryRows(&commonParameter)
    if err == nil && num > 0{
        for _, param := range commonParameter {
            mapCommonParameter[param.ParmKey] = param.ParmValue
        }
        redis.Set("COMMON_PARMETER_VALUE_GO_" + key, mapCommonParameter, 60 * 60 * time.Second)
    }
    return mapCommonParameter
}