package configx

import (
	"testing"
	"time"
)

func TestHttpConfig(t *testing.T) {
	cfg, err := LoadConfigHttp("https://raw.githubusercontent.com/hagbei/remote-config/master/market.ini",
		"https://raw.githubusercontent.com/hagbei/remote-config/master/traffic.ini")
	if err != nil {
		t.Fatal(err)
	}

	v, err := cfg.Bool("Weigh", "switch")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(v)

	s, err := cfg.GetValue("MySQL", "ip")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(s)

	time.Sleep(2 * time.Minute)

	err = cfg.Reload()
	if err != nil {
		t.Fatal(err)
	}

	v, err = cfg.Bool("Weigh", "switch")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(v)

	s, err = cfg.GetValue("MySQL", "ip")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(s)
}
