// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticloadbalancingv2

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// ~> **Note:** `aws_alb_target_group` is known as `aws_lb_target_group`. The functionality is identical.
// 
// Provides information about a Load Balancer Target Group.
// 
// This data source can prove useful when a module accepts an LB Target Group as an
// input variable and needs to know its attributes. It can also be used to get the ARN of
// an LB Target Group for use in other resources, given LB Target Group name.
func LookupTargetGroup(ctx *pulumi.Context, args *GetTargetGroupArgs) (*GetTargetGroupResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["arn"] = args.Arn
		inputs["name"] = args.Name
		inputs["tags"] = args.Tags
	}
	outputs, err := ctx.Invoke("aws:elasticloadbalancingv2/getTargetGroup:getTargetGroup", inputs)
	if err != nil {
		return nil, err
	}
	return &GetTargetGroupResult{
		Arn: outputs["arn"],
		ArnSuffix: outputs["arnSuffix"],
		DeregistrationDelay: outputs["deregistrationDelay"],
		HealthCheck: outputs["healthCheck"],
		Name: outputs["name"],
		Port: outputs["port"],
		Protocol: outputs["protocol"],
		SlowStart: outputs["slowStart"],
		Stickiness: outputs["stickiness"],
		Tags: outputs["tags"],
		VpcId: outputs["vpcId"],
	}, nil
}

// A collection of arguments for invoking getTargetGroup.
type GetTargetGroupArgs struct {
	// The full ARN of the target group.
	Arn interface{}
	// The unique name of the target group.
	Name interface{}
	Tags interface{}
}

// A collection of values returned by getTargetGroup.
type GetTargetGroupResult struct {
	Arn interface{}
	ArnSuffix interface{}
	DeregistrationDelay interface{}
	HealthCheck interface{}
	Name interface{}
	Port interface{}
	Protocol interface{}
	SlowStart interface{}
	Stickiness interface{}
	Tags interface{}
	VpcId interface{}
}