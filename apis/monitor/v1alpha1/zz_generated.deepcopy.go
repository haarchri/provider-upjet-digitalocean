//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Alert) DeepCopyInto(out *Alert) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Alert.
func (in *Alert) DeepCopy() *Alert {
	if in == nil {
		return nil
	}
	out := new(Alert)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Alert) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertInitParameters) DeepCopyInto(out *AlertInitParameters) {
	*out = *in
	if in.Alerts != nil {
		in, out := &in.Alerts, &out.Alerts
		*out = make([]AlertsInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Compare != nil {
		in, out := &in.Compare, &out.Compare
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(float64)
		**out = **in
	}
	if in.Window != nil {
		in, out := &in.Window, &out.Window
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertInitParameters.
func (in *AlertInitParameters) DeepCopy() *AlertInitParameters {
	if in == nil {
		return nil
	}
	out := new(AlertInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertList) DeepCopyInto(out *AlertList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Alert, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertList.
func (in *AlertList) DeepCopy() *AlertList {
	if in == nil {
		return nil
	}
	out := new(AlertList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlertList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertObservation) DeepCopyInto(out *AlertObservation) {
	*out = *in
	if in.Alerts != nil {
		in, out := &in.Alerts, &out.Alerts
		*out = make([]AlertsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Compare != nil {
		in, out := &in.Compare, &out.Compare
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Entities != nil {
		in, out := &in.Entities, &out.Entities
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.UUID != nil {
		in, out := &in.UUID, &out.UUID
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(float64)
		**out = **in
	}
	if in.Window != nil {
		in, out := &in.Window, &out.Window
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertObservation.
func (in *AlertObservation) DeepCopy() *AlertObservation {
	if in == nil {
		return nil
	}
	out := new(AlertObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertParameters) DeepCopyInto(out *AlertParameters) {
	*out = *in
	if in.Alerts != nil {
		in, out := &in.Alerts, &out.Alerts
		*out = make([]AlertsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Compare != nil {
		in, out := &in.Compare, &out.Compare
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Entities != nil {
		in, out := &in.Entities, &out.Entities
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.EntitiesRefs != nil {
		in, out := &in.EntitiesRefs, &out.EntitiesRefs
		*out = make([]v1.Reference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EntitiesSelector != nil {
		in, out := &in.EntitiesSelector, &out.EntitiesSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(float64)
		**out = **in
	}
	if in.Window != nil {
		in, out := &in.Window, &out.Window
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertParameters.
func (in *AlertParameters) DeepCopy() *AlertParameters {
	if in == nil {
		return nil
	}
	out := new(AlertParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertSpec) DeepCopyInto(out *AlertSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertSpec.
func (in *AlertSpec) DeepCopy() *AlertSpec {
	if in == nil {
		return nil
	}
	out := new(AlertSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertStatus) DeepCopyInto(out *AlertStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertStatus.
func (in *AlertStatus) DeepCopy() *AlertStatus {
	if in == nil {
		return nil
	}
	out := new(AlertStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertsInitParameters) DeepCopyInto(out *AlertsInitParameters) {
	*out = *in
	if in.Email != nil {
		in, out := &in.Email, &out.Email
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Slack != nil {
		in, out := &in.Slack, &out.Slack
		*out = make([]SlackInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertsInitParameters.
func (in *AlertsInitParameters) DeepCopy() *AlertsInitParameters {
	if in == nil {
		return nil
	}
	out := new(AlertsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertsObservation) DeepCopyInto(out *AlertsObservation) {
	*out = *in
	if in.Email != nil {
		in, out := &in.Email, &out.Email
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Slack != nil {
		in, out := &in.Slack, &out.Slack
		*out = make([]SlackObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertsObservation.
func (in *AlertsObservation) DeepCopy() *AlertsObservation {
	if in == nil {
		return nil
	}
	out := new(AlertsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertsParameters) DeepCopyInto(out *AlertsParameters) {
	*out = *in
	if in.Email != nil {
		in, out := &in.Email, &out.Email
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Slack != nil {
		in, out := &in.Slack, &out.Slack
		*out = make([]SlackParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertsParameters.
func (in *AlertsParameters) DeepCopy() *AlertsParameters {
	if in == nil {
		return nil
	}
	out := new(AlertsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlackInitParameters) DeepCopyInto(out *SlackInitParameters) {
	*out = *in
	if in.Channel != nil {
		in, out := &in.Channel, &out.Channel
		*out = new(string)
		**out = **in
	}
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlackInitParameters.
func (in *SlackInitParameters) DeepCopy() *SlackInitParameters {
	if in == nil {
		return nil
	}
	out := new(SlackInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlackObservation) DeepCopyInto(out *SlackObservation) {
	*out = *in
	if in.Channel != nil {
		in, out := &in.Channel, &out.Channel
		*out = new(string)
		**out = **in
	}
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlackObservation.
func (in *SlackObservation) DeepCopy() *SlackObservation {
	if in == nil {
		return nil
	}
	out := new(SlackObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlackParameters) DeepCopyInto(out *SlackParameters) {
	*out = *in
	if in.Channel != nil {
		in, out := &in.Channel, &out.Channel
		*out = new(string)
		**out = **in
	}
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlackParameters.
func (in *SlackParameters) DeepCopy() *SlackParameters {
	if in == nil {
		return nil
	}
	out := new(SlackParameters)
	in.DeepCopyInto(out)
	return out
}
