package cmds

import (
	"fmt"

	"github.com/appscode/kutil/meta"
	"github.com/kube-ci/cli/pkg/trigger"
	ext_api "github.com/kube-ci/engine/apis/extensions/v1alpha1"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func NewCmdTrigger(clientGetter genericclioptions.RESTClientGetter) *cobra.Command {
	var (
		workflow  string
		namespace string
		request   string
		output    string
	)

	cmd := &cobra.Command{
		Use:               "trigger",
		Short:             "Trigger workflow",
		Long:              "Trigger workflow",
		DisableAutoGenTag: true,
		Args:              cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			if len(args) == 0 {
				return fmt.Errorf("workflow name not specified")
			}
			workflow = args[0]

			namespace, _, err = clientGetter.ToRawKubeConfigLoader().Namespace()
			if err != nil {
				return err
			}

			cfg, err := clientGetter.ToRESTConfig()
			if err != nil {
				return err
			}

			var requestUnstructured *unstructured.Unstructured

			if len(request) != 0 {
				requestJSON, err := yaml.ToJSON([]byte(request))
				if err != nil {
					return fmt.Errorf("failed to convert content to json, reason: %s", err)
				}
				obj, _, err := unstructured.UnstructuredJSONScheme.Decode(requestJSON, nil, nil)
				if err != nil {
					return fmt.Errorf("failed to decode content to unstructured, reason: %s", err)
				}
				var ok bool
				requestUnstructured, ok = obj.(*unstructured.Unstructured)
				if !ok {
					return fmt.Errorf("failed to convert object to unstructured")
				}
			}

			trigger, err := trigger.TriggerWorkflow(cfg, workflow, namespace, requestUnstructured)
			if err != nil {
				return fmt.Errorf("failed to trigger workflow, reason: %v", err)
			}
			// if no error, print the trigger object
			return printTrigger(trigger, output)
		},
	}
	cmd.Flags().StringVar(&request, "by", "", "Contents of request object.")
	cmd.Flags().StringVarP(&output, "output", "o", "", "Output format. One of: json|yaml")
	return cmd
}

func printTrigger(trigger *ext_api.Trigger, output string) error {
	switch output {
	case "json":
		triggerJSON, err := meta.MarshalToJson(trigger, ext_api.SchemeGroupVersion)
		if err != nil {
			return fmt.Errorf("failed to print output, reason: %v", err)
		}
		fmt.Printf(string(triggerJSON))
	case "yaml":
		triggerYAML, err := meta.MarshalToYAML(trigger, ext_api.SchemeGroupVersion)
		if err != nil {
			return fmt.Errorf("failed to print output, reason: %v", err)
		}
		fmt.Printf(string(triggerYAML))
	default:
		fmt.Printf("trigger.extensions.kube.ci/%s created\n", trigger.Name)
	}
	return nil
}

// kubectl ci trigger sample-workflow -n demo
// kubectl ci trigger sample-workflow -n demo -o yaml
// kubectl ci trigger sample-workflow -n demo --by "$(cat docs/examples/engine/hello-world/configmap.yaml)"
// kubectl ci trigger sample-workflow -n demo --by "$(kubectl get configmap sample-config -n demo -o yaml)"
