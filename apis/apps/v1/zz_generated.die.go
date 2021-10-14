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
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

type DaemonSetDie struct {
	metav1.FrozenObjectMeta
	mutable bool
	r       appsv1.DaemonSet
}

var DaemonSetBlank = (&DaemonSetDie{}).DieFeed(appsv1.DaemonSet{})

func (d *DaemonSetDie) DieImmutable(immutable bool) *DaemonSetDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

func (d *DaemonSetDie) DieFeed(r appsv1.DaemonSet) *DaemonSetDie {
	if d.mutable {
		d.FrozenObjectMeta = metav1.FreezeObjectMeta(r.ObjectMeta)
		d.r = r
		return d
	}
	return &DaemonSetDie{
		FrozenObjectMeta: metav1.FreezeObjectMeta(r.ObjectMeta),
		mutable:          d.mutable,
		r:                r,
	}
}

func (d *DaemonSetDie) DieRelease() appsv1.DaemonSet {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

func (d *DaemonSetDie) DieStamp(fn func(r *appsv1.DaemonSet)) *DaemonSetDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

func (d *DaemonSetDie) DeepCopy() *DaemonSetDie {
	r := *d.r.DeepCopy()
	return &DaemonSetDie{
		FrozenObjectMeta: metav1.FreezeObjectMeta(r.ObjectMeta),
		mutable:          d.mutable,
		r:                r,
	}
}

func (d *DaemonSetDie) DeepCopyObject() runtime.Object {
	return d.r.DeepCopy()
}

func (d *DaemonSetDie) GetObjectKind() schema.ObjectKind {
	r := d.DieRelease()
	return r.GetObjectKind()
}

func (d *DaemonSetDie) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.r)
}

func (d *DaemonSetDie) UnmarshalJSON(b []byte) error {
	if d == DaemonSetBlank {
		return fmtx.Errorf("cannot unmarshal into the root object, create a copy first")
	}
	r := &appsv1.DaemonSet{}
	err := json.Unmarshal(b, r)
	*d = *d.DieFeed(*r)
	return err
}

func (d *DaemonSetDie) MetadataDie(fn func(d *metav1.ObjectMetaDie)) *DaemonSetDie {
	return d.DieStamp(func(r *appsv1.DaemonSet) {
		d := metav1.ObjectMetaBlank.DieImmutable(false).DieFeed(r.ObjectMeta)
		fn(d)
		r.ObjectMeta = d.DieRelease()
	})
}

func (d *DaemonSetDie) Spec(v appsv1.DaemonSetSpec) *DaemonSetDie {
	return d.DieStamp(func(r *appsv1.DaemonSet) {
		r.Spec = v
	})
}

func (d *DaemonSetDie) SpecDie(fn func(d *DaemonSetSpecDie)) *DaemonSetDie {
	return d.DieStamp(func(r *appsv1.DaemonSet) {
		d := DaemonSetSpecBlank.DieImmutable(false).DieFeed(r.Spec)
		fn(d)
		r.Spec = d.DieRelease()
	})
}

func (d *DaemonSetDie) Status(v appsv1.DaemonSetStatus) *DaemonSetDie {
	return d.DieStamp(func(r *appsv1.DaemonSet) {
		r.Status = v
	})
}

func (d *DaemonSetDie) StatusDie(fn func(d *DaemonSetStatusDie)) *DaemonSetDie {
	return d.DieStamp(func(r *appsv1.DaemonSet) {
		d := DaemonSetStatusBlank.DieImmutable(false).DieFeed(r.Status)
		fn(d)
		r.Status = d.DieRelease()
	})
}

var _ apismetav1.Object = (*DaemonSetDie)(nil)
var _ apismetav1.ObjectMetaAccessor = (*DaemonSetDie)(nil)
var _ runtime.Object = (*DaemonSetDie)(nil)

type DaemonSetSpecDie struct {
	mutable bool
	r       appsv1.DaemonSetSpec
}

var DaemonSetSpecBlank = (&DaemonSetSpecDie{}).DieFeed(appsv1.DaemonSetSpec{})

func (d *DaemonSetSpecDie) DieImmutable(immutable bool) *DaemonSetSpecDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

func (d *DaemonSetSpecDie) DieFeed(r appsv1.DaemonSetSpec) *DaemonSetSpecDie {
	if d.mutable {
		d.r = r
		return d
	}
	return &DaemonSetSpecDie{
		mutable: d.mutable,
		r:       r,
	}
}

func (d *DaemonSetSpecDie) DieRelease() appsv1.DaemonSetSpec {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

func (d *DaemonSetSpecDie) DieStamp(fn func(r *appsv1.DaemonSetSpec)) *DaemonSetSpecDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

func (d *DaemonSetSpecDie) DeepCopy() *DaemonSetSpecDie {
	r := *d.r.DeepCopy()
	return &DaemonSetSpecDie{
		mutable: d.mutable,
		r:       r,
	}
}

func (d *DaemonSetSpecDie) Selector(v *apismetav1.LabelSelector) *DaemonSetSpecDie {
	return d.DieStamp(func(r *appsv1.DaemonSetSpec) {
		r.Selector = v
	})
}

func (d *DaemonSetSpecDie) Template(v corev1.PodTemplateSpec) *DaemonSetSpecDie {
	return d.DieStamp(func(r *appsv1.DaemonSetSpec) {
		r.Template = v
	})
}

func (d *DaemonSetSpecDie) UpdateStrategy(v appsv1.DaemonSetUpdateStrategy) *DaemonSetSpecDie {
	return d.DieStamp(func(r *appsv1.DaemonSetSpec) {
		r.UpdateStrategy = v
	})
}

func (d *DaemonSetSpecDie) MinReadySeconds(v int32) *DaemonSetSpecDie {
	return d.DieStamp(func(r *appsv1.DaemonSetSpec) {
		r.MinReadySeconds = v
	})
}

func (d *DaemonSetSpecDie) RevisionHistoryLimit(v *int32) *DaemonSetSpecDie {
	return d.DieStamp(func(r *appsv1.DaemonSetSpec) {
		r.RevisionHistoryLimit = v
	})
}

type DaemonSetStatusDie struct {
	mutable bool
	r       appsv1.DaemonSetStatus
}

var DaemonSetStatusBlank = (&DaemonSetStatusDie{}).DieFeed(appsv1.DaemonSetStatus{})

func (d *DaemonSetStatusDie) DieImmutable(immutable bool) *DaemonSetStatusDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

func (d *DaemonSetStatusDie) DieFeed(r appsv1.DaemonSetStatus) *DaemonSetStatusDie {
	if d.mutable {
		d.r = r
		return d
	}
	return &DaemonSetStatusDie{
		mutable: d.mutable,
		r:       r,
	}
}

func (d *DaemonSetStatusDie) DieRelease() appsv1.DaemonSetStatus {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

func (d *DaemonSetStatusDie) DieStamp(fn func(r *appsv1.DaemonSetStatus)) *DaemonSetStatusDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

func (d *DaemonSetStatusDie) DeepCopy() *DaemonSetStatusDie {
	r := *d.r.DeepCopy()
	return &DaemonSetStatusDie{
		mutable: d.mutable,
		r:       r,
	}
}

func (d *DaemonSetStatusDie) CurrentNumberScheduled(v int32) *DaemonSetStatusDie {
	return d.DieStamp(func(r *appsv1.DaemonSetStatus) {
		r.CurrentNumberScheduled = v
	})
}

func (d *DaemonSetStatusDie) NumberMisscheduled(v int32) *DaemonSetStatusDie {
	return d.DieStamp(func(r *appsv1.DaemonSetStatus) {
		r.NumberMisscheduled = v
	})
}

func (d *DaemonSetStatusDie) DesiredNumberScheduled(v int32) *DaemonSetStatusDie {
	return d.DieStamp(func(r *appsv1.DaemonSetStatus) {
		r.DesiredNumberScheduled = v
	})
}

func (d *DaemonSetStatusDie) NumberReady(v int32) *DaemonSetStatusDie {
	return d.DieStamp(func(r *appsv1.DaemonSetStatus) {
		r.NumberReady = v
	})
}

func (d *DaemonSetStatusDie) ObservedGeneration(v int64) *DaemonSetStatusDie {
	return d.DieStamp(func(r *appsv1.DaemonSetStatus) {
		r.ObservedGeneration = v
	})
}

func (d *DaemonSetStatusDie) UpdatedNumberScheduled(v int32) *DaemonSetStatusDie {
	return d.DieStamp(func(r *appsv1.DaemonSetStatus) {
		r.UpdatedNumberScheduled = v
	})
}

func (d *DaemonSetStatusDie) NumberAvailable(v int32) *DaemonSetStatusDie {
	return d.DieStamp(func(r *appsv1.DaemonSetStatus) {
		r.NumberAvailable = v
	})
}

func (d *DaemonSetStatusDie) NumberUnavailable(v int32) *DaemonSetStatusDie {
	return d.DieStamp(func(r *appsv1.DaemonSetStatus) {
		r.NumberUnavailable = v
	})
}

func (d *DaemonSetStatusDie) CollisionCount(v *int32) *DaemonSetStatusDie {
	return d.DieStamp(func(r *appsv1.DaemonSetStatus) {
		r.CollisionCount = v
	})
}

func (d *DaemonSetStatusDie) Conditions(v ...appsv1.DaemonSetCondition) *DaemonSetStatusDie {
	return d.DieStamp(func(r *appsv1.DaemonSetStatus) {
		r.Conditions = v
	})
}

type DeploymentDie struct {
	metav1.FrozenObjectMeta
	mutable bool
	r       appsv1.Deployment
}

var DeploymentBlank = (&DeploymentDie{}).DieFeed(appsv1.Deployment{})

func (d *DeploymentDie) DieImmutable(immutable bool) *DeploymentDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

func (d *DeploymentDie) DieFeed(r appsv1.Deployment) *DeploymentDie {
	if d.mutable {
		d.FrozenObjectMeta = metav1.FreezeObjectMeta(r.ObjectMeta)
		d.r = r
		return d
	}
	return &DeploymentDie{
		FrozenObjectMeta: metav1.FreezeObjectMeta(r.ObjectMeta),
		mutable:          d.mutable,
		r:                r,
	}
}

func (d *DeploymentDie) DieRelease() appsv1.Deployment {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

func (d *DeploymentDie) DieStamp(fn func(r *appsv1.Deployment)) *DeploymentDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

func (d *DeploymentDie) DeepCopy() *DeploymentDie {
	r := *d.r.DeepCopy()
	return &DeploymentDie{
		FrozenObjectMeta: metav1.FreezeObjectMeta(r.ObjectMeta),
		mutable:          d.mutable,
		r:                r,
	}
}

func (d *DeploymentDie) DeepCopyObject() runtime.Object {
	return d.r.DeepCopy()
}

func (d *DeploymentDie) GetObjectKind() schema.ObjectKind {
	r := d.DieRelease()
	return r.GetObjectKind()
}

func (d *DeploymentDie) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.r)
}

func (d *DeploymentDie) UnmarshalJSON(b []byte) error {
	if d == DeploymentBlank {
		return fmtx.Errorf("cannot unmarshal into the root object, create a copy first")
	}
	r := &appsv1.Deployment{}
	err := json.Unmarshal(b, r)
	*d = *d.DieFeed(*r)
	return err
}

func (d *DeploymentDie) MetadataDie(fn func(d *metav1.ObjectMetaDie)) *DeploymentDie {
	return d.DieStamp(func(r *appsv1.Deployment) {
		d := metav1.ObjectMetaBlank.DieImmutable(false).DieFeed(r.ObjectMeta)
		fn(d)
		r.ObjectMeta = d.DieRelease()
	})
}

func (d *DeploymentDie) Spec(v appsv1.DeploymentSpec) *DeploymentDie {
	return d.DieStamp(func(r *appsv1.Deployment) {
		r.Spec = v
	})
}

func (d *DeploymentDie) SpecDie(fn func(d *DeploymentSpecDie)) *DeploymentDie {
	return d.DieStamp(func(r *appsv1.Deployment) {
		d := DeploymentSpecBlank.DieImmutable(false).DieFeed(r.Spec)
		fn(d)
		r.Spec = d.DieRelease()
	})
}

func (d *DeploymentDie) Status(v appsv1.DeploymentStatus) *DeploymentDie {
	return d.DieStamp(func(r *appsv1.Deployment) {
		r.Status = v
	})
}

func (d *DeploymentDie) StatusDie(fn func(d *DeploymentStatusDie)) *DeploymentDie {
	return d.DieStamp(func(r *appsv1.Deployment) {
		d := DeploymentStatusBlank.DieImmutable(false).DieFeed(r.Status)
		fn(d)
		r.Status = d.DieRelease()
	})
}

var _ apismetav1.Object = (*DeploymentDie)(nil)
var _ apismetav1.ObjectMetaAccessor = (*DeploymentDie)(nil)
var _ runtime.Object = (*DeploymentDie)(nil)

type DeploymentSpecDie struct {
	mutable bool
	r       appsv1.DeploymentSpec
}

var DeploymentSpecBlank = (&DeploymentSpecDie{}).DieFeed(appsv1.DeploymentSpec{})

func (d *DeploymentSpecDie) DieImmutable(immutable bool) *DeploymentSpecDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

func (d *DeploymentSpecDie) DieFeed(r appsv1.DeploymentSpec) *DeploymentSpecDie {
	if d.mutable {
		d.r = r
		return d
	}
	return &DeploymentSpecDie{
		mutable: d.mutable,
		r:       r,
	}
}

func (d *DeploymentSpecDie) DieRelease() appsv1.DeploymentSpec {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

func (d *DeploymentSpecDie) DieStamp(fn func(r *appsv1.DeploymentSpec)) *DeploymentSpecDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

func (d *DeploymentSpecDie) DeepCopy() *DeploymentSpecDie {
	r := *d.r.DeepCopy()
	return &DeploymentSpecDie{
		mutable: d.mutable,
		r:       r,
	}
}

func (d *DeploymentSpecDie) Replicas(v *int32) *DeploymentSpecDie {
	return d.DieStamp(func(r *appsv1.DeploymentSpec) {
		r.Replicas = v
	})
}

func (d *DeploymentSpecDie) Selector(v *apismetav1.LabelSelector) *DeploymentSpecDie {
	return d.DieStamp(func(r *appsv1.DeploymentSpec) {
		r.Selector = v
	})
}

func (d *DeploymentSpecDie) Template(v corev1.PodTemplateSpec) *DeploymentSpecDie {
	return d.DieStamp(func(r *appsv1.DeploymentSpec) {
		r.Template = v
	})
}

func (d *DeploymentSpecDie) Strategy(v appsv1.DeploymentStrategy) *DeploymentSpecDie {
	return d.DieStamp(func(r *appsv1.DeploymentSpec) {
		r.Strategy = v
	})
}

func (d *DeploymentSpecDie) MinReadySeconds(v int32) *DeploymentSpecDie {
	return d.DieStamp(func(r *appsv1.DeploymentSpec) {
		r.MinReadySeconds = v
	})
}

func (d *DeploymentSpecDie) RevisionHistoryLimit(v *int32) *DeploymentSpecDie {
	return d.DieStamp(func(r *appsv1.DeploymentSpec) {
		r.RevisionHistoryLimit = v
	})
}

func (d *DeploymentSpecDie) Paused(v bool) *DeploymentSpecDie {
	return d.DieStamp(func(r *appsv1.DeploymentSpec) {
		r.Paused = v
	})
}

func (d *DeploymentSpecDie) ProgressDeadlineSeconds(v *int32) *DeploymentSpecDie {
	return d.DieStamp(func(r *appsv1.DeploymentSpec) {
		r.ProgressDeadlineSeconds = v
	})
}

type DeploymentStatusDie struct {
	mutable bool
	r       appsv1.DeploymentStatus
}

var DeploymentStatusBlank = (&DeploymentStatusDie{}).DieFeed(appsv1.DeploymentStatus{})

func (d *DeploymentStatusDie) DieImmutable(immutable bool) *DeploymentStatusDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

func (d *DeploymentStatusDie) DieFeed(r appsv1.DeploymentStatus) *DeploymentStatusDie {
	if d.mutable {
		d.r = r
		return d
	}
	return &DeploymentStatusDie{
		mutable: d.mutable,
		r:       r,
	}
}

func (d *DeploymentStatusDie) DieRelease() appsv1.DeploymentStatus {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

func (d *DeploymentStatusDie) DieStamp(fn func(r *appsv1.DeploymentStatus)) *DeploymentStatusDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

func (d *DeploymentStatusDie) DeepCopy() *DeploymentStatusDie {
	r := *d.r.DeepCopy()
	return &DeploymentStatusDie{
		mutable: d.mutable,
		r:       r,
	}
}

func (d *DeploymentStatusDie) ObservedGeneration(v int64) *DeploymentStatusDie {
	return d.DieStamp(func(r *appsv1.DeploymentStatus) {
		r.ObservedGeneration = v
	})
}

func (d *DeploymentStatusDie) Replicas(v int32) *DeploymentStatusDie {
	return d.DieStamp(func(r *appsv1.DeploymentStatus) {
		r.Replicas = v
	})
}

func (d *DeploymentStatusDie) UpdatedReplicas(v int32) *DeploymentStatusDie {
	return d.DieStamp(func(r *appsv1.DeploymentStatus) {
		r.UpdatedReplicas = v
	})
}

func (d *DeploymentStatusDie) ReadyReplicas(v int32) *DeploymentStatusDie {
	return d.DieStamp(func(r *appsv1.DeploymentStatus) {
		r.ReadyReplicas = v
	})
}

func (d *DeploymentStatusDie) AvailableReplicas(v int32) *DeploymentStatusDie {
	return d.DieStamp(func(r *appsv1.DeploymentStatus) {
		r.AvailableReplicas = v
	})
}

func (d *DeploymentStatusDie) UnavailableReplicas(v int32) *DeploymentStatusDie {
	return d.DieStamp(func(r *appsv1.DeploymentStatus) {
		r.UnavailableReplicas = v
	})
}

func (d *DeploymentStatusDie) Conditions(v ...appsv1.DeploymentCondition) *DeploymentStatusDie {
	return d.DieStamp(func(r *appsv1.DeploymentStatus) {
		r.Conditions = v
	})
}

func (d *DeploymentStatusDie) CollisionCount(v *int32) *DeploymentStatusDie {
	return d.DieStamp(func(r *appsv1.DeploymentStatus) {
		r.CollisionCount = v
	})
}

type ReplicaSetDie struct {
	metav1.FrozenObjectMeta
	mutable bool
	r       appsv1.ReplicaSet
}

var ReplicaSetBlank = (&ReplicaSetDie{}).DieFeed(appsv1.ReplicaSet{})

func (d *ReplicaSetDie) DieImmutable(immutable bool) *ReplicaSetDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

func (d *ReplicaSetDie) DieFeed(r appsv1.ReplicaSet) *ReplicaSetDie {
	if d.mutable {
		d.FrozenObjectMeta = metav1.FreezeObjectMeta(r.ObjectMeta)
		d.r = r
		return d
	}
	return &ReplicaSetDie{
		FrozenObjectMeta: metav1.FreezeObjectMeta(r.ObjectMeta),
		mutable:          d.mutable,
		r:                r,
	}
}

func (d *ReplicaSetDie) DieRelease() appsv1.ReplicaSet {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

func (d *ReplicaSetDie) DieStamp(fn func(r *appsv1.ReplicaSet)) *ReplicaSetDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

func (d *ReplicaSetDie) DeepCopy() *ReplicaSetDie {
	r := *d.r.DeepCopy()
	return &ReplicaSetDie{
		FrozenObjectMeta: metav1.FreezeObjectMeta(r.ObjectMeta),
		mutable:          d.mutable,
		r:                r,
	}
}

func (d *ReplicaSetDie) DeepCopyObject() runtime.Object {
	return d.r.DeepCopy()
}

func (d *ReplicaSetDie) GetObjectKind() schema.ObjectKind {
	r := d.DieRelease()
	return r.GetObjectKind()
}

func (d *ReplicaSetDie) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.r)
}

func (d *ReplicaSetDie) UnmarshalJSON(b []byte) error {
	if d == ReplicaSetBlank {
		return fmtx.Errorf("cannot unmarshal into the root object, create a copy first")
	}
	r := &appsv1.ReplicaSet{}
	err := json.Unmarshal(b, r)
	*d = *d.DieFeed(*r)
	return err
}

func (d *ReplicaSetDie) MetadataDie(fn func(d *metav1.ObjectMetaDie)) *ReplicaSetDie {
	return d.DieStamp(func(r *appsv1.ReplicaSet) {
		d := metav1.ObjectMetaBlank.DieImmutable(false).DieFeed(r.ObjectMeta)
		fn(d)
		r.ObjectMeta = d.DieRelease()
	})
}

func (d *ReplicaSetDie) Spec(v appsv1.ReplicaSetSpec) *ReplicaSetDie {
	return d.DieStamp(func(r *appsv1.ReplicaSet) {
		r.Spec = v
	})
}

func (d *ReplicaSetDie) SpecDie(fn func(d *ReplicaSetSpecDie)) *ReplicaSetDie {
	return d.DieStamp(func(r *appsv1.ReplicaSet) {
		d := ReplicaSetSpecBlank.DieImmutable(false).DieFeed(r.Spec)
		fn(d)
		r.Spec = d.DieRelease()
	})
}

func (d *ReplicaSetDie) Status(v appsv1.ReplicaSetStatus) *ReplicaSetDie {
	return d.DieStamp(func(r *appsv1.ReplicaSet) {
		r.Status = v
	})
}

func (d *ReplicaSetDie) StatusDie(fn func(d *ReplicaSetStatusDie)) *ReplicaSetDie {
	return d.DieStamp(func(r *appsv1.ReplicaSet) {
		d := ReplicaSetStatusBlank.DieImmutable(false).DieFeed(r.Status)
		fn(d)
		r.Status = d.DieRelease()
	})
}

var _ apismetav1.Object = (*ReplicaSetDie)(nil)
var _ apismetav1.ObjectMetaAccessor = (*ReplicaSetDie)(nil)
var _ runtime.Object = (*ReplicaSetDie)(nil)

type ReplicaSetSpecDie struct {
	mutable bool
	r       appsv1.ReplicaSetSpec
}

var ReplicaSetSpecBlank = (&ReplicaSetSpecDie{}).DieFeed(appsv1.ReplicaSetSpec{})

func (d *ReplicaSetSpecDie) DieImmutable(immutable bool) *ReplicaSetSpecDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

func (d *ReplicaSetSpecDie) DieFeed(r appsv1.ReplicaSetSpec) *ReplicaSetSpecDie {
	if d.mutable {
		d.r = r
		return d
	}
	return &ReplicaSetSpecDie{
		mutable: d.mutable,
		r:       r,
	}
}

func (d *ReplicaSetSpecDie) DieRelease() appsv1.ReplicaSetSpec {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

func (d *ReplicaSetSpecDie) DieStamp(fn func(r *appsv1.ReplicaSetSpec)) *ReplicaSetSpecDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

func (d *ReplicaSetSpecDie) DeepCopy() *ReplicaSetSpecDie {
	r := *d.r.DeepCopy()
	return &ReplicaSetSpecDie{
		mutable: d.mutable,
		r:       r,
	}
}

func (d *ReplicaSetSpecDie) Replicas(v *int32) *ReplicaSetSpecDie {
	return d.DieStamp(func(r *appsv1.ReplicaSetSpec) {
		r.Replicas = v
	})
}

func (d *ReplicaSetSpecDie) MinReadySeconds(v int32) *ReplicaSetSpecDie {
	return d.DieStamp(func(r *appsv1.ReplicaSetSpec) {
		r.MinReadySeconds = v
	})
}

func (d *ReplicaSetSpecDie) Selector(v *apismetav1.LabelSelector) *ReplicaSetSpecDie {
	return d.DieStamp(func(r *appsv1.ReplicaSetSpec) {
		r.Selector = v
	})
}

func (d *ReplicaSetSpecDie) Template(v corev1.PodTemplateSpec) *ReplicaSetSpecDie {
	return d.DieStamp(func(r *appsv1.ReplicaSetSpec) {
		r.Template = v
	})
}

type ReplicaSetStatusDie struct {
	mutable bool
	r       appsv1.ReplicaSetStatus
}

var ReplicaSetStatusBlank = (&ReplicaSetStatusDie{}).DieFeed(appsv1.ReplicaSetStatus{})

func (d *ReplicaSetStatusDie) DieImmutable(immutable bool) *ReplicaSetStatusDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

func (d *ReplicaSetStatusDie) DieFeed(r appsv1.ReplicaSetStatus) *ReplicaSetStatusDie {
	if d.mutable {
		d.r = r
		return d
	}
	return &ReplicaSetStatusDie{
		mutable: d.mutable,
		r:       r,
	}
}

func (d *ReplicaSetStatusDie) DieRelease() appsv1.ReplicaSetStatus {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

func (d *ReplicaSetStatusDie) DieStamp(fn func(r *appsv1.ReplicaSetStatus)) *ReplicaSetStatusDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

func (d *ReplicaSetStatusDie) DeepCopy() *ReplicaSetStatusDie {
	r := *d.r.DeepCopy()
	return &ReplicaSetStatusDie{
		mutable: d.mutable,
		r:       r,
	}
}

func (d *ReplicaSetStatusDie) Replicas(v int32) *ReplicaSetStatusDie {
	return d.DieStamp(func(r *appsv1.ReplicaSetStatus) {
		r.Replicas = v
	})
}

func (d *ReplicaSetStatusDie) FullyLabeledReplicas(v int32) *ReplicaSetStatusDie {
	return d.DieStamp(func(r *appsv1.ReplicaSetStatus) {
		r.FullyLabeledReplicas = v
	})
}

func (d *ReplicaSetStatusDie) ReadyReplicas(v int32) *ReplicaSetStatusDie {
	return d.DieStamp(func(r *appsv1.ReplicaSetStatus) {
		r.ReadyReplicas = v
	})
}

func (d *ReplicaSetStatusDie) AvailableReplicas(v int32) *ReplicaSetStatusDie {
	return d.DieStamp(func(r *appsv1.ReplicaSetStatus) {
		r.AvailableReplicas = v
	})
}

func (d *ReplicaSetStatusDie) ObservedGeneration(v int64) *ReplicaSetStatusDie {
	return d.DieStamp(func(r *appsv1.ReplicaSetStatus) {
		r.ObservedGeneration = v
	})
}

func (d *ReplicaSetStatusDie) Conditions(v ...appsv1.ReplicaSetCondition) *ReplicaSetStatusDie {
	return d.DieStamp(func(r *appsv1.ReplicaSetStatus) {
		r.Conditions = v
	})
}

type StatefulSetDie struct {
	metav1.FrozenObjectMeta
	mutable bool
	r       appsv1.StatefulSet
}

var StatefulSetBlank = (&StatefulSetDie{}).DieFeed(appsv1.StatefulSet{})

func (d *StatefulSetDie) DieImmutable(immutable bool) *StatefulSetDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

func (d *StatefulSetDie) DieFeed(r appsv1.StatefulSet) *StatefulSetDie {
	if d.mutable {
		d.FrozenObjectMeta = metav1.FreezeObjectMeta(r.ObjectMeta)
		d.r = r
		return d
	}
	return &StatefulSetDie{
		FrozenObjectMeta: metav1.FreezeObjectMeta(r.ObjectMeta),
		mutable:          d.mutable,
		r:                r,
	}
}

func (d *StatefulSetDie) DieRelease() appsv1.StatefulSet {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

func (d *StatefulSetDie) DieStamp(fn func(r *appsv1.StatefulSet)) *StatefulSetDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

func (d *StatefulSetDie) DeepCopy() *StatefulSetDie {
	r := *d.r.DeepCopy()
	return &StatefulSetDie{
		FrozenObjectMeta: metav1.FreezeObjectMeta(r.ObjectMeta),
		mutable:          d.mutable,
		r:                r,
	}
}

func (d *StatefulSetDie) DeepCopyObject() runtime.Object {
	return d.r.DeepCopy()
}

func (d *StatefulSetDie) GetObjectKind() schema.ObjectKind {
	r := d.DieRelease()
	return r.GetObjectKind()
}

func (d *StatefulSetDie) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.r)
}

func (d *StatefulSetDie) UnmarshalJSON(b []byte) error {
	if d == StatefulSetBlank {
		return fmtx.Errorf("cannot unmarshal into the root object, create a copy first")
	}
	r := &appsv1.StatefulSet{}
	err := json.Unmarshal(b, r)
	*d = *d.DieFeed(*r)
	return err
}

func (d *StatefulSetDie) MetadataDie(fn func(d *metav1.ObjectMetaDie)) *StatefulSetDie {
	return d.DieStamp(func(r *appsv1.StatefulSet) {
		d := metav1.ObjectMetaBlank.DieImmutable(false).DieFeed(r.ObjectMeta)
		fn(d)
		r.ObjectMeta = d.DieRelease()
	})
}

func (d *StatefulSetDie) Spec(v appsv1.StatefulSetSpec) *StatefulSetDie {
	return d.DieStamp(func(r *appsv1.StatefulSet) {
		r.Spec = v
	})
}

func (d *StatefulSetDie) SpecDie(fn func(d *StatefulSetSpecDie)) *StatefulSetDie {
	return d.DieStamp(func(r *appsv1.StatefulSet) {
		d := StatefulSetSpecBlank.DieImmutable(false).DieFeed(r.Spec)
		fn(d)
		r.Spec = d.DieRelease()
	})
}

func (d *StatefulSetDie) Status(v appsv1.StatefulSetStatus) *StatefulSetDie {
	return d.DieStamp(func(r *appsv1.StatefulSet) {
		r.Status = v
	})
}

func (d *StatefulSetDie) StatusDie(fn func(d *StatefulSetStatusDie)) *StatefulSetDie {
	return d.DieStamp(func(r *appsv1.StatefulSet) {
		d := StatefulSetStatusBlank.DieImmutable(false).DieFeed(r.Status)
		fn(d)
		r.Status = d.DieRelease()
	})
}

var _ apismetav1.Object = (*StatefulSetDie)(nil)
var _ apismetav1.ObjectMetaAccessor = (*StatefulSetDie)(nil)
var _ runtime.Object = (*StatefulSetDie)(nil)

type StatefulSetSpecDie struct {
	mutable bool
	r       appsv1.StatefulSetSpec
}

var StatefulSetSpecBlank = (&StatefulSetSpecDie{}).DieFeed(appsv1.StatefulSetSpec{})

func (d *StatefulSetSpecDie) DieImmutable(immutable bool) *StatefulSetSpecDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

func (d *StatefulSetSpecDie) DieFeed(r appsv1.StatefulSetSpec) *StatefulSetSpecDie {
	if d.mutable {
		d.r = r
		return d
	}
	return &StatefulSetSpecDie{
		mutable: d.mutable,
		r:       r,
	}
}

func (d *StatefulSetSpecDie) DieRelease() appsv1.StatefulSetSpec {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

func (d *StatefulSetSpecDie) DieStamp(fn func(r *appsv1.StatefulSetSpec)) *StatefulSetSpecDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

func (d *StatefulSetSpecDie) DeepCopy() *StatefulSetSpecDie {
	r := *d.r.DeepCopy()
	return &StatefulSetSpecDie{
		mutable: d.mutable,
		r:       r,
	}
}

func (d *StatefulSetSpecDie) Replicas(v *int32) *StatefulSetSpecDie {
	return d.DieStamp(func(r *appsv1.StatefulSetSpec) {
		r.Replicas = v
	})
}

func (d *StatefulSetSpecDie) Selector(v *apismetav1.LabelSelector) *StatefulSetSpecDie {
	return d.DieStamp(func(r *appsv1.StatefulSetSpec) {
		r.Selector = v
	})
}

func (d *StatefulSetSpecDie) Template(v corev1.PodTemplateSpec) *StatefulSetSpecDie {
	return d.DieStamp(func(r *appsv1.StatefulSetSpec) {
		r.Template = v
	})
}

func (d *StatefulSetSpecDie) VolumeClaimTemplates(v ...corev1.PersistentVolumeClaim) *StatefulSetSpecDie {
	return d.DieStamp(func(r *appsv1.StatefulSetSpec) {
		r.VolumeClaimTemplates = v
	})
}

func (d *StatefulSetSpecDie) ServiceName(v string) *StatefulSetSpecDie {
	return d.DieStamp(func(r *appsv1.StatefulSetSpec) {
		r.ServiceName = v
	})
}

func (d *StatefulSetSpecDie) PodManagementPolicy(v appsv1.PodManagementPolicyType) *StatefulSetSpecDie {
	return d.DieStamp(func(r *appsv1.StatefulSetSpec) {
		r.PodManagementPolicy = v
	})
}

func (d *StatefulSetSpecDie) UpdateStrategy(v appsv1.StatefulSetUpdateStrategy) *StatefulSetSpecDie {
	return d.DieStamp(func(r *appsv1.StatefulSetSpec) {
		r.UpdateStrategy = v
	})
}

func (d *StatefulSetSpecDie) RevisionHistoryLimit(v *int32) *StatefulSetSpecDie {
	return d.DieStamp(func(r *appsv1.StatefulSetSpec) {
		r.RevisionHistoryLimit = v
	})
}

func (d *StatefulSetSpecDie) MinReadySeconds(v int32) *StatefulSetSpecDie {
	return d.DieStamp(func(r *appsv1.StatefulSetSpec) {
		r.MinReadySeconds = v
	})
}

type StatefulSetStatusDie struct {
	mutable bool
	r       appsv1.StatefulSetStatus
}

var StatefulSetStatusBlank = (&StatefulSetStatusDie{}).DieFeed(appsv1.StatefulSetStatus{})

func (d *StatefulSetStatusDie) DieImmutable(immutable bool) *StatefulSetStatusDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

func (d *StatefulSetStatusDie) DieFeed(r appsv1.StatefulSetStatus) *StatefulSetStatusDie {
	if d.mutable {
		d.r = r
		return d
	}
	return &StatefulSetStatusDie{
		mutable: d.mutable,
		r:       r,
	}
}

func (d *StatefulSetStatusDie) DieRelease() appsv1.StatefulSetStatus {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

func (d *StatefulSetStatusDie) DieStamp(fn func(r *appsv1.StatefulSetStatus)) *StatefulSetStatusDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

func (d *StatefulSetStatusDie) DeepCopy() *StatefulSetStatusDie {
	r := *d.r.DeepCopy()
	return &StatefulSetStatusDie{
		mutable: d.mutable,
		r:       r,
	}
}

func (d *StatefulSetStatusDie) ObservedGeneration(v int64) *StatefulSetStatusDie {
	return d.DieStamp(func(r *appsv1.StatefulSetStatus) {
		r.ObservedGeneration = v
	})
}

func (d *StatefulSetStatusDie) Replicas(v int32) *StatefulSetStatusDie {
	return d.DieStamp(func(r *appsv1.StatefulSetStatus) {
		r.Replicas = v
	})
}

func (d *StatefulSetStatusDie) ReadyReplicas(v int32) *StatefulSetStatusDie {
	return d.DieStamp(func(r *appsv1.StatefulSetStatus) {
		r.ReadyReplicas = v
	})
}

func (d *StatefulSetStatusDie) CurrentReplicas(v int32) *StatefulSetStatusDie {
	return d.DieStamp(func(r *appsv1.StatefulSetStatus) {
		r.CurrentReplicas = v
	})
}

func (d *StatefulSetStatusDie) UpdatedReplicas(v int32) *StatefulSetStatusDie {
	return d.DieStamp(func(r *appsv1.StatefulSetStatus) {
		r.UpdatedReplicas = v
	})
}

func (d *StatefulSetStatusDie) CurrentRevision(v string) *StatefulSetStatusDie {
	return d.DieStamp(func(r *appsv1.StatefulSetStatus) {
		r.CurrentRevision = v
	})
}

func (d *StatefulSetStatusDie) UpdateRevision(v string) *StatefulSetStatusDie {
	return d.DieStamp(func(r *appsv1.StatefulSetStatus) {
		r.UpdateRevision = v
	})
}

func (d *StatefulSetStatusDie) CollisionCount(v *int32) *StatefulSetStatusDie {
	return d.DieStamp(func(r *appsv1.StatefulSetStatus) {
		r.CollisionCount = v
	})
}

func (d *StatefulSetStatusDie) Conditions(v ...appsv1.StatefulSetCondition) *StatefulSetStatusDie {
	return d.DieStamp(func(r *appsv1.StatefulSetStatus) {
		r.Conditions = v
	})
}

func (d *StatefulSetStatusDie) AvailableReplicas(v int32) *StatefulSetStatusDie {
	return d.DieStamp(func(r *appsv1.StatefulSetStatus) {
		r.AvailableReplicas = v
	})
}
