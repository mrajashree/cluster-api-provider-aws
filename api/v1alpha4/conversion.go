/*
Copyright 2020 The Kubernetes Authors.
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

package v1alpha4

// Hub marks AWSCluster as a conversion hub.
func (*AWSCluster) Hub() {}

// Hub marks AWSClusterList as a conversion hub.
func (*AWSClusterList) Hub() {}

// Hub marks AWSMachine as a conversion hub.
func (*AWSMachine) Hub() {}

// Hub marks AWSMachineList as a conversion hub.
func (*AWSMachineList) Hub() {}

// Hub marks AWSMachineTemplate as a conversion hub.
func (*AWSMachineTemplate) Hub() {}

// Hub marks AWSMachineTemplateList as a conversion hub.
func (*AWSMachineTemplateList) Hub() {}
