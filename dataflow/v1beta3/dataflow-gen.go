// Package dataflow provides access to the Google Dataflow API.
//
// Usage example:
//
//   import "google.golang.org/api/dataflow/v1beta3"
//   ...
//   dataflowService, err := dataflow.New(oauthHttpClient)
package dataflow

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/api/googleapi"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Background

const apiId = "dataflow:v1beta3"
const apiName = "dataflow"
const apiVersion = "v1beta3"
const basePath = "https://www.googleapis.com/dataflow/v1b3/projects/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"

	// View your email address
	UserinfoEmailScope = "https://www.googleapis.com/auth/userinfo.email"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.V1b3 = NewV1b3Service(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	V1b3 *V1b3Service
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewV1b3Service(s *Service) *V1b3Service {
	rs := &V1b3Service{s: s}
	rs.Projects = NewV1b3ProjectsService(s)
	return rs
}

type V1b3Service struct {
	s *Service

	Projects *V1b3ProjectsService
}

func NewV1b3ProjectsService(s *Service) *V1b3ProjectsService {
	rs := &V1b3ProjectsService{s: s}
	rs.Jobs = NewV1b3ProjectsJobsService(s)
	return rs
}

type V1b3ProjectsService struct {
	s *Service

	Jobs *V1b3ProjectsJobsService
}

func NewV1b3ProjectsJobsService(s *Service) *V1b3ProjectsJobsService {
	rs := &V1b3ProjectsJobsService{s: s}
	rs.Messages = NewV1b3ProjectsJobsMessagesService(s)
	rs.WorkItems = NewV1b3ProjectsJobsWorkItemsService(s)
	return rs
}

type V1b3ProjectsJobsService struct {
	s *Service

	Messages *V1b3ProjectsJobsMessagesService

	WorkItems *V1b3ProjectsJobsWorkItemsService
}

func NewV1b3ProjectsJobsMessagesService(s *Service) *V1b3ProjectsJobsMessagesService {
	rs := &V1b3ProjectsJobsMessagesService{s: s}
	return rs
}

type V1b3ProjectsJobsMessagesService struct {
	s *Service
}

func NewV1b3ProjectsJobsWorkItemsService(s *Service) *V1b3ProjectsJobsWorkItemsService {
	rs := &V1b3ProjectsJobsWorkItemsService{s: s}
	return rs
}

type V1b3ProjectsJobsWorkItemsService struct {
	s *Service
}

type ApproximateProgress struct {
	PercentComplete float64 `json:"percentComplete,omitempty"`

	Position *Position `json:"position,omitempty"`

	RemainingTime string `json:"remainingTime,omitempty"`
}

type AutoscalingSettings struct {
	Algorithm string `json:"algorithm,omitempty"`

	MaxNumWorkers int64 `json:"maxNumWorkers,omitempty"`
}

type ComputationTopology struct {
	ComputationId string `json:"computationId,omitempty"`

	Inputs []*StreamLocation `json:"inputs,omitempty"`

	KeyRanges []*KeyRangeLocation `json:"keyRanges,omitempty"`

	Outputs []*StreamLocation `json:"outputs,omitempty"`
}

type DataDiskAssignment struct {
	DataDisks []string `json:"dataDisks,omitempty"`

	VmInstance string `json:"vmInstance,omitempty"`
}

type DerivedSource struct {
	DerivationMode string `json:"derivationMode,omitempty"`

	Source *Source `json:"source,omitempty"`
}

type Disk struct {
	DiskType string `json:"diskType,omitempty"`

	MountPoint string `json:"mountPoint,omitempty"`

	SizeGb int64 `json:"sizeGb,omitempty"`
}

type DynamicSourceSplit struct {
	Primary *DerivedSource `json:"primary,omitempty"`

	Residual *DerivedSource `json:"residual,omitempty"`
}

type Environment struct {
	ClusterManagerApiService string `json:"clusterManagerApiService,omitempty"`

	Dataset string `json:"dataset,omitempty"`

	Experiments []string `json:"experiments,omitempty"`

	SdkPipelineOptions *EnvironmentSdkPipelineOptions `json:"sdkPipelineOptions,omitempty"`

	TempStoragePrefix string `json:"tempStoragePrefix,omitempty"`

	UserAgent *EnvironmentUserAgent `json:"userAgent,omitempty"`

	Version *EnvironmentVersion `json:"version,omitempty"`

	WorkerPools []*WorkerPool `json:"workerPools,omitempty"`
}

type EnvironmentSdkPipelineOptions struct {
}

type EnvironmentUserAgent struct {
}

type EnvironmentVersion struct {
}

type FlattenInstruction struct {
	Inputs []*InstructionInput `json:"inputs,omitempty"`
}

type GoogleprotobufValue interface{}

type InstructionInput struct {
	OutputNum int64 `json:"outputNum,omitempty"`

	ProducerInstructionIndex int64 `json:"producerInstructionIndex,omitempty"`
}

type InstructionOutput struct {
	Codec *InstructionOutputCodec `json:"codec,omitempty"`

	Name string `json:"name,omitempty"`
}

type InstructionOutputCodec struct {
}

type Job struct {
	CreateTime string `json:"createTime,omitempty"`

	CurrentState string `json:"currentState,omitempty"`

	CurrentStateTime string `json:"currentStateTime,omitempty"`

	Environment *Environment `json:"environment,omitempty"`

	ExecutionInfo *JobExecutionInfo `json:"executionInfo,omitempty"`

	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	ProjectId string `json:"projectId,omitempty"`

	RequestedState string `json:"requestedState,omitempty"`

	Steps []*Step `json:"steps,omitempty"`

	Type string `json:"type,omitempty"`
}

type JobExecutionInfo struct {
	Stages map[string]JobExecutionStageInfo `json:"stages,omitempty"`
}

type JobExecutionStageInfo struct {
	StepName []string `json:"stepName,omitempty"`
}

type JobMessage struct {
	Id string `json:"id,omitempty"`

	MessageImportance string `json:"messageImportance,omitempty"`

	MessageText string `json:"messageText,omitempty"`

	Time string `json:"time,omitempty"`
}

type JobMetrics struct {
	MetricTime string `json:"metricTime,omitempty"`

	Metrics []*MetricUpdate `json:"metrics,omitempty"`
}

type KeyRangeDataDiskAssignment struct {
	DataDisk string `json:"dataDisk,omitempty"`

	End string `json:"end,omitempty"`

	Start string `json:"start,omitempty"`
}

type KeyRangeLocation struct {
	DataDisk string `json:"dataDisk,omitempty"`

	DeliveryEndpoint string `json:"deliveryEndpoint,omitempty"`

	End string `json:"end,omitempty"`

	PersistentDirectory string `json:"persistentDirectory,omitempty"`

	Start string `json:"start,omitempty"`
}

type LeaseWorkItemRequest struct {
	CurrentWorkerTime string `json:"currentWorkerTime,omitempty"`

	RequestedLeaseDuration string `json:"requestedLeaseDuration,omitempty"`

	WorkItemTypes []string `json:"workItemTypes,omitempty"`

	WorkerCapabilities []string `json:"workerCapabilities,omitempty"`

	WorkerId string `json:"workerId,omitempty"`
}

type LeaseWorkItemResponse struct {
	WorkItems []*WorkItem `json:"workItems,omitempty"`
}

type ListJobMessagesResponse struct {
	JobMessages []*JobMessage `json:"jobMessages,omitempty"`

	NextPageToken string `json:"nextPageToken,omitempty"`
}

type ListJobsResponse struct {
	Jobs []*Job `json:"jobs,omitempty"`

	NextPageToken string `json:"nextPageToken,omitempty"`
}

type MapTask struct {
	Instructions []*ParallelInstruction `json:"instructions,omitempty"`

	StageName string `json:"stageName,omitempty"`

	SystemName string `json:"systemName,omitempty"`
}

type MetricStructuredName struct {
	Context map[string]string `json:"context,omitempty"`

	Name string `json:"name,omitempty"`

	Origin string `json:"origin,omitempty"`
}

type MetricUpdate struct {
	Cumulative bool `json:"cumulative,omitempty"`

	Internal interface{} `json:"internal,omitempty"`

	Kind string `json:"kind,omitempty"`

	MeanCount interface{} `json:"meanCount,omitempty"`

	MeanSum interface{} `json:"meanSum,omitempty"`

	Name *MetricStructuredName `json:"name,omitempty"`

	Scalar interface{} `json:"scalar,omitempty"`

	Set interface{} `json:"set,omitempty"`

	UpdateTime string `json:"updateTime,omitempty"`
}

type MountedDataDisk struct {
	DataDisk string `json:"dataDisk,omitempty"`
}

type MultiOutputInfo struct {
	Tag string `json:"tag,omitempty"`
}

type Package struct {
	Location string `json:"location,omitempty"`

	Name string `json:"name,omitempty"`
}

type ParDoInstruction struct {
	Input *InstructionInput `json:"input,omitempty"`

	MultiOutputInfos []*MultiOutputInfo `json:"multiOutputInfos,omitempty"`

	NumOutputs int64 `json:"numOutputs,omitempty"`

	SideInputs []*SideInputInfo `json:"sideInputs,omitempty"`

	UserFn *ParDoInstructionUserFn `json:"userFn,omitempty"`
}

type ParDoInstructionUserFn struct {
}

type ParallelInstruction struct {
	Flatten *FlattenInstruction `json:"flatten,omitempty"`

	Name string `json:"name,omitempty"`

	Outputs []*InstructionOutput `json:"outputs,omitempty"`

	ParDo *ParDoInstruction `json:"parDo,omitempty"`

	PartialGroupByKey *PartialGroupByKeyInstruction `json:"partialGroupByKey,omitempty"`

	Read *ReadInstruction `json:"read,omitempty"`

	SystemName string `json:"systemName,omitempty"`

	Write *WriteInstruction `json:"write,omitempty"`
}

type PartialGroupByKeyInstruction struct {
	Input *InstructionInput `json:"input,omitempty"`

	InputElementCodec *PartialGroupByKeyInstructionInputElementCodec `json:"inputElementCodec,omitempty"`

	ValueCombiningFn *PartialGroupByKeyInstructionValueCombiningFn `json:"valueCombiningFn,omitempty"`
}

type PartialGroupByKeyInstructionInputElementCodec struct {
}

type PartialGroupByKeyInstructionValueCombiningFn struct {
}

type Position struct {
	ByteOffset int64 `json:"byteOffset,omitempty,string"`

	End bool `json:"end,omitempty"`

	Key string `json:"key,omitempty"`

	RecordIndex int64 `json:"recordIndex,omitempty,string"`

	ShufflePosition string `json:"shufflePosition,omitempty"`
}

type PubsubLocation struct {
	DropLateData bool `json:"dropLateData,omitempty"`

	IdLabel string `json:"idLabel,omitempty"`

	Subscription string `json:"subscription,omitempty"`

	TimestampLabel string `json:"timestampLabel,omitempty"`

	Topic string `json:"topic,omitempty"`

	TrackingSubscription string `json:"trackingSubscription,omitempty"`
}

type ReadInstruction struct {
	Source *Source `json:"source,omitempty"`
}

type ReportWorkItemStatusRequest struct {
	CurrentWorkerTime string `json:"currentWorkerTime,omitempty"`

	WorkItemStatuses []*WorkItemStatus `json:"workItemStatuses,omitempty"`

	WorkerId string `json:"workerId,omitempty"`
}

type ReportWorkItemStatusResponse struct {
	WorkItemServiceStates []*WorkItemServiceState `json:"workItemServiceStates,omitempty"`
}

type SeqMapTask struct {
	Inputs []*SideInputInfo `json:"inputs,omitempty"`

	Name string `json:"name,omitempty"`

	OutputInfos []*SeqMapTaskOutputInfo `json:"outputInfos,omitempty"`

	StageName string `json:"stageName,omitempty"`

	SystemName string `json:"systemName,omitempty"`

	UserFn *SeqMapTaskUserFn `json:"userFn,omitempty"`
}

type SeqMapTaskUserFn struct {
}

type SeqMapTaskOutputInfo struct {
	Sink *Sink `json:"sink,omitempty"`

	Tag string `json:"tag,omitempty"`
}

type ShellTask struct {
	Command string `json:"command,omitempty"`

	ExitCode int64 `json:"exitCode,omitempty"`
}

type SideInputInfo struct {
	Kind *SideInputInfoKind `json:"kind,omitempty"`

	Sources []*Source `json:"sources,omitempty"`

	Tag string `json:"tag,omitempty"`
}

type SideInputInfoKind struct {
}

type Sink struct {
	Codec *SinkCodec `json:"codec,omitempty"`

	Spec *SinkSpec `json:"spec,omitempty"`
}

type SinkCodec struct {
}

type SinkSpec struct {
}

type Source struct {
	BaseSpecs []*SourceBaseSpecs `json:"baseSpecs,omitempty"`

	Codec *SourceCodec `json:"codec,omitempty"`

	DoesNotNeedSplitting bool `json:"doesNotNeedSplitting,omitempty"`

	Metadata *SourceMetadata `json:"metadata,omitempty"`

	Spec *SourceSpec `json:"spec,omitempty"`
}

type SourceBaseSpecs struct {
}

type SourceCodec struct {
}

type SourceSpec struct {
}

type SourceFork struct {
	Primary *SourceSplitShard `json:"primary,omitempty"`

	PrimarySource *DerivedSource `json:"primarySource,omitempty"`

	Residual *SourceSplitShard `json:"residual,omitempty"`

	ResidualSource *DerivedSource `json:"residualSource,omitempty"`
}

type SourceGetMetadataRequest struct {
	Source *Source `json:"source,omitempty"`
}

type SourceGetMetadataResponse struct {
	Metadata *SourceMetadata `json:"metadata,omitempty"`
}

type SourceMetadata struct {
	EstimatedSizeBytes int64 `json:"estimatedSizeBytes,omitempty,string"`

	Infinite bool `json:"infinite,omitempty"`

	ProducesSortedKeys bool `json:"producesSortedKeys,omitempty"`
}

type SourceOperationRequest struct {
	GetMetadata *SourceGetMetadataRequest `json:"getMetadata,omitempty"`

	Split *SourceSplitRequest `json:"split,omitempty"`
}

type SourceOperationResponse struct {
	GetMetadata *SourceGetMetadataResponse `json:"getMetadata,omitempty"`

	Split *SourceSplitResponse `json:"split,omitempty"`
}

type SourceSplitOptions struct {
	DesiredBundleSizeBytes int64 `json:"desiredBundleSizeBytes,omitempty,string"`

	DesiredShardSizeBytes int64 `json:"desiredShardSizeBytes,omitempty,string"`
}

type SourceSplitRequest struct {
	Options *SourceSplitOptions `json:"options,omitempty"`

	Source *Source `json:"source,omitempty"`
}

type SourceSplitResponse struct {
	Bundles []*DerivedSource `json:"bundles,omitempty"`

	Outcome string `json:"outcome,omitempty"`

	Shards []*SourceSplitShard `json:"shards,omitempty"`
}

type SourceSplitShard struct {
	DerivationMode string `json:"derivationMode,omitempty"`

	Source *Source `json:"source,omitempty"`
}

type Status struct {
	Code int64 `json:"code,omitempty"`

	Details []*StatusDetails `json:"details,omitempty"`

	Message string `json:"message,omitempty"`
}

type StatusDetails struct {
}

type Step struct {
	Kind string `json:"kind,omitempty"`

	Name string `json:"name,omitempty"`

	Properties *StepProperties `json:"properties,omitempty"`
}

type StepProperties struct {
}

type StreamLocation struct {
	PubsubLocation *PubsubLocation `json:"pubsubLocation,omitempty"`

	SideInputLocation *StreamingSideInputLocation `json:"sideInputLocation,omitempty"`

	StreamingStageLocation *StreamingStageLocation `json:"streamingStageLocation,omitempty"`
}

type StreamingComputationRanges struct {
	ComputationId string `json:"computationId,omitempty"`

	RangeAssignments []*KeyRangeDataDiskAssignment `json:"rangeAssignments,omitempty"`
}

type StreamingComputationTask struct {
	ComputationRanges []*StreamingComputationRanges `json:"computationRanges,omitempty"`

	DataDisks []*MountedDataDisk `json:"dataDisks,omitempty"`

	TaskType string `json:"taskType,omitempty"`
}

type StreamingSetupTask struct {
	ReceiveWorkPort int64 `json:"receiveWorkPort,omitempty"`

	StreamingComputationTopology *TopologyConfig `json:"streamingComputationTopology,omitempty"`

	WorkerHarnessPort int64 `json:"workerHarnessPort,omitempty"`
}

type StreamingSideInputLocation struct {
	Tag string `json:"tag,omitempty"`
}

type StreamingStageLocation struct {
	StreamId string `json:"streamId,omitempty"`
}

type TaskRunnerSettings struct {
	Alsologtostderr bool `json:"alsologtostderr,omitempty"`

	BaseTaskDir string `json:"baseTaskDir,omitempty"`

	BaseUrl string `json:"baseUrl,omitempty"`

	CommandlinesFileName string `json:"commandlinesFileName,omitempty"`

	ContinueOnException bool `json:"continueOnException,omitempty"`

	DataflowApiVersion string `json:"dataflowApiVersion,omitempty"`

	HarnessCommand string `json:"harnessCommand,omitempty"`

	LanguageHint string `json:"languageHint,omitempty"`

	LogDir string `json:"logDir,omitempty"`

	LogToSerialconsole bool `json:"logToSerialconsole,omitempty"`

	LogUploadLocation string `json:"logUploadLocation,omitempty"`

	OauthScopes []string `json:"oauthScopes,omitempty"`

	ParallelWorkerSettings *WorkerSettings `json:"parallelWorkerSettings,omitempty"`

	StreamingWorkerMainClass string `json:"streamingWorkerMainClass,omitempty"`

	TaskGroup string `json:"taskGroup,omitempty"`

	TaskUser string `json:"taskUser,omitempty"`

	TempStoragePrefix string `json:"tempStoragePrefix,omitempty"`

	VmId string `json:"vmId,omitempty"`

	WorkflowFileName string `json:"workflowFileName,omitempty"`
}

type TopologyConfig struct {
	Computations []*ComputationTopology `json:"computations,omitempty"`

	DataDiskAssignments []*DataDiskAssignment `json:"dataDiskAssignments,omitempty"`
}

type WorkItem struct {
	Configuration string `json:"configuration,omitempty"`

	Id int64 `json:"id,omitempty,string"`

	InitialReportIndex int64 `json:"initialReportIndex,omitempty,string"`

	JobId string `json:"jobId,omitempty"`

	LeaseExpireTime string `json:"leaseExpireTime,omitempty"`

	MapTask *MapTask `json:"mapTask,omitempty"`

	Packages []*Package `json:"packages,omitempty"`

	ProjectId string `json:"projectId,omitempty"`

	ReportStatusInterval string `json:"reportStatusInterval,omitempty"`

	SeqMapTask *SeqMapTask `json:"seqMapTask,omitempty"`

	ShellTask *ShellTask `json:"shellTask,omitempty"`

	SourceOperationTask *SourceOperationRequest `json:"sourceOperationTask,omitempty"`

	StreamingComputationTask *StreamingComputationTask `json:"streamingComputationTask,omitempty"`

	StreamingSetupTask *StreamingSetupTask `json:"streamingSetupTask,omitempty"`
}

type WorkItemServiceState struct {
	HarnessData *WorkItemServiceStateHarnessData `json:"harnessData,omitempty"`

	LeaseExpireTime string `json:"leaseExpireTime,omitempty"`

	NextReportIndex int64 `json:"nextReportIndex,omitempty,string"`

	ReportStatusInterval string `json:"reportStatusInterval,omitempty"`

	SuggestedStopPoint *ApproximateProgress `json:"suggestedStopPoint,omitempty"`

	SuggestedStopPosition *Position `json:"suggestedStopPosition,omitempty"`
}

type WorkItemServiceStateHarnessData struct {
}

type WorkItemStatus struct {
	Completed bool `json:"completed,omitempty"`

	DynamicSourceSplit *DynamicSourceSplit `json:"dynamicSourceSplit,omitempty"`

	Errors []*Status `json:"errors,omitempty"`

	MetricUpdates []*MetricUpdate `json:"metricUpdates,omitempty"`

	Progress *ApproximateProgress `json:"progress,omitempty"`

	ReportIndex int64 `json:"reportIndex,omitempty,string"`

	RequestedLeaseDuration string `json:"requestedLeaseDuration,omitempty"`

	SourceFork *SourceFork `json:"sourceFork,omitempty"`

	SourceOperationResponse *SourceOperationResponse `json:"sourceOperationResponse,omitempty"`

	StopPosition *Position `json:"stopPosition,omitempty"`

	WorkItemId string `json:"workItemId,omitempty"`
}

type WorkerPool struct {
	AutoscalingSettings *AutoscalingSettings `json:"autoscalingSettings,omitempty"`

	DataDisks []*Disk `json:"dataDisks,omitempty"`

	DefaultPackageSet string `json:"defaultPackageSet,omitempty"`

	DiskSizeGb int64 `json:"diskSizeGb,omitempty"`

	DiskSourceImage string `json:"diskSourceImage,omitempty"`

	Kind string `json:"kind,omitempty"`

	MachineType string `json:"machineType,omitempty"`

	Metadata map[string]string `json:"metadata,omitempty"`

	NumWorkers int64 `json:"numWorkers,omitempty"`

	OnHostMaintenance string `json:"onHostMaintenance,omitempty"`

	Packages []*Package `json:"packages,omitempty"`

	PoolArgs *WorkerPoolPoolArgs `json:"poolArgs,omitempty"`

	TaskrunnerSettings *TaskRunnerSettings `json:"taskrunnerSettings,omitempty"`

	TeardownPolicy string `json:"teardownPolicy,omitempty"`

	Zone string `json:"zone,omitempty"`
}

type WorkerPoolPoolArgs struct {
}

type WorkerSettings struct {
	BaseUrl string `json:"baseUrl,omitempty"`

	ReportingEnabled bool `json:"reportingEnabled,omitempty"`

	ServicePath string `json:"servicePath,omitempty"`

	ShuffleServicePath string `json:"shuffleServicePath,omitempty"`

	TempStoragePrefix string `json:"tempStoragePrefix,omitempty"`

	WorkerId string `json:"workerId,omitempty"`
}

type WriteInstruction struct {
	Input *InstructionInput `json:"input,omitempty"`

	Sink *Sink `json:"sink,omitempty"`
}

// method id "dataflow.v1b3.projects.jobs.create":

type V1b3ProjectsJobsCreateCall struct {
	s         *Service
	projectId string
	job       *Job
	opt_      map[string]interface{}
}

// Create: Creates a dataflow job.
func (r *V1b3ProjectsJobsService) Create(projectId string, job *Job) *V1b3ProjectsJobsCreateCall {
	c := &V1b3ProjectsJobsCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.job = job
	return c
}

// View sets the optional parameter "view":
func (c *V1b3ProjectsJobsCreateCall) View(view string) *V1b3ProjectsJobsCreateCall {
	c.opt_["view"] = view
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *V1b3ProjectsJobsCreateCall) Fields(s ...googleapi.Field) *V1b3ProjectsJobsCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *V1b3ProjectsJobsCreateCall) Do() (*Job, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.job)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["view"]; ok {
		params.Set("view", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{projectId}/jobs")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Job
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a dataflow job.",
	//   "httpMethod": "POST",
	//   "id": "dataflow.v1b3.projects.jobs.create",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "view": {
	//       "enum": [
	//         "JOB_VIEW_ALL",
	//         "JOB_VIEW_SUMMARY",
	//         "JOB_VIEW_UNKNOWN"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "{projectId}/jobs",
	//   "request": {
	//     "$ref": "Job"
	//   },
	//   "response": {
	//     "$ref": "Job"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "dataflow.v1b3.projects.jobs.get":

type V1b3ProjectsJobsGetCall struct {
	s         *Service
	projectId string
	jobId     string
	opt_      map[string]interface{}
}

// Get: Gets the state of the specified dataflow job.
func (r *V1b3ProjectsJobsService) Get(projectId string, jobId string) *V1b3ProjectsJobsGetCall {
	c := &V1b3ProjectsJobsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.jobId = jobId
	return c
}

// View sets the optional parameter "view":
func (c *V1b3ProjectsJobsGetCall) View(view string) *V1b3ProjectsJobsGetCall {
	c.opt_["view"] = view
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *V1b3ProjectsJobsGetCall) Fields(s ...googleapi.Field) *V1b3ProjectsJobsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *V1b3ProjectsJobsGetCall) Do() (*Job, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["view"]; ok {
		params.Set("view", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{projectId}/jobs/{jobId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"jobId":     c.jobId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Job
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the state of the specified dataflow job.",
	//   "httpMethod": "GET",
	//   "id": "dataflow.v1b3.projects.jobs.get",
	//   "parameterOrder": [
	//     "projectId",
	//     "jobId"
	//   ],
	//   "parameters": {
	//     "jobId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "view": {
	//       "enum": [
	//         "JOB_VIEW_ALL",
	//         "JOB_VIEW_SUMMARY",
	//         "JOB_VIEW_UNKNOWN"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "{projectId}/jobs/{jobId}",
	//   "response": {
	//     "$ref": "Job"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "dataflow.v1b3.projects.jobs.getMetrics":

type V1b3ProjectsJobsGetMetricsCall struct {
	s         *Service
	projectId string
	jobId     string
	opt_      map[string]interface{}
}

// GetMetrics: Request the job status.
func (r *V1b3ProjectsJobsService) GetMetrics(projectId string, jobId string) *V1b3ProjectsJobsGetMetricsCall {
	c := &V1b3ProjectsJobsGetMetricsCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.jobId = jobId
	return c
}

// StartTime sets the optional parameter "startTime":
func (c *V1b3ProjectsJobsGetMetricsCall) StartTime(startTime string) *V1b3ProjectsJobsGetMetricsCall {
	c.opt_["startTime"] = startTime
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *V1b3ProjectsJobsGetMetricsCall) Fields(s ...googleapi.Field) *V1b3ProjectsJobsGetMetricsCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *V1b3ProjectsJobsGetMetricsCall) Do() (*JobMetrics, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["startTime"]; ok {
		params.Set("startTime", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{projectId}/jobs/{jobId}/metrics")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"jobId":     c.jobId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *JobMetrics
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Request the job status.",
	//   "httpMethod": "GET",
	//   "id": "dataflow.v1b3.projects.jobs.getMetrics",
	//   "parameterOrder": [
	//     "projectId",
	//     "jobId"
	//   ],
	//   "parameters": {
	//     "jobId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "startTime": {
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "{projectId}/jobs/{jobId}/metrics",
	//   "response": {
	//     "$ref": "JobMetrics"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "dataflow.v1b3.projects.jobs.list":

type V1b3ProjectsJobsListCall struct {
	s         *Service
	projectId string
	opt_      map[string]interface{}
}

// List: List the jobs of a project
func (r *V1b3ProjectsJobsService) List(projectId string) *V1b3ProjectsJobsListCall {
	c := &V1b3ProjectsJobsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	return c
}

// PageSize sets the optional parameter "pageSize":
func (c *V1b3ProjectsJobsListCall) PageSize(pageSize int64) *V1b3ProjectsJobsListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken":
func (c *V1b3ProjectsJobsListCall) PageToken(pageToken string) *V1b3ProjectsJobsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// View sets the optional parameter "view":
func (c *V1b3ProjectsJobsListCall) View(view string) *V1b3ProjectsJobsListCall {
	c.opt_["view"] = view
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *V1b3ProjectsJobsListCall) Fields(s ...googleapi.Field) *V1b3ProjectsJobsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *V1b3ProjectsJobsListCall) Do() (*ListJobsResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["pageSize"]; ok {
		params.Set("pageSize", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["view"]; ok {
		params.Set("view", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{projectId}/jobs")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ListJobsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List the jobs of a project",
	//   "httpMethod": "GET",
	//   "id": "dataflow.v1b3.projects.jobs.list",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "view": {
	//       "enum": [
	//         "JOB_VIEW_ALL",
	//         "JOB_VIEW_SUMMARY",
	//         "JOB_VIEW_UNKNOWN"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "{projectId}/jobs",
	//   "response": {
	//     "$ref": "ListJobsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "dataflow.v1b3.projects.jobs.patch":

type V1b3ProjectsJobsPatchCall struct {
	s         *Service
	projectId string
	jobId     string
	job       *Job
	opt_      map[string]interface{}
}

// Patch: Updates the state of an existing dataflow job. This method
// supports patch semantics.
func (r *V1b3ProjectsJobsService) Patch(projectId string, jobId string, job *Job) *V1b3ProjectsJobsPatchCall {
	c := &V1b3ProjectsJobsPatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.jobId = jobId
	c.job = job
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *V1b3ProjectsJobsPatchCall) Fields(s ...googleapi.Field) *V1b3ProjectsJobsPatchCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *V1b3ProjectsJobsPatchCall) Do() (*Job, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.job)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{projectId}/jobs/{jobId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"jobId":     c.jobId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Job
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates the state of an existing dataflow job. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "dataflow.v1b3.projects.jobs.patch",
	//   "parameterOrder": [
	//     "projectId",
	//     "jobId"
	//   ],
	//   "parameters": {
	//     "jobId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{projectId}/jobs/{jobId}",
	//   "request": {
	//     "$ref": "Job"
	//   },
	//   "response": {
	//     "$ref": "Job"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "dataflow.v1b3.projects.jobs.update":

type V1b3ProjectsJobsUpdateCall struct {
	s         *Service
	projectId string
	jobId     string
	job       *Job
	opt_      map[string]interface{}
}

// Update: Updates the state of an existing dataflow job.
func (r *V1b3ProjectsJobsService) Update(projectId string, jobId string, job *Job) *V1b3ProjectsJobsUpdateCall {
	c := &V1b3ProjectsJobsUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.jobId = jobId
	c.job = job
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *V1b3ProjectsJobsUpdateCall) Fields(s ...googleapi.Field) *V1b3ProjectsJobsUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *V1b3ProjectsJobsUpdateCall) Do() (*Job, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.job)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{projectId}/jobs/{jobId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"jobId":     c.jobId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Job
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates the state of an existing dataflow job.",
	//   "httpMethod": "PUT",
	//   "id": "dataflow.v1b3.projects.jobs.update",
	//   "parameterOrder": [
	//     "projectId",
	//     "jobId"
	//   ],
	//   "parameters": {
	//     "jobId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{projectId}/jobs/{jobId}",
	//   "request": {
	//     "$ref": "Job"
	//   },
	//   "response": {
	//     "$ref": "Job"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "dataflow.v1b3.projects.jobs.messages.list":

type V1b3ProjectsJobsMessagesListCall struct {
	s         *Service
	projectId string
	jobId     string
	opt_      map[string]interface{}
}

// List: Request the job status.
func (r *V1b3ProjectsJobsMessagesService) List(projectId string, jobId string) *V1b3ProjectsJobsMessagesListCall {
	c := &V1b3ProjectsJobsMessagesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.jobId = jobId
	return c
}

// EndTime sets the optional parameter "endTime":
func (c *V1b3ProjectsJobsMessagesListCall) EndTime(endTime string) *V1b3ProjectsJobsMessagesListCall {
	c.opt_["endTime"] = endTime
	return c
}

// MinimumImportance sets the optional parameter "minimumImportance":
func (c *V1b3ProjectsJobsMessagesListCall) MinimumImportance(minimumImportance string) *V1b3ProjectsJobsMessagesListCall {
	c.opt_["minimumImportance"] = minimumImportance
	return c
}

// PageSize sets the optional parameter "pageSize":
func (c *V1b3ProjectsJobsMessagesListCall) PageSize(pageSize int64) *V1b3ProjectsJobsMessagesListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken":
func (c *V1b3ProjectsJobsMessagesListCall) PageToken(pageToken string) *V1b3ProjectsJobsMessagesListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// StartTime sets the optional parameter "startTime":
func (c *V1b3ProjectsJobsMessagesListCall) StartTime(startTime string) *V1b3ProjectsJobsMessagesListCall {
	c.opt_["startTime"] = startTime
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *V1b3ProjectsJobsMessagesListCall) Fields(s ...googleapi.Field) *V1b3ProjectsJobsMessagesListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *V1b3ProjectsJobsMessagesListCall) Do() (*ListJobMessagesResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["endTime"]; ok {
		params.Set("endTime", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["minimumImportance"]; ok {
		params.Set("minimumImportance", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageSize"]; ok {
		params.Set("pageSize", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["startTime"]; ok {
		params.Set("startTime", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{projectId}/jobs/{jobId}/messages")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"jobId":     c.jobId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ListJobMessagesResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Request the job status.",
	//   "httpMethod": "GET",
	//   "id": "dataflow.v1b3.projects.jobs.messages.list",
	//   "parameterOrder": [
	//     "projectId",
	//     "jobId"
	//   ],
	//   "parameters": {
	//     "endTime": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "jobId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "minimumImportance": {
	//       "enum": [
	//         "JOB_MESSAGE_DEBUG",
	//         "JOB_MESSAGE_DETAILED",
	//         "JOB_MESSAGE_ERROR",
	//         "JOB_MESSAGE_IMPORTANCE_UNKNOWN",
	//         "JOB_MESSAGE_WARNING"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         "",
	//         "",
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "startTime": {
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "{projectId}/jobs/{jobId}/messages",
	//   "response": {
	//     "$ref": "ListJobMessagesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "dataflow.v1b3.projects.jobs.workItems.lease":

type V1b3ProjectsJobsWorkItemsLeaseCall struct {
	s                    *Service
	projectId            string
	jobId                string
	leaseworkitemrequest *LeaseWorkItemRequest
	opt_                 map[string]interface{}
}

// Lease: Leases a dataflow WorkItem to run.
func (r *V1b3ProjectsJobsWorkItemsService) Lease(projectId string, jobId string, leaseworkitemrequest *LeaseWorkItemRequest) *V1b3ProjectsJobsWorkItemsLeaseCall {
	c := &V1b3ProjectsJobsWorkItemsLeaseCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.jobId = jobId
	c.leaseworkitemrequest = leaseworkitemrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *V1b3ProjectsJobsWorkItemsLeaseCall) Fields(s ...googleapi.Field) *V1b3ProjectsJobsWorkItemsLeaseCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *V1b3ProjectsJobsWorkItemsLeaseCall) Do() (*LeaseWorkItemResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.leaseworkitemrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{projectId}/jobs/{jobId}/workItems:lease")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"jobId":     c.jobId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *LeaseWorkItemResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Leases a dataflow WorkItem to run.",
	//   "httpMethod": "POST",
	//   "id": "dataflow.v1b3.projects.jobs.workItems.lease",
	//   "parameterOrder": [
	//     "projectId",
	//     "jobId"
	//   ],
	//   "parameters": {
	//     "jobId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{projectId}/jobs/{jobId}/workItems:lease",
	//   "request": {
	//     "$ref": "LeaseWorkItemRequest"
	//   },
	//   "response": {
	//     "$ref": "LeaseWorkItemResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "dataflow.v1b3.projects.jobs.workItems.reportStatus":

type V1b3ProjectsJobsWorkItemsReportStatusCall struct {
	s                           *Service
	projectId                   string
	jobId                       string
	reportworkitemstatusrequest *ReportWorkItemStatusRequest
	opt_                        map[string]interface{}
}

// ReportStatus: Reports the status of dataflow WorkItems leased by a
// worker.
func (r *V1b3ProjectsJobsWorkItemsService) ReportStatus(projectId string, jobId string, reportworkitemstatusrequest *ReportWorkItemStatusRequest) *V1b3ProjectsJobsWorkItemsReportStatusCall {
	c := &V1b3ProjectsJobsWorkItemsReportStatusCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.jobId = jobId
	c.reportworkitemstatusrequest = reportworkitemstatusrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *V1b3ProjectsJobsWorkItemsReportStatusCall) Fields(s ...googleapi.Field) *V1b3ProjectsJobsWorkItemsReportStatusCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *V1b3ProjectsJobsWorkItemsReportStatusCall) Do() (*ReportWorkItemStatusResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.reportworkitemstatusrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{projectId}/jobs/{jobId}/workItems:reportStatus")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"jobId":     c.jobId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ReportWorkItemStatusResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Reports the status of dataflow WorkItems leased by a worker.",
	//   "httpMethod": "POST",
	//   "id": "dataflow.v1b3.projects.jobs.workItems.reportStatus",
	//   "parameterOrder": [
	//     "projectId",
	//     "jobId"
	//   ],
	//   "parameters": {
	//     "jobId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{projectId}/jobs/{jobId}/workItems:reportStatus",
	//   "request": {
	//     "$ref": "ReportWorkItemStatusRequest"
	//   },
	//   "response": {
	//     "$ref": "ReportWorkItemStatusResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}
