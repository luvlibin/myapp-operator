package abc

import (
        appv1alpha1 "github.com/huzefa51/myapp-operator/pkg/apis/myapp/v1alpha1"
        "fmt"
        "io/ioutil"
        "log"
        "gopkg.in/yaml.v2"
        //"github.com/spf13/viper"
)



type Config struct {
	IMAGE string `yaml:"IMAGE"`
	PORT int `yaml:"PORT"`
	APP_NAME string `yaml:"APP_NAME"`
	NAMESPACE string `yaml:"NAMESPACE"`
	SECRET_FILES string `yaml:"SECRET_FILES"`
	CONFIGMAP_FILES string `yaml:"CONFIGMAP_FILES"`
	JAVA_OPTS string
}




func readConfigFile(cr *appv1alpha1.MyApp) *Config {
        inputYamlFile, err := ioutil.ReadFile("/var/tmp/test/env.yaml")
	if err != nil {
		log.Printf("Error while reading env config file %v ", err)
	}

	var config Config
	err = yaml.Unmarshal(inputYamlFile, &config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	fmt.Printf("config Data: \n%v\n", config.SECRET_FILES)
	return &config
}
