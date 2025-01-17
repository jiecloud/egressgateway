// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package common

import (
	. "github.com/onsi/gomega"
	"github.com/spidernet-io/e2eframework/framework"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/spidernet-io/egressgateway/pkg/utils"
	"github.com/spidernet-io/egressgateway/test/e2e/tools"
)

func GetAllNodes(f *framework.Framework) (nodes []string, err error) {
	nodelist, err := f.GetNodeList()
	if err != nil {
		return nil, err
	}
	for _, node := range nodelist.Items {
		nodes = append(nodes, node.Name)
	}
	return nodes, nil
}

func GetNodesByMatchLabels(f *framework.Framework, matchLabels map[string]string) (nodes []string, err error) {
	nodeList := &corev1.NodeList{}
	err = f.ListResource(nodeList, client.MatchingLabels(matchLabels))
	if err != nil {
		return nil, err
	}
	for _, item := range nodeList.Items {
		nodes = append(nodes, item.Name)
	}
	return
}

func GetUnmatchedNodes(f *framework.Framework, matchedNodes []string) (nodes []string, err error) {
	nodes, err = GetAllNodes(f)
	if err != nil {
		return nil, err
	}
	nodes = tools.SubtractionSlice(nodes, matchedNodes)
	return nodes, nil
}

func LabelNodes(f *framework.Framework, nodes []string, labels map[string]string) error {
	for _, nodeName := range nodes {
		node, err := f.GetNode(nodeName)
		if err != nil {
			return err
		}
		for k, v := range labels {
			node.Labels[k] = v
		}
		node.SetLabels(node.Labels)
		err = f.UpdateResource(node)
		if err != nil {
			return err
		}
	}
	return nil
}

func UnLabelNodes(f *framework.Framework, nodes []string, labels map[string]string) error {
	for _, nodeName := range nodes {
		node, err := f.GetNode(nodeName)
		if err != nil {
			return err
		}
		nodeLabels := node.Labels
		for k := range labels {
			delete(nodeLabels, k)
		}
		node.SetLabels(nodeLabels)
		err = f.UpdateResource(node)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetAllNodesIP(f *framework.Framework) (nodesIPv4, nodesIPv6 []string) {
	nodesIPv4, nodesIPv6 = make([]string, 0), make([]string, 0)
	allNodes := f.Info.KindNodeList
	for _, node := range allNodes {
		getNode, err := f.GetNode(node)
		Expect(err).NotTo(HaveOccurred())
		ipv4, ipv6 := utils.GetNodeIP(getNode)
		if len(ipv4) != 0 {
			nodesIPv4 = append(nodesIPv4, ipv4)

		}
		if len(ipv6) != 0 {
			nodesIPv6 = append(nodesIPv6, ipv6)
		}
	}
	return
}
