package compreconciler

import (
	"bufio"
	"bytes"
	b64 "encoding/base64"
	"github.com/kyma-incubator/reconciler/pkg/compreconciler/kubeclient"
	"github.com/kyma-incubator/reconciler/pkg/compreconciler/types"
	"github.com/pkg/errors"
	"io"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilyaml "k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes"
	yamlToJson "sigs.k8s.io/yaml"
)

type goClient struct {
	kubeClient kubeclient.KubeClient
}

func newGoClient(kubeconfig string) (kubernetesClient, error) {

	base64kubeConfig := b64.StdEncoding.EncodeToString([]byte(kubeconfig))
	client, err := kubeclient.NewKubeClient(base64kubeConfig)
	if err != nil {
		return nil, err
	}

	return &goClient{
		kubeClient: *client,
	}, nil
}

func (g goClient) Deploy(manifest string) (results []string, resources []types.Metadata, err error) {
	chanMes, chanErr := asyncReadYaml([]byte(manifest))
	for {
		select {
		case dataBytes, ok := <-chanMes:
			{
				if !ok {
					return results, resources, err
				}
				if err != nil {
					results = append(results, err.Error())
					continue
				}

				json, err := yamlToJson.YAMLToJSON(dataBytes)
				if err != nil {
					continue
				}
				toUnstructured, err := kubeclient.ToUnstructured(json)
				if err == nil {
					resource, err := g.kubeClient.Apply(&toUnstructured)
					if err == nil {
						resources = append(resources, resource)
						continue
					}
					results = append(results, err.Error())
				}
			}
		case err, ok := <-chanErr:
			if !ok {
				return results, resources, err
			}
			if err == nil {
				continue
			}
			results = append(results, err.Error())
		}
	}
}

func (g goClient) Clientset() (*kubernetes.Clientset, error) {
	return g.kubeClient.GetClientSet()
}

func (g goClient) Delete(manifest string) (results []string, err error) {
	yamls, err := syncReadYaml([]byte(manifest))
	if err != nil {
		results = append(results, err.Error())
		return results, err
	}

	//delete resource in reverse order
	for i := len(yamls) - 1; i >= 0; i-- {
		json, err := yamlToJson.YAMLToJSON(yamls[i])
		if err != nil {
			continue
		}
		toUnstructured, err := kubeclient.ToUnstructured(json)
		if err != nil {
			results = append(results, err.Error())
			return results, err
		}
		err = g.kubeClient.DeleteResourceByKindAndNameAndNamespace(toUnstructured.GetKind(), toUnstructured.GetName(), toUnstructured.GetNamespace(), v1.DeleteOptions{})
		if err != nil {
			results = append(results, err.Error())
			return results, err
		}
	}
	return results, nil
}

func asyncReadYaml(data []byte) (<-chan []byte, <-chan error) {
	var (
		chanErr        = make(chan error)
		chanBytes      = make(chan []byte)
		multidocReader = utilyaml.NewYAMLReader(bufio.NewReader(bytes.NewReader(data)))
	)

	go func() {
		defer close(chanErr)
		defer close(chanBytes)

		for {
			buf, err := multidocReader.Read()
			if err != nil {
				if err == io.EOF {
					return
				}
				chanErr <- errors.Wrap(err, "failed to read yaml data")
				return
			}
			chanBytes <- buf
		}
	}()
	return chanBytes, chanErr
}

func syncReadYaml(data []byte) (results [][]byte, err error) {
	var (
		multidocReader = utilyaml.NewYAMLReader(bufio.NewReader(bytes.NewReader(data)))
	)

	for {
		buf, err := multidocReader.Read()
		if err != nil {
			if err == io.EOF {
				return results, nil
			}
			return results, err
		}
		results = append(results, buf)
	}
	return results, err
}
