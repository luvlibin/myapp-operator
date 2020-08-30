module github.com/huzefa51/myapp-operator

go 1.13

require (
	github.com/Masterminds/semver v1.5.0 // indirect
	github.com/boltdb/bolt v1.3.1 // indirect
	github.com/go-git/go-git/v5 v5.1.0
	github.com/golang/dep v0.5.4 // indirect
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/jmank88/nuts v0.4.0 // indirect
	github.com/nightlyone/lockfile v1.0.0 // indirect
	github.com/operator-framework/operator-sdk v0.18.2
	github.com/pelletier/go-toml v1.8.0 // indirect
	github.com/sdboyer/constext v0.0.0-20170321163424-836a14457353 // indirect
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.4.0
	golang.org/x/sync v0.0.0-20200625203802-6e8e738ad208 // indirect
	golang.org/x/sys v0.0.0-20200814200057-3d37ad5750ed // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/yaml.v2 v2.3.0
	k8s.io/api v0.18.2
	k8s.io/apimachinery v0.18.2
	k8s.io/client-go v12.0.0+incompatible
	sigs.k8s.io/controller-runtime v0.6.0
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v13.3.2+incompatible // Required by OLM
	k8s.io/client-go => k8s.io/client-go v0.18.2 // Required by prometheus-operator
)
