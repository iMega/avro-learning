package acceptance

import (
	"bytes"
	"fmt"

	"github.com/imega/avro-learning/nid"
	"github.com/linkedin/goavro"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Make record", func() {

	Context("Create record in mysql", func() {
		It("done", func() {
			err := db.Ping()
			Expect(err).NotTo(HaveOccurred())

			schema := `
				{
					"namespace": "ru.imega",
					"type": "record",
					"name": "person",
					"aliases": ["version", "1.0"],
					"doc": "personal data",
					"fields" : [
						{"name": "firstname", "type": "string"},
						{"name": "lastname", "type": "string"}
					]
				}
			`

			buf := new(bytes.Buffer)

			err = encode(schema, buf)
			Expect(err).NotTo(HaveOccurred())

			fmt.Println(buf)

			id := nid.NewNID()

			_, err = db.Exec(
				"insert entities (account_id, entity_id, entity_type, `schema`, `entity`) values (?, ?, ?, ?, ?)",
				1, id, "", []byte{'1'}, buf.Bytes(),
			)
			Expect(err).NotTo(HaveOccurred())
		})
	})

})

func encode(schema string, buf *bytes.Buffer) error {
	codec, err := goavro.NewCodec(schema)
	if err != nil {
		return err
	}

	textual := []byte(`{"firstname":"Имя", "lastname":"Фамилия"}`)

	native, _, err := codec.NativeFromTextual(textual)
	if err != nil {
		return err
	}

	w, err := goavro.NewOCFWriter(goavro.OCFConfig{
		W:               buf,
		Codec:           codec,
		Schema:          schema,
		CompressionName: goavro.CompressionDeflateLabel,
	})

	if err := w.Append([]interface{}{native}); err != nil {
		return err
	}

	return nil
}
