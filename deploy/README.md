# deploy

## terraform

### On MacOS

Install terraform:

    brew tap hashicorp/tap
    brew install hashicorp/tap/terraform

Verify installation:

    terraform -v

Add AWS aws_secret_key and aws_access_key to a file `.tfvars`

Create the plan:

    terraform plan -var-file=.tfvars -out plan.tfplan
