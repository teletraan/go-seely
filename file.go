package seely

import (
	"context"
	"fmt"

	"github.com/syfun/go-graphql"
)

// FileService represents service about file.
type FileService service

// UploadImage upload image from remote url.
func (fs *FileService) UploadImage(ctx context.Context, selection string, f graphql.NamedReader) (*Photo, error) {
	uploadErr := errFactory("upload image error")
	query := fmt.Sprintf(uploadImageQuery, selection)
	resp, err := fs.client.graphqlClient.SingleUpload(ctx, query, "uploadImage", f)
	if err != nil {
		return nil, uploadErr(err)
	}
	uploadImage := UploadImage{&Photo{}}
	if err := resp.Guess("uploadImage", &uploadImage); err != nil {
		return nil, uploadErr(err)
	}
	return uploadImage.Image, nil
}
