/*
** Copyright [2013-2015] [Megam Systems]
**
** Licensed under the Apache License, Version 2.0 (the "License");
** you may not use this file except in compliance with the License.
** You may obtain a copy of the License at
**
** http://www.apache.org/licenses/LICENSE-2.0
**
** Unless required by applicable law or agreed to in writing, software
** distributed under the License is distributed on an "AS IS" BASIS,
** WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
** See the License for the specific language governing permissions and
** limitations under the License.
*/
package mesos

import (
	"github.com/megamsys/libgo/cmd"
	"gopkg.in/check.v1"
)

type S struct{}

var _ = check.Suite(&S{})

func (s *S) TestMesosmasterStartInfo(c *check.C) {
	desc := `starts megdc.

If you use the '--quiet' flag megdc doesn't print the logs.

`

	expected := &cmd.Info{
		Name:    "mesosmasterinstall",
		Usage:   `mesosmasterinstall [--masterip]...`,
		Desc:    desc,
		MinArgs: 0,
	}
	command := MesosMasterInstall{}
	c.Assert(command.Info(), check.DeepEquals, expected)
}
