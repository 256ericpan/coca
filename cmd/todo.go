package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/phodal/coca/cmd/cmd_util"
	"github.com/phodal/coca/pkg/context/todo"
	"github.com/spf13/cobra"
	"strings"
)

type RootCmdConfig struct {
	Path    string
	WithGit bool
}

var (
	todoCmdConfig RootCmdConfig
)

var todoCmd = &cobra.Command{
	Use:   "todo",
	Short: "scan all todo, and list with time",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		path := cmd.Flag("path").Value.String()
		app := todo.NewTodoApp()
		todos := app.AnalysisPath(path)

		simple, _ := json.MarshalIndent(todos, "", "\t")
		cmd_util.WriteToCocaFile("simple-todos.json", string(simple))

		fmt.Fprintf(output, "Todos Count %d\n", len(todos))

		if todoCmdConfig.WithGit {
			gitTodos := app.BuildWithGitHistory(todos)

			cModel, _ := json.MarshalIndent(todos, "", "\t")
			cmd_util.WriteToCocaFile("todos.json", string(cModel))

			table := tablewriter.NewWriter(output)
			table.SetHeader([]string{"Date", "Author", "Messages", "FileName", "Line"})
			for _, todo := range gitTodos {
				table.Append([]string{todo.Date, todo.Author, strings.Join(todo.Message, "\n"), todo.FileName, todo.Line})
			}

			table.Render()
		}
	},
}

func init() {
	todoCmd.SetOut(output)
	todoCmd.PersistentFlags().StringVarP(&todoCmdConfig.Path, "path", "p", ".", "path")
	todoCmd.PersistentFlags().BoolVarP(&todoCmdConfig.WithGit, "git", "g", false, "path")

	rootCmd.AddCommand(todoCmd)
}
