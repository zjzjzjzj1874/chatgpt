package image

import (
	"bytes"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/zjzjzjzj1874/chatgpt/pkg"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
)

func init() {
	vCmd.Flags().StringVarP(&image, "image", "i", "", "The image to edit. Must be a valid PNG file, less than 4MB, and square. If mask is not provided, image must have transparency, which will be used as the mask.")
	vCmd.Flags().IntVarP(&num, "num", "n", defaultNum, "The number of images to generate. Must be between 1 and 10")
	vCmd.Flags().StringVarP(&size, "size", "s", "256x256", "The size of the generated images. Must be one of 256x256, 512x512, or 1024x1024")
	vCmd.MarkFlagsRequiredTogether("image")
}

var (
	vCmd = &cobra.Command{
		Use:   "var",
		Short: "Creates a variation of a given image.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(image) == 0 {
				color.Red("%s", "Please input your image path with -i")
				return
			}
			if num < minNum || num > maxNum {
				color.Red("The number of images to generate. Must be between %s and %s", minNum, maxNum)
				return
			}
			if _, ok := imgGenSizeMap[size]; !ok {
				color.Red("%s", "The size of the generated images. Must be one of 256x256, 512x512, or 1024x1024")
				return
			}

			fi, err := os.Open(image)
			if err != nil {
				color.Red("Open image(%s) failure:%s", image, err.Error())
				return
			}

			body := new(bytes.Buffer)
			writer := multipart.NewWriter(body)
			part, err := writer.CreateFormFile("image", image)
			if err != nil {
				color.Red("CreateFormFile image(%s) failure:%s", image, err.Error())
				return
			}
			_, err = io.Copy(part, fi)
			if err != nil {
				color.Red("Copy image(%s) failure:%s", image, err.Error())
				return
			}

			if len(mask) != 0 {
				fi, err := os.Open(mask)
				if err != nil {
					color.Red("Open mask image(%s) failure:%s", image, err.Error())
					return
				}

				writer := multipart.NewWriter(body)
				part, err := writer.CreateFormFile("mask", mask)
				if err != nil {
					color.Red("CreateFormFile mask(%s) failure:%s", mask, err.Error())
					return
				}
				_, err = io.Copy(part, fi)
				if err != nil {
					color.Red("Copy mask image(%s) failure:%s", mask, err.Error())
					return
				}
			}
			_ = writer.WriteField("n", strconv.Itoa(num))
			_ = writer.WriteField("size", size)
			_ = writer.Close()

			var (
				resp pkg.ImageCreateResponse
			)
			client, err := pkg.NewClient(pkg.WithMethod(http.MethodPost), pkg.WithContentType(writer.FormDataContentType()), pkg.WithUrl(pkg.IMG_VAR_URL), pkg.WithBody(body))
			if err != nil {
				color.Red("New Client Err:%s", err.Error())
				return
			}

			err = client.Send(&resp)
			if err != nil {
				color.Red("Send Chat Err:%s", err.Error())
				return
			}

			color.Cyan("Total Image:%v", len(resp.Data))
			for _, item := range resp.Data {
				color.Cyan(item.URL)
			}
		},
	}
)
