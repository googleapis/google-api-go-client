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
	// View and manage your data in Google BigQuery
	BigqueryScope = "https://www.googleapis.com/auth/bigquery"

	// Manage your data in Google Cloud Storage
	DevstorageRead_writeScope = "https://www.googleapis.com/auth/devstorage.read_write"

	// View and manage Genomics data
	GenomicsScope = "https://www.googleapis.com/auth/genomics"

	// View Genomics data
	GenomicsReadonlyScope = "https://www.googleapis.com/auth/genomics.readonly"
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
	rs.Coveragebuckets = NewReadsetsCoveragebucketsService(s)
	return rs
}

type ReadsetsService struct {
	s *Service

	Coveragebuckets *ReadsetsCoveragebucketsService
}

func NewReadsetsCoveragebucketsService(s *Service) *ReadsetsCoveragebucketsService {
	rs := &ReadsetsCoveragebucketsService{s: s}
	return rs
}

type ReadsetsCoveragebucketsService struct {
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
	// phaseset field is present. If a genotype is not called (that is, a
	// "." is present in the GT string) -1 is returned.
	Genotype []int64 `json:"genotype,omitempty"`

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
	// value. When importing data from VCF, if the genotype data was phased
	// but no phase set was specified this field will be set to "*".
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

type CoverageBucket struct {
	// MeanCoverage: The average number of reads which are aligned to each
	// individual reference base in this bucket.
	MeanCoverage float64 `json:"meanCoverage,omitempty"`

	// Range: The genomic coordinate range spanned by this bucket.
	Range *GenomicRange `json:"range,omitempty"`
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

	// PairedSourceUris: A list of Google Cloud Storage URIs of paired end
	// .fastq files to operate upon. If specified, this represents the
	// second file of each paired .fastq file. The first file of each pair
	// should be specified in "sourceUris".
	PairedSourceUris []string `json:"pairedSourceUris,omitempty"`

	// ProjectId: Required. The Google Cloud Project ID with which to
	// associate the request.
	ProjectId int64 `json:"projectId,omitempty,string"`

	// SourceUris: A list of Google Cloud Storage URIs of data files to
	// operate upon. These can be .bam, interleaved .fastq, or paired
	// .fastq. If specifying paired .fastq files, the first of each pair of
	// files should be listed here, and the second of each pair should be
	// listed in "pairedSourceUris".
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
	// BigqueryDataset: The BigQuery dataset to export data to. Note that
	// this is distinct from the Genomics concept of "dataset". The caller
	// must have WRITE access to this BigQuery dataset.
	BigqueryDataset string `json:"bigqueryDataset,omitempty"`

	// BigqueryTable: The BigQuery table to export data to. The caller must
	// have WRITE access to this BigQuery table.
	BigqueryTable string `json:"bigqueryTable,omitempty"`

	// CallsetIds: If provided, only variant call information from the
	// specified callsets will be exported. By default all variant calls are
	// exported.
	CallsetIds []string `json:"callsetIds,omitempty"`

	// DatasetId: Required. The ID of the dataset that contains variant data
	// which should be exported. The caller must have READ access to this
	// dataset.
	DatasetId string `json:"datasetId,omitempty"`

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

type GenomicRange struct {
	// SequenceEnd: The end position of the range on the reference, 1-based
	// exclusive. If specified, sequenceName must also be specified.
	SequenceEnd uint64 `json:"sequenceEnd,omitempty,string"`

	// SequenceName: The reference sequence name, for example "chr1", "1",
	// or "chrX".
	SequenceName string `json:"sequenceName,omitempty"`

	// SequenceStart: The start position of the range on the reference,
	// 1-based inclusive. If specified, sequenceName must also be specified.
	SequenceStart uint64 `json:"sequenceStart,omitempty,string"`
}

type GetVariantsSummaryResponse struct {
	// ContigBounds: A list of all contigs used by the variants in a dataset
	// with associated coordinate upper bounds for each one.
	ContigBounds []*ContigBound `json:"contigBounds,omitempty"`

	// Metadata: The metadata associated with this dataset.
	Metadata []*Metadata `json:"metadata,omitempty"`
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

	// FileUri: [Deprecated] This field is deprecated and will no longer be
	// populated. Please use filename instead.
	FileUri string `json:"fileUri,omitempty"`

	// Filename: The name of the file from which this data was imported.
	Filename string `json:"filename,omitempty"`

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
	// to. The caller must have WRITE permissions to this dataset.
	DatasetId string `json:"datasetId,omitempty"`

	// SourceUris: A list of URIs pointing at BAM files in Google Cloud
	// Storage.
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
	// Created: The date this job was created, in milliseconds from the
	// epoch.
	Created int64 `json:"created,omitempty,string"`

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

type ListCoverageBucketsResponse struct {
	// BucketWidth: The length of each coverage bucket in base pairs. Note
	// that buckets at the end of a reference sequence may be shorter. This
	// value is omitted if the bucket width is infinity (the default
	// behaviour, with no range or target_bucket_width).
	BucketWidth uint64 `json:"bucketWidth,omitempty,string"`

	// CoverageBuckets: The coverage buckets. The list of buckets is sparse;
	// a bucket with 0 overlapping reads is not returned. A bucket never
	// crosses more than one reference sequence. Each bucket has width
	// bucket_width, unless its end is is the end of the reference sequence.
	CoverageBuckets []*CoverageBucket `json:"coverageBuckets,omitempty"`

	// NextPageToken: The continuation token, which is used to page through
	// large result sets. Provide this value in a subsequent request to
	// return the next page of results. This field will be empty if there
	// aren't any additional results.
	NextPageToken string `json:"nextPageToken,omitempty"`
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

type Metadata struct {
	// Description: A textual description of this metadata.
	Description string `json:"description,omitempty"`

	// Id: User-provided ID field, not enforced by this API. Two or more
	// pieces of structured metadata with identical id and key fields are
	// considered equivalent.
	Id string `json:"id,omitempty"`

	// Info: Remaining structured metadata key-value pairs.
	Info map[string][]string `json:"info,omitempty"`

	// Key: The top-level key.
	Key string `json:"key,omitempty"`

	// Number: The number of values that can be included in a field
	// described by this metadata.
	Number string `json:"number,omitempty"`

	// Type: The type of data. Possible types include: Integer, Float, Flag,
	// Character, and String.
	Type string `json:"type,omitempty"`

	// Value: The value field for simple metadata
	Value string `json:"value,omitempty"`
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
	// DatasetId: The ID of the dataset this readset belongs to.
	DatasetId string `json:"datasetId,omitempty"`

	// FileData: File information from the original BAM import. See the BAM
	// format specification for additional information on each field.
	FileData []*HeaderSection `json:"fileData,omitempty"`

	// Id: The Google generated ID of the readset, immutable.
	Id string `json:"id,omitempty"`

	// Name: The readset name, typically this is the sample name.
	Name string `json:"name,omitempty"`
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

type SearchJobsRequest struct {
	// CreatedAfter: If specified, only jobs created on or after this date,
	// given in milliseconds since Unix epoch, will be returned.
	CreatedAfter int64 `json:"createdAfter,omitempty,string"`

	// CreatedBefore: If specified, only jobs created prior to this date,
	// given in milliseconds since Unix epoch, will be returned.
	CreatedBefore int64 `json:"createdBefore,omitempty,string"`

	// MaxResults: Specifies the number of results to return in a single
	// page. Defaults to 128. The maximum value is 256.
	MaxResults uint64 `json:"maxResults,omitempty,string"`

	// PageToken: The continuation token which is used to page through large
	// result sets. To get the next page of results, set this parameter to
	// the value of the "nextPageToken" from the previous response.
	PageToken string `json:"pageToken,omitempty"`

	// ProjectId: Required. Only return jobs which belong to this Google
	// Developers Console project. Only accepts project numbers.
	ProjectId int64 `json:"projectId,omitempty,string"`

	// Status: Only return jobs which have a matching status.
	Status []string `json:"status,omitempty"`
}

type SearchJobsResponse struct {
	// Jobs: The list of jobs results, ordered newest to oldest.
	Jobs []*Job `json:"jobs,omitempty"`

	// NextPageToken: The continuation token which is used to page through
	// large result sets. Provide this value is a subsequent request to
	// return the next page of results. This field will be empty if there
	// are no more results.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type SearchReadsRequest struct {
	// MaxResults: Specifies number of results to return in a single page.
	// If unspecified, it will default to 256. The maximum value is 1024.
	MaxResults uint64 `json:"maxResults,omitempty,string"`

	// PageToken: The continuation token, which is used to page through
	// large result sets. To get the next page of results, set this
	// parameter to the value of "nextPageToken" from the previous response.
	PageToken string `json:"pageToken,omitempty"`

	// ReadsetIds: The readsets within which to search for reads. At least
	// one readset ID must be provided. All specified readsets must be
	// aligned against a common set of reference sequences; this defines the
	// genomic coordinates for the query.
	ReadsetIds []string `json:"readsetIds,omitempty"`

	// SequenceEnd: The end position (1-based, inclusive) of the target
	// range. If specified, "sequenceName" must also be specified. Defaults
	// to the end of the target reference sequence, if any.
	SequenceEnd uint64 `json:"sequenceEnd,omitempty,string"`

	// SequenceName: Restricts the results to a particular reference
	// sequence such as '1', 'chr1', or 'X'. The set of valid references
	// sequences depends on the readsets specified. If set to "*", only
	// unmapped Reads are returned.
	SequenceName string `json:"sequenceName,omitempty"`

	// SequenceStart: The start position (1-based, inclusive) of the target
	// range. If specified, "sequenceName" must also be specified. Defaults
	// to the start of the target reference sequence, if any.
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

	// MaxResults: Specifies number of results to return in a single page.
	// If unspecified, it will default to 128. The maximum value is 256.
	MaxResults uint64 `json:"maxResults,omitempty,string"`

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

	// End: The end position (1-based) of this variant. This corresponds to
	// the first base after the last base in the reference allele. So, the
	// length of the reference allele is (end - position). This is useful
	// for variants that don't explicitly give alternate bases, for example
	// large deletions.
	End int64 `json:"end,omitempty,string"`

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
	googleapi.Expand(req.URL, map[string]string{
		"datasetId": c.datasetId,
	})
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
	googleapi.Expand(req.URL, map[string]string{
		"callsetId": c.callsetId,
	})
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
	googleapi.Expand(req.URL, map[string]string{
		"callsetId": c.callsetId,
	})
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
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
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
	googleapi.Expand(req.URL, map[string]string{
		"callsetId": c.callsetId,
	})
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
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
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
	googleapi.Expand(req.URL, map[string]string{
		"callsetId": c.callsetId,
	})
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
	googleapi.Expand(req.URL, map[string]string{
		"datasetId": c.datasetId,
	})
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
	googleapi.Expand(req.URL, map[string]string{
		"datasetId": c.datasetId,
	})
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
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
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
// datasets which belong to this Google Developers Console project. Only
// accepts project numbers.
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
	//       "description": "Only return datasets which belong to this Google Developers Console project. Only accepts project numbers.",
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
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
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
	googleapi.Expand(req.URL, map[string]string{
		"datasetId": c.datasetId,
	})
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

// method id "genomics.datasets.undelete":

type DatasetsUndeleteCall struct {
	s         *Service
	datasetId string
	opt_      map[string]interface{}
}

// Undelete: Undeletes a dataset by restoring a dataset which was
// deleted via this API. This operation is only possible for a week
// after the deletion occurred.
func (r *DatasetsService) Undelete(datasetId string) *DatasetsUndeleteCall {
	c := &DatasetsUndeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.datasetId = datasetId
	return c
}

func (c *DatasetsUndeleteCall) Do() (*Dataset, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative(c.s.BasePath, "datasets/{datasetId}/undelete")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"datasetId": c.datasetId,
	})
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
	//   "description": "Undeletes a dataset by restoring a dataset which was deleted via this API. This operation is only possible for a week after the deletion occurred.",
	//   "httpMethod": "POST",
	//   "id": "genomics.datasets.undelete",
	//   "parameterOrder": [
	//     "datasetId"
	//   ],
	//   "parameters": {
	//     "datasetId": {
	//       "description": "The ID of the dataset to be undeleted.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "datasets/{datasetId}/undelete",
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
	googleapi.Expand(req.URL, map[string]string{
		"datasetId": c.datasetId,
	})
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

// method id "genomics.jobs.cancel":

type JobsCancelCall struct {
	s     *Service
	jobId string
	opt_  map[string]interface{}
}

// Cancel: Cancels a job by ID. Note that it is possible for partial
// results to be generated and stored for cancelled jobs.
func (r *JobsService) Cancel(jobId string) *JobsCancelCall {
	c := &JobsCancelCall{s: r.s, opt_: make(map[string]interface{})}
	c.jobId = jobId
	return c
}

func (c *JobsCancelCall) Do() error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative(c.s.BasePath, "jobs/{jobId}/cancel")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"jobId": c.jobId,
	})
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
	//   "description": "Cancels a job by ID. Note that it is possible for partial results to be generated and stored for cancelled jobs.",
	//   "httpMethod": "POST",
	//   "id": "genomics.jobs.cancel",
	//   "parameterOrder": [
	//     "jobId"
	//   ],
	//   "parameters": {
	//     "jobId": {
	//       "description": "Required. The ID of the job.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "jobs/{jobId}/cancel",
	//   "scopes": [
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
	googleapi.Expand(req.URL, map[string]string{
		"jobId": c.jobId,
	})
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
	//       "description": "Required. The ID of the job.",
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
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
	//   ]
	// }

}

// method id "genomics.jobs.search":

type JobsSearchCall struct {
	s                 *Service
	searchjobsrequest *SearchJobsRequest
	opt_              map[string]interface{}
}

// Search: Gets a list of jobs matching the criteria.
func (r *JobsService) Search(searchjobsrequest *SearchJobsRequest) *JobsSearchCall {
	c := &JobsSearchCall{s: r.s, opt_: make(map[string]interface{})}
	c.searchjobsrequest = searchjobsrequest
	return c
}

func (c *JobsSearchCall) Do() (*SearchJobsResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.searchjobsrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative(c.s.BasePath, "jobs/search")
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
	var ret *SearchJobsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a list of jobs matching the criteria.",
	//   "httpMethod": "POST",
	//   "id": "genomics.jobs.search",
	//   "path": "jobs/search",
	//   "request": {
	//     "$ref": "SearchJobsRequest"
	//   },
	//   "response": {
	//     "$ref": "SearchJobsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
	//   ]
	// }

}

// method id "genomics.reads.search":

type ReadsSearchCall struct {
	s                  *Service
	searchreadsrequest *SearchReadsRequest
	opt_               map[string]interface{}
}

// Search: Gets a list of reads for one or more readsets. Reads search
// operates over a genomic coordinate space of reference sequence &
// position defined over the reference sequences to which the requested
// readsets are aligned. If a target positional range is specified,
// search returns all reads whose alignment to the reference genome
// overlap the range. A query which specifies only readset IDs yields
// all reads in those readsets, including unmapped reads. All reads
// returned (including reads on subsequent pages) are ordered by genomic
// coordinate (reference sequence & position). Reads with equivalent
// genomic coordinates are returned in a deterministic order.
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
	//   "description": "Gets a list of reads for one or more readsets. Reads search operates over a genomic coordinate space of reference sequence \u0026 position defined over the reference sequences to which the requested readsets are aligned. If a target positional range is specified, search returns all reads whose alignment to the reference genome overlap the range. A query which specifies only readset IDs yields all reads in those readsets, including unmapped reads. All reads returned (including reads on subsequent pages) are ordered by genomic coordinate (reference sequence \u0026 position). Reads with equivalent genomic coordinates are returned in a deterministic order.",
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
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
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
	googleapi.Expand(req.URL, map[string]string{
		"readsetId": c.readsetId,
	})
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
	//       "description": "The ID of the readset to be deleted. The caller must have WRITE permissions to the dataset associated with this readset.",
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

// Export: Exports readsets to a BAM file in Google Cloud Storage. Note
// that currently there may be some differences between exported BAM
// files and the original BAM file at the time of import. In particular,
// comments in the input file header will not be preserved, and some
// custom tags will be converted to strings.
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
	//   "description": "Exports readsets to a BAM file in Google Cloud Storage. Note that currently there may be some differences between exported BAM files and the original BAM file at the time of import. In particular, comments in the input file header will not be preserved, and some custom tags will be converted to strings.",
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
	googleapi.Expand(req.URL, map[string]string{
		"readsetId": c.readsetId,
	})
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
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
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
// information. Note that currently comments in the input file header
// are not imported and some custom tags will be converted to strings,
// rather than preserving tag types. The caller must have WRITE
// permissions to the dataset.
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
	//   "description": "Creates readsets by asynchronously importing the provided information. Note that currently comments in the input file header are not imported and some custom tags will be converted to strings, rather than preserving tag types. The caller must have WRITE permissions to the dataset.",
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
	googleapi.Expand(req.URL, map[string]string{
		"readsetId": c.readsetId,
	})
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
	//       "description": "The ID of the readset to be updated. The caller must have WRITE permissions to the dataset associated with this readset.",
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
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
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
	googleapi.Expand(req.URL, map[string]string{
		"readsetId": c.readsetId,
	})
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
	//       "description": "The ID of the readset to be updated. The caller must have WRITE permissions to the dataset associated with this readset.",
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

// method id "genomics.readsets.coveragebuckets.list":

type ReadsetsCoveragebucketsListCall struct {
	s         *Service
	readsetId string
	opt_      map[string]interface{}
}

// List: Lists fixed width coverage buckets for a readset, each of which
// correspond to a range of a reference sequence. Each bucket summarizes
// coverage information across its corresponding genomic range. Coverage
// is defined as the number of reads which are aligned to a given base
// in the reference sequence. Coverage buckets are available at various
// bucket widths, enabling various coverage "zoom levels". The caller
// must have READ permissions for the target readset.
func (r *ReadsetsCoveragebucketsService) List(readsetId string) *ReadsetsCoveragebucketsListCall {
	c := &ReadsetsCoveragebucketsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.readsetId = readsetId
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of results to return in a single page. If unspecified,
// defaults to 1024. The maximum value is 2048.
func (c *ReadsetsCoveragebucketsListCall) MaxResults(maxResults uint64) *ReadsetsCoveragebucketsListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// PageToken sets the optional parameter "pageToken": The continuation
// token, which is used to page through large result sets. To get the
// next page of results, set this parameter to the value of
// "nextPageToken" from the previous response.
func (c *ReadsetsCoveragebucketsListCall) PageToken(pageToken string) *ReadsetsCoveragebucketsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// RangeSequenceEnd sets the optional parameter "range.sequenceEnd": The
// end position of the range on the reference, 1-based exclusive. If
// specified, sequenceName must also be specified.
func (c *ReadsetsCoveragebucketsListCall) RangeSequenceEnd(rangeSequenceEnd uint64) *ReadsetsCoveragebucketsListCall {
	c.opt_["range.sequenceEnd"] = rangeSequenceEnd
	return c
}

// RangeSequenceName sets the optional parameter "range.sequenceName":
// The reference sequence name, for example "chr1", "1", or "chrX".
func (c *ReadsetsCoveragebucketsListCall) RangeSequenceName(rangeSequenceName string) *ReadsetsCoveragebucketsListCall {
	c.opt_["range.sequenceName"] = rangeSequenceName
	return c
}

// RangeSequenceStart sets the optional parameter "range.sequenceStart":
// The start position of the range on the reference, 1-based inclusive.
// If specified, sequenceName must also be specified.
func (c *ReadsetsCoveragebucketsListCall) RangeSequenceStart(rangeSequenceStart uint64) *ReadsetsCoveragebucketsListCall {
	c.opt_["range.sequenceStart"] = rangeSequenceStart
	return c
}

// TargetBucketWidth sets the optional parameter "targetBucketWidth":
// The desired width of each reported coverage bucket in base pairs.
// This will be rounded down to the nearest precomputed bucket width;
// the value of which is returned as bucket_width in the response.
// Defaults to infinity (each bucket spans an entire reference sequence)
// or the length of the target range, if specified. The smallest
// precomputed bucket_width is currently 2048 base pairs; this is
// subject to change.
func (c *ReadsetsCoveragebucketsListCall) TargetBucketWidth(targetBucketWidth uint64) *ReadsetsCoveragebucketsListCall {
	c.opt_["targetBucketWidth"] = targetBucketWidth
	return c
}

func (c *ReadsetsCoveragebucketsListCall) Do() (*ListCoverageBucketsResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["range.sequenceEnd"]; ok {
		params.Set("range.sequenceEnd", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["range.sequenceName"]; ok {
		params.Set("range.sequenceName", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["range.sequenceStart"]; ok {
		params.Set("range.sequenceStart", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["targetBucketWidth"]; ok {
		params.Set("targetBucketWidth", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "readsets/{readsetId}/coveragebuckets")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"readsetId": c.readsetId,
	})
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ListCoverageBucketsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists fixed width coverage buckets for a readset, each of which correspond to a range of a reference sequence. Each bucket summarizes coverage information across its corresponding genomic range. Coverage is defined as the number of reads which are aligned to a given base in the reference sequence. Coverage buckets are available at various bucket widths, enabling various coverage \"zoom levels\". The caller must have READ permissions for the target readset.",
	//   "httpMethod": "GET",
	//   "id": "genomics.readsets.coveragebuckets.list",
	//   "parameterOrder": [
	//     "readsetId"
	//   ],
	//   "parameters": {
	//     "maxResults": {
	//       "default": "1024",
	//       "description": "The maximum number of results to return in a single page. If unspecified, defaults to 1024. The maximum value is 2048.",
	//       "format": "uint64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "The continuation token, which is used to page through large result sets. To get the next page of results, set this parameter to the value of \"nextPageToken\" from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "range.sequenceEnd": {
	//       "description": "The end position of the range on the reference, 1-based exclusive. If specified, sequenceName must also be specified.",
	//       "format": "uint64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "range.sequenceName": {
	//       "description": "The reference sequence name, for example \"chr1\", \"1\", or \"chrX\".",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "range.sequenceStart": {
	//       "description": "The start position of the range on the reference, 1-based inclusive. If specified, sequenceName must also be specified.",
	//       "format": "uint64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "readsetId": {
	//       "description": "Required. The ID of the readset over which coverage is requested.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "targetBucketWidth": {
	//       "description": "The desired width of each reported coverage bucket in base pairs. This will be rounded down to the nearest precomputed bucket width; the value of which is returned as bucket_width in the response. Defaults to infinity (each bucket spans an entire reference sequence) or the length of the target range, if specified. The smallest precomputed bucket_width is currently 2048 base pairs; this is subject to change.",
	//       "format": "uint64",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "readsets/{readsetId}/coveragebuckets",
	//   "response": {
	//     "$ref": "ListCoverageBucketsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
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
	googleapi.Expand(req.URL, map[string]string{
		"variantId": c.variantId,
	})
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
	//     "https://www.googleapis.com/auth/bigquery",
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
	googleapi.Expand(req.URL, map[string]string{
		"variantId": c.variantId,
	})
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
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
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
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
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
	googleapi.Expand(req.URL, map[string]string{
		"variantId": c.variantId,
	})
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
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
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
	googleapi.Expand(req.URL, map[string]string{
		"variantId": c.variantId,
	})
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
