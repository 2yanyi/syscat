package cat

import (
	"github.com/json-iterator/go"
	"github.com/matsuwin/syscat/cat/internal"
	"os"
)

func SizeFormat(bytes float64) string { return internal.SizeFormat(bytes) }

func Stderr(err string) { internal.Stderr(err) }

func String(fp string) string { return internal.String(&fp) }

func Bytes(fp string) []byte { return internal.Bytes(&fp) }

func Wcl(fp string) int { return internal.Wcl(&fp) }

func MD5sumChunked(fp string) (os.FileInfo, string, error) { return internal.MD5sumChunked(&fp) }

func FileExist(fp string) bool { return internal.FileExist(fp) }

func CommandArgs(dir string, args []string) string { return internal.Commandline(dir, args) }

func Json(a interface{}) []byte { return internal.Json(a) }

// extends

func JsonIter() jsoniter.API { return jsoniter.ConfigFastest }

func Syscat() *internal.Environment { return internal.Syscat() }

func Sysctl(action, name string) error { return internal.Sysctl(action, name) }
