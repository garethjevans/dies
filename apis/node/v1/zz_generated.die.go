//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021-2022 the original author or authors.

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
	metav1 "dies.dev/apis/meta/v1"
	json "encoding/json"
	fmtx "fmt"
	corev1 "k8s.io/api/core/v1"
	nodev1 "k8s.io/api/node/v1"
	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	jsonpath "k8s.io/client-go/util/jsonpath"
	osx "os"
	reflectx "reflect"
	yaml "sigs.k8s.io/yaml"
)

var RuntimeClassBlank = (&RuntimeClassDie{}).DieFeed(nodev1.RuntimeClass{})

type RuntimeClassDie struct {
	metav1.FrozenObjectMeta
	mutable bool
	r       nodev1.RuntimeClass
}

// DieImmutable returns a new die for the current die's state that is either mutable (`false`) or immutable (`true`).
func (d *RuntimeClassDie) DieImmutable(immutable bool) *RuntimeClassDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

// DieFeed returns a new die with the provided resource.
func (d *RuntimeClassDie) DieFeed(r nodev1.RuntimeClass) *RuntimeClassDie {
	if d.mutable {
		d.FrozenObjectMeta = metav1.FreezeObjectMeta(r.ObjectMeta)
		d.r = r
		return d
	}
	return &RuntimeClassDie{
		FrozenObjectMeta: metav1.FreezeObjectMeta(r.ObjectMeta),
		mutable:          d.mutable,
		r:                r,
	}
}

// DieFeedPtr returns a new die with the provided resource pointer. If the resource is nil, the empty value is used instead.
func (d *RuntimeClassDie) DieFeedPtr(r *nodev1.RuntimeClass) *RuntimeClassDie {
	if r == nil {
		r = &nodev1.RuntimeClass{}
	}
	return d.DieFeed(*r)
}

// DieFeedJSON returns a new die with the provided JSON. Panics on error.
func (d *RuntimeClassDie) DieFeedJSON(j []byte) *RuntimeClassDie {
	r := nodev1.RuntimeClass{}
	if err := json.Unmarshal(j, &r); err != nil {
		panic(err)
	}
	return d.DieFeed(r)
}

// DieFeedYAML returns a new die with the provided YAML. Panics on error.
func (d *RuntimeClassDie) DieFeedYAML(y []byte) *RuntimeClassDie {
	r := nodev1.RuntimeClass{}
	if err := yaml.Unmarshal(y, &r); err != nil {
		panic(err)
	}
	return d.DieFeed(r)
}

// DieFeedYAMLFile returns a new die loading YAML from a file path. Panics on error.
func (d *RuntimeClassDie) DieFeedYAMLFile(name string) *RuntimeClassDie {
	y, err := osx.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return d.DieFeedYAML(y)
}

// DieFeedRawExtension returns the resource managed by the die as an raw extension. Panics on error.
func (d *RuntimeClassDie) DieFeedRawExtension(raw runtime.RawExtension) *RuntimeClassDie {
	j, err := json.Marshal(raw)
	if err != nil {
		panic(err)
	}
	return d.DieFeedJSON(j)
}

// DieRelease returns the resource managed by the die.
func (d *RuntimeClassDie) DieRelease() nodev1.RuntimeClass {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

// DieReleasePtr returns a pointer to the resource managed by the die.
func (d *RuntimeClassDie) DieReleasePtr() *nodev1.RuntimeClass {
	r := d.DieRelease()
	return &r
}

// DieReleaseUnstructured returns the resource managed by the die as an unstructured object. Panics on error.
func (d *RuntimeClassDie) DieReleaseUnstructured() *unstructured.Unstructured {
	r := d.DieReleasePtr()
	u, err := runtime.DefaultUnstructuredConverter.ToUnstructured(r)
	if err != nil {
		panic(err)
	}
	return &unstructured.Unstructured{
		Object: u,
	}
}

// DieReleaseJSON returns the resource managed by the die as JSON. Panics on error.
func (d *RuntimeClassDie) DieReleaseJSON() []byte {
	r := d.DieReleasePtr()
	j, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return j
}

// DieReleaseYAML returns the resource managed by the die as YAML. Panics on error.
func (d *RuntimeClassDie) DieReleaseYAML() []byte {
	r := d.DieReleasePtr()
	y, err := yaml.Marshal(r)
	if err != nil {
		panic(err)
	}
	return y
}

// DieReleaseRawExtension returns the resource managed by the die as an raw extension. Panics on error.
func (d *RuntimeClassDie) DieReleaseRawExtension() runtime.RawExtension {
	j := d.DieReleaseJSON()
	raw := runtime.RawExtension{}
	if err := json.Unmarshal(j, &raw); err != nil {
		panic(err)
	}
	return raw
}

// DieStamp returns a new die with the resource passed to the callback function. The resource is mutable.
func (d *RuntimeClassDie) DieStamp(fn func(r *nodev1.RuntimeClass)) *RuntimeClassDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

// Experimental: DieStampAt uses a JSON path (http://goessner.net/articles/JsonPath/) expression to stamp portions of the resource. The callback is invoked with each JSON path match. Panics if the callback function does not accept a single argument of the same type or a pointer to that type as found on the resource at the target location.
//
// Future iterations will improve type coercion from the resource to the callback argument.
func (d *RuntimeClassDie) DieStampAt(jp string, fn interface{}) *RuntimeClassDie {
	return d.DieStamp(func(r *nodev1.RuntimeClass) {
		if ni := reflectx.ValueOf(fn).Type().NumIn(); ni != 1 {
			panic(fmtx.Errorf("callback function must have 1 input parameters, found %d", ni))
		}
		if no := reflectx.ValueOf(fn).Type().NumOut(); no != 0 {
			panic(fmtx.Errorf("callback function must have 0 output parameters, found %d", no))
		}

		cp := jsonpath.New("")
		if err := cp.Parse(fmtx.Sprintf("{%s}", jp)); err != nil {
			panic(err)
		}
		cr, err := cp.FindResults(r)
		if err != nil {
			// errors are expected if a path is not found
			return
		}
		for _, cv := range cr[0] {
			arg0t := reflectx.ValueOf(fn).Type().In(0)

			var args []reflectx.Value
			if cv.Type().AssignableTo(arg0t) {
				args = []reflectx.Value{cv}
			} else if cv.CanAddr() && cv.Addr().Type().AssignableTo(arg0t) {
				args = []reflectx.Value{cv.Addr()}
			} else {
				panic(fmtx.Errorf("callback function must accept value of type %q, found type %q", cv.Type(), arg0t))
			}

			reflectx.ValueOf(fn).Call(args)
		}
	})
}

// DieWith returns a new die after passing the current die to the callback function. The passed die is mutable.
func (d *RuntimeClassDie) DieWith(fns ...func(d *RuntimeClassDie)) *RuntimeClassDie {
	nd := RuntimeClassBlank.DieFeed(d.DieRelease()).DieImmutable(false)
	for _, fn := range fns {
		if fn != nil {
			fn(nd)
		}
	}
	return d.DieFeed(nd.DieRelease())
}

// DeepCopy returns a new die with equivalent state. Useful for snapshotting a mutable die.
func (d *RuntimeClassDie) DeepCopy() *RuntimeClassDie {
	r := *d.r.DeepCopy()
	return &RuntimeClassDie{
		FrozenObjectMeta: metav1.FreezeObjectMeta(r.ObjectMeta),
		mutable:          d.mutable,
		r:                r,
	}
}

var _ runtime.Object = (*RuntimeClassDie)(nil)

func (d *RuntimeClassDie) DeepCopyObject() runtime.Object {
	return d.r.DeepCopy()
}

func (d *RuntimeClassDie) GetObjectKind() schema.ObjectKind {
	r := d.DieRelease()
	return r.GetObjectKind()
}

func (d *RuntimeClassDie) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.r)
}

func (d *RuntimeClassDie) UnmarshalJSON(b []byte) error {
	if d == RuntimeClassBlank {
		return fmtx.Errorf("cannot unmarshal into the blank die, create a copy first")
	}
	if !d.mutable {
		return fmtx.Errorf("cannot unmarshal into immutable dies, create a mutable version first")
	}
	r := &nodev1.RuntimeClass{}
	err := json.Unmarshal(b, r)
	*d = *d.DieFeed(*r)
	return err
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (d *RuntimeClassDie) APIVersion(v string) *RuntimeClassDie {
	return d.DieStamp(func(r *nodev1.RuntimeClass) {
		r.APIVersion = v
	})
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (d *RuntimeClassDie) Kind(v string) *RuntimeClassDie {
	return d.DieStamp(func(r *nodev1.RuntimeClass) {
		r.Kind = v
	})
}

// MetadataDie stamps the resource's ObjectMeta field with a mutable die.
func (d *RuntimeClassDie) MetadataDie(fn func(d *metav1.ObjectMetaDie)) *RuntimeClassDie {
	return d.DieStamp(func(r *nodev1.RuntimeClass) {
		d := metav1.ObjectMetaBlank.DieImmutable(false).DieFeed(r.ObjectMeta)
		fn(d)
		r.ObjectMeta = d.DieRelease()
	})
}

// handler specifies the underlying runtime and configuration that the CRI implementation will use to handle pods of this class. The possible values are specific to the node & CRI configuration.  It is assumed that all handlers are available on every node, and handlers of the same name are equivalent on every node. For example, a handler called "runc" might specify that the runc OCI runtime (using native Linux containers) will be used to run the containers in a pod. The Handler must be lowercase, conform to the DNS Label (RFC 1123) requirements, and is immutable.
func (d *RuntimeClassDie) Handler(v string) *RuntimeClassDie {
	return d.DieStamp(func(r *nodev1.RuntimeClass) {
		r.Handler = v
	})
}

// overhead represents the resource overhead associated with running a pod for a given RuntimeClass. For more details, see https://kubernetes.io/docs/concepts/scheduling-eviction/pod-overhead/
func (d *RuntimeClassDie) Overhead(v *nodev1.Overhead) *RuntimeClassDie {
	return d.DieStamp(func(r *nodev1.RuntimeClass) {
		r.Overhead = v
	})
}

// scheduling holds the scheduling constraints to ensure that pods running with this RuntimeClass are scheduled to nodes that support it. If scheduling is nil, this RuntimeClass is assumed to be supported by all nodes.
func (d *RuntimeClassDie) Scheduling(v *nodev1.Scheduling) *RuntimeClassDie {
	return d.DieStamp(func(r *nodev1.RuntimeClass) {
		r.Scheduling = v
	})
}

var OverheadBlank = (&OverheadDie{}).DieFeed(nodev1.Overhead{})

type OverheadDie struct {
	mutable bool
	r       nodev1.Overhead
}

// DieImmutable returns a new die for the current die's state that is either mutable (`false`) or immutable (`true`).
func (d *OverheadDie) DieImmutable(immutable bool) *OverheadDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

// DieFeed returns a new die with the provided resource.
func (d *OverheadDie) DieFeed(r nodev1.Overhead) *OverheadDie {
	if d.mutable {
		d.r = r
		return d
	}
	return &OverheadDie{
		mutable: d.mutable,
		r:       r,
	}
}

// DieFeedPtr returns a new die with the provided resource pointer. If the resource is nil, the empty value is used instead.
func (d *OverheadDie) DieFeedPtr(r *nodev1.Overhead) *OverheadDie {
	if r == nil {
		r = &nodev1.Overhead{}
	}
	return d.DieFeed(*r)
}

// DieFeedJSON returns a new die with the provided JSON. Panics on error.
func (d *OverheadDie) DieFeedJSON(j []byte) *OverheadDie {
	r := nodev1.Overhead{}
	if err := json.Unmarshal(j, &r); err != nil {
		panic(err)
	}
	return d.DieFeed(r)
}

// DieFeedYAML returns a new die with the provided YAML. Panics on error.
func (d *OverheadDie) DieFeedYAML(y []byte) *OverheadDie {
	r := nodev1.Overhead{}
	if err := yaml.Unmarshal(y, &r); err != nil {
		panic(err)
	}
	return d.DieFeed(r)
}

// DieFeedYAMLFile returns a new die loading YAML from a file path. Panics on error.
func (d *OverheadDie) DieFeedYAMLFile(name string) *OverheadDie {
	y, err := osx.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return d.DieFeedYAML(y)
}

// DieFeedRawExtension returns the resource managed by the die as an raw extension. Panics on error.
func (d *OverheadDie) DieFeedRawExtension(raw runtime.RawExtension) *OverheadDie {
	j, err := json.Marshal(raw)
	if err != nil {
		panic(err)
	}
	return d.DieFeedJSON(j)
}

// DieRelease returns the resource managed by the die.
func (d *OverheadDie) DieRelease() nodev1.Overhead {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

// DieReleasePtr returns a pointer to the resource managed by the die.
func (d *OverheadDie) DieReleasePtr() *nodev1.Overhead {
	r := d.DieRelease()
	return &r
}

// DieReleaseJSON returns the resource managed by the die as JSON. Panics on error.
func (d *OverheadDie) DieReleaseJSON() []byte {
	r := d.DieReleasePtr()
	j, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return j
}

// DieReleaseYAML returns the resource managed by the die as YAML. Panics on error.
func (d *OverheadDie) DieReleaseYAML() []byte {
	r := d.DieReleasePtr()
	y, err := yaml.Marshal(r)
	if err != nil {
		panic(err)
	}
	return y
}

// DieReleaseRawExtension returns the resource managed by the die as an raw extension. Panics on error.
func (d *OverheadDie) DieReleaseRawExtension() runtime.RawExtension {
	j := d.DieReleaseJSON()
	raw := runtime.RawExtension{}
	if err := json.Unmarshal(j, &raw); err != nil {
		panic(err)
	}
	return raw
}

// DieStamp returns a new die with the resource passed to the callback function. The resource is mutable.
func (d *OverheadDie) DieStamp(fn func(r *nodev1.Overhead)) *OverheadDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

// Experimental: DieStampAt uses a JSON path (http://goessner.net/articles/JsonPath/) expression to stamp portions of the resource. The callback is invoked with each JSON path match. Panics if the callback function does not accept a single argument of the same type or a pointer to that type as found on the resource at the target location.
//
// Future iterations will improve type coercion from the resource to the callback argument.
func (d *OverheadDie) DieStampAt(jp string, fn interface{}) *OverheadDie {
	return d.DieStamp(func(r *nodev1.Overhead) {
		if ni := reflectx.ValueOf(fn).Type().NumIn(); ni != 1 {
			panic(fmtx.Errorf("callback function must have 1 input parameters, found %d", ni))
		}
		if no := reflectx.ValueOf(fn).Type().NumOut(); no != 0 {
			panic(fmtx.Errorf("callback function must have 0 output parameters, found %d", no))
		}

		cp := jsonpath.New("")
		if err := cp.Parse(fmtx.Sprintf("{%s}", jp)); err != nil {
			panic(err)
		}
		cr, err := cp.FindResults(r)
		if err != nil {
			// errors are expected if a path is not found
			return
		}
		for _, cv := range cr[0] {
			arg0t := reflectx.ValueOf(fn).Type().In(0)

			var args []reflectx.Value
			if cv.Type().AssignableTo(arg0t) {
				args = []reflectx.Value{cv}
			} else if cv.CanAddr() && cv.Addr().Type().AssignableTo(arg0t) {
				args = []reflectx.Value{cv.Addr()}
			} else {
				panic(fmtx.Errorf("callback function must accept value of type %q, found type %q", cv.Type(), arg0t))
			}

			reflectx.ValueOf(fn).Call(args)
		}
	})
}

// DieWith returns a new die after passing the current die to the callback function. The passed die is mutable.
func (d *OverheadDie) DieWith(fns ...func(d *OverheadDie)) *OverheadDie {
	nd := OverheadBlank.DieFeed(d.DieRelease()).DieImmutable(false)
	for _, fn := range fns {
		if fn != nil {
			fn(nd)
		}
	}
	return d.DieFeed(nd.DieRelease())
}

// DeepCopy returns a new die with equivalent state. Useful for snapshotting a mutable die.
func (d *OverheadDie) DeepCopy() *OverheadDie {
	r := *d.r.DeepCopy()
	return &OverheadDie{
		mutable: d.mutable,
		r:       r,
	}
}

// podFixed represents the fixed resource overhead associated with running a pod.
func (d *OverheadDie) PodFixed(v corev1.ResourceList) *OverheadDie {
	return d.DieStamp(func(r *nodev1.Overhead) {
		r.PodFixed = v
	})
}

var SchedulingBlank = (&SchedulingDie{}).DieFeed(nodev1.Scheduling{})

type SchedulingDie struct {
	mutable bool
	r       nodev1.Scheduling
}

// DieImmutable returns a new die for the current die's state that is either mutable (`false`) or immutable (`true`).
func (d *SchedulingDie) DieImmutable(immutable bool) *SchedulingDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

// DieFeed returns a new die with the provided resource.
func (d *SchedulingDie) DieFeed(r nodev1.Scheduling) *SchedulingDie {
	if d.mutable {
		d.r = r
		return d
	}
	return &SchedulingDie{
		mutable: d.mutable,
		r:       r,
	}
}

// DieFeedPtr returns a new die with the provided resource pointer. If the resource is nil, the empty value is used instead.
func (d *SchedulingDie) DieFeedPtr(r *nodev1.Scheduling) *SchedulingDie {
	if r == nil {
		r = &nodev1.Scheduling{}
	}
	return d.DieFeed(*r)
}

// DieFeedJSON returns a new die with the provided JSON. Panics on error.
func (d *SchedulingDie) DieFeedJSON(j []byte) *SchedulingDie {
	r := nodev1.Scheduling{}
	if err := json.Unmarshal(j, &r); err != nil {
		panic(err)
	}
	return d.DieFeed(r)
}

// DieFeedYAML returns a new die with the provided YAML. Panics on error.
func (d *SchedulingDie) DieFeedYAML(y []byte) *SchedulingDie {
	r := nodev1.Scheduling{}
	if err := yaml.Unmarshal(y, &r); err != nil {
		panic(err)
	}
	return d.DieFeed(r)
}

// DieFeedYAMLFile returns a new die loading YAML from a file path. Panics on error.
func (d *SchedulingDie) DieFeedYAMLFile(name string) *SchedulingDie {
	y, err := osx.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return d.DieFeedYAML(y)
}

// DieFeedRawExtension returns the resource managed by the die as an raw extension. Panics on error.
func (d *SchedulingDie) DieFeedRawExtension(raw runtime.RawExtension) *SchedulingDie {
	j, err := json.Marshal(raw)
	if err != nil {
		panic(err)
	}
	return d.DieFeedJSON(j)
}

// DieRelease returns the resource managed by the die.
func (d *SchedulingDie) DieRelease() nodev1.Scheduling {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

// DieReleasePtr returns a pointer to the resource managed by the die.
func (d *SchedulingDie) DieReleasePtr() *nodev1.Scheduling {
	r := d.DieRelease()
	return &r
}

// DieReleaseJSON returns the resource managed by the die as JSON. Panics on error.
func (d *SchedulingDie) DieReleaseJSON() []byte {
	r := d.DieReleasePtr()
	j, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return j
}

// DieReleaseYAML returns the resource managed by the die as YAML. Panics on error.
func (d *SchedulingDie) DieReleaseYAML() []byte {
	r := d.DieReleasePtr()
	y, err := yaml.Marshal(r)
	if err != nil {
		panic(err)
	}
	return y
}

// DieReleaseRawExtension returns the resource managed by the die as an raw extension. Panics on error.
func (d *SchedulingDie) DieReleaseRawExtension() runtime.RawExtension {
	j := d.DieReleaseJSON()
	raw := runtime.RawExtension{}
	if err := json.Unmarshal(j, &raw); err != nil {
		panic(err)
	}
	return raw
}

// DieStamp returns a new die with the resource passed to the callback function. The resource is mutable.
func (d *SchedulingDie) DieStamp(fn func(r *nodev1.Scheduling)) *SchedulingDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

// Experimental: DieStampAt uses a JSON path (http://goessner.net/articles/JsonPath/) expression to stamp portions of the resource. The callback is invoked with each JSON path match. Panics if the callback function does not accept a single argument of the same type or a pointer to that type as found on the resource at the target location.
//
// Future iterations will improve type coercion from the resource to the callback argument.
func (d *SchedulingDie) DieStampAt(jp string, fn interface{}) *SchedulingDie {
	return d.DieStamp(func(r *nodev1.Scheduling) {
		if ni := reflectx.ValueOf(fn).Type().NumIn(); ni != 1 {
			panic(fmtx.Errorf("callback function must have 1 input parameters, found %d", ni))
		}
		if no := reflectx.ValueOf(fn).Type().NumOut(); no != 0 {
			panic(fmtx.Errorf("callback function must have 0 output parameters, found %d", no))
		}

		cp := jsonpath.New("")
		if err := cp.Parse(fmtx.Sprintf("{%s}", jp)); err != nil {
			panic(err)
		}
		cr, err := cp.FindResults(r)
		if err != nil {
			// errors are expected if a path is not found
			return
		}
		for _, cv := range cr[0] {
			arg0t := reflectx.ValueOf(fn).Type().In(0)

			var args []reflectx.Value
			if cv.Type().AssignableTo(arg0t) {
				args = []reflectx.Value{cv}
			} else if cv.CanAddr() && cv.Addr().Type().AssignableTo(arg0t) {
				args = []reflectx.Value{cv.Addr()}
			} else {
				panic(fmtx.Errorf("callback function must accept value of type %q, found type %q", cv.Type(), arg0t))
			}

			reflectx.ValueOf(fn).Call(args)
		}
	})
}

// DieWith returns a new die after passing the current die to the callback function. The passed die is mutable.
func (d *SchedulingDie) DieWith(fns ...func(d *SchedulingDie)) *SchedulingDie {
	nd := SchedulingBlank.DieFeed(d.DieRelease()).DieImmutable(false)
	for _, fn := range fns {
		if fn != nil {
			fn(nd)
		}
	}
	return d.DieFeed(nd.DieRelease())
}

// DeepCopy returns a new die with equivalent state. Useful for snapshotting a mutable die.
func (d *SchedulingDie) DeepCopy() *SchedulingDie {
	r := *d.r.DeepCopy()
	return &SchedulingDie{
		mutable: d.mutable,
		r:       r,
	}
}

// nodeSelector lists labels that must be present on nodes that support this RuntimeClass. Pods using this RuntimeClass can only be scheduled to a node matched by this selector. The RuntimeClass nodeSelector is merged with a pod's existing nodeSelector. Any conflicts will cause the pod to be rejected in admission.
func (d *SchedulingDie) NodeSelector(v map[string]string) *SchedulingDie {
	return d.DieStamp(func(r *nodev1.Scheduling) {
		r.NodeSelector = v
	})
}

// tolerations are appended (excluding duplicates) to pods running with this RuntimeClass during admission, effectively unioning the set of nodes tolerated by the pod and the RuntimeClass.
func (d *SchedulingDie) Tolerations(v ...corev1.Toleration) *SchedulingDie {
	return d.DieStamp(func(r *nodev1.Scheduling) {
		r.Tolerations = v
	})
}
