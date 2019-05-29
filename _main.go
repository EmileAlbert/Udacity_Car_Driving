package main

import(
  "fmt"
  "./lib"
  "encoding/base64"
  "image"
  _"image/jpeg"
  "image/draw"
  "strings"
  "github.com/google/gxui"
  "github.com/google/gxui/drivers/gl"
  "github.com/google/gxui/themes/dark"
)

func appMain(driver gxui.Driver) {
    newImage,_,_ := image.Decode(base64.NewDecoder(base64.StdEncoding, strings.NewReader(cp.TESTDATA)))
    //Image := GetImageFeature(Image)
    //newImage := resize.Resize(320, 160, midImage, resize.NearestNeighbor)
    m := image.NewRGBA(newImage.Bounds())
    draw.Draw(m, m.Bounds(), newImage, image.ZP, draw.Src)
    // The themes create the content. Currently only a dark theme is offered for GUI elements.
    theme := dark.CreateTheme(driver)
    img := theme.CreateImage()
    window := theme.CreateWindow(newImage.Bounds().Max.X, newImage.Bounds().Max.Y, "Image viewer")
    texture := driver.CreateTexture(m, 1.0)
    img.SetTexture(texture)
    window.AddChild(img)
    window.OnClose(driver.Terminate)
}

func main(){
  //res := cp.GetDataFromCSV("C:/data/driving_log.csv")
  //fmt.Println(res[56].Speed)

  // reader := base64.NewDecoder(base64.StdEncoding,  strings.NewReader(cp.TESTDATA))
  // Image,_,_ := image.Decode(reader)
  // fmt.Println(Image)

  gl.StartDriver(appMain)

}
