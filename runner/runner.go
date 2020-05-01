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

package runner

import (
	"strconv"
	"time"

	"github.com/arnumina/dastum"
	"github.com/arnumina/dastum/logger"

	"github.com/arnumina/hepken/value"
)

// Runner AFAIRE
type Runner struct {
	id        string
	name      string
	version   string
	builtAt   time.Time
	startedAt time.Time
	config    *value.Value
	logger    *logger.Logger
}

// New AFAIRE
func New(name, version, builtAt string) *Runner {
	ts, err := strconv.ParseInt(builtAt, 0, 64)
	if err != nil {
		ts = 0
	}

	return &Runner{
		id:        dastum.NewUUID(),
		name:      name,
		version:   version,
		builtAt:   time.Unix(ts, 0),
		startedAt: time.Now(),
	}
}

// ID AFAIRE
func (r *Runner) ID() string {
	return r.id
}

// Name AFAIRE
func (r *Runner) Name() string {
	return r.name
}

// Version AFAIRE
func (r *Runner) Version() string {
	return r.version
}

// BuiltAt AFAIRE
func (r *Runner) BuiltAt() time.Time {
	return r.builtAt
}

// StartedAt AFAIRE
func (r *Runner) StartedAt() time.Time {
	return r.startedAt
}

// SetConfig AFAIRE
func (r *Runner) SetConfig(config *value.Value) {
	r.config = config
}

// Config AFAIRE
func (r *Runner) Config() *value.Value {
	return r.config
}

// SetLogger AFAIRE
func (r *Runner) SetLogger(logger *logger.Logger) {
	r.logger = logger
}

// Logger AFAIRE
func (r *Runner) Logger() *logger.Logger {
	return r.logger
}

/*
######################################################################################################## @(°_°)@ #######
*/
