package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The repos in github that our particular AWS account should trust
const (
	githubRepoRegex = "repo:amirhajimirsadeghi/*"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create an IAM OpenID Connect provider for GitHub Actions
		githubProvider, err := iam.NewOpenIdConnectProvider(ctx, "github-provider", &iam.OpenIdConnectProviderArgs{
			Url:             pulumi.String("https://token.actions.githubusercontent.com"),
			ClientIdLists:   pulumi.StringArray{pulumi.String("sts.amazonaws.com")},
			ThumbprintLists: pulumi.StringArray{pulumi.String("6938fd4d98bab03faadb97b34396831e3780aea1")},
		})
		if err != nil {
			return err
		}

		// Create an IAM role for GitHub Actions
		githubActionsRole, err := iam.NewRole(ctx, "github-actions-role", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.Sprintf(`{
				"Version": "2012-10-17",
				"Statement": [
					{
						"Effect": "Allow",
						"Principal": {
							"Federated": "%s"
						},
						"Action": "sts:AssumeRoleWithWebIdentity",
						"Condition": {
							"StringEquals": {
								"token.actions.githubusercontent.com:aud": "sts.amazonaws.com"
							},
							"StringLike": {
								"token.actions.githubusercontent.com:sub": "%s"
							}
						}
					}
				]
			}`, githubProvider.Arn, githubRepoRegex),
		})
		if err != nil {
			return err
		}

		// Giving administator access because it will be used by pulumi to do deployments
		_, err = iam.NewRolePolicyAttachment(ctx, "github-actions-s3-policy", &iam.RolePolicyAttachmentArgs{
			Role:      githubActionsRole.Name,
			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/AdministratorAccess"),
		})
		return err
	})
}
