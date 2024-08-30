// Copyright (c) 2021 Nxpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package spicedb

import (
	"github.com/nxpkg/nxpod/installer/pkg/common"

	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func networkpolicy(ctx *common.RenderContext) ([]runtime.Object, error) {
	labels := common.DefaultLabels(Component)

	return []runtime.Object{
		&networkingv1.NetworkPolicy{
			TypeMeta: common.TypeMetaNetworkPolicy,
			ObjectMeta: metav1.ObjectMeta{
				Name:      Component,
				Namespace: ctx.Namespace,
				Labels:    labels,
			},
			Spec: networkingv1.NetworkPolicySpec{
				PodSelector: metav1.LabelSelector{MatchLabels: labels},
				PolicyTypes: []networkingv1.PolicyType{"Ingress"},
				Ingress: []networkingv1.NetworkPolicyIngressRule{
					{
						Ports: []networkingv1.NetworkPolicyPort{
							{
								Protocol: common.TCPProtocol,
								Port:     &intstr.IntOrString{IntVal: ContainerDispatchPort},
							},
						},
						From: []networkingv1.NetworkPolicyPeer{
							{
								PodSelector: &metav1.LabelSelector{
									MatchLabels: map[string]string{
										"component": Component,
									},
								},
							},
						},
					},
					{
						Ports: []networkingv1.NetworkPolicyPort{
							{
								Protocol: common.TCPProtocol,
								Port:     &intstr.IntOrString{IntVal: ContainerHTTPPort},
							},
							{
								Protocol: common.TCPProtocol,
								Port:     &intstr.IntOrString{IntVal: ContainerGRPCPort},
							},
						},
						From: []networkingv1.NetworkPolicyPeer{
							{
								PodSelector: &metav1.LabelSelector{
									MatchLabels: map[string]string{
										"component": common.PublicApiComponent,
									},
								},
							},
						},
					},
					{
						Ports: []networkingv1.NetworkPolicyPort{
							{
								Protocol: common.TCPProtocol,
								Port:     &intstr.IntOrString{IntVal: ContainerHTTPPort},
							},
							{
								Protocol: common.TCPProtocol,
								Port:     &intstr.IntOrString{IntVal: ContainerGRPCPort},
							},
						},
						From: []networkingv1.NetworkPolicyPeer{
							{
								PodSelector: &metav1.LabelSelector{
									MatchLabels: map[string]string{
										"component": common.ServerComponent,
									},
								},
							},
						},
					},
				},
			},
		},
	}, nil
}
