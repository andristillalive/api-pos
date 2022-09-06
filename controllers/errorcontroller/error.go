package errorcontroller

import (
	"os"
	"strconv"
	"time"
)

var now = time.Now()
var tgl = now.Day()
var bulan = now.Month()
var tahun = now.Year()
var jam = now.Hour()
var menit = now.Minute()
var detik = now.Second()

var namafile = "Error_Log_" + strconv.Itoa(tgl) + "_" + strconv.Itoa(int(bulan)) + "_"

func ErrorLogging(err error) string {
	if _, err1 := os.Stat("/Error"); os.IsNotExist(err1) {
		os.Mkdir("/Error", 0755)
	}
	os.Create(namafile)
	return namafile
}
