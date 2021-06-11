package mvc

import (
	"encoding/json"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
)

func ChatRecordList(db *Sql, value string) ([]ChatRecord, error) {
	var crd []ChatRecord
	var result vo.ChatRecordListParams
	err := json.Unmarshal([]byte(value), &result)

	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
		return crd,err
	}
	sugar.Log.Info("Marshal data is  ", result)

	if err != nil {
		sugar.Log.Error("Insert into article table is failed.", err)
		return crd,err
	}
	sugar.Log.Info("result := ", result)


	rows, err := db.DB.Query("SELECT * FROM chat_record where user_id=?", result.UserId)
	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
		//	return arrfile, errors.New("查询下载列表信息失败")
	}
	for rows.Next() {
		var dl ChatRecord
		err = rows.Scan(&dl.Id, &dl.RecordName, &dl.RecordImg, &dl.CreateBy, &dl.Ptime, &dl.LastMsg, &dl.RecordTalker)
		if err != nil {
			sugar.Log.Error("Query scan data is failed.The err is ", err)
			return crd, err
		}
		sugar.Log.Info("Query a entire data is ", dl)
		crd = append(crd, dl)
	}
	if err != nil {
		sugar.Log.Error("Query  chat_record  is Failed.", err)
		return crd,err
	}

	sugar.Log.Info(" 这是 查询数据库 获取的 结果 = ",crd)

	return crd,nil

}
