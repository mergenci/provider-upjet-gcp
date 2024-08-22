// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta2

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	apisresolver "github.com/upbound/provider-gcp/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *ConnectionProfile) ResolveReferences( // ResolveReferences of this ConnectionProfile.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	if mg.Spec.ForProvider.PostgresqlProfile != nil {
		{
			m, l, err = apisresolver.GetManagedResource("sql.gcp.upbound.io", "v1beta1", "Database", "DatabaseList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PostgresqlProfile.Database),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.PostgresqlProfile.DatabaseRef,
				Selector:     mg.Spec.ForProvider.PostgresqlProfile.DatabaseSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.PostgresqlProfile.Database")
		}
		mg.Spec.ForProvider.PostgresqlProfile.Database = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.PostgresqlProfile.DatabaseRef = rsp.ResolvedReference

	}
	if mg.Spec.ForProvider.PostgresqlProfile != nil {
		{
			m, l, err = apisresolver.GetManagedResource("sql.gcp.upbound.io", "v1beta2", "DatabaseInstance", "DatabaseInstanceList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PostgresqlProfile.Hostname),
				Extract:      resource.ExtractParamPath("public_ip_address", true),
				Reference:    mg.Spec.ForProvider.PostgresqlProfile.HostnameRef,
				Selector:     mg.Spec.ForProvider.PostgresqlProfile.HostnameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.PostgresqlProfile.Hostname")
		}
		mg.Spec.ForProvider.PostgresqlProfile.Hostname = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.PostgresqlProfile.HostnameRef = rsp.ResolvedReference

	}
	if mg.Spec.ForProvider.PostgresqlProfile != nil {
		{
			m, l, err = apisresolver.GetManagedResource("sql.gcp.upbound.io", "v1beta2", "User", "UserList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PostgresqlProfile.Username),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.PostgresqlProfile.UsernameRef,
				Selector:     mg.Spec.ForProvider.PostgresqlProfile.UsernameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.PostgresqlProfile.Username")
		}
		mg.Spec.ForProvider.PostgresqlProfile.Username = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.PostgresqlProfile.UsernameRef = rsp.ResolvedReference

	}
	if mg.Spec.ForProvider.PrivateConnectivity != nil {
		{
			m, l, err = apisresolver.GetManagedResource("datastream.gcp.upbound.io", "v1beta2", "PrivateConnection", "PrivateConnectionList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PrivateConnectivity.PrivateConnection),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.PrivateConnectivity.PrivateConnectionRef,
				Selector:     mg.Spec.ForProvider.PrivateConnectivity.PrivateConnectionSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.PrivateConnectivity.PrivateConnection")
		}
		mg.Spec.ForProvider.PrivateConnectivity.PrivateConnection = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.PrivateConnectivity.PrivateConnectionRef = rsp.ResolvedReference

	}
	if mg.Spec.ForProvider.SQLServerProfile != nil {
		{
			m, l, err = apisresolver.GetManagedResource("sql.gcp.upbound.io", "v1beta1", "Database", "DatabaseList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SQLServerProfile.Database),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.SQLServerProfile.DatabaseRef,
				Selector:     mg.Spec.ForProvider.SQLServerProfile.DatabaseSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.SQLServerProfile.Database")
		}
		mg.Spec.ForProvider.SQLServerProfile.Database = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.SQLServerProfile.DatabaseRef = rsp.ResolvedReference

	}
	if mg.Spec.ForProvider.SQLServerProfile != nil {
		{
			m, l, err = apisresolver.GetManagedResource("sql.gcp.upbound.io", "v1beta2", "DatabaseInstance", "DatabaseInstanceList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SQLServerProfile.Hostname),
				Extract:      resource.ExtractParamPath("public_ip_address", true),
				Reference:    mg.Spec.ForProvider.SQLServerProfile.HostnameRef,
				Selector:     mg.Spec.ForProvider.SQLServerProfile.HostnameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.SQLServerProfile.Hostname")
		}
		mg.Spec.ForProvider.SQLServerProfile.Hostname = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.SQLServerProfile.HostnameRef = rsp.ResolvedReference

	}
	if mg.Spec.ForProvider.SQLServerProfile != nil {
		{
			m, l, err = apisresolver.GetManagedResource("sql.gcp.upbound.io", "v1beta2", "User", "UserList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SQLServerProfile.Username),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.SQLServerProfile.UsernameRef,
				Selector:     mg.Spec.ForProvider.SQLServerProfile.UsernameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.SQLServerProfile.Username")
		}
		mg.Spec.ForProvider.SQLServerProfile.Username = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.SQLServerProfile.UsernameRef = rsp.ResolvedReference

	}
	if mg.Spec.InitProvider.PostgresqlProfile != nil {
		{
			m, l, err = apisresolver.GetManagedResource("sql.gcp.upbound.io", "v1beta1", "Database", "DatabaseList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.PostgresqlProfile.Database),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.InitProvider.PostgresqlProfile.DatabaseRef,
				Selector:     mg.Spec.InitProvider.PostgresqlProfile.DatabaseSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.PostgresqlProfile.Database")
		}
		mg.Spec.InitProvider.PostgresqlProfile.Database = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.PostgresqlProfile.DatabaseRef = rsp.ResolvedReference

	}
	if mg.Spec.InitProvider.PostgresqlProfile != nil {
		{
			m, l, err = apisresolver.GetManagedResource("sql.gcp.upbound.io", "v1beta2", "DatabaseInstance", "DatabaseInstanceList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.PostgresqlProfile.Hostname),
				Extract:      resource.ExtractParamPath("public_ip_address", true),
				Reference:    mg.Spec.InitProvider.PostgresqlProfile.HostnameRef,
				Selector:     mg.Spec.InitProvider.PostgresqlProfile.HostnameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.PostgresqlProfile.Hostname")
		}
		mg.Spec.InitProvider.PostgresqlProfile.Hostname = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.PostgresqlProfile.HostnameRef = rsp.ResolvedReference

	}
	if mg.Spec.InitProvider.PostgresqlProfile != nil {
		{
			m, l, err = apisresolver.GetManagedResource("sql.gcp.upbound.io", "v1beta2", "User", "UserList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.PostgresqlProfile.Username),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.InitProvider.PostgresqlProfile.UsernameRef,
				Selector:     mg.Spec.InitProvider.PostgresqlProfile.UsernameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.PostgresqlProfile.Username")
		}
		mg.Spec.InitProvider.PostgresqlProfile.Username = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.PostgresqlProfile.UsernameRef = rsp.ResolvedReference

	}
	if mg.Spec.InitProvider.PrivateConnectivity != nil {
		{
			m, l, err = apisresolver.GetManagedResource("datastream.gcp.upbound.io", "v1beta2", "PrivateConnection", "PrivateConnectionList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.PrivateConnectivity.PrivateConnection),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.PrivateConnectivity.PrivateConnectionRef,
				Selector:     mg.Spec.InitProvider.PrivateConnectivity.PrivateConnectionSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.PrivateConnectivity.PrivateConnection")
		}
		mg.Spec.InitProvider.PrivateConnectivity.PrivateConnection = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.PrivateConnectivity.PrivateConnectionRef = rsp.ResolvedReference

	}
	if mg.Spec.InitProvider.SQLServerProfile != nil {
		{
			m, l, err = apisresolver.GetManagedResource("sql.gcp.upbound.io", "v1beta1", "Database", "DatabaseList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SQLServerProfile.Database),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.InitProvider.SQLServerProfile.DatabaseRef,
				Selector:     mg.Spec.InitProvider.SQLServerProfile.DatabaseSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.SQLServerProfile.Database")
		}
		mg.Spec.InitProvider.SQLServerProfile.Database = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.SQLServerProfile.DatabaseRef = rsp.ResolvedReference

	}
	if mg.Spec.InitProvider.SQLServerProfile != nil {
		{
			m, l, err = apisresolver.GetManagedResource("sql.gcp.upbound.io", "v1beta2", "DatabaseInstance", "DatabaseInstanceList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SQLServerProfile.Hostname),
				Extract:      resource.ExtractParamPath("public_ip_address", true),
				Reference:    mg.Spec.InitProvider.SQLServerProfile.HostnameRef,
				Selector:     mg.Spec.InitProvider.SQLServerProfile.HostnameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.SQLServerProfile.Hostname")
		}
		mg.Spec.InitProvider.SQLServerProfile.Hostname = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.SQLServerProfile.HostnameRef = rsp.ResolvedReference

	}
	if mg.Spec.InitProvider.SQLServerProfile != nil {
		{
			m, l, err = apisresolver.GetManagedResource("sql.gcp.upbound.io", "v1beta2", "User", "UserList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SQLServerProfile.Username),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.InitProvider.SQLServerProfile.UsernameRef,
				Selector:     mg.Spec.InitProvider.SQLServerProfile.UsernameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.SQLServerProfile.Username")
		}
		mg.Spec.InitProvider.SQLServerProfile.Username = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.SQLServerProfile.UsernameRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this PrivateConnection.
func (mg *PrivateConnection) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	if mg.Spec.ForProvider.VPCPeeringConfig != nil {
		{
			m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta1", "Network", "NetworkList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPCPeeringConfig.VPC),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.VPCPeeringConfig.VPCRef,
				Selector:     mg.Spec.ForProvider.VPCPeeringConfig.VPCSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.VPCPeeringConfig.VPC")
		}
		mg.Spec.ForProvider.VPCPeeringConfig.VPC = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.VPCPeeringConfig.VPCRef = rsp.ResolvedReference

	}
	if mg.Spec.InitProvider.VPCPeeringConfig != nil {
		{
			m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta1", "Network", "NetworkList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.VPCPeeringConfig.VPC),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.VPCPeeringConfig.VPCRef,
				Selector:     mg.Spec.InitProvider.VPCPeeringConfig.VPCSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.VPCPeeringConfig.VPC")
		}
		mg.Spec.InitProvider.VPCPeeringConfig.VPC = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.VPCPeeringConfig.VPCRef = rsp.ResolvedReference

	}

	return nil
}