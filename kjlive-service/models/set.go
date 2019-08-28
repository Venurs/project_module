package models

import (
    "github.com/astaxie/beego/orm"
    "fmt"
    "strings"
)


func IsBuySet(setIdList []int, userId int) (isBuy bool) {
    o := orm.NewOrm()
    var strQuestionMark = strings.Repeat("?,", len(setIdList))
    strQuestionMark = strQuestionMark[0:len(strQuestionMark)-1]
    sql := fmt.Sprintf(`
        SELECT
          Count(*) as count
        FROM set_userset T0
        WHERE T0.set_id in (%v)
              AND T0.user_id = ?
              AND T0.delete_flag = 0
    `, strQuestionMark)
    var count int
    err := o.Raw(sql, setIdList, userId).QueryRow(&count)
    if err == nil && count > 0{
        return true
    } else {
        return false
    }
}