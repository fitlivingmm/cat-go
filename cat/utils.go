package cat

import (
	"fmt"
	"net"
	"runtime"
	"strings"
	"time"
)

func getLocalhostIp() (ip net.IP, err error) {
	ip = net.IPv4(127, 0, 0, 1)

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP
				return
			}
		}
	}
	return
}

func ip2String(ip net.IP) string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[12], ip[13], ip[14], ip[15])
}

func ip2HexString(ip net.IP) string {
	return fmt.Sprintf("%02x%02x%02x%02x", ip[12], ip[13], ip[14], ip[15])
}

func duration2Millis(duration time.Duration) int64 {
	return duration.Nanoseconds() / time.Millisecond.Nanoseconds()
}

// SafeGo 错误处理go
func SafeGo(fn func()) {
	go func() {
		defer SimpleRecover()
		fn()
	}()
}

func SimpleRecover() {
	if err := recover(); err != nil {
		stacks := DumpStacks(1, 10)
		logger.Error("%s\n%s", err, stacks)
		return
	}
	return
}

func DumpStacks(skip, max int) string {
	var stacks []string
	for i := skip; i <= max; i += 1 {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		stacks = append(stacks, fmt.Sprintf("\t%s:%d", file, line))
	}
	return strings.Join(stacks, "\n")
}
