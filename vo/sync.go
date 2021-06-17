package vo

//同步 文章

type SyncArticleAddParams struct {
	Method string           `json:"method"`
	Data   ArticleAddParams `json:"data"`
}

// 同步 用户

type SyncUserParams struct {
	Method string           `json:"method"`
	Data   ArticleAddParams `json:"data"`
}
type SyncParams struct {
	Method string           `json:"type"`
	Data   ArticleAddParams `json:"data"`
}

type SyncMsgParams struct {
	Method string           `json:"type"`
	Data   interface{}     `json:"data"`
}