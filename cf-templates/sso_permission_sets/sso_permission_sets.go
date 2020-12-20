package main

import (
	cf "github.com/awslabs/goformation/v4/cloudformation"
	"github.com/awslabs/goformation/v4/cloudformation/sso"
	"github.com/awslabs/goformation/v4/cloudformation/tags"
)

// Enable as plugin
type template bool
var Template template

// Globals
var templateName = "sso_permission_sets"

// Generate the template
func (t template) Render() cf.Template {
	// Create Template
	template := cf.NewTemplate()

	// Parameters
	template.Parameters["SsoInstanceArn"] = cf.Parameter{
		Type:        "String",
		Description: "The SSO Instance ARN to attach Permission Sets to",
	}

	template.Parameters["AdministratorSsoGroupId"] = cf.Parameter{
		Type:        "String",
		Description: "The ID of the Administrator SSO group",
	}

	template.Parameters["DeveloperSsoGroupId"] = cf.Parameter{
		Type:        "String",
		Description: "The ID of the Developer SSO group",
	}

	// Create PermissionSets
	template.Resources["SsoPermissionSetAdministrator"] = &sso.PermissionSet{
		Description: "A PermissionSet used for Administrator users",
		InstanceArn: cf.Ref("SsoInstanceArn"),
		ManagedPolicies: []string{
			"arn:aws:iam::aws:policy/AdministratorAccess",
		},
		Name: "SSO-Administrator",
		Tags: []tags.Tag{
			{
				Key:   "TemplateName",
				Value: "sso_permission_sets",
			},
		},
	}

	template.Resources["SsoPermissionSetDeveloper"] = &sso.PermissionSet{
		Description: "A PermissionSet used for Developer users",
		InstanceArn: cf.Ref("SsoInstanceArn"),
		ManagedPolicies: []string{
			"arn:aws:iam::aws:policy/PowerUserAccess",
		},
		Name: "SSO-Developer",
		Tags: []tags.Tag{
			{
				Key:   "TemplateName",
				Value: templateName,
			},
		},
	}

	// Create Assignments
	template.Resources["SsoPermissionSetAdministratorAssignment"] = &sso.Assignment{
		InstanceArn:      cf.Ref("SsoInstanceArn"),
		PermissionSetArn: cf.GetAtt("SsoPermissionSetAdministrator", "PermissionSetArn"),
		PrincipalId:      cf.Ref("AdministratorSsoGroupId"),
		PrincipalType:    "GROUP",
		TargetId:         cf.Ref("AWS::AccountId"),
		TargetType:       "AWS_ACCOUNT",
	}

	template.Resources["SsoPermissionSetDeveloperAssignment"] = &sso.Assignment{
		InstanceArn:      cf.Ref("SsoInstanceArn"),
		PermissionSetArn: cf.GetAtt("SsoPermissionSetDeveloper", "PermissionSetArn"),
		PrincipalId:      cf.Ref("DeveloperSsoGroupId"),
		PrincipalType:    "GROUP",
		TargetId:         cf.Ref("AWS::AccountId"),
		TargetType:       "AWS_ACCOUNT",
	}

	return *template
}
