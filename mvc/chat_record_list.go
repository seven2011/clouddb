package mvc

import (
	"encoding/json"

	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
)

func ChatRecordList(db *Sql, value string) ([]vo.ChatRecordRespListParams, error) {
	var crd []vo.ChatRecordInfo
	var crdToid []vo.ChatRecordInfo
	var link []vo.ChatRecordRespListParams

	var result vo.ChatRecordListParams

	sugar.Log.Debug("Request Param: ", value)

	err := json.Unmarshal([]byte(value), &result)

	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
		return link, err
	}
	sugar.Log.Info("Marshal data is  ", result)

	if err != nil {
		sugar.Log.Error("Insert into article table is failed.", err)
		return link, err
	}
	sugar.Log.Info("result := ", result)

	//token verify
	claim, b := jwt.JwtVeriyToken(result.Token)
	if !b {
		return link, err
	}

	sugar.Log.Info("claim := ", claim)
	//userId:=claim["UserId"]

	rows, err := db.DB.Query("SELECT a.id, a.name, a.from_id, a.ptime, a.last_msg, a.to_id, b.name as username, b.nickname, b.peer_id, b.phone, b.sex, b.img from chat_record as a LEFT JOIN sys_user as b on a.from_id=b.id where a.from_id=?", result.FromId)
	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
		//	return arrfile, errors.New("查询下载列表信息失败")
	}
	for rows.Next() {
		/*
			UserName string  `json:"userName"`
			Phone    string `json:"phone"`
			PeerId   string `json:"peerId"`
			NickName string `json:"nickName"`
			Sex      int64  `json:"sex"
		*/
		var dl vo.ChatRecordInfo
		err = rows.Scan(&dl.Id, &dl.Name, &dl.FromId, &dl.Ptime, &dl.LastMsg, &dl.Toid, &dl.UserName, &dl.NickName, &dl.PeerId, &dl.Phone, &dl.Sex, &dl.Img)
		if err != nil {
			sugar.Log.Error("Query scan data is failed.The err is ", err)
			return link, err
		}
		sugar.Log.Info("Query a entire data is ", dl)
		crd = append(crd, dl)
	}
	if err != nil {
		sugar.Log.Error("Query  chat_record  is Failed.", err)
		return link, err
	}

	sugar.Log.Info("  这是 from  id 的信息  ", crd)

	//在去 查出 结果集对应的 toId
	//

	for _, v := range crd {
		sugar.Log.Info("  查看  to  id 的信息 v ------ ", v)

		rows, err := db.DB.Query("SELECT a.id, a.name, a.from_id, a.ptime, a.last_msg, a.to_id, b.name as username,b.nickname, b.peer_id, b.phone, b.sex, b.img from chat_record as a LEFT JOIN sys_user as b on a.to_id=b.id where a.from_id = ? and a.to_id=?", v.FromId, v.Toid)
		if err != nil {
			sugar.Log.Error("Query data is failed.Err is ", err)
			//	return arrfile, errors.New("查询下载列表信息失败")
		}
		for rows.Next() {
			/*
				UserName string  `json:"userName"`
				Phone    string `json:"phone"`
				PeerId   string `json:"peerId"`
				NickName string `json:"nickName"`
				Sex      int64  `json:"sex"
			*/
			var dl vo.ChatRecordInfo
			err = rows.Scan(&dl.Id, &dl.Name, &dl.FromId, &dl.Ptime, &dl.LastMsg, &dl.Toid, &dl.UserName, &dl.Phone, &dl.PeerId, &dl.NickName, &dl.Sex, &dl.Img)
			if err != nil {
				sugar.Log.Error("Query scan data is failed.The err is ", err)
				return link, err
			}
			sugar.Log.Info("Query a entire data is ", dl)
			crdToid = append(crdToid, dl)
		}

	}
	sugar.Log.Info("  这是 to id 的信息  ", crdToid)

	//拼接

	//
	var lin vo.ChatRecordRespListParams

	for _, v := range crdToid {
		lin.ToId = v.Toid
		lin.ToImg = v.Img
		lin.ToName = v.UserName
		lin.ToNickName = v.NickName
		lin.ToPeerId = v.PeerId
		lin.ToPhone = v.Phone
		lin.ToSex = v.Sex
		sugar.Log.Info("  这是  v name 的信息  ", v.Name)

		if len(crd) > 0 {
			sugar.Log.Info("  这是  from 的  name 的信息crd[0].Name  ", crd[0].Name)

			lin.FromImg = crd[0].Img
			lin.FromName = crd[0].UserName
			lin.FromNickName = crd[0].NickName
			lin.FromPeerId = crd[0].PeerId
			lin.FromPhone = crd[0].Phone
			lin.FromSex = crd[0].Sex
			lin.FromId = crd[0].FromId
		}

		link = append(link, lin)
	}

	sugar.Log.Info("  最终结果 ---------- link = ", link)

	return link, nil

}
