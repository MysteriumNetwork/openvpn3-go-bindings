/*
 * Copyright (C) 2017 The "MysteriumNetwork/node" Authors.
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

package openvpn

import (
	"math/rand"
	"os"
	"time"
)

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()),
)

const finenameCharset = "1234567890abcdefghijklmnopqrstuvwxyz"

func tempFilename(directory, filePrefix, fileExtension string) (filePath string) {
	for i := 0; i < 10000; i++ {
		filePath = directory + "/" + filePrefix + randomString(10, finenameCharset) + fileExtension
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			break
		}
	}
	return
}

func randomString(size int, charset string) string {
	filename := make([]byte, size)
	for i := range filename {
		charsetIndex := seededRand.Intn(len(charset))
		filename[i] = charset[charsetIndex]
	}

	return string(filename)
}
