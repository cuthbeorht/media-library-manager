package audiofiles

import (
	"database/sql"
	"fmt"
	"log"
)

type AudioFile struct {
	Name string
	Size uint64
	Path string
}

type AudioFileEntity struct {
	Id   int64
	File AudioFile
}

func NewAudioFile(name string, size uint64, path string) AudioFile {
	return AudioFile{
		Name: name,
		Size: size,
		Path: path,
	}
}

func NewAudioFileEntity(name string, size uint64, path string, databaseId int64) AudioFileEntity {
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

func FindAudioFileByName(connection *sql.DB, name string) AudioFileEntity {

	preparedSelect, _ := connection.Prepare("select name from audio_files where name = ?")

	res, _ := preparedSelect.Query(name)
	res.Next()
	var fileName AudioFileEntity
	res.Scan(&fileName)
	res.Close()

	return fileName

}

func PersistAudioFile(connection *sql.DB, audioFile AudioFile) AudioFileEntity {
	tx, _ := connection.Begin()

	preparedInsert, _ := tx.Prepare("insert into audio_files (name, size, path) values (?, ?, ?)")
	defer preparedInsert.Close()

	res, err := preparedInsert.Exec(audioFile.Name, audioFile.Size, audioFile.Path)
	if err != nil {
		fmt.Println("could not insert: ", err)
	}

	var newAudioId int64
	newAudioId, _ = res.LastInsertId()
	tx.Commit()
	return AudioFileEntity{
		File: audioFile,
		Id:   newAudioId,
	}
}
