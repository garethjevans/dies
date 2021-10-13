//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 the original author or authors.

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

// Code generated by diegen. DO NOT EDIT.

package v1

import (
	json "encoding/json"
	fmtx "fmt"
	metav1 "github.com/scothis/dies/apis/meta/v1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	GroupVersion  = schema.GroupVersion{Group: "apiextensions.k8s.io", Version: "v1"}
	SchemeBuilder = runtime.NewSchemeBuilder()
	AddToScheme   = SchemeBuilder.AddToScheme
)

type CustomResourceDefinitionDie struct {
	metav1.FrozenObjectMeta
	mutable bool
	r       apiextensionsv1.CustomResourceDefinition
}

var CustomResourceDefinitionBlank = (&CustomResourceDefinitionDie{}).DieFeed(apiextensionsv1.CustomResourceDefinition{})

func (d *CustomResourceDefinitionDie) DieImmutable(immutable bool) *CustomResourceDefinitionDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

func (d *CustomResourceDefinitionDie) DieFeed(r apiextensionsv1.CustomResourceDefinition) *CustomResourceDefinitionDie {
	if d.mutable {
		d.FrozenObjectMeta = metav1.FreezeObjectMeta(r.ObjectMeta)
		d.r = r
		return d
	}
	return &CustomResourceDefinitionDie{
		FrozenObjectMeta: metav1.FreezeObjectMeta(r.ObjectMeta),
		mutable:          d.mutable,
		r:                r,
	}
}

func (d *CustomResourceDefinitionDie) DieRelease() apiextensionsv1.CustomResourceDefinition {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

func (d *CustomResourceDefinitionDie) DieStamp(fn func(r *apiextensionsv1.CustomResourceDefinition)) *CustomResourceDefinitionDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

func (d *CustomResourceDefinitionDie) DeepCopy() *CustomResourceDefinitionDie {
	r := *d.r.DeepCopy()
	return &CustomResourceDefinitionDie{
		FrozenObjectMeta: metav1.FreezeObjectMeta(r.ObjectMeta),
		mutable:          d.mutable,
		r:                r,
	}
}

func (d *CustomResourceDefinitionDie) DeepCopyObject() runtime.Object {
	return d.DeepCopy()
}

func (d *CustomResourceDefinitionDie) GetObjectKind() schema.ObjectKind {
	r := d.DieRelease()
	return r.GetObjectKind()
}

func (d *CustomResourceDefinitionDie) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.r)
}

func (d *CustomResourceDefinitionDie) UnmarshalJSON(b []byte) error {
	if d == CustomResourceDefinitionBlank {
		return fmtx.Errorf("cannot unmarshing into the root object, create a copy first")
	}
	r := &apiextensionsv1.CustomResourceDefinition{}
	err := json.Unmarshal(b, r)
	*d = *d.DieFeed(*r)
	return err
}

func (d *CustomResourceDefinitionDie) MetadataDie(fn func(d *metav1.ObjectMetaDie)) *CustomResourceDefinitionDie {
	return d.DieStamp(func(r *apiextensionsv1.CustomResourceDefinition) {
		d := metav1.ObjectMetaBlank.DieImmutable(false).DieFeed(r.ObjectMeta)
		fn(d)
		r.ObjectMeta = d.DieRelease()
	})
}

func (d *CustomResourceDefinitionDie) Spec(v apiextensionsv1.CustomResourceDefinitionSpec) *CustomResourceDefinitionDie {
	return d.DieStamp(func(r *apiextensionsv1.CustomResourceDefinition) {
		r.Spec = v
	})
}

func (d *CustomResourceDefinitionDie) SpecDie(fn func(d *CustomResourceDefinitionSpecDie)) *CustomResourceDefinitionDie {
	return d.DieStamp(func(r *apiextensionsv1.CustomResourceDefinition) {
		d := CustomResourceDefinitionSpecBlank.DieImmutable(false).DieFeed(r.Spec)
		fn(d)
		r.Spec = d.DieRelease()
	})
}

func (d *CustomResourceDefinitionDie) Status(v apiextensionsv1.CustomResourceDefinitionStatus) *CustomResourceDefinitionDie {
	return d.DieStamp(func(r *apiextensionsv1.CustomResourceDefinition) {
		r.Status = v
	})
}

func (d *CustomResourceDefinitionDie) StatusDie(fn func(d *CustomResourceDefinitionStatusDie)) *CustomResourceDefinitionDie {
	return d.DieStamp(func(r *apiextensionsv1.CustomResourceDefinition) {
		d := CustomResourceDefinitionStatusBlank.DieImmutable(false).DieFeed(r.Status)
		fn(d)
		r.Status = d.DieRelease()
	})
}

var _ apismetav1.Object = (*CustomResourceDefinitionDie)(nil)
var _ apismetav1.ObjectMetaAccessor = (*CustomResourceDefinitionDie)(nil)
var _ runtime.Object = (*CustomResourceDefinitionDie)(nil)

func init() {
	SchemeBuilder.Register(func(s *runtime.Scheme) error {
		s.AddKnownTypeWithName(GroupVersion.WithKind("CustomResourceDefinition"), &CustomResourceDefinitionDie{})
		return nil
	})
}

type CustomResourceDefinitionSpecDie struct {
	mutable bool
	r       apiextensionsv1.CustomResourceDefinitionSpec
}

var CustomResourceDefinitionSpecBlank = (&CustomResourceDefinitionSpecDie{}).DieFeed(apiextensionsv1.CustomResourceDefinitionSpec{})

func (d *CustomResourceDefinitionSpecDie) DieImmutable(immutable bool) *CustomResourceDefinitionSpecDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

func (d *CustomResourceDefinitionSpecDie) DieFeed(r apiextensionsv1.CustomResourceDefinitionSpec) *CustomResourceDefinitionSpecDie {
	if d.mutable {
		d.r = r
		return d
	}
	return &CustomResourceDefinitionSpecDie{
		mutable: d.mutable,
		r:       r,
	}
}

func (d *CustomResourceDefinitionSpecDie) DieRelease() apiextensionsv1.CustomResourceDefinitionSpec {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

func (d *CustomResourceDefinitionSpecDie) DieStamp(fn func(r *apiextensionsv1.CustomResourceDefinitionSpec)) *CustomResourceDefinitionSpecDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

func (d *CustomResourceDefinitionSpecDie) DeepCopy() *CustomResourceDefinitionSpecDie {
	r := *d.r.DeepCopy()
	return &CustomResourceDefinitionSpecDie{
		mutable: d.mutable,
		r:       r,
	}
}

func (d *CustomResourceDefinitionSpecDie) Group(v string) *CustomResourceDefinitionSpecDie {
	return d.DieStamp(func(r *apiextensionsv1.CustomResourceDefinitionSpec) {
		r.Group = v
	})
}

func (d *CustomResourceDefinitionSpecDie) Names(v apiextensionsv1.CustomResourceDefinitionNames) *CustomResourceDefinitionSpecDie {
	return d.DieStamp(func(r *apiextensionsv1.CustomResourceDefinitionSpec) {
		r.Names = v
	})
}

func (d *CustomResourceDefinitionSpecDie) Scope(v apiextensionsv1.ResourceScope) *CustomResourceDefinitionSpecDie {
	return d.DieStamp(func(r *apiextensionsv1.CustomResourceDefinitionSpec) {
		r.Scope = v
	})
}

func (d *CustomResourceDefinitionSpecDie) Versions(v ...apiextensionsv1.CustomResourceDefinitionVersion) *CustomResourceDefinitionSpecDie {
	return d.DieStamp(func(r *apiextensionsv1.CustomResourceDefinitionSpec) {
		r.Versions = v
	})
}

func (d *CustomResourceDefinitionSpecDie) Conversion(v *apiextensionsv1.CustomResourceConversion) *CustomResourceDefinitionSpecDie {
	return d.DieStamp(func(r *apiextensionsv1.CustomResourceDefinitionSpec) {
		r.Conversion = v
	})
}

func (d *CustomResourceDefinitionSpecDie) PreserveUnknownFields(v bool) *CustomResourceDefinitionSpecDie {
	return d.DieStamp(func(r *apiextensionsv1.CustomResourceDefinitionSpec) {
		r.PreserveUnknownFields = v
	})
}

type CustomResourceDefinitionStatusDie struct {
	mutable bool
	r       apiextensionsv1.CustomResourceDefinitionStatus
}

var CustomResourceDefinitionStatusBlank = (&CustomResourceDefinitionStatusDie{}).DieFeed(apiextensionsv1.CustomResourceDefinitionStatus{})

func (d *CustomResourceDefinitionStatusDie) DieImmutable(immutable bool) *CustomResourceDefinitionStatusDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

func (d *CustomResourceDefinitionStatusDie) DieFeed(r apiextensionsv1.CustomResourceDefinitionStatus) *CustomResourceDefinitionStatusDie {
	if d.mutable {
		d.r = r
		return d
	}
	return &CustomResourceDefinitionStatusDie{
		mutable: d.mutable,
		r:       r,
	}
}

func (d *CustomResourceDefinitionStatusDie) DieRelease() apiextensionsv1.CustomResourceDefinitionStatus {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

func (d *CustomResourceDefinitionStatusDie) DieStamp(fn func(r *apiextensionsv1.CustomResourceDefinitionStatus)) *CustomResourceDefinitionStatusDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

func (d *CustomResourceDefinitionStatusDie) DeepCopy() *CustomResourceDefinitionStatusDie {
	r := *d.r.DeepCopy()
	return &CustomResourceDefinitionStatusDie{
		mutable: d.mutable,
		r:       r,
	}
}

func (d *CustomResourceDefinitionStatusDie) Conditions(v ...apiextensionsv1.CustomResourceDefinitionCondition) *CustomResourceDefinitionStatusDie {
	return d.DieStamp(func(r *apiextensionsv1.CustomResourceDefinitionStatus) {
		r.Conditions = v
	})
}

func (d *CustomResourceDefinitionStatusDie) AcceptedNames(v apiextensionsv1.CustomResourceDefinitionNames) *CustomResourceDefinitionStatusDie {
	return d.DieStamp(func(r *apiextensionsv1.CustomResourceDefinitionStatus) {
		r.AcceptedNames = v
	})
}

func (d *CustomResourceDefinitionStatusDie) StoredVersions(v ...string) *CustomResourceDefinitionStatusDie {
	return d.DieStamp(func(r *apiextensionsv1.CustomResourceDefinitionStatus) {
		r.StoredVersions = v
	})
}
