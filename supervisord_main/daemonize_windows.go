//go:build windows
// +build windows

package supervisord_main

func Daemonize(logfile string, proc func()) {
	proc()
}
