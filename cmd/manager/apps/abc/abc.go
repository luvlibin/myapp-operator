package abc

import (
	appv1alpha1 "github.com/huzefa51/myapp-operator/pkg/apis/myapp/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	core1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//"k8s.io/apimachinery/pkg/util/intstr"
)


func CreateAbcDeployment(cr *appv1alpha1.MyApp) *appsv1.Deployment {
	label := map[string]string{
		"run": "abc",
	}

	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "abc",
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
							Name:  "abc",
							Image: "tomcat:latest",
							Ports: []core1.ContainerPort{
								{
									ContainerPort: 8080,
								},
							},
						},
					},
				},
			},
		},
	}
}
