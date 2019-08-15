// Copyright 2014 The Gogs Authors. All rights reserved.
// Copyright 2018 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package geojson

import (
	"code.gitea.io/gitea/modules/markup"
)

func init() {
	markup.RegisterParser(Parser{})
}

// Parser implements markup.Parser
type Parser struct {
}

// Name implements markup.Parser
func (Parser) Name() string {
	return "geojson"
}

// Extensions implements markup.Parser
func (Parser) Extensions() []string {
	return []string{".geojson"}
}

// DisablePostprocess implements markup.Parser
func (Parser) DisablePostprocess() bool {
	// As we're dealing with json, we don't need html postprocessing
	// (it breaks json with '[[' link replacement)
	return true
}

// Render implements markup.Parser
func (Parser) Render(rawBytes []byte, urlPrefix string, metas map[string]string, isWiki bool) []byte {
	// output geojson as is, all logic runs client side.
	// this parser is just to make the templating engine aware of the geojson special case
	return rawBytes
}
