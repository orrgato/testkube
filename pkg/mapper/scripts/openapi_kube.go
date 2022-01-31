package scripts

import (
	scriptsv2 "github.com/kubeshop/testkube-operator/apis/script/v2"
	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func MapScriptToScriptSpec(request testkube.ScriptUpsertRequest) *scriptsv2.Script {

	script := &scriptsv2.Script{
		ObjectMeta: metav1.ObjectMeta{
			Name:      request.Name,
			Namespace: request.Namespace,
		},
		Spec: scriptsv2.ScriptSpec{
			Type_:   request.Type_,
			Content: MapScriptContentToScriptSpecContent(request.Content),
			Tags:    request.Tags,
		},
	}

	return script

}

func MapScriptContentToScriptSpecContent(content *testkube.ScriptContent) (specContent *scriptsv2.ScriptContent) {
	if content == nil {
		return
	}

	return &scriptsv2.ScriptContent{
		// assuming same data structure
		Repository: (*scriptsv2.Repository)(content.Repository),
		Content:    content.Data,
		Uri:        content.Uri,
		Type_:      content.Type_,
	}
}