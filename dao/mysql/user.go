package mysql

import (
	"database/sql"
	"fmt"
	"furumvv2/models"
	"strings"
)

const secret = "forum-v2"

func CheckUserExist(userAddress string) (bool bool, err error) {
	sqlStr := "select count(*) from user where user_address=?"
	var count int

	if err := Db.Get(&count, sqlStr, userAddress); err != nil {
		return false, err
	}

	if count > 0 { //已经注册 进数据库
		return true, nil
	}

	//没有注册
	return false, nil
}

func InsertUser(user *models.User) (data *models.ResponseLogin, err error) {
	data = new(models.ResponseLogin)
	//执行SQL语句入库
	InsqlStr := "insert into user(user_address,user_name,balance,picture_url) values(?,?,?,?)"
	SelectSql := `select user_name,picture_url,user_address from user where user_address = ?`
	_, err = Db.Exec(InsqlStr, user.UserAddress, user.UserName, user.Balance, user.Picture)

	err = Db.Get(data, SelectSql, user.UserAddress)
	fmt.Println("2", err)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func Login(user *models.User) (data *models.ResponseLogin, err error) {
	data = new(models.ResponseLogin)
	sqlStr := "select user_address,user_name,picture_url from user where user_address=?"
	err = Db.Get(data, sqlStr, user.UserAddress)
	fmt.Println(*data)
	if err != nil {
		if err == sql.ErrNoRows {
			//没有查询到
			return nil, ErrorUserNotExist
		}
		//查询数据库失败
		return nil, err
	}
	data.UserAddress = strings.ToLower(data.UserAddress)
	user.UserAddress = strings.ToLower(user.UserAddress)
	//判断是否相等
	if data.UserAddress != user.UserAddress {
		return nil, ErrorInvalidPassword
	}
	return data, nil
}

func GetUserBalance(user_address string) (data *models.GetBalance, err error) {
	data = new(models.GetBalance)
	sqlStr := "select user_name,user_address,balance from user where user_address=?"
	err = Db.Get(data, sqlStr, user_address)
	if err == sql.ErrNoRows {
		//没有查询到
		return nil, ErrorUserNotExist
	}
	if err != nil {
		//查询数据库失败
		return nil, err
	}

	return
}

func AddUserBalance(user_address string, amount int) (data *models.GetBalance, err error) {
	data = new(models.GetBalance)
	//fmt.Println(transcation)
	sqlStr := "select user_name,user_address,balance from user where user_address=?"
	err = Db.Get(data, sqlStr, user_address)
	if err != nil {
		return nil, err
	}
	data.Balance += amount
	sqlStr2 := "update user set balance=? where user_address=?"
	_, err = Db.Exec(sqlStr2, data.Balance, data.UserAddress)
	if err != nil {
		return nil, err
	}

	err = Db.Get(data, sqlStr, user_address)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func SubUserBalance(user_address string, amount int) (data *models.GetBalance, err error) {
	data = new(models.GetBalance)
	//fmt.Println(transcation)
	sqlStr := "select user_name,user_address,balance from user where user_address=?"
	err = Db.Get(data, sqlStr, user_address)
	if err != nil {
		return nil, err
	}
	data.Balance -= amount
	sqlStr2 := "update user set balance=? where user_address=?"
	_, err = Db.Exec(sqlStr2, data.Balance, data.UserAddress)
	if err != nil {
		return nil, err
	}

	err = Db.Get(data, sqlStr, user_address)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//
func GetUserInformation(user_address string) (data *models.UserInformation, err error) {
	data = new(models.UserInformation)
	sqlStr := `select user_address,user_name,create_time,email,age,signature,gender,picture_url,experience,level 
				from user where user_address=? `
	err = Db.Get(data, sqlStr, user_address)
	if err != nil {
		return nil, err
	}
	return
}
func GetUserInformationInside(user_address string) (data *models.UserInformationInside, err error) {
	data = new(models.UserInformationInside)
	sqlStr := `select user_address,user_name,balance,create_time,email,age,signature,gender,picture_url,experience,level 
				from user where user_address=? `
	err = Db.Get(data, sqlStr, user_address)
	if err != nil {
		return nil, err
	}
	return
}
func GetAllSkinByUser(user_address string) (data []*models.SkinListByUser, err error) {
	data = make([]*models.SkinListByUser, 0)
	sqlStr := `select 
		u.user_address,s.skin_url,s.status,s.skin_id
		from skin as s 
		join user_skin as u
		on s.skin_id=u.skin_id
 		where u.user_address=?`
	err = Db.Select(&data, sqlStr, user_address)
	if err != nil {
		return nil, err
	}
	return
}

func ChangeUserInformation(update *models.UpdateProfile) (err error) {
	sqlStr2 := "update user set user_name=?,email=?,gender=?,signature=?,picture_url=?,age=? where user_address=?"
	_, err = Db.Exec(sqlStr2, update.Username, update.Email, update.Gender, update.Signature, update.HeadPicture, update.Age, update.UserAddress)
	if err != nil {
		return err
	}
	return nil
}

func GetPostFromUserAdd(user_address string) (data []*models.GetPostByList, err error) {
	data = make([]*models.GetPostByList, 0)
	sqlStr := `select title,content,user_name,post.post_id,user.user_address from post 
			join user on user.user_address=post.author_address 
			join postpicture on postpicture.post_id = post.post_id
			where user.user_address=? and status=1 ORDER BY post.id DESC;`
	err = Db.Select(&data, sqlStr, user_address)
	if err != nil {
		return nil, err
	}
	return
}
