// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc

import (
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
)

func ResourceIBMSccRuleAttachmentValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 1)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "scope_type",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              "enterprise, enterprise.account_group, enterprise.account, account, account.resource_group",
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_scc_rule_attachment", Schema: validateSchema}
	return &resourceValidator
}
