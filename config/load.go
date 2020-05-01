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
	"strings"

	"github.com/arnumina/dastum/failure"
	"github.com/arnumina/dastum/options"

	"github.com/arnumina/hepken/value"
)

// Loader AFAIRE
type Loader func(string, options.Options) (*value.Value, error)

var _loaders = make(map[string]Loader)

// AddLoader AFAIRE
func AddLoader(t string, fn Loader) {
	_loaders[t] = fn
}

func defaultLoader(t string, opts options.Options) (*value.Value, error) {
	switch t {
	case "empty":
		return value.Empty(), nil
	case "json":
		return loadJSONFile(opts)
	case "yaml":
		return loadYAMLFile(opts)
	default:
		return nil,
			failure.New(nil).
				Set("type", t).
				Msg("there is no configuration loader for this type") //////////////////////////////////////////////////
	}
}

func init() {
	AddLoader("empty", defaultLoader)
	AddLoader("json", defaultLoader)
	AddLoader("yaml", defaultLoader)
}

func parseCfgString(cs string) (string, options.Options, error) {
	if cs == "" {
		return "", nil,
			failure.New(nil).
				Msg("the configuration string is empty") ///////////////////////////////////////////////////////////////
	}

	opts := make(options.Options)

	ls := strings.Split(cs, ":")

	if len(ls) != 1 {
		if len(ls) != 2 {
			return "", nil,
				failure.New(nil).
					Set("string", cs).
					Msg("this configuration string is not valid") //////////////////////////////////////////////////////
		}

		for _, opt := range strings.Split(ls[1], ",") {
			kv := strings.Split(opt, "=")
			if len(kv) != 2 {
				return "", nil,
					failure.New(nil).
						Set("string", cs).
						Set("option", opt).
						Msg("this option of this configuration string is not valid") ///////////////////////////////////
			}

			opts[kv[0]] = kv[1]
		}
	}

	return ls[0], opts, nil
}

// Load AFAIRE
func Load(cs string) (*value.Value, error) {
	t, opts, err := parseCfgString(cs)
	if err != nil {
		return nil, err
	}

	loader, ok := _loaders[t]
	if !ok {
		return nil,
			failure.New(nil).
				Set("type", t).
				Msg("there is no configuration loader for this type") //////////////////////////////////////////////////
	}

	return loader(t, opts)
}

/*
######################################################################################################## @(°_°)@ #######
*/
