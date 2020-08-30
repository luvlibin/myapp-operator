package abc 

import (
	"bytes"
	"text/template"

	appv1alpha1 "github.com/huzefa51/myapp-operator/pkg/apis/myapp/v1alpha1"
	//appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	appsv1 "k8s.io/api/apps/v1"
	//"k8s.io/apimachinery/pkg/util/intstr"
	"fmt"
        "io/ioutil"
        "log"
	"strings"
)


func CreateAbcDeployment(cr *appv1alpha1.MyApp) *appsv1.Deployment  {
        content, err := ioutil.ReadFile("/var/tmp/test/deployment.json")

        if err != nil {
                log.Fatal(err)
        }

        fmt.Println(string(content))

        /*label := map[string]string{
                "run": "abc",
        }*/
	d := &appsv1.Deployment{}
        return d
}
// CreateConfigMap for FluentD
func CreateConfigMap(cr *appv1alpha1.MyApp) *corev1.ConfigMap {

	templateInput := TemplateInput{}

	configMap := generateConfig(templateInput, configmapTemplate)

	var cm = &corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},

		ObjectMeta: metav1.ObjectMeta{
			Name:      "fluentd-config",
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
type TemplateInput struct {
	Inputs []Input
}

// Input structure
type Input struct {
	Tag     string
	Outputs []Output
}

// Output spec
type Output struct {
	Type         string
	IndexPattern string
}

func generateConfig(TemplateInput TemplateInput, configmapTemplate string) *string {
	output := new(bytes.Buffer)
	tmpl, err := template.ParseFiles("/var/tmp/java-opts")
	if err != nil {
		return nil
	}
	err = tmpl.Execute(output, TemplateInput)
	outputString := output.String()
	fmt.Println("outputString",outputString)
	return &outputString
}

func generateSecret(TemplateInput TemplateInput, secret_Files string) *string {
//func generateSecret(TemplateInput TemplateInput) *string {
	output := new(bytes.Buffer)
	//tmpl, err := template.ParseFiles("/var/tmp/test/abc.txt")
	tmpl, err := template.ParseFiles("/var/tmp/test/"+secret_Files)
	if err != nil {
		return nil
	}
	err = tmpl.Execute(output, TemplateInput)
	//outputString := output.Bytes()
	outputString := output.String()
	fmt.Println("outputSecretString",outputString)
	return &outputString
}



// Create Secret
func CreateSecret(cr *appv1alpha1.MyApp) *corev1.Secret {

	//Get the SECRET FILES name by reading from config file
	con := readConfigFile(cr)
	fmt.Println("conf data=",con.SECRET_FILES)
	secret_Files := strings.Split(con.SECRET_FILES, ",")
	secret_Values := make(map[string]string)

	templateInput := TemplateInput{}
	secret := make(map[string]interface{})
	for key,value := range secret_Files {
	
		fmt.Println("key is:", key)
		fmt.Println("value is:", value)
		secret_Values[value] = *generateSecret(templateInput, value)
		//secret_Values[value] = "hello"
		fmt.Println("output is ", *generateSecret(templateInput, value))
		//secret_Values := *value[key]

	}
	fmt.Println("map values-",secret_Values)



	//templateInput := TemplateInput{}

	//secret := generateSecret(templateInput)
	fmt.Println("secret is:",secret)
	var sec = &corev1.Secret{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Secret",
			APIVersion: "v1",
		},

		ObjectMeta: metav1.ObjectMeta{
			Name:      "myapp-secret",
			Namespace: cr.ObjectMeta.Namespace,
		},

		/*Data: map[string][]byte{
			"abc.txt": []byte(*secret),
		},*/
		/*StringData: map[string]string{
			"abc.txt": "abcd",
			//secret_Values,
		},*/
		StringData: secret_Values,
	}
	fmt.Println("cm=",sec)
	return sec
}

