/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xinlianit/kit-scaffold/cli/common"
	"github.com/xinlianit/kit-scaffold/cli/util"
	"strings"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "new",
	Short: "创建应用",
	Long:  `创建一个项目应用`,
	Run: func(cmd *cobra.Command, args []string) {
		// 参数解析
		argsMap := common.ArgsToMap(args, []string{"name"})

		if argsMap["name"] == "" {
			cobra.CheckErr("项目名称为空")
		}

		// 项目路径
		projectPath, _ := cmd.Flags().GetString("path")

		// 项目目录
		projectDir := strings.TrimRight(projectPath, "/") + "/" + argsMap["name"]

		// 检测目录是否存在
		if util.DirUtil().IsDir(projectDir) && !util.DirUtil().IsEmptyDir(projectDir) {
			cobra.CheckErr(projectDir + " 已存在且目录不为空")
		}

		// 创建项目目录
		if err := util.DirUtil().CreateDir(projectDir, true); err != nil {
			cobra.CheckErr(fmt.Sprintf("%s 目录创建失败: %s", projectDir, err))
		}

		// 部署项目结构
		projectStruct := viper.GetStringSlice("project.struct")

		for _, childDir := range projectStruct {
			// 创建子目录
			childDir = projectDir + "/" + strings.TrimLeft(childDir, "/")
			if err := util.DirUtil().CreateDir(childDir, true); err != nil {
				cobra.CheckErr(fmt.Sprintf("%s 目录创建失败: %s", childDir, err))
			}
		}

		// 生成框架基础代码

		fmt.Println(projectDir, " 项目创建成功！")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// 项目路径
	generateCmd.Flags().StringP("path", "p", "./", "指定生成项目路径")
}
