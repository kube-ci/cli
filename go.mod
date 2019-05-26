module github.com/kube-ci/cli

go 1.12

require (
	github.com/Azure/go-ansiterm v0.0.0-20170929234023-d6e3b3328b78 // indirect
	github.com/Azure/go-autorest v12.0.0+incompatible // indirect
	github.com/MakeNowJust/heredoc v0.0.0-20171113091838-e9091a26100e // indirect
	github.com/Netflix/go-expect v0.0.0-20180928190340-9d1f4485533b // indirect
	github.com/appscode/go v0.0.0-20190523031839-1468ee3a76e8
	github.com/docker/docker v0.7.3-0.20190327010347-be7ac8be2ae0 // indirect
	github.com/exponent-io/jsonpath v0.0.0-20151013193312-d6023ce2651d // indirect
	github.com/hinshun/vt10x v0.0.0-20180809195222-d55458df857c // indirect
	github.com/kube-ci/engine v0.0.0-20190526105551-8a8da8d1fd1e
	github.com/kube-ci/git-apiserver v0.0.0-20190517164833-b5321becdb4a
	github.com/mgutz/ansi v0.0.0-20170206155736-9520e82c474b // indirect
	github.com/mitchellh/go-wordwrap v1.0.0 // indirect
	github.com/spf13/cobra v0.0.4
	gopkg.in/AlecAivazis/survey.v1 v1.6.1
	k8s.io/apimachinery v0.0.0-20190515023456-b74e4c97951f
	k8s.io/cli-runtime v0.0.0-20190515024640-178667528169
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	k8s.io/component-base v0.0.0-20190515024022-2354f2393ad4
	k8s.io/kubernetes v1.14.2
	kmodules.xyz/client-go v0.0.0-20190524133821-9c8a87771aea
)

replace (
	github.com/graymeta/stow => github.com/appscode/stow v0.0.0-20190506085026-ca5baa008ea3
	gopkg.in/robfig/cron.v2 => github.com/appscode/cron v0.0.0-20170717094345-ca60c6d796d4
	k8s.io/api => k8s.io/api v0.0.0-20190313235455-40a48860b5ab
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20190315093550-53c4693659ed
	k8s.io/apimachinery => github.com/kmodules/apimachinery v0.0.0-20190508045248-a52a97a7a2bf
	k8s.io/apiserver => github.com/kmodules/apiserver v0.0.0-20190508082252-8397d761d4b5
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.0.0-20190314001948-2899ed30580f
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.0.0-20190314002645-c892ea32361a
	k8s.io/component-base => k8s.io/component-base v0.0.0-20190314000054-4a91899592f4
	k8s.io/klog => k8s.io/klog v0.3.0
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.0.0-20190314000639-da8327669ac5
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20190228160746-b3a7cee44a30
	k8s.io/metrics => k8s.io/metrics v0.0.0-20190314001731-1bd6a4002213
	k8s.io/utils => k8s.io/utils v0.0.0-20190221042446-c2654d5206da
)
