package mysql

import (
	"furumvv2/models"
	"github.com/jmoiron/sqlx"
)

func GetFollowerList(user_address string) (data []*models.Follower, err error) {
	data = make([]*models.Follower, 0)

	sqlStr := `select user_address,follower_address from user_following where user_address=?`
	err = Db.Select(&data, sqlStr, user_address)
	//fmt.Println("mysql:", data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func ISExitFollower(followerInfo *models.Follower) (exit bool, err error) {
	sqlStr := `select count(*) from user_following where user_address=? and follower_address=?`

	var count int
	err = Db.Get(&count, sqlStr, followerInfo.UserAddress, followerInfo.FollowerAddress)
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func AddFollowerByUserID(followerInfo *models.Follower) (err error) {
	sqlStr := `insert into user_following(user_address,follower_address) values(?,?)`
	_, err = Db.Exec(sqlStr, followerInfo.UserAddress, followerInfo.FollowerAddress)
	if err != nil {
		return err
	}
	return nil
}

func GetFollowerInfo(FollowerAddList []string) (data []*models.FollowerInfo, err error) {
	data = make([]*models.FollowerInfo, len(FollowerAddList))
	sqlStr := `select user_address,user_name,signature,picture_url from user where user_address in(?)`
	query, args, err := sqlx.In(sqlStr, FollowerAddList)
	if err != nil {
		return nil, err
	}
	err = Db.Select(&data, query, args...)
	if err != nil {
		return nil, err
	}
	return
}
