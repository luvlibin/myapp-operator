package abc

import (
	appv1alpha1 "github.com/huzefa51/myapp-operator/pkg/apis/myapp/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	core1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//"k8s.io/apimachinery/pkg/util/intstr"
	//"os"
	"fmt"
	//"strconv"
	//"io/ioutil"
	//"log"

)



func CreateAbcDeployment(cr *appv1alpha1.MyApp) *appsv1.Deployment {

	configProperty := readConfigFile(cr)
	fmt.Println("name from cr ",cr.ObjectMeta.Name)
	fmt.Println("namespace from cr", cr.ObjectMeta.Namespace)
	fmt.Println("conf data=",configProperty.IMAGE)
        fmt.Println("conf port=",configProperty.PORT)

	label := map[string]string{
		"run": "abc",
	}
	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.ObjectMeta.Name,
			Namespace: cr.ObjectMeta.Namespace,
			Labels:    label,
		},

		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: label,
			},

			Template: core1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: label,
				},

				Spec: core1.PodSpec{
					Containers: []core1.Container{
						{
							Name:  cr.ObjectMeta.Name,
							Image: configProperty.IMAGE,
							Ports: []core1.ContainerPort{
								{
									ContainerPort: int32(configProperty.PORT),
								},
							},
						},
					},
				},
			},
		},
	}
}
