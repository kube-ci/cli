package cmds

import (
	"github.com/appscode/go/log"
	workplan_logs "github.com/kube-ci/cli/pkg/workplan-logs"
	"github.com/kube-ci/engine/pkg/logs"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func NewCmdWorkplanLogs(clientGetter genericclioptions.RESTClientGetter) *cobra.Command {
	var query logs.Query

	cmd := &cobra.Command{
		Use:               "workplan-logs",
		Short:             "Get workplan logs",
		Long:              "Get workplan logs",
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := clientGetter.ToRESTConfig()
			if err != nil {
				return err
			}
			if err = workplan_logs.GetLogs(cfg, query); err != nil {
				log.Errorf("error collecting log, reason: %v", err)
			}
			return nil
		},
	}
	cmd.Flags().StringVar(&query.Namespace, "namespace", "", "Namespace of the workflow.") // TODO: set default ns?
	cmd.Flags().StringVar(&query.Workflow, "workflow", "", "Name of the workflow.")
	cmd.Flags().StringVar(&query.Workplan, "workplan", "", "Name of the workplan.")
	cmd.Flags().StringVar(&query.Step, "step", "", "Name of the step.")
	return cmd
}

// engine workplan-logs --namespace default --workplan wf-pr-test-9dtw8 --step step-test
// using kubectl-plugin
// cp /home/dipta/go/bin/engine /home/dipta/go/bin/kubectl-kubeci
// kubectl kubeci workplan-logs
