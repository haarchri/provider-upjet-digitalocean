/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	v1alpha1 "github.com/crossplane-contrib/provider-upjet-digitalocean/apis/droplet/v1alpha1"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Attachment.
func (mg *Attachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.DropletID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.DropletIDRef,
		Selector:     mg.Spec.ForProvider.DropletIDSelector,
		To: reference.To{
			List:    &v1alpha1.DropletList{},
			Managed: &v1alpha1.Droplet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DropletID")
	}
	mg.Spec.ForProvider.DropletID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DropletIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VolumeID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VolumeIDRef,
		Selector:     mg.Spec.ForProvider.VolumeIDSelector,
		To: reference.To{
			List:    &VolumeList{},
			Managed: &Volume{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VolumeID")
	}
	mg.Spec.ForProvider.VolumeID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VolumeIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Snapshot.
func (mg *Snapshot) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VolumeID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VolumeIDRef,
		Selector:     mg.Spec.ForProvider.VolumeIDSelector,
		To: reference.To{
			List:    &VolumeList{},
			Managed: &Volume{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VolumeID")
	}
	mg.Spec.ForProvider.VolumeID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VolumeIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Volume.
func (mg *Volume) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SnapshotID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SnapshotIDRef,
		Selector:     mg.Spec.ForProvider.SnapshotIDSelector,
		To: reference.To{
			List:    &SnapshotList{},
			Managed: &Snapshot{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SnapshotID")
	}
	mg.Spec.ForProvider.SnapshotID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SnapshotIDRef = rsp.ResolvedReference

	return nil
}
