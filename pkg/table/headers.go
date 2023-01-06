package table

type Headers struct {
	parent   *Headers
	label    string
	elements map[string]*Headers
	sorting  *Sorting
}

func NewChildHeaders(parent *Headers, label string) *Headers {
	return &Headers{
		parent:   parent,
		label:    label,
		elements: nil,
		sorting:  parent.sorting,
	}
}

func NewRootHeaders(sorting Sorting) *Headers {
	return &Headers{
		parent:   nil,
		label:    "",
		elements: nil,
		sorting:  &sorting,
	}
}

func (h *Headers) walk(label string) *Headers {
	if h.elements == nil {
		h.elements = make(map[string]*Headers)
	}
	re, ok := h.elements[label]
	if !ok {
		re = NewChildHeaders(h, label)
		h.elements[label] = re
	}
	return re
}

func (h *Headers) string() string {
	if h.parent == nil {
		return ""
	}
	parentLabel := h.parent.string()
	if len(parentLabel) > 0 {
		return parentLabel + "/" + h.label
	} else {
		return h.label
	}
}

func (h *Headers) labels() []string {
	labels := make([]string, 0)
	if h.elements != nil {
		keys := make([]string, 0, len(h.elements))
		for k := range h.elements {
			keys = append(keys, k)
		}
		for _, k := range (*h.sorting)(keys) {
			labels = append(labels, h.elements[k].string())
		}
	}
	return labels
}
