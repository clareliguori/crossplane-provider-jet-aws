// +build !ignore_autogenerated

/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcmpcaCertificate) DeepCopyInto(out *AcmpcaCertificate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcmpcaCertificate.
func (in *AcmpcaCertificate) DeepCopy() *AcmpcaCertificate {
	if in == nil {
		return nil
	}
	out := new(AcmpcaCertificate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AcmpcaCertificate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcmpcaCertificateAuthority) DeepCopyInto(out *AcmpcaCertificateAuthority) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcmpcaCertificateAuthority.
func (in *AcmpcaCertificateAuthority) DeepCopy() *AcmpcaCertificateAuthority {
	if in == nil {
		return nil
	}
	out := new(AcmpcaCertificateAuthority)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AcmpcaCertificateAuthority) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcmpcaCertificateAuthorityCertificate) DeepCopyInto(out *AcmpcaCertificateAuthorityCertificate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcmpcaCertificateAuthorityCertificate.
func (in *AcmpcaCertificateAuthorityCertificate) DeepCopy() *AcmpcaCertificateAuthorityCertificate {
	if in == nil {
		return nil
	}
	out := new(AcmpcaCertificateAuthorityCertificate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AcmpcaCertificateAuthorityCertificate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcmpcaCertificateAuthorityCertificateList) DeepCopyInto(out *AcmpcaCertificateAuthorityCertificateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AcmpcaCertificateAuthorityCertificate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcmpcaCertificateAuthorityCertificateList.
func (in *AcmpcaCertificateAuthorityCertificateList) DeepCopy() *AcmpcaCertificateAuthorityCertificateList {
	if in == nil {
		return nil
	}
	out := new(AcmpcaCertificateAuthorityCertificateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AcmpcaCertificateAuthorityCertificateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcmpcaCertificateAuthorityCertificateObservation) DeepCopyInto(out *AcmpcaCertificateAuthorityCertificateObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcmpcaCertificateAuthorityCertificateObservation.
func (in *AcmpcaCertificateAuthorityCertificateObservation) DeepCopy() *AcmpcaCertificateAuthorityCertificateObservation {
	if in == nil {
		return nil
	}
	out := new(AcmpcaCertificateAuthorityCertificateObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcmpcaCertificateAuthorityCertificateParameters) DeepCopyInto(out *AcmpcaCertificateAuthorityCertificateParameters) {
	*out = *in
	if in.CertificateChain != nil {
		in, out := &in.CertificateChain, &out.CertificateChain
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcmpcaCertificateAuthorityCertificateParameters.
func (in *AcmpcaCertificateAuthorityCertificateParameters) DeepCopy() *AcmpcaCertificateAuthorityCertificateParameters {
	if in == nil {
		return nil
	}
	out := new(AcmpcaCertificateAuthorityCertificateParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcmpcaCertificateAuthorityCertificateSpec) DeepCopyInto(out *AcmpcaCertificateAuthorityCertificateSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcmpcaCertificateAuthorityCertificateSpec.
func (in *AcmpcaCertificateAuthorityCertificateSpec) DeepCopy() *AcmpcaCertificateAuthorityCertificateSpec {
	if in == nil {
		return nil
	}
	out := new(AcmpcaCertificateAuthorityCertificateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcmpcaCertificateAuthorityCertificateStatus) DeepCopyInto(out *AcmpcaCertificateAuthorityCertificateStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcmpcaCertificateAuthorityCertificateStatus.
func (in *AcmpcaCertificateAuthorityCertificateStatus) DeepCopy() *AcmpcaCertificateAuthorityCertificateStatus {
	if in == nil {
		return nil
	}
	out := new(AcmpcaCertificateAuthorityCertificateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcmpcaCertificateAuthorityList) DeepCopyInto(out *AcmpcaCertificateAuthorityList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AcmpcaCertificateAuthority, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcmpcaCertificateAuthorityList.
func (in *AcmpcaCertificateAuthorityList) DeepCopy() *AcmpcaCertificateAuthorityList {
	if in == nil {
		return nil
	}
	out := new(AcmpcaCertificateAuthorityList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AcmpcaCertificateAuthorityList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcmpcaCertificateAuthorityObservation) DeepCopyInto(out *AcmpcaCertificateAuthorityObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcmpcaCertificateAuthorityObservation.
func (in *AcmpcaCertificateAuthorityObservation) DeepCopy() *AcmpcaCertificateAuthorityObservation {
	if in == nil {
		return nil
	}
	out := new(AcmpcaCertificateAuthorityObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcmpcaCertificateAuthorityParameters) DeepCopyInto(out *AcmpcaCertificateAuthorityParameters) {
	*out = *in
	if in.CertificateAuthorityConfiguration != nil {
		in, out := &in.CertificateAuthorityConfiguration, &out.CertificateAuthorityConfiguration
		*out = make([]CertificateAuthorityConfigurationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.PermanentDeletionTimeInDays != nil {
		in, out := &in.PermanentDeletionTimeInDays, &out.PermanentDeletionTimeInDays
		*out = new(int64)
		**out = **in
	}
	if in.RevocationConfiguration != nil {
		in, out := &in.RevocationConfiguration, &out.RevocationConfiguration
		*out = make([]RevocationConfigurationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcmpcaCertificateAuthorityParameters.
func (in *AcmpcaCertificateAuthorityParameters) DeepCopy() *AcmpcaCertificateAuthorityParameters {
	if in == nil {
		return nil
	}
	out := new(AcmpcaCertificateAuthorityParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcmpcaCertificateAuthoritySpec) DeepCopyInto(out *AcmpcaCertificateAuthoritySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcmpcaCertificateAuthoritySpec.
func (in *AcmpcaCertificateAuthoritySpec) DeepCopy() *AcmpcaCertificateAuthoritySpec {
	if in == nil {
		return nil
	}
	out := new(AcmpcaCertificateAuthoritySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcmpcaCertificateAuthorityStatus) DeepCopyInto(out *AcmpcaCertificateAuthorityStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcmpcaCertificateAuthorityStatus.
func (in *AcmpcaCertificateAuthorityStatus) DeepCopy() *AcmpcaCertificateAuthorityStatus {
	if in == nil {
		return nil
	}
	out := new(AcmpcaCertificateAuthorityStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcmpcaCertificateList) DeepCopyInto(out *AcmpcaCertificateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AcmpcaCertificate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcmpcaCertificateList.
func (in *AcmpcaCertificateList) DeepCopy() *AcmpcaCertificateList {
	if in == nil {
		return nil
	}
	out := new(AcmpcaCertificateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AcmpcaCertificateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcmpcaCertificateObservation) DeepCopyInto(out *AcmpcaCertificateObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcmpcaCertificateObservation.
func (in *AcmpcaCertificateObservation) DeepCopy() *AcmpcaCertificateObservation {
	if in == nil {
		return nil
	}
	out := new(AcmpcaCertificateObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcmpcaCertificateParameters) DeepCopyInto(out *AcmpcaCertificateParameters) {
	*out = *in
	if in.TemplateArn != nil {
		in, out := &in.TemplateArn, &out.TemplateArn
		*out = new(string)
		**out = **in
	}
	if in.Validity != nil {
		in, out := &in.Validity, &out.Validity
		*out = make([]ValidityParameters, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcmpcaCertificateParameters.
func (in *AcmpcaCertificateParameters) DeepCopy() *AcmpcaCertificateParameters {
	if in == nil {
		return nil
	}
	out := new(AcmpcaCertificateParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcmpcaCertificateSpec) DeepCopyInto(out *AcmpcaCertificateSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcmpcaCertificateSpec.
func (in *AcmpcaCertificateSpec) DeepCopy() *AcmpcaCertificateSpec {
	if in == nil {
		return nil
	}
	out := new(AcmpcaCertificateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcmpcaCertificateStatus) DeepCopyInto(out *AcmpcaCertificateStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcmpcaCertificateStatus.
func (in *AcmpcaCertificateStatus) DeepCopy() *AcmpcaCertificateStatus {
	if in == nil {
		return nil
	}
	out := new(AcmpcaCertificateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateAuthorityConfigurationObservation) DeepCopyInto(out *CertificateAuthorityConfigurationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateAuthorityConfigurationObservation.
func (in *CertificateAuthorityConfigurationObservation) DeepCopy() *CertificateAuthorityConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(CertificateAuthorityConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateAuthorityConfigurationParameters) DeepCopyInto(out *CertificateAuthorityConfigurationParameters) {
	*out = *in
	if in.Subject != nil {
		in, out := &in.Subject, &out.Subject
		*out = make([]SubjectParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateAuthorityConfigurationParameters.
func (in *CertificateAuthorityConfigurationParameters) DeepCopy() *CertificateAuthorityConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(CertificateAuthorityConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CrlConfigurationObservation) DeepCopyInto(out *CrlConfigurationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CrlConfigurationObservation.
func (in *CrlConfigurationObservation) DeepCopy() *CrlConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(CrlConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CrlConfigurationParameters) DeepCopyInto(out *CrlConfigurationParameters) {
	*out = *in
	if in.CustomCname != nil {
		in, out := &in.CustomCname, &out.CustomCname
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.S3BucketName != nil {
		in, out := &in.S3BucketName, &out.S3BucketName
		*out = new(string)
		**out = **in
	}
	if in.S3ObjectACL != nil {
		in, out := &in.S3ObjectACL, &out.S3ObjectACL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CrlConfigurationParameters.
func (in *CrlConfigurationParameters) DeepCopy() *CrlConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(CrlConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RevocationConfigurationObservation) DeepCopyInto(out *RevocationConfigurationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RevocationConfigurationObservation.
func (in *RevocationConfigurationObservation) DeepCopy() *RevocationConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(RevocationConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RevocationConfigurationParameters) DeepCopyInto(out *RevocationConfigurationParameters) {
	*out = *in
	if in.CrlConfiguration != nil {
		in, out := &in.CrlConfiguration, &out.CrlConfiguration
		*out = make([]CrlConfigurationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RevocationConfigurationParameters.
func (in *RevocationConfigurationParameters) DeepCopy() *RevocationConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(RevocationConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubjectObservation) DeepCopyInto(out *SubjectObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubjectObservation.
func (in *SubjectObservation) DeepCopy() *SubjectObservation {
	if in == nil {
		return nil
	}
	out := new(SubjectObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubjectParameters) DeepCopyInto(out *SubjectParameters) {
	*out = *in
	if in.CommonName != nil {
		in, out := &in.CommonName, &out.CommonName
		*out = new(string)
		**out = **in
	}
	if in.Country != nil {
		in, out := &in.Country, &out.Country
		*out = new(string)
		**out = **in
	}
	if in.DistinguishedNameQualifier != nil {
		in, out := &in.DistinguishedNameQualifier, &out.DistinguishedNameQualifier
		*out = new(string)
		**out = **in
	}
	if in.GenerationQualifier != nil {
		in, out := &in.GenerationQualifier, &out.GenerationQualifier
		*out = new(string)
		**out = **in
	}
	if in.GivenName != nil {
		in, out := &in.GivenName, &out.GivenName
		*out = new(string)
		**out = **in
	}
	if in.Initials != nil {
		in, out := &in.Initials, &out.Initials
		*out = new(string)
		**out = **in
	}
	if in.Locality != nil {
		in, out := &in.Locality, &out.Locality
		*out = new(string)
		**out = **in
	}
	if in.Organization != nil {
		in, out := &in.Organization, &out.Organization
		*out = new(string)
		**out = **in
	}
	if in.OrganizationalUnit != nil {
		in, out := &in.OrganizationalUnit, &out.OrganizationalUnit
		*out = new(string)
		**out = **in
	}
	if in.Pseudonym != nil {
		in, out := &in.Pseudonym, &out.Pseudonym
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.Surname != nil {
		in, out := &in.Surname, &out.Surname
		*out = new(string)
		**out = **in
	}
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubjectParameters.
func (in *SubjectParameters) DeepCopy() *SubjectParameters {
	if in == nil {
		return nil
	}
	out := new(SubjectParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidityObservation) DeepCopyInto(out *ValidityObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidityObservation.
func (in *ValidityObservation) DeepCopy() *ValidityObservation {
	if in == nil {
		return nil
	}
	out := new(ValidityObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidityParameters) DeepCopyInto(out *ValidityParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidityParameters.
func (in *ValidityParameters) DeepCopy() *ValidityParameters {
	if in == nil {
		return nil
	}
	out := new(ValidityParameters)
	in.DeepCopyInto(out)
	return out
}