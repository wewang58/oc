// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeSchedulerConfiguration) DeepCopyInto(out *KubeSchedulerConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.AlgorithmSource.DeepCopyInto(&out.AlgorithmSource)
	in.LeaderElection.DeepCopyInto(&out.LeaderElection)
	out.ClientConnection = in.ClientConnection
	out.DebuggingConfiguration = in.DebuggingConfiguration
	if in.BindTimeoutSeconds != nil {
		in, out := &in.BindTimeoutSeconds, &out.BindTimeoutSeconds
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeSchedulerConfiguration.
func (in *KubeSchedulerConfiguration) DeepCopy() *KubeSchedulerConfiguration {
	if in == nil {
		return nil
	}
	out := new(KubeSchedulerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeSchedulerConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeSchedulerLeaderElectionConfiguration) DeepCopyInto(out *KubeSchedulerLeaderElectionConfiguration) {
	*out = *in
	in.LeaderElectionConfiguration.DeepCopyInto(&out.LeaderElectionConfiguration)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeSchedulerLeaderElectionConfiguration.
func (in *KubeSchedulerLeaderElectionConfiguration) DeepCopy() *KubeSchedulerLeaderElectionConfiguration {
	if in == nil {
		return nil
	}
	out := new(KubeSchedulerLeaderElectionConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulerAlgorithmSource) DeepCopyInto(out *SchedulerAlgorithmSource) {
	*out = *in
	if in.Policy != nil {
		in, out := &in.Policy, &out.Policy
		*out = new(SchedulerPolicySource)
		(*in).DeepCopyInto(*out)
	}
	if in.Provider != nil {
		in, out := &in.Provider, &out.Provider
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulerAlgorithmSource.
func (in *SchedulerAlgorithmSource) DeepCopy() *SchedulerAlgorithmSource {
	if in == nil {
		return nil
	}
	out := new(SchedulerAlgorithmSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulerPolicyConfigMapSource) DeepCopyInto(out *SchedulerPolicyConfigMapSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulerPolicyConfigMapSource.
func (in *SchedulerPolicyConfigMapSource) DeepCopy() *SchedulerPolicyConfigMapSource {
	if in == nil {
		return nil
	}
	out := new(SchedulerPolicyConfigMapSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulerPolicyFileSource) DeepCopyInto(out *SchedulerPolicyFileSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulerPolicyFileSource.
func (in *SchedulerPolicyFileSource) DeepCopy() *SchedulerPolicyFileSource {
	if in == nil {
		return nil
	}
	out := new(SchedulerPolicyFileSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulerPolicySource) DeepCopyInto(out *SchedulerPolicySource) {
	*out = *in
	if in.File != nil {
		in, out := &in.File, &out.File
		*out = new(SchedulerPolicyFileSource)
		**out = **in
	}
	if in.ConfigMap != nil {
		in, out := &in.ConfigMap, &out.ConfigMap
		*out = new(SchedulerPolicyConfigMapSource)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulerPolicySource.
func (in *SchedulerPolicySource) DeepCopy() *SchedulerPolicySource {
	if in == nil {
		return nil
	}
	out := new(SchedulerPolicySource)
	in.DeepCopyInto(out)
	return out
}
