module google.golang.org/api

go 1.26.0

retract v0.258.0 // due to https://github.com/googleapis/google-cloud-go/issues/13503

require (
	cloud.google.com/go/auth v0.23.4-0.20260922060409-6c19dd8b7152
	cloud.google.com/go/auth/oauth2adapt v0.2.8
	cloud.google.com/go/compute/metadata v0.9.1
	github.com/google/go-cmp v0.7.0
	github.com/google/s2a-go v0.1.10
	github.com/google/uuid v1.6.0
	github.com/googleapis/enterprise-certificate-proxy v0.3.22
	github.com/googleapis/gax-go/v2 v2.24.1
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.67.0
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.69.0
	golang.org/x/net v0.59.0
	golang.org/x/oauth2 v0.37.0
	golang.org/x/sync v0.23.0
	golang.org/x/time v0.16.0
	google.golang.org/genproto/googleapis/bytestream v0.0.0-20260921155816-b14227669459
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260921155816-b14227669459
	google.golang.org/grpc v1.84.0
	google.golang.org/protobuf v1.36.12
)

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/felixge/httpsnoop v1.1.0 // indirect
	github.com/go-logr/logr v1.4.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	go.opentelemetry.io/auto/sdk v1.2.1 // indirect
	go.opentelemetry.io/otel v1.44.0 // indirect
	go.opentelemetry.io/otel/metric v1.44.0 // indirect
	go.opentelemetry.io/otel/trace v1.44.0 // indirect
	golang.org/x/crypto v0.57.0 // indirect
	golang.org/x/sys v0.48.0 // indirect
	golang.org/x/text v0.42.0 // indirect
)
