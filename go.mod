module github.com/fmy1993/go-gin-admain

go 1.16

require (
	github.com/OwnLocal/goes v1.0.0 // indirect
	github.com/astaxie/beego v1.12.3 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.62.0
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-sqlite3 v2.0.3+incompatible // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18 // indirect
	github.com/siddontang/go v0.0.0-20180604090527-bdc77568d726 // indirect
	github.com/siddontang/ledisdb v0.0.0-20181029004158-becf5f38d373 // indirect
	github.com/ugorji/go v1.2.5 // indirect
	github.com/unknwon/com v1.0.1
	golang.org/x/crypto v0.0.0-20210322153248-0c34fe9e7dc2 // indirect
	golang.org/x/sys v0.0.0-20210331175145-43e1dd70ce54 // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/ini.v1 v1.62.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect

)

//因为还没有上传到git,所以先将使用到的每个包指向本地的路径
replace (
	// 相对go.mod的路径
	// 步骤： 1.在require 添加项目的总包（就是该项目git的域名）  2. 在下方replace 告诉这个总包的路经，最后完全上传就可以去掉replace了
	github.com/fmy1993/go-gin-admain => ./
	//github.com/unknwon/com => github.com/unknwon/com
	github.com/fmy1993/go-gin-admain/conf => ./go/src/go-gin-admin/pkg/conf
	// github.com/fmy1993/go-gin-admain/conf    	  => ~/go-application/go-gin-admain/pkg/conf
	github.com/fmy1993/go-gin-admain/middleware => ./go/src/go-gin-admin/middleware
	github.com/fmy1993/go-gin-admain/models => ./go/src/go-gin-admin/models
	github.com/fmy1993/go-gin-admain/pkg/setting => ./go/src/go-gin-admin/pkg/setting
	github.com/fmy1993/go-gin-admain/pkg/util => ./go/src/go-gin-admin/pkg/util
	github.com/fmy1993/go-gin-admain/routers => ./go/src/go-gin-admin/routers
//github.com/fmy1993/go-gin-admain/routers/api/v1 => ./go/src/go-gin-admin/routers/api/v1

)
