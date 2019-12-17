package component

import (
	"github.com/3scale/3scale-operator/pkg/common"
	integreatlyv1alpha1 "github.com/integr8ly/grafana-operator/pkg/apis/integreatly/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	apicastGrafanaDashboardJSON = `{
  "__inputs": [],
  "__requires": [
    {
      "type": "grafana",
      "id": "grafana",
      "name": "Grafana",
      "version": "4.2.0"
    },
    {
      "type": "panel",
      "id": "graph",
      "name": "Graph",
      "version": ""
    }
  ],
  "annotations": {
    "list": []
  },
  "editable": true,
  "gnetId": null,
  "graphTooltip": 0,
  "hideControls": false,
  "id": null,
  "links": [],
  "refresh": "30s",
  "rows": [
    {
      "collapse": false,
      "height": "250px",
      "panels": [
        {
          "aliasColors": {},
          "bars": false,
          "datasource": "$datasource",
          "fill": 1,
          "id": 1,
          "legend": {
            "avg": false,
            "current": false,
            "max": false,
            "min": false,
            "show": true,
            "total": false,
            "values": false
          },
          "lines": true,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "span": 6,
          "stack": false,
          "steppedLine": false,
          "targets": [
            {
              "expr": "sum(increase(upstream_status{kubernetes_namespace=\"apicast-[[environment]]\", status=~\"5.*\"}[$interval]))",
              "intervalFactor": 1,
              "legendFormat": "5xx",
              "metric": "",
              "refId": "B",
              "step": 20
            },
            {
              "expr": "sum(increase(upstream_status{kubernetes_namespace=\"apicast-[[environment]]\", status=~\"4.*\"}[$interval]))",
              "intervalFactor": 1,
              "legendFormat": "4xx",
              "metric": "",
              "refId": "C",
              "step": 20
            },
            {
              "expr": "sum(increase(upstream_status{kubernetes_namespace=\"apicast-[[environment]]\", status=~\"3.*\"}[$interval]))",
              "intervalFactor": 1,
              "legendFormat": "3xx",
              "metric": "",
              "refId": "D",
              "step": 20
            },
            {
              "expr": "sum(increase(upstream_status{kubernetes_namespace=\"apicast-[[environment]]\", status=~\"2.*\"}[$interval]))",
              "intervalFactor": 1,
              "legendFormat": "2xx",
              "metric": "",
              "refId": "E",
              "step": 20
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "Upstream Status (rx/min)",
          "tooltip": {
            "shared": true,
            "sort": 0,
            "value_type": "individual"
          },
          "type": "graph",
          "xaxis": {
            "mode": "time",
            "name": null,
            "show": true,
            "values": []
          },
          "yaxes": [
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            },
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            }
          ]
        },
        {
          "aliasColors": {},
          "bars": false,
          "datasource": "$datasource",
          "fill": 1,
          "id": 2,
          "legend": {
            "avg": false,
            "current": false,
            "max": false,
            "min": false,
            "show": true,
            "total": false,
            "values": false
          },
          "lines": true,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "span": 6,
          "stack": false,
          "steppedLine": false,
          "targets": [
            {
              "expr": "sum(increase(apicast_status{kubernetes_namespace=\"apicast-[[environment]]\", status=~\"5.*\"}[$interval]))",
              "intervalFactor": 1,
              "legendFormat": "5xx",
              "metric": "",
              "refId": "B",
              "step": 20
            },
            {
              "expr": "sum(increase(apicast_status{kubernetes_namespace=\"apicast-[[environment]]\", status=~\"4.*\"}[$interval]))",
              "intervalFactor": 1,
              "legendFormat": "4xx",
              "metric": "",
              "refId": "C",
              "step": 20
            },
            {
              "expr": "sum(increase(apicast_status{kubernetes_namespace=\"apicast-[[environment]]\", status=~\"3.*\"}[$interval]))",
              "intervalFactor": 1,
              "legendFormat": "3xx",
              "metric": "",
              "refId": "D",
              "step": 20
            },
            {
              "expr": "sum(increase(apicast_status{kubernetes_namespace=\"apicast-[[environment]]\", status=~\"2.*\"}[$interval]))",
              "intervalFactor": 1,
              "legendFormat": "2xx",
              "metric": "",
              "refId": "E",
              "step": 20
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "Apicast Status (rx/min)",
          "tooltip": {
            "shared": true,
            "sort": 0,
            "value_type": "individual"
          },
          "type": "graph",
          "xaxis": {
            "mode": "time",
            "name": null,
            "show": true,
            "values": []
          },
          "yaxes": [
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            },
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            }
          ]
        }
      ],
      "repeat": null,
      "repeatIteration": null,
      "repeatRowId": null,
      "showTitle": false,
      "title": "Dashboard Row",
      "titleSize": "h6"
    },
    {
      "collapse": false,
      "height": 250,
      "panels": [
        {
          "aliasColors": {},
          "bars": false,
          "datasource": "$datasource",
          "fill": 1,
          "id": 3,
          "legend": {
            "avg": false,
            "current": false,
            "max": false,
            "min": false,
            "show": true,
            "total": false,
            "values": false
          },
          "lines": true,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "span": 12,
          "stack": false,
          "steppedLine": false,
          "targets": [
            {
              "expr": "sum(increase(nginx_http_connections{kubernetes_namespace=~\"apicast-[[environment]]\"}[$interval])) by (state)",
              "hide": false,
              "interval": "",
              "intervalFactor": 1,
              "legendFormat": "{{state}}",
              "metric": "",
              "refId": "A",
              "step": 10
            },
            {
              "expr": "abs(sum(increase(nginx_http_connections{kubernetes_namespace=\"apicast-[[environment]]\",state=\"accepted\"}[$interval])) - sum(increase(nginx_http_connections{kubernetes_namespace=\"apicast-[[environment]]\",state=\"handled\"}[$interval])))",
              "hide": false,
              "intervalFactor": 1,
              "legendFormat": "dropped",
              "refId": "B",
              "step": 10
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "Connections/min",
          "tooltip": {
            "shared": true,
            "sort": 0,
            "value_type": "individual"
          },
          "type": "graph",
          "xaxis": {
            "mode": "time",
            "name": null,
            "show": true,
            "values": []
          },
          "yaxes": [
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            },
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            }
          ]
        }
      ],
      "repeat": null,
      "repeatIteration": null,
      "repeatRowId": null,
      "showTitle": false,
      "title": "Dashboard Row",
      "titleSize": "h6"
    },
    {
      "collapse": false,
      "height": 250,
      "panels": [
        {
          "aliasColors": {},
          "bars": false,
          "datasource": "$datasource",
          "fill": 1,
          "id": 4,
          "legend": {
            "avg": false,
            "current": false,
            "max": false,
            "min": false,
            "show": true,
            "total": false,
            "values": false
          },
          "lines": true,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "span": 12,
          "stack": false,
          "steppedLine": false,
          "targets": [
            {
              "expr": "sum(increase(nginx_error_log{kubernetes_namespace=\"apicast-[[environment]]\"}[$interval])) by (level)",
              "interval": "",
              "intervalFactor": 1,
              "legendFormat": "{{level}} logs",
              "refId": "A",
              "step": 10
            },
            {
              "expr": "sum(increase(nginx_metric_errors_total{kubernetes_namespace=\"apicast-[[environment]]\"}[$interval]))",
              "interval": "",
              "intervalFactor": 1,
              "legendFormat": "Nginx metrics errors",
              "refId": "B",
              "step": 10
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "Error logs / min",
          "tooltip": {
            "shared": true,
            "sort": 0,
            "value_type": "individual"
          },
          "type": "graph",
          "xaxis": {
            "mode": "time",
            "name": null,
            "show": true,
            "values": []
          },
          "yaxes": [
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            },
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            }
          ]
        }
      ],
      "repeat": null,
      "repeatIteration": null,
      "repeatRowId": null,
      "showTitle": false,
      "title": "Dashboard Row",
      "titleSize": "h6"
    },
    {
      "collapse": false,
      "height": 250,
      "panels": [
        {
          "aliasColors": {},
          "bars": false,
          "datasource": "$datasource",
          "fill": 1,
          "id": 5,
          "legend": {
            "avg": false,
            "current": false,
            "max": false,
            "min": false,
            "show": true,
            "total": false,
            "values": false
          },
          "lines": true,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "span": 6,
          "stack": false,
          "steppedLine": false,
          "targets": [
            {
              "expr": "100 - (sum(openresty_shdict_free_space{kubernetes_namespace=\"apicast-[[environment]]\"}) by (dict,instance)/sum(openresty_shdict_capacity{kubernetes_namespace=\"apicast-[[environment]]\"}) by (dict,instance)*100)",
              "intervalFactor": 1,
              "legendFormat": "{{dict}} - {{instance}}",
              "refId": "A",
              "step": 20
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "Shdict Used Percentage per instance",
          "tooltip": {
            "shared": true,
            "sort": 0,
            "value_type": "individual"
          },
          "type": "graph",
          "xaxis": {
            "mode": "time",
            "name": null,
            "show": true,
            "values": []
          },
          "yaxes": [
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            },
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            }
          ]
        },
        {
          "aliasColors": {},
          "bars": false,
          "datasource": "$datasource",
          "fill": 1,
          "id": 7,
          "legend": {
            "avg": false,
            "current": false,
            "max": false,
            "min": false,
            "show": true,
            "total": false,
            "values": false
          },
          "lines": true,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "span": 6,
          "stack": false,
          "steppedLine": false,
          "targets": [
            {
              "expr": "100 - (sum(openresty_shdict_free_space{kubernetes_namespace=\"apicast-[[environment]]\"}) by (dict)/sum(openresty_shdict_capacity{kubernetes_namespace=\"apicast-[[environment]]\"}) by (dict)*100)",
              "intervalFactor": 1,
              "legendFormat": "{{dict}} - {{instance}}",
              "refId": "A",
              "step": 20
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "Shdict Used Percentage Totals",
          "tooltip": {
            "shared": true,
            "sort": 0,
            "value_type": "individual"
          },
          "type": "graph",
          "xaxis": {
            "mode": "time",
            "name": null,
            "show": true,
            "values": []
          },
          "yaxes": [
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            },
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            }
          ]
        }
      ],
      "repeat": null,
      "repeatIteration": null,
      "repeatRowId": null,
      "showTitle": false,
      "title": "Dashboard Row",
      "titleSize": "h6"
    },
    {
      "collapse": false,
      "height": 250,
      "panels": [
        {
          "aliasColors": {},
          "bars": false,
          "datasource": "$datasource",
          "fill": 1,
          "id": 8,
          "legend": {
            "avg": false,
            "current": false,
            "max": false,
            "min": false,
            "show": true,
            "total": false,
            "values": false
          },
          "lines": true,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "span": 6,
          "stack": false,
          "steppedLine": false,
          "targets": [
            {
              "expr": "sum(increase(cloud_hosted_balancer{kubernetes_namespace=\"apicast-[[environment]]\"}[$interval])) by (status)",
              "intervalFactor": 1,
              "legendFormat": "{{ status }}",
              "refId": "A",
              "step": 20
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "Cloud hosted balancer / min",
          "tooltip": {
            "shared": true,
            "sort": 0,
            "value_type": "individual"
          },
          "type": "graph",
          "xaxis": {
            "mode": "time",
            "name": null,
            "show": true,
            "values": []
          },
          "yaxes": [
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            },
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            }
          ]
        },
        {
          "aliasColors": {},
          "bars": false,
          "datasource": "$datasource",
          "fill": 1,
          "id": 9,
          "legend": {
            "avg": false,
            "current": false,
            "max": false,
            "min": false,
            "show": true,
            "total": false,
            "values": false
          },
          "lines": true,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "span": 6,
          "stack": false,
          "steppedLine": false,
          "targets": [
            {
              "expr": "sum(increase(cloud_hosted_rate_limit{kubernetes_namespace=\"apicast-[[environment]]\"}[$interval])) by (state)",
              "intervalFactor": 1,
              "legendFormat": "{{state}}",
              "refId": "B",
              "step": 20
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "Cloud hosted limit / min",
          "tooltip": {
            "shared": true,
            "sort": 0,
            "value_type": "individual"
          },
          "type": "graph",
          "xaxis": {
            "mode": "time",
            "name": null,
            "show": true,
            "values": []
          },
          "yaxes": [
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            },
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            }
          ]
        }
      ],
      "repeat": null,
      "repeatIteration": null,
      "repeatRowId": null,
      "showTitle": false,
      "title": "Dashboard Row",
      "titleSize": "h6"
    }
  ],
  "schemaVersion": 14,
  "style": "dark",
  "tags": [
    "apicast",
    "nginx",
    "app",
    "middleware"
  ],
  "templating": {
    "list": [
      {
        "current": {
          "text": "prometheus-k8s-prod",
          "value": "prometheus-k8s-prod"
        },
        "hide": 0,
        "label": null,
        "name": "datasource",
        "options": [],
        "query": "prometheus",
        "refresh": 1,
        "regex": "",
        "type": "datasource"
      },
      {
        "allValue": null,
        "current": {
          "selected": true,
          "tags": [],
          "text": "production",
          "value": "production"
        },
        "hide": 0,
        "includeAll": false,
        "label": null,
        "multi": false,
        "name": "environment",
        "options": [
          {
            "selected": true,
            "text": "production",
            "value": "production"
          },
          {
            "selected": false,
            "text": "staging",
            "value": "staging"
          }
        ],
        "query": "production, staging",
        "type": "custom"
      },
      {
        "auto": true,
        "auto_count": 200,
        "auto_min": "1s",
        "current": {
          "text": "1m",
          "value": "1m"
        },
        "hide": 0,
        "label": "interval",
        "name": "interval",
        "options": [
          {
            "selected": false,
            "text": "auto",
            "value": "$__auto_interval"
          },
          {
            "selected": false,
            "text": "1s",
            "value": "1s"
          },
          {
            "selected": false,
            "text": "5s",
            "value": "5s"
          },
          {
            "selected": true,
            "text": "1m",
            "value": "1m"
          },
          {
            "selected": false,
            "text": "5m",
            "value": "5m"
          },
          {
            "selected": false,
            "text": "1h",
            "value": "1h"
          },
          {
            "selected": false,
            "text": "6h",
            "value": "6h"
          },
          {
            "selected": false,
            "text": "1d",
            "value": "1d"
          }
        ],
        "query": "1s,5s,1m,5m,1h,6h,1d",
        "refresh": 2,
        "type": "interval"
      }
    ]
  },
  "time": {
    "from": "now-6h",
    "to": "now"
  },
  "timepicker": {
    "refresh_intervals": [
      "5s",
      "10s",
      "30s",
      "1m",
      "5m",
      "15m",
      "30m",
      "1h",
      "2h",
      "1d"
    ],
    "time_options": [
      "5m",
      "15m",
      "1h",
      "6h",
      "12h",
      "24h",
      "2d",
      "7d",
      "30d"
    ]
  },
  "timezone": "browser",
  "title": "3scale Apicast",
  "version": 4
}`
)

func ApicastGrafanaDashboard() *integreatlyv1alpha1.GrafanaDashboard {
	return &integreatlyv1alpha1.GrafanaDashboard{
		ObjectMeta: metav1.ObjectMeta{
			Name: "apicast",
			Labels: map[string]string{
				"monitoring-key": common.MonitoringKey,
			},
		},
		Spec: integreatlyv1alpha1.GrafanaDashboardSpec{
			Json: apicastGrafanaDashboardJSON,
			Name: "apicast.json",
			Plugins: []integreatlyv1alpha1.GrafanaPlugin{
				{
					Name:    "grafana-piechart-panel",
					Version: "1.3.9",
				},
			},
		},
	}
}
