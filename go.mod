module github.com/percona/pmm

go 1.20

// Update saas with
// go get -v github.com/percona-platform/saas@latest

replace github.com/grpc-ecosystem/go-grpc-prometheus => github.com/Percona-Lab/go-grpc-prometheus v0.0.0-20230116133345-3487748d4592

replace github.com/go-openapi/spec => github.com/Percona-Lab/spec v0.20.5-percona

replace gopkg.in/alecthomas/kingpin.v2 => github.com/Percona-Lab/kingpin v2.2.6-percona+incompatible

replace golang.org/x/crypto => github.com/percona-lab/crypto v0.0.0-20220811043533-d164de3c7f08

replace github.com/ClickHouse/clickhouse-go/151 => github.com/ClickHouse/clickhouse-go v1.5.1 // clickhouse-go/v2 cannot work with 1.5.1 which we need for QAN-API

require (
	github.com/AlekSi/pointer v1.2.0
	github.com/ClickHouse/clickhouse-go/151 v0.0.0-00010101000000-000000000000
	github.com/ClickHouse/clickhouse-go/v2 v2.9.2
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/alecthomas/kong v0.7.1
	github.com/alecthomas/units v0.0.0-20211218093645-b94a6e3cc137
	github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2
	github.com/aws/aws-sdk-go v1.44.259
	github.com/blang/semver v3.5.1+incompatible
	github.com/brianvoe/gofakeit/v6 v6.21.0
	github.com/charmbracelet/bubbles v0.15.0
	github.com/charmbracelet/bubbletea v0.23.2
	github.com/charmbracelet/lipgloss v0.7.1
	github.com/davecgh/go-spew v1.1.1
	github.com/docker/docker v20.10.23+incompatible
	github.com/docker/go-connections v0.4.0
	github.com/envoyproxy/protoc-gen-validate v1.0.0
	github.com/go-co-op/gocron v1.25.0
	github.com/go-openapi/errors v0.20.3
	github.com/go-openapi/runtime v0.26.0
	github.com/go-openapi/strfmt v0.21.7
	github.com/go-openapi/swag v0.22.3
	github.com/go-openapi/validate v0.22.1
	github.com/go-sql-driver/mysql v1.7.1
	github.com/golang-migrate/migrate/v4 v4.15.2
	github.com/golang/protobuf v1.5.3
	github.com/google/uuid v1.3.0
	github.com/grafana/grafana-api-golang-client v0.21.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.2
	github.com/hashicorp/go-version v1.6.0
	github.com/jmoiron/sqlx v1.3.5
	github.com/jotaen/kong-completion v0.0.5
	github.com/lib/pq v1.10.6
	github.com/minio/minio-go/v7 v7.0.52
	github.com/operator-framework/api v0.17.5
	github.com/operator-framework/operator-lifecycle-manager v0.24.0
	github.com/percona-platform/dbaas-api v0.0.0-20230103182808-d79c449a9f4c
	github.com/percona-platform/saas v0.0.0-20230306173543-c223f9a47342
	github.com/percona/dbaas-operator v0.1.6
	github.com/percona/exporter_shared v0.7.4
	github.com/percona/go-mysql v0.0.0-20210427141028-73d29c6da78c
	github.com/percona/percona-toolkit v3.2.1+incompatible
	github.com/percona/promconfig v0.2.5
	github.com/pganalyze/pg_query_go/v2 v2.2.0
	github.com/pkg/errors v0.9.1
	github.com/pmezard/go-difflib v1.0.0
	github.com/prometheus/alertmanager v0.25.0
	github.com/prometheus/client_golang v1.15.1
	github.com/prometheus/common v0.43.0
	github.com/ramr/go-reaper v0.2.1
	github.com/robfig/cron/v3 v3.0.1
	github.com/sirupsen/logrus v1.9.0
	github.com/stretchr/objx v0.5.0
	github.com/stretchr/testify v1.8.2
	go.mongodb.org/mongo-driver v1.11.6
	go.starlark.net v0.0.0-20220328144851-d1966c6b9fcd
	golang.org/x/crypto v0.6.0
	golang.org/x/sync v0.2.0
	golang.org/x/sys v0.8.0
	golang.org/x/text v0.9.0
	golang.org/x/tools v0.8.0
	google.golang.org/genproto v0.0.0-20230306155012-7f2fa6fef1f4
	google.golang.org/grpc v1.56.0-dev
	google.golang.org/protobuf v1.30.0
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	gopkg.in/reform.v1 v1.5.1
	gopkg.in/yaml.v3 v3.0.1
	k8s.io/api v0.27.1
	k8s.io/apiextensions-apiserver v0.26.2
	k8s.io/apimachinery v0.27.1
	k8s.io/cli-runtime v0.27.1
	k8s.io/client-go v0.27.1
	k8s.io/kubectl v0.27.1
	modernc.org/sqlite v1.22.1
	sigs.k8s.io/controller-runtime v0.14.6
)

require (
	github.com/Azure/azure-pipeline-go v0.2.3 // indirect
	github.com/Azure/azure-storage-blob-go v0.14.0 // indirect
	github.com/aymanbagabas/go-osc52/v2 v2.0.1 // indirect
	github.com/benbjohnson/clock v1.3.0 // indirect
	github.com/blang/semver/v4 v4.0.0 // indirect
	github.com/coreos/go-systemd/v22 v22.4.0 // indirect
	github.com/emicklei/go-restful/v3 v3.9.0 // indirect
	github.com/evanphx/json-patch v5.6.0+incompatible // indirect
	github.com/evanphx/json-patch/v5 v5.6.0 // indirect
	github.com/go-errors/errors v1.4.2 // indirect
	github.com/go-ini/ini v1.67.0 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/google/gnostic v0.6.9 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510 // indirect
	github.com/hashicorp/go-uuid v1.0.2 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51 // indirect
	github.com/mattn/go-ieproxy v0.0.1 // indirect
	github.com/minio/minio-go v6.0.14+incompatible // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/monochromegane/go-gitignore v0.0.0-20200626010858-205db1a8cc00 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/mwitkow/go-proto-validators v0.3.2 // indirect
	github.com/percona/percona-backup-mongodb v1.8.1 // indirect
	github.com/pierrec/lz4 v2.6.1+incompatible // indirect
	github.com/posener/complete v1.2.3 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	github.com/riywo/loginshell v0.0.0-20200815045211-7d26008be1ab // indirect
	github.com/sergi/go-diff v1.2.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/xlab/treeprint v1.1.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	k8s.io/klog/v2 v2.90.1 // indirect
	k8s.io/kube-openapi v0.0.0-20230308215209-15aac26d736a // indirect
	k8s.io/utils v0.0.0-20230209194617-a36077c30491 // indirect
	lukechampine.com/uint128 v1.2.0 // indirect
	modernc.org/cc/v3 v3.40.0 // indirect
	modernc.org/ccgo/v3 v3.16.13 // indirect
	modernc.org/libc v1.22.5 // indirect
	modernc.org/mathutil v1.5.0 // indirect
	modernc.org/memory v1.5.0 // indirect
	modernc.org/opt v0.1.3 // indirect
	modernc.org/strutil v1.1.3 // indirect
	modernc.org/token v1.0.1 // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/kustomize/api v0.13.2 // indirect
	sigs.k8s.io/kustomize/kyaml v0.14.1 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.3 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)

require (
	github.com/Azure/azure-sdk-for-go/sdk/azcore v1.4.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/azidentity v1.2.2
	github.com/Azure/azure-sdk-for-go/sdk/internal v1.2.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcegraph/armresourcegraph v0.7.1
	github.com/AzureAD/microsoft-authentication-library-for-go v0.9.0 // indirect
	github.com/ClickHouse/ch-go v0.52.1 // indirect
	github.com/ClickHouse/clickhouse-go v1.5.4 // indirect
	github.com/HdrHistogram/hdrhistogram-go v1.1.2
	github.com/Microsoft/go-winio v0.5.2 // indirect
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751 // indirect
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/armon/go-metrics v0.4.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cenkalti/backoff/v4 v4.2.0 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/charmbracelet/harmonica v0.2.0 // indirect
	github.com/cloudflare/golz4 v0.0.0-20150217214814-ef862a3cdc58 // indirect
	github.com/containerd/console v1.0.3 // indirect
	github.com/docker/distribution v2.8.1+incompatible // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/go-faster/city v1.0.1 // indirect
	github.com/go-faster/errors v0.6.1 // indirect
	github.com/go-kit/log v0.2.1 // indirect
	github.com/go-logfmt/logfmt v0.5.1 // indirect
	github.com/go-openapi/analysis v0.21.4 // indirect
	github.com/go-openapi/jsonpointer v0.19.6 // indirect
	github.com/go-openapi/jsonreference v0.20.1 // indirect
	github.com/go-openapi/loads v0.21.2 // indirect
	github.com/go-openapi/spec v0.20.8 // indirect
	github.com/gofrs/uuid v4.3.1+incompatible // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt/v4 v4.5.0 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/btree v1.1.2 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-msgpack v0.5.5 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-sockaddr v1.0.2 // indirect
	github.com/hashicorp/golang-lru v0.6.0 // indirect
	github.com/hashicorp/memberlist v0.5.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/jpillora/backoff v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.16.0 // indirect
	github.com/klauspost/cpuid/v2 v2.2.4 // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/mattn/go-localereader v0.0.1 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/miekg/dns v1.1.50 // indirect
	github.com/minio/md5-simd v1.1.2 // indirect
	github.com/minio/sha256-simd v1.0.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/montanaflynn/stats v0.7.0 // indirect
	github.com/muesli/ansi v0.0.0-20211018074035-2e021307bc4b // indirect
	github.com/muesli/cancelreader v0.2.2 // indirect
	github.com/muesli/reflow v0.3.0 // indirect
	github.com/muesli/termenv v0.15.1 // indirect
	github.com/mwitkow/go-conntrack v0.0.0-20190716064945-2f068394615f // indirect
	github.com/oklog/run v1.1.0 // indirect
	github.com/oklog/ulid v1.3.1 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.3-0.20211202183452-c5a74bcca799 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/paulmach/orb v0.9.0 // indirect
	github.com/pierrec/lz4/v4 v4.1.17 // indirect
	github.com/pkg/browser v0.0.0-20210911075715-681adbf594b8 // indirect
	github.com/prometheus/client_model v0.4.0 // indirect
	github.com/prometheus/common/sigv4 v0.1.0 // indirect
	github.com/prometheus/exporter-toolkit v0.8.2 // indirect
	github.com/prometheus/procfs v0.9.0 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/rs/xid v1.4.0 // indirect
	github.com/sean-/seed v0.0.0-20170313163322-e2103e2c3529 // indirect
	github.com/segmentio/asm v1.2.0 // indirect
	github.com/shopspring/decimal v1.3.1 // indirect
	github.com/shurcooL/httpfs v0.0.0-20190707220628-8d4bc4ba7749 // indirect
	github.com/shurcooL/vfsgen v0.0.0-20200824052919-0d455de96546 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.1 // indirect
	github.com/xdg-go/stringprep v1.0.3 // indirect
	github.com/youmark/pkcs8 v0.0.0-20201027041543-1326539a0a0a // indirect
	go.opentelemetry.io/otel v1.14.0 // indirect
	go.opentelemetry.io/otel/trace v1.14.0 // indirect
	golang.org/x/mod v0.10.0 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/oauth2 v0.7.0 // indirect
	golang.org/x/term v0.7.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gotest.tools/v3 v3.3.0 // indirect
)
