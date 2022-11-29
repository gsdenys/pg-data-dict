package writer

import (
	"bytes"
	"io"
	"os"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/gsdenys/pdgen/pkg/models"
	"github.com/gsdenys/pdgen/pkg/services/translate"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/message"
)

var baseTest models.Describe = models.Describe{
	Database: models.Basic{
		Name: "postgres",
		Desc: "default database",
	},
	Schema: models.Basic{
		Name: "public",
		Desc: "default database",
	},
	Tables: []models.Table{
		{
			Name: "test",
			Desc: "somme test",
			Columns: []models.Columns{
				{
					Column:  "test",
					Type:    "text",
					Allow:   "",
					Comment: "nothing",
				},
			},
		},
	},
}

// func TestPrinterJson_Done(t *testing.T) {
// 	path := os.TempDir() + uuid.NewString()
// 	file, _ := os.Create(path)

// 	p := &PrinterJson{
// 		Out:       file,
// 		Translate: translate.GetTranslation("en"),
// 	}

// 	p.Done(baseTest)

// 	want := "{\n    \"database\": {\n        \"name\": \"postgres\",\n        \"description\": \"default database\"\n    },\n    \"schema\": {\n        \"name\": \"public\",\n        \"description\": \"default database\"\n    },\n    \"tables\": [\n        {\n            \"name\": \"test\",\n            \"description\": \"somme test\",\n            \"columns\": [\n                {\n                    \"column\": \"test\",\n                    \"type\": \"text\",\n                    \"allow\": \"\",\n                    \"comment\": \"nothing\"\n                }\n            ]\n        }\n    ]\n}"

// 	b, err := os.ReadFile(path)
// 	if err != nil {
// 		t.Error(err.Error())
// 	}

// 	assert.Equal(t, string(b), want)
// }

func TestPrinterJson_GetLanguage(t *testing.T) {
	type fields struct {
		Out       io.Writer
		Translate *message.Printer
	}
	tests := []struct {
		name   string
		fields fields
		want   *message.Printer
	}{
		{
			name: "en",
			fields: fields{
				Out:       bytes.NewBuffer([]byte{}),
				Translate: translate.GetTranslation("en"),
			},
			want: translate.GetTranslation("en"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PrinterJson{
				Out:       tt.fields.Out,
				Translate: tt.fields.Translate,
			}
			if got := p.GetLanguage(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PrinterJson.GetLanguage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrinterJson_Init(t *testing.T) {
	p := &PrinterJson{
		Out:       bytes.NewBuffer([]byte{}),
		Translate: translate.GetTranslation("en"),
	}
	p.Init(baseTest)
}

func TestPrinterJson_Title(t *testing.T) {
	p := &PrinterJson{
		Out:       bytes.NewBuffer([]byte{}),
		Translate: translate.GetTranslation("en"),
	}
	p.Title("test")
}

func TestPrinterJson_Subtitle(t *testing.T) {
	p := &PrinterJson{
		Out:       bytes.NewBuffer([]byte{}),
		Translate: translate.GetTranslation("en"),
	}
	p.Subtitle("test")
}

func TestPrinterJson_SubSubtitle(t *testing.T) {
	p := &PrinterJson{
		Out:       bytes.NewBuffer([]byte{}),
		Translate: translate.GetTranslation("en"),
	}
	p.SubSubtitle("test")
}

func TestPrinterJson_LineBreak(t *testing.T) {
	p := &PrinterJson{
		Out:       bytes.NewBuffer([]byte{}),
		Translate: translate.GetTranslation("en"),
	}
	p.LineBreak()
}

func TestPrinterJson_Body(t *testing.T) {
	p := &PrinterJson{
		Out:       bytes.NewBuffer([]byte{}),
		Translate: translate.GetTranslation("en"),
	}
	p.Body("test")
}

func TestPrinterJson_Columns(t *testing.T) {
	p := &PrinterJson{
		Out:       bytes.NewBuffer([]byte{}),
		Translate: translate.GetTranslation("en"),
	}
	p.Columns([]models.Columns{})
}

func TestPrinterJson_Table(t *testing.T) {
	p := &PrinterJson{
		Out:       bytes.NewBuffer([]byte{}),
		Translate: translate.GetTranslation("en"),
	}
	p.Table(models.Table{})
}

func TestPrinterJson_Done(t *testing.T) {
	createFile := func(path string) *os.File {
		f, err := os.Create(path)

		if err != nil {
			t.Error(err.Error())
		}

		return f
	}

	type fields struct {
		Path      string
		Translate *message.Printer
	}
	type args struct {
		desc models.Describe
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "successful",
			fields: fields{
				Path:      os.TempDir() + uuid.NewString(),
				Translate: translate.GetTranslation("en"),
			},
			args: args{
				desc: baseTest,
			},
			want: "{\n    \"database\": {\n        \"name\": \"postgres\",\n        \"description\": \"default database\"\n    },\n    \"schema\": {\n        \"name\": \"public\",\n        \"description\": \"default database\"\n    },\n    \"tables\": [\n        {\n            \"name\": \"test\",\n            \"description\": \"somme test\",\n            \"columns\": [\n                {\n                    \"column\": \"test\",\n                    \"type\": \"text\",\n                    \"allow\": \"\",\n                    \"comment\": \"nothing\"\n                }\n            ]\n        }\n    ]\n}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PrinterJson{
				Out:       createFile(tt.fields.Path),
				Translate: tt.fields.Translate,
			}
			p.Done(tt.args.desc)

			b, err := os.ReadFile(tt.fields.Path)
			if err != nil {
				t.Error(err.Error())
			}

			assert.Equal(t, string(b), tt.want)
		})
	}
}
