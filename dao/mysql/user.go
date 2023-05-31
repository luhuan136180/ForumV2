package mysql

import (
	"database/sql"
	"furumvv2/models"
)

const secret = "forum-v2"

func CheckUserExist(userAddress string) (bool bool, err error) {
	sqlStr := "select count(*) from user where user_address=?"
	var count int
	if err := Db.Get(&count, sqlStr, userAddress); err != nil {
		return false, err
	}
	if count > 0 { //已经注册 进数据库
		return true, ErrorUserExist
	}
	//没有注册
	return false, nil
}

func InsertUser(user *models.User) (err error) {
	//对密码加密

	//执行SQL语句入库
	sqlStr := "insert into user(user_address,user_name,balance) values(?,?,?)"
	_, err = Db.Exec(sqlStr, user.UserAddress, user.UserName, user.Balance)
	return
}

func Login(user *models.User) (err error) {
	getuser := new(models.User)
	sqlStr := "select user_address,user_name from user where user_address=?"
	err = Db.Get(getuser, sqlStr, user.UserAddress)

	if err != nil {
		if err == sql.ErrNoRows {
			//没有查询到
			return ErrorUserNotExist
		}
		//查询数据库失败
		return err
	}
	//判断是否相等
	if getuser.UserAddress != user.UserAddress {
		return ErrorInvalidPassword
	}
	return nil
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
	sqlStr := `select user_address,user_name,balance,create_time from user where user_address=?`
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
