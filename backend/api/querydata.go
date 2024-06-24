package api

import (
	"fmt"
	"os"
	"path"
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
	s1, _ := os.Getwd()
	//s1, _ := filepath.Abs(picsPath)
	fmt.Println(s1)
	filePath := path.Join(s1, "resources", "pics")
	fmt.Println(filePath)
	rd, err := os.ReadDir(filePath)
	if err != nil {
		return nil, err
	}
	for _, fi := range rd {
		if !fi.IsDir() {
			f, err := os.ReadFile(filePath + string(os.PathSeparator) + fi.Name())
			if err != nil {
				return nil, err
			}

			product := Product{
				Name: fi.Name(),
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
