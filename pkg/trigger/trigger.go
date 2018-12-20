package trigger

import (
	"fmt"

	ext_api "github.com/kube-ci/engine/apis/extensions/v1alpha1"
	cs "github.com/kube-ci/engine/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/client-go/rest"
)

func TriggerWorkflow(clientConfig *rest.Config, workflow, namespace string, request *unstructured.Unstructured) (*ext_api.Trigger, error) {
	kubeciClient, err := cs.NewForConfig(clientConfig)
	if err != nil {
		return nil, err
	}
	trigger := &ext_api.Trigger{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("%s-trigger", workflow),
			Namespace: namespace,
		},
		Workflows: []string{workflow},
		Request:   request,
	}
	return kubeciClient.ExtensionsV1alpha1().Triggers(namespace).Create(trigger)
}
