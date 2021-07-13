module github.com/aws/aws-sdk-go-v2/service/dynamodb

go 1.15

require (
	github.com/aws/aws-sdk-go-v2 v1.7.0
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.0.0
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.2.0
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.0.0
	github.com/aws/smithy-go v1.5.1-0.20210713172319-36596d2cf389
	github.com/jmespath/go-jmespath v0.4.0
)

replace github.com/aws/aws-sdk-go-v2 => ../../

replace github.com/aws/aws-sdk-go-v2/internal/configsources => ../../internal/configsources/

replace github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding => ../../service/internal/accept-encoding/

replace github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery => ../../service/internal/endpoint-discovery/
