package models

import (
	"github.com/dombine/mm-wiki/app/utils"
	"github.com/snail007/go-activerecord/mysql"
	"time"
	"sort"
)

const Table_Link_Name = "link"

type Link struct {
}

var LinkModel = Link{}

// get link by link_id
func (l *Link) GetLinkByLinkId(linkId string) (link map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Link_Name).Where(map[string]interface{}{
		"link_id": linkId,
	}))
	if err != nil {
		return
	}
	link = rs.Row()
	return
}

// link_id and name is exists
func (l *Link) HasSameName(linkId, name string) (has bool, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Link_Name).Where(map[string]interface{}{
		"link_id <>": linkId,
		"name":       name,
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
func (l *Link) HasLinkName(userId string, name string) (has bool, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Link_Name).Where(map[string]interface{}{
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

// get link by name
func (l *Link) GetLinkByName(userId string, name string) (link map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Link_Name).Where(map[string]interface{}{
		"name":    name,
		"user_id": userId,
	}).Limit(0, 1))
	if err != nil {
		return
	}
	link = rs.Row()
	return
}

// delete link by link_id
func (l *Link) Delete(linkId string) (err error) {
	db := G.DB()
	_, err = db.Exec(db.AR().Delete(Table_Link_Name, map[string]interface{}{
		"link_id": linkId,
	}))
	if err != nil {
		return
	}
	return
}

// insert link
func (l *Link) Insert(linkValue map[string]interface{}) (id int64, err error) {

	linkValue["create_time"] = time.Now().Unix()
	linkValue["update_time"] = time.Now().Unix()
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Exec(db.AR().Insert(Table_Link_Name, linkValue))
	if err != nil {
		return
	}
	id = rs.LastInsertId
	return
}

// update link by link_id
func (l *Link) Update(linkId string, linkValue map[string]interface{}) (id int64, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	linkValue["update_time"] = time.Now().Unix()
	rs, err = db.Exec(db.AR().Update(Table_Link_Name, linkValue, map[string]interface{}{
		"link_id": linkId,
	}))
	if err != nil {
		return
	}
	id = rs.LastInsertId
	return
}

// get limit links by search keyword
func (l *Link) GetLinksByKeywordAndLimit(keyword string, userId string, limit int, number int) (links []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Link_Name).Where(map[string]interface{}{
		"name LIKE": "%" + keyword + "%",
		"user_id":   userId,
	}).Limit(limit, number).OrderBy("link_id", "DESC"))
	if err != nil {
		return
	}
	links = rs.Rows()

	return
}

// get limit links
func (l *Link) GetLinksByLimit(userId string, limit int, number int) (links []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().
			From(Table_Link_Name).
			Where(map[string]interface{}{
				"user_id": userId,
			}).
			Limit(limit, number).
			OrderBy("link_id", "DESC"))
	if err != nil {
		return
	}
	links = rs.Rows()

	return
}

// get all links
func (l *Link) GetLinks() (links []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().From(Table_Link_Name))
	if err != nil {
		return
	}
	links = rs.Rows()
	return
}

// get all links by sequence
func (l *Link) GetLinksOrderBySequence(userId string) (groups []map[string]interface{}, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().From(Table_Link_Name).Where(map[string]interface{}{
			"user_id": userId,
		}).OrderBy("sequence", "ASC"))
	if err != nil {
		return
	}

	//var links = []map[string]string
	var links = rs.Rows()

    // 封装为string，link格式
	var mapLink = make(map[string][]map[string]string)
	for i := 0; i < len(links); i++ {
        var link = links[i]
        var groupName = link ["group_name"]
        mapLink [groupName] = append(mapLink[groupName], link)
    }

    var keys []string
    for k := range mapLink {
        keys = append(keys, k)
    }

    sort.Strings(keys)
    
    // 循环map，封装为数组
    for _, k := range keys {
        var groupLink = make(map[string]interface{})
        groupLink ["groupName"] = k
        groupLink ["links"] = mapLink[k]
        groups = append(groups, groupLink)
    }

	return
}

// get link count
func (l *Link) CountLinks(userId string) (count int64, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().
		Select("count(*) as total").
		From(Table_Link_Name).
		Where(map[string]interface{}{
			"user_id": userId,
		}))
	if err != nil {
		return
	}
	count = utils.NewConvert().StringToInt64(rs.Value("total"))
	return
}

// get link count by keyword
func (l *Link) CountLinksByKeyword(keyword string, userId string) (count int64, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().
		Select("count(*) as total").
		From(Table_Link_Name).
		Where(map[string]interface{}{
			"name LIKE": "%" + keyword + "%",
			"user_id":   userId,
		}))
	if err != nil {
		return
	}
	count = utils.NewConvert().StringToInt64(rs.Value("total"))
	return
}

// get links by like name
func (l *Link) GetLinksByLikeName(name string) (links []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Link_Name).Where(map[string]interface{}{
		"name Like": "%" + name + "%",
	}).Limit(0, 1))
	if err != nil {
		return
	}
	links = rs.Rows()
	return
}

// get link by many link_id
func (l *Link) GetLinkByLinkIds(linkIds []string) (links []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Link_Name).Where(map[string]interface{}{
		"link_id": linkIds,
	}))
	if err != nil {
		return
	}
	links = rs.Rows()
	return
}
