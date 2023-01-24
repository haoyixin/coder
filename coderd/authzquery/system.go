package authzquery

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/coder/coder/coderd/database"
	"github.com/coder/coder/coderd/rbac"
)

// TODO: @emyrk should we name system functions differently to indicate a user
// cannot call them? Maybe we should have a separate interface for system functions?
// So you'd do `authzQ.System().GetDERPMeshKey(ctx)` or something like that?
// Cian: yes. Let's do it.

func (q *AuthzQuerier) GetLatestWorkspaceBuilds(ctx context.Context) ([]database.WorkspaceBuild, error) {
	// This function is a system function until we implement a join for workspace builds.
	// This is because we need to query for all related workspaces to the returned builds.
	// This is a very inefficient method of fetching the latest workspace builds.
	// We should just join the rbac properties.
	return q.database.GetLatestWorkspaceBuilds(ctx)
}

// GetWorkspaceAgentByAuthToken is used in http middleware to get the workspace agent.
// This should only be used by a system user in that middleware.
func (q *AuthzQuerier) GetWorkspaceAgentByAuthToken(ctx context.Context, authToken uuid.UUID) (database.WorkspaceAgent, error) {
	return q.GetWorkspaceAgentByAuthToken(ctx, authToken)
}

func (q *AuthzQuerier) GetActiveUserCount(ctx context.Context) (int64, error) {
	return q.GetActiveUserCount(ctx)
}

func (q *AuthzQuerier) GetAuthorizationUserRoles(ctx context.Context, userID uuid.UUID) (database.GetAuthorizationUserRolesRow, error) {
	return q.GetAuthorizationUserRoles(ctx, userID)
}

func (q *AuthzQuerier) GetDERPMeshKey(ctx context.Context) (string, error) {
	// TODO Implement authz check for system user.
	return q.database.GetDERPMeshKey(ctx)
}

func (q *AuthzQuerier) InsertDERPMeshKey(ctx context.Context, value string) error {
	// TODO Implement authz check for system user.
	return q.InsertDERPMeshKey(ctx, value)
}

func (q *AuthzQuerier) InsertDeploymentID(ctx context.Context, value string) error {
	// TODO Implement authz check for system user.
	return q.InsertDeploymentID(ctx, value)
}

func (q *AuthzQuerier) InsertReplica(ctx context.Context, arg database.InsertReplicaParams) (database.Replica, error) {
	// TODO Implement authz check for system user.
	return q.InsertReplica(ctx, arg)
}

func (q *AuthzQuerier) UpdateReplica(ctx context.Context, arg database.UpdateReplicaParams) (database.Replica, error) {
	// TODO Implement authz check for system user.
	return q.UpdateReplica(ctx, arg)
}

func (q *AuthzQuerier) DeleteReplicasUpdatedBefore(ctx context.Context, updatedAt time.Time) error {
	// TODO Implement authz check for system user.
	return q.DeleteReplicasUpdatedBefore(ctx, updatedAt)
}

func (q *AuthzQuerier) GetReplicasUpdatedAfter(ctx context.Context, updatedAt time.Time) ([]database.Replica, error) {
	// TODO Implement authz check for system user.
	return q.GetReplicasUpdatedAfter(ctx, updatedAt)
}

func (q *AuthzQuerier) GetTemplates(ctx context.Context, arg database.GetTemplatesWithFilterParams, prepared rbac.PreparedAuthorized) ([]database.Template, error) {
	// TODO Implement authz check for system user.
	return q.GetTemplates(ctx, arg, prepared)
}

// UpdateWorkspaceBuildCostByID is used by the provisioning system to update the cost of a workspace build.
func (q *AuthzQuerier) UpdateWorkspaceBuildCostByID(ctx context.Context, arg database.UpdateWorkspaceBuildCostByIDParams) (database.WorkspaceBuild, error) {
	return q.UpdateWorkspaceBuildCostByID(ctx, arg)
}

// Telemetry related functions. These functions are system functions for returning
// telemetry data. Never called by a user.

func (q *AuthzQuerier) GetWorkspaceBuildsCreatedAfter(ctx context.Context, createdAt time.Time) ([]database.WorkspaceBuild, error) {
	return q.GetWorkspaceBuildsCreatedAfter(ctx, createdAt)
}

func (q *AuthzQuerier) GetWorkspaceAgentsCreatedAfter(ctx context.Context, createdAt time.Time) ([]database.WorkspaceAgent, error) {
	return q.GetWorkspaceAgentsCreatedAfter(ctx, createdAt)
}

func (q *AuthzQuerier) GetWorkspaceAppsCreatedAfter(ctx context.Context, createdAt time.Time) ([]database.WorkspaceApp, error) {
	return q.GetWorkspaceAppsCreatedAfter(ctx, createdAt)
}

func (q *AuthzQuerier) GetWorkspaceResourcesCreatedAfter(ctx context.Context, createdAt time.Time) ([]database.WorkspaceResource, error) {
	return q.GetWorkspaceResourcesCreatedAfter(ctx, createdAt)
}

func (q *AuthzQuerier) GetWorkspaceResourceMetadataCreatedAfter(ctx context.Context, createdAt time.Time) ([]database.WorkspaceResourceMetadatum, error) {
	return q.database.GetWorkspaceResourceMetadataCreatedAfter(ctx, createdAt)
}

// Provisionerd server functions

func (q *AuthzQuerier) InsertWorkspaceAgent(ctx context.Context, arg database.InsertWorkspaceAgentParams) (database.WorkspaceAgent, error) {
	return q.InsertWorkspaceAgent(ctx, arg)
}

func (q *AuthzQuerier) InsertWorkspaceApp(ctx context.Context, arg database.InsertWorkspaceAppParams) (database.WorkspaceApp, error) {
	return q.InsertWorkspaceApp(ctx, arg)
}

func (q *AuthzQuerier) InsertWorkspaceResourceMetadata(ctx context.Context, arg database.InsertWorkspaceResourceMetadataParams) ([]database.WorkspaceResourceMetadatum, error) {
	return q.InsertWorkspaceResourceMetadata(ctx, arg)
}
