package table

import (
	"sort"
	"strings"
)

type Sorting func(elements []string) []string

var AlphaSorting Sorting = func(elements []string) []string {
	sort.Strings(elements)
	return elements
}

type Headers struct {
	parent   *Headers
	label    string
	elements map[string]*Headers
	sorting  Sorting
}

func NewRootHeaders() *Headers {
	return &Headers{
		parent:   nil,
		label:    "",
		elements: nil,
		sorting:  nil,
	}
}

func newChild(parent *Headers, label string, sorting Sorting) *Headers {
	var childLabel string
	if len(parent.label) > 0 {
		childLabel = parent.label + "/" + label
	} else {
		childLabel = label
	}
	parent.sorting = sorting
	return &Headers{
		parent:   parent,
		label:    childLabel,
		elements: nil,
		sorting:  nil,
	}
}

func (h *Headers) sortedWalk(label string, sorting Sorting) *Headers {
	if h.elements == nil {
		h.elements = make(map[string]*Headers)
	}
	re, ok := h.elements[label]
	if !ok {
		re = newChild(h, label, sorting)
		h.elements[label] = re
	}
	return re
}

func (h *Headers) walk(label string) *Headers {
	return h.sortedWalk(label, AlphaSorting)
}

func parentLabel(label string) string {
	if label == "" {
		return ""
	}
	idx := strings.LastIndex(label, "/")
	if idx <= 0 {
		return ""
	}
	return label[0:idx]
}

func (h *Headers) labels(recursive bool, self bool) []string {
	labels := make([]string, 0)
	if h.elements != nil {
		keys := make([]string, 0, len(h.elements))
		for k := range h.elements {
			keys = append(keys, k)
		}
		if h.sorting != nil {
			keys = h.sorting(keys)
		}
		for _, k := range keys {
			labels = append(labels, h.elements[k].label)
			if recursive {
				subLabels := h.elements[k].labels(recursive, false)
				if len(subLabels) > 0 {
					labels = append(labels, subLabels...)
				}
			}
		}
	}
	if self {
		labels = append(labels, h.label)
	}
	return labels
}
