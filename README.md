# ecs-fargate-app

An AWS Elastic Container Service Fargate example application with Go and Cloudformation using Redis ElastiCache for storage.

## Usage

You will need a docker image stored in a repository in order to get started. I typically use AWS ECS but any other repository will work so long as you have
sufficient permissions to download the image to the Fargate cluster.

To build and deploy the demo web application included in the webapp/ directory, follow the instructions in the Demo webapp segment.

Once you have the sample webapp (or your actual docker container you want to use) available via a repository, you should record the location for use in the CloudFormation template.

## Demo webapp

There is a web application included which provides a demo site for you to test your deployment. To compile and build this docker image:

    #> git clone https://github.com/routebyintuition/ecs-fargate-app.git
    #> ecs-fargate-app/webapp
    #> go build
    #> docker build -t webapp .
    ># docker push <DESTINATION>

You will now want to follow the instructions for your docker repository of choice. You will need the image location during cloudformation initiation.

## AWS CloudFormation

You can use the provided cloudformation template to deploy the needed infrastructure for AWS RDS, ElastiCache, ELB/ALB, and ECS.

There are four required inputs for the AWS CloudFormation deployment.

* DockerImage: this is the repository link to the image being deployed to AWS Fargate
* VpcCIDR: CIDR block of overall VPC
* PublicSubnet1CIDR: CIDR block within VPC for the first subnet
* PublicSubnet1CIDR: CIDR block within VPC for the second subnet

Note: Only the DockerImage entry needs to be changed. The remainding CIDR block entries can be used with their defaults.

### AWS ECS

The provided cloudformation template includes an AWS ECS cluster to provision the ECS service and fargate tasks. The security group associated with this service identifies allowed traffic to the containers.

