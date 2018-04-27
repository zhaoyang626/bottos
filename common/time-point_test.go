// Copyright 2017~2022 The Bottos Authors
// This file is part of the Bottos Chain library.
// Created by Rocket Core Team of Bottos.

//This program is free software: you can distribute it and/or modify
//it under the terms of the GNU General Public License as published by
//the Free Software Foundation, either version 3 of the License, or
//(at your option) any later version.

//This program is distributed in the hope that it will be useful,
//but WITHOUT ANY WARRANTY; without even the implied warranty of
//MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//GNU General Public License for more details.

//You should have received a copy of the GNU General Public License
// along with bottos.  If not, see <http://www.gnu.org/licenses/>.

/*
 * file description:  provide a interface such as time to seconds and epoch time etc.
 * @Author: May Luo
 * @Date:   2017-12-01
 * @Last Modified by:
 * @Last Modified time:
 */
package common

import (
	"fmt"
	"testing"
	"time"
)

func Test_NowToSeconds(t *testing.T) {
	sec := NowToSeconds()
	fmt.Println(sec)
}
func Test_NowToMicroseconds(t *testing.T) {
	now := time.Now()
	fmt.Println(now)
	sec := TimeToMicroseconds(now)
	fmt.Println(sec)
}
