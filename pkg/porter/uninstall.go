package porter

import (
	"fmt"

	"github.com/pkg/errors"
)

// UninstallOptions that may be specified when uninstalling a bundle.
// Porter handles defaulting any missing values.
type UninstallOptions struct {
	BundleLifecycleOpts
}

// UninstallBundle accepts a set of pre-validated UninstallOptions and uses
// them to uninstall a bundle.
func (p *Porter) UninstallBundle(opts UninstallOptions) error {
	err := p.prepullBundleByTag(&opts.BundleLifecycleOpts)
	if err != nil {
		return errors.Wrap(err, "unable to pull bundle before uninstall")
	}

	err = p.applyDefaultOptions(&opts.sharedOptions)
	if err != nil {
		return err
	}

	err = p.ensureLocalBundleIsUpToDate(opts.bundleFileOptions)
	if err != nil {
		return err
	}

	fmt.Fprintf(p.Out, "uninstalling %s...\n", opts.Name)
	return p.CNAB.Uninstall(opts.ToDuffleArgs())
}
