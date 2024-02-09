// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3cust "github.com/aws/aws-sdk-go-v2/service/s3/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This action creates an Amazon S3 bucket. To create an Amazon S3 on Outposts
// bucket, see CreateBucket (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateBucket.html)
// . Creates a new S3 bucket. To create a bucket, you must set up Amazon S3 and
// have a valid Amazon Web Services Access Key ID to authenticate requests.
// Anonymous requests are never allowed to create buckets. By creating the bucket,
// you become the bucket owner. There are two types of buckets: general purpose
// buckets and directory buckets. For more information about these bucket types,
// see Creating, configuring, and working with Amazon S3 buckets (https://docs.aws.amazon.com/AmazonS3/latest/userguide/creating-buckets-s3.html)
// in the Amazon S3 User Guide.
//   - General purpose buckets - If you send your CreateBucket request to the
//     s3.amazonaws.com global endpoint, the request goes to the us-east-1 Region. So
//     the signature calculations in Signature Version 4 must use us-east-1 as the
//     Region, even if the location constraint in the request specifies another Region
//     where the bucket is to be created. If you create a bucket in a Region other than
//     US East (N. Virginia), your application must be able to handle 307 redirect. For
//     more information, see Virtual hosting of buckets (https://docs.aws.amazon.com/AmazonS3/latest/dev/VirtualHosting.html)
//     in the Amazon S3 User Guide.
//   - Directory buckets - For directory buckets, you must make requests for this
//     API operation to the Regional endpoint. These endpoints support path-style
//     requests in the format
//     https://s3express-control.region_code.amazonaws.com/bucket-name .
//     Virtual-hosted-style requests aren't supported. For more information, see
//     Regional and Zonal endpoints (https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-Regions-and-Zones.html)
//     in the Amazon S3 User Guide.
//
// Permissions
//   - General purpose bucket permissions - In addition to the s3:CreateBucket
//     permission, the following permissions are required in a policy when your
//     CreateBucket request includes specific headers:
//   - Access control lists (ACLs) - In your CreateBucket request, if you specify
//     an access control list (ACL) and set it to public-read , public-read-write ,
//     authenticated-read , or if you explicitly specify any other custom ACLs, both
//     s3:CreateBucket and s3:PutBucketAcl permissions are required. In your
//     CreateBucket request, if you set the ACL to private , or if you don't specify
//     any ACLs, only the s3:CreateBucket permission is required.
//   - Object Lock - In your CreateBucket request, if you set
//     x-amz-bucket-object-lock-enabled to true, the
//     s3:PutBucketObjectLockConfiguration and s3:PutBucketVersioning permissions are
//     required.
//   - S3 Object Ownership - If your CreateBucket request includes the
//     x-amz-object-ownership header, then the s3:PutBucketOwnershipControls
//     permission is required. If your CreateBucket request sets BucketOwnerEnforced
//     for Amazon S3 Object Ownership and specifies a bucket ACL that provides access
//     to an external Amazon Web Services account, your request fails with a 400
//     error and returns the InvalidBucketAcLWithObjectOwnership error code. For more
//     information, see Setting Object Ownership on an existing bucket  (https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-ownership-existing-bucket.html)
//     in the Amazon S3 User Guide.
//   - S3 Block Public Access - If your specific use case requires granting public
//     access to your S3 resources, you can disable Block Public Access. Specifically,
//     you can create a new bucket with Block Public Access enabled, then separately
//     call the DeletePublicAccessBlock (https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeletePublicAccessBlock.html)
//     API. To use this operation, you must have the s3:PutBucketPublicAccessBlock
//     permission. For more information about S3 Block Public Access, see Blocking
//     public access to your Amazon S3 storage  (https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-control-block-public-access.html)
//     in the Amazon S3 User Guide.
//   - Directory bucket permissions - You must have the s3express:CreateBucket
//     permission in an IAM identity-based policy instead of a bucket policy.
//     Cross-account access to this API operation isn't supported. This operation can
//     only be performed by the Amazon Web Services account that owns the resource. For
//     more information about directory bucket policies and permissions, see Amazon
//     Web Services Identity and Access Management (IAM) for S3 Express One Zone (https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-security-iam.html)
//     in the Amazon S3 User Guide. The permissions for ACLs, Object Lock, S3 Object
//     Ownership, and S3 Block Public Access are not supported for directory buckets.
//     For directory buckets, all Block Public Access settings are enabled at the
//     bucket level and S3 Object Ownership is set to Bucket owner enforced (ACLs
//     disabled). These settings can't be modified. For more information about
//     permissions for creating and working with directory buckets, see Directory
//     buckets (https://docs.aws.amazon.com/AmazonS3/latest/userguide/directory-buckets-overview.html)
//     in the Amazon S3 User Guide. For more information about supported S3 features
//     for directory buckets, see Features of S3 Express One Zone (https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-one-zone.html#s3-express-features)
//     in the Amazon S3 User Guide.
//
// HTTP Host header syntax Directory buckets - The HTTP Host header syntax is
// s3express-control.region.amazonaws.com . The following operations are related to
// CreateBucket :
//   - PutObject (https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObject.html)
//   - DeleteBucket (https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucket.html)
func (c *Client) CreateBucket(ctx context.Context, params *CreateBucketInput, optFns ...func(*Options)) (*CreateBucketOutput, error) {
	if params == nil {
		params = &CreateBucketInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateBucket", params, optFns, c.addOperationCreateBucketMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateBucketOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateBucketInput struct {

	// The name of the bucket to create. General purpose buckets - For information
	// about bucket naming restrictions, see Bucket naming rules (https://docs.aws.amazon.com/AmazonS3/latest/userguide/bucketnamingrules.html)
	// in the Amazon S3 User Guide. Directory buckets - When you use this operation
	// with a directory bucket, you must use path-style requests in the format
	// https://s3express-control.region_code.amazonaws.com/bucket-name .
	// Virtual-hosted-style requests aren't supported. Directory bucket names must be
	// unique in the chosen Availability Zone. Bucket names must also follow the format
	// bucket_base_name--az_id--x-s3 (for example,  DOC-EXAMPLE-BUCKET--usw2-az2--x-s3
	// ). For information about bucket naming restrictions, see Directory bucket
	// naming rules (https://docs.aws.amazon.com/AmazonS3/latest/userguide/directory-bucket-naming-rules.html)
	// in the Amazon S3 User Guide
	//
	// This member is required.
	Bucket *string

	// The canned ACL to apply to the bucket. This functionality is not supported for
	// directory buckets.
	ACL types.BucketCannedACL

	// The configuration information for the bucket.
	CreateBucketConfiguration *types.CreateBucketConfiguration

	// Allows grantee the read, write, read ACP, and write ACP permissions on the
	// bucket. This functionality is not supported for directory buckets.
	GrantFullControl *string

	// Allows grantee to list the objects in the bucket. This functionality is not
	// supported for directory buckets.
	GrantRead *string

	// Allows grantee to read the bucket ACL. This functionality is not supported for
	// directory buckets.
	GrantReadACP *string

	// Allows grantee to create new objects in the bucket. For the bucket and object
	// owners of existing objects, also allows deletions and overwrites of those
	// objects. This functionality is not supported for directory buckets.
	GrantWrite *string

	// Allows grantee to write the ACL for the applicable bucket. This functionality
	// is not supported for directory buckets.
	GrantWriteACP *string

	// Specifies whether you want S3 Object Lock to be enabled for the new bucket.
	// This functionality is not supported for directory buckets.
	ObjectLockEnabledForBucket *bool

	// The container element for object ownership for a bucket's ownership controls.
	// BucketOwnerPreferred - Objects uploaded to the bucket change ownership to the
	// bucket owner if the objects are uploaded with the bucket-owner-full-control
	// canned ACL. ObjectWriter - The uploading account will own the object if the
	// object is uploaded with the bucket-owner-full-control canned ACL.
	// BucketOwnerEnforced - Access control lists (ACLs) are disabled and no longer
	// affect permissions. The bucket owner automatically owns and has full control
	// over every object in the bucket. The bucket only accepts PUT requests that don't
	// specify an ACL or specify bucket owner full control ACLs (such as the predefined
	// bucket-owner-full-control canned ACL or a custom ACL in XML format that grants
	// the same permissions). By default, ObjectOwnership is set to BucketOwnerEnforced
	// and ACLs are disabled. We recommend keeping ACLs disabled, except in uncommon
	// use cases where you must control access for each object individually. For more
	// information about S3 Object Ownership, see Controlling ownership of objects and
	// disabling ACLs for your bucket (https://docs.aws.amazon.com/AmazonS3/latest/userguide/about-object-ownership.html)
	// in the Amazon S3 User Guide. This functionality is not supported for directory
	// buckets. Directory buckets use the bucket owner enforced setting for S3 Object
	// Ownership.
	ObjectOwnership types.ObjectOwnership

	noSmithyDocumentSerde
}

func (in *CreateBucketInput) bindEndpointParams(p *EndpointParameters) {
	p.Bucket = in.Bucket
	p.UseS3ExpressControlEndpoint = ptr.Bool(true)
	p.DisableAccessPoints = ptr.Bool(true)
}

type CreateBucketOutput struct {

	// A forward slash followed by the name of the bucket.
	Location *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateBucketMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpCreateBucket{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpCreateBucket{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateBucket"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addPutBucketContextMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateBucketValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateBucket(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addCreateBucketUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = disableAcceptEncodingGzip(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSerializeImmutableHostnameBucketMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func (v *CreateBucketInput) bucket() (string, bool) {
	if v.Bucket == nil {
		return "", false
	}
	return *v.Bucket, true
}

func newServiceMetadataMiddleware_opCreateBucket(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateBucket",
	}
}

// getCreateBucketBucketMember returns a pointer to string denoting a provided
// bucket member valueand a boolean indicating if the input has a modeled bucket
// name,
func getCreateBucketBucketMember(input interface{}) (*string, bool) {
	in := input.(*CreateBucketInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addCreateBucketUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: getCreateBucketBucketMember,
		},
		UsePathStyle:                   options.UsePathStyle,
		UseAccelerate:                  options.UseAccelerate,
		SupportsAccelerate:             false,
		TargetS3ObjectLambda:           false,
		EndpointResolver:               options.EndpointResolver,
		EndpointResolverOptions:        options.EndpointOptions,
		UseARNRegion:                   options.UseARNRegion,
		DisableMultiRegionAccessPoints: options.DisableMultiRegionAccessPoints,
	})
}
