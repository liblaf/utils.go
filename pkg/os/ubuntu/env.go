package ubuntu

import "path/filepath"

func EtcApt() string { return "/etc/apt" }

func AptSourcesListD() string {
	return filepath.Join(EtcApt(), "sources.list.d")
}

func AptTrustedGpgD() string {
	return filepath.Join(EtcApt(), "trusted.gpg.d")
}
