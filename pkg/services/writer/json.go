package writer

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/gsdenys/pdgen/pkg/models"
)

type JSON struct {
	Out io.Writer
}

func (p *JSON) SetWriter(path string) {
	p.Out = createFile(path)
}

func (p *JSON) Init(desc models.Describe) {
	// Do nothing because have nothing to initialise
}

func (p *JSON) Title(title string) {
	//Do nothing because the unique action of this writer is Done
}

func (p *JSON) Subtitle(subtitle string) {
	//Do nothing because the unique action of this writer is Done
}

func (p *JSON) SubSubtitle(subSubtitle string) {
	//Do nothing because the unique action of this writer is Done
}

func (p *JSON) LineBreak() {
	//Do nothing because the unique action of this writer is Done
}

func (p *JSON) Body(desc string) {
	//Do nothing because the unique action of this writer is Done
}

func (p *JSON) Columns(columns []models.Columns) {
	//Do nothing because the unique action of this writer is Done
}

func (p *JSON) Table(t models.Table) {
	//Do nothing because the unique action of this writer is Done
}

func (p *JSON) Done(desc models.Describe) {
	b, _ := json.MarshalIndent(desc, "", "    ")

	fmt.Fprintf(p.Out, "%s", string(b))

	_ = p.Out.(*os.File).Close()
}
