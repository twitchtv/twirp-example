// Copyright 2021 Twitch Interactive, Inc.  All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may not
// use this file except in compliance with the License. A copy of the License is
// located at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// or in the "license" file accompanying this file. This file is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package server

// New returns a new Haberdasher which returns random Hats of the requested
// size.
func NewHaberdasherServer() *haberdasherServer {
	return new(haberdasherServer)
}

// randomHaberdasher is our implementation of the generated
// rpc/haberdasher.Haberdasher interface. This is where the real "business logic" lives.
type haberdasherServer struct{}
