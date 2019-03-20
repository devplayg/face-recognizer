// gocv.io sample
//
// What it does:
//
// This example uses the CascadeClassifier class to detect faces,
// and draw a rectangle around each of them, before displaying them within a Window.
//
// How to run:
//
// facedetect [camera ID] [classifier XML file]
//
// 		go run ./cmd/facedetect/main.go 0 data/haarcascade_frontalface_default.xml
//
// +build example
package main

import (
	"fmt"
	"os"
	"gocv.io/x/gocv"
	ts "github.com/tensorflow/tensorflow/tensorflow/go"
)
var faceDb map[string]*gocv.Mat

func init() {
	//faceDb := make(map[string]*gocv.Mat)
	//dir := "img/"
	//files, err := ioutil.ReadDir(dir)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//for _, file := range files {
	//	name := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
	//	imgPath := filepath.Join(dir, name)
	//	fmt.Printf("%s\n", imgPath)
	//
	//	img:= gocv.IMRead(imgPath, gocv.IMReadColor)

		//gocv.im
		//readImage(&img)
		//gocv.IMRead()
	//}
	//for file in glob.glob("images/*"):
	//person_name = os.path.splitext(os.path.basename(file))[0]
	//image_file = cv2.imread(file, 1)
	//input_embeddings[person_name] = image_to_embedding(image_file, model)
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("How to run:\n\tfacedetect [camera ID]")
		return
	}

	// parse args
	deviceID := os.Args[1]
	modelFile := os.Args[2]


	// open webcam
	webcam, err := gocv.OpenVideoCapture(deviceID)
	if err != nil {
		fmt.Printf("error opening video capture device: %v\n", deviceID)
		return
	}
	defer webcam.Close()

	// open display window
	//window := gocv.NewWindow("Face Detect")
	//defer window.Close()

	// prepare image matrix
	img := gocv.NewMat()
	defer img.Close()

	//model := filepath.Join("../", "lib", "20170511-185253.pb")
	if _, err := os.Stat(modelFile); os.IsNotExist(err) {
		fmt.Printf("File not found: %s\n", modelFile)
		return
	}

	tensorflow.Bool
	ts.LoadSavedModel()


	//net := gocv.ReadNetFromTensorflow(modelFile)
	//if net.Empty() {
	//	fmt.Printf("Error reading network model : %v\n", modelFile)
	//	return
	//}

	//model := filepath.Join()
	//net := gocv.ReadNet(model, "")
	//if net.Empty() {
	//	fmt.Printf("Error reading network model : %v\n", model)
	//	return
	//}
	// color for the rect when faces detected
	//blue := color.RGBA{0, 0, 255, 0}

	// load classifier to recognize faces
	//classifier := gocv.NewCascadeClassifier()
	//defer classifier.Close()
	//
	//cf := "../lib/classifier.pk"
	//if !classifier.Load(cf) {
	//	fmt.Printf("Error reading cascade file: %v\n", cf)
	//	return
	//}
	//
	////println("hello")
	//fmt.Printf("Start reading device: %v\n", deviceID)
	//for {
	//	if ok := webcam.Read(&img); !ok {
	//		fmt.Printf("Device closed: %v\n", deviceID)
	//		return
	//	}
	//	if img.Empty() {
	//		continue
	//	}
	//
	//	// detect faces
	//	rects := classifier.DetectMultiScale(img)
	//	fmt.Printf("found %d faces\n", len(rects))
	//
	//	for _, r := range rects {
	//		gocv.Rectangle(&img, r, blue, 3)
	//
	//		// Recognize
	//		//identity = recognize_face(face_image, input_embeddings, model)
	//		//recognizeFace
	//		//size := gocv.GetTextSize("Human", gocv.FontHersheyPlain, 1.2, 2)
	//		//pt := image.Pt(r.Min.X+(r.Min.X/2)-(size.X/2), r.Min.Y-2)
	//		//gocv.PutText(&img, "Human", pt, gocv.FontHersheyPlain, 1.2, blue, 2)
	//	}
	//
	//	// show the image in the window, and wait 1 millisecond
	//	window.IMShow(img)
	//	if window.WaitKey(1) >= 0 {
	//		break
	//	}
	//}
}
