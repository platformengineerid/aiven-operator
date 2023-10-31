//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright (c) 2022 Aiven, Helsinki, Finland. https://aiven.io/

// Code generated by controller-gen. DO NOT EDIT.

package metricsuserconfig

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricsUserConfig) DeepCopyInto(out *MetricsUserConfig) {
	*out = *in
	if in.Database != nil {
		in, out := &in.Database, &out.Database
		*out = new(string)
		**out = **in
	}
	if in.RetentionDays != nil {
		in, out := &in.RetentionDays, &out.RetentionDays
		*out = new(int)
		**out = **in
	}
	if in.RoUsername != nil {
		in, out := &in.RoUsername, &out.RoUsername
		*out = new(string)
		**out = **in
	}
	if in.SourceMysql != nil {
		in, out := &in.SourceMysql, &out.SourceMysql
		*out = new(SourceMysql)
		(*in).DeepCopyInto(*out)
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsUserConfig.
func (in *MetricsUserConfig) DeepCopy() *MetricsUserConfig {
	if in == nil {
		return nil
	}
	out := new(MetricsUserConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceMysql) DeepCopyInto(out *SourceMysql) {
	*out = *in
	if in.Telegraf != nil {
		in, out := &in.Telegraf, &out.Telegraf
		*out = new(Telegraf)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceMysql.
func (in *SourceMysql) DeepCopy() *SourceMysql {
	if in == nil {
		return nil
	}
	out := new(SourceMysql)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Telegraf) DeepCopyInto(out *Telegraf) {
	*out = *in
	if in.GatherEventWaits != nil {
		in, out := &in.GatherEventWaits, &out.GatherEventWaits
		*out = new(bool)
		**out = **in
	}
	if in.GatherFileEventsStats != nil {
		in, out := &in.GatherFileEventsStats, &out.GatherFileEventsStats
		*out = new(bool)
		**out = **in
	}
	if in.GatherIndexIoWaits != nil {
		in, out := &in.GatherIndexIoWaits, &out.GatherIndexIoWaits
		*out = new(bool)
		**out = **in
	}
	if in.GatherInfoSchemaAutoInc != nil {
		in, out := &in.GatherInfoSchemaAutoInc, &out.GatherInfoSchemaAutoInc
		*out = new(bool)
		**out = **in
	}
	if in.GatherInnodbMetrics != nil {
		in, out := &in.GatherInnodbMetrics, &out.GatherInnodbMetrics
		*out = new(bool)
		**out = **in
	}
	if in.GatherPerfEventsStatements != nil {
		in, out := &in.GatherPerfEventsStatements, &out.GatherPerfEventsStatements
		*out = new(bool)
		**out = **in
	}
	if in.GatherProcessList != nil {
		in, out := &in.GatherProcessList, &out.GatherProcessList
		*out = new(bool)
		**out = **in
	}
	if in.GatherSlaveStatus != nil {
		in, out := &in.GatherSlaveStatus, &out.GatherSlaveStatus
		*out = new(bool)
		**out = **in
	}
	if in.GatherTableIoWaits != nil {
		in, out := &in.GatherTableIoWaits, &out.GatherTableIoWaits
		*out = new(bool)
		**out = **in
	}
	if in.GatherTableLockWaits != nil {
		in, out := &in.GatherTableLockWaits, &out.GatherTableLockWaits
		*out = new(bool)
		**out = **in
	}
	if in.GatherTableSchema != nil {
		in, out := &in.GatherTableSchema, &out.GatherTableSchema
		*out = new(bool)
		**out = **in
	}
	if in.PerfEventsStatementsDigestTextLimit != nil {
		in, out := &in.PerfEventsStatementsDigestTextLimit, &out.PerfEventsStatementsDigestTextLimit
		*out = new(int)
		**out = **in
	}
	if in.PerfEventsStatementsLimit != nil {
		in, out := &in.PerfEventsStatementsLimit, &out.PerfEventsStatementsLimit
		*out = new(int)
		**out = **in
	}
	if in.PerfEventsStatementsTimeLimit != nil {
		in, out := &in.PerfEventsStatementsTimeLimit, &out.PerfEventsStatementsTimeLimit
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Telegraf.
func (in *Telegraf) DeepCopy() *Telegraf {
	if in == nil {
		return nil
	}
	out := new(Telegraf)
	in.DeepCopyInto(out)
	return out
}