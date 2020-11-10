module gitlab.com/cloudmanaged/operator/alert-receiver

go 1.13

require (
	github.com/evanphx/json-patch v4.9.0+incompatible // indirect
	github.com/golang/groupcache v0.0.0-20191227052852-215e87163ea7 // indirect
	github.com/googleapis/gnostic v0.4.0 // indirect
	github.com/imdario/mergo v0.3.11 // indirect
	gitlab.com/cloudmanaged/operator v0.0.1
	golang.org/x/net v0.0.0-20201110031124-69a78807bb2b
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
	google.golang.org/protobuf v1.24.0 // indirect
	k8s.io/api v0.19.0-alpha.1 // indirect
	k8s.io/apimachinery v0.19.0-alpha.1
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/gengo v0.0.0-20200413195148-3a45101e95ac // indirect
	k8s.io/klog v1.0.0 // indirect
	k8s.io/klog/v2 v2.2.0 // indirect
	k8s.io/utils v0.0.0-20201104234853-8146046b121e // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.0.1 // indirect
)

replace k8s.io/client-go => k8s.io/client-go v0.18.8
