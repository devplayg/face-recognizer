package main

import (
	"github.com/devplayg/face-recognizer"
	"gocv.io/x/gocv"
	"log"
	"runtime"
	"os"
	"net"
	"encoding/gob"
	"time"
)

var (
	webcam *gocv.VideoCapture
	//conn net.Conn
	deviceID *int
)

type Message struct {
	Text string
	Seq int
}

func main() {
	// CPU 설정
	runtime.GOMAXPROCS(runtime.NumCPU())

	deviceID = face_recognizer.CmdFlags.Int("d", 0, "Device ID")
	var (
		addr     = face_recognizer.CmdFlags.String("addr", "127.0.0.1:8000", "Server address")
	)
	face_recognizer.CmdFlags.Usage = face_recognizer.PrintHelp
	face_recognizer.CmdFlags.Parse(os.Args[1:])

	if len(*addr) < 1 {
		printHelp()
		return
	}

	// 카메라 디바이스 접속
	//webcam, err := gocv.OpenVideoCapture(*deviceID)
	//if err != nil {
	//	fmt.Printf("Error opening capture device: %v, %s\n", deviceID, err)
	//	return
	//}
	//defer webcam.Close()

	// 서버 연결
	conn, err := net.Dial("tcp", *addr)
	if err != nil {
		log.Print(err)
		return
	}

	// 데이터 전송 시작
	go send(conn)
	defer conn.Close()

	face_recognizer.WaitForSignals()
}

func send(conn net.Conn) {
	img := gocv.NewMat()
	defer img.Close()

	seq := 0

	for {
		//if ok := webcam.Read(&img); !ok {
		//	fmt.Printf("Device closed: %v\n", deviceID)
		//	return
		//}
		//if img.Empty() {
		//	continue
		//}

		encoder := gob.NewEncoder(conn)

		msg := Message{"msg", seq}
		if err := encoder.Encode(msg); err != nil {
			log.Print(err)
			return
		}
		seq++
		time.Sleep(1*time.Millisecond)
	}
}

func printHelp() {
	face_recognizer.CmdFlags.PrintDefaults()
}
