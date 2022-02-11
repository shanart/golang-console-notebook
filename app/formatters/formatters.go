package formatters

import (
	"database/sql"
	"notebook/app/config"
	"notebook/app/models"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

/*
	Format raw query results to Notes array
*/
func FromDBToNotesList(db *gorm.DB, list *sql.Rows) []models.Note {
	var res []models.Note

	for list.Next() {
		var note models.Note
		db.ScanRows(list, &note)
		i := models.Note{
			ID:        note.ID,
			Content:   note.Content,
			CreatedAt: note.CreatedAt,
			UpdatedAt: note.UpdatedAt,
		}

		res = append(res, i)
	}

	return res
}

/*
	Print Notes array
*/
func ListNotesItems(notesList []models.Note) []string {
	// result := []string{"ID\tContent\t\t\tCreatedAt\t\tUpdatedAt"}
	result := []string{}
	for _, item := range notesList {
		str := strconv.FormatUint(uint64(item.ID), 10) + "\t"
		str += TruncateText(item.Content, 18) + "\t"
		str += item.CreatedAt.Format(config.GetConfig().TimeFormat) + "\t"
		// str += item.UpdatedAt.Format(config.GetConfig().TimeFormat) + "\t"
		result = append(result, str)
	}
	return result
}

/*
	Truncate string
*/
func TruncateText(text string, length uint8) string {
	if length <= 0 {
		return ""
	}
	truncated := ""
	var count uint8 = 0
	for _, char := range text {
		truncated += string(char)
		count++
		if count >= length {
			break
		}
	}

	if len(text) > int(length) {
		truncated += "..."
	}

	return strings.ReplaceAll(truncated, "\n", "")
}
