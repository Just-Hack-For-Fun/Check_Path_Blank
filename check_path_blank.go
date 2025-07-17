package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// 获取当前系统的 PATHEXT 并返回小写的扩展名（不带点）
func getPathexts() []string {
	pathext := os.Getenv("PATHEXT")
	if pathext == "" {
		pathext = ".COM;.EXE;.BAT;.CMD;.VBS;.VBE;.JS;.JSE;.WSF;.WSH;.MSC;.CPL"
	}
	items := strings.Split(strings.ToLower(pathext), ";")
	var result []string
	for _, ext := range items {
		ext = strings.TrimSpace(ext)
		ext = strings.TrimPrefix(ext, ".")
		if ext != "" {
			result = append(result, ext)
		}
	}
	return result
}

// 获取所有本地盘符（大写且存在的）
func getAllDrives() []string {
	var drives []string
	for c := 'C'; c <= 'Z'; c++ {
		drive := fmt.Sprintf("%c:\\", c)
		if _, err := os.Stat(drive); err == nil {
			drives = append(drives, drive)
		}
	}
	return drives
}

// 针对每个空格前缀检查是否有风险文件
func checkPrefixRisk(fullPath string, baseRoot string, pathexts []string) {
	if !strings.Contains(fullPath, " ") {
		return
	}
	relativePath := fullPath
	// 计算相对路径
	if strings.HasPrefix(strings.ToLower(fullPath), strings.ToLower(baseRoot)) {
		relativePath = fullPath[len(baseRoot):]
	}

	spaceIndexes := []int{}
	for idx, c := range relativePath {
		if c == ' ' {
			spaceIndexes = append(spaceIndexes, idx)
		}
	}
	for _, idx := range spaceIndexes {
		prefix := relativePath[:idx]
		base := filepath.Join(baseRoot, prefix)

		if fi, err := os.Stat(base); err == nil && !fi.IsDir() {
			fmt.Printf("[可疑文件] %s （截断自 %s）\n", base, fullPath)
		}
		for _, ext := range pathexts {
			exePath := base + "." + ext
			if fi, err := os.Stat(exePath); err == nil && !fi.IsDir() {
				fmt.Printf("[可疑可执行文件] %s （截断自 %s）\n", exePath, fullPath)
			}
		}
	}
}

func main() {
	startTime := time.Now()
	args := os.Args[1:]
	var rootDirs []string
	if len(args) == 0 {
		rootDirs = getAllDrives()
		if len(rootDirs) == 0 {
			fmt.Println("未找到任何可用盘符。")
			return
		}
	} else {
		rootDirs = args
	}

	pathexts := getPathexts()

	fmt.Printf("检测带空格路径的截断前缀可执行文件风险\n")
	fmt.Printf("开始扫描，请耐心等待... 开始时间: %v\n", startTime.Format("2006-01-02 15:04:05"))

	for _, root := range rootDirs {
		fmt.Printf("\n扫描路径: %s\n", root)
		filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return nil
			}
			if !d.IsDir() {
				checkPrefixRisk(path, root, pathexts)
			}
			return nil
		})
	}

	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Printf("\n扫描开始时间: %v\n", startTime.Format("2006-01-02 15:04:05"))
	fmt.Printf("扫描结束时间: %v\n", endTime.Format("2006-01-02 15:04:05"))
	fmt.Printf("总耗时: %v\n", duration.Round(time.Second))
}
