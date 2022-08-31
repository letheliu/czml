package czml

import (
	"fmt"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"image"
	"image/png"
	"os"
)

// Billboard is a billboard, or viewport-aligned image. The billboard is positioned in the scene
// by the position property. A billboard is sometimes called a marker.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Billboard
type Billboard struct {
	Show                       bool                      `json:"show,omitempty"`
	Image                      string                    `json:"image"`
	Scale                      float64                   `json:"scale,omitempty"`
	PixelOffset                *PixelOffset              `json:"pixelOffset,omitempty"`
	EyeOffset                  *EyeOffset                `json:"eyeOffset,omitempty"`
	HorizontalOrigin           *HorizontalOrigin         `json:"horizontalOrigin,omitempty"`
	VerticalOrigin             *VerticalOrigin           `json:"verticalOrigin,omitempty"`
	HeightReference            *HeightReference          `json:"heightReference,omitempty"`
	Color                      *Color                    `json:"color,omitempty"`
	Rotation                   float64                   `json:"rotation,omitempty"`
	AlignedAxis                *AlignedAxis              `json:"alignedAxis,omitempty"`
	SizeInMeters               bool                      `json:"sizeInMeters,omitempty"`
	Width                      float64                   `json:"width,omitempty"`
	Height                     float64                   `json:"height,omitempty"`
	ScaleByDistance            *NearFarScaler            `json:"scaleByDistance,omitempty"`
	TranslucencyByDistance     *NearFarScaler            `json:"translucencyByDistance,omitempty"`
	PixelOffsetScaleByDistance *NearFarScaler            `json:"pixelOffsetScaleByDistance,omitempty"`
	ImageSubRegion             *BoundingRectangle        `json:"imageSubRegion,omitempty"`
	DistanceDisplayCondition   *DistanceDisplayCondition `json:"distanceDisplayCondition,omitempty"`
	DisableDepthTestDistance   float64                   `json:"disableDepthTestDistance,omitempty"`
}

var (
	fontKai *truetype.Font
	fontTtf *truetype.Font
)

func GenerateImgFromText(packetId string, text string, fontPath string, savePath string) (url string, err error) {
	// 创建画布
	newTemplateImage := image.NewRGBA(image.Rect(0, 0, 800, 600))

	// 加载字体
	fontKai, err = loadFont(fontPath)
	if err != nil {
		return
	}

	content := freetype.NewContext()
	content.SetClip(newTemplateImage.Bounds())
	content.SetDst(newTemplateImage)
	content.SetSrc(image.Black)
	content.SetDPI(72)

	content.SetFontSize(80)
	content.SetFont(fontKai)

	content.DrawString(text, freetype.Pt(10, 300))
	url, err = saveFile(packetId, newTemplateImage, savePath)

	return
}

// 根据路径加载字体文件
// path 字体的路径
func loadFont(path string) (font *truetype.Font, err error) {
	var fontBytes []byte
	fontBytes, err = os.ReadFile(path) // 读取字体文件
	if err != nil {
		err = fmt.Errorf("加载字体文件出错:%s", err.Error())
		return
	}
	font, err = freetype.ParseFont(fontBytes) // 解析字体文件
	if err != nil {
		err = fmt.Errorf("解析字体文件出错,%s", err.Error())
		return
	}
	return
}

func saveFile(packetId string, pic *image.RGBA, savePath string) (url string, err error) {
	url = savePath + "/" + packetId + ".png"
	dstFile, err := os.Create(url)
	if err != nil {
		fmt.Println(err)
	}
	defer dstFile.Close()
	png.Encode(dstFile, pic)

	return
}
