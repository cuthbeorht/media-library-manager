package audiofiles

import (
	"database/sql"
	"log"
)

type AudioFile struct {
	Name string
	Size int
	Path string
}

type AudioFileEntity struct {
	Id   int
	File AudioFile
}

func NewAudioFile(name string, size int, path string) AudioFile {
	return AudioFile{
		Name: name,
		Size: size,
		Path: path,
	}
}

func NewAudioFileEntity(name string, size int, path string, databaseId int) AudioFileEntity {
	return AudioFileEntity{
		File: NewAudioFile(name, size, path),
		Id:   databaseId,
	}
}

func CreateTables(connection *sql.DB) {
	createAudioFileTable := `
		create table if not exists audio_files (
			id integer primary key autoincrement,
			name varchar(255),
			size integer,
			path varchar(255)
		)
	`

	_, err := connection.Exec(createAudioFileTable)
	if err != nil {
		log.Fatal("Could not create audio file table: ", err)
	}
}
