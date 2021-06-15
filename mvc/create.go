package mvc

import (
	"encoding/json"
	"fmt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
)
//






//------------------------------------------------//-
// 初始化数据库                                      //|

func (db *Sql) Ping() error {                      //|
	err := db.DB.Ping()                            //|
	if err != nil {                                //|
		sugar.Log.Error("Ping is Failed.",err)  //|
	}                                              //|
	return err                                     //|
}                                                  //|
//-------------------------------------------------|-





//  ==============   User  =================
//  用户注册

func (db *Sql) UserRegister(user string) string {
	err:= AddUser(db, user)
	//返回封装成方法
	// 返回的时候 要改东西
	if err!=nil{
		return vo.ResponseErrorMsg(400,err.Error())
	}
	return vo.ResponseSuccess()
}
//




//  用户注销

func (db *Sql) UserLoginOut(user string) string {
	e := UserDel(db, user)

	if e != nil {
		return vo.ResponseErrorMsg(400,e.Error())
	}
	return vo.ResponseSuccess()
}


//   用户登录

func (db *Sql) UserLogin(user string) string {
	token,e := UserLogin(db,user)

	if e != nil {
		return vo.ResponseErrorMsg(400,e.Error())
	}
	return vo.ResponseSuccess(token)
}


//  用户查询

func (db *Sql) UserQuery(user string) string {

	data,e := UserQuery(db, user)

	if e != nil {
		return vo.ResponseErrorMsg(400,e.Error())
	}
	return vo.ResponseSuccess(data)
}
//用户 更新

func (db *Sql) UserUpdate(user string) string {

	e := UserUpdate(db, user)

	if e != nil {
		return vo.ResponseErrorMsg(400,e.Error())
	}
	return vo.ResponseSuccess()
}

//  ==============   User  End=================




//================   Cloud ====================
//  添加文件

func (db *Sql) AddFile(fInfo string) string {
	fileId,e := AddFile(db, fInfo)

	if e != nil {
		return vo.ResponseErrorMsg(400,e.Error())
	}
	return vo.ResponseSuccess(fileId)
}


//   添加文件夹

func (db *Sql) AddFolder(fInfo string) string {
	e := AddFolder(db, fInfo)

	if e != nil {
		return vo.ResponseErrorMsg(400,e.Error())
	}
	return vo.ResponseSuccess()
}

// 删除文件

func (db *Sql) DeleteOneFile(dInfo string) string {
	e := DeleteOneFile(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400,e.Error())
	}

	return vo.ResponseSuccess()
}

//重命名

func (db *Sql) FileRename(dInfo string) string {
	e := CloudFileRename(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400,e.Error())
	}

	return vo.ResponseSuccess()
}

// 获取文件层级列表

func (db *Sql) FileList(dInfo string) string {
	data,e := CloudFileList(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400,e.Error())
	}

	return vo.ResponseSuccess(data)
}

// 获取文件夹层级列表

func (db *Sql) FolderList(dInfo string) string {
	data,e := CloudFolderList(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400,e.Error())
	}

	return vo.ResponseSuccess(data)
}

// 上传记录传输

func (db *Sql) TransferAdd(dInfo string) string {
	e := DownLoadFile(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400,e.Error())
	}
	return vo.ResponseSuccess()
}

//  根据文件进行分类

func (db *Sql) FileCategory(dInfo string) string {
	data,e := CloudFileCategory(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400,e.Error())
	}

	return vo.ResponseSuccess(data)
}
//  删除传输记录

func (db *Sql) TransferDel(dInfo string) string {
	e := TransferDel(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400,e.Error())
	}

	return vo.ResponseSuccess()
}

//  传输列表

func (db *Sql) TransferList(dInfo string) string {
	data,e := TransferList(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400,e.Error())
	}

	return vo.ResponseSuccess(data)
}

//================   Cloud  ENd ====================




// ================  Article  =====================

//  添加 朋友圈文章

func (db *Sql)ArticleAdd(dInfo string)string{
	e :=AddArticle(db,dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400,e.Error())
	}

	return vo.ResponseSuccess()
}
//  查找arttcle list

func (db *Sql)ArticleList(dInfo string)string{
	data,e :=ArticleList(db,dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400,e.Error())
	}

	return vo.ResponseSuccess(data)
}
// 文章列表分类

func (db *Sql)ArticleCategory(dInfo string)string{

	data,e := ArticleCategory(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400,e.Error())
	}

	return vo.ResponseSuccess(data)
}

// 文章增加播放次数

func (db *Sql)ArticlePlayAdd(dInfo string)string{

	e := ArticlePlayAdd(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400,e.Error())
	}

	return vo.ResponseSuccess()
}


// 增加播放量

func (db *Sql)ArticleShareAdd(dInfo string)string{

	e := ArticleShareAdd(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400,e.Error())
	}

	return vo.ResponseSuccess()
}

//  朋友圈 点赞

func (db *Sql)ArticleGiveLike(dInfo string)string{

	e := AddArticleLike(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400,e.Error())
	}

	return vo.ResponseSuccess()
}

// 取消点赞

func (db *Sql)ArticleCancelLike(dInfo string)string{

	e := ArticleCancelLike(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400,e.Error())
	}

	return vo.ResponseSuccess()
}
// 获取文章详情

func (db *Sql)ArticleQuery(dInfo string)string{

	data,e := ArticleQuery(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400,e.Error())
	}

	return vo.ResponseSuccess(data)
}
// 文章查询

func (db *Sql)ArticleSearch(dInfo string)string{

	data,e := ARticleSearch(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400,e.Error())
	}

	return vo.ResponseSuccess(data)
}


// 我发布的 文章信息

func (db *Sql)ArticleAboutMe(dInfo string)string{

	data,e := ArticleAboutMe(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400,e.Error())
	}

	return vo.ResponseSuccess(data)
}




// ================  Article   End   =================

//===========123======   Chat           ==================

//保存消息

func (db *Sql)ChatAddMsg(dInfo string)string{

	e := AddChatMsg(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400,e.Error())
	}

	return vo.ResponseSuccess()
}

//获取 消息 分页

func (db *Sql)ChatMsgList(dInfo string)string{

	data,e := ChatMsgList(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400,e.Error())
	}

	return vo.ResponseSuccess(data)
}

// 删除消息

func (db *Sql)ChatMsgDel(dInfo string)string{

	e := ChatMsgDel(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400,e.Error())
	}
	return vo.ResponseSuccess()
}


// 消息记录 新增

func (db *Sql)ChatRecordAdd(dInfo string)string{

	recordId,e := ChatRecordAdd(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400,e.Error())
	}
	return vo.ResponseSuccess(recordId)
}

// 获取消息记录列表


func (db *Sql)ChatRecordList(dInfo string)string{

	data,e := ChatRecordList(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400,e.Error())
	}
	return vo.ResponseSuccess(data)
}

//  删除记录

func (db *Sql)ChatRecordDel(dInfo string)string{

	e := ChatRecordDel(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400,e.Error())
	}
	return vo.ResponseSuccess()
}

//  撤回消息

func (db *Sql)ChatMsgWithDraw(dInfo string)string{

	e := ChatWithDraw(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400,e.Error())
	}
	return vo.ResponseSuccess()
}
//=================   Chat     End   ==================


//  下载列表

func (db *Sql) DownloadList(dInfo string) string {
	data, e := DownloadList(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400,e.Error())
	}


	return vo.ResponseSuccess(data)
}
//  复制文件

func (db *Sql) CopyFile(dInfo string) string {
	 e := CopyFile(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400,e.Error())
	}
	return vo.ResponseSuccess()
}

//  移动文件

func (db *Sql) MoveFile(dInfo string) string {
	 e := MoveFile(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400,e.Error())
	}

	return vo.ResponseSuccess()
}






//  聊天记录

func (db *Sql)AddChatRecord(dInfo string)string{

	e := AddChatRecord(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400,e.Error())
	}

	return vo.ResponseSuccess()
}


//删除

func (db *Sql)DeleteAll(dInfo string)string{

	e := Delete(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400,e.Error())
	}

	return vo.ResponseSuccess()
}

//
//delete

func (db *Sql)CloudFindList(dInfo string)string{

	result,e := CloudFindList(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400,e.Error())
	}
	return vo.ResponseSuccess(result)
}


//查询

func (db *Sql)CloudSearch(dInfo string)string{

	result,e := Search(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400,e.Error())
	}
	return vo.ResponseSuccess(result)
}


//===================  同步 ==========================
//同步 User 数据

func (db *Sql)SyncUser(dInfo string)string{

	e := SyncUser(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400,e.Error())
	}
	return vo.ResponseSuccess()
}

// 同步 文章表数据

//===================  同步 ==========================

func (db *Sql)SyncArticle(dInfo string)string{

	e := SyncArticle(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400,e.Error())
	}
	return vo.ResponseSuccess()
}

// 同步文章点赞

func (db *Sql)SyncArticleGiveLike(dInfo string)string{

	e := SyncAticlePlay(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400,e.Error())
	}
	return vo.ResponseSuccess()
}

// 同步 文章 取消点赞

func (db *Sql)SyncArticleCancelLike(dInfo string)string{

	e := SyncAticlePlay(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400,e.Error())
	}
	return vo.ResponseSuccess()
}

//



//同步 分享


func (db *Sql)SyncArticleShare(dInfo string)string{

	e := SyncArticleShare(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400,e.Error())
	}
	return vo.ResponseSuccess()
}


//convert

func ConvertString(value string, t interface{}) (res map[string]interface{}) {

	json.Unmarshal([]byte(value), &t)
	fmt.Println(" 这是 获得的结果 ", t)
	fmt.Printf(" 这是 获得的结果  %T\n", t)

	return t.(map[string]interface{})
}

