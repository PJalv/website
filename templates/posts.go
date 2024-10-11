package templates

import (
	"fmt"
	// "github.com/alecthomas/chroma/v2"
	"github.com/russross/blackfriday/v2"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type Post struct {
	id          int
	Date        string
	Title       string
	RawTitle    string
	Description string
	Content     string
	File        string // New field to store the file path
}

var Posts = []Post{
	{
		id:      0,
		Date:    time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC).Format("Jan 2, 2006"),
		Title:   "Test",
		Content: "Testing",
		File:    "",
	},
	{
		id:      1,
		Date:    time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC).Format("Jan 2, 2006"),
		Title:   "Test1",
		Content: "Testing",
		File:    "",
	},
	{
		id:      2,
		Date:    time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC).Format("Jan 2, 2006"),
		Title:   "Test2",
		Content: "Testing",
		File:    "",
	},
}

func MDConvert() {
	dir := "./data/blogs"
	temp := []Post{}
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			postToAdd := convertMarkdownToHTML(path)
			if postToAdd.id != -99 {
				temp = append(temp, postToAdd)
				Posts = temp
			}
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Error walking through directory %s: %v\n", dir, err)
	}
	log.Println("sorting by date...")

	const dateFormat = "Jan 2, 2006"

	type sortablePost struct {
		post Post
		time time.Time
	}

	sortablePosts := make([]sortablePost, len(Posts))

	for i, post := range Posts {
		parsedDate, err := time.Parse(dateFormat, post.Date)
		if err != nil {
			fmt.Println("Error parsing date:", err)
			continue
		}
		sortablePosts[i] = sortablePost{post, parsedDate}
	}

	sort.Slice(sortablePosts, func(i, j int) bool {
		return sortablePosts[i].time.Before(sortablePosts[j].time)
	})

	for i, sp := range sortablePosts {
		Posts[i] = sp.post
	}

}

func convertMarkdownToHTML(inputFile string) Post {
	filename := filepath.Base(inputFile)

	parts := strings.Split(filename, "_")
	if len(parts) < 4 {
		log.Printf("Invalid filename format: %s\n", filename)
		return Post{
			id: -99,
		}
	}
	dateStr := fmt.Sprintf("%s_%s_%s", parts[1], parts[2], parts[3]) // Construct date string
	date, err := time.Parse("01_02_2006", dateStr)
	if err != nil {
		log.Printf("Error parsing date from filename: %v\n", err)
		return Post{
			id: -99,
		}
	}

	input, err := os.ReadFile(inputFile)
	if err != nil {
		log.Printf("Error reading file %s: %v\n", inputFile, err)
		return Post{
			id: -99,
		}
	}
	newlineIndex := strings.Index(string(input), "\n")
	if newlineIndex == -1 {
		fmt.Println("File is empty or unable to find the first line")
		return Post{
			id: -99,
		}
	}
	firstLine := string(input[:newlineIndex])
	// Extract the content after the first line
	restOfFile := input[newlineIndex+1:]

	htmlContent := blackfriday.Run(restOfFile)

	// Extract title from the first <h1> tag
	title, rawTitle := extractTitleFromHTML(string(htmlContent))

	// Convert date to desired format
	dateFormatted := date.Format("Jan 2, 2006")

	post := Post{
		id:          len(Posts),
		Date:        dateFormatted,
		RawTitle:    rawTitle,
		Title:       title,
		Description: firstLine,
		Content:     string(htmlContent),
		File:        filename, // Use filename instead of the full path
	}
	log.Println(post.Description)
	Posts = append(Posts, post)

	fmt.Printf("Converted %s to HTML and added to Posts\n", inputFile)
	return post
}

func extractTitleFromHTML(htmlContent string) (string, string) {
	// Find the first occurrence of <h1> tag
	start := strings.Index(htmlContent, "<h1>")
	end := strings.Index(htmlContent, "</h1>")
	if start == -1 || end == -1 {
		// No <h1> tag found, return empty title
		return "", ""
	}

	// Extract the title text between <h1> and </h1> tags
	title := htmlContent[start+4 : end]
	rawTitle := htmlContent[start+4 : end]
	// Replace spaces with hyphens
	title = strings.ReplaceAll(title, " ", "-")
	title = strings.ReplaceAll(title, ":", "")

	return title, rawTitle
}

func extractTitle(filePath string) string {
	base := filepath.Base(filePath)
	ext := filepath.Ext(base)
	return base[:len(base)-len(ext)]
}

func changeFileExtension(filename, newExt string) string {
	ext := filepath.Ext(filename)
	return filename[:len(filename)-len(ext)] + newExt
}
