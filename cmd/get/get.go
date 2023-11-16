package get

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fatih/color"
	"github.com/flyteorg/flytectl/cmd/config/subcommand/clusterresourceattribute"
	"github.com/flyteorg/flytectl/cmd/config/subcommand/execution"
	"github.com/flyteorg/flytectl/cmd/config/subcommand/executionclusterlabel"
	"github.com/flyteorg/flytectl/cmd/config/subcommand/executionqueueattribute"
	"github.com/flyteorg/flytectl/cmd/config/subcommand/launchplan"
	pluginoverride "github.com/flyteorg/flytectl/cmd/config/subcommand/plugin_override"
	"github.com/flyteorg/flytectl/cmd/config/subcommand/project"
	"github.com/flyteorg/flytectl/cmd/config/subcommand/task"
	"github.com/flyteorg/flytectl/cmd/config/subcommand/taskresourceattribute"
	"github.com/flyteorg/flytectl/cmd/config/subcommand/workflow"
	"github.com/flyteorg/flytectl/cmd/config/subcommand/workflowexecutionconfig"
	cmdcore "github.com/flyteorg/flytectl/cmd/core"

	"github.com/spf13/cobra"
)

type model struct {
	choices  []string
	cursor   int
	selected string
}

var GetResourcesFuncs = map[string]cmdcore.CommandEntry{
	"project": {CmdFunc: getProjectsFunc, Aliases: []string{"projects"}, ProjectDomainNotRequired: true,
		Short: projectShort,
		Long:  projectLong, PFlagProvider: project.DefaultConfig},
	"task": {CmdFunc: getTaskFunc, Aliases: []string{"tasks"}, Short: taskShort,
		Long: taskLong, PFlagProvider: task.DefaultConfig},
	"workflow": {CmdFunc: getWorkflowFunc, Aliases: []string{"workflows"}, Short: workflowShort,
		Long: workflowLong, PFlagProvider: workflow.DefaultConfig},
	"launchplan": {CmdFunc: getLaunchPlanFunc, Aliases: []string{"launchplans"}, Short: launchPlanShort,
		Long: launchPlanLong, PFlagProvider: launchplan.DefaultConfig},
	"execution": {CmdFunc: getExecutionFunc, Aliases: []string{"executions"}, Short: executionShort,
		Long: executionLong, PFlagProvider: execution.DefaultConfig},
	"task-resource-attribute": {CmdFunc: getTaskResourceAttributes, Aliases: []string{"task-resource-attributes"},
		Short: taskResourceAttributesShort,
		Long:  taskResourceAttributesLong, PFlagProvider: taskresourceattribute.DefaultFetchConfig},
	"cluster-resource-attribute": {CmdFunc: getClusterResourceAttributes, Aliases: []string{"cluster-resource-attributes"},
		Short: clusterResourceAttributesShort,
		Long:  clusterResourceAttributesLong, PFlagProvider: clusterresourceattribute.DefaultFetchConfig},
	"execution-queue-attribute": {CmdFunc: getExecutionQueueAttributes, Aliases: []string{"execution-queue-attributes"},
		Short: executionQueueAttributesShort,
		Long:  executionQueueAttributesLong, PFlagProvider: executionqueueattribute.DefaultFetchConfig},
	"execution-cluster-label": {CmdFunc: getExecutionClusterLabel, Aliases: []string{"execution-cluster-labels"},
		Short: executionClusterLabelShort,
		Long:  executionClusterLabelLong, PFlagProvider: executionclusterlabel.DefaultFetchConfig},
	"plugin-override": {CmdFunc: getPluginOverridesFunc, Aliases: []string{"plugin-overrides"},
		Short: pluginOverrideShort,
		Long:  pluginOverrideLong, PFlagProvider: pluginoverride.DefaultFetchConfig},
	"workflow-execution-config": {CmdFunc: getWorkflowExecutionConfigFunc, Aliases: []string{"workflow-execution-config"},
		Short: workflowExecutionConfigShort,
		Long:  workflowExecutionConfigLong, PFlagProvider: workflowexecutionconfig.DefaultFetchConfig, ProjectDomainNotRequired: true},
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter":
			m.selected = m.choices[m.cursor]
			fmt.Println("You selected:", m.selected)
			// if cmdFunc, exists := GetResourcesFuncs[m.selected]; exists {
			// 	cmdFunc[m.selected]
			// }

			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	var s strings.Builder
	blue := color.New(color.FgBlue)

	for i, choice := range m.choices {
		if i == m.cursor {
			blue.Fprintf(&s, "> %s\n", choice)
		} else {
			fmt.Fprintf(&s, "  %s\n", choice)
		}
	}
	return s.String()
}

// Long descriptions are whitespace sensitive when generating docs using sphinx.
const (
	getCmdShort = `Fetches various Flyte resources such as tasks, workflows, launch plans, executions, and projects.`
	getCmdLong  = `
To fetch a project, use the following command:
::

 flytectl get project
`
)

// CreateGetCommand will return get command
func CreateGetCommand() *cobra.Command {
	getCmd := &cobra.Command{
		Use:   "get",
		Short: getCmdShort,
		Long:  getCmdLong,
		Run: func(cmd *cobra.Command, args []string) {
			p := tea.NewProgram(model{
				choices: []string{"project", "task", "workflow", "launchplan",
					"task-resource-attribute", "cluster-resource-attribute", "execution-queue-attribute",
					"execution-cluster-label", "plugin-override", "workflow-execution-config"},
			})
			if err := p.Start(); err != nil {
				fmt.Fprintf(os.Stderr, "Error occurred: %v\n", err)
				os.Exit(1)
			}
		},
	}

	fmt.Println("getCmd:", getCmd)

	cmdcore.AddCommands(getCmd, GetResourcesFuncs)

	return getCmd
}
