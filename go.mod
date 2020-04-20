module github.com/quay/container-security-operator

go 1.13

require (
	github.com/coreos/prometheus-operator v0.38.1
	github.com/go-kit/kit v0.9.0
	github.com/go-openapi/spec v0.19.3
	github.com/imdario/mergo v0.3.8 // indirect
	github.com/pkg/errors v0.8.1
	github.com/prometheus/client_golang v1.2.1
	github.com/sirupsen/logrus v1.4.2
	github.com/stretchr/testify v1.4.0
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45 // indirect
	gopkg.in/yaml.v2 v2.2.8
	k8s.io/api v0.17.3
	k8s.io/apimachinery v0.17.3
	k8s.io/client-go v12.0.0+incompatible
	k8s.io/kube-openapi v0.0.0-20191107075043-30be4d16710a
)
