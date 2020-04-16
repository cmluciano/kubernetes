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

package v1beta1

import (
	unsafe "unsafe"

	v1 "k8s.io/api/core/v1"
	v1beta1 "k8s.io/api/networking/v1beta1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/kubernetes/pkg/apis/core"
	corev1 "k8s.io/kubernetes/pkg/apis/core/v1"
	networking "k8s.io/kubernetes/pkg/apis/networking"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*v1beta1.HTTPIngressPath)(nil), (*networking.HTTPIngressPath)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_HTTPIngressPath_To_networking_HTTPIngressPath(a.(*v1beta1.HTTPIngressPath), b.(*networking.HTTPIngressPath), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*networking.HTTPIngressPath)(nil), (*v1beta1.HTTPIngressPath)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_networking_HTTPIngressPath_To_v1beta1_HTTPIngressPath(a.(*networking.HTTPIngressPath), b.(*v1beta1.HTTPIngressPath), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.HTTPIngressRuleValue)(nil), (*networking.HTTPIngressRuleValue)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_HTTPIngressRuleValue_To_networking_HTTPIngressRuleValue(a.(*v1beta1.HTTPIngressRuleValue), b.(*networking.HTTPIngressRuleValue), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*networking.HTTPIngressRuleValue)(nil), (*v1beta1.HTTPIngressRuleValue)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_networking_HTTPIngressRuleValue_To_v1beta1_HTTPIngressRuleValue(a.(*networking.HTTPIngressRuleValue), b.(*v1beta1.HTTPIngressRuleValue), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.Ingress)(nil), (*networking.Ingress)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_Ingress_To_networking_Ingress(a.(*v1beta1.Ingress), b.(*networking.Ingress), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*networking.Ingress)(nil), (*v1beta1.Ingress)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_networking_Ingress_To_v1beta1_Ingress(a.(*networking.Ingress), b.(*v1beta1.Ingress), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.IngressBackend)(nil), (*networking.IngressBackend)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_IngressBackend_To_networking_IngressBackend(a.(*v1beta1.IngressBackend), b.(*networking.IngressBackend), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*networking.IngressBackend)(nil), (*v1beta1.IngressBackend)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_networking_IngressBackend_To_v1beta1_IngressBackend(a.(*networking.IngressBackend), b.(*v1beta1.IngressBackend), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.IngressClass)(nil), (*networking.IngressClass)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_IngressClass_To_networking_IngressClass(a.(*v1beta1.IngressClass), b.(*networking.IngressClass), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*networking.IngressClass)(nil), (*v1beta1.IngressClass)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_networking_IngressClass_To_v1beta1_IngressClass(a.(*networking.IngressClass), b.(*v1beta1.IngressClass), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.IngressClassList)(nil), (*networking.IngressClassList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_IngressClassList_To_networking_IngressClassList(a.(*v1beta1.IngressClassList), b.(*networking.IngressClassList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*networking.IngressClassList)(nil), (*v1beta1.IngressClassList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_networking_IngressClassList_To_v1beta1_IngressClassList(a.(*networking.IngressClassList), b.(*v1beta1.IngressClassList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.IngressClassSpec)(nil), (*networking.IngressClassSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_IngressClassSpec_To_networking_IngressClassSpec(a.(*v1beta1.IngressClassSpec), b.(*networking.IngressClassSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*networking.IngressClassSpec)(nil), (*v1beta1.IngressClassSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_networking_IngressClassSpec_To_v1beta1_IngressClassSpec(a.(*networking.IngressClassSpec), b.(*v1beta1.IngressClassSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.IngressList)(nil), (*networking.IngressList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_IngressList_To_networking_IngressList(a.(*v1beta1.IngressList), b.(*networking.IngressList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*networking.IngressList)(nil), (*v1beta1.IngressList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_networking_IngressList_To_v1beta1_IngressList(a.(*networking.IngressList), b.(*v1beta1.IngressList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.IngressRule)(nil), (*networking.IngressRule)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_IngressRule_To_networking_IngressRule(a.(*v1beta1.IngressRule), b.(*networking.IngressRule), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*networking.IngressRule)(nil), (*v1beta1.IngressRule)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_networking_IngressRule_To_v1beta1_IngressRule(a.(*networking.IngressRule), b.(*v1beta1.IngressRule), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.IngressRuleValue)(nil), (*networking.IngressRuleValue)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_IngressRuleValue_To_networking_IngressRuleValue(a.(*v1beta1.IngressRuleValue), b.(*networking.IngressRuleValue), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*networking.IngressRuleValue)(nil), (*v1beta1.IngressRuleValue)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_networking_IngressRuleValue_To_v1beta1_IngressRuleValue(a.(*networking.IngressRuleValue), b.(*v1beta1.IngressRuleValue), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.IngressStatus)(nil), (*networking.IngressStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_IngressStatus_To_networking_IngressStatus(a.(*v1beta1.IngressStatus), b.(*networking.IngressStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*networking.IngressStatus)(nil), (*v1beta1.IngressStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_networking_IngressStatus_To_v1beta1_IngressStatus(a.(*networking.IngressStatus), b.(*v1beta1.IngressStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.IngressTLS)(nil), (*networking.IngressTLS)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_IngressTLS_To_networking_IngressTLS(a.(*v1beta1.IngressTLS), b.(*networking.IngressTLS), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*networking.IngressTLS)(nil), (*v1beta1.IngressTLS)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_networking_IngressTLS_To_v1beta1_IngressTLS(a.(*networking.IngressTLS), b.(*v1beta1.IngressTLS), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*networking.IngressSpec)(nil), (*v1beta1.IngressSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_networking_IngressSpec_To_v1beta1_IngressSpec(a.(*networking.IngressSpec), b.(*v1beta1.IngressSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1beta1.IngressSpec)(nil), (*networking.IngressSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_IngressSpec_To_networking_IngressSpec(a.(*v1beta1.IngressSpec), b.(*networking.IngressSpec), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1beta1_HTTPIngressPath_To_networking_HTTPIngressPath(in *v1beta1.HTTPIngressPath, out *networking.HTTPIngressPath, s conversion.Scope) error {
	out.Path = in.Path
	out.PathType = (*networking.PathType)(unsafe.Pointer(in.PathType))
	if err := Convert_v1beta1_IngressBackend_To_networking_IngressBackend(&in.Backend, &out.Backend, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_HTTPIngressPath_To_networking_HTTPIngressPath is an autogenerated conversion function.
func Convert_v1beta1_HTTPIngressPath_To_networking_HTTPIngressPath(in *v1beta1.HTTPIngressPath, out *networking.HTTPIngressPath, s conversion.Scope) error {
	return autoConvert_v1beta1_HTTPIngressPath_To_networking_HTTPIngressPath(in, out, s)
}

func autoConvert_networking_HTTPIngressPath_To_v1beta1_HTTPIngressPath(in *networking.HTTPIngressPath, out *v1beta1.HTTPIngressPath, s conversion.Scope) error {
	out.Path = in.Path
	out.PathType = (*v1beta1.PathType)(unsafe.Pointer(in.PathType))
	if err := Convert_networking_IngressBackend_To_v1beta1_IngressBackend(&in.Backend, &out.Backend, s); err != nil {
		return err
	}
	return nil
}

// Convert_networking_HTTPIngressPath_To_v1beta1_HTTPIngressPath is an autogenerated conversion function.
func Convert_networking_HTTPIngressPath_To_v1beta1_HTTPIngressPath(in *networking.HTTPIngressPath, out *v1beta1.HTTPIngressPath, s conversion.Scope) error {
	return autoConvert_networking_HTTPIngressPath_To_v1beta1_HTTPIngressPath(in, out, s)
}

func autoConvert_v1beta1_HTTPIngressRuleValue_To_networking_HTTPIngressRuleValue(in *v1beta1.HTTPIngressRuleValue, out *networking.HTTPIngressRuleValue, s conversion.Scope) error {
	if in.Paths != nil {
		in, out := &in.Paths, &out.Paths
		*out = make([]networking.HTTPIngressPath, len(*in))
		for i := range *in {
			if err := Convert_v1beta1_HTTPIngressPath_To_networking_HTTPIngressPath(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Paths = nil
	}
	return nil
}

// Convert_v1beta1_HTTPIngressRuleValue_To_networking_HTTPIngressRuleValue is an autogenerated conversion function.
func Convert_v1beta1_HTTPIngressRuleValue_To_networking_HTTPIngressRuleValue(in *v1beta1.HTTPIngressRuleValue, out *networking.HTTPIngressRuleValue, s conversion.Scope) error {
	return autoConvert_v1beta1_HTTPIngressRuleValue_To_networking_HTTPIngressRuleValue(in, out, s)
}

func autoConvert_networking_HTTPIngressRuleValue_To_v1beta1_HTTPIngressRuleValue(in *networking.HTTPIngressRuleValue, out *v1beta1.HTTPIngressRuleValue, s conversion.Scope) error {
	if in.Paths != nil {
		in, out := &in.Paths, &out.Paths
		*out = make([]v1beta1.HTTPIngressPath, len(*in))
		for i := range *in {
			if err := Convert_networking_HTTPIngressPath_To_v1beta1_HTTPIngressPath(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Paths = nil
	}
	return nil
}

// Convert_networking_HTTPIngressRuleValue_To_v1beta1_HTTPIngressRuleValue is an autogenerated conversion function.
func Convert_networking_HTTPIngressRuleValue_To_v1beta1_HTTPIngressRuleValue(in *networking.HTTPIngressRuleValue, out *v1beta1.HTTPIngressRuleValue, s conversion.Scope) error {
	return autoConvert_networking_HTTPIngressRuleValue_To_v1beta1_HTTPIngressRuleValue(in, out, s)
}

func autoConvert_v1beta1_Ingress_To_networking_Ingress(in *v1beta1.Ingress, out *networking.Ingress, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_IngressSpec_To_networking_IngressSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_IngressStatus_To_networking_IngressStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_Ingress_To_networking_Ingress is an autogenerated conversion function.
func Convert_v1beta1_Ingress_To_networking_Ingress(in *v1beta1.Ingress, out *networking.Ingress, s conversion.Scope) error {
	return autoConvert_v1beta1_Ingress_To_networking_Ingress(in, out, s)
}

func autoConvert_networking_Ingress_To_v1beta1_Ingress(in *networking.Ingress, out *v1beta1.Ingress, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_networking_IngressSpec_To_v1beta1_IngressSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_networking_IngressStatus_To_v1beta1_IngressStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_networking_Ingress_To_v1beta1_Ingress is an autogenerated conversion function.
func Convert_networking_Ingress_To_v1beta1_Ingress(in *networking.Ingress, out *v1beta1.Ingress, s conversion.Scope) error {
	return autoConvert_networking_Ingress_To_v1beta1_Ingress(in, out, s)
}

func autoConvert_v1beta1_IngressBackend_To_networking_IngressBackend(in *v1beta1.IngressBackend, out *networking.IngressBackend, s conversion.Scope) error {
	// WARNING: in.ServiceName requires manual conversion: does not exist in peer-type
	// WARNING: in.ServicePort requires manual conversion: does not exist in peer-type
	out.Resource = (*core.TypedLocalObjectReference)(unsafe.Pointer(in.Resource))
	return nil
}

func autoConvert_networking_IngressBackend_To_v1beta1_IngressBackend(in *networking.IngressBackend, out *v1beta1.IngressBackend, s conversion.Scope) error {
	// WARNING: in.Service requires manual conversion: does not exist in peer-type
	out.Resource = (*v1.TypedLocalObjectReference)(unsafe.Pointer(in.Resource))
	return nil
}

func autoConvert_v1beta1_IngressClass_To_networking_IngressClass(in *v1beta1.IngressClass, out *networking.IngressClass, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_IngressClassSpec_To_networking_IngressClassSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_IngressClass_To_networking_IngressClass is an autogenerated conversion function.
func Convert_v1beta1_IngressClass_To_networking_IngressClass(in *v1beta1.IngressClass, out *networking.IngressClass, s conversion.Scope) error {
	return autoConvert_v1beta1_IngressClass_To_networking_IngressClass(in, out, s)
}

func autoConvert_networking_IngressClass_To_v1beta1_IngressClass(in *networking.IngressClass, out *v1beta1.IngressClass, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_networking_IngressClassSpec_To_v1beta1_IngressClassSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_networking_IngressClass_To_v1beta1_IngressClass is an autogenerated conversion function.
func Convert_networking_IngressClass_To_v1beta1_IngressClass(in *networking.IngressClass, out *v1beta1.IngressClass, s conversion.Scope) error {
	return autoConvert_networking_IngressClass_To_v1beta1_IngressClass(in, out, s)
}

func autoConvert_v1beta1_IngressClassList_To_networking_IngressClassList(in *v1beta1.IngressClassList, out *networking.IngressClassList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]networking.IngressClass)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1beta1_IngressClassList_To_networking_IngressClassList is an autogenerated conversion function.
func Convert_v1beta1_IngressClassList_To_networking_IngressClassList(in *v1beta1.IngressClassList, out *networking.IngressClassList, s conversion.Scope) error {
	return autoConvert_v1beta1_IngressClassList_To_networking_IngressClassList(in, out, s)
}

func autoConvert_networking_IngressClassList_To_v1beta1_IngressClassList(in *networking.IngressClassList, out *v1beta1.IngressClassList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]v1beta1.IngressClass)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_networking_IngressClassList_To_v1beta1_IngressClassList is an autogenerated conversion function.
func Convert_networking_IngressClassList_To_v1beta1_IngressClassList(in *networking.IngressClassList, out *v1beta1.IngressClassList, s conversion.Scope) error {
	return autoConvert_networking_IngressClassList_To_v1beta1_IngressClassList(in, out, s)
}

func autoConvert_v1beta1_IngressClassSpec_To_networking_IngressClassSpec(in *v1beta1.IngressClassSpec, out *networking.IngressClassSpec, s conversion.Scope) error {
	out.Controller = in.Controller
	out.Parameters = (*core.TypedLocalObjectReference)(unsafe.Pointer(in.Parameters))
	return nil
}

// Convert_v1beta1_IngressClassSpec_To_networking_IngressClassSpec is an autogenerated conversion function.
func Convert_v1beta1_IngressClassSpec_To_networking_IngressClassSpec(in *v1beta1.IngressClassSpec, out *networking.IngressClassSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_IngressClassSpec_To_networking_IngressClassSpec(in, out, s)
}

func autoConvert_networking_IngressClassSpec_To_v1beta1_IngressClassSpec(in *networking.IngressClassSpec, out *v1beta1.IngressClassSpec, s conversion.Scope) error {
	out.Controller = in.Controller
	out.Parameters = (*v1.TypedLocalObjectReference)(unsafe.Pointer(in.Parameters))
	return nil
}

// Convert_networking_IngressClassSpec_To_v1beta1_IngressClassSpec is an autogenerated conversion function.
func Convert_networking_IngressClassSpec_To_v1beta1_IngressClassSpec(in *networking.IngressClassSpec, out *v1beta1.IngressClassSpec, s conversion.Scope) error {
	return autoConvert_networking_IngressClassSpec_To_v1beta1_IngressClassSpec(in, out, s)
}

func autoConvert_v1beta1_IngressList_To_networking_IngressList(in *v1beta1.IngressList, out *networking.IngressList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]networking.Ingress, len(*in))
		for i := range *in {
			if err := Convert_v1beta1_Ingress_To_networking_Ingress(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1beta1_IngressList_To_networking_IngressList is an autogenerated conversion function.
func Convert_v1beta1_IngressList_To_networking_IngressList(in *v1beta1.IngressList, out *networking.IngressList, s conversion.Scope) error {
	return autoConvert_v1beta1_IngressList_To_networking_IngressList(in, out, s)
}

func autoConvert_networking_IngressList_To_v1beta1_IngressList(in *networking.IngressList, out *v1beta1.IngressList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1beta1.Ingress, len(*in))
		for i := range *in {
			if err := Convert_networking_Ingress_To_v1beta1_Ingress(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_networking_IngressList_To_v1beta1_IngressList is an autogenerated conversion function.
func Convert_networking_IngressList_To_v1beta1_IngressList(in *networking.IngressList, out *v1beta1.IngressList, s conversion.Scope) error {
	return autoConvert_networking_IngressList_To_v1beta1_IngressList(in, out, s)
}

func autoConvert_v1beta1_IngressRule_To_networking_IngressRule(in *v1beta1.IngressRule, out *networking.IngressRule, s conversion.Scope) error {
	out.Host = in.Host
	if err := Convert_v1beta1_IngressRuleValue_To_networking_IngressRuleValue(&in.IngressRuleValue, &out.IngressRuleValue, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_IngressRule_To_networking_IngressRule is an autogenerated conversion function.
func Convert_v1beta1_IngressRule_To_networking_IngressRule(in *v1beta1.IngressRule, out *networking.IngressRule, s conversion.Scope) error {
	return autoConvert_v1beta1_IngressRule_To_networking_IngressRule(in, out, s)
}

func autoConvert_networking_IngressRule_To_v1beta1_IngressRule(in *networking.IngressRule, out *v1beta1.IngressRule, s conversion.Scope) error {
	out.Host = in.Host
	if err := Convert_networking_IngressRuleValue_To_v1beta1_IngressRuleValue(&in.IngressRuleValue, &out.IngressRuleValue, s); err != nil {
		return err
	}
	return nil
}

// Convert_networking_IngressRule_To_v1beta1_IngressRule is an autogenerated conversion function.
func Convert_networking_IngressRule_To_v1beta1_IngressRule(in *networking.IngressRule, out *v1beta1.IngressRule, s conversion.Scope) error {
	return autoConvert_networking_IngressRule_To_v1beta1_IngressRule(in, out, s)
}

func autoConvert_v1beta1_IngressRuleValue_To_networking_IngressRuleValue(in *v1beta1.IngressRuleValue, out *networking.IngressRuleValue, s conversion.Scope) error {
	if in.HTTP != nil {
		in, out := &in.HTTP, &out.HTTP
		*out = new(networking.HTTPIngressRuleValue)
		if err := Convert_v1beta1_HTTPIngressRuleValue_To_networking_HTTPIngressRuleValue(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.HTTP = nil
	}
	return nil
}

// Convert_v1beta1_IngressRuleValue_To_networking_IngressRuleValue is an autogenerated conversion function.
func Convert_v1beta1_IngressRuleValue_To_networking_IngressRuleValue(in *v1beta1.IngressRuleValue, out *networking.IngressRuleValue, s conversion.Scope) error {
	return autoConvert_v1beta1_IngressRuleValue_To_networking_IngressRuleValue(in, out, s)
}

func autoConvert_networking_IngressRuleValue_To_v1beta1_IngressRuleValue(in *networking.IngressRuleValue, out *v1beta1.IngressRuleValue, s conversion.Scope) error {
	if in.HTTP != nil {
		in, out := &in.HTTP, &out.HTTP
		*out = new(v1beta1.HTTPIngressRuleValue)
		if err := Convert_networking_HTTPIngressRuleValue_To_v1beta1_HTTPIngressRuleValue(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.HTTP = nil
	}
	return nil
}

// Convert_networking_IngressRuleValue_To_v1beta1_IngressRuleValue is an autogenerated conversion function.
func Convert_networking_IngressRuleValue_To_v1beta1_IngressRuleValue(in *networking.IngressRuleValue, out *v1beta1.IngressRuleValue, s conversion.Scope) error {
	return autoConvert_networking_IngressRuleValue_To_v1beta1_IngressRuleValue(in, out, s)
}

func autoConvert_v1beta1_IngressSpec_To_networking_IngressSpec(in *v1beta1.IngressSpec, out *networking.IngressSpec, s conversion.Scope) error {
	out.IngressClassName = (*string)(unsafe.Pointer(in.IngressClassName))
	// WARNING: in.Backend requires manual conversion: does not exist in peer-type
	out.TLS = *(*[]networking.IngressTLS)(unsafe.Pointer(&in.TLS))
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]networking.IngressRule, len(*in))
		for i := range *in {
			if err := Convert_v1beta1_IngressRule_To_networking_IngressRule(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Rules = nil
	}
	return nil
}

func autoConvert_networking_IngressSpec_To_v1beta1_IngressSpec(in *networking.IngressSpec, out *v1beta1.IngressSpec, s conversion.Scope) error {
	out.IngressClassName = (*string)(unsafe.Pointer(in.IngressClassName))
	// WARNING: in.DefaultBackend requires manual conversion: does not exist in peer-type
	out.TLS = *(*[]v1beta1.IngressTLS)(unsafe.Pointer(&in.TLS))
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]v1beta1.IngressRule, len(*in))
		for i := range *in {
			if err := Convert_networking_IngressRule_To_v1beta1_IngressRule(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Rules = nil
	}
	return nil
}

func autoConvert_v1beta1_IngressStatus_To_networking_IngressStatus(in *v1beta1.IngressStatus, out *networking.IngressStatus, s conversion.Scope) error {
	if err := corev1.Convert_v1_LoadBalancerStatus_To_core_LoadBalancerStatus(&in.LoadBalancer, &out.LoadBalancer, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_IngressStatus_To_networking_IngressStatus is an autogenerated conversion function.
func Convert_v1beta1_IngressStatus_To_networking_IngressStatus(in *v1beta1.IngressStatus, out *networking.IngressStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_IngressStatus_To_networking_IngressStatus(in, out, s)
}

func autoConvert_networking_IngressStatus_To_v1beta1_IngressStatus(in *networking.IngressStatus, out *v1beta1.IngressStatus, s conversion.Scope) error {
	if err := corev1.Convert_core_LoadBalancerStatus_To_v1_LoadBalancerStatus(&in.LoadBalancer, &out.LoadBalancer, s); err != nil {
		return err
	}
	return nil
}

// Convert_networking_IngressStatus_To_v1beta1_IngressStatus is an autogenerated conversion function.
func Convert_networking_IngressStatus_To_v1beta1_IngressStatus(in *networking.IngressStatus, out *v1beta1.IngressStatus, s conversion.Scope) error {
	return autoConvert_networking_IngressStatus_To_v1beta1_IngressStatus(in, out, s)
}

func autoConvert_v1beta1_IngressTLS_To_networking_IngressTLS(in *v1beta1.IngressTLS, out *networking.IngressTLS, s conversion.Scope) error {
	out.Hosts = *(*[]string)(unsafe.Pointer(&in.Hosts))
	out.SecretName = in.SecretName
	return nil
}

// Convert_v1beta1_IngressTLS_To_networking_IngressTLS is an autogenerated conversion function.
func Convert_v1beta1_IngressTLS_To_networking_IngressTLS(in *v1beta1.IngressTLS, out *networking.IngressTLS, s conversion.Scope) error {
	return autoConvert_v1beta1_IngressTLS_To_networking_IngressTLS(in, out, s)
}

func autoConvert_networking_IngressTLS_To_v1beta1_IngressTLS(in *networking.IngressTLS, out *v1beta1.IngressTLS, s conversion.Scope) error {
	out.Hosts = *(*[]string)(unsafe.Pointer(&in.Hosts))
	out.SecretName = in.SecretName
	return nil
}

// Convert_networking_IngressTLS_To_v1beta1_IngressTLS is an autogenerated conversion function.
func Convert_networking_IngressTLS_To_v1beta1_IngressTLS(in *networking.IngressTLS, out *v1beta1.IngressTLS, s conversion.Scope) error {
	return autoConvert_networking_IngressTLS_To_v1beta1_IngressTLS(in, out, s)
}
