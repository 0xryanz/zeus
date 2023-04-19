package command

import (
	"bufio"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func parse(filename string, platform int) {
	absPath, _ := filepath.Abs(filename)
	file, err := os.Open(absPath)
	commands := []Command{}

	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var tmpExeDir string
	for scanner.Scan() {
		input := scanner.Text()
		subs := strings.Split(input, " ")
		command := Command{}
		if strings.HasPrefix(subs[0], "+") {
			command.IsDir = true
			subs = regexp.MustCompile("[\\+\\,\\ \\s]+").Split(input, -1)
			tmpExeDir = strings.Join(subs, "")
		}
		command.ExecDir = tmpExeDir
		command.Subs = subs
		commands = append(commands, command)
	}
	PlatformCommands[platform] = commands
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

const (
	ANROID = iota + 1
	IOS
	WINDOWS
	MAC
	UNKNOWN
)

type Command struct {
	Subs    []string
	IsDir   bool
	ExecDir string
}

var (
	PlatformCommands = make(map[int][]Command)
)

func init() {
	defaultFiles := make([]interface{}, 0)
	defaultFiles = append(defaultFiles,
		"./static/android",
		ANROID,
		"./static/ios",
		IOS,
		"./static/windows",
		WINDOWS,
		"./static/mac",
		MAC,
	)
	for i, filename := range defaultFiles {
		if i%2 != 1 {
			filename, _ := filename.(string)
			platform, _ := defaultFiles[i+1].(int)
			parse(filename, platform)
		}
	}
}
