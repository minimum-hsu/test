package g

import (
	"time"
)

// changelog:
// 3.1.3: code refactor
// 3.1.4: bugfix ignore configuration
// 5.0.0: 支持通过配置控制是否开启/run接口；收集udp流量数据；du某个目录的大小
// 5.1.0: 同步插件的时候不再使用checksum机制
// 5.1.3: Fix config syntax error when deploying
// 5.1.4: Only trustable ip could access the webpage
// 5.1.5: New policy and plugin mechanism
// 5.1.6: Update cfg.json in release package. Program file is same as 5.1.5.
// 5.1.7: Fix failure of plugin updating.
const (
	VERSION          = "5.1.7"
	COLLECT_INTERVAL = time.Second
	URL_CHECK_HEALTH = "url.check.health"
	NET_PORT_LISTEN  = "net.port.listen"
	DU_BS            = "du.bs"
	PROC_NUM         = "proc.num"
)
