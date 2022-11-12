package xxml

import (
	"archive/zip"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"xmind-go/xmind"
)

type XmindNode struct {
	NodeID       string `json:"node_id"`
	TopicContent string `json:"topic_content"`
	ParentID     string `json:"parent_id,omitempty"`
}

type Properite struct {
	LineColor        string `json:"line-color"`
	FoColor          string `json:"fo:color"`
	ShapeClass       string `json:"shape-class"`
	BorderLineWidth  string `json:"border-line-width"`
	LineClass        string `json:"line-class"`
	LineWidth        string `json:"line-width"`
	FoFontFamily     string `json:"fo:font-family"`
	FoFontStyle      string `json:"fo:font-style"`
	FoFontWeight     int    `json:"fo:font-weight"`
	FoFontSize       string `json:"fo:font-size"`
	FoTextTransform  string `json:"fo:text-transform"`
	FoTextDecoration string `json:"fo:text-decoration"`
	SvgFill          string `json:"svg:fill"`
}

//https://copyfuture.com/blogs-details/202204071548342101

// SaveSheets 保存多个sheet画布到一个xmind文件
func SaveSheets(path string, sheet ...*xmind.Topic) error {
	return (&xmlXmind{WorkBook: &xmind.WorkBook{Topics: sheet}}).Save(path)
}

type xmlXmind struct {
	*xmind.WorkBook
}

func (wk *xmlXmind) check() error {
	if wk == nil || len(wk.Topics) == 0 {
		return errors.New("WorkBook.Topics is null")
	}
	return nil
}

const (
	rootKey xmind.TopicID = "root" // 画布主题地址Key
)

// manifest.json
//{"file-entries":{"content.json":{},"metadata.json":{},"Thumbnails/thumbnail.png":{}}}
type FileEntries struct {
	ContentJSON            interface{} `json:"content.json"`
	MetadataJSON           interface{} `json:"metadata.json"`
	ThumbnailsThumbnailPng interface{} `json:"Thumbnails/thumbnail.png"`
}

type Manifest struct {
	FileEntries FileEntries `json:"file-entries"`
}

//metadata.json
//{"creator":{"name":"Vana","version":"11.0.2.202107130052"},"activeSheetId":"41fa715fe09849fe0aa1b9e514"}
type Creator struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}
type Metadata struct {
	Creator       Creator `json:"creator"`
	ActiveSheetID string  `json:"activeSheetId"`
}

// Save 保存对象为 *.xmind 文件
func (wk *xmlXmind) Save(path string) error {
	err := wk.check()
	if err != nil {
		return err
	}
	if filepath.Ext(path) != ".xmind" {
		return fmt.Errorf("%s: suffix must be .xmind", path)
	}

	cp := make([]*xmind.Topic, 0, len(wk.Topics))
	for _, topic := range wk.Topics {
		// 所有sheet全部切换到根节点,最终使用存入的cp生成xmind文件
		cp = append(cp, topic.On(rootKey))
	}

	fw, err := os.Create(path)
	if err != nil {
		return err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer fw.Close()

	zw := zip.NewWriter(fw)
	//goland:noinspection GoUnhandledErrorResult
	defer zw.Close()

	mi, err := zw.Create(xmind.Manifest)
	if err != nil {
		return err
	}
	if err := json.NewEncoder(mi).Encode(Manifest{
		FileEntries: FileEntries{
			ContentJSON:            struct{}{},
			MetadataJSON:           struct{}{},
			ThumbnailsThumbnailPng: struct{}{},
		},
	}); err != nil {
		return err
	}
	md, err := zw.Create(xmind.Metadata)
	if err != nil {
		return err
	}
	if err := json.NewEncoder(md).Encode(Metadata{
		Creator: Creator{
			Name:    "Vana",
			Version: "11.0.2.202107130052",
		}}); err != nil {
		return err
	}

	wz, err := zw.Create(xmind.ContentJson)
	if err != nil {
		return err
	}
	if err := json.NewEncoder(wz).Encode(cp); err != nil {
		return err
	}
	xz, err := zw.Create(xmind.ContentXml)
	if err != nil {
		return err
	}
	return xml.NewEncoder(xz).Encode(cp)
}

func AddProperty(topic *xmind.Topic) {
	if topic == nil {
		return
	}
	//fmt.Println(topic.Title)
	topic.Style = xmind.Style{
		Id: topic.ID,
		Properties: xmind.Properite{
			LineColor:        "#1414ff",
			FoColor:          "#14ff14",
			ShapeClass:       "org.xmind.topicShape.roundedRect",
			BorderLineWidth:  "1pt",
			LineClass:        "org.xmind.branchConnection.roundedfold",
			LineWidth:        "2pt",
			FoFontFamily:     "NeverMind",
			FoFontStyle:      "normal",
			FoFontWeight:     600,
			FoFontSize:       "14pt",
			FoTextTransform:  "manual",
			FoTextDecoration: "none",
			SvgFill:          "#ffffff",
		},
	}

	if topic.Children != nil {
		for i := range topic.Children.Attached {
			AddProperty(topic.Children.Attached[i])
		}

		for i := range topic.Children.Topics.Topic {
			AddProperty(topic.Children.Topics.Topic[i])
		}

	}

	return
}
