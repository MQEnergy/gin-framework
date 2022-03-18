package util

import (
	"reflect"
)

type TreeList struct {
	ID        uint64     `json:"id"`
	Pid       uint64     `json:"parent_id"`
	Name      string     `json:"name"`
	IsChecked bool       `json:"is_checked"`
	Children  []TreeList `json:"children"`
}

// GenerateTree 无限极分类
func GenerateTree(list interface{}, pid uint64) []TreeList {
	var treeList = make([]TreeList, 0)
	sv := reflect.ValueOf(list)
	svs := sv.Slice(0, sv.Len())
	for i := 0; i < svs.Len(); i++ {
		e := svs.Index(i)
		var currentPid, currentId uint64
		var IsChecked bool
		if e.Kind() == reflect.Struct {
			IsChecked = e.FieldByName("IsChecked").Bool()
			currentPid = e.FieldByName("ParentId").Uint()
			currentId = e.FieldByName("Id").Uint()
		} else {
			IsChecked = e.Elem().FieldByName("IsChecked").Bool()
			currentPid = e.Elem().FieldByName("ParentId").Uint()
			currentId = e.Elem().FieldByName("Id").Uint()
		}
		if currentPid == pid {
			child := GenerateTree(list, currentId)
			node := TreeList{
				ID:        currentId,
				Name:      e.FieldByName("Name").String(),
				Pid:       currentPid,
				IsChecked: IsChecked,
			}
			node.Children = child
			treeList = append(treeList, node)
		}
	}
	return treeList
}
