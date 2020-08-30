package apps

import (
	"github.com/huzefa51/myapp-operator/cmd/manager/apps/abc"
	appv1alpha1 "github.com/huzefa51/myapp-operator/pkg/apis/myapp/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	//rbacv1 "k8s.io/api/rbac/v1"
	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Apps structure declarations
type Apps struct {
	cr            *appv1alpha1.MyApp
	Abc           Abc
}

func (t *Apps) init() {
	t.Abc = Abc{
		cr: t.cr,
	}
}

// ElasticSearch structure
type Abc struct {
	cr *appv1alpha1.MyApp
}

// GetDeployment returns ES deployment
func (e Abc) GetDeployment()(*appsv1.Deployment, *appsv1.Deployment) {
	return &appsv1.Deployment{}, abc.CreateAbcDeployment(e.cr)
}

// GetConfigMap returns FluentD configmap
func (f *Abc) GetConfigMap() (*corev1.ConfigMap, *corev1.ConfigMap) {
	return &corev1.ConfigMap{}, abc.CreateConfigMap(f.cr)
}
// GetSecret returns secret
func (s *Abc) GetSecret() (*corev1.Secret, *corev1.Secret) {
	return &corev1.Secret{}, abc.CreateSecret(s.cr)
}

// GetTools returns an instance of Tools
func GetApps(customResource *appv1alpha1.MyApp) *Apps {
	apps := Apps{
		cr: customResource,
	}
	apps.init()
	return &apps
}
