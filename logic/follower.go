package logic

import (
	"furumvv2/dao/mysql"
	"furumvv2/models"
)

func GetFollowerList(user_address string) (responsedata *models.FollowerList, err error) {
	responseFollwerList := new(models.FollowerList)
	data, err := mysql.GetFollowerList(user_address)
	//fmt.Println("data:", data)
	//
	for _, value := range data {

		responseFollwerList.FollowerList = append(responseFollwerList.FollowerList, value.FollowerAddress)
	}
	responseFollwerList.UserAddress = user_address
	followerInfoList, err := mysql.GetFollowerInfo(responseFollwerList.FollowerList)
	if err != nil {
		return nil, err
	}
	responseFollwerList.FollowerInfoList = followerInfoList
	return responseFollwerList, nil
}

func AddFollowerByUserID(followerInfo *models.Follower) (data string, err error) {
	//验证是否已经在关注链表中（是否已经存储数据库）
	exit, err := mysql.ISExitFollower(followerInfo)
	if err != nil {
		return "", err
	}
	if exit {
		//存在
		return "exit", nil
	}
	//添加
	err = mysql.AddFollowerByUserID(followerInfo)
	if err != nil {
		return "", err
	}
	return "success", nil
}
