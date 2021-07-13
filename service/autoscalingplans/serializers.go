// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscalingplans

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/autoscalingplans/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/encoding/httpbinding"
	smithyjson "github.com/aws/smithy-go/encoding/json"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"math"
)

type awsAwsjson11_serializeOpCreateScalingPlan struct {
}

func (*awsAwsjson11_serializeOpCreateScalingPlan) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpCreateScalingPlan) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*CreateScalingPlanInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AnyScaleScalingPlannerFrontendService.CreateScalingPlan")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentCreateScalingPlanInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpDeleteScalingPlan struct {
}

func (*awsAwsjson11_serializeOpDeleteScalingPlan) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpDeleteScalingPlan) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DeleteScalingPlanInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AnyScaleScalingPlannerFrontendService.DeleteScalingPlan")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentDeleteScalingPlanInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpDescribeScalingPlanResources struct {
}

func (*awsAwsjson11_serializeOpDescribeScalingPlanResources) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpDescribeScalingPlanResources) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DescribeScalingPlanResourcesInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AnyScaleScalingPlannerFrontendService.DescribeScalingPlanResources")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentDescribeScalingPlanResourcesInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpDescribeScalingPlans struct {
}

func (*awsAwsjson11_serializeOpDescribeScalingPlans) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpDescribeScalingPlans) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DescribeScalingPlansInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AnyScaleScalingPlannerFrontendService.DescribeScalingPlans")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentDescribeScalingPlansInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpGetScalingPlanResourceForecastData struct {
}

func (*awsAwsjson11_serializeOpGetScalingPlanResourceForecastData) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpGetScalingPlanResourceForecastData) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetScalingPlanResourceForecastDataInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AnyScaleScalingPlannerFrontendService.GetScalingPlanResourceForecastData")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentGetScalingPlanResourceForecastDataInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpUpdateScalingPlan struct {
}

func (*awsAwsjson11_serializeOpUpdateScalingPlan) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpUpdateScalingPlan) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*UpdateScalingPlanInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AnyScaleScalingPlannerFrontendService.UpdateScalingPlan")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentUpdateScalingPlanInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsAwsjson11_serializeDocumentApplicationSource(v *types.ApplicationSource, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.CloudFormationStackARN != nil {
		ok := object.Key("CloudFormationStackARN")
		ok.String(*v.CloudFormationStackARN)
	}

	if v.TagFilters != nil {
		ok := object.Key("TagFilters")
		if err := awsAwsjson11_serializeDocumentTagFilters(v.TagFilters, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsAwsjson11_serializeDocumentApplicationSources(v []types.ApplicationSource, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsAwsjson11_serializeDocumentApplicationSource(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsAwsjson11_serializeDocumentCustomizedLoadMetricSpecification(v *types.CustomizedLoadMetricSpecification, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Dimensions != nil {
		ok := object.Key("Dimensions")
		if err := awsAwsjson11_serializeDocumentMetricDimensions(v.Dimensions, ok); err != nil {
			return err
		}
	}

	if v.MetricName != nil {
		ok := object.Key("MetricName")
		ok.String(*v.MetricName)
	}

	if v.Namespace != nil {
		ok := object.Key("Namespace")
		ok.String(*v.Namespace)
	}

	if len(v.Statistic) > 0 {
		ok := object.Key("Statistic")
		ok.String(string(v.Statistic))
	}

	if v.Unit != nil {
		ok := object.Key("Unit")
		ok.String(*v.Unit)
	}

	return nil
}

func awsAwsjson11_serializeDocumentCustomizedScalingMetricSpecification(v *types.CustomizedScalingMetricSpecification, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Dimensions != nil {
		ok := object.Key("Dimensions")
		if err := awsAwsjson11_serializeDocumentMetricDimensions(v.Dimensions, ok); err != nil {
			return err
		}
	}

	if v.MetricName != nil {
		ok := object.Key("MetricName")
		ok.String(*v.MetricName)
	}

	if v.Namespace != nil {
		ok := object.Key("Namespace")
		ok.String(*v.Namespace)
	}

	if len(v.Statistic) > 0 {
		ok := object.Key("Statistic")
		ok.String(string(v.Statistic))
	}

	if v.Unit != nil {
		ok := object.Key("Unit")
		ok.String(*v.Unit)
	}

	return nil
}

func awsAwsjson11_serializeDocumentMetricDimension(v *types.MetricDimension, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Name != nil {
		ok := object.Key("Name")
		ok.String(*v.Name)
	}

	if v.Value != nil {
		ok := object.Key("Value")
		ok.String(*v.Value)
	}

	return nil
}

func awsAwsjson11_serializeDocumentMetricDimensions(v []types.MetricDimension, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsAwsjson11_serializeDocumentMetricDimension(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsAwsjson11_serializeDocumentPredefinedLoadMetricSpecification(v *types.PredefinedLoadMetricSpecification, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if len(v.PredefinedLoadMetricType) > 0 {
		ok := object.Key("PredefinedLoadMetricType")
		ok.String(string(v.PredefinedLoadMetricType))
	}

	if v.ResourceLabel != nil {
		ok := object.Key("ResourceLabel")
		ok.String(*v.ResourceLabel)
	}

	return nil
}

func awsAwsjson11_serializeDocumentPredefinedScalingMetricSpecification(v *types.PredefinedScalingMetricSpecification, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if len(v.PredefinedScalingMetricType) > 0 {
		ok := object.Key("PredefinedScalingMetricType")
		ok.String(string(v.PredefinedScalingMetricType))
	}

	if v.ResourceLabel != nil {
		ok := object.Key("ResourceLabel")
		ok.String(*v.ResourceLabel)
	}

	return nil
}

func awsAwsjson11_serializeDocumentScalingInstruction(v *types.ScalingInstruction, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.CustomizedLoadMetricSpecification != nil {
		ok := object.Key("CustomizedLoadMetricSpecification")
		if err := awsAwsjson11_serializeDocumentCustomizedLoadMetricSpecification(v.CustomizedLoadMetricSpecification, ok); err != nil {
			return err
		}
	}

	if v.DisableDynamicScaling != nil {
		ok := object.Key("DisableDynamicScaling")
		ok.Boolean(*v.DisableDynamicScaling)
	}

	if v.MaxCapacity != nil {
		ok := object.Key("MaxCapacity")
		ok.Integer(*v.MaxCapacity)
	}

	if v.MinCapacity != nil {
		ok := object.Key("MinCapacity")
		ok.Integer(*v.MinCapacity)
	}

	if v.PredefinedLoadMetricSpecification != nil {
		ok := object.Key("PredefinedLoadMetricSpecification")
		if err := awsAwsjson11_serializeDocumentPredefinedLoadMetricSpecification(v.PredefinedLoadMetricSpecification, ok); err != nil {
			return err
		}
	}

	if len(v.PredictiveScalingMaxCapacityBehavior) > 0 {
		ok := object.Key("PredictiveScalingMaxCapacityBehavior")
		ok.String(string(v.PredictiveScalingMaxCapacityBehavior))
	}

	if v.PredictiveScalingMaxCapacityBuffer != nil {
		ok := object.Key("PredictiveScalingMaxCapacityBuffer")
		ok.Integer(*v.PredictiveScalingMaxCapacityBuffer)
	}

	if len(v.PredictiveScalingMode) > 0 {
		ok := object.Key("PredictiveScalingMode")
		ok.String(string(v.PredictiveScalingMode))
	}

	if v.ResourceId != nil {
		ok := object.Key("ResourceId")
		ok.String(*v.ResourceId)
	}

	if len(v.ScalableDimension) > 0 {
		ok := object.Key("ScalableDimension")
		ok.String(string(v.ScalableDimension))
	}

	if len(v.ScalingPolicyUpdateBehavior) > 0 {
		ok := object.Key("ScalingPolicyUpdateBehavior")
		ok.String(string(v.ScalingPolicyUpdateBehavior))
	}

	if v.ScheduledActionBufferTime != nil {
		ok := object.Key("ScheduledActionBufferTime")
		ok.Integer(*v.ScheduledActionBufferTime)
	}

	if len(v.ServiceNamespace) > 0 {
		ok := object.Key("ServiceNamespace")
		ok.String(string(v.ServiceNamespace))
	}

	if v.TargetTrackingConfigurations != nil {
		ok := object.Key("TargetTrackingConfigurations")
		if err := awsAwsjson11_serializeDocumentTargetTrackingConfigurations(v.TargetTrackingConfigurations, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsAwsjson11_serializeDocumentScalingInstructions(v []types.ScalingInstruction, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsAwsjson11_serializeDocumentScalingInstruction(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsAwsjson11_serializeDocumentScalingPlanNames(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsAwsjson11_serializeDocumentTagFilter(v *types.TagFilter, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Key != nil {
		ok := object.Key("Key")
		ok.String(*v.Key)
	}

	if v.Values != nil {
		ok := object.Key("Values")
		if err := awsAwsjson11_serializeDocumentTagValues(v.Values, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsAwsjson11_serializeDocumentTagFilters(v []types.TagFilter, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsAwsjson11_serializeDocumentTagFilter(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsAwsjson11_serializeDocumentTagValues(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsAwsjson11_serializeDocumentTargetTrackingConfiguration(v *types.TargetTrackingConfiguration, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.CustomizedScalingMetricSpecification != nil {
		ok := object.Key("CustomizedScalingMetricSpecification")
		if err := awsAwsjson11_serializeDocumentCustomizedScalingMetricSpecification(v.CustomizedScalingMetricSpecification, ok); err != nil {
			return err
		}
	}

	if v.DisableScaleIn != nil {
		ok := object.Key("DisableScaleIn")
		ok.Boolean(*v.DisableScaleIn)
	}

	if v.EstimatedInstanceWarmup != nil {
		ok := object.Key("EstimatedInstanceWarmup")
		ok.Integer(*v.EstimatedInstanceWarmup)
	}

	if v.PredefinedScalingMetricSpecification != nil {
		ok := object.Key("PredefinedScalingMetricSpecification")
		if err := awsAwsjson11_serializeDocumentPredefinedScalingMetricSpecification(v.PredefinedScalingMetricSpecification, ok); err != nil {
			return err
		}
	}

	if v.ScaleInCooldown != nil {
		ok := object.Key("ScaleInCooldown")
		ok.Integer(*v.ScaleInCooldown)
	}

	if v.ScaleOutCooldown != nil {
		ok := object.Key("ScaleOutCooldown")
		ok.Integer(*v.ScaleOutCooldown)
	}

	if v.TargetValue != nil {
		ok := object.Key("TargetValue")
		switch {
		case math.IsNaN(*v.TargetValue):
			ok.String("NaN")

		case math.IsInf(*v.TargetValue, 1):
			ok.String("Infinity")

		case math.IsInf(*v.TargetValue, -1):
			ok.String("-Infinity")

		default:
			ok.Double(*v.TargetValue)

		}
	}

	return nil
}

func awsAwsjson11_serializeDocumentTargetTrackingConfigurations(v []types.TargetTrackingConfiguration, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsAwsjson11_serializeDocumentTargetTrackingConfiguration(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsAwsjson11_serializeOpDocumentCreateScalingPlanInput(v *CreateScalingPlanInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ApplicationSource != nil {
		ok := object.Key("ApplicationSource")
		if err := awsAwsjson11_serializeDocumentApplicationSource(v.ApplicationSource, ok); err != nil {
			return err
		}
	}

	if v.ScalingInstructions != nil {
		ok := object.Key("ScalingInstructions")
		if err := awsAwsjson11_serializeDocumentScalingInstructions(v.ScalingInstructions, ok); err != nil {
			return err
		}
	}

	if v.ScalingPlanName != nil {
		ok := object.Key("ScalingPlanName")
		ok.String(*v.ScalingPlanName)
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentDeleteScalingPlanInput(v *DeleteScalingPlanInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ScalingPlanName != nil {
		ok := object.Key("ScalingPlanName")
		ok.String(*v.ScalingPlanName)
	}

	if v.ScalingPlanVersion != nil {
		ok := object.Key("ScalingPlanVersion")
		ok.Long(*v.ScalingPlanVersion)
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentDescribeScalingPlanResourcesInput(v *DescribeScalingPlanResourcesInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.MaxResults != nil {
		ok := object.Key("MaxResults")
		ok.Integer(*v.MaxResults)
	}

	if v.NextToken != nil {
		ok := object.Key("NextToken")
		ok.String(*v.NextToken)
	}

	if v.ScalingPlanName != nil {
		ok := object.Key("ScalingPlanName")
		ok.String(*v.ScalingPlanName)
	}

	if v.ScalingPlanVersion != nil {
		ok := object.Key("ScalingPlanVersion")
		ok.Long(*v.ScalingPlanVersion)
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentDescribeScalingPlansInput(v *DescribeScalingPlansInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ApplicationSources != nil {
		ok := object.Key("ApplicationSources")
		if err := awsAwsjson11_serializeDocumentApplicationSources(v.ApplicationSources, ok); err != nil {
			return err
		}
	}

	if v.MaxResults != nil {
		ok := object.Key("MaxResults")
		ok.Integer(*v.MaxResults)
	}

	if v.NextToken != nil {
		ok := object.Key("NextToken")
		ok.String(*v.NextToken)
	}

	if v.ScalingPlanNames != nil {
		ok := object.Key("ScalingPlanNames")
		if err := awsAwsjson11_serializeDocumentScalingPlanNames(v.ScalingPlanNames, ok); err != nil {
			return err
		}
	}

	if v.ScalingPlanVersion != nil {
		ok := object.Key("ScalingPlanVersion")
		ok.Long(*v.ScalingPlanVersion)
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentGetScalingPlanResourceForecastDataInput(v *GetScalingPlanResourceForecastDataInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.EndTime != nil {
		ok := object.Key("EndTime")
		ok.Double(smithytime.FormatEpochSeconds(*v.EndTime))
	}

	if len(v.ForecastDataType) > 0 {
		ok := object.Key("ForecastDataType")
		ok.String(string(v.ForecastDataType))
	}

	if v.ResourceId != nil {
		ok := object.Key("ResourceId")
		ok.String(*v.ResourceId)
	}

	if len(v.ScalableDimension) > 0 {
		ok := object.Key("ScalableDimension")
		ok.String(string(v.ScalableDimension))
	}

	if v.ScalingPlanName != nil {
		ok := object.Key("ScalingPlanName")
		ok.String(*v.ScalingPlanName)
	}

	if v.ScalingPlanVersion != nil {
		ok := object.Key("ScalingPlanVersion")
		ok.Long(*v.ScalingPlanVersion)
	}

	if len(v.ServiceNamespace) > 0 {
		ok := object.Key("ServiceNamespace")
		ok.String(string(v.ServiceNamespace))
	}

	if v.StartTime != nil {
		ok := object.Key("StartTime")
		ok.Double(smithytime.FormatEpochSeconds(*v.StartTime))
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentUpdateScalingPlanInput(v *UpdateScalingPlanInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ApplicationSource != nil {
		ok := object.Key("ApplicationSource")
		if err := awsAwsjson11_serializeDocumentApplicationSource(v.ApplicationSource, ok); err != nil {
			return err
		}
	}

	if v.ScalingInstructions != nil {
		ok := object.Key("ScalingInstructions")
		if err := awsAwsjson11_serializeDocumentScalingInstructions(v.ScalingInstructions, ok); err != nil {
			return err
		}
	}

	if v.ScalingPlanName != nil {
		ok := object.Key("ScalingPlanName")
		ok.String(*v.ScalingPlanName)
	}

	if v.ScalingPlanVersion != nil {
		ok := object.Key("ScalingPlanVersion")
		ok.Long(*v.ScalingPlanVersion)
	}

	return nil
}
