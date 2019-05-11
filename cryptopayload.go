// Package cryptopayload provides functions to prepare the payload
// and to retrieve the payload from within the custom client certificate.
// Nothing fancy for now, b64enc(compress(payload)) and reverse.
// @Sourcefrenchy
package cryptopayload

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io"
	"io/ioutil"
	"log"
)

// Prepare takes care of compressing and encoding a string (payload)
func Prepare(payloadDat string) string {
	var buf bytes.Buffer
	err := compress(&buf, []byte(payloadDat))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Payload  --> %.*s...\t(%d bytes)", 10, payloadDat, len(payloadDat))
	log.Printf("PayloadPrepared --> %.*s...\t\t(%d bytes)", 10, buf.String(), len(buf.String()))
	return base64.StdEncoding.EncodeToString([]byte(buf.String()))
}

// Retrieve takes care of retrieving and decoding a string (payload)
func Retrieve(payloadDat string) string {
	decoded, _ := base64.StdEncoding.DecodeString(payloadDat)
	var buf bytes.Buffer
	err := decompress(&buf, decoded)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("PayloadRetrieved -->\n%s<--", buf.String())
	return buf.String()
}

func compress(w io.Writer, data []byte) error {
	gw, err := gzip.NewWriterLevel(w, gzip.BestCompression)
	defer gw.Close()
	gw.Write(data)
	return err
}

func decompress(w io.Writer, data []byte) error {
	gr, err := gzip.NewReader(bytes.NewBuffer(data))
	defer gr.Close()
	data, err = ioutil.ReadAll(gr)
	if err != nil {
		return err
	}
	w.Write(data)
	return nil
}
