module github.com/integreatly/crDetect

go 1.13

require (
	github.com/evanphx/json-patch v4.5.0+incompatible // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/google/btree v1.0.0 // indirect
	github.com/googleapis/gnostic v0.3.1 // indirect
	github.com/gregjones/httpcache v0.0.0-20190611155906-901d90724c79 // indirect
	github.com/json-iterator/go v1.1.8 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/peterbourgon/diskv v2.0.1+incompatible // indirect
	github.com/pkg/errors v0.8.1 // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stretchr/testify v1.4.0 // indirect
	golang.org/x/crypto v0.0.0-20190611184440-5c40567a22f8 // indirect
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45 // indirect
	golang.org/x/time v0.0.0-20181108054448-85acf8d2951c // indirect
	google.golang.org/appengine v1.5.0 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	k8s.io/api v0.0.0-20190918155943-95b840bb6a1f
	k8s.io/apimachinery v0.0.0-20191203211716-adc6f4cd9e7d
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/klog v1.0.0 // indirect
	k8s.io/kube-openapi v0.0.0-20191107075043-30be4d16710a // indirect
)

// Pinned to kubernetes-1.13.4
replace (
	github.com/coreos/prometheus-operator => github.com/coreos/prometheus-operator v0.29.0
	github.com/operator-framework/operator-sdk => github.com/operator-framework/operator-sdk v0.10.0
	github.com/prometheus/prometheus => github.com/prometheus/prometheus v0.0.0-20190424153033-d3245f150225
	k8s.io/api => k8s.io/api v0.0.0-20190222213804-5cb15d344471
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20190228180357-d002e88f6236
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190221213512-86fb29eff628
	k8s.io/client-go => k8s.io/client-go v0.0.0-20190228174230-b40b2a5939e4
	k8s.io/kube-state-metrics => k8s.io/kube-state-metrics v1.6.0
	sigs.k8s.io/controller-runtime => sigs.k8s.io/controller-runtime v0.1.12
	sigs.k8s.io/controller-tools => sigs.k8s.io/controller-tools v0.1.11-0.20190411181648-9d55346c2bde
)
