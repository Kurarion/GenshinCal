package genshindata

import (
	"strconv"
)

//取得角色
func GetAvatar(id interface{}) *Avatar {

	var targeID uint64
	switch v := id.(type) {
	case string:
		targeID, _ = strconv.ParseUint(v, 10, 64)
	case uint64:
		targeID = v
	case int:
		targeID = uint64(v)
	}

	return avatar[targeID]
}

//取得角色
func GetAvatarByName(name string) *Avatar {
	targeID, ok := avatarNameMap[name]
	if !ok {
		return nil
	}
	return GetAvatar(targeID)
}
