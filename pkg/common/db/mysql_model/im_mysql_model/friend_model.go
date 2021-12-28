package im_mysql_model

import (
	"Open_IM/pkg/common/db"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

func InsertToFriend(toInsertFollow *Friend) error {
	dbConn, err := db.DB.MysqlDB.DefaultGormDB()
	if err != nil {
		return err
	}
	toInsertFollow.CreateTime = time.Now()

	err = dbConn.Table("friend").Create(toInsertFollow).Error
	if err != nil {
		return err
	}
	return nil
}

func GetFriendRelationshipFromFriend(OwnerUserID, FriendUserID string) (*Friend, error) {
	dbConn, err := db.DB.MysqlDB.DefaultGormDB()
	if err != nil {
		return nil, err
	}
	var friend Friend
	err = dbConn.Table("friend").Where("owner_user_id=? and friend_user_id=?", OwnerUserID, FriendUserID).Find(&friend).Error
	if err != nil {
		return nil, err
	}
	return &friend, err
}

func GetFriendListByUserID(OwnerUserID string) ([]Friend, error) {
	dbConn, err := db.DB.MysqlDB.DefaultGormDB()
	if err != nil {
		return nil, err
	}
	var friends []Friend
	err = dbConn.Table("friend").Where("owner_user_id=?", OwnerUserID).Find(&friends).Error
	if err != nil {
		return nil, err
	}
	return friends, nil
}

func UpdateFriendComment(OwnerUserID, FriendUserID, Remark string) error {
	dbConn, err := db.DB.MysqlDB.DefaultGormDB()
	if err != nil {
		return err
	}
	err = dbConn.Exec("update friend set remark=? where owner_user_id=? and friend_user_id=?", Remark, OwnerUserID, FriendUserID).Error
	return err
}

func DeleteSingleFriendInfo(OwnerUserID, FriendUserID string) error {
	dbConn, err := db.DB.MysqlDB.DefaultGormDB()
	if err != nil {
		return err
	}
	err = dbConn.Table("friend").Where("owner_user_id=? and friend_user_id=?", OwnerUserID, FriendUserID).Delete(Friend{}).Error
	return err
}

//type Friend struct {
//	OwnerUserID    string    `gorm:"column:owner_user_id;primaryKey;"`
//	FriendUserID   string    `gorm:"column:friend_user_id;primaryKey;"`
//	Remark         string    `gorm:"column:remark"`
//	CreateTime     time.Time `gorm:"column:create_time"`
//	AddSource      int32     `gorm:"column:add_source"`
//	OperatorUserID string    `gorm:"column:operator_user_id"`
//	Ex             string    `gorm:"column:ex"`
//}
