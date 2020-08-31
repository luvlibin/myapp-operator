package abc

import (
        "bytes"
        "text/template"

        appv1alpha1 "github.com/huzefa51/myapp-operator/pkg/apis/myapp/v1alpha1"
        //appsv1 "k8s.io/api/apps/v1"
        corev1 "k8s.io/api/core/v1"
        metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
        //appsv1 "k8s.io/api/apps/v1"
        //"k8s.io/apimachinery/pkg/util/intstr"
        "fmt"
        //"io/ioutil"
        //"log"
        //"strings"
)


func CreateConfigMap(cr *appv1alpha1.MyApp) *corev1.ConfigMap {

        //templateInput := TemplateInput{}

	//templateInput.APP_NAME = cr.ObjectMeta.Name 
	//templateInput.NAMESPACE = cr.ObjectMeta.Namespace
	
	//Config struct is defined in configreader.go
	config.APP_NAME = cr.ObjectMeta.Name 
	config.NAMESPACE = cr.ObjectMeta.Namespace

        configMap := generateConfig(config, configmapTemplate)

        var cm = &corev1.ConfigMap{
                TypeMeta: metav1.TypeMeta{
                        Kind:       "ConfigMap",
                        APIVersion: "v1",
                },

                ObjectMeta: metav1.ObjectMeta{
                        Name:      cr.ObjectMeta.Name,
                        Namespace: cr.ObjectMeta.Namespace,
                },

                Data: map[string]string{
                        "fluent.conf": *configMap,
                },
        }
        fmt.Println("cm=",cm)
        return cm
}

// TemplateInput structure
/*type TemplateInput struct {
	name string
	namespace string
}*/


func generateConfig(Config Config, configmapTemplate string) *string {
        //output := new(bytes.Buffer)
	var output bytes.Buffer
        tmpl, err := template.ParseFiles("/var/tmp/java-opts")
        if err != nil {
                return nil
        }
        err = tmpl.Execute(&output, Config)
        outputString := output.String()
        fmt.Println("outputString",outputString)
        return &outputString
}
