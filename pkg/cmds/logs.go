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
		Use:               "logs",
		Short:             "Get workplan logs",
		Long:              "Get workplan logs",
		DisableAutoGenTag: true,
		Args:              cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			if len(args) == 1 {
				query.Workplan = args[0]
			}
			query.Namespace, _, err = clientGetter.ToRawKubeConfigLoader().Namespace()
			if err != nil {
				return err
			}
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
	cmd.Flags().StringVar(&query.Workflow, "workflow", "", "Name of the workflow.")
	cmd.Flags().StringVar(&query.Step, "step", "", "Name of the step.")
	return cmd
}
