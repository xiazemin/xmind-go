package xjson

import "xmind-go/xmind"

func SaveSheets(path string, data string) error {
	/*
		idKey: 以该json tag字段作为主题ID
		titleKey: 以该json tag字段作为主题内容
		parentKey: 以该json tag字段作为判断父节点的依据
		isRootKey: 以该json tag字段,该字段为bool类型,true表示根节点,false表示普通节点
	*/
	st, err := xmind.LoadCustom(data, "node_id", "topic_content", "parent_id", "false")
	if err != nil {
		panic(err)
	}
	err = xmind.SaveSheets(path, st)
	if err != nil {
		panic(err)
	}
	return err
}
