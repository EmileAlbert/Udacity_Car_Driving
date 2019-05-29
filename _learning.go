package main

import (
  "./lib"
  //"fmt"
  "image"
  "image/png"
  "os"
)



func main() {
    traindata := cp.GetDataFromCSV("C:\\Udacity\\data\\driving_log.csv")
    infile, err := os.Open(traindata[0].ImagePath)
    if err != nil {
        panic(err)
    }
    defer infile.Close()
    src, _, err := image.Decode(infile)
    if err != nil {
      // replace this with real error handling
      panic(err)
    }
    //Grayscale
    processed_img := cp.GetImageFeature(src)
    //processed_img := cp.GetImageFeature(src,320,160,1,0.0,0.0)
    //fmt.Println(processed_img)

    // Encode the grayscale image to the output file
    outfile, err := os.Create("test.png")
    if err != nil {
    panic(err)
    }
    defer outfile.Close()
    png.Encode(outfile, processed_img)
    }
