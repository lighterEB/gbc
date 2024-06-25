package api

import (
	"fmt"
	"gbc/backend/bridge"
	"strings"
)

type Product struct {
	Name string `json:"name"`
	Pic  string `json:"pic"`
}

type Products struct {
	Products []*Product `json:"products"`
}

// 获取展示数据
func QueryData() (*Products, error) {
	var products Products
	//s1, _ := os.Getwd()
	//s1, _ := filepath.Abs(picsPath)
	//fmt.Println(s1)
	//filePath := path.Join(s1, "assets", "resources", "pics")
	//fmt.Println(filePath)
	fDir := "resources/pics"
	rd, err := bridge.Fs.ReadDir(fDir)
	if err != nil {
		return nil, err
	}
	fmt.Println(rd)
	for _, fi := range rd {
		if !fi.IsDir() {
			fData := fDir + "/" + fi.Name()
			f, err := bridge.Fs.ReadFile(fData)
			if err != nil {
				return nil, err
			}

			product := Product{
				Name: strings.Split(fi.Name(), ".")[0],
				Pic:  string(f),
			}
			products.Products = append(products.Products, &product)
		}
	}
	//res, err := json.Marshal(products)
	//if err != nil {
	//	return response.NewErrorResponse[string](500, "Failed to marshal products.")
	//}
	return &products, nil
}
