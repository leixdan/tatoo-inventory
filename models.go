package main

import (
	"context"
)

type Material struct {
	ID          int
	Name        string
	Quantity    int
	Description string
}

func GetAllMaterials() ([]Material, error) {
	rows, err := db.QueryContext(context.Background(), "SELECT id, name, quantity, description FROM materials ORDER BY id DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	materials := []Material{}
	for rows.Next() {
		var m Material
		err := rows.Scan(&m.ID, &m.Name, &m.Quantity, &m.Description)
		if err != nil {
			return nil, err
		}
		materials = append(materials, m)
	}
	return materials, nil
}

func GetMaterialByID(id int) (Material, error) {
	var m Material
	err := db.QueryRowContext(context.Background(), "SELECT id, name, quantity, description FROM materials WHERE id=$1", id).Scan(&m.ID, &m.Name, &m.Quantity, &m.Description)
	return m, err
}

func CreateMaterial(m Material) error {
	_, err := db.ExecContext(context.Background(), "INSERT INTO materials (name, quantity, description) VALUES ($1, $2, $3)", m.Name, m.Quantity, m.Description)
	return err
}

func UpdateMaterial(m Material) error {
	_, err := db.ExecContext(context.Background(), "UPDATE materials SET name=$1, quantity=$2, description=$3 WHERE id=$4", m.Name, m.Quantity, m.Description, m.ID)
	return err
}

func DeleteMaterial(id int) error {
	_, err := db.ExecContext(context.Background(), "DELETE FROM materials WHERE id=$1", id)
	return err
}
