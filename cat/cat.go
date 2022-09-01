package cat

import (
	"github.com/json-iterator/go"
	"github.com/matsuwin/syscat/cat/internal"
	"os"
)

func SizeFormat(bytes float64) string { return internal.SizeFormat(bytes) }

func Stderr(err string) { internal.Stderr(err) }

func String(fp string) string { return internal.String(fp) }

func Bytes(fp string) []byte { return internal.Bytes(fp) }

func Wcl(fp string) int { return internal.Wcl(fp) }

func MD5sumChunked(fp string) (os.FileInfo, string, error) { return internal.MD5sumChunked(fp) }

func FileExist(fp string) bool { return internal.FileExist(fp) }

func RandomChars(n int) []byte { return internal.RandomChars(n) }

func Commandline(dir string, args []string) string { return internal.Commandline(dir, args) }

func BashC(dir, sh string) string { return internal.BashC(dir, sh) }

func JsonFormat(a interface{}) []byte { return internal.JsonFormat(a) }

// extends

func Sysctl(action, name string) error { return internal.Sysctl(action, name) }

func Syscat() *internal.Environment { return internal.Syscat() }

var Json = jsoniter.ConfigFastest
