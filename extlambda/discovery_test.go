/*
 * Copyright 2023 steadybit GmbH. All rights reserved.
 */

package extlambda

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/steadybit/discovery-kit/go/discovery_kit_api"
	"github.com/steadybit/extension-kit/extutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type lambdaClientMock struct {
	mock.Mock
}

func (m lambdaClientMock) ListFunctions(ctx context.Context, params *lambda.ListFunctionsInput, optFns ...func(*lambda.Options)) (*lambda.ListFunctionsOutput, error) {
	args := m.Called(ctx, params)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*lambda.ListFunctionsOutput), args.Error(1)
}

func Test_getAllAwsLambdaFunctions(t *testing.T) {
	api := new(lambdaClientMock)
	listedFunction := lambda.ListFunctionsOutput{
		Functions: []types.FunctionConfiguration{
			{
				Architectures: []types.Architecture{"x86_64"},
				CodeSize:      1024,
				Description:   extutil.Ptr("description"),
				Environment: extutil.Ptr(types.EnvironmentResponse{
					Variables: map[string]string{
						"FAILURE_INJECTION_PARAM": "env-fip",
					},
				}),
				FunctionArn:  extutil.Ptr("arn"),
				FunctionName: extutil.Ptr("name"),
				LastModified: extutil.Ptr("last-modified"),
				MasterArn:    extutil.Ptr("master-arn"),
				MemorySize:   extutil.Ptr(int32(1024)),
				PackageType:  "package-type",
				RevisionId:   extutil.Ptr("revision-id"),
				Role:         extutil.Ptr("role"),
				Runtime:      "runtime",
				Timeout:      extutil.Ptr(int32(10)),
				Version:      extutil.Ptr("version"),
			},
		},
	}
	api.On("ListFunctions", mock.Anything, mock.Anything).Return(&listedFunction, nil)

	// When
	targets, err := getAllAwsLambdaFunctions(context.Background(), api, "42", "us-east-1")

	// Then
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, len(*targets))

	target := (*targets)[0]
	assert.Equal(t, lambdaTargetID, target.TargetType)
	assert.Equal(t, "name", target.Label)
	assert.Equal(t, "arn", target.Id)
	assert.Equal(t, 18, len(target.Attributes))
	assert.Equal(t, []string{"42"}, target.Attributes["aws.account"])
	assert.Equal(t, []string{"us-east-1"}, target.Attributes["aws.region"])
	assert.Equal(t, []string{"name"}, target.Attributes["aws.lambda.function-name"])
	assert.Equal(t, []string{"env-fip"}, target.Attributes["aws.lambda.failure-injection-param"])
}

func Test_getAllAwsLambdaFunctions_withPagination(t *testing.T) {
	// Given
	mockedApi := new(lambdaClientMock)

	withMarker := mock.MatchedBy(func(arg *lambda.ListFunctionsInput) bool {
		return arg.Marker != nil
	})
	withoutMarker := mock.MatchedBy(func(arg *lambda.ListFunctionsInput) bool {
		return arg.Marker == nil
	})
	mockedApi.On("ListFunctions", mock.Anything, withoutMarker).Return(&lambda.ListFunctionsOutput{
		NextMarker: discovery_kit_api.Ptr("marker"),
		Functions: []types.FunctionConfiguration{
			{
				FunctionArn: extutil.Ptr("arn1"),
			},
		},
	}, nil)
	mockedApi.On("ListFunctions", mock.Anything, withMarker).Return(&lambda.ListFunctionsOutput{
		Functions: []types.FunctionConfiguration{
			{
				FunctionArn: extutil.Ptr("arn2"),
			},
		},
	}, nil)

	// When
	targets, err := getAllAwsLambdaFunctions(context.Background(), mockedApi, "42", "us-east-1")

	// Then
	assert.Equal(t, nil, err)
	assert.Equal(t, 2, len(*targets))
	assert.Equal(t, "arn1", (*targets)[0].Id)
	assert.Equal(t, "arn2", (*targets)[1].Id)
}

func Test_getAllAwsLambdaFunctions_withError(t *testing.T) {
	// Given
	api := new(lambdaClientMock)
	api.On("ListFunctions", mock.Anything, mock.Anything).Return(nil, errors.New("error"))

	// When
	_, err := getAllAwsLambdaFunctions(context.Background(), api, "42", "us-east-1")
	assert.Equal(t, "error", err.Error())
}
