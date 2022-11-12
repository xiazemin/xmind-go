给定父子结构关系json，生成对应的xmind文件

path 保存数据的路径
data := `[{"node_id":"1","topic_content":"main topic"},
{"node_id":"2","topic_content":"topic1","parent_id":"1"},{"node_id":"3","topic_content":"topic2","parent_id":"1"},
{"node_id":"4","topic_content":"topic3","parent_id":"2"},{"node_id":"5","topic_content":"topic4","parent_id":"2"},
{"node_id":"6","topic_content":"topic5","parent_id":"3"},{"node_id":"7","topic_content":"topic6","parent_id":"3"},
{"node_id":"8","topic_content":"topic8","parent_id":"7"}
]`
这里定义 node_id 表示节点id, topic_content 表示主题内容, parent_id 表示父节点id

支持json格式和xml格式，兼容高版本的xmind，常用版本的xmind都能打开
使用
```
go install github.com/xiazemin/xmind-go
```