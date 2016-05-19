package patcher

type StartingVersions struct {
	Versions []struct {
		Version    int
		Ref        string
		Submodules map[string]string
	} `yaml:"starting_versions"`
}

type Checkpoint struct {
	Changes     []Changeset
	CheckoutRef string
	FinalBranch string
}

type Changeset struct {
	Patches          []string
	Bumps            map[string]string
	SubmodulePatches map[string][]string
}

type patchSet interface {
	VersionsToApplyFor(version string) ([]Version, error)
	PatchesFor(Version) (patches []string, err error)
	BumpsFor(Version) (bumps map[string]string, err error)
	SubmodulePatchesFor(Version) (submodulePatches map[string][]string, err error)
}

type VersionsParser struct {
	version  string
	patchSet patchSet
}

func NewVersionsParser(version string, patchSet patchSet) VersionsParser {
	return VersionsParser{
		version:  version,
		patchSet: patchSet,
	}
}

func (p VersionsParser) GetCheckpoint() (Checkpoint, error) {
	var checkpoint Checkpoint

	versionsToApply, err := p.patchSet.VersionsToApplyFor(p.version)
	if err != nil {
		return Checkpoint{}, err
	}

	for _, version := range versionsToApply {
		patches, err := p.patchSet.PatchesFor(version)
		if err != nil {
			return Checkpoint{}, err
		}

		bumps, err := p.patchSet.BumpsFor(version)
		if err != nil {
			return Checkpoint{}, err
		}

		submodulePatches, err := p.patchSet.SubmodulePatchesFor(version)
		if err != nil {
			return Checkpoint{}, err
		}

		checkpoint.Changes = append(checkpoint.Changes, Changeset{
			Patches:          patches,
			Bumps:            bumps,
			SubmodulePatches: submodulePatches,
		})
	}

	checkpoint.CheckoutRef = versionsToApply[0].Ref
	checkpoint.FinalBranch = p.version

	return checkpoint, nil
}