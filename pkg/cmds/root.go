package cmds

import (
	"flag"

	v "github.com/appscode/go/version"
	enginescheme "github.com/kube-ci/engine/client/clientset/versioned/scheme"
	gitscheme "github.com/kube-ci/git-apiserver/client/clientset/versioned/scheme"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	clientsetscheme "k8s.io/client-go/kubernetes/scheme"
	cliflag "k8s.io/component-base/cli/flag"
	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"kmodules.xyz/client-go/logs"
	"kmodules.xyz/client-go/tools/cli"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:               "ci [command]",
		Short:             `KubeCI by AppsCode - Kubernetes Native Workflow Engine`,
		DisableAutoGenTag: true,
		PersistentPreRun: func(c *cobra.Command, args []string) {
			cli.SendAnalytics(c, v.Version.Version)
			enginescheme.AddToScheme(clientsetscheme.Scheme)
			gitscheme.AddToScheme(clientsetscheme.Scheme)
		},
	}

	flags := rootCmd.PersistentFlags()
	// Normalize all flags that are coming from other packages or pre-configurations
	// a.k.a. change all "_" to "-". e.g. glog package
	flags.SetNormalizeFunc(cliflag.WordSepNormalizeFunc)

	kubeConfigFlags := genericclioptions.NewConfigFlags(true)
	kubeConfigFlags.AddFlags(flags)
	matchVersionKubeConfigFlags := cmdutil.NewMatchVersionFlags(kubeConfigFlags)
	matchVersionKubeConfigFlags.AddFlags(flags)

	flags.AddGoFlagSet(flag.CommandLine)
	logs.ParseFlags()
	flags.BoolVar(&cli.EnableAnalytics, "enable-analytics", cli.EnableAnalytics, "Send analytical events to Google Analytics")
	flag.Set("stderrthreshold", "ERROR")

	rootCmd.AddCommand(NewCmdWorkplanLogs(matchVersionKubeConfigFlags))
	rootCmd.AddCommand(NewCmdTrigger(matchVersionKubeConfigFlags))
	rootCmd.AddCommand(v.NewCmdVersion())
	return rootCmd
}
