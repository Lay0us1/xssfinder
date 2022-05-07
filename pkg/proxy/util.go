package proxy

import (
	"path/filepath"

	"github.com/gokitx/pkgs/slicex"
)

var (
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types/Common_types
	exts = []string{
		".aac",
		".abw",
		".arc",
		".avif",
		".avi",
		".azw",
		".bin",
		".bmp",
		".bz",
		".bz2",
		".cda",
		".csh",
		".css",
		".csv",
		".doc",
		".docx",
		".eot",
		".epub",
		".gz",
		".gif",
		".ico",
		".ics",
		".jpeg",
		".jpg",
		".js",
		".json",
		".jsonld",
		".mid",
		".midi",
		".mjs",
		".mp3",
		".mp4",
		".mpeg",
		".mpkg",
		".odp",
		".ods",
		".odt",
		".oga",
		".ogv",
		".ogx",
		".opus",
		".otf",
		".png",
		".pdf",
		".ppt",
		".pptx",
		".rar",
		".rtf",
		".sh",
		".svg",
		".swf",
		".tar",
		".tif ",
		".tiff",
		".ts",
		".ttf",
		".txt",
		".vsd",
		".wav",
		".weba",
		".webm",
		".webp",
		".woff",
		".woff2",
		".xls",
		".xlsx",
		".xml",
		".xul",
		".zip",
		".3gp",
		".3g2",
		".7z",
	}
)

func ignoreRequestWithPath(path string) bool {
	return slicex.ContainsIn(exts,
		filepath.Ext(path))
}
