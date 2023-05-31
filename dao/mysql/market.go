package mysql

import (
	"furumvv2/models"
)

func GetAllSkinList(status int) (data []*models.Skin, err error) {
	data = make([]*models.Skin, 0)
	sqlStr := `select skin_id,skin_url,create_time,status,price from skin where status=?`
	err = Db.Select(&data, sqlStr, status)
	if err != nil {
		return nil, err
	}
	return
}

func ISExitSkin(shopInformation *models.Shop) (err error) {
	sqlStr := `select count(*) from user_skin where user_address=? and skin_id=?`
	var count int
	if err := Db.Get(&count, sqlStr, shopInformation.UserAddress, shopInformation.SkinID); err != nil {
		return err
	}
	if count > 0 { //已经注册 进数据库
		return ErrorSkinISExit
	}
	//没有注册
	return nil
}

func ISEnoughAmount(shopInformation *models.Shop) (err error) {
	sqlStr := `select balance from user where user_address=?`
	var balance int
	err = Db.Get(&balance, sqlStr, shopInformation.UserAddress)
	if err != nil {
		return err
	}
	if balance >= shopInformation.Price {
		return ErrorNotEnoughAmount
	}
	return nil
}

func ShopSkinByUser(shopInforamtion *models.Shop) (err error) {
	sqlInsert := `insert into user_skin(user_address,skin_id,status) values(?,?,?)`
	sqlUpdate := `update user set balance=balance-? where user_address=?`

	tx, err := Db.Begin()
	if err != nil {
		// 错误处理
		return err
	}
	_, err = tx.Exec(sqlInsert, shopInforamtion.UserAddress, shopInforamtion.SkinID, shopInforamtion.Status)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.Exec(sqlUpdate, shopInforamtion.Price, shopInforamtion.UserAddress)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		// 回滚事务
		tx.Rollback()
		// 错误处理
		return err
	}
	return nil
}
