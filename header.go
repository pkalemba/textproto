// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package textproto2

// A MIMEHeader represents a MIME-style header mapping
// keys to sets of values.
type MIMEHeader []string

// Add adds the key, value pair to the header.
// It appends to any existing values associated with key.
func (h MIMEHeader) Add(value string) {
	h = append(h, value)
}

// Set sets the header entries associated with key to
// the single element value. It replaces any existing
// values associated with key.
//func (h MIMEHeader) Set(key, value string) {
//for i := range h {
//if h[i].Key == key {
//tmp
//tmp := make(map[string]string)
//tmp[key] = value
//h[i] = tmp
//}
//}
//}

//// Get gets the first value associated with the given key.
//// It is case insensitive; CanonicalMIMEHeaderKey is used
//// to canonicalize the provided key.
//// If there are no values associated with the key, Get returns "".
//// To access multiple values of a key, or to use non-canonical keys,
//// access the map directly.
//func (h MIMEHeader) Get(key string) string {
//if h == nil {
//return ""
//}
//for i := range h {
//if h[i].Key == key {
//v = h[i].Value
//}
//return v
//}
//}

//// Del deletes the values associated with key.
//func (h MIMEHeader) Del(key string) {
//delete(h, CanonicalMIMEHeaderKey(key))
//}
