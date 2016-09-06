package command

import (
	"archive/zip"
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/cwen-coder/ljgo/app/util"
	"github.com/qiniu/log"
	"github.com/urfave/cli"
)

var CmdNew = cli.Command{
	Name:   "new",
	Usage:  "provide the webserver which builds and serves the site",
	Action: runNew,
}

func runNew(c *cli.Context) error {
	new(c)
	return nil
}

func new(c *cli.Context) {
	siteName := c.Args()[0]
	if siteName == "" {
		siteName = "default"
	}
	_, err := os.Stat(siteName)
	if err == nil || !os.IsNotExist(err) {
		log.Fatal("Path Exist?!")
	}
	err = os.MkdirAll(siteName, os.ModePerm)
	if err != nil {
		log.Fatalf("mkdir %v: %v", siteName, err)
	}

	decoder := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(util.INIT_ZIP))
	b, _ := ioutil.ReadAll(decoder)

	z, err := zip.NewReader(bytes.NewReader(b), int64(len(b)))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Unpack init content zip")
	for _, zf := range z.File {
		if zf.FileInfo().IsDir() {
			continue
		}
		filename := strings.TrimPrefix(zf.Name, "template/")
		dst := siteName + "/" + filename
		os.MkdirAll(filepath.Dir(dst), os.ModePerm)
		f, err := os.OpenFile(dst, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		rc, err := zf.Open()
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.Copy(f, rc)
		if err != nil {
			log.Fatal(err)
		}
		f.Sync()
		f.Close()
		rc.Close()
	}
	fmt.Println("Done")
}
