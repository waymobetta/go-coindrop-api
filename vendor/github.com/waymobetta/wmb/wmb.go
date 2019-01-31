// package for random helper utilities

package wmb

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime/debug"
	"time"

	"github.com/fatih/color"
	log "github.com/sirupsen/logrus"
)

// Clear clears the terminal screen
func Clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

// Used to time functions
func Elapsed(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

// JSONToString converts JSON file to string; returns string and error
func JSONToString(fileName string) (string, error) {
	raw, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("[!] Error encountered when reading JSON file\n", err)
	}
	return string(raw), nil
}

// TermNotify notifies terminal with data; returns error
func TermNotify(data string) error {
	if err := exec.Command("terminal-notifier", "-message", data).Run(); err != nil {
		log.Fatal("[!] Error encountered when attemping to notify user via TermNotify\n", err)
	}
	return nil
}

// ReadFileByLine reads a file line-by-line; returns array, length of array, and error
func ReadFileByLine(path string, data []string) ([]string, int, error) {
	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		log.Fatal("[!] Error encountered when opening file\n", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data, len(data), nil
}

// WriteFile writes string (data) to a file, delimited in some fashion; returns error
func WriteFile(path string, data interface{}) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatal("[!] Error encountered when creating file\n", err)
	}
	defer file.Close()
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("[!] Error encountered when opening file\n", err)
	}
	defer f.Close()
	if _, err = f.WriteString(data.(string)); err != nil {
		log.Fatal("[!] Error encountered when writing to file\n", err)
	}
}

// Require is a mirror of Require in Solidity; rather than reverting a transaction, this Require will panic loudly and print stacktrace if arguments do not match
func Require(a interface{}, b interface{}) {
	if a != b {
		fmt.Println("")
		color.Red("[!] Execution halted: requirement not fulfilled!\nDetails: %v != %v\n\n", a, b)
		debug.PrintStack()
		os.Exit(0)
	}
}
