// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package database

import (
	"context"

	"github.com/google/uuid"
)

type querier interface {
	// Acquires the lock for a single job that isn't started, completed,
	// canceled, and that matches an array of provisioner types.
	//
	// SKIP LOCKED is used to jump over locked rows. This prevents
	// multiple provisioners from acquiring the same jobs. See:
	// https://www.postgresql.org/docs/9.5/sql-select.html#SQL-FOR-UPDATE-SHARE
	AcquireProvisionerJob(ctx context.Context, arg AcquireProvisionerJobParams) (ProvisionerJob, error)
	DeleteAPIKeyByID(ctx context.Context, id string) error
	DeleteGitSSHKey(ctx context.Context, userID uuid.UUID) error
	DeleteParameterValueByID(ctx context.Context, id uuid.UUID) error
	GetAPIKeyByID(ctx context.Context, id string) (APIKey, error)
	// GetAuditLogsBefore retrieves `limit` number of audit logs before the provided
	// ID.
	GetAuditLogsBefore(ctx context.Context, arg GetAuditLogsBeforeParams) ([]AuditLog, error)
	// This function
	GetAuthorizationUserRoles(ctx context.Context, userID uuid.UUID) (GetAuthorizationUserRolesRow, error)
	GetFileByHash(ctx context.Context, hash string) (File, error)
	GetGitSSHKey(ctx context.Context, userID uuid.UUID) (GitSSHKey, error)
	GetLatestWorkspaceBuildByWorkspaceID(ctx context.Context, workspaceID uuid.UUID) (WorkspaceBuild, error)
	GetLatestWorkspaceBuildsByWorkspaceIDs(ctx context.Context, ids []uuid.UUID) ([]WorkspaceBuild, error)
	GetOrganizationByID(ctx context.Context, id uuid.UUID) (Organization, error)
	GetOrganizationByName(ctx context.Context, name string) (Organization, error)
	GetOrganizationIDsByMemberIDs(ctx context.Context, ids []uuid.UUID) ([]GetOrganizationIDsByMemberIDsRow, error)
	GetOrganizationMemberByUserID(ctx context.Context, arg GetOrganizationMemberByUserIDParams) (OrganizationMember, error)
	GetOrganizationMembershipsByUserID(ctx context.Context, userID uuid.UUID) ([]OrganizationMember, error)
	GetOrganizations(ctx context.Context) ([]Organization, error)
	GetOrganizationsByUserID(ctx context.Context, userID uuid.UUID) ([]Organization, error)
	GetParameterSchemasByJobID(ctx context.Context, jobID uuid.UUID) ([]ParameterSchema, error)
	GetParameterValueByScopeAndName(ctx context.Context, arg GetParameterValueByScopeAndNameParams) (ParameterValue, error)
	GetParameterValuesByScope(ctx context.Context, arg GetParameterValuesByScopeParams) ([]ParameterValue, error)
	GetProvisionerDaemonByID(ctx context.Context, id uuid.UUID) (ProvisionerDaemon, error)
	GetProvisionerDaemons(ctx context.Context) ([]ProvisionerDaemon, error)
	GetProvisionerJobByID(ctx context.Context, id uuid.UUID) (ProvisionerJob, error)
	GetProvisionerJobsByIDs(ctx context.Context, ids []uuid.UUID) ([]ProvisionerJob, error)
	GetProvisionerLogsByIDBetween(ctx context.Context, arg GetProvisionerLogsByIDBetweenParams) ([]ProvisionerJobLog, error)
	GetTemplateByID(ctx context.Context, id uuid.UUID) (Template, error)
	GetTemplateByOrganizationAndName(ctx context.Context, arg GetTemplateByOrganizationAndNameParams) (Template, error)
	GetTemplateVersionByID(ctx context.Context, id uuid.UUID) (TemplateVersion, error)
	GetTemplateVersionByJobID(ctx context.Context, jobID uuid.UUID) (TemplateVersion, error)
	GetTemplateVersionByTemplateIDAndName(ctx context.Context, arg GetTemplateVersionByTemplateIDAndNameParams) (TemplateVersion, error)
	GetTemplateVersionsByTemplateID(ctx context.Context, arg GetTemplateVersionsByTemplateIDParams) ([]TemplateVersion, error)
	GetTemplatesByIDs(ctx context.Context, ids []uuid.UUID) ([]Template, error)
	GetTemplatesByOrganization(ctx context.Context, arg GetTemplatesByOrganizationParams) ([]Template, error)
	GetUserByEmailOrUsername(ctx context.Context, arg GetUserByEmailOrUsernameParams) (User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (User, error)
	GetUserCount(ctx context.Context) (int64, error)
	GetUsers(ctx context.Context, arg GetUsersParams) ([]User, error)
	GetUsersByIDs(ctx context.Context, ids []uuid.UUID) ([]User, error)
	GetWorkspaceAgentByAuthToken(ctx context.Context, authToken uuid.UUID) (WorkspaceAgent, error)
	GetWorkspaceAgentByID(ctx context.Context, id uuid.UUID) (WorkspaceAgent, error)
	GetWorkspaceAgentByInstanceID(ctx context.Context, authInstanceID string) (WorkspaceAgent, error)
	GetWorkspaceAgentsByResourceIDs(ctx context.Context, ids []uuid.UUID) ([]WorkspaceAgent, error)
	GetWorkspaceBuildByID(ctx context.Context, id uuid.UUID) (WorkspaceBuild, error)
	GetWorkspaceBuildByJobID(ctx context.Context, jobID uuid.UUID) (WorkspaceBuild, error)
	GetWorkspaceBuildByWorkspaceID(ctx context.Context, arg GetWorkspaceBuildByWorkspaceIDParams) ([]WorkspaceBuild, error)
	GetWorkspaceBuildByWorkspaceIDAndName(ctx context.Context, arg GetWorkspaceBuildByWorkspaceIDAndNameParams) (WorkspaceBuild, error)
	GetWorkspaceByID(ctx context.Context, id uuid.UUID) (Workspace, error)
	GetWorkspaceByOwnerIDAndName(ctx context.Context, arg GetWorkspaceByOwnerIDAndNameParams) (Workspace, error)
	GetWorkspaceOwnerCountsByTemplateIDs(ctx context.Context, ids []uuid.UUID) ([]GetWorkspaceOwnerCountsByTemplateIDsRow, error)
	GetWorkspaceResourceByID(ctx context.Context, id uuid.UUID) (WorkspaceResource, error)
	GetWorkspaceResourcesByJobID(ctx context.Context, jobID uuid.UUID) ([]WorkspaceResource, error)
	GetWorkspacesAutostart(ctx context.Context) ([]Workspace, error)
	GetWorkspacesByOrganizationIDs(ctx context.Context, arg GetWorkspacesByOrganizationIDsParams) ([]Workspace, error)
	GetWorkspacesByTemplateID(ctx context.Context, arg GetWorkspacesByTemplateIDParams) ([]Workspace, error)
	GetWorkspacesWithFilter(ctx context.Context, arg GetWorkspacesWithFilterParams) ([]Workspace, error)
	InsertAPIKey(ctx context.Context, arg InsertAPIKeyParams) (APIKey, error)
	InsertAuditLog(ctx context.Context, arg InsertAuditLogParams) (AuditLog, error)
	InsertFile(ctx context.Context, arg InsertFileParams) (File, error)
	InsertGitSSHKey(ctx context.Context, arg InsertGitSSHKeyParams) (GitSSHKey, error)
	InsertOrganization(ctx context.Context, arg InsertOrganizationParams) (Organization, error)
	InsertOrganizationMember(ctx context.Context, arg InsertOrganizationMemberParams) (OrganizationMember, error)
	InsertParameterSchema(ctx context.Context, arg InsertParameterSchemaParams) (ParameterSchema, error)
	InsertParameterValue(ctx context.Context, arg InsertParameterValueParams) (ParameterValue, error)
	InsertProvisionerDaemon(ctx context.Context, arg InsertProvisionerDaemonParams) (ProvisionerDaemon, error)
	InsertProvisionerJob(ctx context.Context, arg InsertProvisionerJobParams) (ProvisionerJob, error)
	InsertProvisionerJobLogs(ctx context.Context, arg InsertProvisionerJobLogsParams) ([]ProvisionerJobLog, error)
	InsertTemplate(ctx context.Context, arg InsertTemplateParams) (Template, error)
	InsertTemplateVersion(ctx context.Context, arg InsertTemplateVersionParams) (TemplateVersion, error)
	InsertUser(ctx context.Context, arg InsertUserParams) (User, error)
	InsertWorkspace(ctx context.Context, arg InsertWorkspaceParams) (Workspace, error)
	InsertWorkspaceAgent(ctx context.Context, arg InsertWorkspaceAgentParams) (WorkspaceAgent, error)
	InsertWorkspaceBuild(ctx context.Context, arg InsertWorkspaceBuildParams) (WorkspaceBuild, error)
	InsertWorkspaceResource(ctx context.Context, arg InsertWorkspaceResourceParams) (WorkspaceResource, error)
	UpdateAPIKeyByID(ctx context.Context, arg UpdateAPIKeyByIDParams) error
	UpdateGitSSHKey(ctx context.Context, arg UpdateGitSSHKeyParams) error
	UpdateMemberRoles(ctx context.Context, arg UpdateMemberRolesParams) (OrganizationMember, error)
	UpdateProvisionerDaemonByID(ctx context.Context, arg UpdateProvisionerDaemonByIDParams) error
	UpdateProvisionerJobByID(ctx context.Context, arg UpdateProvisionerJobByIDParams) error
	UpdateProvisionerJobWithCancelByID(ctx context.Context, arg UpdateProvisionerJobWithCancelByIDParams) error
	UpdateProvisionerJobWithCompleteByID(ctx context.Context, arg UpdateProvisionerJobWithCompleteByIDParams) error
	UpdateTemplateActiveVersionByID(ctx context.Context, arg UpdateTemplateActiveVersionByIDParams) error
	UpdateTemplateDeletedByID(ctx context.Context, arg UpdateTemplateDeletedByIDParams) error
	UpdateTemplateVersionByID(ctx context.Context, arg UpdateTemplateVersionByIDParams) error
	UpdateTemplateVersionDescriptionByJobID(ctx context.Context, arg UpdateTemplateVersionDescriptionByJobIDParams) error
	UpdateUserHashedPassword(ctx context.Context, arg UpdateUserHashedPasswordParams) error
	UpdateUserProfile(ctx context.Context, arg UpdateUserProfileParams) (User, error)
	UpdateUserRoles(ctx context.Context, arg UpdateUserRolesParams) (User, error)
	UpdateUserStatus(ctx context.Context, arg UpdateUserStatusParams) (User, error)
	UpdateWorkspaceAgentConnectionByID(ctx context.Context, arg UpdateWorkspaceAgentConnectionByIDParams) error
	UpdateWorkspaceAutostart(ctx context.Context, arg UpdateWorkspaceAutostartParams) error
	UpdateWorkspaceBuildByID(ctx context.Context, arg UpdateWorkspaceBuildByIDParams) error
	UpdateWorkspaceDeletedByID(ctx context.Context, arg UpdateWorkspaceDeletedByIDParams) error
	UpdateWorkspaceTTL(ctx context.Context, arg UpdateWorkspaceTTLParams) error
}

var _ querier = (*sqlQuerier)(nil)
