package cmd

// 导入需要的包
import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// 定义根命令 ebhelp
var rootCmd = &cobra.Command{
	Use: "ebhelp",  // 命令的使用方式
	Short: "ebhelp is a study schedule generator assistant.",  // 命令的简短描述
	Long: `ebhelp can help you make your study schedule with Ebbinghaus Learning Curve, 
					  built by CodeSingerGnC in Go.`,  // 命令的长描述
	Run: func(cmd *cobra.Command, args []string) {  // 命令的运行函数
		print()
	},
}

// 定义一个子命令 version，用于打印版本号
var versionCmd = &cobra.Command{
	Use: "version",  // 子命令的使用方式
	Short: "Print the version number of ebhelp.",  // 子命令的简短描述
	Run: func(cmd *cobra.Command, args []string) {  // 子命令的运行函数
		fmt.Println("Ebhelp schedule generator assistant v0.1 -- alpha.")
	},
}

// 定义一个子命令 new，用于创建新的学习计划
var newCmd = &cobra.Command{
	Use: "new",
	Short: "Create a new study task.",
	Long: `Create a new study task. First, you need to create a new, 
		properly formatted JSON file in the .ebhelp directory 
		which is under the User account (./ebhelp), 
		if you did'nt find it, then you create one. 
		After that, you use this command to add it to the task list.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		new(args)
	},
}

// 定义一个子命令 schedule，用于查看今日计划
var scheduleCmd = &cobra.Command{
	Use: "schedule",
	Short: "View today's schedule",
	Run: func(cmd *cobra.Command, args []string) {
		m := FileDataMap{}
		m.schedule()
	},
}

var showCmd = &cobra.Command{
	Use: "show",
	Short: "View today's schedule",
	Run: func(cmd *cobra.Command, args []string) {
		m :=FileDataMap{}
		m.show()
	},
}

// 在init函数中，将子命令添加到根命令
func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(newCmd)
	rootCmd.AddCommand(scheduleCmd)
	rootCmd.AddCommand(showCmd)
}

// Execute函数执行根命令，如果有错误，打印错误并退出程序
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}