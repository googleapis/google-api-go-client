// Package prediction provides access to the Prediction API.
//
// See http://code.google.com/apis/predict/docs/developer-guide.html
//
// Usage example:
//
//   import "google-api-go-client.googlecode.com/hg/prediction/v1.4"
//   ...
//   predictionService, err := prediction.New(oauthHttpClient)
package prediction

import (
	"bytes"
	"fmt"
	"http"
	"io"
	"json"
	"os"
	"strings"
	"strconv"
	"url"
	"google-api-go-client.googlecode.com/hg/google-api"
)

var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = googleapi.Version

const apiId = "prediction:v1.4"
const apiName = "prediction"
const apiVersion = "v1.4"
const basePath = "https://www.googleapis.com/prediction/v1.4/"

// OAuth2 scopes used by this API.
const (
	// Manage your data in the Google Prediction API
	PredictionScope = "https://www.googleapis.com/auth/prediction"
)

func New(client *http.Client) (*Service, os.Error) {
	if client == nil {
		return nil, os.NewError("client is nil")
	}
	s := &Service{client: client}
	s.Hostedmodels = &HostedmodelsService{s: s}
	s.Trainedmodels = &TrainedmodelsService{s: s}
	return s, nil
}

type Service struct {
	client *http.Client

	Hostedmodels *HostedmodelsService

	Trainedmodels *TrainedmodelsService
}

type HostedmodelsService struct {
	s *Service
}

type TrainedmodelsService struct {
	s *Service
}

type InputInput struct {
	// CsvInstance: A list of input features, these can be strings or
	// doubles.
	CsvInstance []interface{} `json:"csvInstance,omitempty"`
}

type Output struct {
	// SelfLink: A URL to re-request this resource.
	SelfLink string `json:"selfLink,omitempty"`

	// OutputValue: The estimated regression value [Regression models only].
	OutputValue float64 `json:"outputValue,omitempty"`

	// OutputLabel: The most likely class label [Categorical models only].
	OutputLabel string `json:"outputLabel,omitempty"`

	// Kind: What kind of resource this is.
	Kind string `json:"kind,omitempty"`

	// Id: The unique name for the predictive model.
	Id string `json:"id,omitempty"`

	// OutputMulti: A list of class labels with their estimated
	// probabilities [Categorical models only].
	OutputMulti []*OutputOutputMulti `json:"outputMulti,omitempty"`
}

type TrainingDataAnalysis struct {
	Warnings []string `json:"warnings,omitempty"`
}

type OutputOutputMulti struct {
	// Score: The probability of the class label.
	Score float64 `json:"score,omitempty"`

	// Label: The class label.
	Label string `json:"label,omitempty"`
}

type Training struct {
	// Utility: A class weighting function, which allows the importance
	// weights for class labels to be specified [Categorical models only].
	Utility []*TrainingUtility `json:"utility,omitempty"`

	// SelfLink: A URL to re-request this resource.
	SelfLink string `json:"selfLink,omitempty"`

	// StorageDataLocation: Google storage location of the training data
	// file.
	StorageDataLocation string `json:"storageDataLocation,omitempty"`

	// DataAnalysis: Data Analysis.
	DataAnalysis *TrainingDataAnalysis `json:"dataAnalysis,omitempty"`

	// ModelInfo: Model metadata.
	ModelInfo *TrainingModelInfo `json:"modelInfo,omitempty"`

	// Kind: What kind of resource this is.
	Kind string `json:"kind,omitempty"`

	// TrainingStatus: The current status of the training job. This can be
	// one of following: RUNNING; DONE; ERROR; ERROR: TRAINING JOB NOT FOUND
	TrainingStatus string `json:"trainingStatus,omitempty"`

	// Id: The unique name for the predictive model.
	Id string `json:"id,omitempty"`

	// StoragePMMLLocation: Google storage location of the preprocessing
	// pmml file.
	StoragePMMLLocation string `json:"storagePMMLLocation,omitempty"`
}

type Input struct {
	// Input: Input to the model for a prediction
	Input *InputInput `json:"input,omitempty"`
}

type Update struct {
	// Label: The true class label of this instance
	Label string `json:"label,omitempty"`

	// CsvInstance: The input features for this instance
	CsvInstance []interface{} `json:"csvInstance,omitempty"`
}

type TrainingModelInfoConfusionMatrix struct {

}

type TrainingUtility struct {

}

type TrainingModelInfo struct {
	// ConfusionMatrixRowTotals: A list of the confusion matrix row totals
	ConfusionMatrixRowTotals *TrainingModelInfoConfusionMatrixRowTotals `json:"confusionMatrixRowTotals,omitempty"`

	// NumberInstances: Number of valid data instances used in the trained
	// model.
	NumberInstances int64 `json:"numberInstances,omitempty,string"`

	// ConfusionMatrix: An output confusion matrix. This shows an estimate
	// for how this model will do in predictions. This is first indexed by
	// the true class label. For each true class label, this provides a pair
	// {predicted_label, count}, where count is the estimated number of
	// times the model will predict the predicted label given the true label
	// [Categorical models only].
	ConfusionMatrix *TrainingModelInfoConfusionMatrix `json:"confusionMatrix,omitempty"`

	// MeanSquaredError: An estimated mean squared error. The can be used to
	// measure the quality of the predicted model [Regression models only].
	MeanSquaredError float64 `json:"meanSquaredError,omitempty"`

	// ModelType: Type of predictive model (CLASSIFICATION or REGRESSION)
	ModelType string `json:"modelType,omitempty"`

	// NumberLabels: Number of class labels in the trained model
	// [Categorical models only].
	NumberLabels int64 `json:"numberLabels,omitempty,string"`

	// ClassWeightedAccuracy: Estimated accuracy of model taking utility
	// weights into account [Categorical models only].
	ClassWeightedAccuracy float64 `json:"classWeightedAccuracy,omitempty"`

	// ClassificationAccuracy: A number between 0.0 and 1.0, where 1.0 is
	// 100% accurate. This is an estimate, based on the amount and quality
	// of the training data, of the estimated prediction accuracy. You can
	// use this is a guide to decide whether the results are accurate enough
	// for your needs. This estimate will be more reliable if your real
	// input data is similar to your training data [Categorical models
	// only].
	ClassificationAccuracy float64 `json:"classificationAccuracy,omitempty"`
}

type TrainingModelInfoConfusionMatrixRowTotals struct {

}

// method id "prediction.hostedmodels.predict":

type HostedmodelsPredictCall struct {
	s               *Service
	hostedModelName string
	input           *Input
	opt_            map[string]interface{}
}

// Predict: Submit input and request an output against a hosted model.
func (r *HostedmodelsService) Predict(hostedModelName string, input *Input) *HostedmodelsPredictCall {
	c := &HostedmodelsPredictCall{s: r.s, opt_: make(map[string]interface{})}
	c.hostedModelName = hostedModelName
	c.input = input
	return c
}

func (c *HostedmodelsPredictCall) Do() (*Output, os.Error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.input)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/prediction/v1.4/", "hostedmodels/{hostedModelName}/predict")
	urls = strings.Replace(urls, "{hostedModelName}", cleanPathString(c.hostedModelName), 1)
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(Output)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Submit input and request an output against a hosted model.",
	//   "httpMethod": "POST",
	//   "id": "prediction.hostedmodels.predict",
	//   "parameterOrder": [
	//     "hostedModelName"
	//   ],
	//   "parameters": {
	//     "hostedModelName": {
	//       "description": "The name of a hosted model.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "hostedmodels/{hostedModelName}/predict",
	//   "request": {
	//     "$ref": "Input"
	//   },
	//   "response": {
	//     "$ref": "Output"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/prediction"
	//   ]
	// }

}

// method id "prediction.trainedmodels.update":

type TrainedmodelsUpdateCall struct {
	s      *Service
	id     string
	update *Update
	opt_   map[string]interface{}
}

// Update: Add new data to a trained model.
func (r *TrainedmodelsService) Update(id string, update *Update) *TrainedmodelsUpdateCall {
	c := &TrainedmodelsUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.id = id
	c.update = update
	return c
}

func (c *TrainedmodelsUpdateCall) Do() (*Training, os.Error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.update)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/prediction/v1.4/", "trainedmodels/{id}")
	urls = strings.Replace(urls, "{id}", cleanPathString(c.id), 1)
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(Training)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Add new data to a trained model.",
	//   "httpMethod": "PUT",
	//   "id": "prediction.trainedmodels.update",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The unique name for the predictive model.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "trainedmodels/{id}",
	//   "request": {
	//     "$ref": "Update"
	//   },
	//   "response": {
	//     "$ref": "Training"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/prediction"
	//   ]
	// }

}

// method id "prediction.trainedmodels.insert":

type TrainedmodelsInsertCall struct {
	s        *Service
	training *Training
	opt_     map[string]interface{}
}

// Insert: Begin training your model.
func (r *TrainedmodelsService) Insert(training *Training) *TrainedmodelsInsertCall {
	c := &TrainedmodelsInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.training = training
	return c
}

func (c *TrainedmodelsInsertCall) Do() (*Training, os.Error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.training)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/prediction/v1.4/", "trainedmodels")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(Training)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Begin training your model.",
	//   "httpMethod": "POST",
	//   "id": "prediction.trainedmodels.insert",
	//   "path": "trainedmodels",
	//   "request": {
	//     "$ref": "Training"
	//   },
	//   "response": {
	//     "$ref": "Training"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/prediction"
	//   ]
	// }

}

// method id "prediction.trainedmodels.get":

type TrainedmodelsGetCall struct {
	s    *Service
	id   string
	opt_ map[string]interface{}
}

// Get: Check training status of your model.
func (r *TrainedmodelsService) Get(id string) *TrainedmodelsGetCall {
	c := &TrainedmodelsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.id = id
	return c
}

func (c *TrainedmodelsGetCall) Do() (*Training, os.Error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/prediction/v1.4/", "trainedmodels/{id}")
	urls = strings.Replace(urls, "{id}", cleanPathString(c.id), 1)
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(Training)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Check training status of your model.",
	//   "httpMethod": "GET",
	//   "id": "prediction.trainedmodels.get",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The unique name for the predictive model.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "trainedmodels/{id}",
	//   "response": {
	//     "$ref": "Training"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/prediction"
	//   ]
	// }

}

// method id "prediction.trainedmodels.delete":

type TrainedmodelsDeleteCall struct {
	s    *Service
	id   string
	opt_ map[string]interface{}
}

// Delete: Delete a trained model.
func (r *TrainedmodelsService) Delete(id string) *TrainedmodelsDeleteCall {
	c := &TrainedmodelsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.id = id
	return c
}

func (c *TrainedmodelsDeleteCall) Do() os.Error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/prediction/v1.4/", "trainedmodels/{id}")
	urls = strings.Replace(urls, "{id}", cleanPathString(c.id), 1)
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return err
	}
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Delete a trained model.",
	//   "httpMethod": "DELETE",
	//   "id": "prediction.trainedmodels.delete",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The unique name for the predictive model.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "trainedmodels/{id}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/prediction"
	//   ]
	// }

}

// method id "prediction.trainedmodels.predict":

type TrainedmodelsPredictCall struct {
	s     *Service
	id    string
	input *Input
	opt_  map[string]interface{}
}

// Predict: Submit model id and request a prediction
func (r *TrainedmodelsService) Predict(id string, input *Input) *TrainedmodelsPredictCall {
	c := &TrainedmodelsPredictCall{s: r.s, opt_: make(map[string]interface{})}
	c.id = id
	c.input = input
	return c
}

func (c *TrainedmodelsPredictCall) Do() (*Output, os.Error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.input)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/prediction/v1.4/", "trainedmodels/{id}/predict")
	urls = strings.Replace(urls, "{id}", cleanPathString(c.id), 1)
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(Output)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Submit model id and request a prediction",
	//   "httpMethod": "POST",
	//   "id": "prediction.trainedmodels.predict",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The unique name for the predictive model.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "trainedmodels/{id}/predict",
	//   "request": {
	//     "$ref": "Input"
	//   },
	//   "response": {
	//     "$ref": "Output"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/prediction"
	//   ]
	// }

}

func cleanPathString(s string) string {
	return strings.Map(func(r int) int {
		if r >= 0x30 && r <= 0x7a {
			return r
		}
		return -1
	}, s)
}
