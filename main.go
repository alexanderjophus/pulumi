package main

import (
	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/container"
	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/projects"
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/helm/v3"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	argoproj "github.com/trelore/pulumi/crds/argoproj/v1alpha1"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		containerService, err := projects.NewService(ctx, "project", &projects.ServiceArgs{
			Service: pulumi.String("container.googleapis.com"),
		})
		if err != nil {
			return err
		}

		engineVersions, err := container.GetEngineVersions(ctx, &container.GetEngineVersionsArgs{})
		if err != nil {
			return err
		}
		masterVersion := engineVersions.LatestMasterVersion

		cluster, err := container.NewCluster(ctx, "demo-cluster", &container.ClusterArgs{
			InitialNodeCount: pulumi.Int(2),
			MinMasterVersion: pulumi.String(masterVersion),
			NodeVersion:      pulumi.String(masterVersion),
			NodeConfig: &container.ClusterNodeConfigArgs{
				MachineType: pulumi.String("n1-standard-1"),
				OauthScopes: pulumi.StringArray{
					pulumi.String("https://www.googleapis.com/auth/compute"),
					pulumi.String("https://www.googleapis.com/auth/devstorage.read_only"),
					pulumi.String("https://www.googleapis.com/auth/logging.write"),
					pulumi.String("https://www.googleapis.com/auth/monitoring"),
				},
			},
		}, pulumi.DependsOn([]pulumi.Resource{containerService}))
		if err != nil {
			return err
		}

		ctx.Export("kubeconfig", generateKubeconfig(cluster.Endpoint, cluster.Name, cluster.MasterAuth))

		k8sProvider, err := kubernetes.NewProvider(ctx, "k8sprovider", &kubernetes.ProviderArgs{
			Kubeconfig: generateKubeconfig(cluster.Endpoint, cluster.Name, cluster.MasterAuth),
		}, pulumi.DependsOn([]pulumi.Resource{cluster}))
		if err != nil {
			return err
		}

		if err = createMonitoring(ctx, k8sProvider); err != nil {
			return err
		}

		if err = createArgo(ctx, k8sProvider); err != nil {
			return err
		}

		return nil
	})
}

func createMonitoring(ctx *pulumi.Context, k8sProvider pulumi.ProviderResource) error {
	ns, err := corev1.NewNamespace(ctx, "monitoring", &corev1.NamespaceArgs{
		Metadata: &metav1.ObjectMetaArgs{
			Name: pulumi.String("monitoring"),
		},
	}, pulumi.Provider(k8sProvider))
	if err != nil {
		return err
	}

	_, err = helm.NewChart(ctx, "loki-stack", helm.ChartArgs{
		Chart:     pulumi.String("loki-stack"),
		Version:   pulumi.String("2.5.0"),
		Namespace: ns.Metadata.Name().Elem(),
		FetchArgs: helm.FetchArgs{
			Repo: pulumi.String("https://grafana.github.io/helm-charts"),
		},
		Values: pulumi.Map{
			"loki": pulumi.Map{
				"enabled": pulumi.Bool(true),
			},
			"grafana": pulumi.Map{
				"enabled": pulumi.Bool(true),
			},
			"prometheus": pulumi.Map{
				"enabled": pulumi.Bool(true),
			},
		},
	}, pulumi.Provider(k8sProvider))
	if err != nil {
		return err
	}

	return nil
}

func createArgo(ctx *pulumi.Context, k8sProvider pulumi.ProviderResource) error {
	ns, err := corev1.NewNamespace(ctx, "argo", &corev1.NamespaceArgs{
		Metadata: &metav1.ObjectMetaArgs{
			Name: pulumi.String("argo"),
		},
	}, pulumi.Provider(k8sProvider))
	if err != nil {
		return err
	}

	chart, err := helm.NewChart(ctx, "argo", helm.ChartArgs{
		Chart:     pulumi.String("argo-cd"),
		Version:   pulumi.String("3.26.9"),
		Namespace: ns.Metadata.Name().Elem(),
		FetchArgs: helm.FetchArgs{
			Repo: pulumi.String("https://argoproj.github.io/argo-helm"),
		},
	}, pulumi.Provider(k8sProvider))
	if err != nil {
		return err
	}

	argoproj.NewApplication(ctx, "application", &argoproj.ApplicationArgs{
		Metadata: metav1.ObjectMetaArgs{
			Namespace: ns.Metadata.Name().Elem(),
		},
		Spec: argoproj.ApplicationSpecArgs{
			Destination: argoproj.ApplicationSpecDestinationArgs{
				Namespace: ns.Metadata.Name().Elem(),
				Server:    pulumi.String("https://kubernetes.default.svc"),
			},
			Source: argoproj.ApplicationSpecSourceArgs{
				RepoURL:        pulumi.String("https://github.com/trelore/iris-classification"),
				Path:           pulumi.String("kustomize"),
				TargetRevision: pulumi.String("HEAD"),
			},
			SyncPolicy: argoproj.ApplicationSpecSyncPolicyArgs{
				Automated: argoproj.ApplicationSpecSyncPolicyAutomatedArgs{
					Prune: pulumi.Bool(true),
				},
			},
			Project: pulumi.String("default"),
		},
	}, pulumi.Provider(k8sProvider), pulumi.DependsOn([]pulumi.Resource{chart}))

	return nil
}

func generateKubeconfig(clusterEndpoint pulumi.StringOutput, clusterName pulumi.StringOutput,
	clusterMasterAuth container.ClusterMasterAuthOutput) pulumi.StringOutput {
	context := pulumi.Sprintf("demo_%s", clusterName)

	return pulumi.Sprintf(`apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: %s
    server: https://%s
  name: %s
contexts:
- context:
    cluster: %s
    user: %s
  name: %s
current-context: %s
kind: Config
preferences: {}
users:
- name: %s
  user:
    auth-provider:
      config:
        cmd-args: config config-helper --format=json
        cmd-path: gcloud
        expiry-key: '{.credential.token_expiry}'
        token-key: '{.credential.access_token}'
      name: gcp`,
		clusterMasterAuth.ClusterCaCertificate().Elem(),
		clusterEndpoint, context, context, context, context, context, context)
}
