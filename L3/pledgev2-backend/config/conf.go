package config

var Config *Conf

type Conf struct {
	MySql        MysqlConf
	Redis        RedisConf
	TestNet      TestNetConfig
	MainNet      MainNetConfig
	Token        TokenConfig
	Email        EmailConfig
	DefaultAdmin DefaultAdminConfig
	Threshold    ThresholdConfig
	Jwt          JwtConfig
	Env          EnvConfig
}

type EnvConfig struct {
	Port               string `toml:"port"`
	Version            string `toml:"version"`
	Protocol           string `toml:"protocol"`
	DomainName         string `toml:"domain_name"`
	TaskDuration       int64  `toml:"task_duration"`
	WssTimeoutDuration int64  `toml:"ws_timeout_duration"`
	TaskExtendDuration int64  `toml:"task_extend_duration"`
}

type JwtConfig struct {
	SecretKey  string `toml:"secret_key"`
	ExpireTime int    `toml:"expire_time"`
}

type ThresholdConfig struct {
	PledgePoolTokenThresholdBnb string `toml:"pledge_pool_token_threshold_bnb"`
}

type DefaultAdminConfig struct {
	UserName string `toml:"user_name"`
	Password string `toml:"password"`
}

type EmailConfig struct {
	UserName string   `toml:"user_name"`
	Password string   `toml:"password"`
	Host     string   `toml:"host"`
	Port     string   `toml:"port"`
	From     string   `toml:"from"`
	Subject  string   `toml:"subject"`
	To       []string `toml:"to"`
	Cc       []string `toml:"cc"`
}

type TokenConfig struct {
	LogoUrl string `json:"logo_url" toml:"logo_url"`
}

type MainNetConfig struct {
	ChainId              string `json:"chain_id" toml:"chain_id"`
	NetUrl               string `json:"net_url" toml:"net_url"`
	PlgrAddress          string `json:"plgr_address" toml:"plgr_address"`
	PledgePoolToken      string `json:"pledge_pool_token" toml:"pledge_pool_token"`
	BscPledgeOracleToken string `json:"bsc_pledge_oracle_token" toml:"bsc_pledge_oracle_token"`
}

type TestNetConfig struct {
	ChainId              string `json:"chain_id" toml:"chain_id"`
	NetUrl               string `json:"net_url" toml:"net_url"`
	PlgrAddress          string `json:"plgr_address" toml:"plgr_address"`
	PledgePoolToken      string `json:"pledge_pool_token" json:"pledge_pool_token"`
	BscPledgeOracleToken string `json:"bsc_pledge_oracle_token" json:"bsc_pledge_oracle_token"`
}

type MysqlConf struct {
	Address      string `json:"address" toml:"address"`
	Port         string `json:"port" toml:"port"`
	DbName       string `json:"db_name" toml:"db_name"`
	UserName     string `json:"user_name" toml:"user_name"`
	Password     string `json:"password" toml:"password"`
	MaxOpenConns int    `json:"max_open_conns" toml:"max_open_conns"`
	MaxIdleConns int    `json:"max_idle_conns" toml:"max_idle_conns"`
	MaxLifeTime  int    `json:"max_life_time" toml:"max_life_time"`
}

type RedisConf struct {
	Address     string `json:"address" toml:"address"`
	Port        string `json:"port" toml:"port"`
	Db          string `json:"db" toml:"db"`
	Password    string `json:"password" toml:"password"`
	MaxIdle     int    `json:"max_idle" toml:"max_idle"`
	MaxActive   int    `json:"max_active" toml:"max_active"`
	IdleTimeout int    `json:"idle_timeout" toml:"idle_timeout"`
}
