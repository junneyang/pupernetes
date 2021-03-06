// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2018 Datadog, Inc.

package config

import (
	"github.com/spf13/viper"
	"time"
)

var ViperConfig = viper.New()

const (
	// Job control
	JobTypeKey    = "job-type"
	JobSystemd    = "systemd"
	JobForeground = "fg"
)

func init() {
	ViperConfig.SetDefault("hyperkube-version", "1.10.3")
	ViperConfig.SetDefault("vault-version", "0.9.5")
	ViperConfig.SetDefault("etcd-version", "3.1.11")
	ViperConfig.SetDefault("cni-version", "0.7.0")

	ViperConfig.SetDefault("kubernetes-cluster-ip-range", "192.168.254.0/24")
	ViperConfig.SetDefault("bind-address", "127.0.0.1:8989")
	ViperConfig.SetDefault("api-address", "127.0.0.1:8989")
	ViperConfig.SetDefault("kubelet-root-dir", "/var/lib/p8s-kubelet")
	ViperConfig.SetDefault("systemd-unit-prefix", "p8s-")

	ViperConfig.SetDefault("kubectl-link", "")
	ViperConfig.SetDefault("vault-root-token", "")

	ViperConfig.SetDefault("clean", "etcd,kubelet,mounts,iptables")
	ViperConfig.SetDefault("drain", "all")
	ViperConfig.SetDefault("timeout", time.Hour*6)
	ViperConfig.SetDefault("gc", time.Second*60)

	// The supported job-type are "fg" and "systemd"
	ViperConfig.SetDefault(JobTypeKey, JobForeground)

	ViperConfig.SetDefault("systemd-job-name", "pupernetes")

	ViperConfig.SetDefault("apply", false)

	ViperConfig.SetDefault("logging-since", time.Minute*5)
	ViperConfig.SetDefault("unit-to-watch", "pupernetes.service")
}
