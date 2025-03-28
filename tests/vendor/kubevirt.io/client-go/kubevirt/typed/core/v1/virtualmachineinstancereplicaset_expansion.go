/*
 * This file is part of the KubeVirt project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright The KubeVirt Authors
 *
 */

package v1

import (
	"context"

	autov1 "k8s.io/api/autoscaling/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	v1 "kubevirt.io/api/core/v1"
)

type VirtualMachineInstanceReplicaSetExpansion interface {
	GetScale(ctx context.Context, replicaSetName string, options metav1.GetOptions) (*autov1.Scale, error)
	UpdateScale(ctx context.Context, replicaSetName string, scale *autov1.Scale) (*autov1.Scale, error)
	PatchStatus(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions) (*v1.VirtualMachineInstanceReplicaSet, error)
}

func (c *virtualMachineInstanceReplicaSets) GetScale(ctx context.Context, replicaSetName string, options metav1.GetOptions) (*autov1.Scale, error) {
	result := &autov1.Scale{}
	err := c.GetClient().Get().
		Namespace(c.GetNamespace()).
		Resource("virtualmachineinstancereplicasets").
		Name(replicaSetName).
		SubResource("scale").
		Do(ctx).
		Into(result)
	return result, err
}

func (c *virtualMachineInstanceReplicaSets) UpdateScale(ctx context.Context, replicaSetName string, scale *autov1.Scale) (*autov1.Scale, error) {
	result := &autov1.Scale{}
	err := c.GetClient().Put().
		Namespace(c.GetNamespace()).
		Resource("virtualmachineinstancereplicasets").
		Name(replicaSetName).
		SubResource("scale").
		Body(scale).
		Do(ctx).
		Into(result)
	return result, err
}

func (c *virtualMachineInstanceReplicaSets) PatchStatus(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions) (*v1.VirtualMachineInstanceReplicaSet, error) {
	return c.Patch(ctx, name, pt, data, opts, "status")
}
