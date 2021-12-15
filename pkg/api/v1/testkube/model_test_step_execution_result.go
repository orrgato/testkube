/*
 * TestKube API
 *
 * TestKube provides a Kubernetes-native framework for test definition, execution and results
 *
 * API version: 1.0.0
 * Contact: testkube@kubeshop.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package testkube

// execution result returned from executor
type TestStepExecutionResult struct {
	Script *ObjectRef       `json:"script,omitempty"`
	Result *ExecutionResult `json:"result,omitempty"`
}