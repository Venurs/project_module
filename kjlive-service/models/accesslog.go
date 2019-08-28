package models

import (
    "github.com/astaxie/beego/orm"
    "time"
    "strconv"
    "kjlive-service/redis"
)

func GetLiveAccessNumber(liveId int, numberGrowthFlag bool) (accessNumber int)  {
    accessNumberString := redis.Get("LIVE_CONSTASNT_INFO_USER_COUNT_" +  strconv.Itoa(liveId))
    if accessNumberString == "" {
        o := orm.NewOrm()
        sql := `
		SELECT
            Sum(access_count) access_number
        FROM live_liveaccesslog
        WHERE live_info_id = ?;
	    `
        err := o.Raw(sql, liveId).QueryRow(&accessNumber)
        if err == nil {
            redis.Set("LIVE_CONSTASNT_INFO_USER_COUNT_" +  strconv.Itoa(liveId), strconv.Itoa(accessNumber), 24 * 30 * time.Hour)
        }
    } else {
        var err error
        accessNumber, err = strconv.Atoi(accessNumberString)
        if err == nil &&numberGrowthFlag {
            accessNumber += 1
            redis.Set("LIVE_CONSTASNT_INFO_USER_COUNT_" +  strconv.Itoa(liveId), strconv.Itoa(accessNumber), 24 * 30 * time.Hour)
        }
    }
    return
}

func InsertLiveAccessLog(userId int, liveInfoId int, promotionManageId int) (num int64, err error) {
    o := orm.NewOrm()

    var isExist int
    if promotionManageId == 0 {
        sql := `
            SELECT
                Count(id) isExist
            FROM live_liveaccesslog
            WHERE live_info_id = ?
                AND user_id = ?
                AND access_time = ?
            LIMIT 1
        `
        err = o.Raw(sql, liveInfoId, userId, time.Now().Format("2006-01-02")).QueryRow(&isExist)
    } else {
        sql := `
            SELECT
                Count(id) isExist
            FROM live_liveaccesslog
            WHERE live_info_id = ?
                AND user_id = ?
                AND promotion_manage_id = ?
                AND access_time = ?
            LIMIT 1
        `
        err = o.Raw(sql, liveInfoId, userId, promotionManageId, time.Now().Format("2006-01-02")).QueryRow(&isExist)
    }
    if err != nil  {
        return 0, err
    }
    if isExist > 0 {
        return 0, nil
    } else {
        o.Using("replica")
        var rawSeter orm.RawSeter
        if promotionManageId == 0 {
            insertSql := `
                INSERT INTO
                  live_liveaccesslog (create_time, last_update_time, delete_flag, live_info_id, user_id, access_count, access_time)
                  VALUE (?, ?, ?, ?, ?, ?, ?)
            `
            rawSeter = o.Raw(insertSql, time.Now(), time.Now(), 0, liveInfoId, userId, 1, time.Now())
        } else {
            insertSql := `
                INSERT INTO
                  live_liveaccesslog (create_time, last_update_time, delete_flag, live_info_id, user_id, access_count, access_time, promotion_manage_id)
                  VALUE (?, ?, ?, ?, ?, ?, ?, ?)
            `
            rawSeter = o.Raw(insertSql, time.Now(), time.Now(), 0, liveInfoId, userId, 1, time.Now(), promotionManageId)
        }
        liveinfocmsSql := `UPDATE live_liveinfocms SET online_number = online_number + 1 WHERE liveinfo_id = ?`
        onUpdate := o.Raw(liveinfocmsSql, liveInfoId)
        onUpdate.Exec()
        res, err := rawSeter.Exec()
        if err == nil {
            num, _ := res.RowsAffected()
            GetLiveAccessNumber(liveInfoId, true)
            return num, nil
        }
        return 0, err
    }
}
