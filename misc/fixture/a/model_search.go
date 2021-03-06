// generated by smg -type Sample -output misc/fixture/a/model_search.go misc/fixture/a; DO NOT EDIT

package a

import (
	"bytes"
	"errors"
	"github.com/favclip/smg/smgutils"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/search"
	"time"
)

// SampleSearch best match Search API wrapper for Sample.
type SampleSearch struct {
	src *Sample

	Foo string
}

// Searchfy converts *Sample to *SampleSearch.
func (src *Sample) Searchfy() (*SampleSearch, error) {
	if src == nil {
		return nil, nil
	}
	dest := &SampleSearch{}
	dest.src = src
	dest.Foo = src.Foo
	return dest, nil
}

// NewSampleSearch create new *SampleSearchBuilder.
func NewSampleSearch() *SampleSearchBuilder {
	op := &smgutils.Op{}
	b := &SampleSearchBuilder{
		rootOp:    op,
		currentOp: op,
	}
	b.Foo = &SampleSearchStringPropertyInfo{"Foo", b}

	return b
}

// SampleSearchBuilder builds Search API query.
type SampleSearchBuilder struct {
	rootOp    *smgutils.Op
	currentOp *smgutils.Op // for grouping
	opts      *search.SearchOptions
	query     string
	index     *search.Index
	Foo       *SampleSearchStringPropertyInfo
}

// And append new operant to query.
func (b *SampleSearchBuilder) And() *SampleSearchBuilder {
	b.currentOp.Children = append(b.currentOp.Children, &smgutils.Op{Type: smgutils.And})
	return b
}

// Or append new operant to query.
func (b *SampleSearchBuilder) Or() *SampleSearchBuilder {
	b.currentOp.Children = append(b.currentOp.Children, &smgutils.Op{Type: smgutils.Or})
	return b
}

// Group append new operant to query.
func (b *SampleSearchBuilder) Group(p func()) *SampleSearchBuilder {
	b.StartGroup()
	p()
	b.EndGroup()
	return b
}

// StartGroup append new operant to query.
func (b *SampleSearchBuilder) StartGroup() *SampleSearchBuilder {
	op := &smgutils.Op{Type: smgutils.Group, Parent: b.currentOp}
	b.currentOp.Children = append(b.currentOp.Children, op)
	b.currentOp = op
	return b
}

// EndGroup append new operant to query.
func (b *SampleSearchBuilder) EndGroup() *SampleSearchBuilder {
	b.currentOp = b.currentOp.Parent
	return b
}

// Put document to Index.
func (b *SampleSearchBuilder) Put(c context.Context, src *Sample) (string, error) {
	doc, err := src.Searchfy()
	if err != nil {
		return "", err
	}
	return b.PutDocument(c, doc)
}

// PutDocument to Index.
func (b *SampleSearchBuilder) PutDocument(c context.Context, src *SampleSearch) (string, error) {
	index, err := search.Open("Sample")
	if err != nil {
		return "", err
	}

	docID := ""
	if v, ok := interface{}(src).(smgutils.DocIDer); ok { // TODO can I shorten this cond expression?
		docID, err = v.DocID(c)
		if err != nil {
			return "", err
		}

	}

	log.Debugf(c, "id: %#v, payload: %#v", docID, src)

	docID, err = index.Put(c, docID, src)
	if err != nil {
		return "", err
	}

	return docID, nil
}

// Delete document from Index.
func (b *SampleSearchBuilder) Delete(c context.Context, src *Sample) error {
	doc, err := src.Searchfy()
	if err != nil {
		return err
	}
	return b.DeleteDocument(c, doc)
}

// DeleteDocument from Index.
func (b *SampleSearchBuilder) DeleteDocument(c context.Context, src *SampleSearch) error {
	if v, ok := interface{}(src).(smgutils.DocIDer); ok { // TODO can I shorten this cond expression?
		docID, err := v.DocID(c)
		if err != nil {
			return err
		}
		return b.DeleteByDocID(c, docID)
	}

	return errors.New("src is not implemented DocIDer interface")
}

// DeleteByDocID from Index.
func (b *SampleSearchBuilder) DeleteByDocID(c context.Context, docID string) error {
	index, err := search.Open("Sample")
	if err != nil {
		return err
	}

	return index.Delete(c, docID)
}

// Opts returns *SampleSearchOptions.
func (b *SampleSearchBuilder) Opts() *SampleSearchOptions {
	return &SampleSearchOptions{b: b}
}

// Search returns *SampleSearchIterator, It is result from Index.
func (b *SampleSearchBuilder) Search(c context.Context) (*SampleSearchIterator, error) {
	index, err := search.Open("Sample")
	if err != nil {
		return nil, err
	}
	b.index = index

	buffer := &bytes.Buffer{}
	err = b.rootOp.Query(buffer)
	if err != nil {
		return nil, err
	}
	b.query = buffer.String()
	log.Debugf(c, "query: '%s', opts: %#v", b.query, b.opts)
	iter := b.index.Search(c, b.query, b.opts)

	return &SampleSearchIterator{b, iter}, nil
}

// SampleSearchOptions construct *search.SearchOptions.
type SampleSearchOptions struct {
	b *SampleSearchBuilder
}

// Limit setup opts.
func (b *SampleSearchOptions) Limit(value int) *SampleSearchOptions {
	if b.b.opts == nil {
		b.b.opts = &search.SearchOptions{}
	}
	b.b.opts.Limit = value
	return b
}

// IDsOnly setup opts.
func (b *SampleSearchOptions) IDsOnly() *SampleSearchOptions {
	if b.b.opts == nil {
		b.b.opts = &search.SearchOptions{}
	}
	b.b.opts.IDsOnly = true
	return b
}

// Cursor setup opts.
func (b *SampleSearchOptions) Cursor(cursor search.Cursor) *SampleSearchOptions {
	b.b.opts.Cursor = cursor
	return b
}

// Offset setup opts.
func (b *SampleSearchOptions) Offset(value int) *SampleSearchOptions {
	b.b.opts.Offset = value
	return b
}

// SampleSearchIterator can access to search result.
type SampleSearchIterator struct {
	b    *SampleSearchBuilder
	iter *search.Iterator
}

// Next returns next document from iter.
func (b *SampleSearchIterator) Next(c context.Context) (string, *SampleSearch, error) {
	var s *SampleSearch
	if b.b.opts == nil || b.b.opts.IDsOnly != true {
		s = &SampleSearch{}
	}

	docID, err := b.iter.Next(s)
	if err != nil {
		return "", nil, err
	}

	return docID, s, err
}

// Cursor returns cursor of search.
func (b *SampleSearchIterator) Cursor() search.Cursor {
	return b.iter.Cursor()
}

// SampleSearchStringPropertyInfo hold property info.
type SampleSearchStringPropertyInfo struct {
	Name string
	b    *SampleSearchBuilder
}

// Match add query operand.
func (p *SampleSearchStringPropertyInfo) Match(value string) *SampleSearchBuilder {
	p.b.currentOp.Children = append(p.b.currentOp.Children, &smgutils.Op{FieldName: p.Name, Type: smgutils.Match, Value: value})
	return p.b
}

// Asc add query operand.
func (p *SampleSearchStringPropertyInfo) Asc() *SampleSearchBuilder {
	if p.b.opts == nil {
		p.b.opts = &search.SearchOptions{}
	}
	if p.b.opts.Sort == nil {
		p.b.opts.Sort = &search.SortOptions{}
	}
	p.b.opts.Sort.Expressions = append(p.b.opts.Sort.Expressions, search.SortExpression{
		Expr:    p.Name,
		Reverse: true,
	})

	return p.b
}

// Desc add query operand.
func (p *SampleSearchStringPropertyInfo) Desc() *SampleSearchBuilder {
	if p.b.opts == nil {
		p.b.opts = &search.SearchOptions{}
	}
	if p.b.opts.Sort == nil {
		p.b.opts.Sort = &search.SortOptions{}
	}
	p.b.opts.Sort.Expressions = append(p.b.opts.Sort.Expressions, search.SortExpression{
		Expr:    p.Name,
		Reverse: false,
	})

	return p.b
}

// SampleSearchNgramStringPropertyInfo hold property info.
type SampleSearchNgramStringPropertyInfo struct {
	SampleSearchStringPropertyInfo
}

// NgramMatch add query operand.
func (p *SampleSearchNgramStringPropertyInfo) NgramMatch(value string) *SampleSearchBuilder {
	p.b.currentOp.Children = append(p.b.currentOp.Children, &smgutils.Op{FieldName: p.Name, Type: smgutils.NgramMatch, Value: value})
	return p.b
}

// SampleSearchNumberPropertyInfo hold property info.
type SampleSearchNumberPropertyInfo struct {
	Name string
	b    *SampleSearchBuilder
}

// IntGreaterThanOrEqual add query operand.
func (p *SampleSearchNumberPropertyInfo) IntGreaterThanOrEqual(value int) *SampleSearchBuilder {
	p.b.currentOp.Children = append(p.b.currentOp.Children, &smgutils.Op{FieldName: p.Name, Type: smgutils.GtEq, Value: value})
	return p.b
}

// IntGreaterThan add query operand.
func (p *SampleSearchNumberPropertyInfo) IntGreaterThan(value int) *SampleSearchBuilder {
	p.b.currentOp.Children = append(p.b.currentOp.Children, &smgutils.Op{FieldName: p.Name, Type: smgutils.Gt, Value: value})
	return p.b
}

// IntLessThanOrEqual add query operand.
func (p *SampleSearchNumberPropertyInfo) IntLessThanOrEqual(value int) *SampleSearchBuilder {
	p.b.currentOp.Children = append(p.b.currentOp.Children, &smgutils.Op{FieldName: p.Name, Type: smgutils.LtEq, Value: value})
	return p.b
}

// IntLessThan add query operand.
func (p *SampleSearchNumberPropertyInfo) IntLessThan(value int) *SampleSearchBuilder {
	p.b.currentOp.Children = append(p.b.currentOp.Children, &smgutils.Op{FieldName: p.Name, Type: smgutils.Lt, Value: value})
	return p.b
}

// IntEqual add query operand.
func (p *SampleSearchNumberPropertyInfo) IntEqual(value int) *SampleSearchBuilder {
	p.b.currentOp.Children = append(p.b.currentOp.Children, &smgutils.Op{FieldName: p.Name, Type: smgutils.Eq, Value: value})
	return p.b
}

// Int64GreaterThanOrEqual add query operand.
func (p *SampleSearchNumberPropertyInfo) Int64GreaterThanOrEqual(value int64) *SampleSearchBuilder {
	p.b.currentOp.Children = append(p.b.currentOp.Children, &smgutils.Op{FieldName: p.Name, Type: smgutils.GtEq, Value: value})
	return p.b
}

// Int64GreaterThan add query operand.
func (p *SampleSearchNumberPropertyInfo) Int64GreaterThan(value int64) *SampleSearchBuilder {
	p.b.currentOp.Children = append(p.b.currentOp.Children, &smgutils.Op{FieldName: p.Name, Type: smgutils.Gt, Value: value})
	return p.b
}

// Int64LessThanOrEqual add query operand.
func (p *SampleSearchNumberPropertyInfo) Int64LessThanOrEqual(value int64) *SampleSearchBuilder {
	p.b.currentOp.Children = append(p.b.currentOp.Children, &smgutils.Op{FieldName: p.Name, Type: smgutils.LtEq, Value: value})
	return p.b
}

// Int64LessThan add query operand.
func (p *SampleSearchNumberPropertyInfo) Int64LessThan(value int64) *SampleSearchBuilder {
	p.b.currentOp.Children = append(p.b.currentOp.Children, &smgutils.Op{FieldName: p.Name, Type: smgutils.Lt, Value: value})
	return p.b
}

// Int64Equal add query operand.
func (p *SampleSearchNumberPropertyInfo) Int64Equal(value int64) *SampleSearchBuilder {
	p.b.currentOp.Children = append(p.b.currentOp.Children, &smgutils.Op{FieldName: p.Name, Type: smgutils.Eq, Value: value})
	return p.b
}

// Asc add query operand.
func (p *SampleSearchNumberPropertyInfo) Asc() *SampleSearchBuilder {
	if p.b.opts == nil {
		p.b.opts = &search.SearchOptions{}
	}
	if p.b.opts.Sort == nil {
		p.b.opts.Sort = &search.SortOptions{}
	}
	p.b.opts.Sort.Expressions = append(p.b.opts.Sort.Expressions, search.SortExpression{
		Expr:    p.Name,
		Reverse: true,
	})

	return p.b
}

// Desc add query operand.
func (p *SampleSearchNumberPropertyInfo) Desc() *SampleSearchBuilder {
	if p.b.opts == nil {
		p.b.opts = &search.SearchOptions{}
	}
	if p.b.opts.Sort == nil {
		p.b.opts.Sort = &search.SortOptions{}
	}
	p.b.opts.Sort.Expressions = append(p.b.opts.Sort.Expressions, search.SortExpression{
		Expr:    p.Name,
		Reverse: false,
	})

	return p.b
}

// SampleSearchBoolPropertyInfo hold property info.
type SampleSearchBoolPropertyInfo struct {
	Name string
	b    *SampleSearchBuilder
}

// Equal add query operand.
func (p *SampleSearchNumberPropertyInfo) Equal(value bool) *SampleSearchBuilder {
	p.b.currentOp.Children = append(p.b.currentOp.Children, &smgutils.Op{FieldName: p.Name, Type: smgutils.Eq, Value: value})
	return p.b
}

// SampleSearchTimePropertyInfo hold property info.
type SampleSearchTimePropertyInfo struct {
	Name string
	b    *SampleSearchBuilder
}

// query spec for time.Time.
// https://cloud.google.com/appengine/docs/go/search/query_strings#Go_Queries_on_date_fields
// It using date, not datetime.

// GreaterThanOrEqual add query operand.
func (p *SampleSearchTimePropertyInfo) GreaterThanOrEqual(value time.Time) *SampleSearchBuilder {
	p.b.currentOp.Children = append(p.b.currentOp.Children, &smgutils.Op{FieldName: p.Name, Type: smgutils.GtEq, Value: value.UTC()})
	return p.b
}

// GreaterThan add query operand.
func (p *SampleSearchTimePropertyInfo) GreaterThan(value time.Time) *SampleSearchBuilder {
	p.b.currentOp.Children = append(p.b.currentOp.Children, &smgutils.Op{FieldName: p.Name, Type: smgutils.Gt, Value: value.UTC()})
	return p.b
}

// LessThanOrEqual add query operand.
func (p *SampleSearchTimePropertyInfo) LessThanOrEqual(value time.Time) *SampleSearchBuilder {
	p.b.currentOp.Children = append(p.b.currentOp.Children, &smgutils.Op{FieldName: p.Name, Type: smgutils.LtEq, Value: value.UTC()})
	return p.b
}

// LessThan add query operand.
func (p *SampleSearchTimePropertyInfo) LessThan(value time.Time) *SampleSearchBuilder {
	p.b.currentOp.Children = append(p.b.currentOp.Children, &smgutils.Op{FieldName: p.Name, Type: smgutils.Lt, Value: value.UTC()})
	return p.b
}

// Equal add query operand.
func (p *SampleSearchTimePropertyInfo) Equal(value time.Time) *SampleSearchBuilder {
	p.b.currentOp.Children = append(p.b.currentOp.Children, &smgutils.Op{FieldName: p.Name, Type: smgutils.Eq, Value: value.UTC()})
	return p.b
}

// Asc add query operand.
func (p *SampleSearchTimePropertyInfo) Asc() *SampleSearchBuilder {
	if p.b.opts == nil {
		p.b.opts = &search.SearchOptions{}
	}
	if p.b.opts.Sort == nil {
		p.b.opts.Sort = &search.SortOptions{}
	}
	p.b.opts.Sort.Expressions = append(p.b.opts.Sort.Expressions, search.SortExpression{
		Expr:    p.Name,
		Reverse: true,
	})

	return p.b
}

// Desc add query operand.
func (p *SampleSearchTimePropertyInfo) Desc() *SampleSearchBuilder {
	if p.b.opts == nil {
		p.b.opts = &search.SearchOptions{}
	}
	if p.b.opts.Sort == nil {
		p.b.opts.Sort = &search.SortOptions{}
	}
	p.b.opts.Sort.Expressions = append(p.b.opts.Sort.Expressions, search.SortExpression{
		Expr:    p.Name,
		Reverse: false,
	})

	return p.b
}

// SampleSearchUnixTimePropertyInfo hold property info.
type SampleSearchUnixTimePropertyInfo struct {
	Name string
	b    *SampleSearchBuilder
}

// GreaterThanOrEqual add query operand.
func (p *SampleSearchUnixTimePropertyInfo) GreaterThanOrEqual(value time.Time) *SampleSearchBuilder {
	p.b.currentOp.Children = append(p.b.currentOp.Children, &smgutils.Op{FieldName: p.Name, Type: smgutils.GtEq, Value: value.UTC()})
	return p.b
}

// GreaterThan add query operand.
func (p *SampleSearchUnixTimePropertyInfo) GreaterThan(value time.Time) *SampleSearchBuilder {
	p.b.currentOp.Children = append(p.b.currentOp.Children, &smgutils.Op{FieldName: p.Name, Type: smgutils.Gt, Value: value.UTC()})
	return p.b
}

// LessThanOrEqual add query operand.
func (p *SampleSearchUnixTimePropertyInfo) LessThanOrEqual(value time.Time) *SampleSearchBuilder {
	p.b.currentOp.Children = append(p.b.currentOp.Children, &smgutils.Op{FieldName: p.Name, Type: smgutils.LtEq, Value: value.UTC()})
	return p.b
}

// LessThan add query operand.
func (p *SampleSearchUnixTimePropertyInfo) LessThan(value time.Time) *SampleSearchBuilder {
	p.b.currentOp.Children = append(p.b.currentOp.Children, &smgutils.Op{FieldName: p.Name, Type: smgutils.Lt, Value: value.UTC()})
	return p.b
}

// Equal add query operand.
func (p *SampleSearchUnixTimePropertyInfo) Equal(value time.Time) *SampleSearchBuilder {
	p.b.currentOp.Children = append(p.b.currentOp.Children, &smgutils.Op{FieldName: p.Name, Type: smgutils.Eq, Value: value.UTC()})
	return p.b
}

// Asc add query operand.
func (p *SampleSearchUnixTimePropertyInfo) Asc() *SampleSearchBuilder {
	if p.b.opts == nil {
		p.b.opts = &search.SearchOptions{}
	}
	if p.b.opts.Sort == nil {
		p.b.opts.Sort = &search.SortOptions{}
	}
	p.b.opts.Sort.Expressions = append(p.b.opts.Sort.Expressions, search.SortExpression{
		Expr:    p.Name,
		Reverse: true,
	})

	return p.b
}

// Desc add query operand.
func (p *SampleSearchUnixTimePropertyInfo) Desc() *SampleSearchBuilder {
	if p.b.opts == nil {
		p.b.opts = &search.SearchOptions{}
	}
	if p.b.opts.Sort == nil {
		p.b.opts.Sort = &search.SortOptions{}
	}
	p.b.opts.Sort.Expressions = append(p.b.opts.Sort.Expressions, search.SortExpression{
		Expr:    p.Name,
		Reverse: false,
	})

	return p.b
}

// UnixTimeGreaterThanOrEqual add query operand.
func (p *SampleSearchUnixTimePropertyInfo) UnixTimeGreaterThanOrEqual(value time.Time) *SampleSearchBuilder {
	p.b.currentOp.Children = append(p.b.currentOp.Children, &smgutils.Op{FieldName: p.Name + "UnixTime", Type: smgutils.GtEq, Value: smgutils.Unix(value)})
	return p.b
}

// UnixTimeGreaterThan add query operand.
func (p *SampleSearchUnixTimePropertyInfo) UnixTimeGreaterThan(value time.Time) *SampleSearchBuilder {
	p.b.currentOp.Children = append(p.b.currentOp.Children, &smgutils.Op{FieldName: p.Name + "UnixTime", Type: smgutils.Gt, Value: smgutils.Unix(value)})
	return p.b
}

// UnixTimeLessThanOrEqual add query operand.
func (p *SampleSearchUnixTimePropertyInfo) UnixTimeLessThanOrEqual(value time.Time) *SampleSearchBuilder {
	p.b.currentOp.Children = append(p.b.currentOp.Children, &smgutils.Op{FieldName: p.Name + "UnixTime", Type: smgutils.LtEq, Value: smgutils.Unix(value)})
	return p.b
}

// UnixTimeLessThan add query operand.
func (p *SampleSearchUnixTimePropertyInfo) UnixTimeLessThan(value time.Time) *SampleSearchBuilder {
	p.b.currentOp.Children = append(p.b.currentOp.Children, &smgutils.Op{FieldName: p.Name + "UnixTime", Type: smgutils.Lt, Value: smgutils.Unix(value)})
	return p.b
}

// UnixTimeEqual add query operand.
func (p *SampleSearchUnixTimePropertyInfo) UnixTimeEqual(value time.Time) *SampleSearchBuilder {
	p.b.currentOp.Children = append(p.b.currentOp.Children, &smgutils.Op{FieldName: p.Name + "UnixTime", Type: smgutils.Eq, Value: smgutils.Unix(value)})
	return p.b
}

// UnixTimeAsc add query operand.
func (p *SampleSearchUnixTimePropertyInfo) UnixTimeAsc() *SampleSearchBuilder {
	if p.b.opts == nil {
		p.b.opts = &search.SearchOptions{}
	}
	if p.b.opts.Sort == nil {
		p.b.opts.Sort = &search.SortOptions{}
	}
	p.b.opts.Sort.Expressions = append(p.b.opts.Sort.Expressions, search.SortExpression{
		Expr:    p.Name + "UnixTime",
		Reverse: true,
	})

	return p.b
}

// UnixTimeDesc add query operand.
func (p *SampleSearchUnixTimePropertyInfo) UnixTimeDesc() *SampleSearchBuilder {
	if p.b.opts == nil {
		p.b.opts = &search.SearchOptions{}
	}
	if p.b.opts.Sort == nil {
		p.b.opts.Sort = &search.SortOptions{}
	}
	p.b.opts.Sort.Expressions = append(p.b.opts.Sort.Expressions, search.SortExpression{
		Expr:    p.Name + "UnixTime",
		Reverse: false,
	})

	return p.b
}
