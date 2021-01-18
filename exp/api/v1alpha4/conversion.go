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

// Hub marks AWSMachinePool as a conversion hub.
func (*AWSMachinePool) Hub() {}

// Hub marks AWSMachinePoolList as a conversion hub.
func (*AWSMachinePoolList) Hub() {}

// Hub marks AWSManagedMachinePool as a conversion hub.
func (*AWSManagedMachinePool) Hub() {}

// Hub marks AWSManagedMachinePoolList as a conversion hub.
func (*AWSManagedMachinePoolList) Hub() {}

// Hub marks AWSManagedCluster as a conversion hub.
func (*AWSManagedCluster) Hub() {}

// Hub marks AWSManagedClusterList as a conversion hub.
func (*AWSManagedClusterList) Hub() {}
