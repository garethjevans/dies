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
	testing "dies.dev/testing"
	testingx "testing"
)

func TestHorizontalPodAutoscalerDie_MissingMethods(t *testingx.T) {
	die := HorizontalPodAutoscalerBlank
	ignore := []string{"TypeMeta", "ObjectMeta"}
	diff := testing.DieFieldDiff(die).Delete(ignore...)
	if diff.Len() != 0 {
		t.Errorf("found missing fields for HorizontalPodAutoscalerDie: %s", diff.List())
	}
}

func TestHorizontalPodAutoscalerSpecDie_MissingMethods(t *testingx.T) {
	die := HorizontalPodAutoscalerSpecBlank
	ignore := []string{}
	diff := testing.DieFieldDiff(die).Delete(ignore...)
	if diff.Len() != 0 {
		t.Errorf("found missing fields for HorizontalPodAutoscalerSpecDie: %s", diff.List())
	}
}

func TestCrossVersionObjectReferenceDie_MissingMethods(t *testingx.T) {
	die := CrossVersionObjectReferenceBlank
	ignore := []string{}
	diff := testing.DieFieldDiff(die).Delete(ignore...)
	if diff.Len() != 0 {
		t.Errorf("found missing fields for CrossVersionObjectReferenceDie: %s", diff.List())
	}
}

func TestHorizontalPodAutoscalerStatusDie_MissingMethods(t *testingx.T) {
	die := HorizontalPodAutoscalerStatusBlank
	ignore := []string{}
	diff := testing.DieFieldDiff(die).Delete(ignore...)
	if diff.Len() != 0 {
		t.Errorf("found missing fields for HorizontalPodAutoscalerStatusDie: %s", diff.List())
	}
}
