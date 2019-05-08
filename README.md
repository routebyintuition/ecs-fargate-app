# ecs-fargate-app

An AWS Elastic Container Service Fargate example application with Go and Cloudformation using Redis ElastiCache for storage.

## Usage

You will need a docker image stored in a repository in order to get started. I typically use AWS ECS but any other repository will work so long as you have
sufficient permissions to download the image to the Fargate cluster.

To build and deploy the demo web application included in the webapp/ directory, follow the instructions in the Demo webapp segment.



## Demo webapp

There is a web application included which provides a demo site for you to test your deployment.

## AWS CloudFormation

You can use the provided cloudformation template to deploy the needed infrastructure for AWS RDS, ElastiCache, ELB/ALB, and ECS.

