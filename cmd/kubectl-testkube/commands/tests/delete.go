package tests

import (
	"github.com/kubeshop/testkube/cmd/kubectl-testkube/commands/common"
	"github.com/kubeshop/testkube/cmd/kubectl-testkube/commands/common/validator"
	"github.com/kubeshop/testkube/pkg/ui"
	"github.com/spf13/cobra"
)

func NewDeleteTestsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete <testName>",
		Short: "Delete scripts",
		Args:  validator.TestName,
		Run: func(cmd *cobra.Command, args []string) {
			ui.Logo()

			client, _ := common.GetClient(cmd)
			namespace := cmd.Flag("namespace").Value.String()
			name := args[0]

			err := client.DeleteTest(name, namespace)
			ui.ExitOnError("delete script "+name+" from namespace "+namespace, err)

			ui.Success("Succesfully deleted", name)
		},
	}

	return cmd
}
