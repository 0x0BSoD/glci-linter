package gitlab

import (
	"time"
)

type APILintRequest struct {
	Content string `json:"content"`
}

type APILintResponse struct {
	Valid      bool     `json:"valid,omitempty"`
	Warnings   []string `json:"warnings,omitempty"`
	Errors     []string `json:"errors,omitempty"`
	MergedYaml string   `json:"merged_yaml,omitempty"`
}

type Project struct {
	ID                int       `json:"id"`
	Description       string    `json:"description"`
	Name              string    `json:"name"`
	NameWithNamespace string    `json:"name_with_namespace"`
	Path              string    `json:"path"`
	PathWithNamespace string    `json:"path_with_namespace"`
	CreatedAt         time.Time `json:"created_at"`
	DefaultBranch     string    `json:"default_branch"`
	TagList           []any     `json:"tag_list"`
	Topics            []any     `json:"topics"`
	SSHURLToRepo      string    `json:"ssh_url_to_repo"`
	HTTPURLToRepo     string    `json:"http_url_to_repo"`
	WebURL            string    `json:"web_url"`
	ReadmeURL         string    `json:"readme_url"`
	ForksCount        int       `json:"forks_count"`
	AvatarURL         any       `json:"avatar_url"`
	StarCount         int       `json:"star_count"`
	LastActivityAt    time.Time `json:"last_activity_at"`
	Namespace         struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Path      string `json:"path"`
		Kind      string `json:"kind"`
		FullPath  string `json:"full_path"`
		ParentID  int    `json:"parent_id"`
		AvatarURL any    `json:"avatar_url"`
		WebURL    string `json:"web_url"`
	} `json:"namespace"`
	ContainerRegistryImagePrefix string `json:"container_registry_image_prefix"`
	Links                        struct {
		Self          string `json:"self"`
		Issues        string `json:"issues"`
		MergeRequests string `json:"merge_requests"`
		RepoBranches  string `json:"repo_branches"`
		Labels        string `json:"labels"`
		Events        string `json:"events"`
		Members       string `json:"members"`
		ClusterAgents string `json:"cluster_agents"`
	} `json:"_links"`
	PackagesEnabled                bool   `json:"packages_enabled"`
	EmptyRepo                      bool   `json:"empty_repo"`
	Archived                       bool   `json:"archived"`
	Visibility                     string `json:"visibility"`
	ResolveOutdatedDiffDiscussions bool   `json:"resolve_outdated_diff_discussions"`
	ContainerExpirationPolicy      struct {
		Cadence       string    `json:"cadence"`
		Enabled       bool      `json:"enabled"`
		KeepN         int       `json:"keep_n"`
		OlderThan     string    `json:"older_than"`
		NameRegex     any       `json:"name_regex"`
		NameRegexKeep any       `json:"name_regex_keep"`
		NextRunAt     time.Time `json:"next_run_at"`
	} `json:"container_expiration_policy"`
	RepositoryObjectFormat           string    `json:"repository_object_format"`
	IssuesEnabled                    bool      `json:"issues_enabled"`
	MergeRequestsEnabled             bool      `json:"merge_requests_enabled"`
	WikiEnabled                      bool      `json:"wiki_enabled"`
	JobsEnabled                      bool      `json:"jobs_enabled"`
	SnippetsEnabled                  bool      `json:"snippets_enabled"`
	ContainerRegistryEnabled         bool      `json:"container_registry_enabled"`
	ServiceDeskEnabled               bool      `json:"service_desk_enabled"`
	ServiceDeskAddress               string    `json:"service_desk_address"`
	CanCreateMergeRequestIn          bool      `json:"can_create_merge_request_in"`
	IssuesAccessLevel                string    `json:"issues_access_level"`
	RepositoryAccessLevel            string    `json:"repository_access_level"`
	MergeRequestsAccessLevel         string    `json:"merge_requests_access_level"`
	ForkingAccessLevel               string    `json:"forking_access_level"`
	WikiAccessLevel                  string    `json:"wiki_access_level"`
	BuildsAccessLevel                string    `json:"builds_access_level"`
	SnippetsAccessLevel              string    `json:"snippets_access_level"`
	PagesAccessLevel                 string    `json:"pages_access_level"`
	AnalyticsAccessLevel             string    `json:"analytics_access_level"`
	ContainerRegistryAccessLevel     string    `json:"container_registry_access_level"`
	SecurityAndComplianceAccessLevel string    `json:"security_and_compliance_access_level"`
	ReleasesAccessLevel              string    `json:"releases_access_level"`
	EnvironmentsAccessLevel          string    `json:"environments_access_level"`
	FeatureFlagsAccessLevel          string    `json:"feature_flags_access_level"`
	InfrastructureAccessLevel        string    `json:"infrastructure_access_level"`
	MonitorAccessLevel               string    `json:"monitor_access_level"`
	ModelExperimentsAccessLevel      string    `json:"model_experiments_access_level"`
	ModelRegistryAccessLevel         string    `json:"model_registry_access_level"`
	EmailsDisabled                   bool      `json:"emails_disabled"`
	EmailsEnabled                    bool      `json:"emails_enabled"`
	SharedRunnersEnabled             bool      `json:"shared_runners_enabled"`
	LfsEnabled                       bool      `json:"lfs_enabled"`
	CreatorID                        int       `json:"creator_id"`
	ImportStatus                     string    `json:"import_status"`
	OpenIssuesCount                  int       `json:"open_issues_count"`
	DescriptionHTML                  string    `json:"description_html"`
	UpdatedAt                        time.Time `json:"updated_at"`
	CiConfigPath                     string    `json:"ci_config_path"`
	PublicJobs                       bool      `json:"public_jobs"`
	SharedWithGroups                 []struct {
		GroupID          int    `json:"group_id"`
		GroupName        string `json:"group_name"`
		GroupFullPath    string `json:"group_full_path"`
		GroupAccessLevel int    `json:"group_access_level"`
		ExpiresAt        any    `json:"expires_at"`
	} `json:"shared_with_groups"`
	OnlyAllowMergeIfPipelineSucceeds          bool   `json:"only_allow_merge_if_pipeline_succeeds"`
	AllowMergeOnSkippedPipeline               bool   `json:"allow_merge_on_skipped_pipeline"`
	RequestAccessEnabled                      bool   `json:"request_access_enabled"`
	OnlyAllowMergeIfAllDiscussionsAreResolved bool   `json:"only_allow_merge_if_all_discussions_are_resolved"`
	RemoveSourceBranchAfterMerge              bool   `json:"remove_source_branch_after_merge"`
	PrintingMergeRequestLinkEnabled           bool   `json:"printing_merge_request_link_enabled"`
	MergeMethod                               string `json:"merge_method"`
	SquashOption                              string `json:"squash_option"`
	EnforceAuthChecksOnUploads                bool   `json:"enforce_auth_checks_on_uploads"`
	SuggestionCommitMessage                   string `json:"suggestion_commit_message"`
	MergeCommitTemplate                       any    `json:"merge_commit_template"`
	SquashCommitTemplate                      any    `json:"squash_commit_template"`
	IssueBranchTemplate                       string `json:"issue_branch_template"`
	WarnAboutPotentiallyUnwantedCharacters    bool   `json:"warn_about_potentially_unwanted_characters"`
	AutocloseReferencedIssues                 bool   `json:"autoclose_referenced_issues"`
	ApprovalsBeforeMerge                      int    `json:"approvals_before_merge"`
	Mirror                                    bool   `json:"mirror"`
	ExternalAuthorizationClassificationLabel  string `json:"external_authorization_classification_label"`
	MarkedForDeletionAt                       any    `json:"marked_for_deletion_at"`
	MarkedForDeletionOn                       any    `json:"marked_for_deletion_on"`
	RequirementsEnabled                       bool   `json:"requirements_enabled"`
	RequirementsAccessLevel                   string `json:"requirements_access_level"`
	SecurityAndComplianceEnabled              bool   `json:"security_and_compliance_enabled"`
	ComplianceFrameworks                      []any  `json:"compliance_frameworks"`
	IssuesTemplate                            any    `json:"issues_template"`
	MergeRequestsTemplate                     string `json:"merge_requests_template"`
	MergePipelinesEnabled                     bool   `json:"merge_pipelines_enabled"`
	MergeTrainsEnabled                        bool   `json:"merge_trains_enabled"`
	MergeTrainsSkipTrainAllowed               bool   `json:"merge_trains_skip_train_allowed"`
	AllowPipelineTriggerApproveDeployment     bool   `json:"allow_pipeline_trigger_approve_deployment"`
	Permissions                               struct {
		ProjectAccess any `json:"project_access"`
		GroupAccess   any `json:"group_access"`
	} `json:"permissions"`
}
