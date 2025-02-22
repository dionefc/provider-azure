/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta11 "github.com/upbound/provider-azure/apis/azure/v1beta1"
	v1beta1 "github.com/upbound/provider-azure/apis/network/v1beta1"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this NamespaceAuthorizationRule.
func (mg *NamespaceAuthorizationRule) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NamespaceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.NamespaceIDRef,
		Selector:     mg.Spec.ForProvider.NamespaceIDSelector,
		To: reference.To{
			List:    &ServiceBusNamespaceList{},
			Managed: &ServiceBusNamespace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NamespaceID")
	}
	mg.Spec.ForProvider.NamespaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NamespaceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this NamespaceDisasterRecoveryConfig.
func (mg *NamespaceDisasterRecoveryConfig) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AliasAuthorizationRuleID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.AliasAuthorizationRuleIDRef,
		Selector:     mg.Spec.ForProvider.AliasAuthorizationRuleIDSelector,
		To: reference.To{
			List:    &NamespaceAuthorizationRuleList{},
			Managed: &NamespaceAuthorizationRule{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AliasAuthorizationRuleID")
	}
	mg.Spec.ForProvider.AliasAuthorizationRuleID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AliasAuthorizationRuleIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PartnerNamespaceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.PartnerNamespaceIDRef,
		Selector:     mg.Spec.ForProvider.PartnerNamespaceIDSelector,
		To: reference.To{
			List:    &ServiceBusNamespaceList{},
			Managed: &ServiceBusNamespace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PartnerNamespaceID")
	}
	mg.Spec.ForProvider.PartnerNamespaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PartnerNamespaceIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PrimaryNamespaceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.PrimaryNamespaceIDRef,
		Selector:     mg.Spec.ForProvider.PrimaryNamespaceIDSelector,
		To: reference.To{
			List:    &ServiceBusNamespaceList{},
			Managed: &ServiceBusNamespace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PrimaryNamespaceID")
	}
	mg.Spec.ForProvider.PrimaryNamespaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PrimaryNamespaceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this NamespaceNetworkRuleSet.
func (mg *NamespaceNetworkRuleSet) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NamespaceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.NamespaceIDRef,
		Selector:     mg.Spec.ForProvider.NamespaceIDSelector,
		To: reference.To{
			List:    &ServiceBusNamespaceList{},
			Managed: &ServiceBusNamespace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NamespaceID")
	}
	mg.Spec.ForProvider.NamespaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NamespaceIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.NetworkRules); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkRules[i3].SubnetID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.NetworkRules[i3].SubnetIDRef,
			Selector:     mg.Spec.ForProvider.NetworkRules[i3].SubnetIDSelector,
			To: reference.To{
				List:    &v1beta1.SubnetList{},
				Managed: &v1beta1.Subnet{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.NetworkRules[i3].SubnetID")
		}
		mg.Spec.ForProvider.NetworkRules[i3].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.NetworkRules[i3].SubnetIDRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this Queue.
func (mg *Queue) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NamespaceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.NamespaceIDRef,
		Selector:     mg.Spec.ForProvider.NamespaceIDSelector,
		To: reference.To{
			List:    &ServiceBusNamespaceList{},
			Managed: &ServiceBusNamespace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NamespaceID")
	}
	mg.Spec.ForProvider.NamespaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NamespaceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this QueueAuthorizationRule.
func (mg *QueueAuthorizationRule) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.QueueID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.QueueIDRef,
		Selector:     mg.Spec.ForProvider.QueueIDSelector,
		To: reference.To{
			List:    &QueueList{},
			Managed: &Queue{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.QueueID")
	}
	mg.Spec.ForProvider.QueueID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.QueueIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ServiceBusNamespace.
func (mg *ServiceBusNamespace) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta11.ResourceGroupList{},
			Managed: &v1beta11.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Subscription.
func (mg *Subscription) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TopicID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.TopicIDRef,
		Selector:     mg.Spec.ForProvider.TopicIDSelector,
		To: reference.To{
			List:    &TopicList{},
			Managed: &Topic{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TopicID")
	}
	mg.Spec.ForProvider.TopicID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TopicIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SubscriptionRule.
func (mg *SubscriptionRule) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubscriptionID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.SubscriptionIDRef,
		Selector:     mg.Spec.ForProvider.SubscriptionIDSelector,
		To: reference.To{
			List:    &SubscriptionList{},
			Managed: &Subscription{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubscriptionID")
	}
	mg.Spec.ForProvider.SubscriptionID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubscriptionIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Topic.
func (mg *Topic) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NamespaceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.NamespaceIDRef,
		Selector:     mg.Spec.ForProvider.NamespaceIDSelector,
		To: reference.To{
			List:    &ServiceBusNamespaceList{},
			Managed: &ServiceBusNamespace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NamespaceID")
	}
	mg.Spec.ForProvider.NamespaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NamespaceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TopicAuthorizationRule.
func (mg *TopicAuthorizationRule) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TopicID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.TopicIDRef,
		Selector:     mg.Spec.ForProvider.TopicIDSelector,
		To: reference.To{
			List:    &TopicList{},
			Managed: &Topic{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TopicID")
	}
	mg.Spec.ForProvider.TopicID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TopicIDRef = rsp.ResolvedReference

	return nil
}
