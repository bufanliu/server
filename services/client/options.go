package client

import (
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

func NewFlags() []cli.Flag {
	return []cli.Flag{
		altsrc.NewBoolFlag(&cli.BoolFlag{Name: "debug", Usage: "debug mode"}),
		altsrc.NewStringFlag(&cli.StringFlag{Name: "log_level", Usage: "log level"}),
		altsrc.NewBoolFlag(&cli.BoolFlag{Name: "open_gin", Usage: "open gin server"}),
		altsrc.NewIntFlag(&cli.IntFlag{Name: "client_id", Usage: "client unique id"}),
		altsrc.NewStringFlag(&cli.StringFlag{Name: "client_name", Usage: "client name"}),
		altsrc.NewDurationFlag(&cli.DurationFlag{Name: "heart_beat", Usage: "heart beat seconds"}),
		altsrc.NewBoolFlag(&cli.BoolFlag{Name: "prompt_ui", Usage: "enable prompt ui"}),
		// cert
		altsrc.NewStringFlag(&cli.StringFlag{Name: "cert_path_debug", Usage: "debug tls cert_pem path"}),
		altsrc.NewStringFlag(&cli.StringFlag{Name: "key_path_debug", Usage: "debug tls server_key path"}),
		altsrc.NewStringFlag(&cli.StringFlag{Name: "cert_path_release", Usage: "release tls cert_pem path"}),
		altsrc.NewStringFlag(&cli.StringFlag{Name: "key_path_release", Usage: "release tls server_key path"}),
		altsrc.NewStringFlag(&cli.StringFlag{Name: "http_listen_addr", Usage: "http listen address"}),
		altsrc.NewStringFlag(&cli.StringFlag{Name: "gate_addr", Usage: "gate address"}),

		altsrc.NewStringFlag(&cli.StringFlag{Name: "config_file", Usage: "client config path"}),
	}
}

func NewClientBotsFlags() []cli.Flag {
	return []cli.Flag{
		altsrc.NewBoolFlag(&cli.BoolFlag{Name: "debug", Usage: "debug mode"}),
		altsrc.NewStringFlag(&cli.StringFlag{Name: "log_level", Usage: "log level"}),
		altsrc.NewDurationFlag(&cli.DurationFlag{Name: "heart_beat", Usage: "heart beat seconds"}),
		altsrc.NewBoolFlag(&cli.BoolFlag{Name: "prompt_ui", Usage: "enable prompt ui"}),
		altsrc.NewIntFlag(&cli.IntFlag{Name: "client_bots_num", Usage: "client bots number"}),
		// cert
		altsrc.NewStringFlag(&cli.StringFlag{Name: "cert_path_debug", Usage: "debug tls cert_pem path"}),
		altsrc.NewStringFlag(&cli.StringFlag{Name: "key_path_debug", Usage: "debug tls server_key path"}),
		altsrc.NewStringFlag(&cli.StringFlag{Name: "cert_path_release", Usage: "release tls cert_pem path"}),
		altsrc.NewStringFlag(&cli.StringFlag{Name: "key_path_release", Usage: "release tls server_key path"}),
		altsrc.NewStringFlag(&cli.StringFlag{Name: "http_listen_addr", Usage: "http listen address"}),
		altsrc.NewStringFlag(&cli.StringFlag{Name: "gate_addr", Usage: "gate address"}),
		&cli.StringFlag{
			Name:  "config_file",
			Value: "config/client_bots/config.toml",
		},
	}
}
