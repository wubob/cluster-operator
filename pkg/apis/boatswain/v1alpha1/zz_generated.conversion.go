// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1alpha1

import (
	boatswain "github.com/staebler/boatswain/pkg/apis/boatswain"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	unsafe "unsafe"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_Host_To_boatswain_Host,
		Convert_boatswain_Host_To_v1alpha1_Host,
		Convert_v1alpha1_HostList_To_boatswain_HostList,
		Convert_boatswain_HostList_To_v1alpha1_HostList,
		Convert_v1alpha1_HostSpec_To_boatswain_HostSpec,
		Convert_boatswain_HostSpec_To_v1alpha1_HostSpec,
		Convert_v1alpha1_HostStatus_To_boatswain_HostStatus,
		Convert_boatswain_HostStatus_To_v1alpha1_HostStatus,
	)
}

func autoConvert_v1alpha1_Host_To_boatswain_Host(in *Host, out *boatswain.Host, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_HostSpec_To_boatswain_HostSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_HostStatus_To_boatswain_HostStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_Host_To_boatswain_Host is an autogenerated conversion function.
func Convert_v1alpha1_Host_To_boatswain_Host(in *Host, out *boatswain.Host, s conversion.Scope) error {
	return autoConvert_v1alpha1_Host_To_boatswain_Host(in, out, s)
}

func autoConvert_boatswain_Host_To_v1alpha1_Host(in *boatswain.Host, out *Host, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_boatswain_HostSpec_To_v1alpha1_HostSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_boatswain_HostStatus_To_v1alpha1_HostStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_boatswain_Host_To_v1alpha1_Host is an autogenerated conversion function.
func Convert_boatswain_Host_To_v1alpha1_Host(in *boatswain.Host, out *Host, s conversion.Scope) error {
	return autoConvert_boatswain_Host_To_v1alpha1_Host(in, out, s)
}

func autoConvert_v1alpha1_HostList_To_boatswain_HostList(in *HostList, out *boatswain.HostList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]boatswain.Host)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_HostList_To_boatswain_HostList is an autogenerated conversion function.
func Convert_v1alpha1_HostList_To_boatswain_HostList(in *HostList, out *boatswain.HostList, s conversion.Scope) error {
	return autoConvert_v1alpha1_HostList_To_boatswain_HostList(in, out, s)
}

func autoConvert_boatswain_HostList_To_v1alpha1_HostList(in *boatswain.HostList, out *HostList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Host)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_boatswain_HostList_To_v1alpha1_HostList is an autogenerated conversion function.
func Convert_boatswain_HostList_To_v1alpha1_HostList(in *boatswain.HostList, out *HostList, s conversion.Scope) error {
	return autoConvert_boatswain_HostList_To_v1alpha1_HostList(in, out, s)
}

func autoConvert_v1alpha1_HostSpec_To_boatswain_HostSpec(in *HostSpec, out *boatswain.HostSpec, s conversion.Scope) error {
	return nil
}

// Convert_v1alpha1_HostSpec_To_boatswain_HostSpec is an autogenerated conversion function.
func Convert_v1alpha1_HostSpec_To_boatswain_HostSpec(in *HostSpec, out *boatswain.HostSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_HostSpec_To_boatswain_HostSpec(in, out, s)
}

func autoConvert_boatswain_HostSpec_To_v1alpha1_HostSpec(in *boatswain.HostSpec, out *HostSpec, s conversion.Scope) error {
	return nil
}

// Convert_boatswain_HostSpec_To_v1alpha1_HostSpec is an autogenerated conversion function.
func Convert_boatswain_HostSpec_To_v1alpha1_HostSpec(in *boatswain.HostSpec, out *HostSpec, s conversion.Scope) error {
	return autoConvert_boatswain_HostSpec_To_v1alpha1_HostSpec(in, out, s)
}

func autoConvert_v1alpha1_HostStatus_To_boatswain_HostStatus(in *HostStatus, out *boatswain.HostStatus, s conversion.Scope) error {
	return nil
}

// Convert_v1alpha1_HostStatus_To_boatswain_HostStatus is an autogenerated conversion function.
func Convert_v1alpha1_HostStatus_To_boatswain_HostStatus(in *HostStatus, out *boatswain.HostStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_HostStatus_To_boatswain_HostStatus(in, out, s)
}

func autoConvert_boatswain_HostStatus_To_v1alpha1_HostStatus(in *boatswain.HostStatus, out *HostStatus, s conversion.Scope) error {
	return nil
}

// Convert_boatswain_HostStatus_To_v1alpha1_HostStatus is an autogenerated conversion function.
func Convert_boatswain_HostStatus_To_v1alpha1_HostStatus(in *boatswain.HostStatus, out *HostStatus, s conversion.Scope) error {
	return autoConvert_boatswain_HostStatus_To_v1alpha1_HostStatus(in, out, s)
}
