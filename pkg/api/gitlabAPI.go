package api

import (
	"github.com/BtQuentin/gitlab-mr-following/pkg/configuration"
	"github.com/go-resty/resty/v2"
	"slices"
	"strconv"
	"time"
)

type MergeRequests []struct {
	ID             int         `json:"id"`
	Iid            int         `json:"iid"`
	ProjectID      int         `json:"project_id"`
	Title          string      `json:"title"`
	Description    string      `json:"description"`
	State          string      `json:"state"`
	CreatedAt      time.Time   `json:"created_at"`
	UpdatedAt      time.Time   `json:"updated_at"`
	MergedBy       interface{} `json:"merged_by"`
	MergeUser      interface{} `json:"merge_user"`
	MergedAt       interface{} `json:"merged_at"`
	ClosedBy       interface{} `json:"closed_by"`
	ClosedAt       interface{} `json:"closed_at"`
	TargetBranch   string      `json:"target_branch"`
	SourceBranch   string      `json:"source_branch"`
	UserNotesCount int         `json:"user_notes_count"`
	Upvotes        int         `json:"upvotes"`
	Downvotes      int         `json:"downvotes"`
	Author         struct {
		ID        int    `json:"id"`
		Username  string `json:"username"`
		Name      string `json:"name"`
		State     string `json:"state"`
		AvatarURL string `json:"avatar_url"`
		WebURL    string `json:"web_url"`
	} `json:"author"`
	Assignees                 []interface{} `json:"assignees"`
	Assignee                  interface{}   `json:"assignee"`
	Reviewers                 []interface{} `json:"reviewers"`
	SourceProjectID           int           `json:"source_project_id"`
	TargetProjectID           int           `json:"target_project_id"`
	Labels                    []interface{} `json:"labels"`
	Draft                     bool          `json:"draft"`
	WorkInProgress            bool          `json:"work_in_progress"`
	Milestone                 interface{}   `json:"milestone"`
	MergeWhenPipelineSucceeds bool          `json:"merge_when_pipeline_succeeds"`
	MergeStatus               string        `json:"merge_status"`
	Sha                       string        `json:"sha"`
	MergeCommitSha            interface{}   `json:"merge_commit_sha"`
	SquashCommitSha           interface{}   `json:"squash_commit_sha"`
	DiscussionLocked          interface{}   `json:"discussion_locked"`
	ShouldRemoveSourceBranch  interface{}   `json:"should_remove_source_branch"`
	ForceRemoveSourceBranch   bool          `json:"force_remove_source_branch"`
	Reference                 string        `json:"reference"`
	References                struct {
		Short    string `json:"short"`
		Relative string `json:"relative"`
		Full     string `json:"full"`
	} `json:"references"`
	WebURL    string `json:"web_url"`
	TimeStats struct {
		TimeEstimate        int         `json:"time_estimate"`
		TotalTimeSpent      int         `json:"total_time_spent"`
		HumanTimeEstimate   interface{} `json:"human_time_estimate"`
		HumanTotalTimeSpent interface{} `json:"human_total_time_spent"`
	} `json:"time_stats"`
	Squash               bool `json:"squash"`
	TaskCompletionStatus struct {
		Count          int `json:"count"`
		CompletedCount int `json:"completed_count"`
	} `json:"task_completion_status"`
	HasConflicts                bool `json:"has_conflicts"`
	BlockingDiscussionsResolved bool `json:"blocking_discussions_resolved"`
}

type Projects struct {
	Id   int
	Name string
}

type Project struct {
	ID                int           `json:"id"`
	Description       string        `json:"description"`
	Name              string        `json:"name"`
	NameWithNamespace string        `json:"name_with_namespace"`
	Path              string        `json:"path"`
	PathWithNamespace string        `json:"path_with_namespace"`
	CreatedAt         time.Time     `json:"created_at"`
	DefaultBranch     string        `json:"default_branch"`
	TagList           []interface{} `json:"tag_list"`
	Topics            []interface{} `json:"topics"`
	SSHURLToRepo      string        `json:"ssh_url_to_repo"`
	HTTPURLToRepo     string        `json:"http_url_to_repo"`
	WebURL            string        `json:"web_url"`
	ReadmeURL         interface{}   `json:"readme_url"`
	AvatarURL         interface{}   `json:"avatar_url"`
	ForksCount        int           `json:"forks_count"`
	StarCount         int           `json:"star_count"`
	LastActivityAt    time.Time     `json:"last_activity_at"`
	Namespace         struct {
		ID        int         `json:"id"`
		Name      string      `json:"name"`
		Path      string      `json:"path"`
		Kind      string      `json:"kind"`
		FullPath  string      `json:"full_path"`
		ParentID  interface{} `json:"parent_id"`
		AvatarURL string      `json:"avatar_url"`
		WebURL    string      `json:"web_url"`
	} `json:"namespace"`
	ContainerRegistryImagePrefix string `json:"container_registry_image_prefix"`
	Links                        struct {
		Self          string `json:"self"`
		MergeRequests string `json:"merge_requests"`
		RepoBranches  string `json:"repo_branches"`
		Labels        string `json:"labels"`
		Events        string `json:"events"`
		Members       string `json:"members"`
	} `json:"_links"`
	PackagesEnabled                bool        `json:"packages_enabled"`
	EmptyRepo                      bool        `json:"empty_repo"`
	Archived                       bool        `json:"archived"`
	Visibility                     string      `json:"visibility"`
	ResolveOutdatedDiffDiscussions interface{} `json:"resolve_outdated_diff_discussions"`
	ContainerExpirationPolicy      struct {
		Cadence       string    `json:"cadence"`
		Enabled       bool      `json:"enabled"`
		KeepN         int       `json:"keep_n"`
		OlderThan     string    `json:"older_than"`
		NameRegex     string    `json:"name_regex"`
		NameRegexKeep string    `json:"name_regex_keep"`
		NextRunAt     time.Time `json:"next_run_at"`
	} `json:"container_expiration_policy"`
	IssuesEnabled                    bool        `json:"issues_enabled"`
	MergeRequestsEnabled             bool        `json:"merge_requests_enabled"`
	WikiEnabled                      bool        `json:"wiki_enabled"`
	JobsEnabled                      bool        `json:"jobs_enabled"`
	SnippetsEnabled                  bool        `json:"snippets_enabled"`
	ContainerRegistryEnabled         bool        `json:"container_registry_enabled"`
	ServiceDeskEnabled               bool        `json:"service_desk_enabled"`
	CanCreateMergeRequestIn          bool        `json:"can_create_merge_request_in"`
	IssuesAccessLevel                string      `json:"issues_access_level"`
	RepositoryAccessLevel            string      `json:"repository_access_level"`
	MergeRequestsAccessLevel         string      `json:"merge_requests_access_level"`
	ForkingAccessLevel               string      `json:"forking_access_level"`
	WikiAccessLevel                  string      `json:"wiki_access_level"`
	BuildsAccessLevel                string      `json:"builds_access_level"`
	SnippetsAccessLevel              string      `json:"snippets_access_level"`
	PagesAccessLevel                 string      `json:"pages_access_level"`
	OperationsAccessLevel            string      `json:"operations_access_level"`
	AnalyticsAccessLevel             string      `json:"analytics_access_level"`
	ContainerRegistryAccessLevel     string      `json:"container_registry_access_level"`
	SecurityAndComplianceAccessLevel string      `json:"security_and_compliance_access_level"`
	EmailsDisabled                   bool        `json:"emails_disabled"`
	SharedRunnersEnabled             bool        `json:"shared_runners_enabled"`
	LfsEnabled                       bool        `json:"lfs_enabled"`
	CreatorID                        int         `json:"creator_id"`
	ImportURL                        interface{} `json:"import_url"`
	ImportType                       interface{} `json:"import_type"`
	ImportStatus                     string      `json:"import_status"`
	ImportError                      interface{} `json:"import_error"`
	RunnersToken                     string      `json:"runners_token"`
	CiDefaultGitDepth                interface{} `json:"ci_default_git_depth"`
	CiForwardDeploymentEnabled       interface{} `json:"ci_forward_deployment_enabled"`
	CiJobTokenScopeEnabled           bool        `json:"ci_job_token_scope_enabled"`
	CiSeparatedCaches                bool        `json:"ci_separated_caches"`
	PublicJobs                       bool        `json:"public_jobs"`
	BuildGitStrategy                 string      `json:"build_git_strategy"`
	BuildTimeout                     int         `json:"build_timeout"`
	AutoCancelPendingPipelines       string      `json:"auto_cancel_pending_pipelines"`
	BuildCoverageRegex               string      `json:"build_coverage_regex"`
	CiConfigPath                     interface{} `json:"ci_config_path"`
	SharedWithGroups                 []struct {
		GroupID          int         `json:"group_id"`
		GroupName        string      `json:"group_name"`
		GroupFullPath    string      `json:"group_full_path"`
		GroupAccessLevel int         `json:"group_access_level"`
		ExpiresAt        interface{} `json:"expires_at"`
	} `json:"shared_with_groups"`
	OnlyAllowMergeIfPipelineSucceeds          bool        `json:"only_allow_merge_if_pipeline_succeeds"`
	AllowMergeOnSkippedPipeline               interface{} `json:"allow_merge_on_skipped_pipeline"`
	RestrictUserDefinedVariables              bool        `json:"restrict_user_defined_variables"`
	RequestAccessEnabled                      bool        `json:"request_access_enabled"`
	OnlyAllowMergeIfAllDiscussionsAreResolved bool        `json:"only_allow_merge_if_all_discussions_are_resolved"`
	RemoveSourceBranchAfterMerge              interface{} `json:"remove_source_branch_after_merge"`
	PrintingMergeRequestLinkEnabled           bool        `json:"printing_merge_request_link_enabled"`
	MergeMethod                               string      `json:"merge_method"`
	SquashOption                              string      `json:"squash_option"`
	SuggestionCommitMessage                   interface{} `json:"suggestion_commit_message"`
	MergeCommitTemplate                       interface{} `json:"merge_commit_template"`
	SquashCommitTemplate                      interface{} `json:"squash_commit_template"`
	AutoDevopsEnabled                         bool        `json:"auto_devops_enabled"`
	AutoDevopsDeployStrategy                  string      `json:"auto_devops_deploy_strategy"`
	AutocloseReferencedIssues                 bool        `json:"autoclose_referenced_issues"`
	RepositoryStorage                         string      `json:"repository_storage"`
	KeepLatestArtifact                        bool        `json:"keep_latest_artifact"`
	RunnerTokenExpirationInterval             interface{} `json:"runner_token_expiration_interval"`
	Permissions                               struct {
		ProjectAccess interface{} `json:"project_access"`
		GroupAccess   struct {
			AccessLevel       int `json:"access_level"`
			NotificationLevel int `json:"notification_level"`
		} `json:"group_access"`
	} `json:"permissions"`
}

type Details struct {
	Title       string
	Author      string
	CreatedAt   string
	UpdatedAt   string
	Link        string
	ProjectName string
}

func GetProjectName(g configuration.Gitlab, p int) string {
	return getProjectDetails(g, p).Name
}

func GetMergeRequestsByProject(g configuration.Gitlab, projectID int, orderBy, sort, wip string) []Details {
	client := resty.New()

	var mrs MergeRequests

	_, err := client.R().
		SetResult(&mrs).
		SetQueryParams(map[string]string{
			"state":    "opened",
			"scope":    "all",
			"order_by": orderBy,
			"sort":     sort,
			"wip":      wip,
		}).
		SetHeader("PRIVATE-TOKEN", g.Headers.PrivateToken).
		SetHeader("accept", g.Headers.Accept).
		Get(g.API + "projects/" + strconv.Itoa(projectID) + "/merge_requests")

	if err != nil {
		panic(err)
	}

	var mrDetails []Details

	for i := 0; i < len(mrs); i++ {
		mrDetails = append(mrDetails, Details{
			Title:       mrs[i].Title,
			Author:      mrs[i].Author.Name,
			CreatedAt:   mrs[i].CreatedAt.Format("02/01/2006"),
			UpdatedAt:   mrs[i].UpdatedAt.Format("02/01/2006"),
			Link:        mrs[i].WebURL,
			ProjectName: "",
		})
	}

	return mrDetails
}

func GetMergeRequestsByTeam(g configuration.Gitlab, team, orderBy, sort, wip string) []Details {
	client := resty.New()

	var mrs MergeRequests

	_, err := client.R().
		SetResult(&mrs).
		SetQueryParams(map[string]string{
			"state":    "opened",
			"scope":    "all",
			"order_by": orderBy,
			"search":   team + "-",
			"in":       "title",
			"sort":     sort,
			"wip":      wip,
		}).
		SetHeader("PRIVATE-TOKEN", g.Headers.PrivateToken).
		SetHeader("accept", g.Headers.Accept).
		Get(g.API + "merge_requests")

	if err != nil {
		panic(err)
	}

	p := getProjectsFromMergeRequests(mrs)

	var projects []Projects
	for i := 0; i < len(p); i++ {
		d := getProjectDetails(g, p[i])
		projects = append(projects, Projects{
			Id:   p[i],
			Name: d.Name,
		})
	}

	var mrDetails []Details

	for i := 0; i < len(mrs); i++ {
		mrDetails = append(mrDetails, Details{
			Title:       mrs[i].Title,
			Author:      mrs[i].Author.Name,
			CreatedAt:   mrs[i].CreatedAt.Format("02/01/2006"),
			UpdatedAt:   mrs[i].UpdatedAt.Format("02/01/2006"),
			Link:        mrs[i].WebURL,
			ProjectName: projects[slices.IndexFunc(projects, func(p Projects) bool { return p.Id == mrs[i].ProjectID })].Name,
		})
	}

	return mrDetails
}

func getProjectsFromMergeRequests(mrs MergeRequests) []int {

	var projects []int
	for i := 0; i < len(mrs); i++ {
		if !slices.Contains(projects, mrs[i].ProjectID) {
			projects = append(projects, mrs[i].ProjectID)
		}
	}

	return projects
}

func getProjectDetails(
	g configuration.Gitlab,
	p int,
) Project {
	client := resty.New()

	var project Project

	_, err := client.R().
		SetResult(&project).
		SetHeader("PRIVATE-TOKEN", g.Headers.PrivateToken).
		SetHeader("accept", g.Headers.Accept).
		Get(g.API + "projects/" + strconv.Itoa(p))

	if err != nil {
		panic(err)
	}

	return project
}
