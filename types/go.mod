module github.com/sensu/sensu-go/types

go 1.13

replace (
	github.com/sensu/sensu-go => ../
	github.com/sensu/sensu-go/api/core/v2 => ../api/core/v2
)

require (
	github.com/json-iterator/go v1.1.9
	github.com/robertkrimen/otto v0.0.0-20191219234010-c382bd3c16ff
	github.com/sensu/sensu-go/api/core/v2 v2.4.0
	github.com/stretchr/testify v1.6.0
)
