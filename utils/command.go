/*
 * Copyright (C) 2018 The "MysteriumNetwork/node" Authors.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package utils

import (
	"os/exec"
	"strings"
)

// SplitCommand parses command arguments from string and returns command with split arguments
func SplitCommand(command string, commandArguments string) *exec.Cmd {
	args := strings.Split(commandArguments, " ")
	var trimmedArgs []string
	for _, arg := range args {
		trimmedArgs = append(trimmedArgs, strings.TrimSpace(arg))
	}
	return exec.Command(command, trimmedArgs...)
}
