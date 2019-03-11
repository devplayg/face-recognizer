package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"image"
	"strconv"
)

var (
	dir            = "C:/Temp/pics/jisoo"
	classifierFile = "haarcascade_frontalface_default.xml"
)

func main() {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	img := gocv.NewMat()
	defer img.Close()

	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	if !classifier.Load(classifierFile) {
		fmt.Printf("Error reading cascade file: %v\n", classifierFile)
		return
	}

	for _, f := range files {
		imgPath := filepath.Join(dir, f.Name())
		img := gocv.IMRead(imgPath, gocv.IMReadColor)
		rects := classifier.DetectMultiScale(img)
		fmt.Printf("%s: found %d\n", f.Name(), len(rects))
		fileExt := filepath.Ext(f.Name())
		fileName :=  strings.TrimSuffix(f.Name(), filepath.Ext(f.Name()))
		filePrefix := filepath.Join(dir, "crop_"+fileName )

		for idx, r := range rects {
			if r.Size().X > 95 {
				fmt.Printf("	    -> x: %d, y: %d\n", r.Size().X, r.Size().Y)
				subImage := img.Region(r)
				resizedImage := gocv.NewMat()
				gocv.Resize(subImage, &resizedImage, image.Point{96, 96}, 1, 1, gocv.InterpolationArea)
				println("    " + f.Name() + "_" + strconv.Itoa(idx) + fileExt)
				println(gocv.IMWrite(filePrefix+"_"+strconv.Itoa(idx)+fileExt, resizedImage))
			}
		}
	}
}
