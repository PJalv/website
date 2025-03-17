package components

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type Interview struct {
	id            int
	Date          string
	Title         string
	RawTitle      string
	Description   string
	AudioFile     string
	Transcription string
	File          string // Store the file path
}

var InterviewsList = []Interview{
	{
		id:            0,
		Date:          time.Date(2024, 2, 19, 0, 0, 0, 0, time.UTC).Format("Jan 2, 2006"),
		Title:         "Sample Interview",
		RawTitle:      "Sample Interview",
		Description:   "A sample interview description",
		AudioFile:     "/file/interviews/sample.mp3",
		Transcription: "This is a sample transcription of the interview...",
		File:          "",
	},
}

func InterviewsConvert() {
	dir := "./data/interviews"
	temp := []Interview{}
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			interviewToAdd := convertInterviewMarkdownToHTML(path)
			if interviewToAdd.id != -99 {
				temp = append(temp, interviewToAdd)
				InterviewsList = temp
			}
		}
		return nil
	})

	if err != nil {
		log.Printf("Error walking through directory %s: %v\n", dir, err)
		return
	}
	log.Println("sorting interviews by date...")

	const dateFormat = "Jan 2, 2006"

	type sortableInterview struct {
		interview Interview
		time      time.Time
	}

	sortableInterviews := make([]sortableInterview, len(InterviewsList))

	for i, interview := range InterviewsList {
		parsedDate, err := time.Parse(dateFormat, interview.Date)
		if err != nil {
			fmt.Println("Error parsing date:", err)
			continue
		}
		sortableInterviews[i] = sortableInterview{interview, parsedDate}
	}

	sort.Slice(sortableInterviews, func(i, j int) bool {
		return sortableInterviews[i].time.Before(sortableInterviews[j].time)
	})

	for i, si := range sortableInterviews {
		InterviewsList[i] = si.interview
	}
}

func formatTranscriptionWithHTML(transcription string) string {
	lines := strings.Split(transcription, "\n")
	var formattedLines []string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// Check if line starts with a speaker (contains a colon)
		if idx := strings.Index(line, ":"); idx > 0 {
			speaker := line[:idx]
			content := strings.TrimSpace(line[idx+1:])

			// Format with different styles for interviewer and interviewee
			if strings.ToLower(speaker) == "interviewer" {
				formattedLines = append(formattedLines, fmt.Sprintf(
					`<div class="mb-4">
						<span class="font-semibold text-blue-600 dark:text-blue-400">%s:</span>
						<span class="ml-2 text-gray-700 dark:text-gray-300">%s</span>
					</div>`,
					speaker, content))
			} else {
				formattedLines = append(formattedLines, fmt.Sprintf(
					`<div class="mb-4 ml-4">
						<span class="font-semibold text-green-600 dark:text-green-400">%s:</span>
						<span class="ml-2 text-gray-700 dark:text-gray-300">%s</span>
					</div>`,
					speaker, content))
			}
		} else {
			// For lines without a speaker (actions or notes)
			formattedLines = append(formattedLines, fmt.Sprintf(
				`<div class="mb-4 italic text-gray-600 dark:text-gray-400">%s</div>`,
				line))
		}
	}

	return strings.Join(formattedLines, "\n")
}

func convertInterviewMarkdownToHTML(inputFile string) Interview {
	filename := filepath.Base(inputFile)

	parts := strings.Split(filename, "_")
	if len(parts) < 4 {
		log.Printf("Invalid filename format: %s\n", filename)
		return Interview{
			id: -99,
		}
	}
	dateStr := fmt.Sprintf("%s_%s_%s", parts[1], parts[2], parts[3])
	date, err := time.Parse("01_02_2006", dateStr)
	if err != nil {
		log.Printf("Error parsing date from filename: %v\n", err)
		return Interview{
			id: -99,
		}
	}

	input, err := os.ReadFile(inputFile)
	if err != nil {
		log.Printf("Error reading file %s: %v\n", inputFile, err)
		return Interview{
			id: -99,
		}
	}

	// Split content into sections (you might want to define your own format)
	sections := strings.Split(string(input), "---TRANSCRIPTION---")
	if len(sections) != 2 {
		log.Printf("Invalid interview file format: %s\n", filename)
		return Interview{
			id: -99,
		}
	}

	metadata := sections[0]
	transcription := formatTranscriptionWithHTML(sections[1])

	// Extract title from the first line
	lines := strings.Split(metadata, "\n")
	if len(lines) < 1 {
		log.Printf("Invalid interview file format: %s\n", filename)
		return Interview{
			id: -99,
		}
	}

	title := strings.TrimSpace(lines[0])
	description := ""
	audioFile := ""

	// Parse metadata
	for _, line := range lines[1:] {
		if strings.HasPrefix(line, "Description: ") {
			description = strings.TrimPrefix(line, "Description: ")
		} else if strings.HasPrefix(line, "AudioFile: ") {
			audioFile = strings.TrimPrefix(line, "AudioFile: ")
		}
	}

	dateFormatted := date.Format("Jan 2, 2006")

	interview := Interview{
		id:            len(InterviewsList),
		Date:          dateFormatted,
		RawTitle:      title,
		Title:         strings.ReplaceAll(strings.ToLower(title), " ", "-"),
		Description:   description,
		AudioFile:     audioFile,
		Transcription: transcription,
		File:          filename,
	}

	return interview
}

