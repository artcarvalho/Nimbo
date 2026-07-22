package repository

import (
	"database/sql"
	"fmt"
	"nimbo-api/model"
)

type CaptureRepository struct {
	connection *sql.DB
}

func NewCaptureRepository(connection *sql.DB) CaptureRepository {
	return CaptureRepository{
		connection: connection,
	}
}

func (cr *CaptureRepository) GetCapture() ([]model.Capture, error) {

	query := "SELECT id, file_path, source, captured_at, created_at FROM captures"
	rows, err := cr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Capture{}, err
	}

	var captureList []model.Capture
	var captureObj model.Capture

	for rows.Next() {
		err = rows.Scan(
			&captureObj.ID,
			&captureObj.File_path,
			&captureObj.Source,
			&captureObj.Captured_at,
			&captureObj.Created_at)

		if err != nil {
			fmt.Println(err)
			return []model.Capture{}, err
		}

		captureList = append(captureList, captureObj)
	}

	rows.Close()

	return captureList, nil
}
