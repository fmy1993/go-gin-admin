module github.com/fmy1993/go-gin-example

go 1.16

require (
	github.com/EDDYCJY/go-gin-example v0.0.0-20201228125222-28f372bf41f9
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.62.0
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/ugorji/go v1.2.5 // indirect
	github.com/unknwon/com v1.0.1
	golang.org/x/crypto v0.0.0-20210322153248-0c34fe9e7dc2 // indirect
	golang.org/x/sys v0.0.0-20210331175145-43e1dd70ce54 // indirect
	gopkg.in/ini.v1 v1.62.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

//因为还没有上传到git,所以先将使用到的每个包指向本地的路径
replace (
	// 相对go.mod的路径
	github.com/fmy1993/go-gin-example/conf => ./go/src/go-gin-admin/pkg/conf
	// github.com/fmy1993/go-gin-example/conf    	  => ~/go-application/go-gin-example/pkg/conf
	github.com/fmy1993/go-gin-example/middleware => ./go/src/go-gin-admin/middleware
	github.com/fmy1993/go-gin-example/models => ./go/src/go-gin-admin/models
	github.com/fmy1993/go-gin-example/pkg/setting => ./go/src/go-gin-admin/pkg/setting
	github.com/fmy1993/go-gin-example/pkg/util => ./go/src/go-gin-admin/pkg/util
	github.com/fmy1993/go-gin-example/routers => ./go/src/go-gin-admin/routers
	github.com/fmy1993/go-gin-example/routers/api/v1 => ./go/src/go-gin-admin/routers/api/v1

)
