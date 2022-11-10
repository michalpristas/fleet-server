// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package bulk

import (
	json "encoding/json"
	es "github.com/elastic/fleet-server/v7/internal/pkg/es"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonCef4e921DecodeGithubComElasticFleetServerV7InternalPkgBulk(in *jlexer.Lexer, out *bulkStubItem) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "index":
			if in.IsNull() {
				in.Skip()
				out.Index = nil
			} else {
				if out.Index == nil {
					out.Index = new(BulkIndexerResponseItem)
				}
				(*out.Index).UnmarshalEasyJSON(in)
			}
		case "delete":
			if in.IsNull() {
				in.Skip()
				out.Delete = nil
			} else {
				if out.Delete == nil {
					out.Delete = new(BulkIndexerResponseItem)
				}
				(*out.Delete).UnmarshalEasyJSON(in)
			}
		case "create":
			if in.IsNull() {
				in.Skip()
				out.Create = nil
			} else {
				if out.Create == nil {
					out.Create = new(BulkIndexerResponseItem)
				}
				(*out.Create).UnmarshalEasyJSON(in)
			}
		case "update":
			if in.IsNull() {
				in.Skip()
				out.Update = nil
			} else {
				if out.Update == nil {
					out.Update = new(BulkIndexerResponseItem)
				}
				(*out.Update).UnmarshalEasyJSON(in)
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonCef4e921EncodeGithubComElasticFleetServerV7InternalPkgBulk(out *jwriter.Writer, in bulkStubItem) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"index\":"
		out.RawString(prefix[1:])
		if in.Index == nil {
			out.RawString("null")
		} else {
			(*in.Index).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"delete\":"
		out.RawString(prefix)
		if in.Delete == nil {
			out.RawString("null")
		} else {
			(*in.Delete).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"create\":"
		out.RawString(prefix)
		if in.Create == nil {
			out.RawString("null")
		} else {
			(*in.Create).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"update\":"
		out.RawString(prefix)
		if in.Update == nil {
			out.RawString("null")
		} else {
			(*in.Update).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v bulkStubItem) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCef4e921EncodeGithubComElasticFleetServerV7InternalPkgBulk(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v bulkStubItem) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCef4e921EncodeGithubComElasticFleetServerV7InternalPkgBulk(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *bulkStubItem) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCef4e921DecodeGithubComElasticFleetServerV7InternalPkgBulk(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *bulkStubItem) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCef4e921DecodeGithubComElasticFleetServerV7InternalPkgBulk(l, v)
}
func easyjsonCef4e921DecodeGithubComElasticFleetServerV7InternalPkgBulk1(in *jlexer.Lexer, out *bulkIndexerResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "took":
			out.Took = int(in.Int())
		case "errors":
			out.HasErrors = bool(in.Bool())
		case "items":
			if in.IsNull() {
				in.Skip()
				out.Items = nil
			} else {
				in.Delim('[')
				if out.Items == nil {
					if !in.IsDelim(']') {
						out.Items = make([]bulkStubItem, 0, 2)
					} else {
						out.Items = []bulkStubItem{}
					}
				} else {
					out.Items = (out.Items)[:0]
				}
				for !in.IsDelim(']') {
					var v1 bulkStubItem
					(v1).UnmarshalEasyJSON(in)
					out.Items = append(out.Items, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonCef4e921EncodeGithubComElasticFleetServerV7InternalPkgBulk1(out *jwriter.Writer, in bulkIndexerResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"took\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Took))
	}
	{
		const prefix string = ",\"errors\":"
		out.RawString(prefix)
		out.Bool(bool(in.HasErrors))
	}
	if len(in.Items) != 0 {
		const prefix string = ",\"items\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v2, v3 := range in.Items {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v bulkIndexerResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCef4e921EncodeGithubComElasticFleetServerV7InternalPkgBulk1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v bulkIndexerResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCef4e921EncodeGithubComElasticFleetServerV7InternalPkgBulk1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *bulkIndexerResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCef4e921DecodeGithubComElasticFleetServerV7InternalPkgBulk1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *bulkIndexerResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCef4e921DecodeGithubComElasticFleetServerV7InternalPkgBulk1(l, v)
}
func easyjsonCef4e921DecodeGithubComElasticFleetServerV7InternalPkgBulk2(in *jlexer.Lexer, out *MsearchResponseItem) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "status":
			out.Status = int(in.Int())
		case "took":
			out.Took = uint64(in.Uint64())
		case "timed_out":
			out.TimedOut = bool(in.Bool())
		case "_shards":
			easyjsonCef4e921Decode(in, &out.Shards)
		case "hits":
			easyjsonCef4e921DecodeGithubComElasticFleetServerV7InternalPkgEs(in, &out.Hits)
		case "aggregations":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Aggregations = make(map[string]es.Aggregation)
				} else {
					out.Aggregations = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v4 es.Aggregation
					easyjsonCef4e921DecodeGithubComElasticFleetServerV7InternalPkgEs1(in, &v4)
					(out.Aggregations)[key] = v4
					in.WantComma()
				}
				in.Delim('}')
			}
		case "error":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Error).UnmarshalJSON(data))
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonCef4e921EncodeGithubComElasticFleetServerV7InternalPkgBulk2(out *jwriter.Writer, in MsearchResponseItem) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Status))
	}
	{
		const prefix string = ",\"took\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.Took))
	}
	{
		const prefix string = ",\"timed_out\":"
		out.RawString(prefix)
		out.Bool(bool(in.TimedOut))
	}
	{
		const prefix string = ",\"_shards\":"
		out.RawString(prefix)
		easyjsonCef4e921Encode(out, in.Shards)
	}
	{
		const prefix string = ",\"hits\":"
		out.RawString(prefix)
		easyjsonCef4e921EncodeGithubComElasticFleetServerV7InternalPkgEs(out, in.Hits)
	}
	if len(in.Aggregations) != 0 {
		const prefix string = ",\"aggregations\":"
		out.RawString(prefix)
		{
			out.RawByte('{')
			v5First := true
			for v5Name, v5Value := range in.Aggregations {
				if v5First {
					v5First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v5Name))
				out.RawByte(':')
				easyjsonCef4e921EncodeGithubComElasticFleetServerV7InternalPkgEs1(out, v5Value)
			}
			out.RawByte('}')
		}
	}
	if len(in.Error) != 0 {
		const prefix string = ",\"error\":"
		out.RawString(prefix)
		out.Raw((in.Error).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MsearchResponseItem) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCef4e921EncodeGithubComElasticFleetServerV7InternalPkgBulk2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MsearchResponseItem) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCef4e921EncodeGithubComElasticFleetServerV7InternalPkgBulk2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MsearchResponseItem) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCef4e921DecodeGithubComElasticFleetServerV7InternalPkgBulk2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MsearchResponseItem) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCef4e921DecodeGithubComElasticFleetServerV7InternalPkgBulk2(l, v)
}
func easyjsonCef4e921DecodeGithubComElasticFleetServerV7InternalPkgEs1(in *jlexer.Lexer, out *es.Aggregation) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "value":
			out.Value = float64(in.Float64())
		case "doc_count_error_upper_bound":
			out.DocCountErrorUpperBound = int64(in.Int64())
		case "sum_other_doc_count":
			out.SumOtherDocCount = int64(in.Int64())
		case "buckets":
			if in.IsNull() {
				in.Skip()
				out.Buckets = nil
			} else {
				in.Delim('[')
				if out.Buckets == nil {
					if !in.IsDelim(']') {
						out.Buckets = make([]es.Bucket, 0, 2)
					} else {
						out.Buckets = []es.Bucket{}
					}
				} else {
					out.Buckets = (out.Buckets)[:0]
				}
				for !in.IsDelim(']') {
					var v6 es.Bucket
					if data := in.Raw(); in.Ok() {
						in.AddError((v6).UnmarshalJSON(data))
					}
					out.Buckets = append(out.Buckets, v6)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonCef4e921EncodeGithubComElasticFleetServerV7InternalPkgEs1(out *jwriter.Writer, in es.Aggregation) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"value\":"
		out.RawString(prefix[1:])
		out.Float64(float64(in.Value))
	}
	{
		const prefix string = ",\"doc_count_error_upper_bound\":"
		out.RawString(prefix)
		out.Int64(int64(in.DocCountErrorUpperBound))
	}
	{
		const prefix string = ",\"sum_other_doc_count\":"
		out.RawString(prefix)
		out.Int64(int64(in.SumOtherDocCount))
	}
	if len(in.Buckets) != 0 {
		const prefix string = ",\"buckets\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v7, v8 := range in.Buckets {
				if v7 > 0 {
					out.RawByte(',')
				}
				easyjsonCef4e921EncodeGithubComElasticFleetServerV7InternalPkgEs2(out, v8)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjsonCef4e921DecodeGithubComElasticFleetServerV7InternalPkgEs2(in *jlexer.Lexer, out *es.Bucket) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "key":
			out.Key = string(in.String())
		case "doc_count":
			out.DocCount = int64(in.Int64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonCef4e921EncodeGithubComElasticFleetServerV7InternalPkgEs2(out *jwriter.Writer, in es.Bucket) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"key\":"
		out.RawString(prefix[1:])
		out.String(string(in.Key))
	}
	{
		const prefix string = ",\"doc_count\":"
		out.RawString(prefix)
		out.Int64(int64(in.DocCount))
	}
	out.RawByte('}')
}
func easyjsonCef4e921DecodeGithubComElasticFleetServerV7InternalPkgEs(in *jlexer.Lexer, out *es.HitsT) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "hits":
			if in.IsNull() {
				in.Skip()
				out.Hits = nil
			} else {
				in.Delim('[')
				if out.Hits == nil {
					if !in.IsDelim(']') {
						out.Hits = make([]es.HitT, 0, 0)
					} else {
						out.Hits = []es.HitT{}
					}
				} else {
					out.Hits = (out.Hits)[:0]
				}
				for !in.IsDelim(']') {
					var v9 es.HitT
					easyjsonCef4e921DecodeGithubComElasticFleetServerV7InternalPkgEs3(in, &v9)
					out.Hits = append(out.Hits, v9)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "total":
			easyjsonCef4e921Decode1(in, &out.Total)
		case "max_score":
			if in.IsNull() {
				in.Skip()
				out.MaxScore = nil
			} else {
				if out.MaxScore == nil {
					out.MaxScore = new(float64)
				}
				*out.MaxScore = float64(in.Float64())
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonCef4e921EncodeGithubComElasticFleetServerV7InternalPkgEs(out *jwriter.Writer, in es.HitsT) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"hits\":"
		out.RawString(prefix[1:])
		if in.Hits == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v10, v11 := range in.Hits {
				if v10 > 0 {
					out.RawByte(',')
				}
				easyjsonCef4e921EncodeGithubComElasticFleetServerV7InternalPkgEs3(out, v11)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"total\":"
		out.RawString(prefix)
		easyjsonCef4e921Encode1(out, in.Total)
	}
	{
		const prefix string = ",\"max_score\":"
		out.RawString(prefix)
		if in.MaxScore == nil {
			out.RawString("null")
		} else {
			out.Float64(float64(*in.MaxScore))
		}
	}
	out.RawByte('}')
}
func easyjsonCef4e921Decode1(in *jlexer.Lexer, out *struct {
	Relation string `json:"relation"`
	Value    uint64 `json:"value"`
}) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "relation":
			out.Relation = string(in.String())
		case "value":
			out.Value = uint64(in.Uint64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonCef4e921Encode1(out *jwriter.Writer, in struct {
	Relation string `json:"relation"`
	Value    uint64 `json:"value"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"relation\":"
		out.RawString(prefix[1:])
		out.String(string(in.Relation))
	}
	{
		const prefix string = ",\"value\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.Value))
	}
	out.RawByte('}')
}
func easyjsonCef4e921DecodeGithubComElasticFleetServerV7InternalPkgEs3(in *jlexer.Lexer, out *es.HitT) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "_id":
			out.ID = string(in.String())
		case "_seq_no":
			out.SeqNo = int64(in.Int64())
		case "version":
			out.Version = int64(in.Int64())
		case "_index":
			out.Index = string(in.String())
		case "_source":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Source).UnmarshalJSON(data))
			}
		case "_score":
			if in.IsNull() {
				in.Skip()
				out.Score = nil
			} else {
				if out.Score == nil {
					out.Score = new(float64)
				}
				*out.Score = float64(in.Float64())
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonCef4e921EncodeGithubComElasticFleetServerV7InternalPkgEs3(out *jwriter.Writer, in es.HitT) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"_id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"_seq_no\":"
		out.RawString(prefix)
		out.Int64(int64(in.SeqNo))
	}
	{
		const prefix string = ",\"version\":"
		out.RawString(prefix)
		out.Int64(int64(in.Version))
	}
	{
		const prefix string = ",\"_index\":"
		out.RawString(prefix)
		out.String(string(in.Index))
	}
	{
		const prefix string = ",\"_source\":"
		out.RawString(prefix)
		out.Raw((in.Source).MarshalJSON())
	}
	{
		const prefix string = ",\"_score\":"
		out.RawString(prefix)
		if in.Score == nil {
			out.RawString("null")
		} else {
			out.Float64(float64(*in.Score))
		}
	}
	out.RawByte('}')
}
func easyjsonCef4e921Decode(in *jlexer.Lexer, out *struct {
	Total      uint64 `json:"total"`
	Successful uint64 `json:"successful"`
	Skipped    uint64 `json:"skipped"`
	Failed     uint64 `json:"failed"`
}) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "total":
			out.Total = uint64(in.Uint64())
		case "successful":
			out.Successful = uint64(in.Uint64())
		case "skipped":
			out.Skipped = uint64(in.Uint64())
		case "failed":
			out.Failed = uint64(in.Uint64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonCef4e921Encode(out *jwriter.Writer, in struct {
	Total      uint64 `json:"total"`
	Successful uint64 `json:"successful"`
	Skipped    uint64 `json:"skipped"`
	Failed     uint64 `json:"failed"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"total\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.Total))
	}
	{
		const prefix string = ",\"successful\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.Successful))
	}
	{
		const prefix string = ",\"skipped\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.Skipped))
	}
	{
		const prefix string = ",\"failed\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.Failed))
	}
	out.RawByte('}')
}
func easyjsonCef4e921DecodeGithubComElasticFleetServerV7InternalPkgBulk3(in *jlexer.Lexer, out *MsearchResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "responses":
			if in.IsNull() {
				in.Skip()
				out.Responses = nil
			} else {
				in.Delim('[')
				if out.Responses == nil {
					if !in.IsDelim(']') {
						out.Responses = make([]MsearchResponseItem, 0, 0)
					} else {
						out.Responses = []MsearchResponseItem{}
					}
				} else {
					out.Responses = (out.Responses)[:0]
				}
				for !in.IsDelim(']') {
					var v12 MsearchResponseItem
					(v12).UnmarshalEasyJSON(in)
					out.Responses = append(out.Responses, v12)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "took":
			out.Took = int(in.Int())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonCef4e921EncodeGithubComElasticFleetServerV7InternalPkgBulk3(out *jwriter.Writer, in MsearchResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"responses\":"
		out.RawString(prefix[1:])
		if in.Responses == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v13, v14 := range in.Responses {
				if v13 > 0 {
					out.RawByte(',')
				}
				(v14).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"took\":"
		out.RawString(prefix)
		out.Int(int(in.Took))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MsearchResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCef4e921EncodeGithubComElasticFleetServerV7InternalPkgBulk3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MsearchResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCef4e921EncodeGithubComElasticFleetServerV7InternalPkgBulk3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MsearchResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCef4e921DecodeGithubComElasticFleetServerV7InternalPkgBulk3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MsearchResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCef4e921DecodeGithubComElasticFleetServerV7InternalPkgBulk3(l, v)
}
func easyjsonCef4e921DecodeGithubComElasticFleetServerV7InternalPkgBulk4(in *jlexer.Lexer, out *MgetResponseItem) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "found":
			out.Found = bool(in.Bool())
		case "_source":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Source).UnmarshalJSON(data))
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonCef4e921EncodeGithubComElasticFleetServerV7InternalPkgBulk4(out *jwriter.Writer, in MgetResponseItem) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"found\":"
		out.RawString(prefix[1:])
		out.Bool(bool(in.Found))
	}
	{
		const prefix string = ",\"_source\":"
		out.RawString(prefix)
		out.Raw((in.Source).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MgetResponseItem) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCef4e921EncodeGithubComElasticFleetServerV7InternalPkgBulk4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MgetResponseItem) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCef4e921EncodeGithubComElasticFleetServerV7InternalPkgBulk4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MgetResponseItem) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCef4e921DecodeGithubComElasticFleetServerV7InternalPkgBulk4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MgetResponseItem) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCef4e921DecodeGithubComElasticFleetServerV7InternalPkgBulk4(l, v)
}
func easyjsonCef4e921DecodeGithubComElasticFleetServerV7InternalPkgBulk5(in *jlexer.Lexer, out *MgetResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "docs":
			if in.IsNull() {
				in.Skip()
				out.Items = nil
			} else {
				in.Delim('[')
				if out.Items == nil {
					if !in.IsDelim(']') {
						out.Items = make([]MgetResponseItem, 0, 2)
					} else {
						out.Items = []MgetResponseItem{}
					}
				} else {
					out.Items = (out.Items)[:0]
				}
				for !in.IsDelim(']') {
					var v15 MgetResponseItem
					(v15).UnmarshalEasyJSON(in)
					out.Items = append(out.Items, v15)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonCef4e921EncodeGithubComElasticFleetServerV7InternalPkgBulk5(out *jwriter.Writer, in MgetResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"docs\":"
		out.RawString(prefix[1:])
		if in.Items == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v16, v17 := range in.Items {
				if v16 > 0 {
					out.RawByte(',')
				}
				(v17).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MgetResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCef4e921EncodeGithubComElasticFleetServerV7InternalPkgBulk5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MgetResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCef4e921EncodeGithubComElasticFleetServerV7InternalPkgBulk5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MgetResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCef4e921DecodeGithubComElasticFleetServerV7InternalPkgBulk5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MgetResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCef4e921DecodeGithubComElasticFleetServerV7InternalPkgBulk5(l, v)
}
func easyjsonCef4e921DecodeGithubComElasticFleetServerV7InternalPkgBulk6(in *jlexer.Lexer, out *BulkIndexerResponseItem) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "_id":
			out.DocumentID = string(in.String())
		case "status":
			out.Status = int(in.Int())
		case "error":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Error).UnmarshalJSON(data))
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonCef4e921EncodeGithubComElasticFleetServerV7InternalPkgBulk6(out *jwriter.Writer, in BulkIndexerResponseItem) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"_id\":"
		out.RawString(prefix[1:])
		out.String(string(in.DocumentID))
	}
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix)
		out.Int(int(in.Status))
	}
	if len(in.Error) != 0 {
		const prefix string = ",\"error\":"
		out.RawString(prefix)
		out.Raw((in.Error).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v BulkIndexerResponseItem) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCef4e921EncodeGithubComElasticFleetServerV7InternalPkgBulk6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BulkIndexerResponseItem) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCef4e921EncodeGithubComElasticFleetServerV7InternalPkgBulk6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BulkIndexerResponseItem) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCef4e921DecodeGithubComElasticFleetServerV7InternalPkgBulk6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BulkIndexerResponseItem) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCef4e921DecodeGithubComElasticFleetServerV7InternalPkgBulk6(l, v)
}
