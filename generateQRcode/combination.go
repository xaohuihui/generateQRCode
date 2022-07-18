package generateQRcode

import (
	"github.com/golang/freetype"
	"golang.org/x/image/font"
	"image"
	"image/draw"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"
	"unicode/utf8"
)

const SquareHeight = 300

func AddTextToImg(imgText, outImg string) {
	var dpi float64 = 72
	var fontFile = "msyh.ttf"
	var hinting = "none"
	var fontSize float64 = 12
	var spacing = 1.5
	var wonb = false

	//读取字体
	fontBytes, err := ioutil.ReadFile(fontFile)
	if err != nil {
		log.Println(err)
		return
	}
	//解析字体
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
		return
	}
	// 初始化图片背景
	fg := image.Black
	if wonb {
		fg = image.White
	}
	//初始化一张图片,生成原图
	imgB, _ := os.Open(outImg)
	img, _ := jpeg.Decode(imgB)
	defer imgB.Close()
	b := img.Bounds()
	rgba := image.NewNRGBA(b)
	draw.Draw(rgba, rgba.Bounds(), img, image.ZP, draw.Src)
	//在图片上面添加文字
	c := freetype.NewContext()
	c.SetDPI(dpi)
	//设置字体
	c.SetFont(f)
	//设置大小
	c.SetFontSize(fontSize)
	//设置边界
	c.SetClip(rgba.Bounds())
	//设置背景底图
	c.SetDst(rgba)
	//设置背景图
	c.SetSrc(fg)
	//设置提示
	switch hinting {
	default:
		c.SetHinting(font.HintingNone)
	case "full":
		c.SetHinting(font.HintingFull)
	}
	// 获取字体的尺寸大小
	fixed := c.PointToFixed(fontSize)
	// 画文字
	//pt := freetype.Pt(10, 10+int(c.PointToFixed(imgSize)>>6))
	pt := freetype.Pt(rgba.Rect.Max.X/2-(utf8.RuneCountInString(imgText)/2)*fixed.Ceil(),
		rgba.Rect.Max.Y-SquareHeight+SquareHeight/2+140)

	_, err = c.DrawString(imgText, pt)
	if err != nil {
		log.Println(err)
		return
	}
	pt.Y += c.PointToFixed(fontSize * spacing)

	imgw, _ := os.Create(outImg)
	jpeg.Encode(imgw, rgba, &jpeg.Options{100})
	defer imgw.Close()
}
