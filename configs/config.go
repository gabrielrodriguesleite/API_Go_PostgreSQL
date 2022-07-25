package configs

import "github.com/spf13/viper"

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

type config struct {
	API APIConfig
	DB  DBConfig
}

// ponteiro para impedir edição de fora desse package,
// apenas acesso a leitura dos dados de configuração
var cfg *config

// função de ciclo de vida do go
func init() {
	// dentro dessa função passamos uma configuração padrão que deve ser alterada posteriormente
	// pacote do banco de dados
	viper.SetDefault("api.port", "9000")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
}

// função que carrega configurações
func Load() error {
	viper.SetConfigName("config") // nome do arquivo de configuração
	viper.SetConfigType("toml")   // toml tipo de arquivo de configuração | viper pode ler outros
	viper.AddConfigPath(".")      // caminho do arquivo de configuração | . = sempre estará ao lado do binário
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok { // o err passou por casting
			return err // se ok for false
		}
	}

	cfg = new(config) // new carrega o endereço assim como o &config, cria um ponteiro

	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}

	cfg.DB = DBConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Pass:     viper.GetString("database.pass"),
		Database: viper.GetString("database.name"),
	}

	return nil
}

func GetDB() DBConfig {
	return cfg.DB
}

func GetServerPort() string {
	return cfg.API.Port
}
