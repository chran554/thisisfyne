package selfie

import (
	"image"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type ImageFile struct {
	Image *image.Image `json:"-"`

	FileName string `json:"file_name"`

	Info1 string `json:"info_1"`
	Info2 string `json:"info_2"`
	Info3 string `json:"info_3"`
	Info4 string `json:"info_4"`
	Info5 string `json:"info_5"`
}

func LoadImageFiles(path string) ([]*ImageFile, error) {
	var imageFiles []*ImageFile

	dirEntries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for i, entry := range dirEntries {
		fileInfo, err := entry.Info()
		if err != nil {
			return nil, err
		}
		if !fileInfo.IsDir() && strings.HasSuffix(fileInfo.Name(), ".jpg") {
			img, err := loadImage(filepath.Join(path, fileInfo.Name()))
			if err != nil {
				return nil, err
			}
			imageFile := &ImageFile{
				Image:    &img,
				FileName: fileInfo.Name(),
				Info1:    strconv.Itoa(int(fileInfo.Size())),
				Info2:    fileInfo.ModTime().String(),
				Info3:    strconv.Itoa(i),
				Info4:    path,
				Info5:    "",
			}
			imageFiles = append(imageFiles, imageFile)
		}
	}

	return imageFiles, nil
}

func loadImage(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	return img, err
}
