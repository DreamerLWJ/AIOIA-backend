package config

type RedisConfig struct {
	Addr               string `toml:"addr"`                 // 连接地址
	Username           string `toml:"username"`             // 用户名
	Password           string `toml:"password"`             // 密码
	DB                 int    `toml:"db"`                   // 数据库 index
	MaxActive          int    `toml:"max_active"`           // 连接池最大连接数
	MinIdle            int    `toml:"min_idle"`             // 启动阶段创建的连接数，维持的连接数
	MaxIdle            int    `toml:"max_idle"`             // 最大空闲连接数
	MaxIdleTime        int    `toml:"max_idle_time"`        // 最大空闲时间
	IdleTimeout        int    `toml:"idle_timeout"`         // 闲置超时，默认5分钟，-1表示取消闲置超时检查
	ConnTimeout        int    `toml:"dial_timeout"`         // 连接建立超时时间，默认5秒。
	ReadTimeout        int    `toml:"read_timeout"`         // 读超时，默认3秒， -1表示取消读超时
	WriteTimeout       int    `toml:"write_timeout"`        // 写超时，默认等于读超时
	PoolTimeout        int    `toml:"pool_timeout"`         // 当所有连接都处在繁忙状态时，客户端等待可用连接的最大等待时长，默认为读超时+1秒。
	IdleCheckFrequency int    `toml:"idle_check_frequency"` // 闲置连接检查的周期，默认为1分钟，-1表示不做周期性检查，只在客户端获取连接时对闲置连接进行处理。
	MaxConnAge         int    `toml:"max_conn_age"`         // 连接存活时长，从创建开始计时，超过指定时长则关闭连接，默认为0，即不关闭存活时长较长的连接
	MaxRetries         int    `toml:"max_retries"`          // 命令执行失败时，最多重试多少次，默认为0即不重试
	MinRetryBackoff    int    `toml:"min_retry_back_off"`   // 每次计算重试间隔时间的下限，默认8毫秒，-1表示取消间隔
	MaxRetryBackoff    int    `toml:"max_retry_back_off"`   // 每次计算重试间隔时间的上限，默认512毫秒，-1表示取消间隔
}

func (r RedisConfig) String() string {
	return "redis"
}
