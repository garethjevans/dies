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
	coordinationv1 "k8s.io/api/coordination/v1"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	GroupVersion  = schema.GroupVersion{Group: "coordination.k8s.io", Version: "v1"}
	SchemeBuilder = runtime.NewSchemeBuilder()
	AddToScheme   = SchemeBuilder.AddToScheme
)

type LeaseDie struct {
	metav1.FrozenObjectMeta
	mutable bool
	r       coordinationv1.Lease
}

var LeaseBlank = (&LeaseDie{}).DieFeed(coordinationv1.Lease{})

func (d *LeaseDie) DieImmutable(immutable bool) *LeaseDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

func (d *LeaseDie) DieFeed(r coordinationv1.Lease) *LeaseDie {
	if d.mutable {
		d.FrozenObjectMeta = metav1.FreezeObjectMeta(r.ObjectMeta)
		d.r = r
		return d
	}
	return &LeaseDie{
		FrozenObjectMeta: metav1.FreezeObjectMeta(r.ObjectMeta),
		mutable:          d.mutable,
		r:                r,
	}
}

func (d *LeaseDie) DieRelease() coordinationv1.Lease {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

func (d *LeaseDie) DieStamp(fn func(r *coordinationv1.Lease)) *LeaseDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

func (d *LeaseDie) DeepCopy() *LeaseDie {
	r := *d.r.DeepCopy()
	return &LeaseDie{
		FrozenObjectMeta: metav1.FreezeObjectMeta(r.ObjectMeta),
		mutable:          d.mutable,
		r:                r,
	}
}

func (d *LeaseDie) DeepCopyObject() runtime.Object {
	return d.DeepCopy()
}

func (d *LeaseDie) GetObjectKind() schema.ObjectKind {
	r := d.DieRelease()
	return r.GetObjectKind()
}

func (d *LeaseDie) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.r)
}

func (d *LeaseDie) UnmarshalJSON(b []byte) error {
	if d == LeaseBlank {
		return fmtx.Errorf("cannot unmarshing into the root object, create a copy first")
	}
	r := &coordinationv1.Lease{}
	err := json.Unmarshal(b, r)
	*d = *d.DieFeed(*r)
	return err
}

func (d *LeaseDie) MetadataDie(fn func(d *metav1.ObjectMetaDie)) *LeaseDie {
	return d.DieStamp(func(r *coordinationv1.Lease) {
		d := metav1.ObjectMetaBlank.DieImmutable(false).DieFeed(r.ObjectMeta)
		fn(d)
		r.ObjectMeta = d.DieRelease()
	})
}

func (d *LeaseDie) Spec(v coordinationv1.LeaseSpec) *LeaseDie {
	return d.DieStamp(func(r *coordinationv1.Lease) {
		r.Spec = v
	})
}

func (d *LeaseDie) SpecDie(fn func(d *LeaseSpecDie)) *LeaseDie {
	return d.DieStamp(func(r *coordinationv1.Lease) {
		d := LeaseSpecBlank.DieImmutable(false).DieFeed(r.Spec)
		fn(d)
		r.Spec = d.DieRelease()
	})
}

var _ apismetav1.Object = (*LeaseDie)(nil)
var _ apismetav1.ObjectMetaAccessor = (*LeaseDie)(nil)
var _ runtime.Object = (*LeaseDie)(nil)

func init() {
	SchemeBuilder.Register(func(s *runtime.Scheme) error {
		s.AddKnownTypeWithName(GroupVersion.WithKind("Lease"), &LeaseDie{})
		return nil
	})
}

type LeaseSpecDie struct {
	mutable bool
	r       coordinationv1.LeaseSpec
}

var LeaseSpecBlank = (&LeaseSpecDie{}).DieFeed(coordinationv1.LeaseSpec{})

func (d *LeaseSpecDie) DieImmutable(immutable bool) *LeaseSpecDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

func (d *LeaseSpecDie) DieFeed(r coordinationv1.LeaseSpec) *LeaseSpecDie {
	if d.mutable {
		d.r = r
		return d
	}
	return &LeaseSpecDie{
		mutable: d.mutable,
		r:       r,
	}
}

func (d *LeaseSpecDie) DieRelease() coordinationv1.LeaseSpec {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

func (d *LeaseSpecDie) DieStamp(fn func(r *coordinationv1.LeaseSpec)) *LeaseSpecDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

func (d *LeaseSpecDie) DeepCopy() *LeaseSpecDie {
	r := *d.r.DeepCopy()
	return &LeaseSpecDie{
		mutable: d.mutable,
		r:       r,
	}
}

func (d *LeaseSpecDie) HolderIdentity(v *string) *LeaseSpecDie {
	return d.DieStamp(func(r *coordinationv1.LeaseSpec) {
		r.HolderIdentity = v
	})
}

func (d *LeaseSpecDie) LeaseDurationSeconds(v *int32) *LeaseSpecDie {
	return d.DieStamp(func(r *coordinationv1.LeaseSpec) {
		r.LeaseDurationSeconds = v
	})
}

func (d *LeaseSpecDie) AcquireTime(v *apismetav1.MicroTime) *LeaseSpecDie {
	return d.DieStamp(func(r *coordinationv1.LeaseSpec) {
		r.AcquireTime = v
	})
}

func (d *LeaseSpecDie) RenewTime(v *apismetav1.MicroTime) *LeaseSpecDie {
	return d.DieStamp(func(r *coordinationv1.LeaseSpec) {
		r.RenewTime = v
	})
}

func (d *LeaseSpecDie) LeaseTransitions(v *int32) *LeaseSpecDie {
	return d.DieStamp(func(r *coordinationv1.LeaseSpec) {
		r.LeaseTransitions = v
	})
}
