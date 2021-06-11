package vo

//消息记录新增 (类似房间)

type ChatRecoredAddParams struct {
	Id           string `json:"id"`
	RecordName   string  `json:"recordName"` //       1 文本  2 表情 3 图片 4 文件
	RecordImg    string `json:"recordImg"`     // require    消息头像
	CreateBy     string `json:"createBy"`      //require     创建者
	LastMsg      string `json:"lastMsg"`        //require     最后的消息
	RecordTalker string `json:"recordTalker"`  //require         0 未撤回  1  撤回
	Token string `json:"token"`
}

//获取消息记录列表

type ChatRecordListParams struct {
	UserId string `json:"userId"`

}

// 获取消息记录列表

type ChatRecordDelParams struct {
	Id string `json:"id"`
	Token string `json:"token"`

}
