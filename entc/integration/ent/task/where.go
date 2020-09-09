// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package task

import (
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/entc/integration/ent/predicate"
	"github.com/facebook/ent/entc/integration/ent/schema"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Priority applies equality check predicate on the "priority" field. It's identical to PriorityEQ.
func Priority(v schema.Priority) predicate.Task {
	vc := int(v)
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPriority), vc))
	})
}

// PriorityEQ applies the EQ predicate on the "priority" field.
func PriorityEQ(v schema.Priority) predicate.Task {
	vc := int(v)
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPriority), vc))
	})
}

// PriorityNEQ applies the NEQ predicate on the "priority" field.
func PriorityNEQ(v schema.Priority) predicate.Task {
	vc := int(v)
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPriority), vc))
	})
}

// PriorityIn applies the In predicate on the "priority" field.
func PriorityIn(vs ...schema.Priority) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = int(vs[i])
	}
	return predicate.Task(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPriority), v...))
	})
}

// PriorityNotIn applies the NotIn predicate on the "priority" field.
func PriorityNotIn(vs ...schema.Priority) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = int(vs[i])
	}
	return predicate.Task(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPriority), v...))
	})
}

// PriorityGT applies the GT predicate on the "priority" field.
func PriorityGT(v schema.Priority) predicate.Task {
	vc := int(v)
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPriority), vc))
	})
}

// PriorityGTE applies the GTE predicate on the "priority" field.
func PriorityGTE(v schema.Priority) predicate.Task {
	vc := int(v)
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPriority), vc))
	})
}

// PriorityLT applies the LT predicate on the "priority" field.
func PriorityLT(v schema.Priority) predicate.Task {
	vc := int(v)
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPriority), vc))
	})
}

// PriorityLTE applies the LTE predicate on the "priority" field.
func PriorityLTE(v schema.Priority) predicate.Task {
	vc := int(v)
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPriority), vc))
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Task) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Task) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Task) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		p(s.Not())
	})
}