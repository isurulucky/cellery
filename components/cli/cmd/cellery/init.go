/*
 * Copyright (c) 2018 WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
 *
 * WSO2 Inc. licenses this file to you under the Apache License,
 * Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package main

import (
	"github.com/spf13/cobra"

	"cellery.io/cellery/components/cli/cli"
	"cellery.io/cellery/components/cli/pkg/commands/project"
	"cellery.io/cellery/components/cli/pkg/util"
)

func newInitCommand(cli cli.Cli) *cobra.Command {
	var projectName = ""
	var isBallerinaProject bool
	cmd := &cobra.Command{
		Use:   "init [PROJECT_NAME]",
		Short: "Initialize a cell project",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				projectName = args[0]
			}
			if err := project.RunInit(cli, projectName, isBallerinaProject); err != nil {
				util.ExitWithErrorMessage("Cellery init command failed", err)
			}
		},
		Example: "  cellery init [PROJECT_NAME]\n" +
			"  cellery init test [PROJECT_NAME --project\n" +
			"  cellery init test [PROJECT_NAME -p",
	}
	cmd.Flags().BoolVarP(&isBallerinaProject, "project", "p", false,
		"Create a Ballerina project")
	return cmd
}
