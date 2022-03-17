package util

import (
	"reflect"
)

type TreeList struct {
	ID       uint64
	Pid      uint64
	Name     string
	Children []TreeList
}

// GenerateTree 无限极分类
func GenerateTree(list interface{}, pid uint64) []TreeList {
	var treeList = make([]TreeList, 0)
	sv := reflect.ValueOf(list)
	svs := sv.Slice(0, sv.Len())
	for i := 0; i < svs.Len(); i++ {
		e := svs.Index(i)
		currentPid := e.FieldByName("ParentId").Uint()
		currentId := e.FieldByName("Id").Uint()
		if currentPid == pid {
			child := GenerateTree(list, currentId)
			node := TreeList{
				ID:   currentId,
				Name: e.FieldByName("Name").String(),
				Pid:  currentPid,
			}
			node.Children = child
			treeList = append(treeList, node)
		}
	}
	return treeList
}
