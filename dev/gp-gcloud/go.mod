module github.com/gitpod-io/gitpod/gp-gcloud

go 1.22

require (
	github.com/gitpod-io/gitpod/common-go v0.0.0-00010101000000-000000000000
	github.com/spf13/cobra v1.5.0
	golang.org/x/net v0.23.0
	golang.org/x/oauth2 v0.7.0
	google.golang.org/api v0.114.0
)

require (
	cloud.google.com/go/compute v1.19.1 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/gitpod-io/gitpod/components/scrubber v0.0.0-00010101000000-000000000000 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.2.3 // indirect
	github.com/googleapis/gax-go/v2 v2.7.1 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/hashicorp/golang-lru v1.0.2 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/prometheus/client_golang v1.16.0 // indirect
	github.com/prometheus/client_model v0.3.0 // indirect
	github.com/prometheus/common v0.42.0 // indirect
	github.com/prometheus/procfs v0.10.1 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	go.opencensus.io v0.24.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.20.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1 // indirect
	google.golang.org/grpc v1.56.3 // indirect
	google.golang.org/protobuf v1.35.2 // indirect
)

replace github.com/gitpod-io/gitpod/common-go => ../../components/common-go // leeway

replace github.com/gitpod-io/gitpod/components/scrubber => ../../components/scrubber // leeway

replace k8s.io/api => k8s.io/api v0.29.3 // leeway indirect from components/common-go:lib

replace k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.29.3 // leeway indirect from components/common-go:lib

replace k8s.io/apimachinery => k8s.io/apimachinery v0.29.3 // leeway indirect from components/common-go:lib

replace k8s.io/apiserver => k8s.io/apiserver v0.29.3 // leeway indirect from components/common-go:lib

replace k8s.io/cli-runtime => k8s.io/cli-runtime v0.29.3 // leeway indirect from components/common-go:lib

replace k8s.io/client-go => k8s.io/client-go v0.29.3 // leeway indirect from components/common-go:lib

replace k8s.io/cloud-provider => k8s.io/cloud-provider v0.29.3 // leeway indirect from components/common-go:lib

replace k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.29.3 // leeway indirect from components/common-go:lib

replace k8s.io/code-generator => k8s.io/code-generator v0.29.3 // leeway indirect from components/common-go:lib

replace k8s.io/component-base => k8s.io/component-base v0.29.3 // leeway indirect from components/common-go:lib

replace k8s.io/cri-api => k8s.io/cri-api v0.29.3 // leeway indirect from components/common-go:lib

replace k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.29.3 // leeway indirect from components/common-go:lib

replace k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.29.3 // leeway indirect from components/common-go:lib

replace k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.29.3 // leeway indirect from components/common-go:lib

replace k8s.io/kube-proxy => k8s.io/kube-proxy v0.29.3 // leeway indirect from components/common-go:lib

replace k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.29.3 // leeway indirect from components/common-go:lib

replace k8s.io/kubelet => k8s.io/kubelet v0.29.3 // leeway indirect from components/common-go:lib

replace k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.29.3 // leeway indirect from components/common-go:lib

replace k8s.io/metrics => k8s.io/metrics v0.29.3 // leeway indirect from components/common-go:lib

replace k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.29.3 // leeway indirect from components/common-go:lib

replace k8s.io/component-helpers => k8s.io/component-helpers v0.29.3 // leeway indirect from components/common-go:lib

replace k8s.io/controller-manager => k8s.io/controller-manager v0.29.3 // leeway indirect from components/common-go:lib

replace k8s.io/kubectl => k8s.io/kubectl v0.29.3 // leeway indirect from components/common-go:lib

replace k8s.io/mount-utils => k8s.io/mount-utils v0.29.3 // leeway indirect from components/common-go:lib

replace k8s.io/pod-security-admission => k8s.io/pod-security-admission v0.29.3 // leeway indirect from components/common-go:lib
