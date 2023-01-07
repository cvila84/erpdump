package table

type headers struct {
	parent      *headers
	label       string
	elements    map[string]*headers
	defaultSort Sort
	actualSort  Sort
}

func newRootHeaders(defaultSort Sort) *headers {
	return &headers{
		parent:      nil,
		label:       "",
		elements:    nil,
		defaultSort: defaultSort,
		actualSort:  defaultSort,
	}
}

func newChild(parent *headers, label string) *headers {
	var childLabel string
	if len(parent.label) > 0 {
		childLabel = parent.label + "/" + label
	} else {
		childLabel = label
	}
	return &headers{
		parent:      parent,
		label:       childLabel,
		elements:    nil,
		defaultSort: parent.defaultSort,
		actualSort:  nil,
	}
}

func (h *headers) sort(sort Sort) *headers {
	h.actualSort = sort
	return h
}

func (h *headers) walk(label string) *headers {
	if h.elements == nil {
		h.elements = make(map[string]*headers)
	}
	re, ok := h.elements[label]
	if !ok {
		re = newChild(h, label)
		h.elements[label] = re
	}
	return re
}

func (h *headers) labels(recursive bool, self bool) []string {
	labels := make([]string, 0)
	if h.elements != nil {
		keys := make([]string, 0, len(h.elements))
		for k := range h.elements {
			keys = append(keys, k)
		}
		if h.actualSort != nil {
			keys = h.actualSort(keys)
		} else if h.defaultSort != nil {
			keys = h.defaultSort(keys)
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
