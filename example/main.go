package main

import (
	xmlXmind "github.com/xiazemin/xmind-go"
	"github.com/xiazemin/xmind-go/xjson"
)

func main() {
	data := `[{"node_id":"1","topic_content":"main topic"},
{"node_id":"2","topic_content":"topic1","parent_id":"1"},{"node_id":"3","topic_content":"topic2","parent_id":"1"},
{"node_id":"4","topic_content":"topic3","parent_id":"2"},{"node_id":"5","topic_content":"topic4","parent_id":"2"},
{"node_id":"6","topic_content":"topic5","parent_id":"3"},{"node_id":"7","topic_content":"topic6","parent_id":"3"},
{"node_id":"8","topic_content":"topic8","parent_id":"7"}
]`
	// 这里定义 node_id 表示节点id, topic_content 表示主题内容, parent_id 表示父节点id
	// 传入定好的json字符串,以及指定好json的key字符串就可以将任意json数据转换成xmind
	// 也可用用 data := []byte(`{}`) 传入字节数组

	err := xjson.SaveSheets("./example/custom.xmind", data)
	if err != nil {
		panic(err)
	}

	err = xmlXmind.SaveSheets("./example/example.xmind", data)
	if err != nil {
		panic(err)
	}
}
