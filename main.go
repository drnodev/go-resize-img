package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"
)


func main() {
	files,err := ioutil.ReadDir("demo");
	if err!=nil{
		fmt.Printf("Could not read image the path")
		os.Exit(1)
	}
	for _,file :=range files{
		Resize(file, 700)
	}
}


func Resize(info os.FileInfo, newSize int){

	defer track(time.Now(),info.Name())

	img := gocv.IMRead("demo/"+info.Name(), gocv.IMReadAnyColor)
	if img.Empty() {
		fmt.Printf("Could not read image %s\n", info.Name())
		os.Exit(1)
	}
	matResize  := gocv.NewMat()
	var h float64 = (float64(newSize) / float64(img.Cols())) * float64(img.Rows())
	newMatSize := gocv.NewMatWithSize(newSize,int(h),gocv.MatTypeCV8UC3)
	qualityParm:=[]int{gocv.IMWriteJpegQuality,95}
	gocv.Resize(img,&matResize,image.Pt(newMatSize.Rows(),newMatSize.Cols()),float64(newSize),h,gocv.InterpolationArea)
	gocv.IMWriteWithParams("result/"+(strings.TrimSuffix(info.Name(),path.Ext(info.Name())))+".jpg",matResize,qualityParm)


}



func track(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s : %s\n",name, elapsed)
}

