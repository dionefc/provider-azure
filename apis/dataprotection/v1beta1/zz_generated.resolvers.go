/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta12 "github.com/upbound/provider-azure/apis/azure/v1beta1"
	v1beta11 "github.com/upbound/provider-azure/apis/compute/v1beta1"
	v1beta14 "github.com/upbound/provider-azure/apis/dbforpostgresql/v1beta1"
	v1beta13 "github.com/upbound/provider-azure/apis/keyvault/v1beta1"
	rconfig "github.com/upbound/provider-azure/apis/rconfig"
	v1beta1 "github.com/upbound/provider-azure/apis/storage/v1beta1"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this BackupInstanceBlobStorage.
func (mg *BackupInstanceBlobStorage) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.BackupPolicyID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.BackupPolicyIDRef,
		Selector:     mg.Spec.ForProvider.BackupPolicyIDSelector,
		To: reference.To{
			List:    &BackupPolicyBlobStorageList{},
			Managed: &BackupPolicyBlobStorage{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.BackupPolicyID")
	}
	mg.Spec.ForProvider.BackupPolicyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BackupPolicyIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageAccountID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.StorageAccountIDRef,
		Selector:     mg.Spec.ForProvider.StorageAccountIDSelector,
		To: reference.To{
			List:    &v1beta1.AccountList{},
			Managed: &v1beta1.Account{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.StorageAccountID")
	}
	mg.Spec.ForProvider.StorageAccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.StorageAccountIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VaultID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.VaultIDRef,
		Selector:     mg.Spec.ForProvider.VaultIDSelector,
		To: reference.To{
			List:    &BackupVaultList{},
			Managed: &BackupVault{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VaultID")
	}
	mg.Spec.ForProvider.VaultID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VaultIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this BackupInstanceDisk.
func (mg *BackupInstanceDisk) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.BackupPolicyID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.BackupPolicyIDRef,
		Selector:     mg.Spec.ForProvider.BackupPolicyIDSelector,
		To: reference.To{
			List:    &BackupPolicyDiskList{},
			Managed: &BackupPolicyDisk{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.BackupPolicyID")
	}
	mg.Spec.ForProvider.BackupPolicyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BackupPolicyIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DiskID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.DiskIDRef,
		Selector:     mg.Spec.ForProvider.DiskIDSelector,
		To: reference.To{
			List:    &v1beta11.ManagedDiskList{},
			Managed: &v1beta11.ManagedDisk{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DiskID")
	}
	mg.Spec.ForProvider.DiskID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DiskIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SnapshotResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SnapshotResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.SnapshotResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta12.ResourceGroupList{},
			Managed: &v1beta12.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SnapshotResourceGroupName")
	}
	mg.Spec.ForProvider.SnapshotResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SnapshotResourceGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VaultID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.VaultIDRef,
		Selector:     mg.Spec.ForProvider.VaultIDSelector,
		To: reference.To{
			List:    &BackupVaultList{},
			Managed: &BackupVault{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VaultID")
	}
	mg.Spec.ForProvider.VaultID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VaultIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this BackupInstancePostgreSQL.
func (mg *BackupInstancePostgreSQL) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.BackupPolicyID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.BackupPolicyIDRef,
		Selector:     mg.Spec.ForProvider.BackupPolicyIDSelector,
		To: reference.To{
			List:    &BackupPolicyPostgreSQLList{},
			Managed: &BackupPolicyPostgreSQL{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.BackupPolicyID")
	}
	mg.Spec.ForProvider.BackupPolicyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BackupPolicyIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DatabaseCredentialKeyVaultSecretID),
		Extract:      resource.ExtractParamPath("versionless_id", true),
		Reference:    mg.Spec.ForProvider.DatabaseCredentialKeyVaultSecretIDRef,
		Selector:     mg.Spec.ForProvider.DatabaseCredentialKeyVaultSecretIDSelector,
		To: reference.To{
			List:    &v1beta13.SecretList{},
			Managed: &v1beta13.Secret{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DatabaseCredentialKeyVaultSecretID")
	}
	mg.Spec.ForProvider.DatabaseCredentialKeyVaultSecretID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DatabaseCredentialKeyVaultSecretIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DatabaseID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.DatabaseIDRef,
		Selector:     mg.Spec.ForProvider.DatabaseIDSelector,
		To: reference.To{
			List:    &v1beta14.DatabaseList{},
			Managed: &v1beta14.Database{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DatabaseID")
	}
	mg.Spec.ForProvider.DatabaseID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DatabaseIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VaultID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.VaultIDRef,
		Selector:     mg.Spec.ForProvider.VaultIDSelector,
		To: reference.To{
			List:    &BackupVaultList{},
			Managed: &BackupVault{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VaultID")
	}
	mg.Spec.ForProvider.VaultID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VaultIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this BackupPolicyBlobStorage.
func (mg *BackupPolicyBlobStorage) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VaultID),
		Extract:      rconfig.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.VaultIDRef,
		Selector:     mg.Spec.ForProvider.VaultIDSelector,
		To: reference.To{
			List:    &BackupVaultList{},
			Managed: &BackupVault{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VaultID")
	}
	mg.Spec.ForProvider.VaultID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VaultIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this BackupPolicyDisk.
func (mg *BackupPolicyDisk) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VaultID),
		Extract:      rconfig.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.VaultIDRef,
		Selector:     mg.Spec.ForProvider.VaultIDSelector,
		To: reference.To{
			List:    &BackupVaultList{},
			Managed: &BackupVault{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VaultID")
	}
	mg.Spec.ForProvider.VaultID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VaultIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this BackupPolicyPostgreSQL.
func (mg *BackupPolicyPostgreSQL) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta12.ResourceGroupList{},
			Managed: &v1beta12.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VaultName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VaultNameRef,
		Selector:     mg.Spec.ForProvider.VaultNameSelector,
		To: reference.To{
			List:    &BackupVaultList{},
			Managed: &BackupVault{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VaultName")
	}
	mg.Spec.ForProvider.VaultName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VaultNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this BackupVault.
func (mg *BackupVault) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta12.ResourceGroupList{},
			Managed: &v1beta12.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ResourceGuard.
func (mg *ResourceGuard) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta12.ResourceGroupList{},
			Managed: &v1beta12.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}
