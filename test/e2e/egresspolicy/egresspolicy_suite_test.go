// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package egresspolicy_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

func TestEgresspolicy(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Egresspolicy Suite")
}

//var (
//	frame *framework.Framework
//	err   error
//	c     client.WithWatch
//
//	testV4, testV6                bool
//	controlPlane, worker, worker2 string
//
//	delOpts client.DeleteOption
//)
//
//var _ = BeforeSuite(func() {
//	GinkgoRecover()
//
//	delOpts = client.GracePeriodSeconds(0)
//
//	frame, err = framework.NewFramework(GinkgoT(), []func(scheme *runtime.Scheme) error{egressgatewayv1.AddToScheme})
//	Expect(err).NotTo(HaveOccurred(), "failed to NewFramework, details: %w", err)
//	c = frame.KClient
//
//	// get ip version of cluster
//	v4Enabled, v6Enabled, err := common.GetIPVersion(frame)
//	Expect(err).NotTo(HaveOccurred())
//	GinkgoWriter.Printf("v4Enabled: %v, v6Enabled: %v\n", v4Enabled, v6Enabled)
//	if v4Enabled {
//		testV4 = true
//	}
//	if v6Enabled && !v4Enabled {
//		testV6 = true
//	}
//	GinkgoWriter.Printf("testV4: %v, testV6: %v\n", testV4, testV6)
//
//	// get all nodes
//	nodes := frame.Info.KindNodeList
//	GinkgoWriter.Printf("nodes: %v\n", nodes)
//
//	for _, node := range nodes {
//		GinkgoWriter.Printf("node: %v\n", node)
//
//		switch {
//		case strings.HasSuffix(node, "control-plane"):
//			controlPlane = node
//		case strings.HasSuffix(node, "worker"):
//			worker = node
//		case strings.HasSuffix(node, "worker2"):
//			worker2 = node
//		default:
//		}
//	}
//})
