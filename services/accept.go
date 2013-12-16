package services

import (
	"strings"
	"strconv"
)

// AcceptEntry represents a single entry within an Accept header.  It
// includes both the ContentType and a Quality, parsed from the
// ContentType's parameters.
type AcceptEntry struct {
	ContentType *ContentType

	// Quality is the parsed q value from a ContentType's parameters.
	Quality float32

	// Internal use only, for measuring how "specific" the entry is.
	specificityCount int
}

// NewAcceptEntry sets up an AcceptEntry with the proper default
// quality.
func NewAcceptEntry() *AcceptEntry {
	entry := &AcceptEntry{
		Quality: 1.0,
	}
	return entry
}

func ParseAcceptEntry(accept string) (*AcceptEntry, error) {

	entry := NewAcceptEntry()
	var typeErr error
	entry.ContentType, typeErr = ParseContentType(accept)
	if typeErr != nil {
		return nil, typeErr
	}

	if qualityString, ok := entry.ContentType.Parameters["q"]; ok {
		quality, err := strconv.ParseFloat(qualityString, 32)
		if err != nil {
			return nil, err
		}
		entry.Quality = float32(quality)
	}

	// Parameters are more specific.  Wildcards in the MimeType
	// are less specific.  I can't find anything detailing whether
	// one or the other is more important, so I'm counting them as
	// equals.
	entry.specificityCount += len(entry.ContentType.Parameters)
	entry.specificityCount -= strings.Count(entry.ContentType.MimeType, "*")

	return entry, nil
}

// CompareTo compares two AcceptEntries and returns an integer
// representing which of the two entries is more highly preferred.
// Negative return values mean that the passed in entry is preferred,
// positive values mean that the target entry is preferred, and zero
// values mean that there is no preference.
func (entry *AcceptEntry) CompareTo(otherEntry *AcceptEntry) int {
	if entry.Quality > otherEntry.Quality {
		return 1
	}
	if entry.Quality < otherEntry.Quality {
		return -1
	}
	return entry.specificityCount - otherEntry.specificityCount
}

// AcceptTree is a binary tree that handles Accept header entries.
// The left-most node will always be the most highly preferred entry,
// and preference levels will decrease as you move right in the tree.
type AcceptTree struct {
	Value *AcceptEntry
	Left *AcceptTree
	Right *AcceptTree
}

// Add adds an AcceptEntry to the AcceptTree, putting it in proper
// order of preference.
func (tree *AcceptTree) Add(next *AcceptEntry) {
	if tree.Value == nil {
		tree.Value = next
	} else if next.CompareTo(tree.Value) > 0 {
		if tree.Left == nil {
			tree.Left = &AcceptTree{Value: next}
		} else {
			tree.Left.Add(next)
		}
	} else {
		if tree.Right == nil {
			tree.Right = &AcceptTree{Value: next}
		} else {
			tree.Right.Add(next)
		}
	}
}

// Flatten returns the AcceptTree's values in proper order of
// preference as a []*AcceptEntry value.
func (tree *AcceptTree) Flatten() (entries []*AcceptEntry) {
	if tree.Value != nil {
		if tree.Left != nil {
			entries = append(entries, tree.Left.Flatten()...)
		}
		entries = append(entries, tree.Value)
		if tree.Right != nil {
			entries = append(entries, tree.Right.Flatten()...)
		}
	}
	return
}

// OrderAcceptHeader reads an Accept header and pulls out the various
// MIME types in the order of preference, returning the types in that
// order.
//
// The HTTP spec for the Accept header states that multiple MIME types
// can be specified in the Accept header, and that preferred MIME
// types are chosen based on the following criteria:
//
// 1. The q variable for a MIME type in the Accept header defines a
// 'quality', and higher qualities should be chosen over lower
// qualities.  The default quality is 1.0 for any type that doesn't
// state the quality explicitly.
// 2. More specific MIME types should be chosen over less specific
// MIME types, excepting the presence of a q parameter that counters
// this guideline.  For example, assuming equal qualities,
// application/xml should trump application/*.
// 3. Barring the previous two guidelines, MIME types should be chosen
// based on the order that they appear in the Accept header.
//
// For more information, see
// http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html
func OrderAcceptHeader(accept string) ([]*AcceptEntry, error) {
	acceptTree := new(AcceptTree)
	for _, rawEntry := range strings.Split(accept, ",") {
		rawEntry = strings.TrimSpace(rawEntry)
		if rawEntry != "" {
			entry, err := ParseAcceptEntry(rawEntry)
			if err != nil {
				return nil, err
			}
			acceptTree.Add(entry)
		}
	}

	return acceptTree.Flatten(), nil
}
