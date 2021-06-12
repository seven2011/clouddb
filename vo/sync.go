package vo

//同步 文章

type SyncArticleAddParams struct {
	Method string           `json:"method"`
	Data ArticleAddParams  `json:"data"`
}
// 同步 用户

type SyncUserParams struct {
	Method string           `json:"method"`
	Data ArticleAddParams  `json:"data"`
}