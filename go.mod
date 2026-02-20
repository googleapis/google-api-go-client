module google.golang.org/api

go 1.25.0

retract v0.258.0 // due to https://github.com/googleapis/google-cloud-go/issues/13503

require (
	cloud.google.com/go/auth v0.18.2
	cloud.google.com/go/auth/oauth2adapt v0.2.8
	cloud.google.com/go/compute/metadata v0.9.0
	github.com/google/go-cmp v0.7.0
	github.com/google/s2a-go v0.1.9
	github.com/google/uuid v1.6.0
	github.com/googleapis/enterprise-certificate-proxy v0.3.12
	github.com/googleapis/gax-go/v2 v2.17.0
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.61.0
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.61.0
	golang.org/x/net v0.50.0
	golang.org/x/oauth2 v0.35.0
	golang.org/x/sync v0.19.0
	golang.org/x/time v0.14.0
	google.golang.org/genproto/googleapis/bytestream v0.0.0-20260217215200-42d3e9bedb6d
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260217215200-42d3e9bedb6d
	google.golang.org/grpc v1.79.1
	google.golang.org/protobuf v1.36.11
)

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/go-logr/logr v1.4.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	go.opentelemetry.io/auto/sdk v1.2.1 // indirect
	go.opentelemetry.io/otel v1.39.0 // indirect
	go.opentelemetry.io/otel/metric v1.39.0 // indirect
	go.opentelemetry.io/otel/trace v1.39.0 // indirect
	golang.org/x/crypto v0.48.0 // indirect
	golang.org/x/sys v0.41.0 // indirect
	golang.org/x/text v0.34.0 // indirect
)
