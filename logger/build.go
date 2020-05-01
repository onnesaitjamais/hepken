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

package logger

import (
	"github.com/arnumina/dastum/failure"
	_logger "github.com/arnumina/dastum/logger"

	"github.com/arnumina/hepken/value"
)

const (
	_defaultLevel                = "trace"
	_defaultFormatter            = "default"
	_defaultOutput               = "stdout"
	_defaultOutputFileName       = "/tmp/hepken.log"
	_defaultOutputSyslogFacility = "local0"
)

// FmtBuilder AFAIRE
type FmtBuilder func(string, *value.Value) (_logger.Formatter, error)

// OutBuilder AFAIRE
type OutBuilder func(string, *value.Value) (_logger.Output, error)

var (
	_fmtBuilders = make(map[string]FmtBuilder)
	_outBuilders = make(map[string]OutBuilder)
)

// AddFmtBuilder AFAIRE
func AddFmtBuilder(t string, builder FmtBuilder) {
	_fmtBuilders[t] = builder
}

// AddOutBuilder AFAIRE
func AddOutBuilder(t string, builder OutBuilder) {
	_outBuilders[t] = builder
}

func defaultFmtBuilder(t string, _ *value.Value) (_logger.Formatter, error) {
	switch t {
	case "default":
		return _logger.NewDefaultFormatter(), nil
	default:
		return nil, failure.New(nil).
			Set("type", t).
			Msg("this type of logger formatter does not exist") ////////////////////////////////////////////////////////
	}
}

func buildOutputFile(cfg *value.Value) (_logger.Output, error) {
	fn, err := cfg.DString(_defaultOutputFileName, "file", "name")
	if err != nil {
		return nil, err
	}

	return _logger.NewFileOutput(fn)
}

func buildOutputSyslog(cfg *value.Value) (_logger.Output, error) {
	sf, err := cfg.DString(_defaultOutputSyslogFacility, "syslog", "facility")
	if err != nil {
		return nil, err
	}

	return _logger.NewSyslogOutput(sf)
}

func defaultOutBuilder(t string, cfg *value.Value) (_logger.Output, error) {
	switch t {
	case "file":
		return buildOutputFile(cfg)
	case "stderr":
		return _logger.NewStderrOutput(), nil
	case "stdout":
		return _logger.NewStdoutOutput(), nil
	case "syslog":
		return buildOutputSyslog(cfg)
	default:
		return nil, failure.New(nil).
			Set("type", t).
			Msg("this type of logger output does not exist") ///////////////////////////////////////////////////////////
	}
}

func init() {
	AddFmtBuilder("default", defaultFmtBuilder)
	AddOutBuilder("file", defaultOutBuilder)
	AddOutBuilder("stderr", defaultOutBuilder)
	AddOutBuilder("stdout", defaultOutBuilder)
	AddOutBuilder("syslog", defaultOutBuilder)
}

func formatter(cfg *value.Value) (_logger.Formatter, error) {
	t, err := cfg.DString(_defaultFormatter, "formatter")
	if err != nil {
		return nil, err
	}

	builder, ok := _fmtBuilders[t]
	if !ok {
		return nil,
			failure.New(nil).
				Set("type", t).
				Msg("there is no builder for this type of logger formatter") ///////////////////////////////////////////
	}

	fmt, err := builder(t, cfg)
	if err != nil {
		return nil, err
	}

	return fmt, nil
}

func output(cfg *value.Value) (_logger.Output, error) {
	t, err := cfg.DString(_defaultOutput, "output")
	if err != nil {
		return nil, err
	}

	builder, ok := _outBuilders[t]
	if !ok {
		return nil,
			failure.New(nil).
				Set("type", t).
				Msg("there is no builder for this type of logger output") //////////////////////////////////////////////
	}

	out, err := builder(t, cfg)
	if err != nil {
		return nil, err
	}

	return out, nil
}

// Build AFAIRE
func Build(runner string, cfg *value.Value) (*_logger.Logger, error) {
	level, err := cfg.DString(_defaultLevel, "level")
	if err != nil {
		return nil, err
	}

	fmt, err := formatter(cfg)
	if err != nil {
		return nil, err
	}

	out, err := output(cfg)
	if err != nil {
		return nil, err
	}

	return _logger.New(level, runner, fmt, out), nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
