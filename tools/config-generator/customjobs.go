/*
Copyright 2020 The Knative Authors

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

// Although custom jobs are not generated by this generator, certain testgrid
// configs are needed for certain custom jobs

package main

var (
	customJobnames = []string{
		"ci-knative-backup-artifacts",
		"ci-knative-cleanup",
		"ci-knative-flakes-reporter",
		"ci-knative-flakes-resultsrecorder",
		"ci-knative-prow-jobs-syncer",
		"post-knative-test-infra-image-push",
		"post-knative-sandbox-peribolos",
		"post-knative-test-infra-deploy-tools",
	}
)

func addCustomJobsTestgrid() {
	var (
		extras = map[string]string{
			"num_failures_to_alert": "1",
			"alert_options":         "\n      alert_mail_to_addresses: \"serverless-engprod-sea@google.com\"",
		}
	)
	for _, job := range customJobnames {
		metaData.AddNonAlignedTest(NonAlignedTestGroup{
			DashboardGroup: "maintenance",
			DashboardName:  "utilities",
			HumanTabName:   job,
			CIJobName:      job,
			Extra:          extras,
		})
	}
}
