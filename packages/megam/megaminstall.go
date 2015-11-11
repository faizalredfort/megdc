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
package megam

import (
	"github.com/megamsys/libgo/cmd"
	"github.com/megamsys/megdc/handler"
	"launchpad.net/gnuflag"
)

type MegamInstall struct {
	Fs        *gnuflag.FlagSet
	All       bool
	Nilavu    bool
	Gateway   bool
	Megamd    bool
	Snowflake bool
}

func (c *MegamInstall) Info() *cmd.Info {
	return &cmd.Info{
		Name:  "megaminstall",
		Usage: "megaminstall [--nilavu/-n] [--gateway/-g] [--snowflake/-s]",
		Desc: `Install megam (app orchestrator) on the local machine. For megdc, available install plaform is ubuntu.
We are working to support centos.
In order to install individual packages use the following options.

The [[--nilavu]] parameter defines megam cockpit ui to install.
This code name is nilavu packaged as megamnilavu.

The [[--gateway]] parameter defines megam gateway apiserver to install.
This code name is gateway packaged as megamgateway.

The [[--snowflake]] parameter defines megam uidserver to install.
This code name is snowflake packaged as megamsnowflake.

The [[--megamd]] parameter defines megam omni scheduler to install.
This code name is megamd packaged as megammegamd.

For more information read http://docs.megam.io.`,
		MinArgs: 0,
	}
}

func (c *MegamInstall) Run(context *cmd.Context) error {
	handler.SunSpin(cmd.Colorfy(handler.Logo, "green", "", "bold"), "", "install")
	w := handler.NewWrap(c)
	c.chooseAll(w)
	if h, err := handler.NewHandler(w); err != nil {
		return err
	} else if err := h.Run(); err != nil {
		return err
	}
	return nil
}

func (c *MegamInstall) Flags() *gnuflag.FlagSet {
	if c.Fs == nil {
		c.Fs = gnuflag.NewFlagSet("", gnuflag.ExitOnError)
		nilMsg := "Install megam cockpit ui"
		c.Fs.BoolVar(&c.Nilavu, "nilavu", false, nilMsg)
		c.Fs.BoolVar(&c.Nilavu, "n", false, nilMsg)
		gwyMsg := "Install megam gateway apiserver"
		c.Fs.BoolVar(&c.Gateway, "gateway", false, gwyMsg)
		c.Fs.BoolVar(&c.Gateway, "g", false, gwyMsg)
		megdMsg := "Install megam omni scheduler"
		c.Fs.BoolVar(&c.Megamd, "megamd", false, megdMsg)
		c.Fs.BoolVar(&c.Megamd, "d", false, megdMsg)
		snoMsg := "Install megam uidserver"
		c.Fs.BoolVar(&c.Snowflake, "snowflake", false, snoMsg)
		c.Fs.BoolVar(&c.Snowflake, "s", false, snoMsg)
	}
	return c.Fs
}

func (c *MegamInstall) chooseAll(w *handler.WrappedParms) {
	DEFAULT_PACKAGES := []string{"megamsnowflake",
		"megamnilavu", "megamgateway", "megamd"}

	if w.Empty() {
		for i := range DEFAULT_PACKAGES {
			w.AddPackage(DEFAULT_PACKAGES[i])
		}
	}
}
