// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"
)

func NewCmdServe() *cobra.Command {
	type Options struct {
		Optint int    `validate:"min=0,max=10"`      // optintから変更し、validateタグを追記
		Optstr string `validate:"required,alphanum"` // optstrから変更し、validateタグを追記
	}

	var (
		o = &Options{}
	)

	cmd := &cobra.Command{
		Use:   "serve",
		Short: "A brief description of your command",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return validateParams(*o)
		},
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Printf("serve called: optint: %d, optstr: %s", o.Optint, o.Optstr)
		},
		SilenceErrors: true,
		SilenceUsage:  true,
	}
	cmd.Flags().IntVarP(&o.Optint, "int", "i", 0, "int option")
	cmd.Flags().StringVarP(&o.Optstr, "str", "s", "default", "string option")

	return cmd
}

func init() {
}
