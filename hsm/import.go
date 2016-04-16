package hsm

import (
	"os"

	"github.intel.com/hpdd/lustre"
	"github.intel.com/hpdd/lustre/llapi"
)

func Import(f string, archive uint, fi os.FileInfo, layout *llapi.DataLayout) (*lustre.Fid, error) {
	return llapi.HsmImport(f, archive, fi, layout)
}
