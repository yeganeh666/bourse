package configs

type Configs struct {
	ServiceName string          `mapstructure:"service_name"`
	Version     string          `mapstructure:"version"`
	GRPCPort    uint16          `mapstructure:"grpc_port"`
	HttpPort    uint16          `mapstructure:"http_port"`
	HttpHost    string          `mapstructure:"http_host"`
	Postgres    SectionPostgres `yaml:"postgres"`
}
type SectionPostgres struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	DB   string `yaml:"db"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
}

var (
	DefaultConfig = []byte(`
service_name: "trades"
version: "v1.0.0"
grpc_port: 26000
http_port: 8080
http_host: "localhost"
postgres:
  host: "localhost"
  port: 5432
  db: "bourse"
  user: "postgres"
  pass: "postgres"
`)
)
