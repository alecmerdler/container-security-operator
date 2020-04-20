package k8sutils

import (
	monitoringv1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func createPodMonitor(name, namespace string) *monitoringv1.PodMonitor {
	return &monitoringv1.PodMonitor{
		TypeMeta: metav1.TypeMeta{
			Kind:       monitoringv1.DefaultCrdKinds.PodMonitor.Kind,
			APIVersion: monitoringv1.SchemeGroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			Labels:    map[string]string{},
		},
		Spec: monitoringv1.PodMonitorSpec{
			NamespaceSelector: monitoringv1.NamespaceSelector{MatchNames: []string{"openshift-operators"}},
			Selector:          metav1.LabelSelector{MatchLabels: map[string]string{"name": "container-security-operator-alm-owned"}},
			PodMetricsEndpoints: []monitoringv1.PodMetricsEndpoint{
				{Port: "metrics"},
			},
		},
	}
}

func createPrometheusRule() *monitoringv1.PrometheusRule {

}
