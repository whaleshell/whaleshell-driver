package sidecar

import (
	"slices"
	"strings"
	"testing"

	"github.com/whaleshell/whaleshell-core/defaults"
)

func TestCABundleEnvIncludesGit(t *testing.T) {
	got := CABundleEnv("")
	wantKeys := []string{
		"SSL_CERT_FILE",
		"CURL_CA_BUNDLE",
		"REQUESTS_CA_BUNDLE",
		"NODE_EXTRA_CA_CERTS",
		"GIT_SSL_CAINFO",
	}
	var keys []string
	for _, e := range got {
		k, v, ok := strings.Cut(e, "=")
		if !ok {
			t.Fatalf("bad entry %q", e)
		}
		keys = append(keys, k)
		if v != defaults.GuestCAFile {
			t.Fatalf("%s=%q want %q", k, v, defaults.GuestCAFile)
		}
	}
	for _, k := range wantKeys {
		if !slices.Contains(keys, k) {
			t.Fatalf("missing %s in %v", k, keys)
		}
	}
}
