package xmindgo

import (
	"github.com/xiazemin/xmind-go/xmind"
	"github.com/xiazemin/xmind-go/xxml"
)

/**
path 保存数据的路径
data := `[{"node_id":"1","topic_content":"main topic"},
{"node_id":"2","topic_content":"topic1","parent_id":"1"},{"node_id":"3","topic_content":"topic2","parent_id":"1"},
{"node_id":"4","topic_content":"topic3","parent_id":"2"},{"node_id":"5","topic_content":"topic4","parent_id":"2"},
{"node_id":"6","topic_content":"topic5","parent_id":"3"},{"node_id":"7","topic_content":"topic6","parent_id":"3"},
{"node_id":"8","topic_content":"topic8","parent_id":"7"}
]`
这里定义 node_id 表示节点id, topic_content 表示主题内容, parent_id 表示父节点id
*/
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
	// xmlData, err := xml.Marshal(st)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(xmlData))
	xxml.AddProperty(st)
	return xxml.SaveSheets(path, st)
}
