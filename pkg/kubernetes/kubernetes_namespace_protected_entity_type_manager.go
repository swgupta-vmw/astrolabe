package kubernetes

import (
	"context"
	"github.com/vmware/arachne/pkg/arachne"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type KubernetesNamespaceProtectedEntityTypeManager struct {
	clientset  *kubernetes.Clientset
	namespaces map[string]*KubernetesNamespaceProtectedEntity
}

func NewKubernetesNamespaceProtectedEntityTypeManagerFromConfig(params map[string]interface{}, s3URLBase string) (*KubernetesNamespaceProtectedEntityTypeManager, error) {
	masterURL := params["masterURL"].(string)
	kubeconfigPath := params["kubeconfigPath"].(string)
	config, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfigPath)
	if err != nil {
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	returnTypeManager := KubernetesNamespaceProtectedEntityTypeManager{
		clientset: clientset,
	}
	returnTypeManager.namespaces = make(map[string]*KubernetesNamespaceProtectedEntity)
	err = returnTypeManager.loadNamespaceEntities()
	if err != nil {
		return nil, err
	}
	return &returnTypeManager, nil
}

func (this *KubernetesNamespaceProtectedEntityTypeManager) GetTypeName() string {
	return "kubernetes-ns"
}

func (this *KubernetesNamespaceProtectedEntityTypeManager) GetProtectedEntity(ctx context.Context, id arachne.ProtectedEntityID) (
	arachne.ProtectedEntity, error) {
	return nil, nil
}

func (this *KubernetesNamespaceProtectedEntityTypeManager) GetProtectedEntities(ctx context.Context) ([]arachne.ProtectedEntity, error) {
	protectedEntities := make([]arachne.ProtectedEntity, 0, len(this.namespaces))

	for _, pe := range this.namespaces {
		protectedEntities = append(protectedEntities, pe)
	}
	return protectedEntities, nil
}

func (this *KubernetesNamespaceProtectedEntityTypeManager) loadNamespaceEntities() error {
	namespaceList, err := this.clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		return err
	}
	for _, namespace := range namespaceList.Items {
		k8snsPE, err := NewKubernetesNamespaceProtectedEntity(this, &namespace)
		if err == nil {
			key := k8snsPE.id.String()
			if _, exists := this.namespaces[key]; !exists {
				this.namespaces[key] = k8snsPE
			}
		} else {
			// log it
		}
	}
	return nil
}

func (this *KubernetesNamespaceProtectedEntityTypeManager) Copy(ctx context.Context, pe arachne.ProtectedEntity) error {
	return nil
}
