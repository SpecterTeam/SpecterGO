/**
 *     SpecterGO  Copyright (C) 2018  SpecterTeam
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
	"path/filepath"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"encoding/json"
	"errors"
	"os"
)

const (
	TypeJson = iota // .json
	TypeYaml // .yml & .yaml
)


type(
	Content map[string]string

	Config struct {
		file       string
		config     Content
		configType int
	}
)

func NewConfig(file string, configType int) Config  {
	c := Config{}
	c.SetConfigType(configType)
	c.SetFile(file)
	if FileExists(file) {
		if ext := filepath.Ext(file); ExtMatchType(ext, configType) {
			c.SetConfig(c.Unmarshal())
		} else {
			err := errors.New("Ext of " + file + " doesn't match the configType!")
			HandleError(err)
		}
	} else {
		os.Create(file)
	}
	return c
}

func ExtMatchType(ext string, configType int) bool {
	switch configType {
	case TypeJson:
		if ext == "json" {
			return true
		} else {
			return false
		}
	case TypeYaml:
		if ext == "yml" || ext == "yaml" {
			return true
		} else {
			return false
		}
	}
	return false
}

func (c *Config) Marshal() ([]byte, error) {
	var b []byte

	switch c.ConfigType() {
	case TypeYaml:
		b,_ = yaml.Marshal(c.Config())
	case TypeJson:
		b,_ = json.Marshal(c.Config())
	default:
		err := errors.New("couldn't marshal the Config, Perhaps because the type of the config isn't set")
		return b, err
	}

	return b, nil
}

func (c *Config) Unmarshal() Content {
	var r Content
	switch c.ConfigType() {
	case TypeYaml:
		bts,_ := ioutil.ReadFile(c.File())
		yaml.Unmarshal(bts,&r)
	case TypeJson:
		bts,_ := ioutil.ReadFile(c.File())
		json.Unmarshal(bts,&r)
	}
	return r
}

func (c *Config) Save(goroutines bool) {
	if goroutines == true {
		go func() {
			bts,err := c.Marshal()
			if err != nil {
				HandleError(err)
			} else {
				ioutil.WriteFile(c.File(), bts, 0644)
			}
		}()
	} else {
		bts, err := c.Marshal()
		if err != nil {
			HandleError(err)
		} else {
			ioutil.WriteFile(c.File(), bts, 0644)
		}
	}
}

func (c *Config) ConfigType() int {
	return c.configType
}

func (c *Config) SetConfigType(configType int) {
	c.configType = configType
}

func (c *Config) Config() Content {
	return c.config
}

func (c *Config) SetConfig(config Content) {
	c.config = config
}

func (c *Config) File() string {
	return c.file
}

func (c *Config) SetFile(file string) {
	c.file = file
}