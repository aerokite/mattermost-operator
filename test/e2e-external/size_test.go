package e2e

import (
	"context"
	mmv1alpha "github.com/mattermost/mattermost-operator/apis/mattermost/v1alpha1"
	mmv1beta "github.com/mattermost/mattermost-operator/apis/mattermost/v1beta1"
	"github.com/mattermost/mattermost-operator/test/e2e"
	"github.com/stretchr/testify/require"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"testing"
)

// TestMattermostSize checks defaulting & updating replicas & resources from size.
func TestMattermostSize(t *testing.T) {
	namespace := "e2e-test-size"
	name := "test-mm"
	mmNamespaceName := types.NamespacedName{Namespace: namespace, Name: name}

	testEnv, setupErr := SetupTestEnv(k8sClient, namespace)
	require.NoError(t, setupErr)
	defer testEnv.CleanupFunc()

	mattermost := &mmv1beta.Mattermost{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: mmv1beta.MattermostSpec{
			Ingress: &mmv1beta.Ingress{
				Host: "e2e-test-example.mattermost.dev",
			},
			FileStore: mmv1beta.FileStore{
				External: &testEnv.FileStoreConfig,
			},
			Database: mmv1beta.Database{
				External: &testEnv.DBConfig,
			},
		},
	}

	mmSize := mmv1alpha.CloudSize10String
	mattermost.Spec.Size = mmSize
	instance := e2e.NewMattermostInstance(t, k8sClient, mattermost)

	{
		clusterSize, err := mmv1alpha.GetClusterSize(mmSize)
		require.NoError(t, err)

		t.Log("create and waiting for Mattermost to be stable")
		instance.CreateAndWait()
		defer instance.Destroy()

		t.Log("checking mattermost replicas & resources")
		var newMattermost mmv1beta.Mattermost
		err = k8sClient.Get(context.TODO(), mmNamespaceName, &newMattermost)
		require.NoError(t, err)
		// Size should be erased
		require.Empty(t, newMattermost.Spec.Size)
		// Check Replicas & Resources, set by Size
		require.NotNil(t, newMattermost.Spec.Replicas)
		t.Logf("mattermost replicas & resources should match %s\n", mmSize)
		require.Equal(t, clusterSize.App.Replicas, *newMattermost.Spec.Replicas)
		compareResources(t, clusterSize.App.Resources, newMattermost.Spec.Scheduling.Resources)

		t.Log("checking deployment replicas & resources")
		var mmDeployment appsv1.Deployment
		err = k8sClient.Get(context.TODO(), mmNamespaceName, &mmDeployment)
		require.NoError(t, err)
		require.Equal(t, clusterSize.App.Replicas, *mmDeployment.Spec.Replicas)
		// compare resources in deployment
		compareResources(t, clusterSize.App.Resources, mmDeployment.Spec.Template.Spec.Containers[0].Resources)
	}

	{
		mmSize := mmv1alpha.CloudSize100String
		clusterSize, err := mmv1alpha.GetClusterSize(mmSize)
		require.NoError(t, err)

		t.Logf("updating scheduling resources in mattermost object with %s\n", mmSize)
		var newMattermost mmv1beta.Mattermost
		err = k8sClient.Get(context.TODO(), mmNamespaceName, &newMattermost)
		require.NoError(t, err)
		newMattermost.Spec.Scheduling.Resources = clusterSize.App.Resources
		instance.UpdateAndWait(&newMattermost)

		// compare resources with deployment
		t.Log("checking deployment resources")
		var mmDeployment appsv1.Deployment
		err = k8sClient.Get(context.TODO(), mmNamespaceName, &mmDeployment)
		require.NoError(t, err)
		compareResources(t, clusterSize.App.Resources, mmDeployment.Spec.Template.Spec.Containers[0].Resources)
	}

	{
		mmSize := mmv1alpha.Size100String
		clusterSize, err := mmv1alpha.GetClusterSize(mmSize)
		require.NoError(t, err)

		t.Logf("updating size in mattermost object with %s\n", mmSize)
		var newMattermost mmv1beta.Mattermost
		err = k8sClient.Get(context.TODO(), mmNamespaceName, &newMattermost)
		require.NoError(t, err)

		// update size in mattermost
		newMattermost.Spec.Size = mmSize
		instance.UpdateAndWait(&newMattermost)

		// compare replicas & resources
		t.Log("checking deployment resources")
		var mmDeployment appsv1.Deployment
		err = k8sClient.Get(context.TODO(), mmNamespaceName, &mmDeployment)
		require.NoError(t, err)
		require.Equal(t, clusterSize.App.Replicas, *mmDeployment.Spec.Replicas)
		compareResources(t, clusterSize.App.Resources, mmDeployment.Spec.Template.Spec.Containers[0].Resources)
	}

}