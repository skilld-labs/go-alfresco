package rest

import (
	"bytes"
	"mime/multipart"
	"net/http"
)

func (c *Client) UploadFile(params map[string]string, field string, fileContents []byte, filename string) error {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(field, filename)
	if err != nil {
		return err
	}
	part.Write(fileContents)
	for key, val := range params {
		err = writer.WriteField(key, val)
		if err != nil {
			return err
		}
	}
	err = writer.Close()
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", c.getUrl()+"/alfresco/s/api/upload", body)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())
	if _, _, err := c.doRequest(req, nil); err != nil {
		return err
	}
	return nil
}
