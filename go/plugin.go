package go_flutter_clash

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/Dreamacro/clash/tunnel/statistic"
	"os"
	"path/filepath"
	"time"

	C "github.com/Dreamacro/clash/constant"
	"github.com/Dreamacro/clash/hub/executor"
	"github.com/Dreamacro/clash/hub/route"
	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
	"github.com/mapleafgo/go-flutter-clash/go/config"
)

const channelName = "go_flutter_clash"

// GoFlutterClashPlugin implements flutter.Plugin and handles method.
type GoFlutterClashPlugin struct {
	channel *plugin.MethodChannel
	status  bool
}

var _ flutter.Plugin = new(GoFlutterClashPlugin) // compile-time type check

// InitPlugin initializes the plugin.
func (p *GoFlutterClashPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	p.channel = plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	p.channel.HandleFunc("init", p.initClash)
	p.channel.HandleFunc("start", p.start)
	p.channel.HandleFunc("status", p.getStatus)
	return nil
}

func (p *GoFlutterClashPlugin) initClash(arguments any) (reply any, err error) {
	if homeDir, ok := arguments.(string); ok {
		if !filepath.IsAbs(homeDir) {
			currentDir, _ := os.Getwd()
			homeDir = filepath.Join(currentDir, homeDir)
		}
		C.SetHomeDir(homeDir)
		return nil, nil
	}
	return nil, errors.New("arguments error")
}

func (p *GoFlutterClashPlugin) start(arguments any) (reply any, err error) {
	if params, ok := arguments.([]any); ok {
		var profile, fcc string
		if params[0] != nil {
			profile = params[0].(string)
		}
		if params[1] != nil {
			fcc = params[1].(string)
		}
		cfg, err := config.Parse(profile, fcc)
		if err != nil {
			return nil, err
		}
		go route.Start("127.0.0.1:9090", cfg.General.Secret)
		executor.ApplyConfig(cfg, true)
		go p.trafficHandler()
		p.status = true
		return nil, nil
	}
	return nil, errors.New("props error")
}

func (p *GoFlutterClashPlugin) getStatus(any) (reply any, err error) {
	return p.status, nil
}

// 实时网速
func (p *GoFlutterClashPlugin) trafficHandler() {
	tick := time.NewTicker(time.Second)
	defer tick.Stop()
	t := statistic.DefaultManager
	buf := &bytes.Buffer{}
	for range tick.C {
		buf.Reset()
		up, down := t.Now()
		if err := json.NewEncoder(buf).Encode(route.Traffic{
			Up:   up,
			Down: down,
		}); err != nil {
			break
		}
		_ = p.channel.InvokeMethod("trafficHandler", string(buf.Bytes()))
	}
}
