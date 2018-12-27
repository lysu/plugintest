package main

import (
	"context"
	"fmt"

	"github.com/pingcap/tidb/plugin"
	"github.com/pingcap/tidb/sessionctx/variable"
)

// Validate implements TiDB plugin's Validate SPI.
func Validate(ctx context.Context, m *plugin.Manifest) error {
	fmt.Println("plugin_conn_hello validating")
	return nil
}

// OnInit implements TiDB plugin's OnInit SPI.
func OnInit(ctx context.Context, manifest *plugin.Manifest) error {
	fmt.Println("plugin_conn_hello init")
	return nil
}

// OnShutdown implements TiDB plugin's OnShutdown SPI.
func OnShutdown(ctx context.Context, manifest *plugin.Manifest) error {
	fmt.Println("plugin_conn_hello shutdown")
	return nil
}

// NotifyEvent implements TiDB Audit plugin's NotifyEvent SPI.
func NotifyEvent(ctx context.Context) error {
	fmt.Printf("plugin_conn_hello receive OnConn event from %s and say %s\n",
		ctx.Value("ip"), variable.GetSysVar("conn_hello_message").Value,
	)
	return nil
}
