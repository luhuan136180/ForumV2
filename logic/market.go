package logic

import (
	"furumvv2/dao/mysql"
	"furumvv2/models"
)

func GetAllSkinList(status int) (data []*models.Skin, err error) {
	return mysql.GetAllSkinList(status)
}

func ShopSkinByUser(shopInformation *models.Shop) (err error) {
	//验证
	//查询是否已经买了
	err = mysql.ISExitSkin(shopInformation)
	if err != nil {
		return err
	}
	//查询余额是否充足
	err = mysql.ISEnoughAmount(shopInformation)
	if err != nil {
		return err
	}
	//对数据库操作
	//购买--存入user_skin表中
	//对user的balance减操作
	err = mysql.ShopSkinByUser(shopInformation)
	if err != nil {
		return err
	}
	return nil
}
