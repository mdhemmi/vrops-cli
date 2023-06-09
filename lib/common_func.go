package vrops_cli

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/TylerBrock/colorjson"
	"github.com/sfreiberg/simplessh"
)

func WalkMatch(root, pattern string) ([]string, error) {
	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}

// ==============================================================================================================================

func Untar(sourcefile string, target string) {

	file, err := os.Open(sourcefile)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	var fileReader io.ReadCloser = file

	// just in case we are reading a tar.gz file, add a filter to handle gzipped file
	if strings.HasSuffix(sourcefile, ".gz") {
		if fileReader, err = gzip.NewReader(file); err != nil {

			fmt.Println(err)
			os.Exit(1)
		}
		defer fileReader.Close()
	}

	tarBallReader := tar.NewReader(fileReader)

	// Extracting tarred files

	for {
		header, err := tarBallReader.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}

		// get the individual filename and extract to the current directory
		filename := header.Name

		switch header.Typeflag {
		case tar.TypeDir:
			// handle directory
			fmt.Println("Creating directory :", filename)
			err = os.MkdirAll(filename, os.FileMode(header.Mode)) // or use 0755 if you prefer

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

		case tar.TypeReg:
			// handle normal file
			//fmt.Println("Untarring :", filename)
			writer, err := os.Create(filename)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			io.Copy(writer, tarBallReader)

			err = os.Chmod(filename, os.FileMode(header.Mode))

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			writer.Close()
		default:
			fmt.Printf("Unable to untar type : %c in file %s", header.Typeflag, filename)
		}
	}
}

// ==============================================================================================================================

func Between(value string, a string, b string) string {
	// Get substring between two strings.
	posFirst := strings.Index(value, a)
	if posFirst == -1 {
		return ""
	}
	posLast := strings.Index(value, b)
	if posLast == -1 {
		return ""
	}
	posFirstAdjusted := posFirst + len(a)
	if posFirstAdjusted >= posLast {
		return ""
	}
	return value[posFirstAdjusted:posLast]
}

// ==============================================================================================================================

func Check_Service_Availability(target string, port int) bool {
	//fmt.Print("Check if Target Service is available: ")
	checkTimeout := 5
	serverAddress := net.JoinHostPort(target, strconv.Itoa(port))
	timeout := time.Second * time.Duration(checkTimeout)
	tcpConn, tcpErr := net.DialTimeout("tcp", serverAddress, timeout)
	tcpResult := "FAIL"
	if tcpErr == nil {
		tcpResult = "OK"
	}
	if tcpResult == "OK" {
		tcpConn.Close()
		//fmt.Println("OK")
		//fmt.Println("")
		return true
	} else {
		//fmt.Println("FAIL")
		return false
	}
}

// ==============================================================================================================================

type SlackRequestBody struct {
	Text string `json:"text"`
}

func SendSlackNotification(webhookUrl string, msg string) error {

	slackBody, _ := json.Marshal(SlackRequestBody{Text: msg})
	req, err := http.NewRequest(http.MethodPost, webhookUrl, bytes.NewBuffer(slackBody))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	if buf.String() != "ok" {
		return errors.New("Non-ok response returned from Slack")
	}
	return nil
}

func Ssh_remotecommand(hostname string, username string, password string, command string) string {
	var client *simplessh.Client
	var err error
	var output string

	if client, err = simplessh.ConnectWithPassword(hostname, username, password); err != nil {
		fmt.Println(err)
	}

	defer client.Close()

	// Now run the commands on the remote machine:
	bytes, err := client.Exec(command)
	if err != nil {
		log.Println(err)
	}
	output = string(bytes)
	return output
}

// ==============================================================================================================================

func GetCurrentFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
	//return fmt.Sprintf("%s", runtime.FuncForPC(pc).Name())
}

// ==============================================================================================================================

func Debug(data string, kind string, task string) {
	switch task {
	case "start":
		fmt.Println("")
		fmt.Println("DEBUG OUTPUT START")
	case "print":
		if kind == "json" {
			var obj map[string]interface{}
			json.Unmarshal([]byte(data), &obj)
			f := colorjson.NewFormatter()
			f.Indent = 4
			s, _ := f.Marshal(obj)
			//s, _ := colorjson.Marshal(obj)
			fmt.Println(string(s))
		} else {
			fmt.Println(data)
		}
	case "end":
		fmt.Println("")
		fmt.Println("DEBUG OUTPUT END")
		fmt.Println("")
	}
}

// ==============================================================================================================================

func Get_Username() string {
	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}

	username := user.Username
	return username
}

// ==============================================================================================================================

func IsFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

// ==============================================================================================================================
var (
	suffixes [5]string
)

func Round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

func HumanFileSize(size float64) string {
	//fmt.Println(size)
	suffixes[0] = "B"
	suffixes[1] = "KB"
	suffixes[2] = "MB"
	suffixes[3] = "GB"
	suffixes[4] = "TB"

	base := math.Log(size) / math.Log(1024)
	getSize := Round(math.Pow(1024, base-math.Floor(base)), .5, 2)
	//fmt.Println(int(math.Floor(base)))
	getSuffix := suffixes[int(math.Floor(base))]
	return strconv.FormatFloat(getSize, 'f', -1, 64) + " " + string(getSuffix)
}
