// Copyright 2019 The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package page

import (
	"html/template"

	"github.com/gohugoio/hugo/lazy"
)

// LazyContentProvider initializes itself when read. Each method of the
// ContentProvider interface initializes a content provider and shares it
// with other methods.
//
// Used in cases where we cannot guarantee whether the content provider
// will be needed. Must create via NewLazyContentProvider.
type LazyContentProvider struct {
	init *lazy.Init
	cp   ContentProvider
}

// NewLazyContentProvider returns a LazyContentProvider initialized with
// function f. The resulting LazyContentProvider calls f in order to
// retrieve a ContentProvider
func NewLazyContentProvider(f func() (ContentProvider, error)) *LazyContentProvider {
	lcp := LazyContentProvider{
		init: lazy.New(),
		cp:   NopPage,
	}
	lcp.init.Add(func() (interface{}, error) {
		cp, err := f()
		if err != nil {
			return nil, err
		}
		lcp.cp = cp
		return nil, nil
	})
	return &lcp
}

func (lcp *LazyContentProvider) Init() ContentProvider {
	lcp.init.Do()
	return lcp.cp
}

func (lcp *LazyContentProvider) Reset() {
	lcp.init.Reset()
}

func (lcp *LazyContentProvider) Content() (interface{}, error) {
	lcp.init.Do()
	return lcp.cp.Content()
}

func (lcp *LazyContentProvider) Plain() string {
	lcp.init.Do()
	return lcp.cp.Plain()
}

func (lcp *LazyContentProvider) PlainWords() []string {
	lcp.init.Do()
	return lcp.cp.PlainWords()
}

func (lcp *LazyContentProvider) Summary() template.HTML {
	lcp.init.Do()
	return lcp.cp.Summary()
}

func (lcp *LazyContentProvider) Truncated() bool {
	lcp.init.Do()
	return lcp.cp.Truncated()
}

func (lcp *LazyContentProvider) FuzzyWordCount() int {
	lcp.init.Do()
	return lcp.cp.FuzzyWordCount()
}

func (lcp *LazyContentProvider) WordCount() int {
	lcp.init.Do()
	return lcp.cp.WordCount()
}

func (lcp *LazyContentProvider) ReadingTime() int {
	lcp.init.Do()
	return lcp.cp.ReadingTime()
}

func (lcp *LazyContentProvider) Len() int {
	lcp.init.Do()
	return lcp.cp.Len()
}