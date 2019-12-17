package component

import (
	"fmt"

	"github.com/3scale/3scale-operator/pkg/common"
	monitoringv1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func ApicastPrometheusRules(ns string) *monitoringv1.PrometheusRule {
	return &monitoringv1.PrometheusRule{
		ObjectMeta: v12.ObjectMeta{
			Name: "apicast",
			Labels: map[string]string{
				"monitoring-key": common.MonitoringKey,
				"prometheus":     "application-monitoring",
				"role":           "alert-rules",
			},
		},
		Spec: monitoringv1.PrometheusRuleSpec{
			Groups: []monitoringv1.RuleGroup{
				{
					Name: "apicast",
					Rules: []monitoringv1.Rule{
						{
							Alert: "ApiCastProdRunningPods",
							Annotations: map[string]string{
								"summary":     "{{$labels.container_name}} replica controller on {{$labels.namespace}}: Less than 2 running pods",
								"description": "{{$labels.container_name}} replica controller on {{$labels.namespace}} project: Less than 2 running pods",
							},
							Expr: intstr.FromString(fmt.Sprintf(`label_replace(label_replace(label_replace(sum(clamp_max(container_memory_usage_bytes{container_name="apicast",namespace="%s"},1)), "cluster", "prod", "", ""  ) ,"container_name","apicast","","" ),"namespace", "%s", "", "") < 2`, ns, ns)),
							For:  "2m",
							Labels: map[string]string{
								"severity": "critical",
							},
						}, {
							Alert: "ApiCastErrors",
							Annotations: map[string]string{
								"summary":     "{{$labels.container_name}} replica controleer {{$labels.namespace}}: Has more than 5 errors in the last 5 minutes",
								"description": "{{$labels.container_name}} replica controller on {{$labels.namespace}} project: Has more than 5 errors in the last 5 minutes",
							},
							Expr: intstr.FromString(fmt.Sprintf(`sum(increase(nginx_error_log{kubernetes_namespace="%s",level=~"(error|crit|alert|emerg)"}[5m])) by (kubernetes_name,cluster,kubernetes_namespace) > 100`, ns)),
							For:  "2m",
							Labels: map[string]string{
								"severity": "critical",
							},
						}, {
							Alert: "ApiCast5xx",
							Annotations: map[string]string{
								"summary":     "{{$labels.container_name}} replica controller on {{$labels.namespace}}: Has more than 10 Http 5XX in the last 5 minutes",
								"description": "{{$labels.container_name}} replica controller on {{$labels.namespace}} project: Has more than 10 Http 5XX in the last 5 minutes",
							},
							Expr: intstr.FromString(fmt.Sprintf(`sum(increase(apicast_status{kubernetes_namespace="%s",status=~"5\\d{2}"}[5m])) by (kubernetes_name,cluster,kubernetes_namespace) > 10`, ns)),
							For:  "2m",
							Labels: map[string]string{
								"severity": "warning",
							},
						},
					},
				},
			},
		},
	}
}
