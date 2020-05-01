/*
#######
##         __            __
##        / /  ___ ___  / /_____ ___
##       / _ \/ -_) _ \/  '_/ -_) _ \
##      /_//_/\__/ .__/_/\_\\__/_//_/
##              /_/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package config

import (
	"io/ioutil"

	"github.com/arnumina/dastum/options"

	"github.com/arnumina/hepken/value"
)

func loadYAMLFile(opts options.Options) (*value.Value, error) {
	filename, err := opts.String("file")
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return value.FromYAML(content)
}

/*
######################################################################################################## @(°_°)@ #######
*/
