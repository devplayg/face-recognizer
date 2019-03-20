package main

import (
	"fmt"
	"github.com/devplayg/face-recognizer"
	"gocv.io/x/gocv"
	"log"
	"runtime"
	"os"
	"net"
	"bufio"
)

var (
	webcam *gocv.VideoCapture
)

func main() {
	// CPU 설정
	runtime.GOMAXPROCS(runtime.NumCPU())

	var (
		deviceID = face_recognizer.CmdFlags.Int("d", 0, "Device ID")
		addr     = face_recognizer.CmdFlags.String("addr", "", "Server address")
	)
	face_recognizer.CmdFlags.Usage = face_recognizer.PrintHelp
	face_recognizer.CmdFlags.Parse(os.Args[1:])

	log.Print(*deviceID)
	log.Print("asdfasdf")
	log.Println(*addr)
	if len(*addr) < 1 {
		printHelp()
		return
	}

	// 디바이스 열기
	webcam, err := gocv.OpenVideoCapture(deviceID)
	if err != nil {
		fmt.Printf("Error opening capture device: %v\n", deviceID)
		return
	}
	defer webcam.Close()

	// 스트림 생성
	//stream = mjpeg.NewStream()

	// 서버 연결

	conn, err := net.Dial("tcp", *addr)
	if err != nil {
		log.Println(err)
		return
	}

	for {
	}

	// 전송
	go send()

	face_recognizer.WaitForSignals()

}

func send() {
	img := gocv.NewMat()
	defer img.Close()

	for {
		if ok := webcam.Read(&img); !ok {
			//fmt.Printf("Device closed: %v\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}

		//encoder := gob.NewEncoder(conn)
		//encoder.Encode(studentEncode)
		//buf, _ := gocv.IMEncode(".jpg", img)
		//stream.UpdateJPEG(buf)
	}
}

func printHelp() {
	face_recognizer.CmdFlags.PrintDefaults()
}
