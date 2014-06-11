// Package genomics provides access to the Genomics API.
//
// See https://developers.google.com/genomics/v1beta/reference
//
// Usage example:
//
//   import "code.google.com/p/google-api-go-client/genomics/v1beta"
//   ...
//   genomicsService, err := genomics.New(oauthHttpClient)
package genomics

import (
	"bytes"
	"code.google.com/p/google-api-go-client/googleapi"
	"encoding/json"
	"errors"
	"fmt"
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

const apiId = "genomics:v1beta"
const apiName = "genomics"
const apiVersion = "v1beta"
const basePath = "https://www.googleapis.com/genomics/v1beta/"

// OAuth2 scopes used by this API.
const (
	// Manage your data in Google Cloud Storage
	DevstorageRead_writeScope = "https://www.googleapis.com/auth/devstorage.read_write"

	// View and manage Genomics data
	GenomicsScope = "https://www.googleapis.com/auth/genomics"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Beacons = NewBeaconsService(s)
	s.Callsets = NewCallsetsService(s)
	s.Datasets = NewDatasetsService(s)
	s.Experimental = NewExperimentalService(s)
	s.Jobs = NewJobsService(s)
	s.Reads = NewReadsService(s)
	s.Readsets = NewReadsetsService(s)
	s.Variants = NewVariantsService(s)
	return s, nil
}

type Service struct {
	client   *http.Client
	BasePath string // API endpoint base URL

	Beacons *BeaconsService

	Callsets *CallsetsService

	Datasets *DatasetsService

	Experimental *ExperimentalService

	Jobs *JobsService

	Reads *ReadsService

	Readsets *ReadsetsService

	Variants *VariantsService
}

func NewBeaconsService(s *Service) *BeaconsService {
	rs := &BeaconsService{s: s}
	return rs
}

type BeaconsService struct {
	s *Service
}

func NewCallsetsService(s *Service) *CallsetsService {
	rs := &CallsetsService{s: s}
	return rs
}

type CallsetsService struct {
	s *Service
}

func NewDatasetsService(s *Service) *DatasetsService {
	rs := &DatasetsService{s: s}
	return rs
}

type DatasetsService struct {
	s *Service
}

func NewExperimentalService(s *Service) *ExperimentalService {
	rs := &ExperimentalService{s: s}
	rs.Jobs = NewExperimentalJobsService(s)
	return rs
}

type ExperimentalService struct {
	s *Service

	Jobs *ExperimentalJobsService
}

func NewExperimentalJobsService(s *Service) *ExperimentalJobsService {
	rs := &ExperimentalJobsService{s: s}
	return rs
}

type ExperimentalJobsService struct {
	s *Service
}

func NewJobsService(s *Service) *JobsService {
	rs := &JobsService{s: s}
	return rs
}

type JobsService struct {
	s *Service
}

func NewReadsService(s *Service) *ReadsService {
	rs := &ReadsService{s: s}
	return rs
}

type ReadsService struct {
	s *Service
}

func NewReadsetsService(s *Service) *ReadsetsService {
	rs := &ReadsetsService{s: s}
	return rs
}

type ReadsetsService struct {
	s *Service
}

func NewVariantsService(s *Service) *VariantsService {
	rs := &VariantsService{s: s}
	return rs
}

type VariantsService struct {
	s *Service
}

type Beacon struct {
	// Exists: True if the allele exists on any variant call, false
	// otherwise.
	Exists bool `json:"exists,omitempty"`
}

type Call struct {
	// CallsetId: The ID of the callset this variant call belongs to.
	CallsetId string `json:"callsetId,omitempty"`

	// CallsetName: The name of the callset this variant call belongs to.
	CallsetName string `json:"callsetName,omitempty"`

	// Genotype: The genotype of this variant call. Each value represents
	// either the value of the referenceBases field or a 1-based index into
	// alternateBases. If a variant had a referenceBases field of "T" and an
	// alternateBases value of ["A", "C"], and the genotype was [2, 1], that
	// would mean the call represented the heterozygous value "CA" for this
	// variant. If the genotype was instead [0, 1], the represented value
	// would be "TA". Ordering of the genotype values is important if the
	// phaseset field is present.
	Genotype googleapi.Int64s `json:"genotype,omitempty"`

	// GenotypeLikelihood: The genotype likelihoods for this variant call.
	// Each array entry represents how likely a specific genotype is for
	// this call. The value ordering is defined by the GL tag in the VCF
	// spec.
	GenotypeLikelihood []float64 `json:"genotypeLikelihood,omitempty"`

	// Info: A map of additional variant call information.
	Info map[string][]string `json:"info,omitempty"`

	// Phaseset: If this field is present, this variant call's genotype
	// ordering implies the phase of the bases and is consistent with any
	// other variant calls on the same contig which have the same phaseset
	// value.
	Phaseset string `json:"phaseset,omitempty"`
}

type Callset struct {
	// Created: The date this callset was created in milliseconds from the
	// epoch.
	Created int64 `json:"created,omitempty,string"`

	// DatasetId: The ID of the dataset this callset belongs to.
	DatasetId string `json:"datasetId,omitempty"`

	// Id: The Google generated ID of the callset, immutable.
	Id string `json:"id,omitempty"`

	// Info: A map of additional callset information.
	Info map[string][]string `json:"info,omitempty"`

	// Name: The callset name.
	Name string `json:"name,omitempty"`
}

type ContigBound struct {
	// Contig: The contig the bound is associate with.
	Contig string `json:"contig,omitempty"`

	// UpperBound: An upper bound (inclusive) on the starting coordinate of
	// any variant in the contig.
	UpperBound int64 `json:"upperBound,omitempty,string"`
}

type Dataset struct {
	// Id: The Google generated ID of the dataset, immutable.
	Id string `json:"id,omitempty"`

	// IsPublic: Flag indicating whether or not a dataset is publicly
	// viewable. If a dataset is not public, it inherits viewing permissions
	// from its project.
	IsPublic bool `json:"isPublic,omitempty"`

	// Name: The dataset name.
	Name string `json:"name,omitempty"`

	// ProjectId: The Google Developers Console project number that this
	// dataset belongs to.
	ProjectId int64 `json:"projectId,omitempty,string"`
}

type ExperimentalCreateJobRequest struct {
	// Align: Specifies whether or not to run the alignment pipeline. At
	// least one of align or call_variants must be provided.
	Align bool `json:"align,omitempty"`

	// CallVariants: Specifies whether or not to run the variant calling
	// pipeline. If specified, alignment will be performed first and the
	// aligned BAMs will passed as input to the variant caller. At least one
	// of align or call_variants must be provided.
	CallVariants bool `json:"callVariants,omitempty"`

	// GcsOutputPath: Specifies where to copy the results of certain
	// pipelines. This shoud be in the form of "gs://bucket/path".
	GcsOutputPath string `json:"gcsOutputPath,omitempty"`

	// ProjectId: Required. The Google Cloud Project ID with which to
	// associate the request.
	ProjectId int64 `json:"projectId,omitempty,string"`

	// SourceUris: Required. A list of Google Cloud Storage URIs of data
	// files to operate upon.
	SourceUris []string `json:"sourceUris,omitempty"`
}

type ExperimentalCreateJobResponse struct {
	// JobId: A job ID that can be used to get status information.
	JobId string `json:"jobId,omitempty"`
}

type ExportReadsetsRequest struct {
	// ExportUri: A Google Cloud Storage URI where the exported BAM file
	// will be created. The currently authenticated user must have write
	// access to the new file location. An error will be returned if the URI
	// already contains data.
	ExportUri string `json:"exportUri,omitempty"`

	// ProjectId: The Google Developers Console project number that owns
	// this export. This is the project that will be billed.
	ProjectId int64 `json:"projectId,omitempty,string"`

	// ReadsetIds: The IDs of the readsets to export.
	ReadsetIds []string `json:"readsetIds,omitempty"`
}

type ExportReadsetsResponse struct {
	// JobId: A job ID that can be used to get status information.
	JobId string `json:"jobId,omitempty"`
}

type ExportVariantsRequest struct {
	// BigqueryDataset: The BigQuery dataset for export. Note this is
	// distinct from the Genomics concept of "dataset".
	BigqueryDataset string `json:"bigqueryDataset,omitempty"`

	// BigqueryTable: The BigQuery table for export.
	BigqueryTable string `json:"bigqueryTable,omitempty"`

	// CallsetIds: If provided, only variant call information from the
	// specified callsets will be exported. By default all variant calls are
	// exported.
	CallsetIds []string `json:"callsetIds,omitempty"`

	// DatasetIds: The IDs of the datasets that contain variant data which
	// should be exported. At least one dataset ID must be provided.
	DatasetIds []string `json:"datasetIds,omitempty"`

	// ExportUri: The Google Cloud URI to export to.
	ExportUri string `json:"exportUri,omitempty"`

	// Format: The format for the exported data.
	Format string `json:"format,omitempty"`

	// ProjectId: The Google Cloud project number that owns this export.
	// This is the project that will be billed.
	ProjectId int64 `json:"projectId,omitempty,string"`
}

type ExportVariantsResponse struct {
	// JobId: A job ID that can be used to get status information.
	JobId string `json:"jobId,omitempty"`
}

type GetVariantsSummaryResponse struct {
	// ContigBounds: A list of all contigs used by the variants in a dataset
	// with associated coordinate upper bounds for each one.
	ContigBounds []*ContigBound `json:"contigBounds,omitempty"`
}

type Header struct {
	// SortingOrder: (SO) Sorting order of alignments.
	SortingOrder string `json:"sortingOrder,omitempty"`

	// Version: (VN) BAM format version.
	Version string `json:"version,omitempty"`
}

type HeaderSection struct {
	// Comments: (@CO) One-line text comments.
	Comments []string `json:"comments,omitempty"`

	// FileUri: The file URI that this data was imported from.
	FileUri string `json:"fileUri,omitempty"`

	// Headers: (@HD) The header line.
	Headers []*Header `json:"headers,omitempty"`

	// Programs: (@PG) Programs.
	Programs []*Program `json:"programs,omitempty"`

	// ReadGroups: (@RG) Read group.
	ReadGroups []*ReadGroup `json:"readGroups,omitempty"`

	// RefSequences: (@SQ) Reference sequence dictionary.
	RefSequences []*ReferenceSequence `json:"refSequences,omitempty"`
}

type ImportReadsetsRequest struct {
	// DatasetId: Required. The ID of the dataset these readsets will belong
	// to.
	DatasetId string `json:"datasetId,omitempty"`

	// SourceUris: A list of URIs pointing at BAM or FASTQ files in Google
	// Cloud Storage.
	SourceUris []string `json:"sourceUris,omitempty"`
}

type ImportReadsetsResponse struct {
	// JobId: A job ID that can be used to get status information.
	JobId string `json:"jobId,omitempty"`
}

type ImportVariantsRequest struct {
	// DatasetId: Required. The dataset to which variant data should be
	// imported.
	DatasetId string `json:"datasetId,omitempty"`

	// SourceUris: A list of URIs pointing at VCF files in Google Cloud
	// Storage. See the VCF Specification for more details on the input
	// format.
	SourceUris []string `json:"sourceUris,omitempty"`
}

type ImportVariantsResponse struct {
	// JobId: A job ID that can be used to get status information.
	JobId string `json:"jobId,omitempty"`
}

type Job struct {
	// Description: A more detailed description of this job's current
	// status.
	Description string `json:"description,omitempty"`

	// Errors: Any errors that occurred during processing.
	Errors []string `json:"errors,omitempty"`

	// Id: The job ID.
	Id string `json:"id,omitempty"`

	// ImportedIds: If this Job represents an import, this field will
	// contain the IDs of the objects that were successfully imported.
	ImportedIds []string `json:"importedIds,omitempty"`

	// ProjectId: The Google Developers Console project number to which this
	// job belongs.
	ProjectId int64 `json:"projectId,omitempty,string"`

	// Status: The status of this job.
	Status string `json:"status,omitempty"`

	// Warnings: Any warnings that occurred during processing.
	Warnings []string `json:"warnings,omitempty"`
}

type ListDatasetsResponse struct {
	// Datasets: The list of matching Datasets.
	Datasets []*Dataset `json:"datasets,omitempty"`

	// NextPageToken: The continuation token, which is used to page through
	// large result sets. Provide this value in a subsequent request to
	// return the next page of results. This field will be empty if there
	// aren't any additional results.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type Program struct {
	// CommandLine: (CL) Command line.
	CommandLine string `json:"commandLine,omitempty"`

	// Id: (ID) Program record identifier.
	Id string `json:"id,omitempty"`

	// Name: (PN) Program name.
	Name string `json:"name,omitempty"`

	// PrevProgramId: (PP) Previous program ID.
	PrevProgramId string `json:"prevProgramId,omitempty"`

	// Version: (VN) Program version.
	Version string `json:"version,omitempty"`
}

type Read struct {
	// AlignedBases: The originalBases after the cigar field has been
	// applied. Deletions are represented with '-' and insertions are
	// omitted.
	AlignedBases string `json:"alignedBases,omitempty"`

	// BaseQuality: Represents the quality of each base in this read. Each
	// character represents one base. To get the quality, take the ASCII
	// value of the character and subtract 33. (QUAL)
	BaseQuality string `json:"baseQuality,omitempty"`

	// Cigar: A condensed representation of how this read matches up to the
	// reference. (CIGAR)
	Cigar string `json:"cigar,omitempty"`

	// Flags: Each bit of this number has a different meaning if enabled.
	// See the full BAM spec for more details. (FLAG)
	Flags int64 `json:"flags,omitempty"`

	// Id: The Google generated ID of the read, immutable.
	Id string `json:"id,omitempty"`

	// MappingQuality: A score up to 255 that represents how likely this
	// read's aligned position is to be correct. A higher value is better.
	// (MAPQ)
	MappingQuality int64 `json:"mappingQuality,omitempty"`

	// MatePosition: The 1-based start position of the paired read. (PNEXT)
	MatePosition int64 `json:"matePosition,omitempty"`

	// MateReferenceSequenceName: The name of the sequence that the paired
	// read is aligned to. This is usually the same as
	// referenceSequenceName. (RNEXT)
	MateReferenceSequenceName string `json:"mateReferenceSequenceName,omitempty"`

	// Name: The name of the read. When imported from a BAM file, this is
	// the query template name. (QNAME)
	Name string `json:"name,omitempty"`

	// OriginalBases: The list of bases that this read represents (such as
	// "CATCGA"). (SEQ)
	OriginalBases string `json:"originalBases,omitempty"`

	// Position: The 1-based start position of the aligned read. If the
	// first base starts at the very beginning of the reference sequence,
	// then the position would be '1'. (POS)
	Position int64 `json:"position,omitempty"`

	// ReadsetId: The ID of the readset this read belongs to.
	ReadsetId string `json:"readsetId,omitempty"`

	// ReferenceSequenceName: The name of the sequence that this read is
	// aligned to. This would be, for example, 'X' for the X Chromosome or
	// '20' for Chromosome 20. (RNAME)
	ReferenceSequenceName string `json:"referenceSequenceName,omitempty"`

	// Tags: A map of additional read information. (TAG)
	Tags map[string][]string `json:"tags,omitempty"`

	// TemplateLength: Length of the original piece of DNA that produced
	// both this read and the paired read. (TLEN)
	TemplateLength int64 `json:"templateLength,omitempty"`
}

type ReadGroup struct {
	// Date: (DT) Date the run was produced (ISO8601 date or date/time).
	Date string `json:"date,omitempty"`

	// Description: (DS) Description.
	Description string `json:"description,omitempty"`

	// FlowOrder: (FO) Flow order. The array of nucleotide bases that
	// correspond to the nucleotides used for each flow of each read.
	FlowOrder string `json:"flowOrder,omitempty"`

	// Id: (ID) Read group identifier.
	Id string `json:"id,omitempty"`

	// KeySequence: (KS) The array of nucleotide bases that correspond to
	// the key sequence of each read.
	KeySequence string `json:"keySequence,omitempty"`

	// Library: (LS) Library.
	Library string `json:"library,omitempty"`

	// PlatformUnit: (PU) Platform unit.
	PlatformUnit string `json:"platformUnit,omitempty"`

	// PredictedInsertSize: (PI) Predicted median insert size.
	PredictedInsertSize int64 `json:"predictedInsertSize,omitempty"`

	// ProcessingProgram: (PG) Programs used for processing the read group.
	ProcessingProgram string `json:"processingProgram,omitempty"`

	// Sample: (SM) Sample.
	Sample string `json:"sample,omitempty"`

	// SequencingCenterName: (CN) Name of sequencing center producing the
	// read.
	SequencingCenterName string `json:"sequencingCenterName,omitempty"`

	// SequencingTechnology: (PL) Platform/technology used to produce the
	// reads.
	SequencingTechnology string `json:"sequencingTechnology,omitempty"`
}

type Readset struct {
	// Created: The date this readset was created, in milliseconds from the
	// epoch.
	Created int64 `json:"created,omitempty,string"`

	// DatasetId: The ID of the dataset this readset belongs to.
	DatasetId string `json:"datasetId,omitempty"`

	// FileData: File information from the original BAM import. See the BAM
	// format specification for additional information on each field.
	FileData []*HeaderSection `json:"fileData,omitempty"`

	// Id: The Google generated ID of the readset, immutable.
	Id string `json:"id,omitempty"`

	// Name: The readset name, typically this is the sample name.
	Name string `json:"name,omitempty"`

	// ReadCount: The number of reads in this readset.
	ReadCount uint64 `json:"readCount,omitempty,string"`
}

type ReferenceSequence struct {
	// AssemblyId: (AS) Genome assembly identifier.
	AssemblyId string `json:"assemblyId,omitempty"`

	// Length: (LN) Reference sequence length.
	Length int64 `json:"length,omitempty"`

	// Md5Checksum: (M5) MD5 checksum of the sequence in the uppercase,
	// excluding spaces but including pads as *.
	Md5Checksum string `json:"md5Checksum,omitempty"`

	// Name: (SN) Reference sequence name.
	Name string `json:"name,omitempty"`

	// Species: (SP) Species.
	Species string `json:"species,omitempty"`

	// Uri: (UR) URI of the sequence.
	Uri string `json:"uri,omitempty"`
}

type SearchCallsetsRequest struct {
	// DatasetIds: Restrict the query to callsets within the given datasets.
	// At least one ID must be provided.
	DatasetIds []string `json:"datasetIds,omitempty"`

	// MaxResults: The maximum number of callsets to return.
	MaxResults uint64 `json:"maxResults,omitempty,string"`

	// Name: Only return callsets for which a substring of the name matches
	// this string.
	Name string `json:"name,omitempty"`

	// PageToken: The continuation token, which is used to page through
	// large result sets. To get the next page of results, set this
	// parameter to the value of "nextPageToken" from the previous response.
	PageToken string `json:"pageToken,omitempty"`
}

type SearchCallsetsResponse struct {
	// Callsets: The list of matching callsets.
	Callsets []*Callset `json:"callsets,omitempty"`

	// NextPageToken: The continuation token, which is used to page through
	// large result sets. Provide this value in a subsequent request to
	// return the next page of results. This field will be empty if there
	// aren't any additional results.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type SearchReadsRequest struct {
	// PageToken: The continuation token, which is used to page through
	// large result sets. To get the next page of results, set this
	// parameter to the value of "nextPageToken" from the previous response.
	PageToken string `json:"pageToken,omitempty"`

	// ReadsetIds: If specified, will restrict this query to reads within
	// the given readsets. At least one readset ID must be provided.
	ReadsetIds []string `json:"readsetIds,omitempty"`

	// SequenceEnd: The end position (1-based, inclusive) of this query. If
	// a sequence name is specified, this defaults to the sequence's length.
	SequenceEnd uint64 `json:"sequenceEnd,omitempty,string"`

	// SequenceName: The sequence to query. (e.g. 'X' for the X chromosome)
	// Leaving this blank returns results from all sequences, including
	// unmapped reads.
	SequenceName string `json:"sequenceName,omitempty"`

	// SequenceStart: The start position (1-based, inclusive) of this query.
	// If a sequence name is specified, this defaults to 1.
	SequenceStart uint64 `json:"sequenceStart,omitempty,string"`
}

type SearchReadsResponse struct {
	// NextPageToken: The continuation token, which is used to page through
	// large result sets. Provide this value in a subsequent request to
	// return the next page of results. This field will be empty if there
	// aren't any additional results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Reads: The list of matching Reads. The resulting Reads are sorted by
	// position; the smallest positions are returned first. Unmapped reads,
	// which have no position, are returned last and are further sorted
	// alphabetically by name.
	Reads []*Read `json:"reads,omitempty"`
}

type SearchReadsetsRequest struct {
	// DatasetIds: Restricts this query to readsets within the given
	// datasets. At least one ID must be provided.
	DatasetIds []string `json:"datasetIds,omitempty"`

	// Name: Only return readsets for which a substring of the name matches
	// this string.
	Name string `json:"name,omitempty"`

	// PageToken: The continuation token, which is used to page through
	// large result sets. To get the next page of results, set this
	// parameter to the value of "nextPageToken" from the previous response.
	PageToken string `json:"pageToken,omitempty"`
}

type SearchReadsetsResponse struct {
	// NextPageToken: The continuation token, which is used to page through
	// large result sets. Provide this value in a subsequent request to
	// return the next page of results. This field will be empty if there
	// aren't any additional results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Readsets: The list of matching Readsets.
	Readsets []*Readset `json:"readsets,omitempty"`
}

type SearchVariantsRequest struct {
	// CallsetIds: Only return variant calls which belong to callsets with
	// these ids. Leaving this blank returns all variant calls. At most one
	// of callsetNames or callsetIds should be provided.
	CallsetIds []string `json:"callsetIds,omitempty"`

	// CallsetNames: Only return variant calls which belong to callsets
	// which have exactly these names. Leaving this blank returns all
	// variant calls. At most one of callsetNames or callsetIds should be
	// provided.
	CallsetNames []string `json:"callsetNames,omitempty"`

	// Contig: Required. Only return variants on this contig.
	Contig string `json:"contig,omitempty"`

	// DatasetId: Required. The ID of the dataset to search.
	DatasetId string `json:"datasetId,omitempty"`

	// EndPosition: Required. The end of the window (1-based, inclusive) for
	// which overlapping variants should be returned.
	EndPosition int64 `json:"endPosition,omitempty,string"`

	// MaxResults: The maximum number of variants to return.
	MaxResults uint64 `json:"maxResults,omitempty,string"`

	// PageToken: The continuation token, which is used to page through
	// large result sets. To get the next page of results, set this
	// parameter to the value of "nextPageToken" from the previous response.
	PageToken string `json:"pageToken,omitempty"`

	// StartPosition: Required. The beginning of the window (1-based,
	// inclusive) for which overlapping variants should be returned.
	StartPosition int64 `json:"startPosition,omitempty,string"`

	// VariantName: Only return variants which have exactly this name.
	VariantName string `json:"variantName,omitempty"`
}

type SearchVariantsResponse struct {
	// NextPageToken: The continuation token, which is used to page through
	// large result sets. Provide this value in a subsequent request to
	// return the next page of results. This field will be empty if there
	// aren't any additional results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Variants: The list of matching Variants.
	Variants []*Variant `json:"variants,omitempty"`
}

type Variant struct {
	// AlternateBases: The bases that appear instead of the reference bases.
	AlternateBases []string `json:"alternateBases,omitempty"`

	// Calls: The variant calls for this particular variant. Each one
	// represents the determination of genotype with respect to this
	// variant.
	Calls []*Call `json:"calls,omitempty"`

	// Contig: The contig on which this variant occurs. (such as 'chr20' or
	// 'X')
	Contig string `json:"contig,omitempty"`

	// Created: The date this variant was created, in milliseconds from the
	// epoch.
	Created int64 `json:"created,omitempty,string"`

	// DatasetId: The ID of the dataset this variant belongs to.
	DatasetId string `json:"datasetId,omitempty"`

	// Id: The Google generated ID of the variant, immutable.
	Id string `json:"id,omitempty"`

	// Info: A map of additional variant information.
	Info map[string][]string `json:"info,omitempty"`

	// Names: Names for the variant, for example a RefSNP ID.
	Names []string `json:"names,omitempty"`

	// Position: The position at which this variant occurs (1-based). This
	// corresponds to the first base of the string of reference bases.
	Position int64 `json:"position,omitempty,string"`

	// ReferenceBases: The reference bases for this variant. They start at
	// the given position.
	ReferenceBases string `json:"referenceBases,omitempty"`
}

// method id "genomics.beacons.get":

type BeaconsGetCall struct {
	s         *Service
	datasetId string
	opt_      map[string]interface{}
}

// Get: This is an experimental API that provides a Global Alliance for
// Genomics and Health Beacon. It may change at any time.
func (r *BeaconsService) Get(datasetId string) *BeaconsGetCall {
	c := &BeaconsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.datasetId = datasetId
	return c
}

// Allele sets the optional parameter "allele": Required. The allele to
// look for ('A', 'C', 'G' or 'T').
func (c *BeaconsGetCall) Allele(allele string) *BeaconsGetCall {
	c.opt_["allele"] = allele
	return c
}

// Contig sets the optional parameter "contig": Required. The contig to
// query over.
func (c *BeaconsGetCall) Contig(contig string) *BeaconsGetCall {
	c.opt_["contig"] = contig
	return c
}

// Position sets the optional parameter "position": Required. The
// 1-based position to query at.
func (c *BeaconsGetCall) Position(position int64) *BeaconsGetCall {
	c.opt_["position"] = position
	return c
}

func (c *BeaconsGetCall) Do() (*Beacon, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["allele"]; ok {
		params.Set("allele", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["contig"]; ok {
		params.Set("contig", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["position"]; ok {
		params.Set("position", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "beacons/{datasetId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{datasetId}", url.QueryEscape(c.datasetId), 1)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Beacon
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "This is an experimental API that provides a Global Alliance for Genomics and Health Beacon. It may change at any time.",
	//   "httpMethod": "GET",
	//   "id": "genomics.beacons.get",
	//   "parameterOrder": [
	//     "datasetId"
	//   ],
	//   "parameters": {
	//     "allele": {
	//       "description": "Required. The allele to look for ('A', 'C', 'G' or 'T').",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "contig": {
	//       "description": "Required. The contig to query over.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "datasetId": {
	//       "description": "The ID of the dataset to query over. It must be public. Private datasets will return an unauthorized exception.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "position": {
	//       "description": "Required. The 1-based position to query at.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "beacons/{datasetId}",
	//   "response": {
	//     "$ref": "Beacon"
	//   }
	// }

}

// method id "genomics.callsets.create":

type CallsetsCreateCall struct {
	s       *Service
	callset *Callset
	opt_    map[string]interface{}
}

// Create: Creates a new callset.
func (r *CallsetsService) Create(callset *Callset) *CallsetsCreateCall {
	c := &CallsetsCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.callset = callset
	return c
}

func (c *CallsetsCreateCall) Do() (*Callset, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.callset)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative(c.s.BasePath, "callsets")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Callset
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a new callset.",
	//   "httpMethod": "POST",
	//   "id": "genomics.callsets.create",
	//   "path": "callsets",
	//   "request": {
	//     "$ref": "Callset"
	//   },
	//   "response": {
	//     "$ref": "Callset"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.callsets.delete":

type CallsetsDeleteCall struct {
	s         *Service
	callsetId string
	opt_      map[string]interface{}
}

// Delete: Deletes a callset.
func (r *CallsetsService) Delete(callsetId string) *CallsetsDeleteCall {
	c := &CallsetsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.callsetId = callsetId
	return c
}

func (c *CallsetsDeleteCall) Do() error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative(c.s.BasePath, "callsets/{callsetId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{callsetId}", url.QueryEscape(c.callsetId), 1)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Deletes a callset.",
	//   "httpMethod": "DELETE",
	//   "id": "genomics.callsets.delete",
	//   "parameterOrder": [
	//     "callsetId"
	//   ],
	//   "parameters": {
	//     "callsetId": {
	//       "description": "The ID of the callset to be deleted.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "callsets/{callsetId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.callsets.get":

type CallsetsGetCall struct {
	s         *Service
	callsetId string
	opt_      map[string]interface{}
}

// Get: Gets a callset by ID.
func (r *CallsetsService) Get(callsetId string) *CallsetsGetCall {
	c := &CallsetsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.callsetId = callsetId
	return c
}

func (c *CallsetsGetCall) Do() (*Callset, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative(c.s.BasePath, "callsets/{callsetId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{callsetId}", url.QueryEscape(c.callsetId), 1)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Callset
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a callset by ID.",
	//   "httpMethod": "GET",
	//   "id": "genomics.callsets.get",
	//   "parameterOrder": [
	//     "callsetId"
	//   ],
	//   "parameters": {
	//     "callsetId": {
	//       "description": "The ID of the callset.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "callsets/{callsetId}",
	//   "response": {
	//     "$ref": "Callset"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.callsets.patch":

type CallsetsPatchCall struct {
	s         *Service
	callsetId string
	callset   *Callset
	opt_      map[string]interface{}
}

// Patch: Updates a callset. This method supports patch semantics.
func (r *CallsetsService) Patch(callsetId string, callset *Callset) *CallsetsPatchCall {
	c := &CallsetsPatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.callsetId = callsetId
	c.callset = callset
	return c
}

func (c *CallsetsPatchCall) Do() (*Callset, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.callset)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative(c.s.BasePath, "callsets/{callsetId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{callsetId}", url.QueryEscape(c.callsetId), 1)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Callset
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates a callset. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "genomics.callsets.patch",
	//   "parameterOrder": [
	//     "callsetId"
	//   ],
	//   "parameters": {
	//     "callsetId": {
	//       "description": "The ID of the callset to be updated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "callsets/{callsetId}",
	//   "request": {
	//     "$ref": "Callset"
	//   },
	//   "response": {
	//     "$ref": "Callset"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.callsets.search":

type CallsetsSearchCall struct {
	s                     *Service
	searchcallsetsrequest *SearchCallsetsRequest
	opt_                  map[string]interface{}
}

// Search: Gets a list of callsets matching the criteria.
func (r *CallsetsService) Search(searchcallsetsrequest *SearchCallsetsRequest) *CallsetsSearchCall {
	c := &CallsetsSearchCall{s: r.s, opt_: make(map[string]interface{})}
	c.searchcallsetsrequest = searchcallsetsrequest
	return c
}

func (c *CallsetsSearchCall) Do() (*SearchCallsetsResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.searchcallsetsrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative(c.s.BasePath, "callsets/search")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *SearchCallsetsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a list of callsets matching the criteria.",
	//   "httpMethod": "POST",
	//   "id": "genomics.callsets.search",
	//   "path": "callsets/search",
	//   "request": {
	//     "$ref": "SearchCallsetsRequest"
	//   },
	//   "response": {
	//     "$ref": "SearchCallsetsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.callsets.update":

type CallsetsUpdateCall struct {
	s         *Service
	callsetId string
	callset   *Callset
	opt_      map[string]interface{}
}

// Update: Updates a callset.
func (r *CallsetsService) Update(callsetId string, callset *Callset) *CallsetsUpdateCall {
	c := &CallsetsUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.callsetId = callsetId
	c.callset = callset
	return c
}

func (c *CallsetsUpdateCall) Do() (*Callset, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.callset)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative(c.s.BasePath, "callsets/{callsetId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{callsetId}", url.QueryEscape(c.callsetId), 1)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Callset
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates a callset.",
	//   "httpMethod": "PUT",
	//   "id": "genomics.callsets.update",
	//   "parameterOrder": [
	//     "callsetId"
	//   ],
	//   "parameters": {
	//     "callsetId": {
	//       "description": "The ID of the callset to be updated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "callsets/{callsetId}",
	//   "request": {
	//     "$ref": "Callset"
	//   },
	//   "response": {
	//     "$ref": "Callset"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.datasets.create":

type DatasetsCreateCall struct {
	s       *Service
	dataset *Dataset
	opt_    map[string]interface{}
}

// Create: Creates a new dataset.
func (r *DatasetsService) Create(dataset *Dataset) *DatasetsCreateCall {
	c := &DatasetsCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.dataset = dataset
	return c
}

func (c *DatasetsCreateCall) Do() (*Dataset, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.dataset)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative(c.s.BasePath, "datasets")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Dataset
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a new dataset.",
	//   "httpMethod": "POST",
	//   "id": "genomics.datasets.create",
	//   "path": "datasets",
	//   "request": {
	//     "$ref": "Dataset"
	//   },
	//   "response": {
	//     "$ref": "Dataset"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.datasets.delete":

type DatasetsDeleteCall struct {
	s         *Service
	datasetId string
	opt_      map[string]interface{}
}

// Delete: Deletes a dataset.
func (r *DatasetsService) Delete(datasetId string) *DatasetsDeleteCall {
	c := &DatasetsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.datasetId = datasetId
	return c
}

func (c *DatasetsDeleteCall) Do() error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative(c.s.BasePath, "datasets/{datasetId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{datasetId}", url.QueryEscape(c.datasetId), 1)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Deletes a dataset.",
	//   "httpMethod": "DELETE",
	//   "id": "genomics.datasets.delete",
	//   "parameterOrder": [
	//     "datasetId"
	//   ],
	//   "parameters": {
	//     "datasetId": {
	//       "description": "The ID of the dataset to be deleted.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "datasets/{datasetId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.datasets.get":

type DatasetsGetCall struct {
	s         *Service
	datasetId string
	opt_      map[string]interface{}
}

// Get: Gets a dataset by ID.
func (r *DatasetsService) Get(datasetId string) *DatasetsGetCall {
	c := &DatasetsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.datasetId = datasetId
	return c
}

func (c *DatasetsGetCall) Do() (*Dataset, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative(c.s.BasePath, "datasets/{datasetId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{datasetId}", url.QueryEscape(c.datasetId), 1)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Dataset
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a dataset by ID.",
	//   "httpMethod": "GET",
	//   "id": "genomics.datasets.get",
	//   "parameterOrder": [
	//     "datasetId"
	//   ],
	//   "parameters": {
	//     "datasetId": {
	//       "description": "The ID of the dataset.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "datasets/{datasetId}",
	//   "response": {
	//     "$ref": "Dataset"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.datasets.list":

type DatasetsListCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// List: Lists all datasets.
func (r *DatasetsService) List() *DatasetsListCall {
	c := &DatasetsListCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of results returned by this request.
func (c *DatasetsListCall) MaxResults(maxResults uint64) *DatasetsListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// PageToken sets the optional parameter "pageToken": The continuation
// token, which is used to page through large result sets. To get the
// next page of results, set this parameter to the value of
// "nextPageToken" from the previous response.
func (c *DatasetsListCall) PageToken(pageToken string) *DatasetsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// ProjectId sets the optional parameter "projectId": Only return
// datasets which belong to this Google Developers Console project.
func (c *DatasetsListCall) ProjectId(projectId int64) *DatasetsListCall {
	c.opt_["projectId"] = projectId
	return c
}

func (c *DatasetsListCall) Do() (*ListDatasetsResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["projectId"]; ok {
		params.Set("projectId", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "datasets")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ListDatasetsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists all datasets.",
	//   "httpMethod": "GET",
	//   "id": "genomics.datasets.list",
	//   "parameters": {
	//     "maxResults": {
	//       "default": "50",
	//       "description": "The maximum number of results returned by this request.",
	//       "format": "uint64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "The continuation token, which is used to page through large result sets. To get the next page of results, set this parameter to the value of \"nextPageToken\" from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "Only return datasets which belong to this Google Developers Console project.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "datasets",
	//   "response": {
	//     "$ref": "ListDatasetsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.datasets.patch":

type DatasetsPatchCall struct {
	s         *Service
	datasetId string
	dataset   *Dataset
	opt_      map[string]interface{}
}

// Patch: Updates a dataset. This method supports patch semantics.
func (r *DatasetsService) Patch(datasetId string, dataset *Dataset) *DatasetsPatchCall {
	c := &DatasetsPatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.datasetId = datasetId
	c.dataset = dataset
	return c
}

func (c *DatasetsPatchCall) Do() (*Dataset, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.dataset)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative(c.s.BasePath, "datasets/{datasetId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{datasetId}", url.QueryEscape(c.datasetId), 1)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Dataset
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates a dataset. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "genomics.datasets.patch",
	//   "parameterOrder": [
	//     "datasetId"
	//   ],
	//   "parameters": {
	//     "datasetId": {
	//       "description": "The ID of the dataset to be updated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "datasets/{datasetId}",
	//   "request": {
	//     "$ref": "Dataset"
	//   },
	//   "response": {
	//     "$ref": "Dataset"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.datasets.update":

type DatasetsUpdateCall struct {
	s         *Service
	datasetId string
	dataset   *Dataset
	opt_      map[string]interface{}
}

// Update: Updates a dataset.
func (r *DatasetsService) Update(datasetId string, dataset *Dataset) *DatasetsUpdateCall {
	c := &DatasetsUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.datasetId = datasetId
	c.dataset = dataset
	return c
}

func (c *DatasetsUpdateCall) Do() (*Dataset, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.dataset)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative(c.s.BasePath, "datasets/{datasetId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{datasetId}", url.QueryEscape(c.datasetId), 1)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Dataset
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates a dataset.",
	//   "httpMethod": "PUT",
	//   "id": "genomics.datasets.update",
	//   "parameterOrder": [
	//     "datasetId"
	//   ],
	//   "parameters": {
	//     "datasetId": {
	//       "description": "The ID of the dataset to be updated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "datasets/{datasetId}",
	//   "request": {
	//     "$ref": "Dataset"
	//   },
	//   "response": {
	//     "$ref": "Dataset"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.experimental.jobs.create":

type ExperimentalJobsCreateCall struct {
	s                            *Service
	experimentalcreatejobrequest *ExperimentalCreateJobRequest
	opt_                         map[string]interface{}
}

// Create: Creates and asynchronously runs an ad-hoc job. This is an
// experimental call and may be removed or changed at any time.
func (r *ExperimentalJobsService) Create(experimentalcreatejobrequest *ExperimentalCreateJobRequest) *ExperimentalJobsCreateCall {
	c := &ExperimentalJobsCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.experimentalcreatejobrequest = experimentalcreatejobrequest
	return c
}

func (c *ExperimentalJobsCreateCall) Do() (*ExperimentalCreateJobResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.experimentalcreatejobrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative(c.s.BasePath, "experimental/jobs/create")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ExperimentalCreateJobResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates and asynchronously runs an ad-hoc job. This is an experimental call and may be removed or changed at any time.",
	//   "httpMethod": "POST",
	//   "id": "genomics.experimental.jobs.create",
	//   "path": "experimental/jobs/create",
	//   "request": {
	//     "$ref": "ExperimentalCreateJobRequest"
	//   },
	//   "response": {
	//     "$ref": "ExperimentalCreateJobResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/devstorage.read_write",
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.jobs.get":

type JobsGetCall struct {
	s     *Service
	jobId string
	opt_  map[string]interface{}
}

// Get: Gets a job by ID.
func (r *JobsService) Get(jobId string) *JobsGetCall {
	c := &JobsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.jobId = jobId
	return c
}

func (c *JobsGetCall) Do() (*Job, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative(c.s.BasePath, "jobs/{jobId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{jobId}", url.QueryEscape(c.jobId), 1)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
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
	//   "description": "Gets a job by ID.",
	//   "httpMethod": "GET",
	//   "id": "genomics.jobs.get",
	//   "parameterOrder": [
	//     "jobId"
	//   ],
	//   "parameters": {
	//     "jobId": {
	//       "description": "The ID of the job.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "jobs/{jobId}",
	//   "response": {
	//     "$ref": "Job"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.reads.get":

type ReadsGetCall struct {
	s      *Service
	readId string
	opt_   map[string]interface{}
}

// Get: Gets a read by ID.
func (r *ReadsService) Get(readId string) *ReadsGetCall {
	c := &ReadsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.readId = readId
	return c
}

func (c *ReadsGetCall) Do() (*Read, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative(c.s.BasePath, "reads/{readId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{readId}", url.QueryEscape(c.readId), 1)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Read
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a read by ID.",
	//   "httpMethod": "GET",
	//   "id": "genomics.reads.get",
	//   "parameterOrder": [
	//     "readId"
	//   ],
	//   "parameters": {
	//     "readId": {
	//       "description": "The ID of the read.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "reads/{readId}",
	//   "response": {
	//     "$ref": "Read"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.reads.search":

type ReadsSearchCall struct {
	s                  *Service
	searchreadsrequest *SearchReadsRequest
	opt_               map[string]interface{}
}

// Search: Gets a list of reads matching the criteria.
func (r *ReadsService) Search(searchreadsrequest *SearchReadsRequest) *ReadsSearchCall {
	c := &ReadsSearchCall{s: r.s, opt_: make(map[string]interface{})}
	c.searchreadsrequest = searchreadsrequest
	return c
}

func (c *ReadsSearchCall) Do() (*SearchReadsResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.searchreadsrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative(c.s.BasePath, "reads/search")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *SearchReadsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a list of reads matching the criteria.",
	//   "httpMethod": "POST",
	//   "id": "genomics.reads.search",
	//   "path": "reads/search",
	//   "request": {
	//     "$ref": "SearchReadsRequest"
	//   },
	//   "response": {
	//     "$ref": "SearchReadsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.readsets.create":

type ReadsetsCreateCall struct {
	s       *Service
	readset *Readset
	opt_    map[string]interface{}
}

// Create: Creates a new readset.
func (r *ReadsetsService) Create(readset *Readset) *ReadsetsCreateCall {
	c := &ReadsetsCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.readset = readset
	return c
}

func (c *ReadsetsCreateCall) Do() (*Readset, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.readset)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative(c.s.BasePath, "readsets")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Readset
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a new readset.",
	//   "httpMethod": "POST",
	//   "id": "genomics.readsets.create",
	//   "path": "readsets",
	//   "request": {
	//     "$ref": "Readset"
	//   },
	//   "response": {
	//     "$ref": "Readset"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.readsets.delete":

type ReadsetsDeleteCall struct {
	s         *Service
	readsetId string
	opt_      map[string]interface{}
}

// Delete: Deletes a readset.
func (r *ReadsetsService) Delete(readsetId string) *ReadsetsDeleteCall {
	c := &ReadsetsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.readsetId = readsetId
	return c
}

func (c *ReadsetsDeleteCall) Do() error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative(c.s.BasePath, "readsets/{readsetId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{readsetId}", url.QueryEscape(c.readsetId), 1)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Deletes a readset.",
	//   "httpMethod": "DELETE",
	//   "id": "genomics.readsets.delete",
	//   "parameterOrder": [
	//     "readsetId"
	//   ],
	//   "parameters": {
	//     "readsetId": {
	//       "description": "The ID of the readset to be deleted.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "readsets/{readsetId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.readsets.export":

type ReadsetsExportCall struct {
	s                     *Service
	exportreadsetsrequest *ExportReadsetsRequest
	opt_                  map[string]interface{}
}

// Export: Exports readsets to a file.
func (r *ReadsetsService) Export(exportreadsetsrequest *ExportReadsetsRequest) *ReadsetsExportCall {
	c := &ReadsetsExportCall{s: r.s, opt_: make(map[string]interface{})}
	c.exportreadsetsrequest = exportreadsetsrequest
	return c
}

func (c *ReadsetsExportCall) Do() (*ExportReadsetsResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.exportreadsetsrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative(c.s.BasePath, "readsets/export")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ExportReadsetsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Exports readsets to a file.",
	//   "httpMethod": "POST",
	//   "id": "genomics.readsets.export",
	//   "path": "readsets/export",
	//   "request": {
	//     "$ref": "ExportReadsetsRequest"
	//   },
	//   "response": {
	//     "$ref": "ExportReadsetsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/devstorage.read_write",
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.readsets.get":

type ReadsetsGetCall struct {
	s         *Service
	readsetId string
	opt_      map[string]interface{}
}

// Get: Gets a readset by ID.
func (r *ReadsetsService) Get(readsetId string) *ReadsetsGetCall {
	c := &ReadsetsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.readsetId = readsetId
	return c
}

func (c *ReadsetsGetCall) Do() (*Readset, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative(c.s.BasePath, "readsets/{readsetId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{readsetId}", url.QueryEscape(c.readsetId), 1)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Readset
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a readset by ID.",
	//   "httpMethod": "GET",
	//   "id": "genomics.readsets.get",
	//   "parameterOrder": [
	//     "readsetId"
	//   ],
	//   "parameters": {
	//     "readsetId": {
	//       "description": "The ID of the readset.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "readsets/{readsetId}",
	//   "response": {
	//     "$ref": "Readset"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.readsets.import":

type ReadsetsImportCall struct {
	s                     *Service
	importreadsetsrequest *ImportReadsetsRequest
	opt_                  map[string]interface{}
}

// Import: Creates readsets by asynchronously importing the provided
// information.
func (r *ReadsetsService) Import(importreadsetsrequest *ImportReadsetsRequest) *ReadsetsImportCall {
	c := &ReadsetsImportCall{s: r.s, opt_: make(map[string]interface{})}
	c.importreadsetsrequest = importreadsetsrequest
	return c
}

func (c *ReadsetsImportCall) Do() (*ImportReadsetsResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.importreadsetsrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative(c.s.BasePath, "readsets/import")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ImportReadsetsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates readsets by asynchronously importing the provided information.",
	//   "httpMethod": "POST",
	//   "id": "genomics.readsets.import",
	//   "path": "readsets/import",
	//   "request": {
	//     "$ref": "ImportReadsetsRequest"
	//   },
	//   "response": {
	//     "$ref": "ImportReadsetsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/devstorage.read_write",
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.readsets.patch":

type ReadsetsPatchCall struct {
	s         *Service
	readsetId string
	readset   *Readset
	opt_      map[string]interface{}
}

// Patch: Updates a readset. This method supports patch semantics.
func (r *ReadsetsService) Patch(readsetId string, readset *Readset) *ReadsetsPatchCall {
	c := &ReadsetsPatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.readsetId = readsetId
	c.readset = readset
	return c
}

func (c *ReadsetsPatchCall) Do() (*Readset, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.readset)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative(c.s.BasePath, "readsets/{readsetId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{readsetId}", url.QueryEscape(c.readsetId), 1)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Readset
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates a readset. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "genomics.readsets.patch",
	//   "parameterOrder": [
	//     "readsetId"
	//   ],
	//   "parameters": {
	//     "readsetId": {
	//       "description": "The ID of the readset to be updated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "readsets/{readsetId}",
	//   "request": {
	//     "$ref": "Readset"
	//   },
	//   "response": {
	//     "$ref": "Readset"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.readsets.search":

type ReadsetsSearchCall struct {
	s                     *Service
	searchreadsetsrequest *SearchReadsetsRequest
	opt_                  map[string]interface{}
}

// Search: Gets a list of readsets matching the criteria.
func (r *ReadsetsService) Search(searchreadsetsrequest *SearchReadsetsRequest) *ReadsetsSearchCall {
	c := &ReadsetsSearchCall{s: r.s, opt_: make(map[string]interface{})}
	c.searchreadsetsrequest = searchreadsetsrequest
	return c
}

func (c *ReadsetsSearchCall) Do() (*SearchReadsetsResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.searchreadsetsrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative(c.s.BasePath, "readsets/search")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *SearchReadsetsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a list of readsets matching the criteria.",
	//   "httpMethod": "POST",
	//   "id": "genomics.readsets.search",
	//   "path": "readsets/search",
	//   "request": {
	//     "$ref": "SearchReadsetsRequest"
	//   },
	//   "response": {
	//     "$ref": "SearchReadsetsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.readsets.update":

type ReadsetsUpdateCall struct {
	s         *Service
	readsetId string
	readset   *Readset
	opt_      map[string]interface{}
}

// Update: Updates a readset.
func (r *ReadsetsService) Update(readsetId string, readset *Readset) *ReadsetsUpdateCall {
	c := &ReadsetsUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.readsetId = readsetId
	c.readset = readset
	return c
}

func (c *ReadsetsUpdateCall) Do() (*Readset, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.readset)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative(c.s.BasePath, "readsets/{readsetId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{readsetId}", url.QueryEscape(c.readsetId), 1)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Readset
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates a readset.",
	//   "httpMethod": "PUT",
	//   "id": "genomics.readsets.update",
	//   "parameterOrder": [
	//     "readsetId"
	//   ],
	//   "parameters": {
	//     "readsetId": {
	//       "description": "The ID of the readset to be updated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "readsets/{readsetId}",
	//   "request": {
	//     "$ref": "Readset"
	//   },
	//   "response": {
	//     "$ref": "Readset"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.variants.create":

type VariantsCreateCall struct {
	s       *Service
	variant *Variant
	opt_    map[string]interface{}
}

// Create: Creates a new variant.
func (r *VariantsService) Create(variant *Variant) *VariantsCreateCall {
	c := &VariantsCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.variant = variant
	return c
}

func (c *VariantsCreateCall) Do() (*Variant, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.variant)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative(c.s.BasePath, "variants")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Variant
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a new variant.",
	//   "httpMethod": "POST",
	//   "id": "genomics.variants.create",
	//   "path": "variants",
	//   "request": {
	//     "$ref": "Variant"
	//   },
	//   "response": {
	//     "$ref": "Variant"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.variants.delete":

type VariantsDeleteCall struct {
	s         *Service
	variantId string
	opt_      map[string]interface{}
}

// Delete: Deletes a variant.
func (r *VariantsService) Delete(variantId string) *VariantsDeleteCall {
	c := &VariantsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.variantId = variantId
	return c
}

func (c *VariantsDeleteCall) Do() error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative(c.s.BasePath, "variants/{variantId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{variantId}", url.QueryEscape(c.variantId), 1)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Deletes a variant.",
	//   "httpMethod": "DELETE",
	//   "id": "genomics.variants.delete",
	//   "parameterOrder": [
	//     "variantId"
	//   ],
	//   "parameters": {
	//     "variantId": {
	//       "description": "The ID of the variant to be deleted.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "variants/{variantId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.variants.export":

type VariantsExportCall struct {
	s                     *Service
	exportvariantsrequest *ExportVariantsRequest
	opt_                  map[string]interface{}
}

// Export: Exports variant data to an external destination.
func (r *VariantsService) Export(exportvariantsrequest *ExportVariantsRequest) *VariantsExportCall {
	c := &VariantsExportCall{s: r.s, opt_: make(map[string]interface{})}
	c.exportvariantsrequest = exportvariantsrequest
	return c
}

func (c *VariantsExportCall) Do() (*ExportVariantsResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.exportvariantsrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative(c.s.BasePath, "variants/export")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ExportVariantsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Exports variant data to an external destination.",
	//   "httpMethod": "POST",
	//   "id": "genomics.variants.export",
	//   "path": "variants/export",
	//   "request": {
	//     "$ref": "ExportVariantsRequest"
	//   },
	//   "response": {
	//     "$ref": "ExportVariantsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.variants.get":

type VariantsGetCall struct {
	s         *Service
	variantId string
	opt_      map[string]interface{}
}

// Get: Gets a variant by ID.
func (r *VariantsService) Get(variantId string) *VariantsGetCall {
	c := &VariantsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.variantId = variantId
	return c
}

func (c *VariantsGetCall) Do() (*Variant, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative(c.s.BasePath, "variants/{variantId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{variantId}", url.QueryEscape(c.variantId), 1)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Variant
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a variant by ID.",
	//   "httpMethod": "GET",
	//   "id": "genomics.variants.get",
	//   "parameterOrder": [
	//     "variantId"
	//   ],
	//   "parameters": {
	//     "variantId": {
	//       "description": "The ID of the variant.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "variants/{variantId}",
	//   "response": {
	//     "$ref": "Variant"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.variants.getSummary":

type VariantsGetSummaryCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// GetSummary: Gets a summary of all the variant data in a dataset.
func (r *VariantsService) GetSummary() *VariantsGetSummaryCall {
	c := &VariantsGetSummaryCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

// DatasetId sets the optional parameter "datasetId": Required. The ID
// of the dataset to get variant summary information for.
func (c *VariantsGetSummaryCall) DatasetId(datasetId string) *VariantsGetSummaryCall {
	c.opt_["datasetId"] = datasetId
	return c
}

func (c *VariantsGetSummaryCall) Do() (*GetVariantsSummaryResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["datasetId"]; ok {
		params.Set("datasetId", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "variants/summary")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *GetVariantsSummaryResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a summary of all the variant data in a dataset.",
	//   "httpMethod": "GET",
	//   "id": "genomics.variants.getSummary",
	//   "parameters": {
	//     "datasetId": {
	//       "description": "Required. The ID of the dataset to get variant summary information for.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "variants/summary",
	//   "response": {
	//     "$ref": "GetVariantsSummaryResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.variants.import":

type VariantsImportCall struct {
	s                     *Service
	importvariantsrequest *ImportVariantsRequest
	opt_                  map[string]interface{}
}

// Import: Creates variant data by asynchronously importing the provided
// information.
func (r *VariantsService) Import(importvariantsrequest *ImportVariantsRequest) *VariantsImportCall {
	c := &VariantsImportCall{s: r.s, opt_: make(map[string]interface{})}
	c.importvariantsrequest = importvariantsrequest
	return c
}

func (c *VariantsImportCall) Do() (*ImportVariantsResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.importvariantsrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative(c.s.BasePath, "variants/import")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ImportVariantsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates variant data by asynchronously importing the provided information.",
	//   "httpMethod": "POST",
	//   "id": "genomics.variants.import",
	//   "path": "variants/import",
	//   "request": {
	//     "$ref": "ImportVariantsRequest"
	//   },
	//   "response": {
	//     "$ref": "ImportVariantsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/devstorage.read_write",
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.variants.patch":

type VariantsPatchCall struct {
	s         *Service
	variantId string
	variant   *Variant
	opt_      map[string]interface{}
}

// Patch: Updates a variant. This method supports patch semantics.
func (r *VariantsService) Patch(variantId string, variant *Variant) *VariantsPatchCall {
	c := &VariantsPatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.variantId = variantId
	c.variant = variant
	return c
}

func (c *VariantsPatchCall) Do() (*Variant, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.variant)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative(c.s.BasePath, "variants/{variantId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{variantId}", url.QueryEscape(c.variantId), 1)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Variant
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates a variant. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "genomics.variants.patch",
	//   "parameterOrder": [
	//     "variantId"
	//   ],
	//   "parameters": {
	//     "variantId": {
	//       "description": "The ID of the variant to be updated..",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "variants/{variantId}",
	//   "request": {
	//     "$ref": "Variant"
	//   },
	//   "response": {
	//     "$ref": "Variant"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.variants.search":

type VariantsSearchCall struct {
	s                     *Service
	searchvariantsrequest *SearchVariantsRequest
	opt_                  map[string]interface{}
}

// Search: Gets a list of variants matching the criteria.
func (r *VariantsService) Search(searchvariantsrequest *SearchVariantsRequest) *VariantsSearchCall {
	c := &VariantsSearchCall{s: r.s, opt_: make(map[string]interface{})}
	c.searchvariantsrequest = searchvariantsrequest
	return c
}

func (c *VariantsSearchCall) Do() (*SearchVariantsResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.searchvariantsrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative(c.s.BasePath, "variants/search")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *SearchVariantsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a list of variants matching the criteria.",
	//   "httpMethod": "POST",
	//   "id": "genomics.variants.search",
	//   "path": "variants/search",
	//   "request": {
	//     "$ref": "SearchVariantsRequest"
	//   },
	//   "response": {
	//     "$ref": "SearchVariantsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.variants.update":

type VariantsUpdateCall struct {
	s         *Service
	variantId string
	variant   *Variant
	opt_      map[string]interface{}
}

// Update: Updates a variant.
func (r *VariantsService) Update(variantId string, variant *Variant) *VariantsUpdateCall {
	c := &VariantsUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.variantId = variantId
	c.variant = variant
	return c
}

func (c *VariantsUpdateCall) Do() (*Variant, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.variant)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative(c.s.BasePath, "variants/{variantId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{variantId}", url.QueryEscape(c.variantId), 1)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Variant
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates a variant.",
	//   "httpMethod": "PUT",
	//   "id": "genomics.variants.update",
	//   "parameterOrder": [
	//     "variantId"
	//   ],
	//   "parameters": {
	//     "variantId": {
	//       "description": "The ID of the variant to be updated..",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "variants/{variantId}",
	//   "request": {
	//     "$ref": "Variant"
	//   },
	//   "response": {
	//     "$ref": "Variant"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}
