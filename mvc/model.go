package mvc

import "time"

//sys_user

type SysUser struct {
	Id       string    `json:"id"`
	PeerId   string    `json:"peerId"`
	Name     string    `json:"name"`
	Phone    string    `json:"phone"`
	Sex      int64     `json:"sex"`
	NickName string    `json:"nickName"`
	Ptime    time.Time `json:"-"`
	Utime    time.Time `json:"-"`
}

//cloud_file
type File struct {
	Id       string    `json:"id"`
	UserId   string    `json:"userId"`
	FileName string    `json:"fileName"`
	ParentId string    `json:"parentId"`
	FileCid  string    `json:"fileCid"`
	FileSize int64     `json:"fileSize"`
	FileType int64     `json:"fileType"`
	IsFolder int64     `json:"isFolder"`
	Ptime    time.Time `json:"ptime"`
}

//DownLoadList
type DownLoad struct {
	Id           string    `json:"id"`
	UserId       string    `json:"userId"`
	FileName     string    `json:"fileName"`
	Ptime        time.Time `json:"ptime"`
	FileCid      string    `json:"fileCid"`
	FileSize     int64     `json:"fileSize"`
	DownPath     string    `json:"downPath"`
	FileType     int64     `json:"fileType"`
	TransferType int64     `json:"transferType"`
}

//
//DownLoadList
type TransferDownLoadParams struct {
	Id             string    `json:"id"`
	UserId         string    `json:"userId"`
	FileName       string    `json:"fileName"`
	Ptime          time.Time `json:"ptime"`
	FileCid        string    `json:"fileCid"`
	FileSize       int64     `json:"fileSize"`
	DownPath       string    `json:"downPath"`
	FileType       int64     `json:"fileType"`
	TransferType   int64     `json:"transferType"`
	UploadParentId string    `json:"uploadParentId"`
	UploadFileId   string    `json:"uploadFileId"`
}

//article

type Article struct {
	Id             string    `json:"id"`
	UserId         string    `json:"userId"`
	Accesstory     string    `json:"accesstory"`
	AccesstoryType int64     `json:"accesstoryType"`
	Text           string    `json:"text"`
	Tag            string    `json:"tag"`
	Ptime          time.Time `json:"ptime"`
	PlayNum        int64     `json:"playNum"`
	ShareNum       int64     `json:"shareNum"`
	Title          string    `json:"title"`
	Thumbnail      string    `json:"thumbnail"`
	FileName      string    `json:"fileName"`
	FileType      string    `json:"fileType"`
}

//article like
type ArticleLike struct {
	Id        string `json:"id"`
	UserId    string `json:"userId"`
	ArticleId string `json:"articleId"`
	IsLike    int64  `json:"isLike"`
}

// chat_msg
type ChatMsg struct {
	Id          string    `json:"id"`
	ContentType int64     `json:"contentType"`
	Content     string    `json:"content"`
	FromId      string    `json:"fromId"`
	ToId        string    `json:"toId"`
	Ptime       time.Time `json:"ptime"`
	IsWithdraw  int64     `json:"isWithdraw"` //require     coment 是否撤回         0 未撤回  1  撤回
	IsRead      int64     `json:"isRead"`
	RecordId    string    `json:"recordId"`
}

// chat_record

type ChatRecord struct {
	Id      string    `json:"id"`
	Name    string    `json:"name"`
	Img     string    `json:"img"`
	FromId  string    `json:"fromId"`
	Ptime   time.Time `json:"ptime"`
	LastMsg string    `json:"lastMsg"`
	Toid    string    `json:"toId"`
}

type CopyParams struct {
	Pid      string
	CopyFile []File
}
type MoveParams struct {
	Pid      string
	MoveFile []File
}

//delete one file

type DeleteOneParams struct {
	DropFile []File
}

//delete many file

type DeleteManyParams struct {
	DropFile []File
}

//delete one file

type DeleteOneDirParams struct {
	DropFile []File
}

//delete one file

type DeleteManyDirParams struct {
	DropFile []File
}

//delete
type DeleteParams struct {
	DropFile []File
}
