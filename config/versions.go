package config

import "github.com/darxkies/k8s-tew/utils"

type Versions struct {
	Etcd                       string `yaml:"etcd"`
	K8S                        string `yaml:"kubernetes"`
	Helm                       string `yaml:"helm"`
	Containerd                 string `yaml:"containerd"`
	Runc                       string `yaml:"runc"`
	CriCtl                     string `yaml:"crictl"`
	Gobetween                  string `yaml:"gobetween"`
	Ark                        string `yaml:"ark"`
	MinioServer                string `yaml:"minio-server"`
	MinioClient                string `yaml:"minio-client"`
	Pause                      string `yaml:"pause"`
	CoreDNS                    string `yaml:"core-dns"`
	Elasticsearch              string `yaml:"elasticsearch"`
	ElasticsearchCron          string `yaml:"elasticsearch-cron"`
	ElasticsearchOperator      string `yaml:"elasticsearch-operator"`
	Kibana                     string `yaml:"kibana"`
	Cerebro                    string `yaml:"cerebro"`
	FluentBit                  string `yaml:"fluent-bit"`
	CalicoTypha                string `yaml:"calico-typha"`
	CalicoNode                 string `yaml:"calico-node"`
	CalicoCNI                  string `yaml:"calico-cni"`
	Ceph                       string `yaml:"ceph"`
	Heapster                   string `yaml:"heapster"`
	AddonResizer               string `yaml:"addon-resizer"`
	KubernetesDashboard        string `yaml:"kubernetes-dashboard"`
	CertManagerController      string `yaml:"cert-manager-controller"`
	NginxIngressController     string `yaml:"nginx-ingress-controller"`
	NginxIngressDefaultBackend string `yaml:"nginx-ingress-default-backend"`
	MetricsServer              string `yaml:"metrics-server"`
	PrometheusOperator         string `yaml:"prometheus-operator"`
	PrometheusConfigReloader   string `yaml:"prometheus-config-reloader"`
	ConfigMapReload            string `yaml:"configmap-reload"`
	KubeStateMetrics           string `yaml:"kube-state-metrics"`
	Grafana                    string `yaml:"grafana"`
	GrafanaWatcher             string `yaml:"grafana-watcher"`
	Prometheus                 string `yaml:"prometheus"`
	PrometheusNodeExporter     string `yaml:"prometheus-node-exporter"`
	PrometheusAlertManager     string `yaml:"prometheus-alert-manager"`
	CSIAttacher                string `yaml:"csi-attacher"`
	CSIProvisioner             string `yaml:"csi-provisioner"`
	CSIDriverRegistrar         string `yaml:"csi-driver-registrar"`
	CSICephRBDPlugin           string `yaml:"csi-ceph-rbd-plugin"`
	CSICephFSPlugin            string `yaml:"csi-ceph-fs-plugin"`
	WordPress                  string `yaml:"wordpress"`
	MySQL                      string `yaml:"mysql"`
}

func NewVersions() Versions {
	return Versions{
		Etcd:                       utils.VersionEtcd,
		K8S:                        utils.VersionK8s,
		Helm:                       utils.VersionHelm,
		Containerd:                 utils.VersionContainerd,
		Runc:                       utils.VersionRunc,
		CriCtl:                     utils.VersionCrictl,
		Gobetween:                  utils.VersionGobetween,
		Ark:                        utils.VersionArk,
		MinioServer:                utils.VersionMinioServer,
		MinioClient:                utils.VersionMinioClient,
		Pause:                      utils.VersionPause,
		CoreDNS:                    utils.VersionCoredns,
		Elasticsearch:              utils.VersionElasticsearch,
		ElasticsearchCron:          utils.VersionElasticsearchCron,
		ElasticsearchOperator:      utils.VersionElasticsearchOperator,
		Kibana:                     utils.VersionKibana,
		Cerebro:                    utils.VersionCerebro,
		FluentBit:                  utils.VersionFluentBit,
		CalicoTypha:                utils.VersionCalicoTypha,
		CalicoNode:                 utils.VersionCalicoNode,
		CalicoCNI:                  utils.VersionCalicoCni,
		Ceph:                       utils.VersionCeph,
		Heapster:                   utils.VersionHeapster,
		AddonResizer:               utils.VersionAddonResizer,
		KubernetesDashboard:        utils.VersionKubernetesDashboard,
		CertManagerController:      utils.VersionCertManagerController,
		NginxIngressController:     utils.VersionNginxIngressController,
		NginxIngressDefaultBackend: utils.VersionNginxIngressDefaultBackend,
		MetricsServer:              utils.VersionMetricsServer,
		PrometheusOperator:         utils.VersionPrometheusOperator,
		PrometheusConfigReloader:   utils.VersionPrometheusConfigReloader,
		ConfigMapReload:            utils.VersionConfigmapReload,
		KubeStateMetrics:           utils.VersionKubeStateMetrics,
		Grafana:                    utils.VersionGrafana,
		GrafanaWatcher:             utils.VersionGrafanaWatcher,
		Prometheus:                 utils.VersionPrometheus,
		PrometheusNodeExporter:     utils.VersionPrometheusNodeExporter,
		PrometheusAlertManager:     utils.VersionPrometheusAlertManager,
		CSIAttacher:                utils.VersionCsiAttacher,
		CSIProvisioner:             utils.VersionCsiProvisioner,
		CSIDriverRegistrar:         utils.VersionCsiDriverRegistrar,
		CSICephRBDPlugin:           utils.VersionCsiCephRbdPlugin,
		CSICephFSPlugin:            utils.VersionCsiCephFsPlugin,
		WordPress:                  utils.VersionWordpress,
		MySQL:                      utils.VersionMysql,
	}
}
