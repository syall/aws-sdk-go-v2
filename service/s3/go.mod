module github.com/aws/aws-sdk-go-v2/service/s3

go 1.19

require (
	github.com/aws/aws-sdk-go-v2 v1.23.5
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.5.3
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.2.8
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.5.8
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.2.8
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.10.3
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.2.8
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.10.8
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.16.8
	github.com/aws/smithy-go v1.18.1
	github.com/google/go-cmp v0.5.8
)

replace github.com/aws/aws-sdk-go-v2 => ../../

replace github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream => ../../aws/protocol/eventstream/

replace github.com/aws/aws-sdk-go-v2/internal/configsources => ../../internal/configsources/

replace github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 => ../../internal/endpoints/v2/

replace github.com/aws/aws-sdk-go-v2/internal/v4a => ../../internal/v4a/

replace github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding => ../../service/internal/accept-encoding/

replace github.com/aws/aws-sdk-go-v2/service/internal/checksum => ../../service/internal/checksum/

replace github.com/aws/aws-sdk-go-v2/service/internal/presigned-url => ../../service/internal/presigned-url/

replace github.com/aws/aws-sdk-go-v2/service/internal/s3shared => ../../service/internal/s3shared/
