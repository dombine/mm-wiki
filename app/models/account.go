package models

import (
	"github.com/dombine/mm-wiki/app/utils"
	"github.com/snail007/go-activerecord/mysql"
	"time"
)

const Table_Account_Name = "account"

type Account struct {
}

var AccountModel = Account{}

// get account by account_id
func (a *Account) GetAccountById(accountId string) (account map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Account_Name).Where(map[string]interface{}{
		"account_id": accountId,
	}))
	if err != nil {
		return
	}
	account = rs.Row()
	return
}

// account_id and name is exists
func (a *Account) HasSameName(accountId, name string) (has bool, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Account_Name).Where(map[string]interface{}{
		"account_id <>": accountId,
		"name":          name,
	}).Limit(0, 1))
	if err != nil {
		return
	}
	if rs.Len() > 0 {
		has = true
	}
	return
}

// name is exists
func (a *Account) HasAccountName(userId string, name string) (has bool, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Account_Name).Where(map[string]interface{}{
		"name":    name,
		"user_id": userId,
	}).Limit(0, 1))
	if err != nil {
		return
	}
	if rs.Len() > 0 {
		has = true
	}
	return
}

// get account by name
func (a *Account) GetAccountByName(userId string, name string) (account map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Account_Name).Where(map[string]interface{}{
		"name":    name,
		"user_id": userId,
	}).Limit(0, 1))
	if err != nil {
		return
	}
	account = rs.Row()
	return
}

// delete account by account_id
func (a *Account) Delete(accountId string) (err error) {
	db := G.DB()
	_, err = db.Exec(db.AR().Delete(Table_Account_Name, map[string]interface{}{
		"account_id": accountId,
	}))
	if err != nil {
		return
	}
	return
}

// insert account
func (a *Account) Insert(accountValue map[string]interface{}) (id int64, err error) {

	accountValue["create_time"] = time.Now().Unix()
	accountValue["update_time"] = time.Now().Unix()
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Exec(db.AR().Insert(Table_Account_Name, accountValue))
	if err != nil {
		return
	}
	id = rs.LastInsertId
	return
}

// update account by account_id
func (a *Account) Update(accountId string, accountValue map[string]interface{}) (id int64, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	accountValue["update_time"] = time.Now().Unix()
	rs, err = db.Exec(db.AR().Update(Table_Account_Name, accountValue, map[string]interface{}{
		"account_id": accountId,
	}))
	if err != nil {
		return
	}
	id = rs.LastInsertId
	return
}

// get limit accounts by search keyword
func (a *Account) GetAccountsByKeywordAndLimit(keyword string, keyuname string, keyurl string, userId string, limit int, number int) (accounts []map[string]string, err error) {

	db := G.DB()
	where := make(map[string]interface{})
	if keyword != "" {
		where["name LIKE"] = "%" + keyword + "%"
	}
	if keyuname != "" {
		where["username LIKE"] = "%" + keyuname + "%"
	}
	if keyurl != "" {
		where["url LIKE"] = "%" + keyurl + "%"
	}
	where["user_id"] = userId
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Account_Name).Where(where).Limit(limit, number).OrderBy("account_id", "DESC"))
	if err != nil {
		return
	}
	accounts = rs.Rows()

	return
}

// get limit accounts
func (a *Account) GetAccountsByLimit(userId string, limit int, number int) (accounts []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().
			From(Table_Account_Name).
			Where(map[string]interface{}{
				"user_id": userId,
			}).
			Limit(limit, number).
			OrderBy("account_id", "DESC"))
	if err != nil {
		return
	}
	accounts = rs.Rows()

	return
}

// get all accounts
func (a *Account) GetAccounts() (accounts []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().From(Table_Account_Name))
	if err != nil {
		return
	}
	accounts = rs.Rows()
	return
}

// get all accounts by id
func (a *Account) GetAccountsOrderById(userId string) (accounts []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().From(Table_Account_Name).Where(map[string]interface{}{
			"user_id": userId,
		}).OrderBy("account_id", "ASC"))
	if err != nil {
		return
	}
	accounts = rs.Rows()
	return
}

// get account count
func (a *Account) CountAccounts(userId string) (count int64, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().
		Select("count(*) as total").
		From(Table_Account_Name).
		Where(map[string]interface{}{
			"user_id": userId,
		}))
	if err != nil {
		return
	}
	count = utils.NewConvert().StringToInt64(rs.Value("total"))
	return
}

// get account count by keyword
func (a *Account) CountAccountsByKeyword(keyword string, keyuname string, keyurl string, userId string) (count int64, err error) {

	db := G.DB()
	where := make(map[string]interface{})
	if keyword != "" {
		where["name LIKE"] = "%" + keyword + "%"
	}
	if keyuname != "" {
		where["username LIKE"] = "%" + keyuname + "%"
	}
	if keyurl != "" {
		where["url LIKE"] = "%" + keyurl + "%"
	}
	where["user_id"] = userId
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().
		Select("count(*) as total").
		From(Table_Account_Name).
		Where(where))
	if err != nil {
		return
	}
	count = utils.NewConvert().StringToInt64(rs.Value("total"))
	return
}

// get accounts by like name
func (a *Account) GetAccountsByLikeName(name string) (accounts []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Account_Name).Where(map[string]interface{}{
		"name Like": "%" + name + "%",
	}).Limit(0, 1))
	if err != nil {
		return
	}
	accounts = rs.Rows()
	return
}

// get account by many account_id
func (a *Account) GetAccountByAccountIds(accountIds []string) (accounts []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Account_Name).Where(map[string]interface{}{
		"account_id": accountIds,
	}))
	if err != nil {
		return
	}
	accounts = rs.Rows()
	return
}
