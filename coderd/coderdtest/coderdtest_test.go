package coderdtest_test

import (
	"testing"

	"go.uber.org/goleak"

	"github.com/coder/coder/coderd/coderdtest"
	"github.com/coder/coder/codersdk"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func TestNew(t *testing.T) {
	t.Parallel()
	client := coderdtest.New(t, nil)
	user := coderdtest.CreateFirstUser(t, client)
	closer := coderdtest.NewProvisionerDaemon(t, client)
	version := coderdtest.CreateTemplateVersion(t, client, user.OrganizationID, nil)
	coderdtest.AwaitTemplateVersionJob(t, client, version.ID)
	template := coderdtest.CreateTemplate(t, client, user.OrganizationID, version.ID)
	workspace := coderdtest.CreateWorkspace(t, client, codersdk.Me, template.ID)
	coderdtest.AwaitWorkspaceBuildJob(t, client, workspace.LatestBuild.ID)
	coderdtest.AwaitWorkspaceAgents(t, client, workspace.LatestBuild.ID)
	_, _ = coderdtest.NewGoogleInstanceIdentity(t, "example", false)
	_, _ = coderdtest.NewAWSInstanceIdentity(t, "an-instance")
	closer.Close()
}
