package transfer

import (
	"errors"
	"github.com/viant/toolbox/url"
	"strings"
)



//Rule represents transfer rule
type Rule struct {
	Matcher *Matcher
	Compress bool `description:"flag to compress asset before sending over wire and to decompress (this option is only supported on scp or file scheme)"` //flag to compress asset before sending over wirte and to decompress (this option is only supported on scp or file proto)
	Substitution
	Source  *url.Resource `required:"true" description:"source asset or directory"`
	Dest    *url.Resource `required:"true" description:"destination asset or directory"`
}


//New creates a new transfer
func New(source, dest *url.Resource, compress, expand bool, replace map[string]string) *Rule {
	return &Rule{
		Source:   source,
		Dest:     dest,
		Compress: compress,
		Substitution: Substitution{
			Expand:  expand,
			Replace: replace,
		},
	}
}

//Init initialises transfer
func (r *Rule) Init() error {
	if r.Source != nil {
		if !strings.HasPrefix(r.Source.URL, "$") {
			if err := r.Source.Init();err != nil {
				return err
			}
		}
	}
	if r.Dest != nil {
		if !strings.HasPrefix(r.Dest.URL, "$") {
			if err := r.Dest.Init();err != nil {
				return err
			}
		}
	}
	return nil
}

//Validate checks if request is valid
func (r *Rule) Validate() error {
	if r.Source == nil {
		return errors.New("source was empty")
	}
	if r.Source.URL == "" {
		return errors.New("source.URL was empty")
	}
	if r.Dest == nil {
		return errors.New("dest was empty")
	}
	if r.Dest.URL == "" {
		return errors.New("dest.URL was empty")
	}
	return nil
}


