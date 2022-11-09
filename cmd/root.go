package cmd

import (
	"FilesMap/util"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "FilesMap",
	Short: "重命名文件夹文件并生成前后对照表",
	Long:  `重命名文件夹文件并生成前后对照表,运行传入文件夹路径即可`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			return
		}
		path := args[0]
		fmt.Println("输入了路径", path)
		fmt.Println("路径下所有文件", util.AllSubPath(path))
		filesMap := generateFilesMap(path)
		saveFilesMap(path, filesMap)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func generateFilesMap(path string) map[string]string {
	files := util.AllSubPath(path)
	if files == nil {
		return nil
	}

	filesMap := make(map[string]string)

	for i, filePath := range files {
		oldName := util.FileFullName(filePath)
		if strings.Contains(oldName, "DS_Store") {
			continue
		}
		if !util.IsDir(filePath) {
			_, suffixName := util.FileName(filePath)
			newName := fmt.Sprintf("%d%s", i, suffixName)
			err01 := os.Rename(filePath, path+"/"+newName)
			if err01 != nil {
				fmt.Println("重命名出错:", oldName)
				continue
			}
			filesMap[newName] = oldName
		}
	}
	return filesMap
}

func saveFilesMap(path string, filesMap map[string]string) {
	str := ""
	for key, value := range filesMap {
		str = str + key + "--" + value + "\n"
	}
	textPath := path + "/filesMap.text"

	err := ioutil.WriteFile(textPath, []byte(str), os.ModePerm)
	if err != nil {
		fmt.Println("写入错误", textPath)
	}

	// js, err := json.Marshal(filesMap)
	// if err != nil {
	// 	fmt.Println("转json出错", filesMap)
	// 	return
	// }

	// jsonPath := path + "/filesMap.json"

	// err02 := ioutil.WriteFile(jsonPath, js, os.ModePerm)
	// if err02 != nil {
	// 	fmt.Println("写入json错误", jsonPath)
	// }

}
