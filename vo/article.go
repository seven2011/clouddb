package vo

import "time"

type ArticleAddParams struct {
	Id             string `json:"id"`
	UserId         string `json:"userId"`
	Accesstory     string `json:"accesstory"`
	AccesstoryType int64  `json:"accesstoryType"`
	Text           string `json:"text"`
	Tag            string `json:"tag"`
	Title          string `json:"title"`
	Thumbnail      string `json:"thumbnail"`
}

//  返回 信息

type ArticleResp struct {
	Id             string    `json:"id"`
	UserId         string    `json:"userId"`
	Accesstory     string    `json:"accesstory"`
	AccesstoryType int64     `json:"accesstoryType"`
	Text           string    `json:"text"`
	Tag            string    `json:"tag"`
	Ptime          time.Time `json:"ptime"`
	PlayNum        int64     `json:"playNum"`
	ShareNum       int64     `json:"shareNum"`
	Thumbnail      string    `json:"thumbnail"`
	Title          string    `json:"title"`
	PeerId         string    `json:"peerId"`
	Name           string    `json:"name"`
	Phone          string    `json:"phone"`
	Sex            int64     `json:"sex"`
	NickName       string    `json:"nickName"`
}

//1文本 2图片 3视频 4音乐
/*
id             string     require   后端生成
userId         string     require    用户id
accesstory     string     require    附件 cid  逗号隔开
accesstoryType int64      require    附件类型
text           string     require    正文
tag            string     require    标签

*/

/// article/category

type ArticleCategoryParams struct {
	PageSize       int64 `json:"pageSize"`
	PageNum        int64 `json:"pageNum"`
	AccesstoryType int64 `json:"accesstoryType"`
}

// article play add

type ArticlePlayAddParams struct {
	Id string `json:"id"`
}

//点赞

type ArticleGiveLikeParams struct {
	Id    string `json:"id"`
	Token string `json:"token"`
}
type SyncArticleGiveLikeParams struct {
	Id        string `json:"id"`
	UserId    string `json:"userId"`
	ArticleId string `json:"articleId"`
	Islike    int64  `json:"isLike"`
}

//取消点赞

type ArticleCancelLikeParams struct {
	Id    string `json:"id"`
	Token string `json:"token"`
}

// 获取文章详情信息

type ArticleQueryParams struct {
	Id string `json:"id"`
}
