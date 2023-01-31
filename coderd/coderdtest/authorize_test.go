package coderdtest_test

import (
	"context"
	"testing"

	"github.com/coder/coder/coderd/coderdtest"
)

func TestAuthorizeAllEndpoints(t *testing.T) {
	t.Parallel()
	// TODO: DO NOT MERGE THIS
	t.Skip("TODO: fix all the unit tests that break when this is enabled. ")
	client, _, api := coderdtest.NewWithAPI(t, &coderdtest.Options{
		// Required for any subdomain-based proxy tests to pass.
		AppHostname: "*.test.coder.com",
		Authorizer: &coderdtest.RecordingAuthorizer{
			Wrapped: &coderdtest.FakeAuthorizer{AlwaysReturn: nil},
		},
		IncludeProvisionerDaemon: true,
	})
	admin := coderdtest.CreateFirstUser(t, client)
	a := coderdtest.NewAuthTester(context.Background(), t, client, api, admin)
	skipRoute, assertRoute := coderdtest.AGPLRoutes(a)
	a.Test(context.Background(), assertRoute, skipRoute)
}
