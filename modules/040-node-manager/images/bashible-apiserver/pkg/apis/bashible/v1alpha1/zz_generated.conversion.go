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

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	bashible "d8.io/bashible/pkg/apis/bashible"

	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*Bashible)(nil), (*bashible.Bashible)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Bashible_To_bashible_Bashible(a.(*Bashible), b.(*bashible.Bashible), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*bashible.Bashible)(nil), (*Bashible)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_bashible_Bashible_To_v1alpha1_Bashible(a.(*bashible.Bashible), b.(*Bashible), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*BashibleList)(nil), (*bashible.BashibleList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_BashibleList_To_bashible_BashibleList(a.(*BashibleList), b.(*bashible.BashibleList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*bashible.BashibleList)(nil), (*BashibleList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_bashible_BashibleList_To_v1alpha1_BashibleList(a.(*bashible.BashibleList), b.(*BashibleList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*NodeGroupBundle)(nil), (*bashible.NodeGroupBundle)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_NodeGroupBundle_To_bashible_NodeGroupBundle(a.(*NodeGroupBundle), b.(*bashible.NodeGroupBundle), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*bashible.NodeGroupBundle)(nil), (*NodeGroupBundle)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_bashible_NodeGroupBundle_To_v1alpha1_NodeGroupBundle(a.(*bashible.NodeGroupBundle), b.(*NodeGroupBundle), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*NodeGroupBundleList)(nil), (*bashible.NodeGroupBundleList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_NodeGroupBundleList_To_bashible_NodeGroupBundleList(a.(*NodeGroupBundleList), b.(*bashible.NodeGroupBundleList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*bashible.NodeGroupBundleList)(nil), (*NodeGroupBundleList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_bashible_NodeGroupBundleList_To_v1alpha1_NodeGroupBundleList(a.(*bashible.NodeGroupBundleList), b.(*NodeGroupBundleList), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_Bashible_To_bashible_Bashible(in *Bashible, out *bashible.Bashible, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Data = *(*map[string]string)(unsafe.Pointer(&in.Data))
	return nil
}

// Convert_v1alpha1_Bashible_To_bashible_Bashible is an autogenerated conversion function.
func Convert_v1alpha1_Bashible_To_bashible_Bashible(in *Bashible, out *bashible.Bashible, s conversion.Scope) error {
	return autoConvert_v1alpha1_Bashible_To_bashible_Bashible(in, out, s)
}

func autoConvert_bashible_Bashible_To_v1alpha1_Bashible(in *bashible.Bashible, out *Bashible, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Data = *(*map[string]string)(unsafe.Pointer(&in.Data))
	return nil
}

// Convert_bashible_Bashible_To_v1alpha1_Bashible is an autogenerated conversion function.
func Convert_bashible_Bashible_To_v1alpha1_Bashible(in *bashible.Bashible, out *Bashible, s conversion.Scope) error {
	return autoConvert_bashible_Bashible_To_v1alpha1_Bashible(in, out, s)
}

func autoConvert_v1alpha1_BashibleList_To_bashible_BashibleList(in *BashibleList, out *bashible.BashibleList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]bashible.Bashible)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_BashibleList_To_bashible_BashibleList is an autogenerated conversion function.
func Convert_v1alpha1_BashibleList_To_bashible_BashibleList(in *BashibleList, out *bashible.BashibleList, s conversion.Scope) error {
	return autoConvert_v1alpha1_BashibleList_To_bashible_BashibleList(in, out, s)
}

func autoConvert_bashible_BashibleList_To_v1alpha1_BashibleList(in *bashible.BashibleList, out *BashibleList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Bashible)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_bashible_BashibleList_To_v1alpha1_BashibleList is an autogenerated conversion function.
func Convert_bashible_BashibleList_To_v1alpha1_BashibleList(in *bashible.BashibleList, out *BashibleList, s conversion.Scope) error {
	return autoConvert_bashible_BashibleList_To_v1alpha1_BashibleList(in, out, s)
}

func autoConvert_v1alpha1_NodeGroupBundle_To_bashible_NodeGroupBundle(in *NodeGroupBundle, out *bashible.NodeGroupBundle, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Data = *(*map[string]string)(unsafe.Pointer(&in.Data))
	return nil
}

// Convert_v1alpha1_NodeGroupBundle_To_bashible_NodeGroupBundle is an autogenerated conversion function.
func Convert_v1alpha1_NodeGroupBundle_To_bashible_NodeGroupBundle(in *NodeGroupBundle, out *bashible.NodeGroupBundle, s conversion.Scope) error {
	return autoConvert_v1alpha1_NodeGroupBundle_To_bashible_NodeGroupBundle(in, out, s)
}

func autoConvert_bashible_NodeGroupBundle_To_v1alpha1_NodeGroupBundle(in *bashible.NodeGroupBundle, out *NodeGroupBundle, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Data = *(*map[string]string)(unsafe.Pointer(&in.Data))
	return nil
}

// Convert_bashible_NodeGroupBundle_To_v1alpha1_NodeGroupBundle is an autogenerated conversion function.
func Convert_bashible_NodeGroupBundle_To_v1alpha1_NodeGroupBundle(in *bashible.NodeGroupBundle, out *NodeGroupBundle, s conversion.Scope) error {
	return autoConvert_bashible_NodeGroupBundle_To_v1alpha1_NodeGroupBundle(in, out, s)
}

func autoConvert_v1alpha1_NodeGroupBundleList_To_bashible_NodeGroupBundleList(in *NodeGroupBundleList, out *bashible.NodeGroupBundleList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]bashible.NodeGroupBundle)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_NodeGroupBundleList_To_bashible_NodeGroupBundleList is an autogenerated conversion function.
func Convert_v1alpha1_NodeGroupBundleList_To_bashible_NodeGroupBundleList(in *NodeGroupBundleList, out *bashible.NodeGroupBundleList, s conversion.Scope) error {
	return autoConvert_v1alpha1_NodeGroupBundleList_To_bashible_NodeGroupBundleList(in, out, s)
}

func autoConvert_bashible_NodeGroupBundleList_To_v1alpha1_NodeGroupBundleList(in *bashible.NodeGroupBundleList, out *NodeGroupBundleList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]NodeGroupBundle)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_bashible_NodeGroupBundleList_To_v1alpha1_NodeGroupBundleList is an autogenerated conversion function.
func Convert_bashible_NodeGroupBundleList_To_v1alpha1_NodeGroupBundleList(in *bashible.NodeGroupBundleList, out *NodeGroupBundleList, s conversion.Scope) error {
	return autoConvert_bashible_NodeGroupBundleList_To_v1alpha1_NodeGroupBundleList(in, out, s)
}
