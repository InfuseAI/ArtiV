/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "art",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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
	cobra.EnableCommandSorting = false
	rootCmd.SetUsageTemplate(usageTemplate)

	addCommandWithGroup("basic",
		getCmd,
		putCmd,
	)

	addCommandWithGroup("workspace",
		initCommand,
		pullCmd,
		pushCmd,
		tagCommand,
		listCommand,
		logCommand,
		diffCommand,
		configCommand,
	)
}

func addCommandWithGroup(group string, cmds ...*cobra.Command) {
	for _, cmd := range cmds {
		cmd.Annotations = map[string]string{
			"group": group,
		}
	}

	rootCmd.AddCommand(cmds...)
}

var usageTemplate = `{{- /* usage template */ -}}
{{define "command" -}}
{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}
{{- end -}}
{{- end -}}
{{- /*
	Body
*/
-}}
Usage:{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command]{{end}}{{if gt (len .Aliases) 0}}

Aliases:
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

Examples:
{{.Example}}{{end}}
{{if .HasAvailableSubCommands}}
{{- if not .HasParent}}
Basic Commands:{{range .Commands}}{{if (eq .Annotations.group "basic")}}{{template "command" .}}{{end}}{{end}}

Workspace Commands:{{range .Commands}}{{if (eq .Annotations.group "workspace")}}{{template "command" .}}{{end}}{{end}}

Other Commands:{{range .Commands}}{{if not .Annotations.group}}{{template "command" .}}{{end}}{{end}}
{{- else}}
Available Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
    {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}
{{- end -}}
{{end}}
{{if .HasAvailableLocalFlags}}  
Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

Global Flags:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`
