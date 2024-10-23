// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
// Code generated by internal/generate/namevaluesfilters/v2/main.go; DO NOT EDIT.

package v2

import ( // nosemgrep:ci.semgrep.aws.multiple-service-imports
	"github.com/aws/aws-sdk-go-v2/aws"
	imagebuildertypes "github.com/aws/aws-sdk-go-v2/service/imagebuilder/types"
	licensemanagertypes "github.com/aws/aws-sdk-go-v2/service/licensemanager/types"
	rdstypes "github.com/aws/aws-sdk-go-v2/service/rds/types"
	route53resolvertypes "github.com/aws/aws-sdk-go-v2/service/route53resolver/types"
	secretsmanagertypes "github.com/aws/aws-sdk-go-v2/service/secretsmanager/types"
)

// []*SERVICE.Filter handling

// ImageBuilderFilters returns imagebuilder service filters.
func (filters NameValuesFilters) ImageBuilderFilters() []imagebuildertypes.Filter {
	m := filters.Map()

	if len(m) == 0 {
		return nil
	}

	result := make([]imagebuildertypes.Filter, 0, len(m))

	for k, v := range m {
		filter := imagebuildertypes.Filter{
			Name:   aws.String(k),
			Values: v,
		}

		result = append(result, filter)
	}

	return result
}

// LicenseManagerFilters returns licensemanager service filters.
func (filters NameValuesFilters) LicenseManagerFilters() []licensemanagertypes.Filter {
	m := filters.Map()

	if len(m) == 0 {
		return nil
	}

	result := make([]licensemanagertypes.Filter, 0, len(m))

	for k, v := range m {
		filter := licensemanagertypes.Filter{
			Name:   aws.String(k),
			Values: v,
		}

		result = append(result, filter)
	}

	return result
}

// RDSFilters returns rds service filters.
func (filters NameValuesFilters) RDSFilters() []rdstypes.Filter {
	m := filters.Map()

	if len(m) == 0 {
		return nil
	}

	result := make([]rdstypes.Filter, 0, len(m))

	for k, v := range m {
		filter := rdstypes.Filter{
			Name:   aws.String(k),
			Values: v,
		}

		result = append(result, filter)
	}

	return result
}

// Route53ResolverFilters returns route53resolver service filters.
func (filters NameValuesFilters) Route53ResolverFilters() []route53resolvertypes.Filter {
	m := filters.Map()

	if len(m) == 0 {
		return nil
	}

	result := make([]route53resolvertypes.Filter, 0, len(m))

	for k, v := range m {
		filter := route53resolvertypes.Filter{
			Name:   aws.String(k),
			Values: v,
		}

		result = append(result, filter)
	}

	return result
}

// SecretsManagerFilters returns secretsmanager service filters.
func (filters NameValuesFilters) SecretsManagerFilters() []secretsmanagertypes.Filter {
	m := filters.Map()

	if len(m) == 0 {
		return nil
	}

	result := make([]secretsmanagertypes.Filter, 0, len(m))

	for k, v := range m {
		filter := secretsmanagertypes.Filter{
			Key:    secretsmanagertypes.FilterNameStringType(k),
			Values: v,
		}

		result = append(result, filter)
	}

	return result
}
