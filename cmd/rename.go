package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

var RenameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename files",
	Long:  "Rename files",
	Run:   renameRun,
}

var (
	input     string
	output    string
	startIdx  int
	startTime = "2020-08-01 00:00:00"
)

func init() {
	RenameCmd.Flags().StringVarP(&input, "input", "i", "./", "Input directory")
	RenameCmd.Flags().StringVarP(&output, "output", "o", "./out", "Output directory")
	RenameCmd.Flags().IntVarP(&startIdx, "start", "s", 1, "Start index")
}

func renameRun(cmd *cobra.Command, args []string) {
	input = "D:\\FFOutput"
	startIdx = 20000

	startTimestamp, err := time.Parse("2006-01-02 15:04:05", startTime)
	if err != nil {
		panic(err)
	}

	_, err = os.Stat(input)
	if err != nil {
		panic(err)
	}
	_, err = os.Stat(output)
	if err != nil {
		if err = os.Mkdir(output, os.ModePerm); err != nil {
			panic(err)
		}
	}

	renameMap := make([]NameMap, 0, 100)
	filepath.Walk(input, func(filePath string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		oldName := info.Name()
		if i := strings.LastIndex(oldName, "."); i != -1 {
			oldName = oldName[:i]
		}
		nameMap := NameMap{
			OldName: oldName,
			ok:      false,
		}
		index := strings.Index(oldName, "-")
		if index == -1 {
			nameMap.NewName = "[ERROR]index=-1"
		} else {
			indexStr := oldName[:index]
			idx, err := strconv.Atoi(indexStr)
			if err != nil {
				nameMap.NewName = "[ERROR]not number:" + indexStr
			} else {
				nameMap.idx = startIdx + idx + 1
				nameMap.NewName = strconv.Itoa(nameMap.idx)
				nameMap.ok = true
			}
		}

		renameMap = append(renameMap, nameMap)
		if !nameMap.ok {
			return nil
		}
		baseDir := filepath.Dir(filePath)
		newName := path.Join(baseDir, nameMap.NewName+".mp3")
		_ = newName
		os.Rename(filePath, newName)
		modTime := startTimestamp.Add(time.Duration(nameMap.idx) * time.Second)
		_ = modTime
		os.Chtimes(newName, modTime, modTime)
		return nil
	})

	renameFile := path.Join(output, "rename.txt")
	file, err := os.Create(renameFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	// 写入文件
	sort.SliceStable(renameMap, func(i, j int) bool {
		return renameMap[i].idx < renameMap[j].idx
	})
	for _, nameMap := range renameMap {
		line := nameMap.NewName + "#" + nameMap.OldName + "\n"
		fmt.Print(line)
		writer.WriteString(line)
	}
	writer.Flush()
}

type NameMap struct {
	OldName string
	NewName string
	ok      bool
	idx     int
}
