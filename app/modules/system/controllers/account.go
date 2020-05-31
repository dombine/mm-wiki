package controllers

import (
	"strings"

	"github.com/dombine/mm-wiki/app/models"
	"github.com/dombine/mm-wiki/app/utils"

	valid "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type AccountController struct {
	BaseController
}

func (this *AccountController) Add() {
	this.viewLayout("account/form", "account")
}

func (this *AccountController) Save() {

	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/system/account/list")
	}
	name := strings.TrimSpace(this.GetString("name", ""))
	url := strings.TrimSpace(this.GetString("url", ""))
	username := strings.TrimSpace(this.GetString("username", ""))
	password := strings.TrimSpace(this.GetString("password", ""))
	comment := strings.TrimSpace(this.GetString("comment", ""))
	if name == "" {
		this.jsonError("账户名称不能为空！")
	}
	if url != "" && valid.Validate(url, is.URL) != nil {
		this.jsonError("账户网址格式不正确！")
	}
	if username == "" {
		this.jsonError("用户名不能为空！")
	}
	if password == "" {
		this.jsonError("密码不能为空！")
	}
	ok, err := models.AccountModel.HasAccountName(this.UserId, name)
	if err != nil {
		this.ErrorLog("添加账户失败：" + err.Error())
		this.jsonError("添加账户失败！")
	}
	if ok {
		this.jsonError("账户名已经存在！")
	}

	accountId, err := models.AccountModel.Insert(map[string]interface{}{
		"name":     name,
		"url":      url,
		"comment":  comment,
		"user_id":  this.UserId,
		"username": username,
		"password": password,
	})

	if err != nil {
		this.ErrorLog("添加账户失败：" + err.Error())
		this.jsonError("添加账户失败：" + err.Error())
	}
	this.InfoLog("添加账户 " + utils.Convert.IntToString(accountId, 10) + " 成功")
	this.jsonSuccess("添加账户成功", nil, "/system/account/list")
}

func (this *AccountController) List() {

	page, _ := this.GetInt("page", 1)
	keyword := strings.TrimSpace(this.GetString("keyword", ""))
	keyuname := strings.TrimSpace(this.GetString("keyuname", ""))
	keyurl := strings.TrimSpace(this.GetString("keyurl", ""))
	number, _ := this.GetRangeInt("number", 20, 10, 100)
	limit := (page - 1) * number

	var err error
	var count int64
	var accounts []map[string]string
	if keyword != "" || keyuname != "" || keyurl != "" {
		count, err = models.AccountModel.CountAccountsByKeyword(keyword, keyuname, keyurl, this.UserId)
		accounts, err = models.AccountModel.GetAccountsByKeywordAndLimit(keyword, keyuname, keyurl, this.UserId, limit, number)
	} else {
		count, err = models.AccountModel.CountAccounts(this.UserId)
		accounts, err = models.AccountModel.GetAccountsByLimit(this.UserId, limit, number)
	}
	if err != nil {
		this.ErrorLog("获取账户列表失败: " + err.Error())
		this.ViewError("获取账户列表失败", "/system/main/index")
	}

	this.Data["accounts"] = accounts
	this.Data["keyword"] = keyword
	this.Data["keyuname"] = keyuname
	this.Data["keyurl"] = keyurl
	this.SetPaginator(number, count)
	this.viewLayout("account/list", "account")
}

func (this *AccountController) Edit() {

	accountId := this.GetString("account_id", "")
	if accountId == "" {
		this.ViewError("账户不存在", "/system/account/list")
	}

	account, err := models.AccountModel.GetAccountById(accountId)
	if err != nil {
		this.ViewError("账户不存在", "/system/account/list")
	}

	this.Data["account"] = account
	this.viewLayout("account/form", "account")
}

func (this *AccountController) Modify() {

	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/system/account/list")
	}
	accountId := this.GetString("account_id", "")
	name := strings.TrimSpace(this.GetString("name", ""))
	url := strings.TrimSpace(this.GetString("url", ""))
	username := strings.TrimSpace(this.GetString("username", ""))
	password := strings.TrimSpace(this.GetString("password", ""))
	comment := strings.TrimSpace(this.GetString("sequence", ""))

	if accountId == "" {
		this.jsonError("账户不存在！")
	}
	if name == "" {
		this.jsonError("账户名称不能为空！")
	}
	if url != "" && valid.Validate(url, is.URL) != nil {
		this.jsonError("账户网址格式不正确！")
	}
	if username == "" {
		this.jsonError("用户名不能为空！")
	}
	if password == "" {
		this.jsonError("密码不能为空！")
	}

	account, err := models.AccountModel.GetAccountById(accountId)
	if err != nil {
		this.ErrorLog("修改账户 " + accountId + " 失败: " + err.Error())
		this.jsonError("修改账户失败！")
	}
	if len(account) == 0 {
		this.jsonError("账户不存在！")
	}

	ok, _ := models.AccountModel.HasSameName(accountId, name)
	if ok {
		this.jsonError("账户名已经存在！")
	}
	_, err = models.AccountModel.Update(accountId, map[string]interface{}{
		"name":     name,
		"url":      url,
		"comment":  comment,
		"username": username,
		"password": password,
	})

	if err != nil {
		this.ErrorLog("修改账户 " + accountId + " 失败：" + err.Error())
		this.jsonError("修改账户失败")
	}
	this.InfoLog("修改账户 " + accountId + " 成功")
	this.jsonSuccess("修改账户成功", nil, "/system/account/list")
}

func (this *AccountController) Delete() {

	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/system/account/list")
	}
	accountId := this.GetString("account_id", "")
	if accountId == "" {
		this.jsonError("没有选择账户！")
	}

	account, err := models.AccountModel.GetAccountById(accountId)
	if err != nil {
		this.ErrorLog("删除账户 " + accountId + " 失败: " + err.Error())
		this.jsonError("删除账户失败")
	}
	if len(account) == 0 {
		this.jsonError("账户不存在")
	}

	err = models.AccountModel.Delete(accountId)
	if err != nil {
		this.ErrorLog("删除账户 " + accountId + " 失败: " + err.Error())
		this.jsonError("删除账户失败")
	}

	this.InfoLog("删除账户 " + accountId + " 成功")
	this.jsonSuccess("删除账户成功", nil, "/system/account/list")
}
