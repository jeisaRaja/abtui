package scanner

var audioExtensions = map[string]bool{
	".mp3":  true,
	".m4b":  true,
	".m4a":  true,
	".aac":  true,
	".flac": true,
	".wav":  true,
}

type Audiobook struct {
	Name string
	Path string
}

func scanAudiobooks(baseDir string) ([]Audiobook, error) {
	return nil, nil
}
