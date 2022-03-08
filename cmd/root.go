package cmd

import (
	"FilesMap/util"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "FilesMap",
	Short: "重命名文件夹文件并生成前后对照表",
	Long: `重命名文件夹文件并生成前后对照表,运行传入文件夹路径即可`,
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
		fmt.Println(filesMap)
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

func generateFilesMap(path string) map[string]interface{} {
	files := util.AllSubPath(path)
	if files == nil {
		return nil
	}

	filesMap := make(map[string]interface{})

	for i, filePath := range files {
		oldName := util.FileFullName(filePath)
		if oldName == ".DS_Store" {
			continue
		}
		if util.IsDir(filePath) {
			subFilesMap := generateFilesMap(filePath)
			filesMap[oldName] = subFilesMap
		}else {
			_ ,suffixName := util.FileName(filePath)
			newName := fmt.Sprintf("%d%s",i, suffixName)
			err01 := os.Rename(filePath, path+ "/" + newName)
			if err01 != nil {
				fmt.Println("重命名出错:", oldName)
				continue
			}
			filesMap[newName] = oldName
		}
	}
	return filesMap
}


