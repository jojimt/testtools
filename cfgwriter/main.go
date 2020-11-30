package main

import (
	"context"
	"flag"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	var ns string
	var name string

	flag.StringVar(&ns, "name-space", "default", "Name space for secret")
	flag.StringVar(&name, "secret-name", "config", "Name of secret")
	flag.Parse()
	cfg, err := rest.InClusterConfig()
	if err != nil {
		panic(err)
	}

	cfgStr := []byte(cfg.GoString())
	data := map[string][]byte{
			"config": cfgStr,
		}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		panic(err)
	}

	secret := &v1.Secret{
		Data: data,
	}

	secret.ObjectMeta.Name = name
	_, err = kubeClient.CoreV1().Secrets(ns).Create(context.TODO(), secret, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
}
