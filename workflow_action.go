package main

import "gopkg.in/yaml.v3"

const (
	BranchProtectionConfiguration WorkflowAction = iota
	BranchProtectionRule
	CheckRun
	CheckSuite
	CodeScanningAlert
	CommitComment
	Create
	CustomProperty
	CustomPropertyValues
	Delete
	DependabotAlert
	DeployKey
	Deployment
	DeploymentProtectionRule
	DeploymentReview
	DeploymentStatus
	Discussion
	DiscussionComment
	Fork
	GithubAppAuthorization
	Gollum
	Installation
	InstallationRepositories
	InstallationTarget
	IssueComment
	Issues
	Label
	MarketplacePurchase
	Member
	Membership
	MergeGroup
	Meta
	Milestone
	OrgBlock
	Organization
	Package
	PageBuild
	PersonalAccessTokenRequest
	Ping
	ProjectCard
	Project
	ProjectColumn
	ProjectsV2
	ProjectsV2Item
	ProjectsV2StatusUpdate
	Public
	PullRequest
	PullRequestReviewComment
	PullRequestReview
	PullRequestReviewThread
	Push
	RegistryPackage
	Release
	RepositoryAdvisory
	Repository
	RepositoryDispatch
	RepositoryImport
	RepositoryRuleset
	RepositoryVulnerabilityAlert
	SecretScanningAlert
	SecretScanningAlertLocation
	SecurityAdvisory
	SecurityAndAnalysis
	Sponsorship
	Star
	Status
	SubIssues
	TeamAdd
	Team
	Watch
	WorkflowDispatch
	WorkflowJob
	WorkflowRun
)

type WorkflowAction uint8

var (
	actionStrToEnum = map[string]WorkflowAction{
		"branch_protection_configuration": BranchProtectionConfiguration,
		"branch_protection_rule":          BranchProtectionRule,
		"check_run":                       CheckRun,
		"check_suite":                     CheckSuite,
		"code_scanning_alert":             CodeScanningAlert,
		"commit_comment":                  CommitComment,
		"create":                          Create,
		"custom_property":                 CustomProperty,
		"custom_property_values":          CustomPropertyValues,
		"delete":                          Delete,
		"dependabot_alert":                DependabotAlert,
		"deploy_key":                      DeployKey,
		"deployment":                      Deployment,
		"deployment_protection_rule":      DeploymentProtectionRule,
		"deployment_review":               DeploymentReview,
		"deployment_status":               DeploymentStatus,
		"discussion":                      Discussion,
		"discussion_comment":              DiscussionComment,
		"fork":                            Fork,
		"github_app_authorization":        GithubAppAuthorization,
		"gollum":                          Gollum,
		"installation":                    Installation,
		"installation_repositories":       InstallationRepositories,
		"installation_target":             InstallationTarget,
		"issue_comment":                   IssueComment,
		"issues":                          Issues,
		"label":                           Label,
		"marketplace_purchase":            MarketplacePurchase,
		"member":                          Member,
		"membership":                      Membership,
		"merge_group":                     MergeGroup,
		"meta":                            Meta,
		"milestone":                       Milestone,
		"org_block":                       OrgBlock,
		"organization":                    Organization,
		"package":                         Package,
		"page_build":                      PageBuild,
		"personal_access_token_request":   PersonalAccessTokenRequest,
		"ping":                            Ping,
		"project_card":                    ProjectCard,
		"project":                         Project,
		"project_column":                  ProjectColumn,
		"projects_v2":                     ProjectsV2,
		"projects_v2_item":                ProjectsV2Item,
		"projects_v2_status_update":       ProjectsV2StatusUpdate,
		"public":                          Public,
		"pull_request":                    PullRequest,
		"pull_request_review_comment":     PullRequestReviewComment,
		"pull_request_review":             PullRequestReview,
		"pull_request_review_thread":      PullRequestReviewThread,
		"push":                            Push,
		"registry_package":                RegistryPackage,
		"release":                         Release,
		"repository_advisory":             RepositoryAdvisory,
		"repository":                      Repository,
		"repository_dispatch":             RepositoryDispatch,
		"repository_import":               RepositoryImport,
		"repository_ruleset":              RepositoryRuleset,
		"repository_vulnerability_alert":  RepositoryVulnerabilityAlert,
		"secret_scanning_alert":           SecretScanningAlert,
		"secret_scanning_alert_location":  SecretScanningAlertLocation,
		"security_advisory":               SecurityAdvisory,
		"security_and_analysis":           SecurityAndAnalysis,
		"sponsorship":                     Sponsorship,
		"star":                            Star,
		"status":                          Status,
		"sub_issues":                      SubIssues,
		"team_add":                        TeamAdd,
		"team":                            Team,
		"watch":                           Watch,
		"workflow_dispatch":               WorkflowDispatch,
		"workflow_job":                    WorkflowJob,
		"workflow_run":                    WorkflowRun,
	}
)

func FromActionString(action string) WorkflowAction {
	return actionStrToEnum[action]
}

func (action WorkflowAction) ToActionString() (string, bool) {
	for k, v := range actionStrToEnum {
		if v == action {
			return k, true
		}
	}
	return "", false
}

func (action WorkflowAction) String() string {
	res, _ := action.ToActionString()
	return res
}

func (action WorkflowAction) GoString() string {
	return action.String()
}

func (a WorkflowAction) MarshalYAML() (interface{}, error) {
	value, _ := a.ToActionString()
	return yaml.Marshal(value)
}

func (a *WorkflowAction) UnmarshalYAML(value *yaml.Node) error {
	var actionStr string
	err := value.Decode(&actionStr)
	if err != nil {
		return err
	}

	*a = FromActionString(actionStr)
	return nil
}
