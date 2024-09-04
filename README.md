# GitHub AWS OIDC Integration

This repository contains a Pulumi program that sets up OpenID Connect (OIDC) integration between AWS and GitHub Actions. This integration allows GitHub Actions workflows to securely authenticate with AWS without the need for long-lived AWS credentials.

## Contents

- `main.go`: The main Pulumi program written in Go.
- `Pulumi.yaml`: The Pulumi project configuration file.
- `Pulumi.main.yaml`: The Pulumi stack configuration file.

## What it does

This Pulumi program creates the following AWS resources:

1. An IAM OpenID Connect provider for GitHub Actions.
2. An IAM role that can be assumed by GitHub Actions workflows.
3. A policy attachment giving the IAM role administrator access (for Pulumi deployments).

## Configuration

The program is configured to trust repositories matching the pattern defined by `githubRepoRegex`. When forking, make sure to update this variable before deploying the program.

## Usage

To use this Pulumi program:

1. Ensure you have Pulumi and Go installed and configured on your local machine
2. Clone this repository.
3. Update the `githubRepoRegex` variable in `main.go` to match your GitHub repository.
4. Run `pulumi up` to create the resources in your AWS account.

After deployment, you can use the created IAM role in your GitHub Actions workflows to authenticate with AWS.

## Security Note

The IAM role created by this program has administrator access. While this is suitable for Pulumi deployments that need broad permissions, it's recommended to review and potentially restrict these permissions based on your specific needs and security requirements.
