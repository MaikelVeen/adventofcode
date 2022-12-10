package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type File struct {
	Name   string
	Size   int
	Parent *Directory
}

type Directory struct {
	Name   string
	Files  map[string]*File
	Dirs   map[string]*Directory
	Parent *Directory
}

func (d *Directory) Size() int {
	size := 0
	for _, file := range d.Files {
		size += file.Size
	}

	for _, child := range d.Dirs {
		size += child.Size()
	}

	return size
}

type FeedLineType int8
type CommandType int8
type OutputType int8

const (
	UnknownFeedLineType FeedLineType = iota
	Command
	Output
)

const (
	UnknownCommandType CommandType = iota
	ChangeDir
	List
)

const (
	UnkownnOutputType OutputType = iota
	DirectoryOut
	FileOut
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func parse(feed []string) *Directory {
	root := &Directory{
		Name:  "/",
		Files: make(map[string]*File),
		Dirs:  make(map[string]*Directory),
	}

	cursor := root

	for _, line := range feed {
		typ := parseFeedLineType(line)

		if typ == Command {
			cmd, arg := parseCommand(line)

			switch cmd {
			case ChangeDir:
				switch arg {
				case "/":
					cursor = root
				case "..":
					cursor = cursor.Parent
				default:
					dir, ok := cursor.Dirs[arg]
					if !ok {
						panic(fmt.Sprintf("dir %s does not exist", arg))
					}

					cursor = dir
				}
				continue
			case List:
				continue
			default:
				panic("could not parse command type")
			}
		}

		if typ == Output {
			parts := strings.Split(line, " ")
			if parts[0] == "dir" {
				dir := &Directory{
					Name:   parts[1],
					Files:  make(map[string]*File),
					Dirs:   make(map[string]*Directory),
					Parent: cursor,
				}

				cursor.Dirs[dir.Name] = dir
			} else {
				size, err := strconv.Atoi(parts[0])
				if err != nil {
					panic(err)
				}

				file := &File{
					Name:   parts[1],
					Size:   size,
					Parent: cursor,
				}

				cursor.Files[file.Name] = file
			}

		}
	}

	return root
}

func parseFeedLineType(s string) FeedLineType {
	if s[0] == '$' {
		return Command
	} else {
		return Output
	}
}

func parseCommand(s string) (CommandType, string) {
	parts := strings.Split(strings.TrimPrefix(s, "$ "), " ")
	command := parts[0]

	switch command {
	case "cd":
		return ChangeDir, parts[1]
	case "ls":
		return List, ""
	default:
		return UnknownCommandType, ""
	}
}

func sizes(dir *Directory) int {
	if dir == nil {
		return 0
	}

	sum := 0

	if dir.Size() <= 100000 {
		sum += dir.Size()
	}

	for _, childDir := range dir.Dirs {
		sum += sizes(childDir)
	}
	return sum
}

func indexCandidates(dir *Directory, candidates map[*Directory]int, total int) map[*Directory]int {
	if dir == nil {
		return candidates
	}

	fmt.Printf("%s: %d\n", dir.Name, dir.Size())

	size := dir.Size()
	if 70000000-(total-dir.Size()) >= 30000000 {
		candidates[dir] = size
	}

	for _, child := range dir.Dirs {
		indexCandidates(child, candidates, total)
	}

	return candidates
}

func smallestCandidateDelete(dir *Directory) int {
	candidates := indexCandidates(dir, make(map[*Directory]int), dir.Size())

	min := int(^uint(0) >> 1)

	// iterate over the map and update the minimum value if a smaller value is found
	for _, value := range candidates {
		if value < min {
			min = value
		}
	}

	return min
}

func main() {
	lines, err := readLines("input7.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	dir := parse(lines)
	fmt.Printf("answer: %v\n", sizes(dir))

	smallest := smallestCandidateDelete(dir)
	fmt.Println(smallest)
}
