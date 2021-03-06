/*
 * Copyright 2017 Google Inc. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cli

import (
	"context"
	"flag"
	"strings"

	xpb "kythe.io/kythe/proto/xref_proto"
)

type docsCommand struct {
	nodeFilters     string
	includeChildren bool
}

func (docsCommand) Name() string     { return "docs" }
func (docsCommand) Synopsis() string { return "display documentation for a node" }
func (docsCommand) Usage() string    { return "" }
func (c *docsCommand) SetFlags(flag *flag.FlagSet) {
	flag.StringVar(&c.nodeFilters, "filters", "", "Comma-separated list of node fact filters (default returns all)")
	flag.BoolVar(&c.includeChildren, "include_children", false, "Include documentation for children of the given node")
}
func (c docsCommand) Run(ctx context.Context, flag *flag.FlagSet, api API) error {
	req := &xpb.DocumentationRequest{
		Ticket:          flag.Args(),
		IncludeChildren: c.includeChildren,
	}
	if c.nodeFilters != "" {
		req.Filter = strings.Split(c.nodeFilters, ",")
	}
	LogRequest(req)
	reply, err := api.XRefService.Documentation(ctx, req)
	if err != nil {
		return err
	}
	return c.displayDocumentation(reply)
}

func (c docsCommand) displayDocumentation(reply *xpb.DocumentationReply) error {
	// TODO(zarko): Emit formatted data for -json=false.
	return PrintJSONMessage(reply)
}
