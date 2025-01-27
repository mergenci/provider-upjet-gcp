// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this Release.
	apisresolver "github.com/upbound/provider-gcp/internal/apis"
)

func (mg *Release) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("firebaserules.gcp.upbound.io", "v1beta1", "Ruleset", "RulesetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RulesetName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.RulesetNameRef,
			Selector:     mg.Spec.ForProvider.RulesetNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RulesetName")
	}
	mg.Spec.ForProvider.RulesetName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RulesetNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("firebaserules.gcp.upbound.io", "v1beta1", "Ruleset", "RulesetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RulesetName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.RulesetNameRef,
			Selector:     mg.Spec.InitProvider.RulesetNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RulesetName")
	}
	mg.Spec.InitProvider.RulesetName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RulesetNameRef = rsp.ResolvedReference

	return nil
}
