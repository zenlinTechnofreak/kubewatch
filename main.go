package main

import (
	"fmt"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/client/restclient"
	"k8s.io/kubernetes/pkg/client/unversioned"
	"k8s.io/kubernetes/pkg/labels"
)

/*const (
	RUNNING = "running"
	ERROR   = "error"
	DELETED = "deleted"
	PENDING = "pending"
)*/

func main() {
	clientConfig := restclient.Config{}
	clientConfig.Host = "127.0.0.1:8080"
	client, err := unversioned.New(&clientConfig)
	if err != nil {
		fmt.Errorf("New unversioned client err: %v!\n", err.Error())
	}

	//opts := api.ListOptions{FieldSelector: fields.Set{"kind": "pod"}.AsSelector()}
	opts := api.ListOptions{LabelSelector: labels.Set{"app": "nginx"}.AsSelector()}

	w, err := client.Pods("").Watch(opts)
	if err != nil {
		fmt.Errorf("Get watch interface err")
	}
	for {
		event, ok := <-w.ResultChan()

		if !ok {
			fmt.Errorf("Watch err\n")
			// Resource was deleted, and chanle closed, so return to main programme
			//	return
		}

		fmt.Println(event.Type)
		if event.Type == "DELETED" {
			w.Stop()
			return
		}
	}
}
