package workplan_logs

import (
	"bufio"
	"fmt"

	api "github.com/kube-ci/engine/apis/engine/v1alpha1"
	"github.com/kube-ci/engine/pkg/logs"
	survey "gopkg.in/AlecAivazis/survey.v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/rest"
)

func GetLogs(clientConfig *rest.Config, query logs.Query) error {
	logController, err := logs.NewLogController(clientConfig)
	if err != nil {
		return fmt.Errorf("error initializing log-controller, reason: %s", err)
	}
	if err := prepare(logController, &query); err != nil {
		return err
	}
	reader, err := logController.LogReader(query)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // write log to stdout
	}
	return nil
}

// if Workplan name not provided, ask Workflow name, then ask to select a Workplan
// if Workplan name provided, we don't need Workflow
func prepare(c *logs.LogController, query *logs.Query) error {
	if query.Namespace == "" {
		if err := askNamespace(c, query); err != nil {
			return err
		}
	}
	if query.Workplan == "" {
		if query.Workflow == "" {
			if err := askWorkflow(c, query); err != nil {
				return err
			}
		}
		if err := askWorkplan(c, query); err != nil {
			return err
		}
	}
	if query.Step == "" {
		return askStep(c, query)
	}
	return nil
}

func askNamespace(c *logs.LogController, query *logs.Query) error {
	var namespaceNames []string
	namespaces, err := c.KubeClient.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		return err
	}
	for _, ns := range namespaces.Items {
		namespaceNames = append(namespaceNames, ns.Name)
	}
	if len(namespaceNames) == 0 {
		return fmt.Errorf("no namespace found")
	}

	qs := []*survey.Question{
		{
			Name: "Namespace",
			Prompt: &survey.Select{
				Message: "Choose a Namespace:",
				Options: namespaceNames,
			},
		},
	}
	return survey.Ask(qs, query)
}

func askWorkflow(c *logs.LogController, query *logs.Query) error {
	// list all workflows in the given Namespace
	var workflowNames []string
	workflows, err := c.KubeciClient.EngineV1alpha1().Workflows(query.Namespace).List(metav1.ListOptions{})
	if err != nil {
		return err
	}
	for _, wf := range workflows.Items {
		workflowNames = append(workflowNames, wf.Name)
	}
	if len(workflowNames) == 0 {
		return fmt.Errorf("no workflow found")
	}

	qs := []*survey.Question{
		{
			Name: "Workflow",
			Prompt: &survey.Select{
				Message: "Choose a Workflow:",
				Options: workflowNames,
			},
		},
	}
	return survey.Ask(qs, query)
}

func askWorkplan(c *logs.LogController, query *logs.Query) error {
	// list all workplans in the given Namespace for a specific Workflow
	var workplanNames []string
	workplans, err := c.KubeciClient.EngineV1alpha1().Workplans(query.Namespace).List(metav1.ListOptions{
		LabelSelector: labels.FormatLabels(map[string]string{"workflow": query.Workflow}),
	})
	if err != nil {
		return err
	}
	for _, wp := range workplans.Items {
		workplanNames = append(workplanNames, wp.Name)
	}
	if len(workplanNames) == 0 {
		return fmt.Errorf("no workplan found")
	}

	qs := []*survey.Question{
		{
			Name: "Workplan",
			Prompt: &survey.Select{
				Message: "Choose a Workplan:",
				Options: workplanNames,
			},
		},
	}
	return survey.Ask(qs, query)
}

func askStep(c *logs.LogController, query *logs.Query) error {
	// list all steps in the given Workplan along with their status
	workplanStatus, err := c.WorkplanStatus(*query)
	if err != nil {
		return err
	}

	// list all running and terminated steps, logs are available only for those steps
	var stepNames []string
	for _, stepEntries := range workplanStatus.StepTree {
		for _, stepEntry := range stepEntries {
			if stepEntry.Status == api.ContainerRunning || stepEntry.Status == api.ContainerTerminated {
				stepNames = append(stepNames, stepEntry.Name)
			}
		}
	}

	qs := []*survey.Question{
		{
			Name: "Step",
			Prompt: &survey.Select{
				Message: "Choose a running/terminated step:",
				Options: stepNames,
			},
		},
	}
	return survey.Ask(qs, query)
}
