/*
Copyright 2021 Flant JSC

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

package hooks

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/deckhouse/deckhouse/testing/hooks"
)

var _ = Describe("Deckhouse hooks :: enable_extended_monitoring ::", func() {
	const (
		kubeSystemNS = `
---
apiVersion: v1
kind: Namespace
metadata:
  name: kube-system
  annotations:
  extended-monitoring.deckhouse.io: ""
`
		d8SystemNS = `
---
apiVersion: v1
kind: Namespace
metadata:
  name: d8-system
  labels:
    extended-monitoring.deckhouse.io/enabled: ""
`
	)

	f := HookExecutionConfigInit("", "")

	When("ns/kube-system and ns/d8-system are present", func() {
		BeforeEach(func() {
			f.BindingContexts.Set(f.KubeStateSet(kubeSystemNS + d8SystemNS))
			f.RunHook()
		})

		It("Label should be present on both namespaces", func() {
			Expect(f).To(ExecuteSuccessfully())

			label := f.KubernetesGlobalResource("Namespace", "d8-system").
				Field(`metadata.labels.extended-monitoring\.deckhouse\.io/enabled`)
			Expect(label.Exists()).To(BeTrue())
			Expect(label.String()).To(Equal(""))
			Expect(f.KubernetesGlobalResource("Namespace", "d8-system").
				Field(`metadata.annotations.extended-monitoring\.flant\.com/enabled`).Exists()).To(BeFalse())

			label = f.KubernetesGlobalResource("Namespace", "kube-system").
				Field(`metadata.labels.extended-monitoring\.deckhouse\.io/enabled`)
			Expect(label.Exists()).To(BeTrue())
			Expect(label.String()).To(Equal(""))
			Expect(f.KubernetesGlobalResource("Namespace", "kube-system").
				Field(`metadata.annotations.extended-monitoring\.flant\.com/enabled`).Exists()).To(BeFalse())
		})
	})
})
